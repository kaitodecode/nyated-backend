package AppError

import "github.com/kaitodecode/nyated-backend/constants"



const (
	ErrJwtInvalidToken ErrorCode = "ERR_JWT_INVALID_TOKEN"
)

var jwtError ErrorMessage = ErrorMessage{
	ErrJwtInvalidToken:{
		constants.EN: "Invalid token",
		constants.ID: "Token authentikasi tidak valid",
	},
}