package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
This sample is used to realize the area and perimeter of the rectangle.
*/

const port string = ":8080"

type ArithService struct{}

type ArithRequest struct {
	A, B int
}
type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func (arith ArithService) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (arith *ArithService) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("Divisor cannot be 0")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	// 1. create service and register
	rpc.Register(new(ArithService))

	// 2. start to listening
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			log.Println("new client connection")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
