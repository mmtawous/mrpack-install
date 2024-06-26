package web

import (
	"fmt"
	"github.com/nothub/mrpack-install/buildinfo"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	c  http.Client
	ua string
}

var DefaultClient = NewClient()

func NewClient() *Client {
	c := &Client{c: http.Client{}}
	c.c.Transport = NewTransport()
	c.ua = UserAgent()
	return c
}

func NewTransport() *http.Transport {
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   20 * time.Second,
		ResponseHeaderTimeout: 25 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	}
}

func UserAgent() string {
	return fmt.Sprintf(
		"%s (+https://%s)",
		buildinfo.Name(),
		buildinfo.Module(),
	)
}

func (c *Client) SetProxy(fixedURL string) error {
	proxy, err := url.Parse(fixedURL)
	if err != nil {
		return err
	}

	transport := NewTransport()
	transport.Proxy = http.ProxyURL(proxy)
	c.c.Transport = transport

	// Test proxy
	httpUrl := "https://api.modrinth.com/"
	response, err := c.c.Get(httpUrl)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return err
	}

	return nil
}
