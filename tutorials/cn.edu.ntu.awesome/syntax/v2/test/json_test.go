package test

import (
	"testing"

	_json "cn.edu.ntu.awesome/syntax/v2/json"
)

func TestDecoder(t *testing.T) {
	_json.JsonDecoder()
}

func TestSerialize(t *testing.T) {

	_json.Serialize()
}

func TestDeSerialize(t *testing.T) {

	_json.DeSerialize()
}
