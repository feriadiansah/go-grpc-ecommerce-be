package utils

import "github.com/feriadiansah/go-grpc-ecommerce-be/pb/common"

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message,
	}
}
