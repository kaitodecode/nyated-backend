package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/constants"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
	Token *string `json:"token,omitempty"`
}

type ParamHttpRes struct {
	Code int
	Err error
	Message *string
	Gin *gin.Context
	Data any
	Token *string
}

func HttpResponse(param ParamHttpRes) {
	if param.Err == nil {
		param.Gin.AbortWithStatusJSON(param.Code, Response{
			Status: constants.SUCCESS,
			Message: http.StatusText(http.StatusOK),
			Data: param.Data,
			Token: param.Token,
		})
		return
	}
	message := AppError.GetMessage(param.Gin.Request.Context(), AppError.ErrInternalServerError)
	if param.Message != nil {
		message = *param.Message
	}else if param.Err != nil {
		message = param.Err.Error()
	}

	param.Gin.AbortWithStatusJSON(param.Code, Response{
		Status: constants.ERROR,
		Message: message,
		Data: param.Data,
	})
	return
}