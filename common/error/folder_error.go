package AppError

import "github.com/kaitodecode/nyated-backend/constants"

const (
	ErrFolderNotFound ErrorCode = "ERR_FOLDER_NOT_FOUND"
)

var folderErrors ErrorMessage = ErrorMessage{
	ErrFolderNotFound: {constants.ID:"Data folder tidak ditemukan", constants.EN:"Folder not found"},
}