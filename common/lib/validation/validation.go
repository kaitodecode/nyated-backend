package validation

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/common/response"
	"github.com/sirupsen/logrus"
)
func ValidateBodyJson(c *gin.Context, request any) (res response.ParamHttpRes,err error){
	if err = c.ShouldBindJSON(request); err != nil {
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
	logrus.Errorf("%v", request)
	if err = validate.Struct(request); err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := AppError.ErrValidationResponse(c, err)
		res = response.ParamHttpRes{
			Code: http.StatusUnprocessableEntity,
			Data: &errResponse,
			Message: &errMessage,
			Gin: c,
		}
		logrus.Errorf("%v", request)
		return
	}
	logrus.Errorf("%v", request)
	return
}