package registration

import (
	"github.com/LeilaBeken/golang_ass_1/pck"
	"golang.org/x/crypto/bcrypt"
)

type Registration struct {
	Name       string
	SecondName string
	Age        int
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

type dbUsers struct {
	*pck.DatabaseUsers
}

func (d *dbUsers) Register(name string, password string) {
	pass, err := HashPassword(password)
	if err != nil {
		return
	}
	d.Users = append(d.Users, pck.User{Name: name, Password: pass})
}
