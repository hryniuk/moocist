package cmd

import (
	"fmt"
	"os"
	"strconv"

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

		if priority < 1 || priority > 4 {
			log.Errorf("priority should have integer value between 1 and 4, got %d", priority)
			os.Exit(1)
		}
		priorityString := strconv.Itoa(priority)

		exporter := mooc.CSVExporter{TaskPriority: priorityString}
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
	templateCmd.PersistentFlags().IntVar(&priority, "priority", 1, "Tasks priority (1/2/3/4)")
}
