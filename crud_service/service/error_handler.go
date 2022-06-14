package service

import (
	"github.com/toshkentov01/task/crud_service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorHandler(l logger.Logger, err error, message string) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)

	if !ok || st.Code() == codes.Internal {
		l.Error("Error while deleting post: " + err.Error())
		return status.Error(codes.Internal, "Internal Server Error")

	} else if st.Code() == codes.NotFound {
		l.Error("Error while deleting post: " + err.Error())
		return status.Error(codes.NotFound, "Not Found")
	}

	return nil
}
