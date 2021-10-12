package model
// the label for available currencies
type Label string

const (
	One Label = "ONE"
	Two Label = "TWO"
	Five Label = "FIVE"
	Ten Label = "TEN"
	Fifty Label = "FIFTY"
	Hundred Label = "Hundred"
)

// struct contains information about each currency bill available
type CoinInfo struct {
	Label Label
	Value int
}

// the currently available coins/bills in the machine (assuming the count of each bill to be infinite)
var AvailableCoins = map[int]*CoinInfo{
	1: {
		Label: One,
		Value: 1,
	},
	2: {
		Label: Two,
		Value: 2,
	},
	5: {
		Label: Five,
		Value: 5,
	},
	10: {
		Label: Ten,
		Value: 10,
	},
}

