package hpbclient

import (
	"fmt"
	"github.com/hpb-project/chain-watching/types"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	client     *http.Client
	serverAddr string
	rpcUrl     string
}

func NewClient(rpcUrl string, serverAddr string) *Client {
	return &Client{
		client:     &http.Client{},
		rpcUrl:     rpcUrl,
		serverAddr: serverAddr,
	}
}

func (c *Client) Do(q RCPQuest) ([]byte, error) {
	req, _ := http.NewRequest("POST", c.rpcUrl, strings.NewReader(q.body()))
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

func (c *Client) PostStatus(status string) error {
	path := c.serverAddr + types.NODESTATUS
	req, _ := http.NewRequest("POST", path, strings.NewReader(status))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	_, err := c.client.Do(req)
	if err != nil {
		fmt.Println("post to server error:", err)
		return err
	}

	return nil
}

func (c *Client) PostPeerInfo(peers string) error {
	path := c.serverAddr + types.PEERLIST
	req, _ := http.NewRequest("POST", path, strings.NewReader(peers))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	_, err := c.client.Do(req)
	if err != nil {
		fmt.Println("post to server error:", err)
		return err
	}

	return nil
}
