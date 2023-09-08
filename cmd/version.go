package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version   string = "unset"
	buildTime string = "unset"
	commit    string = "unset"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of subzero",
	Long:  `All software has versions. This is subzero's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Build Time: %s\n", buildTime)
		fmt.Printf("Commit: %s\n", commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
