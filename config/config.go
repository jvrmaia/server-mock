package config

import (
	"encoding/json"
	"time"
)

type Route struct {
	Type        string
	Path        string
	StatusCode  int    `toml:"status_code"`
	ContentType string `toml:"content_type"`
	Headers     httpHeaders
	Body        string
}

type Routes struct {
	Debug        bool
	Listen       string
	WriteTimeout duration `toml:"write_timeout"`
	ReadTimeout  duration `toml:"read_timeout"`
	IdleTimeout  duration `toml:"idle_timeout"`
	Routes       []Route
}

type httpHeaders struct {
	header map[string]string
}

type duration struct {
	time.Duration
}

func (h *httpHeaders) UnmarshalText(text []byte) error {
	err := json.Unmarshal(text, &h.header)
	return err
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (h *httpHeaders) ToMap() map[string]string {
	return h.header
}
