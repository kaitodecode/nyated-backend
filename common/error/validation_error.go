package AppError

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
	"github.com/kaitodecode/nyated-backend/common/util"
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/sirupsen/logrus"
)


type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

 
func ErrValidationResponse(ctx context.Context, err error) (validationResponse []ValidationResponse) {
	var fieldErrors validator.ValidationErrors
	lang := GetLangFromContext(ctx)
	if errors.As(err, &fieldErrors) {
		for _, fieldError := range fieldErrors {
			switch fieldError.Tag() {
			case "required":
				var msg string
				if lang == constants.ID {
					msg = fmt.Sprintf("Kolom %s wajib diisi", util.CamelToSnake(fieldError.Field()))
					} else {
					msg = fmt.Sprintf("Field %s is required", util.CamelToSnake(fieldError.Field()))
				}
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   util.CamelToSnake(fieldError.Field()),
					Message: msg,
				})
			case "email":
				var msg string
				if lang == constants.ID {
					msg = fmt.Sprintf("Kolom %s bukanlah email yang valid", util.CamelToSnake(fieldError.Field()))
				} else {
					msg = fmt.Sprintf("Field %s is not valid email", util.CamelToSnake(fieldError.Field()))
				}
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   util.CamelToSnake(fieldError.Field()),
					Message: msg,
				})
			case "min":
				var msg string
				if lang == constants.ID {
					msg = fmt.Sprintf("Kolom %s minimal berisi %s karakter", util.CamelToSnake(fieldError.Field()), fieldError.Param())
				} else {
					msg = fmt.Sprintf("Field %s must be at least %s characters", util.CamelToSnake(fieldError.Field()), fieldError.Param())
				}
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   util.CamelToSnake(fieldError.Field()),
					Message: msg,
				})
			case "max":
				var msg string
				if lang == constants.ID {
					msg = fmt.Sprintf("Kolom %s maksimal berisi %s karakter", util.CamelToSnake(fieldError.Field()), fieldError.Param())
				} else {
					msg = fmt.Sprintf("Field %s must be at most %s characters", util.CamelToSnake(fieldError.Field()), fieldError.Param())
				}
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   util.CamelToSnake(fieldError.Field()),
					Message: msg,
				})
			default:
				errValidator, ok := ErrValidator[fieldError.Tag()]
				if ok {
					count := strings.Count(errValidator, "%s")
					if count == 1 {
						validationResponse = append(validationResponse, ValidationResponse{
							Field:   util.CamelToSnake(fieldError.Field()),
							Message: fmt.Sprintf(errValidator, util.CamelToSnake(fieldError.Field())),
						})
					} else {
						validationResponse = append(validationResponse, ValidationResponse{
							Field:   util.CamelToSnake(fieldError.Field()),
							Message: fmt.Sprintf(errValidator, util.CamelToSnake(fieldError.Field()), fieldError.Param()),
						})
					}
				}else{
					var msg string
					if lang == constants.ID {
						msg = fmt.Sprintf("Kolom %s tidak valid", util.CamelToSnake(fieldError.Field()))
					} else {
						msg = fmt.Sprintf("Field %s is not valid", util.CamelToSnake(fieldError.Field()))
					}
					validationResponse = append(validationResponse, ValidationResponse{
						Field:   util.CamelToSnake(fieldError.Field()),
						Message: msg,
					})
				}
			}
		}
	}
	return validationResponse
}


func WrapError(err error) error {
	logrus.Errorf("error %v", err)
	return err
}

