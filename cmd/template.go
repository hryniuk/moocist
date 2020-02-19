package cmd

import (
	"fmt"
	"os"

	"github.com/hryniuk/moocist/pkg/coursera"
	"github.com/hryniuk/moocist/pkg/mooc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create a Todoist CSV template for given MOOC",
	Long:  `Create a Todoist CSV template for given MOOC`,
	Run: func(cmd *cobra.Command, args []string) {
		var importer mooc.Importer

		if len(courseraSlug) == 0 {
			log.Errorf("pass course slug")
			return
		}

		importer = &coursera.SlugImporter{Slug: courseraSlug}

		courseSyllabus, err := importer.Import()
		if err != nil {
			log.Errorf("cannot import course syllabus: %s", err)
			os.Exit(1)
		}

		exporter := mooc.CSVExporter{}
		jsonBytes, err := exporter.Export(courseSyllabus)
		if err != nil {
			log.Errorf("cannot marshal course syllabus as JSON: %s", err)
			os.Exit(1)
		}

		if quiet {
			os.Exit(0)
		}

		fmt.Println(string(jsonBytes))
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.PersistentFlags().StringVar(&courseraSlug, "coursera-slug", "", "Get Coursera MOOC with a given slug (e.g. creative-writing)")
	templateCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "Quiet run (no output)")
}
