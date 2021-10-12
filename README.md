# Vending Machine

Vending Machine is an API based system. The items present in the vending machine are limited although new items can be added to the machine. The currency available in the machine to return are in the bills of `1,2,5 and 10`.The total currencies for each bill are considered to be infinite. The total purchase made is stored in the machine as `TotalAmount`.

The following API end points are available :

* `/itemlist` : `GET` operation to get the list of available items in the machine
* `/availablecurrency` : `GET` operation to get all the available bills in the system
* `/getitem` : `POST` operation to extract an item from the machine. The item to be extracted needs to be in the request body as a `json Object` in the following format `{"Name" : "<NAME OF THE ITEM>", "Amount" : "<AMOUNT YOU ARE INSERTING>"}`
* `/additem` : `POST` operation to add a new item to the vending machine or increase the quantity of an existing item.

```
NOTE : After the end of each POST request the final state of the vending machine is returned containing the total amount
in the machine and the money to be returned to the customer in the bills of 1,2,5 and 10.

```

## How to run
In the vending-machine directory run :

`go run main.go`

The API will be available for requests at `http://localhost:8009/`