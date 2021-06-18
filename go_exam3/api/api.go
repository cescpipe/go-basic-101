package api

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"time"
)

type Photo struct {
	AlbumId      int    `json:"albumId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	ApiUrl       string
	RoutineId    string
	Error        error
}

func CallPhotoApi(url string, routineId string, photoChannel chan Photo) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("error pre-call JSON API %s", err)
		photoChannel <- Photo{
			ApiUrl:    url,
			RoutineId: routineId,
			Error:     err,
		}
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("error call JSON API %s", err)
		photoChannel <- Photo{
			ApiUrl:    url,
			RoutineId: routineId,
			Error:     err,
		}
	}

	defer resp.Body.Close()

	bArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error read response from JSON API %s", err)
		photoChannel <- Photo{
			ApiUrl:    url,
			RoutineId: routineId,
			Error:     err,
		}
	}

	var photo Photo

	if err := json.Unmarshal(bArr, &photo); err != nil {
		log.Errorf("error unmarshall response from JSON API %s", err)
		photoChannel <- Photo{
			ApiUrl:    url,
			RoutineId: routineId,
			Error:     err,
		}
	}
	photo.RoutineId = routineId
	photo.ApiUrl = url
	photoChannel <- photo
}
