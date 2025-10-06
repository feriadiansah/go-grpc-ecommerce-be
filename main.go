package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/feriadiansah/go-grpc-ecommerce-be/internal/handler"
	"github.com/feriadiansah/go-grpc-ecommerce-be/pb/service"
	"github.com/feriadiansah/go-grpc-ecommerce-be/pkg/database"
	grpcmiddleware "github.com/feriadiansah/go-grpc-ecommerce-be/pkg/gpcmiddleware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//udah di pindahin ke pkg/grpcmiddleware/error_middleware.go
// func errorMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Println(r)
// 			debug.PrintStack()
// 			err = status.Errorf(codes.Internal, "Internal Server Error")
// 		}
// 	}()
// 	res, err := handler(ctx, req)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, status.Errorf(codes.Internal, "Internal Server Error")
// 	}
// 	return res, err
// }

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
		grpcmiddleware.ErrorMiddleware,
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
