package conc

import "testing"

func Test_poolDemo(t *testing.T) {
	poolDemo()
}

func Test_poolWithContextDemoCancelOnError(t *testing.T) {
	poolWithContextDemoCancelOnError()
}

func Test_poolWithError(t *testing.T) {
	poolWithError()
}

func Test_poolWithResult(t *testing.T) {

	poolWithResult()
}