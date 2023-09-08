package services

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchByHash(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	hash := "somehash"
	key := "someapikey"

	responseBody := `{"total_pages": 1, "total_count": 1, "per_page": 1, "page": 1, "data": []}`
	httpmock.RegisterResponder("GET", "https://api.opensubtitles.com/api/v1/subtitles?moviehash=somehash", httpmock.NewStringResponder(200, responseBody))

	subtitleData, err := SearchByHash(hash, key)

	assert.NoError(t, err)
	assert.Equal(t, 1, subtitleData.TotalPages)
	assert.Equal(t, 1, subtitleData.TotalCount)
}

func TestDownloadSubtitle(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	fileID := "somefileid"
	key := "someapikey"
	outputPath := "/some/output/path"

	responseBody := `{
		"link": "https://www.opensubtitles.com/download/A184A5EA6302F2CA7FD9D49BCEA49A1F36662BBEFB8C9B0ECDC9BB6CAF4BF09A5AA8D7B95C7FBD01615021D1973BAC18D431A8E6A1F627E4617341E8508A6968532088A68B6DDDA996C0116E2CE6F778ED9096A9CAB942B42B59C4EA93F1A7D61FCD6CBBC29C720EBD40CE674A55375862F00981E5D2F315A0982766A2004E0ED0AD9ADABEB506A638F1B829DBC2BE15979F22DA123523967F4D4069BC32098F1086F09AAA776CC365ED744633FD5FA7160B65A2C83539DF30134F5BE6272E46019AF9FD2423AFE12E1DC8642CDB56B8FEB9A4C1F30BF68EF431A3D4ABD3A7E44559E3E572210E5A5A33EC282D3445C537C5DA9DA598300A9900FA1B3B92983FD1504FDDFB34F89E409BF03EC662FC5734F25843C277A64B7C603156926FC6C74AA1D14AABEA6E20/subfile/castle.rock.s01e03.webrip.x264-tbs.ettv.-eng.ro.srt",
		"file_name": "castle.rock.s01e03.webrip.x264-tbs.ettv.-eng.ro.srt",
		"requests": 3,
		"remaining": 97,
		"message": "Your quota will be renewed in 07 hours and 30 minutes (2022-04-08 13:03:16 UTC) ",
		"reset_time": "07 hours and 30 minutes",
		"reset_time_utc": "2022-04-08T13:03:16.000Z"
	}`
	httpmock.RegisterResponder("POST", "https://api.opensubtitles.com/api/v1/download", httpmock.NewStringResponder(200, responseBody))

	err := DownloadSubtitle(fileID, key, outputPath)

	assert.NoError(t, err)
}