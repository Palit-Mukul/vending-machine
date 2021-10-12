package model

import "fmt"

// Item represents all the available products in the vending machine
type ItemInfo struct {
	Count int
	Price int
}

// the map represent the already present items in the vending machine
var AvailableItems = map[string]*ItemInfo {
	"Chocolate" : {
		Count: 5,
		Price: 10,
	},
	"Chips": {
		Count: 30,
		Price: 20,
	},
	"ColdDrink" : {
		Count: 40,
		Price: 15,
	},
	"CupCake": {
		Count:40,
		Price: 25,
	},
}

// interface to define function specific to items section of vending machine
type ItemInterface interface {
	AddItem(string, int, int) error
}

type Item struct {
	Item map[string]*ItemInfo
}

// creates a new instance to add the items in the currently available item list
func New() ItemInterface {
	return &Item{
		Item: AvailableItems,
	}
}

// function to add new item or update the quantity of already existing item
func (n *Item) AddItem (name string, count int, price int) error {
	if _, found := n.Item[name]; !found {
		n.Item[name]= &ItemInfo{
			Count: count,
			Price: price,
		}
		return nil
		}
		if n.Item[name].Price == price {
			n.Item[name].Count = n.Item[name].Count+count
			return nil
		}
		return fmt.Errorf("the item is already available in the machine but the price of the item has changed")
}
