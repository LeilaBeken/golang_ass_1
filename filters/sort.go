package filters

import "net/http"

var Sort bool = false

func SortPrice(writer http.ResponseWriter, request *http.Request) {
	if Sort { Sort = false } else { Sort = true}

	http.Redirect(writer, request, "/home", http.StatusFound)

}