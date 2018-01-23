package hyper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewClient() *Client {
	return &Client{
		httpClient:       &http.Client{},
		additionalHeader: http.Header{},
	}
}

type Client struct {
	httpClient       *http.Client
	additionalHeader http.Header
}

func (c *Client) AdditionalHeader() http.Header {
	return c.additionalHeader
}

func (c *Client) Fetch(url string) (Item, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Item{}, fmt.Errorf("create: %v", err)
	}
	req.Header.Set(HeaderContentType, ContentTypeHyperItem)
	for k, v := range c.additionalHeader {
		if k == HeaderContentType {
			continue
		}
		req.Header[k] = v
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Item{}, fmt.Errorf("do: %v", err)
	}
	defer resp.Body.Close()
	res := Item{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return Item{}, fmt.Errorf("decode: %v", err)
	}
	return res, nil
}
