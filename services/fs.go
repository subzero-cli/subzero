package services

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	cmd_utils "github.com/subzero-cli/subzero/cmd/utils"
	"github.com/subzero-cli/subzero/domain"
	"github.com/subzero-cli/subzero/infra"
	"github.com/subzero-cli/subzero/utils"
)

var videoExtensions = []string{
	".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm",
	".3gp", ".m4v", ".mpeg", ".mpg", ".ogv", ".ts", ".vob",
	".m2ts", ".mts", ".asf", ".divx", ".xvid", ".f4v", ".h264",
	".h265", ".vp8", ".vp9", ".avchd", ".swf", ".opus", ".ogg",
	".mp3", ".wav", ".flac", ".aac",
}

var logger *utils.Logger

func StartFileScan(directoryPath string) {
	scanPath, err := utils.GetFullPath(directoryPath)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to find for files: %s", err.Error()))
	}
	logger = utils.GetLogger()
	database := infra.GetDatabaseInstance()

	c := infra.NewConfigurationInstance()
	cfg, err := c.GetConfig()
	if err != nil {
		logger.Error(fmt.Sprintf("Error loading config: %s", err.Error()))
	}

	var videoFiles []string

	utils.AsyncTaskLoading(func() {

		err = FindVideoFiles(scanPath, &videoFiles)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to find for files: %s", err.Error()))
		}
	}, fmt.Sprintf("Scanning for video files in folder %s you can specify scan directory, run `subzero help` for details", scanPath))

	logger.Info(fmt.Sprintf("%d files found", len(videoFiles)))

	var fileInfoList []domain.FileInfo
	for _, file := range videoFiles {
		fmt.Sprintln(file)
		fileInfo := GetFileInfo(file)
		fileInfoList = append(fileInfoList, fileInfo)
	}

	logger.Info(fmt.Sprintf("%d files updated to database", len(fileInfoList)))

	answer := cmd_utils.AskYesNo("Would automatic find and download subtitles for this files?")

	if answer {
		for _, fileInfo := range fileInfoList {
			if len(fileInfo.OpenSubtitlesHash) > 0 {
				subtitles, err := SearchByHash(fileInfo.OpenSubtitlesHash, cfg.OpenSubtitlesApiKey)
				if err != nil {
					logger.Error(err.Error())
				}
				logger.Info(fmt.Sprintf("Found %d subtitles for file %s", subtitles.TotalCount, fileInfo.FileName))
				fileInfo.Subtitles = subtitles
				database.Update(fileInfo.ID, fileInfo)

				for _, subtitle := range subtitles.Data {
					// logger.Info(fmt.Sprintf("Downloading subtitle for (%s) for language %s", fileInfo.FileName, subtitle.Attributes.Language))
					err := DownloadSubtitle(subtitle.ID, cfg.OpenSubtitlesApiKey, fileInfo.Path)
					if err != nil {
						fmt.Errorf(err.Error())
					}
				}
			}
		}
		logger.Info("Yeeeeeeeeeeeaaaaaaaaaaah! Subtitles download done. 🍿🎥")
	} else {
		logger.Info("Now, your database was updated, you can run `subzero download` any time to download subtitles to your files.")
	}

}

func FindVideoFiles(dirPath string, videoFiles *[]string) error {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())

		if file.IsDir() {
			err := FindVideoFiles(filePath, videoFiles)
			if err != nil {
				return err
			}
		} else {
			ext := strings.ToLower(filepath.Ext(filePath))
			for _, videoExt := range videoExtensions {
				if ext == videoExt {
					*videoFiles = append(*videoFiles, filePath)
					break
				}
			}
		}
	}

	return nil
}
