package validation

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/common/response"
)
func  ValidateBodyJson(c *gin.Context, request any) (iserror bool, res response.ParamHttpRes){
	iserror = false
	res = response.ParamHttpRes{}
	if err := c.ShouldBindJSON(&request); err != nil {
		iserror = true
		if err.Error() == "EOF" {
			res = response.ParamHttpRes{
				Code: http.StatusBadRequest,
				Err:  errors.New(AppError.GetMessage(c,AppError.ErrJsonBodyIsNotSet)),
				Gin:  c,
			}
			return
		}
		res = response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		}
		return
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		iserror = true
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := AppError.ErrValidationResponse(c, err)
		res = response.ParamHttpRes{
			Code: http.StatusUnprocessableEntity,
			Data: errResponse,
			Message: &errMessage,
			Gin: c,
		}
		return
	}
	return
}