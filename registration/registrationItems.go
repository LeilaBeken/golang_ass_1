package registration

import (
	"github.com/LeilaBeken/golang_ass_1/pck"
)

type dbItems struct {
	*pck.DatabaseItems
}

func (d *dbItems) RegisterItem(item pck.Item) {
	d.Items = append(d.Items, item)
}