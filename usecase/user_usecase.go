package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/trailstem/go-nextjs-rest-api/model"
	"github.com/trailstem/go-nextjs-rest-api/repository"
	"golang.org/x/crypto/bcrypt"
)

// usecase層のinterface
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

/*
以下usecase具体実装
*/

// SignUp implements IUserUsecase.
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}

	return resUser, nil
}

// Login implements IUserUsecase.
func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}
