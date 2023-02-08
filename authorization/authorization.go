package authorization

import (
	"golang.org/x/crypto/bcrypt"
)

type Authorization struct {
	Login    string
	Password string
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (a *Authorization) SignIn(login, password string) string{
	token := "token"

	if a.Login == login && a.Password == password{
		return token
	} else{
		return "not authorized"
	}
}