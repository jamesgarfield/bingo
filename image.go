package bingo

import (
	"encoding/json"
)

//go:generate goast write impl baserequest.go

type ImageRequest struct {
	Query      string
	Adult      Adult
	accountKey string
}

func (imr ImageRequest) Base() string {
	return "Image"
}

func (imr ImageRequest) Params() map[string]string {
	return imr.params()
}

func (imr ImageRequest) Do() (results []ImageResult, err error) {
	raw, err := executeRequest(imr)
	if err != nil {
		return
	}

	var resp ImageResponse
	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {
		return
	}
	results = resp.Data.Results
	return
}

type ImageResponse struct {
	Data ImageResponseData `json:"d"`
}

type ImageResponseData struct {
	Results []ImageResult `json:"results"`
}

type ImageResult struct {
	ID          string
	Title       string
	MediaUrl    string
	SourceUrl   string
	DisplayUrl  string
	Width       string
	Height      string
	FileSize    string
	ContentType string
}
