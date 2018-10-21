// jp.go
//
// sequental edition
//
// Version: 0.1.
// Date: 2017-05-13.
// Author: McArcher.
//
// This is a test Program.
// It does not use any Optimizations or any Caching.

package main

import "fmt"
import "log"
import "flag"
import "time"
import "strconv"
import "database/sql"

// Global Parameters

const table_users = "Users"
const table_jp = "JP"
const column_userName = "UserName"
const column_persSum = "PersonalSum"
const column_JPSum = "JPSum"

var db *sql.DB

//--------------------------------------------------------------------------------

func main() {

	var db_host, db_user, db_user_pwd, db_database, db_ssl_mode string
	var db_port, test_count, i int
	var flag_create_tables_ptr, flag_delete_tables_ptr *bool
	var time_1, time_2 time.Time

	flag_create_tables_ptr = flag.Bool("ct", false, "Create Tables.")
	flag_delete_tables_ptr = flag.Bool("dt", false, "Delete Tables.")

	db_host = "localhost"
	db_port = 5432
	db_database = "test"
	db_ssl_mode = "disable"

	db_user = "test"
	db_user_pwd = "test"

	test_count = 100

	// Flags
	flag.Parse()

	// Connect to DB
	db = db_connect_pg(db_host, db_port, db_user, db_user_pwd, db_database,
		db_ssl_mode)

	// Delete Tables
	if *flag_delete_tables_ptr {

		log.Println("Deleteing Tables...") //
		db_tables_delete(db, table_users, table_jp)
	}

	// Create Tables
	if *flag_create_tables_ptr {

		log.Println("Creating Tables...") //
		db_tables_init(db, table_users, table_jp, column_userName,
			column_persSum, column_JPSum)
	}

	if !db_tables_exist(db, table_users, table_jp) {

		log.Println("Tables do not exist!") //
		return
	}

	log.Println("Starting Job...") //
	time_1 = time.Now()

	for i = 1; i <= test_count; i++ {

		addBet("Петя", 20.5, 5.1)
		addBet("Коля", 30, 10)
		addBet("Катя", 50, 40)
		addBet("Петя", 7, 7)
		addBet("Иннокентий", 13, 10)
		addBet("Иоанн Васильевич", 23.45, 19.78)
		addBet("Дмитрий Иванович Менделеев", 16.34, 12.83)
		addBet("John Smith", 28.3, 19.90)
		addBet("John Lennon", 90, 89)
	}

	time_2 = time.Now()
	fmt.Println("Time spent (sec.):", time_2.Sub(time_1).Seconds()) //

	// Disconnect from DB
	db_close(db)

}

//--------------------------------------------------------------------------------

func addBet(userName string, client_total_sum, client_jp_sum float64) bool {

	// Saves the Bet in DataBase.
	// This Function must be run solely!

	var userExists, result bool
	var db_cmd string
	var reply *string
	var err error
	var client_personal_sum, old_personal_sum, new_personal_sum float64
	var old_jp_sum, new_jp_sum float64

	reply = new(string)

	// Fool-check
	if (client_total_sum < 0) || (client_jp_sum < 0) ||
		(client_jp_sum > client_total_sum) || (len(userName) == 0) {

		log.Println("Bad Parameters") //
		return false
	}

	// Personal Sum Delta
	client_personal_sum = client_total_sum - client_jp_sum

	// Create User if it does not exist
	userExists = user_exists(db, userName)
	if !userExists {
		add_user(db, userName)
	}

	// Start Transaction
	db_cmd = "START TRANSACTION;"
	result = db_action(db, db_cmd)
	if !result {
		return false
	}

	// Read Sum of JP
	db_cmd = fmt.Sprintf("SELECT %s FROM %s;", column_JPSum, table_jp)
	result = db_query_value(db, db_cmd, reply)
	if !result {
		db_transaction_rollback(db)
		return false
	}

	// String -> Float
	old_jp_sum, err = strconv.ParseFloat(*reply, 64)
	if err != nil {
		log.Println(err) //
		db_transaction_rollback(db)
		return false
	}

	// Read old Personal Sum
	db_cmd = fmt.Sprintf("SELECT %s FROM %s WHERE %s='%s';",
		column_persSum, table_users, column_userName, userName)
	result = db_query_value(db, db_cmd, reply)
	if !result {
		db_transaction_rollback(db)
		return false
	}

	// String -> Float
	old_personal_sum, err = strconv.ParseFloat(*reply, 64)
	if err != nil {
		log.Println(err) //
		db_transaction_rollback(db)
		return false
	}

	// Calculate new Sums
	new_personal_sum = old_personal_sum + client_personal_sum
	new_jp_sum = old_jp_sum + client_jp_sum

	// Save (update) User's Sum
	db_cmd = fmt.Sprintf("UPDATE %s SET %s = %f WHERE %s='%s';",
		table_users, column_persSum, new_personal_sum, column_userName, userName)
	result = db_action(db, db_cmd)
	if !result {
		db_transaction_rollback(db)
		return false
	}

	// Save (update) JP Sum
	db_cmd = fmt.Sprintf("UPDATE %s SET %s = '%f';", table_jp, column_JPSum,
		new_jp_sum)
	result = db_action(db, db_cmd)
	if !result {
		db_transaction_rollback(db)
		return false
	}

	// Commit Transaction
	db_cmd = "COMMIT TRANSACTION;"
	result = db_action(db, db_cmd)
	if !result {
		db_transaction_rollback(db)
		return false
	}

	return result
}

//--------------------------------------------------------------------------------

func user_exists(db *sql.DB, name string) bool {

	// Checks if UserName exists in Users Table.

	var db_qr string
	var qr_row *sql.Row
	var err error
	var column, table, reply string

	column = column_userName
	table = table_users

	db_qr = fmt.Sprintf("SELECT %s FROM %s WHERE %s='%s';", column, table,
		column, name)
	qr_row = db.QueryRow(db_qr)
	err = qr_row.Scan(&reply)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Println(err) //
			return false
		}
	}

	return true
}

//--------------------------------------------------------------------------------

func add_user(db *sql.DB, name string) {

	// Adds new UserName to Users Table.

	db_table_value_insert(db, table_users, column_userName, name)
}

//--------------------------------------------------------------------------------
