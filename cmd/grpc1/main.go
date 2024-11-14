package main

import "goplayground/pkg/grpc1/server"

func main() {
	err := server.Server()
	if err != nil {
		panic(err)
	}
}
