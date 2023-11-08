package checker

import (
	"fmt"
	"github.com/alice52/proxy/common/constants"
	"net/http"
	"strings"
)

func CheckReqAndParseObjectName(req *http.Request) (string, error) {

	params := req.URL.Query()
	url := params.Get("url")
	token := params.Get("token")

	if url == "" {
		return "", fmt.Errorf("url is empty!")
	}
	if token == "" {
		return "", fmt.Errorf("token is empty!")
	}

	return strings.Split(url, constants.Endpoint+"/")[1], nil
}
