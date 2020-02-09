package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/hryniuk/moocist/pkg/coursera"
	"github.com/hryniuk/moocist/pkg/mooc"
)

var (
	courseraSlug string
	quiet        bool
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import MOOC into JSON file",
	Long:  `Import MOOC into JSON file`,
	Run: func(cmd *cobra.Command, args []string) {
		var importer mooc.Importer

		if len(courseraSlug) > 0 {
			importer = &coursera.SlugImporter{Slug: courseraSlug}
		}

		courseSyllabus, err := importer.Import()
		if err != nil {
			log.Errorf("cannot import course syllabus: %s", err)
			os.Exit(1)
		}

		exporter := mooc.JsonExporter{}
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
	rootCmd.AddCommand(importCmd)

	importCmd.PersistentFlags().StringVar(&courseraSlug, "coursera-slug", "", "Get Coursera MOOC with a given slug (e.g. creative-writing)")
	importCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "Quiet run (no output)")
}
