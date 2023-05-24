package errrors

import (
	"encoding/json"
	"net/http"

	"github.com/lib/pq"
)

type Error struct {
	Statuscode int
	ErrMsg     string
}


var (
	InternalServerError = GenerateError(500, "Something went wrong! Please try again later")
	DuplicateValue      = GenerateError(409, "Category already present")
	CategoryNotFoundErr = GenerateError(404, "Category does not exists")
	InvalidDataErr      = GenerateError(400, "Invalid data input")
	BadRequest          = GenerateError(400, "Bad request")
	Unauthorized        = GenerateError(401, "User not Uthorized")
	ForeignKeyViolates  = GenerateError(409, "Foreign key violates")
	InvalidRequestBody  = GenerateError(400, "Invalid value passed in request body")
	InvalidQueryParams  = GenerateError(400, "Invalid or NULL value in URL")
)

func ErrorGenerator(w http.ResponseWriter, err error) {
	errObj := Error{}
	if pqErr, ok := err.(*pq.Error); ok {
		dynamicError := ErrorGenerateFromPq(pqErr)
		errObj.ErrorResponseGenerator(w, dynamicError.Statuscode, dynamicError.ErrMsg)
	} else {
		dynamicErrorForCustom := ErrorGeneratorFromRequest(err)
		errObj.ErrorResponseGenerator(w, dynamicErrorForCustom.Statuscode, dynamicErrorForCustom.ErrMsg)
	}
	return 
}

func ErrorGenerateFromPq(pqErr error) *Error {
	switch pqErr.(*pq.Error).Code {
	case "23503":
		return &Error{
			Statuscode: 409,
			ErrMsg:     ForeignKeyViolates.Error(),
		}
	case "23505":
		return &Error{
			Statuscode: 404,
			ErrMsg:     DuplicateValue.Error(),
		}
	case "22P02":
		return &Error{
			Statuscode: 422,
			ErrMsg:     InvalidRequestBody.Error(),
		}
	case "42804":
		return &Error{
			Statuscode: 400,
			ErrMsg:     InvalidQueryParams.Error(),
		}
	default:
		return &Error{
			Statuscode: 500,
			ErrMsg:     InternalServerError.Error(),
		}
	}
}

func ErrorGeneratorFromRequest(err error) *Error {
	switch err.Error() {
	case "User not Uthorized":
		return &Error{
			Statuscode: 401,
			ErrMsg:     Unauthorized.Error(),
		}
	case "Category already present":
		return &Error{
			Statuscode: 409,
			ErrMsg:     DuplicateValue.Error(),
		}
	case "Invalid value passed in body":
		return &Error{
			Statuscode: 400,
			ErrMsg:     InvalidRequestBody.Error(),
		}
	case "Invalid value passed in URL":
		return &Error{
			Statuscode: 400,
			ErrMsg:     InvalidQueryParams.Error(),
		}
	case "category does not exists":
		return &Error{
			Statuscode: 404,
			ErrMsg:     CategoryNotFoundErr.Error(),
		}
	default:
		return &Error{
			Statuscode: 500,
			ErrMsg:     InternalServerError.Error(),
		}
	}
}

func GenerateError(statusCode int, errMsg string) error {
	return &Error{
		Statuscode: statusCode,
		ErrMsg:     errMsg,
	}
}

func (e Error) Error() string {
	return e.ErrMsg
}

func (e *Error) ErrorResponseGenerator(w http.ResponseWriter, errCode int, errMsg string) {
	e.WriteError(w, errCode, errMsg)

}

func (e Error) WriteError(w http.ResponseWriter, errorCode int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err})
	return
}
