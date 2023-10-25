package google

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const BaseURL = "https://youtube.googleapis.com/youtube/v3"

type YoutubeAPI interface {
	SearchVideos(query string) (Results, error)
}

type youtubeAPI struct {
	apiKey string
}

func NewYoutubeAPI(apiKey string) YoutubeAPI {
	return &youtubeAPI{
		apiKey: apiKey,
	}
}

type Results struct {
	Info  PageInfo `json:"pageInfo"`
	Items []Item   `json:"items"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type Item struct {
	ID ItemID `json:"id"`
}

type ItemID struct {
	VideoID string `json:"videoId"`
}

type Snippet struct {
	Title        string `json:"title"`
	ChannelTitle string `json:"channelTitle"`
}

func (y *youtubeAPI) SearchVideos(query string) (Results, error) {
	params := url.Values{}
	params.Set("key", y.apiKey)
	params.Set("part", "snippet")
	params.Set("q", query)

	u, _ := url.ParseRequestURI(BaseURL + "/search?")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, u.String()+params.Encode(), nil)
	if err != nil {
		return Results{}, nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return Results{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Results{}, errors.New("fail to fetch youtube data")
	}

	var res Results
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Results{}, nil
	}

	return res, nil
}
