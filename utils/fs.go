package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFullPath(relativePath string) (string, error) {
	absPath := relativePath

	if !filepath.IsAbs(absPath) {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Erro ao obter o diretório de trabalho:", err)
			return "", err
		}
		absPath = filepath.Join(wd, relativePath)
	}

	return absPath, nil
}
