package itemfilters

import(
	"github.com/LeilaBeken/golang_ass_1/pck"
)

type items struct{
	*pck.DatabaseItems
}

func (s *items) FilterByPrice(price float64) []pck.Item {   
	result := make([]pck.Item, 0)
	for _, item := range s.Items {      
		if item.Price <= price {result = append(result, item)}
	}   
	return result
 }