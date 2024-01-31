package reflect

import (
	"errors"
	"reflect"
	"strconv"
)

type basicEvent struct {
	// event name
	name string
	// user data.
	data map[string]any
	// target
	target any
	// mark is aborted
	aborted bool
}

type baseEvent struct {
	basicEvent
	MemberId string
}

// tryParseStruct2Struct: todo
// e is pointer
// e is struct
func tryParseStruct2Struct(e any, name ...string) (*basicEvent, error) {
	beName := ""
	if len(name) == 0 {
		beName = "BaseEvent"
	} else {
		beName = name[0]
	}

	v := reflect.ValueOf(e)
	var be any
	if v.Kind() == reflect.Ptr {
		be = v.Elem().FieldByName(beName).Interface()
	} else if v.Kind() == reflect.Struct {
		be = v.FieldByName(beName).Interface()
	} else {
		return nil, errors.New("unknown event type: " + strconv.Itoa(int(v.Kind())))
	}

	if a, ok := be.(basicEvent); ok {
		return &a, nil
	} else {
		return nil, errors.New("event cast to BaseEvent failed")
	}
}
