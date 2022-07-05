package order

type Points struct {
	Address       string  `json:"address"`
	ContactPerson Contact `json:"contact_person"`
}
