package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//Page Used for sending Json from database to
type Page struct {
	UserName        string
	AveragePerMonth float32
	Increased       []PriceIncrease
	Recurring       []Transactions
}
type PriceIncrease struct {
	Vendorid   string
	Oldprice   float32
	Newprice   float32
	Vendorname string
}
type Transactions struct {
	Vendorname string
	Price      float32
}

var Database *sql.DB

//LoadMainPage Searches the databases for price discrepencies and Sends payload to mobile app
func LoadMainPage(w http.ResponseWriter, r *http.Request) {

	page := CheckForPriceChance()
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(page); err != nil {
		panic(err)
	}

}
func PassContext(db *sql.DB) {
	Database = db
}

//CheckForPriceChange in DB for a user
func CheckForPriceChance() Page {
	var alltransactions = []Transactions{}
	var allincreases = []PriceIncrease{}
	singleAccount := "75fae0a3-6286-47b2-8cad-9b97a3a1d982"
	singleCustomer := "3cc6a3ee-c820-4afe-8d3a-00a251739287"
	_, userName, _ := GetCustomerByCustomerID(Database, singleCustomer)
	statement, _ := Database.Prepare("SELECT ammount, vendorid, date FROM transactions WHERE accountid =?")
	rows, _ := statement.Query(singleAccount)
	var monthlyaverage float32
	var amount float32
	var vendorid string
	var date string
	var vendorPrice = make(map[string]float32)
	var increasedPrice = make(map[string]PriceIncrease)
	for rows.Next() {
		rows.Scan(&amount, &vendorid, &date)
		_, vendorName := GetVendorByID(Database, vendorid)
		i, ok := vendorPrice[vendorName]
		//if vendor exits in map then it check to see if the price is equal to the old value
		if ok == true {
			if i != amount {

				p := PriceIncrease{Vendorid: vendorid, Oldprice: i, Newprice: amount, Vendorname: vendorName}
				increasedPrice[vendorName] = p
				allincreases = append(allincreases, p)
			}
			vendorPrice[vendorName] = amount
		} else {
			vendorPrice[vendorName] = amount
		}
		fmt.Println("Amount: " + fmt.Sprintf("%f", amount) + " vendorID: " + vendorid + " date: " + date)
	}

	for key, value := range increasedPrice {
		fmt.Println("Key:", key, "Value:", value)
	}

	//Not enough time to write logic to determine frequency so will be hard coding values
	for key, value := range vendorPrice {
		trans := Transactions{Vendorname: key, Price: value}
		alltransactions = append(alltransactions, trans)
		println(key)
		println(value)
		monthlyaverage += value
	}
	println(monthlyaverage)
	var page = Page{UserName: userName, AveragePerMonth: monthlyaverage, Increased: allincreases, Recurring: alltransactions}
	return page
}

//TODO Make a demo handler that can be called remotely to trigger alerts in the app
