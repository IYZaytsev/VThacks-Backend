package main

import (
	"database/sql"
	"fmt"
	"strconv"

	guuid "github.com/google/uuid"
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
	getAccountByCustomerID(database, "4fc541f0-91f4-49df-9fb2-c968d705cbd7")
	//log.Fatal(http.ListenAndServe(":8080", router))
}
func genUUID() string {
	id := guuid.New()
	return id.String()
}

//updates account balance
func updateAccount(datab *sql.DB, accountID string, deltaBalance int) {
	statement, _ := datab.Prepare("UPDATE accounts SET balance =? WHERE accountid =?")
	res, err := statement.Exec(deltaBalance, accountID)
	rows, _ := res.RowsAffected()
	println((rows))
	println(err)
}
func getAccountByCustomerID(datab *sql.DB, customerID string) {
	var id string
	var aID string
	var atype string
	var balance int
	statement, _ := datab.Prepare("SELECT accountid, customerid, type, balance FROM accounts WHERE customerid =?")
	rows, _ := statement.Query(customerID)
	for rows.Next() {
		rows.Scan(&aID, &id, &atype, &balance)
		fmt.Println("AccountID:" + aID + ": " + "CustomerID:" + id + "Type: " + atype + " Balance: " + strconv.Itoa(balance))
	}
}
func printAllCustomers(datab *sql.DB) {
	rows, _ := datab.Query("SELECT customerid, firstname, lastname FROM customers")
	var id string
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println("CustomerID: " + id + " FirstName: " + firstname + " LastName: " + lastname)
	}

}

// id is customer id, aID is account id
func printAllAccounts(datab *sql.DB) {
	var id string
	var aID string
	var atype string
	var balance int
	rows, _ := datab.Query("SELECT accountid, customerid, type, balance FROM accounts")
	for rows.Next() {
		rows.Scan(&aID, &id, &atype, &balance)
		fmt.Println("AccountID:" + aID + ": " + "CustomerID:" + id + "Type: " + atype + " Balance: " + strconv.Itoa(balance))
	}
}

func resetScheme(datab *sql.DB) {
	statement, _ := datab.Prepare("CREATE TABLE IF NOT EXISTS customers (customerid TEXT PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = datab.Prepare("CREATE TABLE IF NOT EXISTS accounts (accountid TEXT PRIMARY KEY,customerid INTEGER, type TEXT, balance INTEGER)")
	statement.Exec()
}
