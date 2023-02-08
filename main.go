package main

import (
	"fmt"

	"github.com/LeilaBeken/golang_ass_1/pck"
	"golang.org/x/crypto/bcrypt"
)

type Registration struct {
	Name       string
	SecondName string
	Age        int
}

func main(){
	pas := "secret"
	has, err := HashPassword(pas)
	match := CheckPasswordHash(pas, has)
	if err != nil{
		fmt.Println("err")
	} else if !match{
		fmt.Println("not match")
	} else {
		fmt.Println(pas, has)
	}
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

type dbUsers struct {
	*pck.DatabaseUsers
}

func (d *dbUsers) Register(name string, password string) {
	pass, err := HashPassword(password)
	if err != nil{
		return
	}
	d.Users = append(d.Users, pck.User{Name: name, Password: pass})
}