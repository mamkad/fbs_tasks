/*
 * Сервис на основе gRPC
 */
package GRPCapi

import (
	"./fibonacci"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

//структура сервера
type GRPCServer struct{}

//Вызов Get
//Вычислить числа фибоначчи на диапазон
//Волучает начальный и конечный номера, затем возвращает значения
func (c *GRPCServer) Get(cnx context.Context, req *FibRequest) (*FibResponse, error) {
	res, err := fibonacci.Fibonacci(int64(req.GetX()), int64(req.GetY()))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &FibResponse{Result: res}, nil
}

//Создание grpc сервера и структуры GRPCServer
//Создание сервера Fibonacci, а также запуск соединения
func Start() {
	s := grpc.NewServer()
	srv := &GRPCServer{}
	RegisterFibonacciServer(s, srv)
	
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Start... ")

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
