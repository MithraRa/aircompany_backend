package models

import (
	"airline/database"
	models "airline/models/queries"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	Id         uint   `json:id`
	Name       string `json:name`
	Lastname   string `json:lastname`
	Patronymic string `json:patronymic`
	Document   string `json:document`
	Phone      string `json:phone`
	Password   string `json:password`
	Email      string `json:email`
}

const passwordLen = 6
const emailMusk = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

func Validate(email, password string) error {
	if len(password) < passwordLen {
		return fmt.Errorf("incorrect password")
	}

	matched, _ := regexp.MatchString(emailMusk, email)
	if !matched {
		return fmt.Errorf("incorrect email")
	}

	return nil
}

func createPassword(password string) (string, error) {
	re, err := bcrypt.GenerateFromPassword([]byte(password), 16)

	return string(re), err
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) AddUser() error {
	_, err := database.GetDB().Exec(models.InsertUserQuery, u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(name, password, email string) *User {
	hashedPassword, _ := createPassword(password)

	user := &User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return user
}

func FindUserByEmail(email string) (*User, error) {
	var user User

	err := database.GetDB().QueryRow(models.FindUserByEmailQuery, email).Scan(
		&user.Id,
		&user.Name,
		&user.Lastname,
		&user.Patronymic,
		&user.Document,
		&user.Phone,
		&user.Password,
		&user.Email,
	)
	if err != nil {
		return &User{}, fmt.Errorf("can't find the user")
	}

	return &user, nil
}

func FindUserById(id uint) (*User, error) {
	var user User

	err := database.GetDB().QueryRow(models.FindUserByIdQuery, id).Scan(
		&user.Id,
		&user.Name,
		&user.Lastname,
		&user.Patronymic,
		&user.Document,
		&user.Phone,
		&user.Password,
		&user.Email,
	)
	if err != nil {
		return &User{}, fmt.Errorf("can't find the user")
	}

	return &user, nil
}

func (u *User) UpdateInfo(name, lastname, patronymic, document, phone string) error {
	if name != "" {
		u.Name = name
	}
	if lastname != "" {
		u.Lastname = lastname
	}
	if document != "" {
		u.Document = document
	}
	if patronymic != "" {
		u.Patronymic = patronymic
	}
	if phone != "" {
		u.Phone = phone
	}

	_, err := database.GetDB().Exec(models.UpdateUserQuery,
		u.Name,
		u.Lastname,
		u.Patronymic,
		u.Document,
		u.Phone,
		u.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserById(id int) error {
	_, err := database.GetDB().Exec(models.DeleteUserIdQuery,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
