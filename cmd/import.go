package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import MOOC into JSON file",
	Long:  `Import MOOC into JSON file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.PersistentFlags().String("--coursera-slug", "", "Get Coursera MOOC with a given slug (e.g. creative-writing)")
}
