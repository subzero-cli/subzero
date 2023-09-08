package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
)

func DownloadFile(url string, fileName string, outputPath string) error {
	_check := "\u2713"
	_green := "\x1b[32m"
	_reset := "\x1b[0m"
	_red := "\x1b[31m"
	_x := "✘"

	var check string = _green + _check + _reset
	var x string = _red + _x + _reset

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " └" + fmt.Sprintf("Downloading \"%s\"", fileName)
	s.Start()
	response, err := http.Get(url)
	if err != nil {
		logger.Error(fmt.Sprintf(" └ %s Error downloading \"%s\"", x, fileName))
		return err
	}
	defer response.Body.Close()

	fullPath := filepath.Join(outputPath, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		logger.Error(fmt.Sprintf(" └ %s Error downloading \"%s\"", x, fileName))
		return err
	}
	defer file.Close()

	var reader io.Reader = response.Body

	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}

	s.Stop()
	logger.Info(fmt.Sprintf(" └ %s Downloaded \"%s\"", check, fileName))
	return nil
}
