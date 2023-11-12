package growi_client

import (
	"encoding/json"
	"github.com/ictsc/growi_client/entity"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client interface {
	GetSubordinatedPage(path string) ([]entity.SubordinatedPage, error)
	GetPage(path string) (*entity.Page, error)
}

type GrowiClientOption struct {
	URL         *url.URL
	AccessToken string
}

type GrowiClient struct {
	Jar    *cookiejar.Jar
	Option *GrowiClientOption
}

var client *http.Client

var _ Client = (*GrowiClient)(nil)

func NewGrowiClient(option *GrowiClientOption) *GrowiClient {
	client = &http.Client{}

	return &GrowiClient{
		Option: option,
	}
}

type SubordinatedPagesResponse struct {
	SubordinatedPages []entity.SubordinatedPage `json:"subordinatedPages"`
}

func (c *GrowiClient) GetSubordinatedPage(path string) ([]entity.SubordinatedPage, error) {
	u := *c.Option.URL
	u.Path = "_api/v3/pages/subordinated-list"

	q := u.Query()
	q.Set("access_token", c.Option.AccessToken)
	q.Set("path", path)

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var subordinatedPagesResponse SubordinatedPagesResponse
	err = json.Unmarshal(body, &subordinatedPagesResponse)
	if err != nil {
		return nil, err
	}

	return subordinatedPagesResponse.SubordinatedPages, nil
}

type PageResponse struct {
	Page entity.Page `json:"page"`
}

func (c *GrowiClient) GetPage(path string) (*entity.Page, error) {
	u := *c.Option.URL
	u.Path = "_api/v3/page"

	q := u.Query()
	q.Set("access_token", c.Option.AccessToken)
	q.Set("path", path)

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pageResponse PageResponse
	err = json.Unmarshal(body, &pageResponse)
	if err != nil {
		return nil, err
	}

	return &pageResponse.Page, nil
}
