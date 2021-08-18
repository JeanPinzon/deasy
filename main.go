package main

import (
	"embed"

	"github.com/jeanpinzon/deasy/cmd"
)

//go:embed templates
var templateFs embed.FS

func main() {
	cmd.TemplateFs = templateFs
	cmd.Execute()
}
