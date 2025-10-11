package handler

import (
	"context"
	"fmt"

	protovalidate "buf.build/go/protovalidate"
	"github.com/feriadiansah/go-grpc-ecommerce-be/internal/utils"
	"github.com/feriadiansah/go-grpc-ecommerce-be/pb/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//	type IServiceHandler interface{
//		 HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error)
//	}
type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Validation error: %v", err)
	}
	// panic(errors.New("Pointer nil"))
	return &service.HelloWorldResponse{
		Message: fmt.Sprintf("Hello %s", request.Name),
		Base:    utils.SuccessResponse("Success"),
		//sebelum menggunakan utils
		// Base: &common.BaseResponse{
		// 	StatusCode: 200,
		// 	Message:    "Success",
		// },
	}, nil
}

// func NewServiceHandler() IServiceHandler {
// 	return &serviceHandler{}
// }

func NewServiceHandler() *serviceHandler {
	return &serviceHandler{}
}
