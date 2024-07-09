package entities

type Item struct {
	Name  string
	Price float64
}

type MenuSection struct {
	Name  string
	Items []Item
}

type Menu struct {
	Sections []MenuSection
}
