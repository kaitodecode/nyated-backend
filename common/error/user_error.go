package AppError

import "github.com/kaitodecode/nyated-backend/constants"


const (
	ErrUserNotFound ErrorCode = "ERR_USER_NOT_FOUND"
	ErrUserAlreadyExist ErrorCode = "ERR_USER_ALREADY_EXIST"
	ErrUserPasswordDoestNotMatch ErrorCode = "ERR_USER_PASSWORD_DOES_NOT_MATCH"
)

var userError ErrorMessage = ErrorMessage{
	ErrUserNotFound:{
		constants.EN: "User not found",
		constants.ID: "User tidak ditemukan",
	},
	ErrUserAlreadyExist:{
		constants.EN: "There is data that has been used by other users",
		constants.ID: "Terdapat data yang sudah digunakan user lain",
	},
	ErrUserPasswordDoestNotMatch:{
		constants.EN: "Password does not match",
		constants.ID: "password dan konfirmasi password tidak sama",
	},
}