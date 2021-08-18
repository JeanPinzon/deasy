package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

type API struct {
	Name string
	Slug string
}

var api API

var apisAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add REST APIs",
	Long:  `Scaffold a file structure for a new REST API`,
	Run: func(cmd *cobra.Command, args []string) {
		templateDir := "templates/apis/rest-memory"

		templates := []string{
			"api.yml",
			"Dockerfile",
			"k8s.yaml",
			"Makefile",
			"README.md",
			".github/workflows/push.yml",
		}

		os.Mkdir(api.Slug, 0777)
		os.Mkdir(fmt.Sprintf("%s/.github", api.Slug), 0777)
		os.Mkdir(fmt.Sprintf("%s/.github/workflows", api.Slug), 0777)

		for _, fileName := range templates {
			templateContent, err := fs.ReadFile(ApiTemplateFs, fmt.Sprintf("%s/%s", templateDir, fileName))

			if err != nil {
				log.Panic(err)
			}

			t := template.New(fileName)

			t.Delims("<<", ">>")

			t.Parse(string(templateContent))

			f, err := os.Create(fmt.Sprintf("%s/%s", api.Slug, fileName))

			if err != nil {
				log.Panic(err)
			}

			err = t.Execute(f, api)

			if err != nil {
				log.Panic(err)
			}
		}
	},
}

func init() {
	apisAddCmd.Flags().StringVarP(&api.Name, "name", "", "", "API name (required)")
	apisAddCmd.MarkFlagRequired("name")

	apisAddCmd.Flags().StringVarP(&api.Slug, "slug", "", "", "API slug (required)")
	apisAddCmd.MarkFlagRequired("slug")

	apisCmd.AddCommand(apisAddCmd)
}
