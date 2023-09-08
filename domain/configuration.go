package domain

type Configuration struct {
	EnableOpenSubtitles bool   `yaml:"enableOpenSubtitles"`
	OpenSubtitlesApiKey string `yaml:"openSubtitlesApiKey"`
}
