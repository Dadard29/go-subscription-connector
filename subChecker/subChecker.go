package subChecker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var checkRoute = "/subs"

type checkResponse struct {
	Status bool
	Message string
	Content Content
}

type Content struct {
	AccessToken string `json:"AccessToken"`
	Profile     struct {
		ProfileKey  string    `json:"ProfileKey"`
		Username    string    `json:"Username"`
		DateCreated time.Time `json:"DateCreated"`
	} `json:"Profile"`
	Api struct {
		Name             string    `json:"Name"`
		DocumentationURL string    `json:"DocumentationUrl"`
		Hostname         string    `json:"Hostname"`
		VCSURL           string    `json:"VCSUrl"`
		BuildURL         string    `json:"BuildUrl"`
		IconURL          string    `json:"IconUrl"`
		Image            string    `json:"Image"`
		CreationDate     time.Time `json:"CreationDate"`
		IsStandard       bool      `json:"IsStandard"`
	} `json:"Api"`
	DateSubscribed time.Time `json:"DateSubscribed"`
}

type SubChecker struct {
	// the host of the core
	host string

	client *http.Client
}

func NewSubChecker(host string) *SubChecker {
	return &SubChecker{
		host: host,
		client: &http.Client{},
	}
}

func (s *SubChecker) CheckToken(token string, apiName string) (string, error) {
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf("%s%s", s.host, checkRoute), nil)

	if err != nil {
		return "error creating request", err
	}

	q, _ := url.ParseQuery(req.URL.RawQuery)
	q.Add("accessToken", token)
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return "error processing request", err
	}

	var r checkResponse
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &r)
	if err != nil {
		return "failed to deserialize response body", err
	}

	if r.Status {
		rApiName := r.Content.Api.Name
		if rApiName != apiName {
			msg := "bad api name"
			return msg, errors.New(msg)
		} else {
			return "sub checked", nil
		}
	}

	return r.Message, errors.New(r.Message)

}
