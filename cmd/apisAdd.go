package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type API struct {
	Name string
	Slug string
}

var apisAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add REST APIs",
	Long:  `Scaffold a file structure for a new REST API`,
	Run: func(cmd *cobra.Command, args []string) {
		api := API{Name: "New REST API", Slug: "new-rest-api"}
		templateDir := "templates/apis/rest-memory"

		os.Mkdir(api.Slug, 0777)

		templates, err := fs.ReadDir(TemplateFs, templateDir)

		if err != nil {
			log.Panic(err)
		}

		for _, template := range templates {
			content, err := fs.ReadFile(TemplateFs, fmt.Sprintf("%s/%s", templateDir, template.Name()))

			if err != nil {
				log.Panic(err)
			}

			os.WriteFile(fmt.Sprintf("./%s/%s", api.Slug, template.Name()), content, 0666)
		}
	},
}

func init() {
	apisCmd.AddCommand(apisAddCmd)
}
