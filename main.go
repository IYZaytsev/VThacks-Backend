package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	lib "github.com/IYZaytsev/VThacks-Backend/lib"
	_ "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

//Database is used within the whole application

func main() {
	database, _ := sql.Open("sqlite3", "./recurT.db")
	lib.PassContext(database)
	fmt.Println("starting router")
	router := NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", router))
	/*
		statement, _ := database.Prepare("INSERT INTO customers (customerid, firstname, lastname) VALUES (?,?,?)")
		customerID := databasemethods.GenUUID()
		accountID := databasemethods.GenUUID()
		statement.Exec(customerID, "John", "Raboy")
		statement, _ = database.Prepare("INSERT INTO accounts (accountid, customerid, type, balance) VALUES (?,?,?,?)")
		statement.Exec(accountID, customerID, "Checking", 1000)
	*/
	//vendorid TEXT PRIMARY KEY, vendorName TEXT )"
	/*vendorID := databasemethods.GenUUID()
	statement, _ := database.Prepare("INSERT INTO vendors (vendorid, vendorName) VALUES (?,?)")
	statement.Exec(vendorID, "Target")
	*/

	//statement, _ := database.Prepare("DROP TABLE transactions")
	//statement.Exec()
	//statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS transactions (transactionID TEXT PRIMARY KEY, ammount REAL, vendorid TEXT, accountid TEXT, date TEXT)")
	//statement.Exec()
	//databasemethods.GetALlVendors(Database)
	//databasemethods.GetAllTransActions(Database)

	//log.Fatal(http.ListenAndServe(":8080", router))
}
