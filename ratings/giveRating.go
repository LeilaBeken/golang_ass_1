package ratings

import(
	"github.com/LeilaBeken/golang_ass_1/pck"
)

type items struct{
	*pck.DatabaseItems
}

func (s *items) GiveRating(item pck.Item, rating float64) {   
	for i, it := range s.Items {
		if it.Name == item.Name { 
			s.Items[i].Rating = rating
			break     
		}
    }
}