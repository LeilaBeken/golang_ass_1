package pck

type Item struct {
	Name   string
	Price  int
	Rating int
}
type DatabaseItems struct {
	Items []Item
}

type AccessToItems interface {
	GetListOfItems() []string
}

func (items *DatabaseItems) GetListOfItems() []string {
	var list []string
	for _, item := range items.Items {
		list = append(list, item.Name)
	}
	return list
}
