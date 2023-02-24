package pck

import(
	db "github.com/LeilaBeken/golang_ass_1/database"
)

type User struct {
	Id      string
	Name    string
	Email   string
	Balance float64
	Password string
}

func (user *User) setName(name string) {
	user.Name = name
}

func CreateUser(id string, name string, email string, password string, balance float64) {

	insertDynStmt := `insert into "users"("id","name", "email","password","balance") values($1, $2,$3,$4,$5)`
	_, e := db.Db.Exec(insertDynStmt, id, name, email, password, 0.0)
	db.CheckError(e)
}

func GetUser(email string, pass string) (bool, User) {
	rows, err := db.Db.Query(`SELECT "id", "name", "email","balance" FROM "users" where "email"=$1 and "password"=$2`, email, pass)
	db.CheckError(err)

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &email, &balance)
		db.CheckError(err)
		return true, User{id, name, email, balance, pass}
	}
	return false, User{}
}

func IsEmailFree(email string) bool {
	rows, err := db.Db.Query(`SELECT "email" FROM "users" where "email"=$1`, email)
	db.CheckError(err)

	defer rows.Close()
	for rows.Next() {
		return false
	}
	return true
}