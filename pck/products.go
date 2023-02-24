package pck

import (
	"database/sql"
	rt "github.com/LeilaBeken/golang_ass_1/ratings"
	db "github.com/LeilaBeken/golang_ass_1/database"
)

var id string
var name string
var email string
var balance float64
var price float64

type Product struct {
	Id       string
	Name     string
	Price    float64
	Rating   float64
	OrderCnt int
}

func GetProducts(sort bool) []Product {
	var err error
	var rows *sql.Rows
	if sort {
		rows, err = db.Db.Query(`SELECT "id","name","price" FROM "products" order by "price" desc`)
	} else {
		rows, err = db.Db.Query(`SELECT "id","name","price" FROM "products" order by "price"`)
	}
	db.CheckError(err)

	defer rows.Close()
	products := make([]Product, 0, 100)
	for rows.Next() {
		err = rows.Scan(&id, &name, &price)
		cnt, avg := rt.AvgProductRate(id)
		db.CheckError(err)
		product := Product{id, name, price, avg, cnt}
		products = append(products, product)
	}

	return products
}

func GetProduct(id string) Product {
	rows, err := db.Db.Query(`SELECT "id","name","price" FROM "products" where "id"=$1`, id)
	db.CheckError(err)
	var product Product
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &price)
		db.CheckError(err)
		cnt, avg := rt.AvgProductRate(id)
		product = Product{id, name, price, avg, cnt}
	}
	return product
}

func GetProductsByName(name string) []Product {
	rows, err := db.Db.Query(`SELECT "id","name","price" FROM "products" where "name"=$1`, name)
	db.CheckError(err)
	defer rows.Close()
	products := make([]Product, 0, 100)
	for rows.Next() {
		err = rows.Scan(&id, &name, &price)
		db.CheckError(err)
		cnt, avg := rt.AvgProductRate(id)
		product := Product{id, name, price, avg, cnt}
		products = append(products, product)
	}

	return products
}