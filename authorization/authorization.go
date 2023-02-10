package authorization

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/LeilaBeken/golang_ass_1/pck"
)

type Users struct{
	*pck.DatabaseUsers
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (users *Users) SignIn(user, password string) string{
	token := "token"
	for _, u := range users.Users{
		if u.Name == user && u.Password == password{
			return token
		}
	}
	return token
}