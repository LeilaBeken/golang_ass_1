package itemsearch

import (
	"net/http"
	"text/template"
	db "github.com/LeilaBeken/golang_ass_1/database"
	pck "github.com/LeilaBeken/golang_ass_1/pck"
)

func Search(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("template/search.html")
	db.CheckError(err)
	search := request.FormValue("search")
	products := pck.GetProductsByName(search)
	var data = struct {
		Search   string
		Products []pck.Product
	}{
		Search:   search,
		Products: products,
	}
	err = html.Execute(writer, data)
	db.CheckError(err)
}