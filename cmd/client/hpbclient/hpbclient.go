package hpbclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	client *http.Client
	url    string
}

func NewClient(url string) *Client {
	return &Client{
		client: &http.Client{},
		url:    url,
	}
}

func (c *Client) Do(q RCPQuest) ([]byte, error) {
	req, _ := http.NewRequest("POST", c.url, strings.NewReader(q.body()))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Http Send HPB Error:", err)
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
