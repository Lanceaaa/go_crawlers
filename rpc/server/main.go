package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"example.com/go-http-demo/crawler/rpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
