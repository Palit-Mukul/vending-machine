package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vending-machine/model"
)

// make the available information globally available
var newVendingMachine *model.VendingMachine

func initVendingMachine() {
	newVendingMachine = model.NewVendingMachine()
}

// function the display all the available items in the vending machine
func DisplayItemList(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: Item Page")
	SetupResponse(&w,r)

	w.Header().Set("Content-Type", "application/json")

	// encode the response in the given structure
	err := json.NewEncoder(w).Encode(newVendingMachine.Items)
	if err != nil {
		fmt.Fprintf(w,"could not encode json : %v", err)
		log.Fatalf("could not encode json : %v", err)
	}
}

// function the display all the available currencies/bills in the vending machine
func DisplayAvailableCurrency(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: Available Currency Page")
	SetupResponse(&w,r)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(newVendingMachine.AvailableCurrency)
	if err != nil {
		fmt.Fprintf(w,"could not encode json : %v", err)
		log.Fatalf("could not encode json : %v", err)
	}
}

// function to extract an item from the machine
func GetItemFromMachine(w http.ResponseWriter, r *http.Request){
	SetupResponse(&w,r)
	var newRequest model.ItemRequest
	newVendingMachine.TotalChangeDue=0
	// read the item needed and money inserted into the machine
	err := json.NewDecoder(r.Body).Decode(&newRequest)
	if err != nil {
		fmt.Fprintf(w,"Could not make request for item : %v", err)
		log.Fatalf("Could not make request for item : %v", err)
	}

	// item not available in the machine
	if _, found := newVendingMachine.Items[newRequest.Name]; !found {
		fmt.Println("Hello")
		fmt.Fprintf(w, "Item not available\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// item is available and the quantity is not 0
	if newVendingMachine.Items[newRequest.Name].Count > 0 {
		// the amount inserted is compared with the item price
		if newRequest.Amount > newVendingMachine.Items[newRequest.Name].Price {
			newVendingMachine.TotalMoney += newVendingMachine.Items[newRequest.Name].Price

			// calculate the currencies/bills to be given to the customer
			CalculateChange(newRequest.Amount-newVendingMachine.Items[newRequest.Name].Price)
			// reduce the item count by 1
			newVendingMachine.Items[newRequest.Name].Count--
		}else if newRequest.Amount == newVendingMachine.Items[newRequest.Name].Price {
			newVendingMachine.TotalMoney += newRequest.Amount
			newVendingMachine.Items[newRequest.Name].Count--
		}else{
			fmt.Fprintf(w, "Inserted amount is less than the item value\n")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}else {
		fmt.Fprintf(w, "No more items for [%s] are available in machine\n", newRequest.Name)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newVendingMachine.TotalChangeDue = newRequest.Amount - newVendingMachine.Items[newRequest.Name].Price
	w.Header().Set("Content-Type", "application/json")

	// display the state of vending machine after every request
	err = json.NewEncoder(w).Encode(newVendingMachine)
	if err != nil {
		fmt.Fprintf(w,"could not encode json : %v", err)
		log.Fatalf("could not encode json : %v", err)
	}
}

func CalculateChange(amountDue int) {
	newVendingMachine.MoneyToReturn = map[int]int{}
	// money is returned starting from the highest available bill i.e 10 in our case
	totalTenValueCoins := amountDue/10
	totalFiveValueCoins := (amountDue - (totalTenValueCoins*10))/5
	totalTwoValueCoins := (amountDue - (totalTenValueCoins*10 + totalFiveValueCoins*5))/2
	totalOneValueCoins := (amountDue - (totalTenValueCoins*10 + totalFiveValueCoins*5 + totalTwoValueCoins*2))

	newVendingMachine.MoneyToReturn[10]=totalTenValueCoins
	newVendingMachine.MoneyToReturn[5]=totalFiveValueCoins
	newVendingMachine.MoneyToReturn[2]=totalTwoValueCoins
	newVendingMachine.MoneyToReturn[1]=totalOneValueCoins
}

// function the update the item list in the vending machine
func AddItemToMachine(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w,r)
	var newAddition model.ItemAdditionRequest
	err := json.NewDecoder(r.Body).Decode(&newAddition)
	if err != nil {
		fmt.Fprintf(w,"Could not make request for item : %v", err)
		log.Fatalf("Could not make request for item : %v", err)
	}
	fmt.Println(newAddition)
	itemSection := model.New()
	err = itemSection.AddItem(newAddition.Name,newAddition.Count,newAddition.Price)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newVendingMachine.Items)
	if err != nil {
		fmt.Fprintf(w,"could not encode json : %v", err)
		log.Fatalf("could not encode json : %v", err)
	}
}

// route the request to the respective handler functions
func handleRequests() {
	http.HandleFunc("/itemlist", DisplayItemList)
	http.HandleFunc("/availablecurrency", DisplayAvailableCurrency)
	http.HandleFunc("/getitem", GetItemFromMachine)
	http.HandleFunc("/additem", AddItemToMachine)
	log.Fatal(http.ListenAndServe(":8009", nil))
}

func main() {
	log.Println("Starting Vending Machine")
	// initial the vending machine
	initVendingMachine()
	handleRequests()
}

//SetupResponse : to setup access control on requests
func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*req).Method == "OPTIONS" {
		(*w).Header().Set("Access-Control-Max-Age", "86400")
		(*w).WriteHeader(http.StatusOK)
		return
	}
}
