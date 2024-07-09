package entities

type PaymentMethod int

const (
	Cash PaymentMethod = iota
	CreditCard
)

type Order struct {
	Items       []Item
	PaymentType PaymentMethod
}
