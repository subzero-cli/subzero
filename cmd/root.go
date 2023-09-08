/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/ceelsoin/subzero/utils"

	"github.com/spf13/cobra"
)

var verbose bool
var logger *utils.Logger

var rootCmd = &cobra.Command{
	Use:   "subzero",
	Short: "❄️ Manage and download subtitles for movies, series and tv shows. Made over by the best subtitle databases and so much coffee",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = utils.NewLogger(verbose)
		logger.Info("❄️ Welcome to subzero CLI. The most cold subtitles downloader and manager.")
		logger.Debug("Starting in verbose mode, a lot of text saying bla bla bla will appear in your screen")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose log")
}
