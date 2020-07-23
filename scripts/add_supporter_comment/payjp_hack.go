package main

import (
	"net/http"
	"strings"

	"github.com/payjp/payjp-go/v1"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

type uglyRoundTripper struct {
	underlying http.RoundTripper
}

func buildPayjpService(config *config.Config) *payjp.Service {
	httpClient := &http.Client{
		Transport: &uglyRoundTripper{
			underlying: http.DefaultTransport.(*http.Transport).Clone(),
		},
	}
	return payjp.New(config.Payjp.SecretKey, httpClient)
}

func (rt uglyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.Path, "/tokens") {
		req.Header.Add("X-Payjp-Direct-Token-Generate", "true")
	}

	return rt.underlying.RoundTrip(req)
}
