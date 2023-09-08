package domain

type FileInfo struct {
	ID                string `yaml:"id"`
	Title             string `yaml:"title"`
	FileName          string `yaml:"fileName"`
	Path              string `yaml:"path"`
	Year              string `yaml:"year"`
	Season            string `yaml:"season"`
	Episode           string `yaml:"episode"`
	SanitizedName     string `yaml:"sanitizedName"`
	Resolution        string `yaml:"resolution"`
	OpenSubtitlesHash string `yaml:"openSubtitlesHash"`
	Codec             string `yaml:"codec"`
	Subtitles         any    `yaml:"subtitles"`
}

type SubtitleInfo struct {
	FileName string `yaml:"fileName"`
	Path     string `yaml:"path"`
	Provider string `yaml:"provider"`
}
