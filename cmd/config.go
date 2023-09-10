package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/subzero-cli/subzero/domain"
	"github.com/subzero-cli/subzero/infra"
	"github.com/subzero-cli/subzero/utils"
)

var checkboxList []string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure your experience using subzero cli",

	Run: func(cmd *cobra.Command, args []string) {
		prompt := &survey.MultiSelect{
			Message: "Select the subtitle providers that you wanna enable:",
			Options: []string{"OpenSubtitles.com"},
		}
		err := survey.AskOne(prompt, &checkboxList)
		if err != nil {
			log.Fatal(err)
		}

		if utils.Contains(checkboxList, "OpenSubtitles.com") {
			var apiKey string
			promptAPIKey := &survey.Input{
				Message: "Please paste your OpenSubtitles.org API key:",
			}
			err := survey.AskOne(promptAPIKey, &apiKey, survey.WithValidator(survey.Required))
			if err != nil {
				log.Fatal(err)
			}

			// Agora você pode usar a apiKey como necessário
			fmt.Println("API Key:", apiKey)

			c := infra.NewConfigurationInstance()

			cfg := domain.Configuration{}
			cfg.EnableOpenSubtitles = true
			cfg.OpenSubtitlesApiKey = apiKey

			err = c.SaveConfig(cfg)
			if err != nil {
				logger.Error(fmt.Sprintf("Error while saving configuration: %s", err.Error()))
			}
		}

		logger.Info("Configuration finished, welcome to Subzero CLI, you can reconfigure running `subzero config`")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
