package AppError

import "github.com/kaitodecode/nyated-backend/constants"

const (
	ErrNoteNotFound ErrorCode = "ERR_NOTE_NOT_FOUND"
)

var noteError ErrorMessage = ErrorMessage{
	ErrNoteNotFound: {
		constants.ID: "Catatan tidak ditemukan",
		constants.EN: "Note not found",
	},
}