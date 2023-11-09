package util

import (
	"encoding/json"
	"fmt"
	"github.com/alice52/proxy/common/model"
	"net/http"
)

func RespondError(code int, w http.ResponseWriter, err error) bool {
	err = json.NewEncoder(w).Encode(&model.R{
		ErrMsg:  err.Error(),
		ErrCode: code,
	})

	if err != nil {
		fmt.Println("Error responding error response")
	}

	return true
}
