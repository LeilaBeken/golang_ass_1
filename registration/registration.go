package registration

import (
	"pck"
)

type Registration struct {
	Name       string
	SecondName string
	Age        int
}

func (s *Database) Register(name string, password string) {
    s.Users = append(s.Users, User{Name: name, Password: password})
}