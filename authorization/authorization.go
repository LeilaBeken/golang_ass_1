package authorization

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
	pck "github.com/LeilaBeken/golang_ass_1/pck"
)

var Account pck.User

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func signIn(writer http.ResponseWriter, request *http.Request) {
	login := request.FormValue("login")
	password := request.FormValue("password")
	found, user := pck.GetUser(login, password)
	if !found  || !CheckPasswordHash(password, user.Password){
		http.Redirect(writer, request, "/error", http.StatusFound)
	}
	Account = user
	http.Redirect(writer, request, "/home", http.StatusFound)
}