package util

import (
	"context"
	"errors"

	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/kaitodecode/nyated-backend/domain/dto"
)

func GetUser(ctx context.Context) (*dto.UserResponse, error){
	value := ctx.Value(constants.CONTEXT_USER)

	if value == nil {
		return nil, errors.New("user not found in context")
	}

	user, ok := value.(*dto.UserResponse)

	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return user, nil
}