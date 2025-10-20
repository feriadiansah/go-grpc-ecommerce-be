package handler

import (
	"context"
	"errors"
	"fmt"

	protovalidate "buf.build/go/protovalidate"
	"github.com/feriadiansah/go-grpc-ecommerce-be/internal/utils"
	"github.com/feriadiansah/go-grpc-ecommerce-be/pb/common"
	"github.com/feriadiansah/go-grpc-ecommerce-be/pb/service"
)

//	type IServiceHandler interface{
//		 HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error)
//	}
type serviceHandler struct {
	service.UnimplementedHelloWorldServiceServer
}

func (sh *serviceHandler) HelloWorld(ctx context.Context, request *service.HelloWorldRequest) (*service.HelloWorldResponse, error) {
	if err := protovalidate.Validate(request); err != nil {
		var validationError *protovalidate.ValidationError
		if errors.As(err, &validationError) {

			var validationErrorResponse []*common.ValidationError = make([]*common.ValidationError, 0)
			for _, violation := range validationError.Violations {
				validationErrorResponse = append(validationErrorResponse, &common.ValidationError{
					Field:   *violation.Proto.Field.Elements[0].FieldName,
					Message: *violation.Proto.Message,
				})
			}
			return &service.HelloWorldResponse{
				Base: &common.BaseResponse{
					ValidationErrors: validationErrorResponse,
					StatusCode:       400,
					Message:          "Validation Error",
					IsError:          true,
				},
			}, nil
		}
		return nil, err
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
