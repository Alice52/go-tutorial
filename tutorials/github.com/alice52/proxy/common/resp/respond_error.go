package resp

import (
	"encoding/json"
	"fmt"
	"github.com/alice52/proxy/oss/model"
	"net/http"
)

func RespondErrorWithStatus(status int, w http.ResponseWriter, err error) bool {
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(&model.ErrorResult{
		Msg: err.Error(),
	})

	if err != nil {
		fmt.Println("Error responding error response")
	}

	return true
}

func RespondError(w http.ResponseWriter, err error) bool {

	return RespondErrorWithStatus(http.StatusInternalServerError, w, err)
}