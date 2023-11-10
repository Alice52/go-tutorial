package oop

import (
	"time"
)

// @see gin_practice/routers.go
type Option func(config *RpcConfig)

type RpcConfig struct {
	timeout time.Duration
	cluster string
	host    string
}

// func as setXx
func Timeout(duration time.Duration) Option {
	return func(config *RpcConfig) {
		config.timeout = duration
	}
}
func Cluster(cluster string) Option {
	return func(config *RpcConfig) {
		config.cluster = cluster
	}
}
func Host(host string) Option {
	return func(config *RpcConfig) {
		config.host = host
	}
}

func RpcMethod(ops ...Option) {
	// build args by executing args-func
	var rpcConf RpcConfig
	for _, opFunc := range ops {
		opFunc(&rpcConf)
	}
	// invoke rpc method
}

func OverrideUsage() {
	RpcMethod()
	RpcMethod(Timeout(100))
	RpcMethod(Timeout(100), Host("127.0.0.1"))
}
