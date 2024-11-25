package frontend

import "embed"

const (
	Path = "dist"
)

var (
	//go:embed dist
	Dist embed.FS
)
