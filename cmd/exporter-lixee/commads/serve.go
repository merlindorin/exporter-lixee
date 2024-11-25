package commads

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/exporter-toolkit/web"
	"go.uber.org/zap"

	"github.com/merlindorin/go-shared/pkg/buildinfo"
	"github.com/merlindorin/go-shared/pkg/cmd"
	u "github.com/merlindorin/go-shared/pkg/net/url"
	"github.com/merlindorin/go-shared/pkg/zapadapter"

	"github.com/merlindorin/exporter-lixee/frontend"
	"github.com/merlindorin/exporter-lixee/internal"
)

type Serve struct {
	MQTTHost           *url.URL `env:"MQTT_HOST" help:"Host of the MQTT Server" default:"mqtt://localhost:1883" required:""`
	MQTTTopic          string   `env:"MQTT_TOPIC" help:"MQTT topic to subscribe Lixee event" default:"zigbee2mqtt/LiXee"  required:""`
	MQTTClientID       string   `env:"MQTT_CLIENT_ID" help:"MQTT Client ID use for subscription" default:"exporter-lixee"`
	MQTTGracefulPeriod uint     `name:"MQTT_GRACEFUL_PERIOD" help:"Graceful period for disconnecting when the application is turning off." default:"250"`

	Timeout time.Duration `default:"5s" help:"Max duration for collecting data"`

	Web struct {
		ExternalURL     string   `name:"external-url" help:"The URL under which the exporter is externally reachable (for example, if the exporter is served via a reverse proxy). Used for generating relative and absolute links back to the exporter itself. If the URL has a path portion, it will be used to prefix all HTTP endpoints served by the exporter. If omitted, relevant URL components will be derived automatically."`
		RoutePrefix     *string  `name:"route-prefix" help:"Prefix for the internal routes of web endpoints. Defaults to path of --web.external-url."`
		SystemdSocket   bool     `name:"systemd-socket" help:"Use systemd socket activation listeners instead of port listeners (Linux only)."`
		ListenAddresses []string `name:"listen-addresses" default:":9090" help:"Addresses on which to expose metrics and web interface. Repeatable for multiple addresses."`
		ConfigFile      string   `name:"config.file" help:"Path to configuration file that can enable TLS or authentication. See: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md"`
	} `embed:"" prefix:"web."`

	MetricsReadTimeout       time.Duration `default:"1s"`
	MetricsWriteTimeout      time.Duration `default:"1s"`
	MetricsIdleTimeout       time.Duration `default:"30s"`
	MetricsReadHeaderTimeout time.Duration `default:"2s"`

	externalURL *url.URL
}

func (s *Serve) GetExternalURL() *url.URL {
	if s.externalURL == nil {
		eurl, err := u.ComputeExternalURL(s.Web.ExternalURL, (s.Web.ListenAddresses)[0])
		if err != nil {
			panic(fmt.Errorf("cannot compute external url based on the listening address: %w", err))
		}

		s.externalURL = eurl
	}

	return s.externalURL
}

func (s *Serve) prefixRoute(routes ...string) string {
	var prefixedRoute string
	var err error

	if s.Web.RoutePrefix != nil {
		prefixedRoute, err = url.JoinPath(*s.Web.RoutePrefix, routes...)
	} else {
		prefixedRoute, err = url.JoinPath(s.GetExternalURL().Path, routes...)
	}

	if err != nil {
		prefixedRoute = ""
	}

	prefixedRoute = "/" + strings.Trim(prefixedRoute, "/")

	if prefixedRoute != "/" {
		prefixedRoute += "/"
	}

	return prefixedRoute
}

func (s Serve) Run(common *cmd.Commons) error {
	logger := common.MustLogger()

	logger.Info(
		"Starting server...",
		zap.String("name", common.Version.Name()),
		zap.String("version", common.Version.Version()),
		zap.String("commit", common.Version.Commit()),
		zap.String("date", common.Version.Date()),
		zap.String("build-source", common.Version.BuildSource()),
	)

	reg := prometheus.NewRegistry()
	logger.Debug("Registering Lixee collector", zap.Duration("timeout", s.Timeout))

	collector := internal.NewCollector(s.Timeout, true)
	reg.MustRegister(collector)

	logger.Debug("Registering common build info collector")
	reg.MustRegister(buildinfo.NewCollector(common.Version.BuildInfo))

	logger.Debug("Registering golang build info collector")
	reg.MustRegister(collectors.NewBuildInfoCollector())

	logger.Debug("Registering golang collector")
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	if s.Web.ExternalURL == "" && s.Web.SystemdSocket {
		return fmt.Errorf(
			"cannot automatically infer external URL with systemd socket listener. Please provide --web.external-url",
		)
	}

	if s.prefixRoute() != "/" {
		http.HandleFunc("/", redirectOverExternalURL(s))
	}

	http.Handle(s.prefixRoute("/metrics"), promhttp.HandlerFor(reg, promhttp.HandlerOpts{EnableOpenMetrics: true}))
	http.HandleFunc(s.prefixRoute("-", "healthy"), AlwaysHealthy(logger))
	http.HandleFunc(s.prefixRoute(), StatusPage())

	lixeeState := &internal.LixeeState{}

	http.HandleFunc("/api/v1/lixee", lixeeAPIHandler(lixeeState))

	srv := &http.Server{
		ReadTimeout:       s.MetricsReadTimeout,
		WriteTimeout:      s.MetricsWriteTimeout,
		IdleTimeout:       s.MetricsIdleTimeout,
		ReadHeaderTimeout: s.MetricsReadHeaderTimeout,
	}

	srvc := make(chan error)
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	go func() {
		flagConfig := &web.FlagConfig{
			WebListenAddresses: &s.Web.ListenAddresses,
			WebSystemdSocket:   &s.Web.SystemdSocket,
			WebConfigFile:      &s.Web.ConfigFile,
		}

		logger.Info(fmt.Sprintf("HTTP Server started on %s", s.GetExternalURL()))
		if er := web.ListenAndServe(srv, flagConfig, zapadapter.ZapAdapter("HTTP server", logger)); er != nil {
			defer close(srvc)
			srvc <- er
		}
	}()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(s.MQTTHost.String())
	opts.SetClientID(s.MQTTClientID)

	client := mqtt.NewClient(opts)
	defer client.Disconnect(s.MQTTGracefulPeriod)
	defer client.Unsubscribe(s.MQTTTopic)

	listenerLogger := logger.With(
		zap.String("host", s.MQTTHost.String()),
		zap.String("clientId", s.MQTTClientID),
		zap.String("topic", s.MQTTTopic),
	)
	go lixeeListener(listenerLogger, client, srvc, s, lixeeState)

	for {
		select {
		case <-term:
			logger.Info("Received SIGTERM, exiting gracefully...")
			return nil
		case er := <-srvc:
			return fmt.Errorf("unexpected end: %w", er)
		}
	}
}

func lixeeListener(l *zap.Logger, cl mqtt.Client, srvc chan error, s Serve, state *internal.LixeeState) chan error {
	l.Info("MQTT Client connecting")
	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		defer close(srvc)
		srvc <- token.Error()
	}

	l.Info("MQTT Client subscribing...")
	if t := cl.Subscribe(s.MQTTTopic, 0, MessageHandler(state, l, srvc)); t.Wait() && t.Error() != nil {
		defer close(srvc)
		srvc <- t.Error()
	}
	return srvc
}

func lixeeAPIHandler(state *internal.LixeeState) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, _ *http.Request) {
		if err := json.NewEncoder(writer).Encode(state); err != nil {
			fmt.Fprintf(writer, "error encoding lixee state: %v", err)
		}
	}
}

func MessageHandler(lixeeState *internal.LixeeState, logger *zap.Logger, srvc chan error) mqtt.MessageHandler {
	return func(_ mqtt.Client, message mqtt.Message) {
		logger.Debug("new message received", zap.ByteString("payload", message.Payload()))
		err := json.Unmarshal(message.Payload(), lixeeState)
		if err != nil {
			defer close(srvc)
			srvc <- err
		}
	}
}

func redirectOverExternalURL(s Serve) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, s.GetExternalURL().String(), http.StatusFound)
	}
}

func StatusPage() func(http.ResponseWriter, *http.Request) {
	dist, _ := fs.Sub(frontend.Dist, frontend.Path)
	return http.FileServer(http.FS(dist)).ServeHTTP
}

func AlwaysHealthy(logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Healthy"))

		if err != nil {
			logger.Error("cannot write", zap.Error(err))
		}
	}
}
