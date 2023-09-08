package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
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

func GetFileMd5Hash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:16]
	md5Hash := hex.EncodeToString(hashInBytes)

	return md5Hash, nil
}

func CalculateChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}

func GetCRC32(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := crc32.NewIEEE()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	checksum := hash.Sum32()

	return fmt.Sprintf("%d", checksum), nil
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
