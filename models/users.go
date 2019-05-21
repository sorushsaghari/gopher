package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserDb interface {
	FindById(id uint) (*User, error)
	Insert(user *User) error
	FindByEmail(email string) (*User, error)
	Find(field, value string) (*User, error)
}

type UserService interface {
	UserDb
	Authenticate(password, email string) (bool, error)
}
type userService struct {
	db *gorm.DB
	UserService
}
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `json:"password"`
}


func NewUserService(db *gorm.DB) *userService {
	return &userService{db: db}
}



func (us *userService) Insert(user *User) error {
	password, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	return us.db.Create(user).Error
}

func (us* userService) Find(field, value string) (*User, error){
	var user User
	err := us.db.Where(fmt.Sprintf("%s = ?", field), value).Find(&user).Error
	fmt.Println(user)
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (us *userService) FindById(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = $1", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (us* userService) FindByEmail(email string) (*User, error)  {
	var user User
	fmt.Println(email)
	err := us.db.Where("email = ?", email).Find(&user).Error
	fmt.Println(user)
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (us* userService) Authenticate(email string, password string) (bool, error) {
	u, err := us.Find("email", email)
	fmt.Println(u)
	switch err {
	case nil:
		return checkPasswordHash(password, u.Password), nil
	case gorm.ErrRecordNotFound:
		fmt.Println("sik")
		return false, ErrNotFound
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
