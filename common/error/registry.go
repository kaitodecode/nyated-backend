package AppError

import (
	"context"

	"github.com/kaitodecode/nyated-backend/constants"
)

var AllErrors ErrorMessage = ErrorMessage{}

func RegisterErrors(messages ...ErrorMessage) {
	for _, m := range messages {
		for code, langs := range m {
			AllErrors[code] = langs
		}
	}
}

func Init(){
	RegisterErrors(
		commonError, 
		jwtError, 
		userError,
		roleError,
	)
}

func GetMessage(ctx context.Context, code ErrorCode) string {
	lang := GetLangFromContext(ctx)
	if msgLangMap, ok := AllErrors[code]; ok {
		if msg, ok := msgLangMap[lang]; ok {
			return msg
		}
	}
	return "Unknown error"
}


func GetLangFromContext(ctx context.Context) constants.AppLanguage {
	if val := ctx.Value(constants.CONTEXT_LANG); val != nil {
		if lang, ok := val.(constants.AppLanguage); ok && lang != "" {
			return lang
		}
	}
	return constants.ID
}
