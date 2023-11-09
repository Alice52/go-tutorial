package request

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type LogInterceptor struct {
	li http.RoundTripper
}

func (t LogInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("Making request to", req.URL, req.Body)
	res, err := t.li.RoundTrip(req)
	if err != nil {
		return res, err
	}
	data, _ := httputil.DumpResponse(res, true)
	fmt.Println("Response:", string(data))

	return res, nil
}
