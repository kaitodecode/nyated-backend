package AppError

import (
	"github.com/kaitodecode/nyated-backend/constants"
)

type ErrorCode string

const (
	ErrInternalServerError ErrorCode = "ERR_INTERNAL_SERVER_ERROR"
	ErrUnAuthenticateError ErrorCode = "ERR_UNAUTHENTICATE_ERROR"
	ErrToManyRequest       ErrorCode = "ERR_TO_MANY_REQUEST"
	ErrSqlError			   ErrorCode = "ERR_SQL_ERROR"
	ErrJsonBodyIsNotSet		ErrorCode = "ERR_JSON_BODY_IS_NOT_SET"
)

type ErrorMessageLanguage map[constants.AppLanguage]string
type ErrorMessage map[ErrorCode]ErrorMessageLanguage

var commonError ErrorMessage = ErrorMessage{
	ErrInternalServerError: {
		constants.EN: "Internal Server Error",
		constants.ID: "Terjadi Error Server",
	},
	ErrUnAuthenticateError: {
		constants.EN: "Unauthenticated",
		constants.ID: "Pengguna belum terauthentikasi",
	},
	ErrToManyRequest:{
		constants.EN: "To many request",
		constants.ID: "Terlalu banyak permintaan ke server",
	},
	ErrSqlError:{
		constants.EN: "Error in sql",
		constants.ID: "Terdapat error pada proses query",
	},
	ErrJsonBodyIsNotSet:{
		constants.ID: "Pastikan anda mengatur body json terlebih dahulu",
		constants.EN: "There is data that has been used by other users",
	},
}
