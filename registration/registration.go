package registration

import (
	"github.com/LeilaBeken/golang_ass_1/pck"
)

type Registration struct {
	Name       string
	SecondName string
	Age        int
}

type db struct {
	*pck.Database
}

func (d *db) Register(name string, password string) {
	d.Users = append(d.Users, pck.User{Name: name, Password: password})
}

func (d *db) RegisterItem(item pck.Item) {
	d.Items = append(d.Items, item)
}
