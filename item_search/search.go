package itemsearch

import (
	"strings"

	"github.com/LeilaBeken/golang_ass_1/pck"
)

type Items struct{
	*pck.DatabaseItems
}

func (items *Items)ItemSearch(it string) *pck.Item{
	for _, i := range items.Items{
		if strings.Contains(i.Name, it) {
			return &i
		}
	}
	return nil
}