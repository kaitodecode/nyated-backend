package services

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	AppError "github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/common/lib"
	"github.com/kaitodecode/nyated-backend/common/util"
	"github.com/kaitodecode/nyated-backend/config"
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/kaitodecode/nyated-backend/domain/dto"
	"github.com/kaitodecode/nyated-backend/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repositories repositories.IRepositoryRegistry
}

type IUserService interface {
	Login(context.Context, *dto.UserLoginRequest) (*dto.UserLoginResponse, error) 
	Register(context.Context, *dto.UserRegisterRequest) (error)
	GetUserLogin(context.Context) (*dto.UserResponse, error)
}

func NewUserService(repositories repositories.IRepositoryRegistry) IUserService {
	return &UserService{repositories}
}

func (s *UserService) GetUserLogin(c context.Context) (user *dto.UserResponse, err error){
	user, err = util.GetUser(c)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *UserService) isEmailExist(ctx context.Context, email string) bool {
	_, err := u.repositories.UserRepository().FindByEmail(ctx, email)

	return err == nil
}
// func (u *UserService) isIDExist(ctx context.Context, id string) bool {
// 	_, err := u.repositories.UserRepository().FindByID(ctx, id)

// 	return err == nil
// }

func (s *UserService) Register(c context.Context, req *dto.UserRegisterRequest) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 
	}
	role, err := s.repositories.RoleRepository().FindByCode(c, constants.ROLE_USER)

	if err!= nil {
		return 
	}

	if s.isEmailExist(c, req.Email) {
		err = errors.New(AppError.GetMessage(c, AppError.ErrUserAlreadyExist))
		return 
	}

	if req.Password != req.ConfirmPassword {
		err = errors.New(AppError.GetMessage(c, AppError.ErrUserPasswordDoestNotMatch))
		return 
	}

	err = s.repositories.UserRepository().Register(c, &dto.UserRegisterRequest{
		Name:        req.Name,
		Email:       req.Email,
		Password:    string(hashedPassword),
		RoleID: role.ID,
	})

	if err != nil {
		return
	}

	return 
}

func (s *UserService) Login(c context.Context, req *dto.UserLoginRequest) (response *dto.UserLoginResponse, err error){
	user, err := s.repositories.UserRepository().FindByEmail(c, req.Email)

	if err != nil {
		return nil, err
	}

	if compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); compareErr != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(time.Duration(config.Config.JWTExpirationTime) * time.Minute).Unix()

	data := &dto.UserResponse{
		ID: user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role.Code,
	}

	claims := &lib.Claims{
		User: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JWTSecretKey))

	if err != nil {
		return nil, err
	}

	response = &dto.UserLoginResponse{
		Token: tokenString,
		User:  *data,
	}

	return
}


