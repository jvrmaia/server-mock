package config

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestNoRoutesConfig(t *testing.T) {
	var server Routes

	if _, err := toml.Decode(noRoutesConfigBlob, &server); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, server.Debug, true, "debug must be enable")
	assert.Equal(t, server.Listen, "0.0.0.0:8080", "listen must be 0.0.0.0:8080")
	assert.Equal(t, server.WriteTimeout.Seconds(), float64(15), "write timeout must be 15s")
	assert.Equal(t, server.ReadTimeout.Seconds(), float64(15), "read timeout must be 15s")
	assert.Equal(t, server.IdleTimeout.Seconds(), float64(120), "idle timeout must be 2m")
}

func TestGenericRoutesConfig(t *testing.T) {
	var server Routes

	if _, err := toml.Decode(genericRouteConfigBlob, &server); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, server.Debug, true, "debug must be enable")
	assert.Equal(t, server.Listen, "0.0.0.0:8080", "listen must be 0.0.0.0:8080")
	assert.Equal(t, server.WriteTimeout.Seconds(), float64(15), "write timeout must be 15s")
	assert.Equal(t, server.ReadTimeout.Seconds(), float64(15), "read timeout must be 15s")
	assert.Equal(t, server.IdleTimeout.Seconds(), float64(120), "idle timeout must be 2m")

	assert.Equal(t, 1, len(server.Routes), "must be one route")

	r := server.Routes[0]
	assert.Equal(t, r.Type, "generic", "route type must be generic")
	assert.Equal(t, r.Path, "/", "path must be root")
	assert.Equal(t, r.StatusCode, 200, "status_code must be 200")
	assert.Equal(t, r.Headers.ToMap(), map[string]string{"Content-Type": "text/plain"}, "header must contain content-type")
	assert.Equal(t, r.Body, "oi", "body must be 'oi'")
}

func TestEchoRoutesConfig(t *testing.T) {
	var server Routes

	if _, err := toml.Decode(echoRouteConfigBlob, &server); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, server.Debug, true, "debug must be enable")
	assert.Equal(t, server.Listen, "0.0.0.0:8080", "listen must be 0.0.0.0:8080")
	assert.Equal(t, server.WriteTimeout.Seconds(), float64(15), "write timeout must be 15s")
	assert.Equal(t, server.ReadTimeout.Seconds(), float64(15), "read timeout must be 15s")
	assert.Equal(t, server.IdleTimeout.Seconds(), float64(120), "idle timeout must be 2m")

	assert.Equal(t, 1, len(server.Routes), "must be one route")

	r := server.Routes[0]
	assert.Equal(t, r.Type, "echo", "route type must be echo")
	assert.Equal(t, r.Path, "/echo", "path must be root")
	assert.Equal(t, r.StatusCode, 200, "status_code must be 200")
	assert.Equal(t, r.Headers.ToMap(), map[string]string{"Content-Type": "text/plain"}, "header must contain content-type")
}

func TestMultipleRoutesConfig(t *testing.T) {
	var server Routes

	if _, err := toml.Decode(multipleRoutesConfigBlob, &server); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 5, len(server.Routes), "must have five routes")
}
