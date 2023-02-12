package ratings

import(
	"github.com/LeilaBeken/golang_ass_1/pck"
)

type item struct{
	*pck.Item
}

func (s *item) GiveRating(rating int) {   
	s.Rating = rating
}