package authorization

import (
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{}

type Authorization struct {
	User    string 
	Password string 
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (a *Authorization) SignIn(user, password string) string{
	token := "token"

	if a.User == user && a.Password == password{
		return token
	} else{
		return "not authorized"
	}
}