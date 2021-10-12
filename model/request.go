package model

// request format for getting an item from the vending machine
type ItemRequest struct {
	Name string
	Amount int
}

// request format for putting new items or updating items in the vending machine
type ItemAdditionRequest struct {
	Name string
	Count int
	Price int
}
