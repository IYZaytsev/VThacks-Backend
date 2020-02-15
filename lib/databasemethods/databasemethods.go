package databasemethods

import (
	"database/sql"
	"fmt"
	"strconv"

	guuid "github.com/google/uuid"
)

func genUUID() string {
	id := guuid.New()
	return id.String()
}
func GetAllTransActions(datab *sql.DB) {
	//CREATE TABLE IF NOT EXISTS transactions (transactionID TEXT PRIMARY KEY, ammount int, vendorid TEXT, accountid TEXT)
	rows, _ := datab.Query("SELECT transactionID, ammount, vendorid, accountid FROM transactions")
	var tid string
	var ammount int
	var vendorid string
	var accountid string
	for rows.Next() {
		rows.Scan(&tid, &ammount, &vendorid, &accountid)
		fmt.Println("VendorId: " + tid + " Ammount: " + strconv.Itoa(ammount) + " VendorId: " + vendorid + " AccountId:" + accountid)
	}
}

//updates account balance
func GetALlVendors(datab *sql.DB) {
	rows, _ := datab.Query("SELECT vendorid, vendorname FROM vendors")
	var id string
	var vendorname string
	for rows.Next() {
		rows.Scan(&id, &vendorname)
		fmt.Println("VendorID: " + id + " VendorName: " + vendorname)
	}
}
func GetVendorByID(datab *sql.DB, vendorID string) (string, string) {
	statement, _ := datab.Prepare("SELECT vendorid, vendorName FROM vendors WHERE vendorid =?")
	rows, _ := statement.Query(vendorID)
	var id string
	var vendorName string
	for rows.Next() {
		rows.Scan(&id, &vendorName)
		fmt.Println("VendorID: " + id + " VendorName: " + vendorName)
	}
	return id, vendorName
}
func UpdateAccount(datab *sql.DB, accountID string, deltaBalance int) {
	statement, _ := datab.Prepare("UPDATE accounts SET balance =? WHERE accountid =?")
	statement.Exec(deltaBalance, accountID)

}
func ChargeAccount(datab *sql.DB, accountID string, chargeAmount int, vendorID string) {
	_, _, _, ballance := GetAccountByAccountID(datab, accountID)
	newBalance := ballance - chargeAmount
	UpdateAccount(datab, accountID, newBalance)
	statement, _ := datab.Prepare("INSERT INTO transactions (transactionid, accountid, vendorid, ammount) VALUES (?,?,?,?)")
	statement.Exec(genUUID(), accountID, vendorID, chargeAmount)

}
func GetAccountByCustomerID(datab *sql.DB, customerID string) (string, string, string, int) {
	var id string
	var aID string
	var atype string
	var balance int
	statement, _ := datab.Prepare("SELECT accountid, customerid, type, balance FROM accounts WHERE customerid =?")
	rows, _ := statement.Query(customerID)
	for rows.Next() {
		rows.Scan(&aID, &id, &atype, &balance)
		fmt.Println("AccountID:" + aID + " CustomerID:" + id + " Type: " + atype + " Balance: " + strconv.Itoa(balance))
	}
	return id, aID, atype, balance
}
func GetAccountByAccountID(datab *sql.DB, accountID string) (string, string, string, int) {
	var id string
	var aID string
	var atype string
	var balance int
	statement, _ := datab.Prepare("SELECT accountid, customerid, type, balance FROM accounts WHERE accountid =?")
	rows, _ := statement.Query(accountID)
	for rows.Next() {
		rows.Scan(&aID, &id, &atype, &balance)
		//fmt.Println("AccountID:" + aID + ": " + "CustomerID:" + id + " Type: " + atype + " Balance: " + strconv.Itoa(balance))
	}
	return id, aID, atype, balance
}
func GetCustomerByID(datab *sql.DB, customerID string) (string, string, string) {
	statement, _ := datab.Prepare("SELECT customerid, firstname, lastname FROM customers WHERE customerid =?")
	rows, _ := statement.Query(customerID)
	var id string
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println("CustomerID: " + id + " FirstName: " + firstname + " LastName: " + lastname)
	}
	return id, firstname, lastname
}
func PrintAllCustomers(datab *sql.DB) {
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
func PrintAllAccounts(datab *sql.DB) {
	var id string
	var aID string
	var atype string
	var balance int
	rows, _ := datab.Query("SELECT accountid, customerid, type, balance FROM accounts")
	for rows.Next() {
		rows.Scan(&aID, &id, &atype, &balance)
		fmt.Println("AccountID:" + aID + ": " + "CustomerID:" + id + " Type: " + atype + " Balance: " + strconv.Itoa(balance))
	}
}

func ResetScheme(datab *sql.DB) {
	statement, _ := datab.Prepare("CREATE TABLE IF NOT EXISTS customers (customerid TEXT PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = datab.Prepare("CREATE TABLE IF NOT EXISTS accounts (accountid TEXT PRIMARY KEY,customerid INTEGER, type TEXT, balance INTEGER)")
	statement.Exec()
	statement, _ = datab.Prepare("CREATE TABLE IF NOT EXISTS vendors (vendorid TEXT PRIMARY KEY, vendorName TEXT )")
	statement.Exec()
	statement, _ = datab.Prepare("CREATE TABLE IF NOT EXISTS transactions (transactionID TEXT PRIMARY KEY, ammount int, vendorid TEXT, accountid TEXT)")
	statement.Exec()

}
