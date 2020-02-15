package main

import (
	"database/sql"

	"./lib/databasemethods"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./recurT.db")
	/*statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS customers (customerid TEXT PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS accounts (accountid TEXT PRIMARY KEY,customerid INTEGER, type TEXT, balance INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO customers (customerid, firstname, lastname) VALUES (?,?,?)")
	customerID := genUUID()
	accountID := genUUID()
	statement.Exec(customerID, "John", "Raboy")
	statement, _ = database.Prepare("INSERT INTO accounts (accountid, customerid, type, balance) VALUES (?,?,?,?)")
	statement.Exec(accountID, customerID, "Checking", 1000)
	*/
	//vendorid TEXT PRIMARY KEY, vendorName TEXT )"
	/*vendorID := genUUID()
	statement, _ := database.Prepare("INSERT INTO vendors (vendorid, vendorName) VALUES (?,?)")
	statement.Exec(vendorID, "Walmart")
	*/

	databasemethods.GetALlVendors(database)
	//chargeAccount(database, "e718c118-68be-4e29-9437-a66a0ce00a4e", 100, "03fdea46-3264-4e25-9e27-0fba0fd80182")
	databasemethods.GetAccountByCustomerID(database, "4fc541f0-91f4-49df-9fb2-c968d705cbd7")
	databasemethods.GetCustomerByID(database, "4fc541f0-91f4-49df-9fb2-c968d705cbd7")
	database.Prepare("DROP TABLE transactions")
	databasemethods.ResetScheme(database)
	databasemethods.GetAllTransActions(database)
	//log.Fatal(http.ListenAndServe(":8080", router))
}
