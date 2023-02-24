package ratings

import(
	db "github.com/LeilaBeken/golang_ass_1/database"
)

func SetRate(id string, rate int) {
	updateStmt := `update "orders" set "user_rate"=$1 where "order_id"=$2`
	_, e := db.Db.Exec(updateStmt, rate, id)
	db.CheckError(e)
}