package main

import (
	"context"
	"fmt"
	"goplayground/pkg/grpc1/proto/sample"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	ctx := context.Background()
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8081", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := sample.NewRunnerClient(conn)

	req := &sample.DoRequest{Item: "test"}

	resp, err := client.Do(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Item)
}
