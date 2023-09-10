package cmd

import (
	"github.com/spf13/cobra"
	"github.com/subzero-cli/subzero/services"
)

var dirFlag string

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan files for identify movies, series and tvshows",
	Run: func(cmd *cobra.Command, args []string) {
		services.StartFileScan(dirFlag)

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&dirFlag, "dir", "d", "./", "Directory to scan")
}
