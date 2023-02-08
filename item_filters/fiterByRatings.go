package itemfilters

import(
	"github.com/LeilaBeken/golang_ass_1/pck"
)


func (s *items) FilterByRatings(rating float64) []pck.Item {   
	result := make([]pck.Item, 0)
	for _, item := range s.Items {      
		if item.Rating <= rating {result = append(result, item)}
	}   
	return result
 }