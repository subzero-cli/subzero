/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	scanCmd.Flags().BoolP("dir", "d", false, "Directory to scan")
	scanCmd.Flags().BoolP("verbose", "v", false, "Enable verbose log")
}
