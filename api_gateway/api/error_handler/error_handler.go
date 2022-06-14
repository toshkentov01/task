package errorhandler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/toshkentov01/task/api_gateway/api/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInvalidParams ...
	ErrorCodeInvalidParams = "INVALID_PARAMS"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	//ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	//ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
	//ErrorCodeWrongClub ...
	ErrorCodeWrongClub = "WRONG_CLUB"
	//ErrorCodePasswordsNotEqual ...
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
	// ErrorExpectationFailed ...
	ErrorExpectationFailed = "EXPECTATION_FAILED"
	// ErrorUpgradeRequired ...
	ErrorUpgradeRequired = "UPGRADE_REQUIRED"
	// ErrorInvalidCredentials ...
	ErrorInvalidCredentials = "INVALID_CREDENTIALS"
)

// HandleGrpcErrWithMessage ...
func HandleGrpcErrWithMessage(c *fiber.Ctx, err error, message string, args ...interface{}) error {
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		log.Println(message+",", err)
		return c.Status(http.StatusInternalServerError).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorCodeInternal,
					Message: st.Message(),
				},
			},
		)

	} else if st.Code() == codes.NotFound {
		log.Println(message+", ", err)
		return c.Status(http.StatusNotFound).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorCodeNotFound,
					Message: "invalid cridentials, not found from database",
				},
			},
		)

	} else if st.Code() == codes.Unavailable {
		log.Println(message+", ", err)
		return c.Status(http.StatusInternalServerError).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorCodeInternal,
					Message: "Internal Server Error",
				},
			},
		)

	} else if st.Code() == codes.AlreadyExists {
		log.Println(message+", ", err)
		return c.Status(http.StatusConflict).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorCodeAlreadyExists,
					Message: st.Message(),
				},
			},
		)

	} else if st.Code() == codes.InvalidArgument {
		log.Println(message+", ", err)
		return c.Status(http.StatusBadRequest).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorBadRequest,
					Message: st.Message(),
				},
			},
		)

	} else if st.Code() == codes.DataLoss {
		log.Println(message+", ", err)
		return c.Status(http.StatusBadRequest).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorBadRequest,
					Message: st.Message(),
				},
			},
		)

	} else if st.Code() == codes.PermissionDenied || st.Message() == "forbidden" {
		log.Println(message+", ", err)
		return c.Status(http.StatusForbidden).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorCodeForbidden,
					Message: st.Message(),
				},
			},
		)

	} else if err != nil {
		log.Println(message+", ", err)
		return c.Status(http.StatusInternalServerError).JSON(
			models.Response{
				Error: true,
				Data: models.Error{
					Status:  ErrorBadRequest,
					Message: st.Message(),
				},
			},
		)

	}

	return nil
}

// HandleInternalWithMessage ...
func HandleInternalWithMessage(c *fiber.Ctx, err error, message string) error {
	log.Println(message+" ", err)
	return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
		Error: true,
		Data: models.Error{
			Status:  "Internal Server Error",
			Message: message,
		},
	})
}

// HandleBadRequestErrWithMessage ...
func HandleBadRequestErrWithMessage(c *fiber.Ctx, err error, message string) error {
	log.Println(message+" ", err)
	return c.Status(fiber.StatusBadRequest).JSON(models.Response{
		Error: true,
		Data: models.Error{
			Status:  "Bad Request",
			Message: message,
		},
	})
}
