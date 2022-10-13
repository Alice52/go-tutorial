package main

import (
	"log"
	"net/rpc/jsonrpc"
)

const port = ":8080"

type ArithRequest struct {
	A, B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func main1() {
	// 1. create connection to rpc server
	conn, err := jsonrpc.Dial("tcp", port)
	if err != nil {
		log.Fatalf("connected to rpc server failed: %v", err)
	}

	// 2. call rpc method to get area
	req := ArithRequest{9, 2}
	var res ArithResponse
	err = conn.Call("ArithService.Multiply", req, &res)
	if err != nil {
		log.Fatalf("call rect rpc service error: %v", err)
	}
	log.Printf("request: %v and res: %v", req, res)
}
