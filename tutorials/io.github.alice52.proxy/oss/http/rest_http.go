package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

type loggingTransport struct {
	transport http.RoundTripper
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("Making request to", req.URL, req.Body)
	res, err := t.transport.RoundTrip(req)
	if err != nil {
		return res, err
	}
	data, _ := httputil.DumpResponse(res, true)
	fmt.Println("Response:", string(data))
	return res, nil
}

type CustomClient struct {
	*http.Client
}

func DoReq(client *http.Client, method, url string, body []byte) ([]byte, error) {

	if client == nil {
		client = &http.Client{
			Transport: &loggingTransport{http.DefaultTransport},
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error close response body")
		}
	}(resp.Body)

	respb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	return respb, nil
}
