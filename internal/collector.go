//nolint:lll
package internal

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	mutex        sync.Mutex
	state        *LixeeState
	timeout      time.Duration
	reportErrors bool

	// Metrics related to LixeeState
	infoGauge                 *prometheus.Desc
	apparentPowerGauge        *prometheus.Desc
	availablePowerGauge       *prometheus.Desc
	currentSummDeliveredGauge *prometheus.Desc
	currentTarifGauge         *prometheus.Desc
	linkQualityGauge          *prometheus.Desc
	meterSerialNumberGauge    *prometheus.Desc
	motDEtatGauge             *prometheus.Desc
	rmsCurrentGauge           *prometheus.Desc
	rmsCurrentMaxGauge        *prometheus.Desc
	warnDPSGauge              *prometheus.Desc
}

func (c *Collector) State() *LixeeState {
	return c.state
}

func (c *Collector) SetState(state *LixeeState) {
	c.mutex.Lock()
	*c.state = *state
	c.mutex.Unlock()
}

func NewCollector(timeout time.Duration, reportError bool) *Collector {
	return &Collector{
		reportErrors: reportError,
		timeout:      timeout,
		state:        &LixeeState{},

		infoGauge:                 prometheus.NewDesc("lixee_info", "Lixee info.", []string{"meter_serial_number", "active_register_tier_delivered", "current_tarif", "mot_d_etat"}, nil),
		apparentPowerGauge:        prometheus.NewDesc("lixee_apparent_power", "Apparent power (W)", []string{"meter_serial_number"}, nil),
		availablePowerGauge:       prometheus.NewDesc("lixee_available_power", "Available power (W)", []string{"meter_serial_number"}, nil),
		currentSummDeliveredGauge: prometheus.NewDesc("lixee_current_summ_delivered", "Current sum delivered (kWh)", []string{"meter_serial_number"}, nil),
		currentTarifGauge:         prometheus.NewDesc("lixee_current_tarif", "Current tariff applied", []string{"meter_serial_number"}, nil),
		linkQualityGauge:          prometheus.NewDesc("lixee_link_quality", "Link quality", []string{"meter_serial_number"}, nil),
		meterSerialNumberGauge:    prometheus.NewDesc("lixee_meter_serial_number", "Meter serial number", []string{"meter_serial_number"}, nil),
		motDEtatGauge:             prometheus.NewDesc("lixee_mot_d_etat", "State of the meter", []string{"meter_serial_number"}, nil),
		rmsCurrentGauge:           prometheus.NewDesc("lixee_rms_current", "RMS current (A)", []string{"meter_serial_number"}, nil),
		rmsCurrentMaxGauge:        prometheus.NewDesc("lixee_rms_current_max", "Max RMS current (A)", []string{"meter_serial_number"}, nil),
		warnDPSGauge:              prometheus.NewDesc("lixee_warn_dps", "Warning DPS status", []string{"meter_serial_number"}, nil),
	}
}

func (c *Collector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.infoGauge
	descs <- c.apparentPowerGauge
	descs <- c.availablePowerGauge
	descs <- c.currentSummDeliveredGauge
	descs <- c.currentTarifGauge
	descs <- c.linkQualityGauge
	descs <- c.meterSerialNumberGauge
	descs <- c.motDEtatGauge
	descs <- c.rmsCurrentGauge
	descs <- c.rmsCurrentMaxGauge
	descs <- c.warnDPSGauge
}
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	if c.state == nil {
		return
	}

	ch <- prometheus.MustNewConstMetric(c.infoGauge, prometheus.GaugeValue, 1, c.state.MeterSerialNumber, c.state.ActiveRegisterTierDelivered, c.state.CurrentTarif, c.state.MotDEtat)
	ch <- prometheus.MustNewConstMetric(c.apparentPowerGauge, prometheus.GaugeValue, float64(c.state.ApparentPower), c.state.MeterSerialNumber)
	ch <- prometheus.MustNewConstMetric(c.availablePowerGauge, prometheus.GaugeValue, float64(c.state.AvailablePower), c.state.MeterSerialNumber)
	ch <- prometheus.MustNewConstMetric(c.currentSummDeliveredGauge, prometheus.GaugeValue, float64(c.state.CurrentSummDelivered), c.state.MeterSerialNumber)
	ch <- prometheus.MustNewConstMetric(c.linkQualityGauge, prometheus.GaugeValue, float64(c.state.Linkquality), c.state.MeterSerialNumber)
	ch <- prometheus.MustNewConstMetric(c.meterSerialNumberGauge, prometheus.GaugeValue, 1, c.state.MeterSerialNumber)
	ch <- prometheus.MustNewConstMetric(c.rmsCurrentGauge, prometheus.GaugeValue, float64(c.state.RmsCurrent), c.state.MeterSerialNumber)
	ch <- prometheus.MustNewConstMetric(c.rmsCurrentMaxGauge, prometheus.GaugeValue, float64(c.state.RmsCurrentMax), c.state.MeterSerialNumber)
	ch <- prometheus.MustNewConstMetric(c.warnDPSGauge, prometheus.GaugeValue, float64(c.state.WarnDPS), c.state.MeterSerialNumber)
}
