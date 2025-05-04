package lib

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kaitodecode/nyated-backend/domain/dto"
)


type Claims struct {
	User *dto.UserResponse
	jwt.RegisteredClaims
}