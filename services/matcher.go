package services

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/subzero-cli/subzero/domain"
	"github.com/subzero-cli/subzero/infra"
	"github.com/subzero-cli/subzero/utils"
)

func GetFileInfo(filePath string) domain.FileInfo {

	logger := utils.GetLogger()
	database := infra.GetDatabaseInstance()

	fileInfo := domain.FileInfo{}

	fileNameSplited := strings.Split(filePath, "/")
	filenameWithoutPath := strings.ToLower(strings.Trim(fileNameSplited[len(fileNameSplited)-1], " "))
	fileNameWithoutExt := removeExtension(filenameWithoutPath)

	sanitizedName := sanitizeName(fileNameWithoutExt)

	seasonEpisodeRegex := regexp.MustCompile(`[sS](\d{2})[eE](\d{2})`)
	yearRegex := regexp.MustCompile(`\b(\d{4})\b`)
	resolutionRegex := regexp.MustCompile(`\b(\d{3,4}p)\b`)
	codecRegex := regexp.MustCompile(`(?i)\b(?:` + strings.Join(knownCodecs, "|") + `)\b`)

	logger.Debug(fmt.Sprintf("Using path: %s", filePath))

	fileInfo.FileName = filenameWithoutPath
	fileInfo.Path = filepath.Dir(filePath)
	fileInfo.SanitizedName = sanitizedName

	// md5, err := utils.CalculateChecksum(filePath)
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("Erro ao calcular o hash md5: %s", err))
	// }

	match := seasonEpisodeRegex.FindStringSubmatch(sanitizedName)
	if match != nil && len(match) >= 3 {
		fileInfo.Season = match[1]
		fileInfo.Episode = match[2]
	}

	match = resolutionRegex.FindStringSubmatch(sanitizedName)
	if match != nil && len(match) >= 2 {
		fileInfo.Resolution = match[1]
	}

	matches := yearRegex.FindAllString(sanitizedName, -1)
	if matches != nil && len(matches) > 0 {
		fileInfo.Year = matches[0]
	}

	codecMatch := codecRegex.FindStringSubmatch(sanitizedName)
	if codecMatch != nil && len(codecMatch) >= 1 {
		fileInfo.Codec = codecMatch[0]
	}

	hash, err := utils.HashOpenSubtitles(filePath)
	if len(hash) > 0 {
		fileInfo.OpenSubtitlesHash = hash
	}

	fileInfo.ID = hash

	if err != nil {
		logger.Error(fmt.Sprintf("Erro ao calcular o hash: %s", err))
	}

	parts := strings.Split(sanitizedName, " ")
	var cleanParts []string
	for _, part := range parts {
		if !yearRegex.MatchString(part) && !resolutionRegex.MatchString(part) && !codecRegex.MatchString(part) {
			cleanParts = append(cleanParts, part)
		}
	}

	fileInfo.Title = strings.Join(cleanParts, " ")

	boldPrint := color.New(color.Bold).SprintfFunc()

	logger.Info(fmt.Sprintf("[%s] (%s) SEASON: %s, EPISODE: %s", fileInfo.OpenSubtitlesHash, boldPrint(filenameWithoutPath), fileInfo.Season, fileInfo.Episode))

	database.Create(fileInfo)

	return fileInfo
}

func removeExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func sanitizeName(name string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(name, " ")
}

var knownCodecs = []string{
	"H.261", "H.262", "H.263", "H.264", "H.265", "MPEG-1", "MPEG-2", "MPEG-4", "MJPEG", "VP8",
	"VP9", "VC-1", "DivX", "Xvid", "Theora", "WMV", "AVC", "HEVC", "AV1", "RV40", "RV30", "RV20",
	"Sorenson", "DV", "Indeo", "Cinepak", "FLV1", "Dirac", "SMPTE VC-1", "Apple ProRes", "DNxHD",
	"FFV1", "Ut Video", "Lagarith", "Huffyuv", "Cineform", "JPEG 2000", "Apple Intermediate Codec (AIC)",
	"GoPro CineForm", "HEIF", "AVS", "Flash Video", "RealVideo", "Windows Media Video", "QuickTime Animation (RLE)",
	"Cavs", "Daala", "Thor", "FieldPlus", "Surround Video", "VP10", "AVS2", "OMAF", "OMAF Stereo", "MPEG-H 3D Audio",
	"MPEG-H 3D Audio Baseline Profile", "HE-AAC", "AAC", "MP3", "Vorbis", "Opus", "AC3", "E-AC-3", "DTS", "DTS-HD",
	"PCM", "WMA", "FLAC", "ALAC", "AMR", "AMR-WB", "AMR-WB+", "EVRC", "EVRC-B", "EVRC-WB", "G.711", "G.722", "G.723",
	"G.728", "G.729", "G.729.1", "G.729E", "G.729I", "GSM-FR", "GSM-HR", "GSM-EFR", "GSM-AMR", "GSM-AMR-WB", "GSM-AMR-WB+",
	"ADPCM", "CELT", "SILK", "AMBE", "MELP", "SMV", "EVRC", "MS GSM", "MS ADPCM", "IMA ADPCM", "G.722.1", "G.722.1C",
	"Opus", "G.728", "iLBC", "Siren", "G.711", "G.711.1", "G.711.0", "AAC-LD", "AAC-ELD", "AAC-LC", "AMR-NB", "AMR-WB",
	"Speex", "Comfort noise", "DTS", "DTS-HD", "DTS-X", "AC3", "E-AC-3", "ATRAC", "ATRAC3", "E-AC-3", "SDDS", "ATRAC3plus",
	"G.722.2", "AC-4", "ILBC", "ADPCM", "G.722.1", "G.722.1C", "OPUS", "SIREN", "L8", "L16", "L24", "PCMA", "PCMU", "GSM",
	"SPEEX", "G729", "AACLD", "AMR", "AMR-WB", "SILK", "EVRC", "G.726", "G.722", "G.722.1C", "G.722.2", "MELP", "AMBE",
	"LPC", "CELT", "SBC", "MSBC", "ADPCM", "MP3", "DVI4", "L16", "G722", "G728", "G729", "G726", "ADPCMW", "GSM", "SLIN",
	"LPC", "SILK", "SPEEX", "ILBC", "G726-32", "OPUS", "EVS", "G.711", "G.722", "G.722.1", "G.722.1C", "G.722.2", "G.723.1",
	"G.726", "G.727", "G.728", "G.729", "G.729.1", "GSM EFR", "GSM AMR", "SPEEX", "SILK", "CELT", "AMR-WB", "AMR-WB+",
	"EVRC", "EVRC-B", "EVRC-WB", "MS GSM", "MS ADPCM", "GSM-EFR", "GSM-AMR", "GSM-AMR-WB", "GSM-AMR-WB+", "G.722.1",
	"G.722.1C", "iLBC", "Siren", "G.711", "G.711.1", "G.711.0", "AAC-LD", "AAC-ELD", "AAC-LC", "AMR-NB", "AMR-WB", "Speex",
	"Comfort noise", "DTS", "DTS-HD", "DTS-X", "AC3", "E-AC-3", "ATRAC", "ATRAC3", "E-AC-3", "SDDS", "ATRAC3plus", "G.722.2",
	"AC-4", "ILBC", "ADPCM", "G.722.1", "G.722.1C", "OPUS", "SIREN", "L8", "L16", "L24", "PCMA", "PCMU", "GSM", "SPEEX", "G729",
	"AACLD", "AMR", "AMR-WB", "SILK", "EVRC", "G.726", "G.722", "G.722.1C", "G.722.2", "MELP", "AMBE", "LPC", "CELT", "SBC", "MSBC",
	"ADPCM", "MP3", "DVI4", "L16", "G722", "G728", "G729", "G726", "ADPCMW", "GSM", "SLIN", "LPC", "SILK", "SPEEX", "ILBC", "G726-32",
	"OPUS", "EVS",
}
