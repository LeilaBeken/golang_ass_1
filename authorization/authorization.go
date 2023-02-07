package authorization

type Authorization struct {
	Login    string
	Password string
}

func (a *Authorization) SignIn(login, password string) string{
	token := "token"

	if a.Login == login && a.Password == password{
		return token
	} else{
		return "not authorized"
	}
}