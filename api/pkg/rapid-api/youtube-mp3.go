package rapidapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	BaseURL = "https://youtube-mp36.p.rapidapi.com/dl"
	APIHost = "youtube-mp36.p.rapidapi.com"
)

type RapidAPI interface {
	DownloadYoutubeMP3(videoID string) (MP3Data, error)
}

type rapidAPI struct {
	apiKey string
}

func NewRapidAPI(apiKey string) RapidAPI {
	return &rapidAPI{
		apiKey: apiKey,
	}
}

type MP3Data struct {
	Link     string  `json:"link"`
	Title    string  `json:"title"`
	Duration float32 `json:"duration"`
}

type ErrorResponse struct {
	Message      string `json:"msg"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Error: %v - %v", e.Message, e.ErrorMessage)
}

func (r *rapidAPI) DownloadYoutubeMP3(videoID string) (MP3Data, error) {
	url := fmt.Sprintf("%v?id=%v", BaseURL, videoID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return MP3Data{}, err
	}

	req.Header.Add("X-RapidAPI-Key", r.apiKey)
	req.Header.Add("X-RapidAPI-Host", APIHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return MP3Data{}, err
	}

	if res.StatusCode != http.StatusOK {
		return MP3Data{}, errors.New("fail to download track")
	}

	defer res.Body.Close()
	var data MP3Data
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return MP3Data{}, err
	}

	return data, nil
}
