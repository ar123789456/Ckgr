package usecase

import (
	"cgr/models"
	"cgr/user"
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	repository user.Repository
}

func NewUsecase(repository user.Repository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

var errWrPass = errors.New("Wrong Password")

func (uc *Usecase) LogIn(c context.Context, p models.User) (string, error) {
	user, err := uc.repository.Get(c, p.FullName)
	if err != nil {
		return "", err
	}
	if !CheckPasswordHash(p.Password, user.Password) {
		return "", errWrPass
	}
	uuID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	err = uc.repository.UpdateSesion(c, user.ID, uuID.String())
	return uuID.String(), err
}

func (uc *Usecase) Create(c context.Context, u models.User) error {
	var err error
	u.Password, err = HashPassword(u.Password)
	if err != nil {
		return err
	}
	return uc.repository.Create(c, u)
}

func (uc *Usecase) Delete(c context.Context, id int) error {
	return uc.repository.Delete(c, id)
}

func (uc *Usecase) Update(c context.Context, u models.User) error {
	return uc.repository.Update(c, u)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
