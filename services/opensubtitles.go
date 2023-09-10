package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/subzero-cli/subzero/utils"
)

type SubtitleData struct {
	TotalPages int            `json:"total_pages"`
	TotalCount int            `json:"total_count"`
	PerPage    int            `json:"per_page"`
	Page       int            `json:"page"`
	Data       []SubtitleInfo `json:"data"`
}

type SubtitleInfo struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes SubtitleAttributes `json:"attributes"`
}

type SubtitleAttributes struct {
	SubtitleID        string        `json:"subtitle_id"`
	Language          string        `json:"language"`
	DownloadCount     int           `json:"download_count"`
	NewDownloadCount  int           `json:"new_download_count"`
	HearingImpaired   bool          `json:"hearing_impaired"`
	HD                bool          `json:"hd"`
	FPS               float64       `json:"fps"`
	Votes             int           `json:"votes"`
	Ratings           float64       `json:"ratings"`
	FromTrusted       bool          `json:"from_trusted"`
	ForeignPartsOnly  bool          `json:"foreign_parts_only"`
	UploadDate        string        `json:"upload_date"`
	AITranslated      bool          `json:"ai_translated"`
	MachineTranslated bool          `json:"machine_translated"`
	Release           string        `json:"release"`
	Comments          string        `json:"comments"`
	LegacySubtitleID  int           `json:"legacy_subtitle_id"`
	Uploader          UploaderInfo  `json:"uploader"`
	FeatureDetails    FeatureInfo   `json:"feature_details"`
	URL               string        `json:"url"`
	RelatedLinks      []RelatedLink `json:"related_links"`
	Files             []FileInfo    `json:"files"`
	MoviehashMatch    bool          `json:"moviehash_match"`
}

type UploaderInfo struct {
	UploaderID int    `json:"uploader_id"`
	Name       string `json:"name"`
	Rank       string `json:"rank"`
}

type FeatureInfo struct {
	FeatureID       int    `json:"feature_id"`
	FeatureType     string `json:"feature_type"`
	Year            int    `json:"year"`
	Title           string `json:"title"`
	MovieName       string `json:"movie_name"`
	IMDBID          int    `json:"imdb_id"`
	TMDBID          int    `json:"tmdb_id"`
	SeasonNumber    int    `json:"season_number"`
	EpisodeNumber   int    `json:"episode_number"`
	ParentIMDBID    int    `json:"parent_imdb_id"`
	ParentTitle     string `json:"parent_title"`
	ParentTMDBID    int    `json:"parent_tmdb_id"`
	ParentFeatureID int    `json:"parent_feature_id"`
}

type RelatedLink struct {
	Label  string `json:"label"`
	URL    string `json:"url"`
	ImgURL string `json:"img_url"`
}

type FileInfo struct {
	FileID   int    `json:"file_id"`
	CDNumber int    `json:"cd_number"`
	FileName string `json:"file_name"`
}

type SubtitleDownloadInfo struct {
	Link            string        `json:"link"`
	FileName        string        `json:"file_name"`
	Requests        int           `json:"requests"`
	Remaining       int           `json:"remaining"`
	Message         string        `json:"message"`
	ResetTime       string        `json:"reset_time"`
	ResetTimeUTC    string        `json:"reset_time_utc"`
	Uk              string        `json:"uk"`
	UID             int           `json:"uid"`
	Timestamp       int64         `json:"ts"`
	ResetTimeParsed time.Duration // Campo extra para o tempo de reset convertido em um formato Go
}

var userAgent string = "Subzero CLI over Go HTTP"

func SearchByHash(hash string, key string) (SubtitleData, error) {
	logger := utils.GetLogger()

	logger.Debug(fmt.Sprintf("Using opensubtitles.com api key: %s", key))

	url := fmt.Sprintf("https://api.opensubtitles.com/api/v1/subtitles?moviehash=%s", hash)

	logger.Debug(fmt.Sprintf("Requesting url: %s", url))
	time.Sleep(250 * time.Millisecond)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Api-Key", key)

	logger.Debug(fmt.Sprintf("%v", req.Header))

	res, err := http.DefaultClient.Do(req)

	logger.Debug(fmt.Sprintf("Subtitle search from opensubtitles.com API returned status code: %s", res.Status))

	if err != nil {
		return SubtitleData{}, err
	}

	defer res.Body.Close()

	var subtitles SubtitleData

	err = json.NewDecoder(res.Body).Decode(&subtitles)
	if err != nil {
		return SubtitleData{}, err
	}

	return subtitles, nil
}

func DownloadSubtitle(fileId string, key string, outputPath string) error {
	logger := utils.GetLogger()

	url := "https://api.opensubtitles.com/api/v1/download"

	payload := strings.NewReader(fmt.Sprintf("{\"file_id\": %s, \"force_download\": true}", fileId))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Api-Key", key)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }

	var download SubtitleDownloadInfo

	err = json.NewDecoder(res.Body).Decode(&download)
	if err != nil {
		return err
	}
	logger.Debug(fmt.Sprintf("(File ID: %s) Opensubtitles.com quota %d/%d message: %s", fileId, download.Requests, download.Remaining, download.Message))

	if len(download.FileName) > 0 {
		utils.DownloadFile(download.Link, download.FileName, outputPath)
	}

	return nil
}
