package resp

import (
	"encoding/json"
	"fmt"
	"github.com/alice52/proxy/oss/model"
	"net/http"
)

func RespondError(w http.ResponseWriter, err error) bool {
	err = json.NewEncoder(w).Encode(&model.ErrorResult{
		Msg: err.Error(),
	})

	if err != nil {
		fmt.Println("Error responding error response")
	}

	return true
}
