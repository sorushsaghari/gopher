package services

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"fmt"
	"../models"
)



type UserDb interface {
	FindById(id uint) (*models.User, error)
	Insert(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	Find(field, value string) (*models.User, error)
}

type UserService interface {
	UserDb
	Authenticate(password, email string) (bool, error)
}
type userService struct {
	db *gorm.DB
	UserService
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{db: db}
}

func (us *userService) Insert(user *models.User) error {
	password, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	return us.db.Create(user).Error
}

func (us *userService) Find(field, value string) (*models.User, error) {
	var user models.User
	err := us.db.Where(fmt.Sprintf("%s = ?", field), value).Find(&user).Error
	fmt.Println(user)
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

func (us *userService) FindById(id uint) (*models.User, error) {
	var user models.User
	err := us.db.Where("id = $1", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

func (us *userService) FindByEmail(email string) (*models.User, error) {
	var user models.User
	fmt.Println(email)
	err := us.db.Where("email = ?", email).Find(&user).Error
	fmt.Println(user)
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, err
	}
}

func (us *userService) Authenticate(email string, password string) (bool, error) {
	u, err := us.Find("email", email)
	fmt.Println(u)
	switch err {
	case nil:
		return checkPasswordHash(password, u.Password), nil
	case gorm.ErrRecordNotFound:
		fmt.Println("sik")
		return false, models.ErrNotFound
	default:
		fmt.Println("pashm")
		return false, err
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
