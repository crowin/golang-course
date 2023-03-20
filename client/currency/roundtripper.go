package currency

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type authRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
	apiKey string
}

func (a *authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	params := r.URL.Query()
	if a.apiKey == "" {
		panic(errors.New("apiKey is missed"))
	}
	params.Add("apikey", defaultApiKey)
	r.URL.RawQuery = params.Encode()
	fmt.Fprintf(a.logger, r.URL.String()+"\n")
	return a.next.RoundTrip(r)
}
