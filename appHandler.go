package main

import (
	"database/sql"
	"fmt"

	"./lib/databasemethods"
)

//Page Used for sending Json from database to
type Page struct {
}
type PriceIncrease struct {
	vendorid   string
	oldprice   float32
	newprice   float32
	vendorname string
}

//LoadMainPage Searches the databases for price discrepencies and Sends payload to mobile app
func LoadMainPage() {

}

//CheckForPriceChange in DB for a user
func CheckForPriceChance(db *sql.DB) {
	singleAccount := "75fae0a3-6286-47b2-8cad-9b97a3a1d982"
	//singleCustomer := "3cc6a3ee-c820-4afe-8d3a-00a251739287"
	statement, _ := db.Prepare("SELECT ammount, vendorid, date FROM transactions WHERE accountid =?")
	rows, _ := statement.Query(singleAccount)
	var amount float32
	var vendorid string
	var date string
	var vendorPrice = make(map[string]float32)
	var increasedPrice = make(map[string]PriceIncrease)
	for rows.Next() {
		rows.Scan(&amount, &vendorid, &date)
		i, ok := vendorPrice[vendorid]
		//if vendor exits in map then it check to see if the price is equal to the old value
		if ok == true {
			if i != amount {
				_, vendorName := databasemethods.GetVendorByID(db, vendorid)
				p := PriceIncrease{vendorid: vendorid, oldprice: i, newprice: amount, vendorname: vendorName}
				increasedPrice[vendorName] = p
			}
		} else {
			vendorPrice[vendorid] = amount
		}
		fmt.Println("Amount: " + fmt.Sprintf("%f", amount) + " vendorID: " + vendorid + " date: " + date)
	}

	for key, value := range increasedPrice {
		fmt.Println("Key:", key, "Value:", value)
	}
}

//TODO Make a demo handler that can be called remotely to trigger alerts in the app
