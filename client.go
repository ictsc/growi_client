package growi_client

import (
	"encoding/json"
	"errors"
	"github.com/ictsc/growi_client/entity"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Client interface {
	Init() error
	GetSubordinatedPage(path string) ([]entity.SubordinatedPage, error)
	GetPage(path string) (*entity.Page, error)
}

type GrowiClientOption struct {
	URL         *url.URL
	Username    string
	Password    string
	AccessToken string
}

type GrowiClient struct {
	Jar    *cookiejar.Jar
	Option *GrowiClientOption
}

var client *http.Client

var _ Client = (*GrowiClient)(nil)

func (c *GrowiClient) Init() error {
	client = &http.Client{}
	client.Jar = c.Jar

	csrfToken, err := getCsrfToken(*c.Option.URL, *client)
	if err != nil {
		return errors.New("failed to get csrf token")
	}
	err = doLogin(*c.Option.URL, *client, c.Option.Username, c.Option.Password, csrfToken)
	if err != nil {
		return errors.New("failed to login")
	}

	return nil
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

// getCsrfToken は CSRF Token を取得する
func getCsrfToken(u url.URL, client http.Client) (string, error) {
	u.Path = "/login"

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	node, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	// html > body attr の中から csrfToken を取得
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "html" {
			for d := c.FirstChild; d != nil; d = d.NextSibling {
				if d.Data == "body" {
					for _, attr := range d.Attr {
						if attr.Key == "data-csrftoken" {
							return attr.Val, nil
						}
					}
				}
			}
		}
	}

	return "", nil
}

func doLogin(u url.URL, client http.Client, username string, password string, csrfToken string) error {
	u.Path = "/login"

	form := url.Values{}
	form.Add("loginForm[username]", username)
	form.Add("loginForm[password]", password)
	form.Add("_csrf", csrfToken)

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
