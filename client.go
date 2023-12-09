package http

import (
	"fmt"
	nethttp "net/http"
	"strings"
	"time"
)

// Client is an HTTP client.
type Client interface {
	// Do sends the HTTP request and obtains an HTTP response.
	//
	// `req` may be modified by `Do`.
	Do(req *nethttp.Request) (*nethttp.Response, error)
}

type customClient struct {
	userAgent string
	client    *nethttp.Client
}

func (cc *customClient) Do(req *nethttp.Request) (*nethttp.Response, error) {
	req.Header.Set("User-Agent", cc.userAgent)

	return cc.client.Do(req)
}

func NewClient(userAgent string) (Client, error) {
	const timeout = 15 * time.Second
	stdlibClient := &nethttp.Client{
		Timeout: timeout,
	}

	return NewClientFromStandardLibrary(userAgent, stdlibClient)
}

func NewClientFromStandardLibrary(userAgent string, stdlibClient *nethttp.Client) (Client, error) {
	if strings.TrimSpace(userAgent) == "" {
		return nil, fmt.Errorf("invalid user agent: %q", userAgent)
	}

	c := &customClient{
		userAgent: userAgent,
		client:    stdlibClient,
	}

	return c, nil
}
