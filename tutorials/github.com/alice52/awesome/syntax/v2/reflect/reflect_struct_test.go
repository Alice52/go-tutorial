package reflect

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	ee := &baseEvent{
		basicEvent: basicEvent{},
		MemberId:   "uid",
	}

	event, err := tryParseStruct2Struct(ee, "basicEvent")
	if err != nil {
		fmt.Printf("parse failed: %v\n", err)
		return
	}

	fmt.Printf("parse success: %v\n", event)

}