package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	AppError "github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/common/lib"
	"github.com/kaitodecode/nyated-backend/common/response"
	"github.com/kaitodecode/nyated-backend/config"
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/sirupsen/logrus"
)

func HandlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("Recovered from error panic: %v", r)
				ctx.JSON(http.StatusInternalServerError, response.Response{
					Status:  constants.ERROR,
					Message: AppError.GetMessage(ctx, AppError.ErrInternalServerError),
					Data: r,
				})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}

func RateLimiter(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusTooManyRequests, response.Response{
				Status:  constants.ERROR,
				Message: AppError.GetMessage(ctx, AppError.ErrInternalServerError),
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func extractBearerToken(token string) string {
	arrayToken := strings.Split(token, " ")
	if len(arrayToken) == 2 {
		return arrayToken[1]
	}
	return ""
}

func responseUnAuthorize(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, response.Response{
		Status:  constants.ERROR,
		Message: message,
	})
	c.Abort()
}

// func validationApiKey(c *gin.Context) error {
// 	apiKey := c.GetHeader(constants.XApiKey)
// 	requestAt := c.GetHeader(constants.XRequestAt)
// 	serviceName := c.GetHeader(constants.XServiceName)
// 	signatureKey := config.Config.SignatureKey

// 	valdateKey := fmt.Sprintf("%s:%s:%s", serviceName, signatureKey, requestAt)

// 	hash := sha256.Sum256([]byte(valdateKey))
// 	resultHash := hex.EncodeToString(hash[:])
	
// 	if apiKey != resultHash {
// 		return errorConstant.ErrUnAuthorized
// 	}

// 	return nil
// }

func validatedBearerToken(c *gin.Context, token string) error {
	if !strings.Contains(token, "Bearer") {
		return errors.New(AppError.GetMessage(c, AppError.ErrUnAuthenticateError))
	}

	tokenString := extractBearerToken(token)
	if tokenString == "" {
		return errors.New(AppError.GetMessage(c, AppError.ErrUnAuthenticateError))
	}

	claims := &lib.Claims{}
	tokenJwt, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(AppError.GetMessage(c, AppError.ErrJwtInvalidToken))
		}

		jwtSecret := []byte(config.Config.JWTSecretKey)

		return jwtSecret, nil
	})

	if err != nil || !tokenJwt.Valid {
		return errors.New(AppError.GetMessage(c, AppError.ErrUnAuthenticateError))
	}

	userLogin := c.Request.WithContext(context.WithValue(c.Request.Context(), constants.CONTEXT_USER, claims.User))
	c.Request = userLogin
	c.Set(constants.CONTEXT_TOKEN, token)

	return nil
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		token := c.GetHeader(constants.Authorization)


		if token == "" {
			responseUnAuthorize(c, AppError.GetMessage(c, AppError.ErrUnAuthenticateError))
			return
		}

		err = validatedBearerToken(c, token)

		if err != nil {
			responseUnAuthorize(c, err.Error())
			return
		}

		// err = validationApiKey(c)
		// if err != nil {
		// 	responseUnAuthorize(c, err.Error())
		// 	return
		// }

		c.Next()
	}
}
