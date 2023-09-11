package services

import (
	"testing"
)

func TestGetFileInfo(t *testing.T) {
	fileInfo := GetFileInfo("1988_example_video_S01E01_720p_h.264.mp4")
	if fileInfo.Year != "1988" {
		t.Errorf("Expected empty year, got '%s'", fileInfo.Year)
	}

	if fileInfo.Season != "01" {
		t.Errorf("Expected season '01', got '%s'", fileInfo.Season)
	}

	if fileInfo.Episode != "01" {
		t.Errorf("Expected episode '01', got '%s'", fileInfo.Episode)
	}

	if fileInfo.SanitizedName != "1988 example video s01e01 720p h 264" {
		t.Errorf("Expected sanitized name 'example video S01E01 720p', got '%s'", fileInfo.SanitizedName)
	}

	if fileInfo.Resolution != "720p" {
		t.Errorf("Expected resolution '720p', got '%s'", fileInfo.Resolution)
	}

	if fileInfo.Codec != "h 264" {
		t.Errorf("Expected empty codec, got '%s'", fileInfo.Codec)
	}
}
