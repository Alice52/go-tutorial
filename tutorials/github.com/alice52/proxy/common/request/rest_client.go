package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type RestClient struct {
	*http.Client
}

func DoReq(client *http.Client, method, url string, body []byte) ([]byte, error) {

	if client == nil {
		client = &http.Client{
			Transport: &LogInterceptor{http.DefaultTransport},
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

	respb, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	return respb, nil
}
