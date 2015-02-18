package bingo

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
)

const URL = "https://api.datamarket.azure.com/Bing/Search/v1/"

var (
	client *http.Client = &http.Client{}
)

type Adult string

const (
	Moderate Adult = "Moderate"
	Strict         = "Strict"
	Off            = "Off"
)

type Request interface {
	Params() map[string]string
	Base() string
	AccountKey() string
}

func RequestUrl(r Request) (result string, err error) {
	searchUrl, err := url.Parse(URL + r.Base())

	if err != nil {
		return
	}
	query := searchUrl.Query()
	for key, value := range r.Params() {
		query.Add(key, "'"+value+"'")
	}
	searchUrl.RawQuery = query.Encode()
	result = searchUrl.String()
	return
}

func ExecuteRequest(r Request) (result string, err error) {

	searchUrl, err := RequestUrl(r)
	if err != nil {
		return
	}
	req, err := http.NewRequest("GET", searchUrl, nil)

	auth := base64.StdEncoding.EncodeToString([]byte(r.AccountKey() + ":" + r.AccountKey()))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	result = string(body)
	return
}
