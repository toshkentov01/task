package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/toshkentov01/task/data_service/pkg/errs"
	l "github.com/toshkentov01/task/data_service/pkg/logger"
)

const (
	// InternalServerError ...
	InternalServerError = "Internal Server Error"

	// NotFoundError ...
	NotFoundError = "Not Found"

	// InvalidArgumentError ...
	InvalidArgumentError = "Invalid Argument"
)

// errorHandler function for handling errors in service
func errorHandler(logger l.Logger, err error, message string, req ...interface{}) error {

	if err == nil {
		return nil

	} else if err == errs.ErrNotFound {
		logger.Error("Not Found, error: "+err.Error()+message, l.Any("req", req))
		return status.Error(codes.NotFound, NotFoundError)

	} else if err == errs.ErrInternal {
		logger.Error("Internal Server Error: "+err.Error()+message, l.Any("req", req))
		return status.Error(codes.Internal, InternalServerError)
	}

	return nil
}
