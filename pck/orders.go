package pck

import(
	db "github.com/LeilaBeken/golang_ass_1/database"
)

type Order struct {
	Product  Product
	OrderId  string
	UserRate int
}

func GetOrders(id string) []Order {
	rows, err := db.Db.Query(`SELECT "order_id","product_id","user_rate" from "orders" inner join "users" on "user_id"="users"."id" where "users"."id"=$1`, id)
	db.CheckError(err)
	defer rows.Close()
	orders := make([]Order, 0, 100)
	for rows.Next() {
		var user_rate int
		var order_id string
		err = rows.Scan(&order_id, &id, &user_rate)
		db.CheckError(err)
		order := Order{GetProduct(id), order_id, user_rate}
		orders = append(orders, order)
	}
	return orders
}