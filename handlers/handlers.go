package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func mapToDict(m http.Header) *zerolog.Event {
	dict := zerolog.Dict()
	for k, v := range map[string][]string(m) {
		dict.Str(k, strings.Join(append(v[1:], v...), " "))
	}

	return dict
}

func GenericHandler(statusCode int, path string, contentType string, headers map[string]string, body string, debug bool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		if debug {
			b, _ := ioutil.ReadAll(req.Body)
			log.Info().Dict("headers", mapToDict(req.Header)).Str("body", string(b)).Msg("request received")
		}

		for k, v := range headers {
			rw.Header().Set(k, v)
		}
		rw.Header().Set("Content-Type", contentType)
		rw.WriteHeader(statusCode)
		switch contentType {
		case "application/json":
			raw := json.RawMessage(body)
			marshalled, err := raw.MarshalJSON()
			if err != nil {
				log.Fatal().Err(err)
			}
			rw.Write(marshalled)
		default:
			rw.Write([]byte(body))
		}
	}
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%s", b)
}
