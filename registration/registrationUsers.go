package registration

import (
	"net/http"

	pck "github.com/LeilaBeken/golang_ass_1/pck"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return bytes, err
}


func CreateAccount(writer http.ResponseWriter, request *http.Request) {
	login := request.FormValue("login")
	password, err := HashPassword(request.FormValue("password"))
	name := request.FormValue("nickname")
	id := uuid.New(password)

	if err != nil{
		http.Redirect(writer, request, "/error", http.StatusFound) 
		return 
	}

	if pck.IsEmailFree(login) {
		pck.CreateUser(id.String(), name, login, string(password), 0)
		http.Redirect(writer, request, "/home", http.StatusFound)
	}
	http.Redirect(writer, request, "/error", http.StatusFound)
}
