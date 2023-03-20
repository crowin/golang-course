package coincap

import (
	"fmt"
	"io"
	"net/http"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l *loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "url: %s, method: %s\n", r.URL.String(), r.Method)
	return l.next.RoundTrip(r)
}

type bearerRoundTripper struct {
	next http.RoundTripper
}

func (l *bearerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Test-Auth", "Bearer 123")
	fmt.Println("Current headers: ", r.Header)
	return l.next.RoundTrip(r)
}
