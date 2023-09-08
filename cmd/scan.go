package cmd

import (
	"github.com/ceelsoin/subzero/services"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan files for identify movies, series and tvshows",
	Run: func(cmd *cobra.Command, args []string) {
		directoryPath := "./"
		services.StartFileScan(directoryPath)

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().BoolP("dir", "d", false, "Directory to scan")
}
