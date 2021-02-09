package client

import (
	"net/http"
	"time"
)

type Params struct {
	MaxIdleConns,
	MaxConnsPerHost,
	MaxIdleConnsPerHost int
	Timeout int
}

func Default() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	const timeout time.Duration = 10

	return &http.Client{
		Timeout:   timeout * time.Second,
		Transport: t,
	}
}

func WithParams(p Params) *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = p.MaxIdleConns
	t.MaxConnsPerHost = p.MaxConnsPerHost
	t.MaxIdleConnsPerHost = p.MaxIdleConns

	return &http.Client{
		Timeout:   time.Duration(p.Timeout) * time.Second,
		Transport: t,
	}
}
