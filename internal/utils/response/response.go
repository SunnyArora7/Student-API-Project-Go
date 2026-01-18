package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	typesFile "studentPackage/internal/type"

	"github.com/go-playground/validator/v10"
)

func GeneralErrorResponse(error string, statusCode int) typesFile.GeneralError {
	return typesFile.GeneralError{
		StatusCode: statusCode,
		Error:      error,
	}
}
func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json") //This tell the users the content is of type json
	w.WriteHeader(status)                              //Sends HTTP status code to client
	return json.NewEncoder(w).Encode(data)

}

func ValidationError(errs validator.ValidationErrors, status int) typesFile.GeneralError {
	var errors []string
	for _, e := range errs {
		switch e.ActualTag() {
		case "required":
			errors = append(errors, fmt.Sprintf("%s field is required", e.Field()))
		default:
			errors = append(errors, fmt.Sprintf("%s field is invalid", e.Field()))

		}
	}
	return typesFile.GeneralError{
		StatusCode: status,
		Error:      strings.Join(errors, ","),
	}
}
