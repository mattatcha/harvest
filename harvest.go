package harvest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func NewHarvestClient(domain, username, password string) (*HarvestClient, error) {
	base := fmt.Sprint("https://", domain, ".harvestapp.com")

	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	userPass := fmt.Sprint(username, ":", password)
	encoded := base64.StdEncoding.EncodeToString([]byte(userPass))
	return &HarvestClient{
		encodedAuth: encoded,
		baseURL:     baseURL,
		client:      &http.Client{},
	}, nil

}

func (c *HarvestClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := *c.baseURL
	url.Path = path

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Accept":        {"application/json"},
		"Content-Type":  {"application/json"},
		"Authorization": {"Basic " + c.encodedAuth},
	}

	return req, nil
}

func (c *HarvestClient) do(request *http.Request) (*http.Response, error) {
	return c.client.Do(request)
}

func (c *HarvestClient) Daily() (*Daily, error) {
	req, err := c.NewRequest("GET", "/daily", nil)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	daily := &Daily{}
	err = json.Unmarshal(body, daily)
	if err != nil {
		return nil, err
	}

	return daily, nil

}
