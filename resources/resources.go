package resources

import (
	"embed"
)

//go:embed templates/*
var resTemplates embed.FS

//go:embed messages/*
var resMessages embed.FS

//go:embed web/*
var resWeb embed.FS

//go:embed keys/*
var resKeys embed.FS

//go:embed images/*
var resImages embed.FS

//go:embed data/*
var resData embed.FS

//go:embed release/*
var resRelease embed.FS

//go:embed build/*
var resBuildInfo embed.FS
