package growi_client

import (
	"errors"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type GrowiClientOption struct {
	URL      *url.URL
	Username string
	Password string
}

type GrowiClient struct {
	Option *GrowiClientOption
}

var client *http.Client

func (c *GrowiClient) Init() (*GrowiClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client = &http.Client{}
	client.Jar = jar

	csrfToken, err := getCsrfToken(*c.Option.URL, *client)
	if err != nil {
		return nil, errors.New("failed to get csrf token")
	}
	err = doLogin(*c.Option.URL, *client, c.Option.Username, c.Option.Password, csrfToken)
	if err != nil {
		return nil, errors.New("failed to login")
	}

	return &GrowiClient{Option: c.Option}, nil
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
