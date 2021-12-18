package main

import (
	"./grpcserver"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	x := 6
	y := 456

	con, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	c := GRPCapi.NewFibonacciClient(con)
	res, err := c.Get(context.Background(), &GRPCapi.FibRequest{X: int32(x), Y: int32(y)})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
