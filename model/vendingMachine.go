package model

// the currently available features in the vending machine
type VendingMachine struct{
 	Items map[string]*ItemInfo
 	// list of available currencies
 	AvailableCurrency map[int]*CoinInfo
 	// the map of each bill type to be returned to the customer
 	MoneyToReturn map[int]int
 	// total amount in the machine
 	TotalMoney int
 	// the total amount to be returned to the customer
 	TotalChangeDue int
}

// initialize the vending machine
func NewVendingMachine () *VendingMachine{
	return &VendingMachine{
		Items: AvailableItems,
		AvailableCurrency: AvailableCoins,
		MoneyToReturn: make(map[int]int),
		TotalMoney: 0,
	}
}