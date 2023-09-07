package utils

import (
	"encoding/binary"
	"fmt"
	"os"
)

func HashOpenSubtitles(videoPath string) (string, error) {
	file, err := os.Open(videoPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var fileSize int64
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize = fileInfo.Size()

	const (
		bufferSize = 65536
		bytesize   = 8 // Tamanho de int64 em bytes
	)

	var filehash uint64 = uint64(fileSize)
	if fileSize < bufferSize*2 {
		return "", nil
	}

	buffer := make([]byte, bytesize)
	for i := 0; i < bufferSize/bytesize; i++ {
		_, err := file.Read(buffer)
		if err != nil {
			return "", err
		}

		lValue := binary.LittleEndian.Uint64(buffer)
		filehash += lValue
		filehash &= 0xFFFFFFFFFFFFFFFF
	}

	_, err = file.Seek(max(0, fileSize-bufferSize), 0)
	if err != nil {
		return "", err
	}

	for i := 0; i < bufferSize/bytesize; i++ {
		_, err := file.Read(buffer)
		if err != nil {
			return "", err
		}

		lValue := binary.LittleEndian.Uint64(buffer)
		filehash += lValue
		filehash &= 0xFFFFFFFFFFFFFFFF
	}

	returnedhash := fmt.Sprintf("%016x", filehash)
	return returnedhash, nil
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}