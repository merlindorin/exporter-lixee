package main

import (
	_ "embed"

	kongyaml "github.com/alecthomas/kong-yaml"

	"github.com/alecthomas/kong"
	"github.com/merlindorin/exporter-lixee/cmd/exporter-lixee/commads"
	c "github.com/merlindorin/go-shared/pkg/cmd"
)

const (
	name        = "lixee"
	description = "Exporter for Lixee"
)

//nolint:gochecknoglobals // these global variables exist to be overridden during build
var (
	license string

	version     = "dev"
	commit      = "dirty"
	date        = "latest"
	buildSource = "source"
)

func main() {
	cli := CMD{
		Commons: &c.Commons{
			Version: c.NewVersion(name, version, commit, buildSource, date),
			Licence: c.NewLicence(license),
		},
		Serve: &commads.Serve{},
	}

	ctx := kong.Parse(
		&cli,
		kong.Name(name),
		kong.Description(description),
		kong.UsageOnError(),
		kong.Configuration(kongyaml.Loader, "/etc/lixee/config.yaml", "~/.hoomy/lixee.yaml"),
	)

	ctx.FatalIfErrorf(ctx.Run(cli.Commons))
}

type CMD struct {
	*c.Commons
	Serve *commads.Serve `cmd:"serve"`
}
