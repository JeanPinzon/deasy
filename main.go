package main

import (
	"embed"

	"github.com/jeanpinzon/deasy/cmd"
)

//go:embed templates/apis/rest-memory/*
var apiTemplateFs embed.FS

func main() {
	cmd.ApiTemplateFs = apiTemplateFs
	cmd.Execute()
}
