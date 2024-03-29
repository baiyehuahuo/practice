package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Calc int

func (c *Calc) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	rpc.Register(new(Calc))
	rpc.HandleHTTP()
	log.Printf("Server RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
