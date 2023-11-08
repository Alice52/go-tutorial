package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
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

	// 2. handle request
	rpc.HandleHTTP()

	// 3. start listening
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic(err)
	}
}
