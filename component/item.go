package component

type Item struct {
	t     ItemType
	gotAt uint32
}

type ItemType string

const (
	SHIELD ItemType = "shield"
	SWORD  ItemType = "sword"
	CURSE  ItemType = "curse"
	FAN    ItemType = "fan"
	STAFF  ItemType = "staff"
	QUILL  ItemType = "quill"

	// secret items
	TRUE_CURSE  ItemType = "true_curse"
	FALSE_CURSE ItemType = "false_curse"
)

func NewItem(t ItemType, gotAt uint32) *Item {
	return &Item{t: t, gotAt: gotAt}
}
