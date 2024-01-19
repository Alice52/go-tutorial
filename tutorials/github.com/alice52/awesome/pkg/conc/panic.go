package conc

import (
	"errors"
	"fmt"
	"github.com/sourcegraph/conc/panics"
)

func RecoveryPanicAsError() func(err error) {
	return func(err error) {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recovered from panic: %v", r))
		}
	}
}
func Xxx(message string) (a string, err error) {
	defer RecoveryPanicAsError()(err)
	return a, nil
}

// panicDemo recover可能出现的异常
func panicDemo() {
	var pc panics.Catcher
	i := 0

	pc.Try(func() { i += 1 })
	pc.Try(func() { panic("abort!") })
	pc.Try(func() { i += 1 })

	// recover可能出现的panic
	rc := pc.Recovered()

	fmt.Println(i)
	fmt.Println(rc.Value.(string))
	fmt.Println(rc.AsError())

	// 重新panic
	pc.Repanic()
}
