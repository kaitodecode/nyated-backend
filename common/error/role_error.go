package AppError

import "github.com/kaitodecode/nyated-backend/constants"

const (
	ErrRoleNotFound ErrorCode = "ERR_ROLE_NOT_FOUND"
)

var roleError ErrorMessage = ErrorMessage{
	ErrRoleNotFound:{
		constants.ID: "Data role tidak ditemukan",
		constants.EN: "Role not found",
	},
} 