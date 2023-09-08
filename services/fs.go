package services

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ceelsoin/subzero/utils"
)

var videoExtensions = []string{
	".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm",
	".3gp", ".m4v", ".mpeg", ".mpg", ".ogv", ".ts", ".vob",
	".m2ts", ".mts", ".asf", ".divx", ".xvid", ".f4v",
	".mxf", ".rm", ".rmvb", ".dat", ".nut", ".h264",
	".h265", ".vp8", ".vp9", ".avchd", ".swf",
	".mp3", ".wav", ".flac", ".aac", ".ogg", ".wma", ".mka",
	".opus",
}

var logger *utils.Logger

func StartFileScan(directoryPath string) {
	logger = utils.GetLogger()
	var videoFiles []string

	utils.AsyncTaskLoading(func() {
		err := FindVideoFiles(directoryPath, &videoFiles)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to find for files: %s", err.Error()))
		}
	}, fmt.Sprintf("Scanning for video files in folder %s you can specify scan directory, run `subzero help` for details", directoryPath))

	logger.Debug(fmt.Sprintf("%b files found", len(videoFiles)))
	for _, file := range videoFiles {
		GetFileInfo(file, directoryPath)
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
