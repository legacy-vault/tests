// jp.go
//
// concurrent edition
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
import "strconv"
import "time"
import "database/sql"

type tUserManagerJob struct {
	userName      string
	returnChannel chan bool
}

type tBetManagerJob struct {
	userName           string
	personalSum, jpSum float64
	returnChannel      chan bool
}

// Global Parameters

const table_users = "Users"
const table_jp = "JP"
const column_userName = "UserName"
const column_persSum = "PersonalSum"
const column_JPSum = "JPSum"

var jobsTodo int

// Channels
var jobsChan chan bool
var userManagerChan chan tUserManagerJob
var betManagerChan chan tBetManagerJob

// Server Settings
var db_host, db_database, db_ssl_mode string
var db_port int
var db_user_admin, db_user_1, db_user_2 string
var db_pwd_admin, db_pwd_1, db_pwd_2 string

//--------------------------------------------------------------------------------

func main() {

	// Main.

	var test_count, i int
	var flag_create_tables_ptr, flag_delete_tables_ptr *bool
	var time_1, time_2 time.Time
	var db *sql.DB // Connection to DataBase from main Function

	flag_create_tables_ptr = flag.Bool("ct", false, "Create Tables.")
	flag_delete_tables_ptr = flag.Bool("dt", false, "Delete Tables.")

	db_host = "localhost"
	db_port = 5432
	db_database = "test"
	db_ssl_mode = "disable"

	db_user_admin = "test"
	db_pwd_admin = "test"
	db_user_1 = "test"
	db_pwd_1 = "test"
	db_user_2 = "test"
	db_pwd_2 = "test"

	test_count = 100

	// Flags
	flag.Parse()

	// Connect to DB
	db = db_connect_pg(db_host, db_port, db_user_admin, db_pwd_admin,
		db_database, db_ssl_mode)

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

	// Disconnect from DB
	db_close(db)

	log.Println("Starting Job...") //
	time_1 = time.Now()

	// Create Channels
	jobsChan = make(chan bool)
	userManagerChan = make(chan tUserManagerJob)
	betManagerChan = make(chan tBetManagerJob)

	// Start Managers
	go user_manager()
	go bet_manager()

	for i = 1; i <= test_count; i++ {

		// Do Action
		jobsTodo++
		go addBet("Петя", 20.5, 5.1)

		jobsTodo++
		go addBet("Коля", 30, 10)

		jobsTodo++
		go addBet("Катя", 50, 40)

		jobsTodo++
		go addBet("Петя", 7, 7)

		jobsTodo++
		go addBet("Иннокентий", 13, 10)

		jobsTodo++
		go addBet("Иоанн Васильевич", 23.45, 19.78)

		jobsTodo++
		go addBet("Дмитрий Иванович Менделеев", 16.34, 12.83)

		jobsTodo++
		go addBet("John Smith", 28.3, 19.90)

		jobsTodo++
		go addBet("John Lennon", 90, 89)

		// Wait for all Jobs to finish
		for jobsTodo > 0 {

			<-jobsChan
			jobsTodo--
		}
	}

	time_2 = time.Now()
	fmt.Println("Time spent (sec.):", time_2.Sub(time_1).Seconds()) //
}

//--------------------------------------------------------------------------------

func addBet(userName string, client_total_sum, client_jp_sum float64) {

	// Saves the Bet in DataBase.
	// This Function must be run solely!

	var client_personal_sum float64
	var umJob *tUserManagerJob
	var bmJob *tBetManagerJob
	var rcvChan chan bool
	var result bool

	// Create Channels
	rcvChan = make(chan bool)

	// Fool-check
	if (client_total_sum < 0) || (client_jp_sum < 0) ||
		(client_jp_sum > client_total_sum) || (len(userName) == 0) {

		log.Println("Bad Parameters") //
		jobsChan <- false
	}

	// Personal Sum Delta
	client_personal_sum = client_total_sum - client_jp_sum

	// Create User if it does not exist //================

	// Create Job
	umJob = new(tUserManagerJob)
	umJob.userName = userName
	umJob.returnChannel = rcvChan

	// Send Job
	userManagerChan <- *umJob

	// Get Result
	result = <-rcvChan

	// Place Bet //=======================================

	// Create Job
	bmJob = new(tBetManagerJob)
	bmJob.userName = userName
	bmJob.personalSum = client_personal_sum
	bmJob.jpSum = client_jp_sum
	bmJob.returnChannel = rcvChan

	// Send Job
	betManagerChan <- *bmJob

	// Get Result
	result = <-rcvChan

	// Jobs done //=======================================

	// Send Result
	jobsChan <- result
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

func user_manager() {

	// Checks if User exists and registers User if it does not exist.

	var job tUserManagerJob
	var userName string
	var userExists bool
	var db_um *sql.DB // Connection to DataBase from UserManager

	// Connect to DB
	db_um = db_connect_pg(db_host, db_port, db_user_1, db_pwd_1, db_database,
		db_ssl_mode)

	for true {

		// Get Job
		job = <-userManagerChan

		// Do Job
		userName = job.userName
		userExists = user_exists(db_um, userName)
		if !userExists {
			add_user(db_um, userName)
		}

		// Send Feedback
		job.returnChannel <- true
	}

	// Disconnect from DB
	db_close(db_um)
}

//--------------------------------------------------------------------------------

func bet_manager() {

	// Places new Bets.

	var job tBetManagerJob
	var userName, db_cmd string
	var result bool
	var reply *string
	var old_jp_sum, new_jp_sum, old_personal_sum, new_personal_sum float64
	var client_personal_sum, client_jp_sum float64
	var err error
	//var tx *sql.Tx
	var db_bm *sql.DB // Connection to DataBase from BetManager

	// Connect to DB
	db_bm = db_connect_pg(db_host, db_port, db_user_2, db_pwd_2, db_database,
		db_ssl_mode)

	for true {

		// Get Job
		job = <-betManagerChan

		// Do Job

		reply = new(string)

		// Unpack Data from Job
		client_personal_sum = job.personalSum
		client_jp_sum = job.jpSum
		userName = job.userName

		// Start Transaction
		db_cmd = "START TRANSACTION;"
		result = db_action(db_bm, db_cmd)
		if !result {
			job.returnChannel <- false
			continue
		}

		/*
			tx, err = db_bm.Begin()
			if err != nil {
				log.Println("tx:", tx) //
				log.Println(err)       //
				job.returnChannel <- false
				continue
			}
		*/

		// Read Sum of JP
		db_cmd = fmt.Sprintf("SELECT %s FROM %s;", column_JPSum, table_jp)
		result = db_query_value(db_bm, db_cmd, reply)
		if !result {
			db_transaction_rollback(db_bm)
			job.returnChannel <- false
			continue
		}

		// String -> Float
		old_jp_sum, err = strconv.ParseFloat(*reply, 64)
		if err != nil {
			log.Println(err) //
			db_transaction_rollback(db_bm)
			job.returnChannel <- false
			continue
		}

		// Read old Personal Sum
		db_cmd = fmt.Sprintf("SELECT %s FROM %s WHERE %s='%s';",
			column_persSum, table_users, column_userName, userName)
		result = db_query_value(db_bm, db_cmd, reply)
		if !result {
			db_transaction_rollback(db_bm)
			job.returnChannel <- false
			continue
		}

		// String -> Float
		old_personal_sum, err = strconv.ParseFloat(*reply, 64)
		if err != nil {
			log.Println(err) //
			db_transaction_rollback(db_bm)
			job.returnChannel <- false
			continue
		}

		// Calculate new Sums
		new_personal_sum = old_personal_sum + client_personal_sum
		new_jp_sum = old_jp_sum + client_jp_sum

		// Save (update) User's Sum
		db_cmd = fmt.Sprintf("UPDATE %s SET %s = %f WHERE %s='%s';",
			table_users, column_persSum, new_personal_sum, column_userName,
			userName)
		result = db_action(db_bm, db_cmd)
		if !result {
			db_transaction_rollback(db_bm)
			job.returnChannel <- false
			continue
		}

		// Save (update) JP Sum
		db_cmd = fmt.Sprintf("UPDATE %s SET %s = '%f';", table_jp, column_JPSum,
			new_jp_sum)
		result = db_action(db_bm, db_cmd)
		if !result {
			db_transaction_rollback(db_bm)
			job.returnChannel <- false
			continue
		}

		// Commit Transaction
		db_cmd = "COMMIT TRANSACTION;"
		result = db_action(db_bm, db_cmd)
		if !result {
			db_transaction_rollback(db_bm)
			job.returnChannel <- false
			continue
		}

		// Send positive Feedback if no Errors
		job.returnChannel <- true
	}

	// Disconnect from DB
	db_close(db_bm)
}

//--------------------------------------------------------------------------------
