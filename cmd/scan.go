package cmd

import (
	"github.com/spf13/cobra"
	"github.com/subzero-cli/subzero/services"
)

var directory string
var defaultDirectory = "./"

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan files for identify movies, series and tvshows",
	Run: func(cmd *cobra.Command, args []string) {
		services.StartFileScan(directory)

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.PersistentFlags().StringVarP(&directory, "dir", "d", defaultDirectory, "Directory path to scan")

}
