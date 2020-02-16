package lib

import (
	"database/sql"
	"fmt"
	"strconv"

	guuid "github.com/google/uuid"
)

//GenUUID Generates a unique ID
func GenUUID() string {
	id := guuid.New()
	return id.String()
}

//GetAllTransActions gets transactions
func GetAllTransActions(datab *sql.DB) {
	//CREATE TABLE IF NOT EXISTS transactions (transactionID TEXT PRIMARY KEY, ammount int, vendorid TEXT, accountid TEXT)
	rows, _ := datab.Query("SELECT transactionID, ammount, vendorid, accountid, date FROM transactions")
	var tid string
	var ammount float32
	var vendorid string
	var accountid string
	var date string
	for rows.Next() {
		rows.Scan(&tid, &ammount, &vendorid, &accountid, &date)
		fmt.Println("TransactionID: " + tid + " Ammount: " + fmt.Sprintf("%f", ammount) + " VendorId: " + vendorid + " AccountId:" + accountid + " Date: " + date)
	}
}

//GetALlVendors gets all vendors
func GetALlVendors(datab *sql.DB) {
	rows, _ := datab.Query("SELECT vendorid, vendorname FROM vendors")
	var id string
	var vendorname string
	for rows.Next() {
		rows.Scan(&id, &vendorname)
		fmt.Println("VendorID: " + id + " VendorName: " + vendorname)
	}
}

//GetVendorByID gets all vendors by ID
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

//UpdateAccount updates accounts by ID to chance balance
func UpdateAccount(datab *sql.DB, accountID string, deltaBalance float32) {
	statement, _ := datab.Prepare("UPDATE accounts SET balance =? WHERE accountid =?")
	statement.Exec(deltaBalance, accountID)

}

//ChargeAccount charges an account
func ChargeAccount(datab *sql.DB, accountID string, chargeAmount float32, vendorID string, date string) {
	//(transactionID TEXT PRIMARY KEY, ammount REAL, vendorid TEXT, accountid TEXT, date TEXT)
	_, _, _, ballance := GetAccountByAccountID(datab, accountID)
	println(ballance)
	newBalance := ballance - chargeAmount
	println(newBalance)
	println(accountID)
	println(chargeAmount)
	println(vendorID)
	println(date)
	UpdateAccount(datab, accountID, newBalance)
	statement, _ := datab.Prepare("INSERT INTO transactions (transactionID, accountid, vendorid, ammount, date) VALUES (?,?,?,?,?)")
	rslt, _ := statement.Exec(GenUUID(), accountID, vendorID, chargeAmount, date)
	rowaff, _ := rslt.RowsAffected()
	println(rowaff)

}

//GetAccountByCustomerID gets an account by customer id
func GetAccountByCustomerID(datab *sql.DB, customerID string) (string, string, string, float32) {
	var id string
	var aID string
	var atype string
	var balance float32
	statement, _ := datab.Prepare("SELECT accountid, customerid, type, balance FROM accounts WHERE customerid =?")
	rows, _ := statement.Query(customerID)
	for rows.Next() {
		rows.Scan(&aID, &id, &atype, &balance)
		fmt.Println("AccountID:" + aID + " CustomerID:" + id + " Type: " + atype + " Balance: " + fmt.Sprintf("%f", balance))
	}
	return id, aID, atype, balance
}

//GetAccountByAccountID gets an account by account id
func GetAccountByAccountID(datab *sql.DB, accountID string) (string, string, string, float32) {
	var id string
	var aID string
	var atype string
	var balance float32
	statement, _ := datab.Prepare("SELECT accountid, customerid, type, balance FROM accounts WHERE accountid =?")
	rows, _ := statement.Query(accountID)
	for rows.Next() {
		rows.Scan(&aID, &id, &atype, &balance)
		//fmt.Println("AccountID:" + aID + ": " + "CustomerID:" + id + " Type: " + atype + " Balance: " + fmt.Sprintf("%f", balance))
	}
	return id, aID, atype, balance
}

//GetCustomerByCustomerID gets customer by Customer id
func GetCustomerByCustomerID(datab *sql.DB, customerID string) (string, string, string) {
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

//PrintAllCustomers prints all customers to console
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

//PrintAllAccounts prints all accounts to console
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

//ResetScheme Used to reset the scheme
func ResetScheme(datab *sql.DB) {
	statement, _ := datab.Prepare("CREATE TABLE IF NOT EXISTS customers (customerid TEXT PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = datab.Prepare("CREATE TABLE IF NOT EXISTS accounts (accountid TEXT PRIMARY KEY,customerid INTEGER, type TEXT, balance REAL)")
	statement.Exec()
	statement, _ = datab.Prepare("CREATE TABLE IF NOT EXISTS vendors (vendorid TEXT PRIMARY KEY, vendorName TEXT )")
	statement.Exec()
	statement, _ = datab.Prepare("CREATE TABLE IF NOT EXISTS transactions (transactionID TEXT PRIMARY KEY, ammount REAL, vendorid TEXT, accountid TEXT, date TEXT)")
	statement.Exec()

}
