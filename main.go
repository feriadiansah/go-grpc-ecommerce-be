package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/feriadiansah/go-grpc-ecommerce-be/internal/handler"
	"github.com/feriadiansah/go-grpc-ecommerce-be/pb/service"
	"github.com/feriadiansah/go-grpc-ecommerce-be/pkg/database"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func errorMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)

			err = status.Errorf(codes.Internal, "Internal Server Error")
		}
	}()
	res, err := handler(ctx, req)
	return res, err
}

func main() {
	ctx := context.Background()
	godotenv.Load()
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Panicf("Error when listening: %v", err)
	}

	database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Database connected")

	serviceHandler := handler.NewServiceHandler()
	serv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		errorMiddleware,
	))

	//get value from .env file
	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("reflection registered")

	}

	service.RegisterHelloWorldServiceServer(serv, serviceHandler)
	log.Println("gRPC server is running on port 50051")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("Serving is error: %v", err)
	}
}
