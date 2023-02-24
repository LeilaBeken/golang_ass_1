package ratings

import (
	db "github.com/LeilaBeken/golang_ass_1/database"
)

func AvgProductRate(id string) (int, float64) {
	rows, err := db.Db.Query(`SELECT count("product_id"),avg("user_rate") from "orders"  where "user_rate"!=0  group by "product_id" having "product_id"=$1`, id)
	db.CheckError(err)
	var avg float64
	var cnt int
	for rows.Next() {

		err = rows.Scan(&cnt, &avg)
		db.CheckError(err)
		return cnt, avg
	}
	return cnt, avg
}