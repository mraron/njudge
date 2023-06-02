package templates

import "embed"

//go:embed all:*.gohtml all:*/*
var FS embed.FS
