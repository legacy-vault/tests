// db.go
//
// concurrent edition

package main

import "fmt"
import "strings"
import "log"
import "database/sql"
import _ "github.com/lib/pq"

//--------------------------------------------------------------------------------

func db_connect_pg(host string, port int, user, pwd, dbase, sslm string) *sql.DB {

	// Connects to a PostgreSQL DataBase.

	var psqlInfo string
	var db *sql.DB
	var err error

	// Connect to PostgreSQL DataBase
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pwd, dbase, sslm)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err) //
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Println(err) //
		return nil
	}
	return db
}

//--------------------------------------------------------------------------------

func db_close(db *sql.DB) {

	// Closes Connection to the DataBase.

	var err error

	// Close DataBase
	err = db.Close()
	if err != nil {
		log.Println(err) //
	}
}

//--------------------------------------------------------------------------------

func db_tables_init(db *sql.DB, tbl_users, tbl_jp, col_userName, col_persSum, col_JPSum string) {

	// Creates Tables and Indices if needed.

	var table_exists bool
	var fields string

	// tbl_users
	table_exists = true
	table_exists = db_table_exists(db, tbl_users)
	if !table_exists {

		fields = col_userName + " text PRIMARY KEY, " + col_persSum +
			" numeric NOT NULL DEFAULT 0"
		db_table_create(db, tbl_users, fields)

		fields = col_userName + ", " + col_persSum
		db_index_create(db, tbl_users, fields)

	} else {

		log.Println("Table already exists!") //
	}

	// tbl_jp
	table_exists = true
	table_exists = db_table_exists(db, tbl_jp)
	if !table_exists {

		fields = col_JPSum + " numeric PRIMARY KEY"
		db_table_create(db, tbl_jp, fields)
		db_table_value_insert(db, tbl_jp, column_JPSum, "0")

		fields = column_JPSum
		db_index_create(db, table_jp, fields)

	} else {

		log.Println("Table already exists!") //
	}
}

//--------------------------------------------------------------------------------

func db_tables_delete(db *sql.DB, tbl_users, tbl_jp string) {

	// Checks if Tables exist and if yes, then deletes them.

	var table_exists bool

	// tbl_users
	table_exists = true
	table_exists = db_table_exists(db, tbl_users)
	if table_exists {

		db_table_delete(db, tbl_users)

	} else {

		log.Println("Table does not exist!") //
	}

	// tbl_jp
	table_exists = true
	table_exists = db_table_exists(db, tbl_jp)
	if table_exists {

		db_table_delete(db, tbl_jp)

	} else {

		log.Println("Table does not exist!") //
	}
}

//--------------------------------------------------------------------------------

func db_table_create(db *sql.DB, table_name, fields string) {

	// Creates a Table.

	var db_cmd, user string
	var err error

	user = db_user_2

	db_cmd = fmt.Sprintf("CREATE TABLE %s ( %s );", table_name, fields)

	_, err = db.Exec(db_cmd)
	if err != nil {
		log.Println(err) //
	}

	db_cmd = fmt.Sprintf("GRANT SELECT, UPDATE, INSERT ON TABLE %s TO %s;",
		table_name, user)

	_, err = db.Exec(db_cmd)
	if err != nil {
		log.Println(err) //
	}
}

//--------------------------------------------------------------------------------

func db_index_create(db *sql.DB, table_name, fields string) {

	// Creates an Index.

	var db_cmd string
	var err error

	db_cmd = fmt.Sprintf("CREATE INDEX ON %s ( %s );", table_name, fields)

	_, err = db.Exec(db_cmd)
	if err != nil {
		log.Println(err) //
	}
}

//--------------------------------------------------------------------------------

func db_table_delete(db *sql.DB, table_name string) {

	// Deletes a Table.

	var db_cmd string
	var err error

	db_cmd = fmt.Sprintf("DROP TABLE %s;", table_name)

	_, err = db.Exec(db_cmd)
	if err != nil {
		log.Println(err) //
	}
}

//--------------------------------------------------------------------------------

func db_table_exists(db *sql.DB, table_name string) bool {

	// Checks if Table exists.

	var table_exists bool
	var tables []string
	var search_name string
	var name, name_low string

	table_exists = false
	search_name = strings.ToLower(table_name)

	// List existing Tables in the DB
	tables = db_tables_list(db)

	// Check if the Table exists
	for _, name = range tables {

		name_low = strings.ToLower(name)
		if name_low == search_name {

			table_exists = true
			break
		}
	}

	return table_exists
}

//--------------------------------------------------------------------------------

func db_tables_list(db *sql.DB) (tables []string) {

	// Lists Tables in the DataBase.

	var name string
	var db_qr string
	var qr_rows *sql.Rows
	var err error
	var qr_rows_n int

	db_qr = "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';"

	qr_rows, err = db.Query(db_qr)
	if err != nil {
		log.Println(err) //
	}
	defer qr_rows.Close()

	qr_rows_n = 0
	for qr_rows.Next() {

		err = qr_rows.Scan(&name)
		if err != nil {
			log.Println(err) //
		}
		tables = append(tables, name)

		qr_rows_n++
	}
	err = qr_rows.Err()
	if err != nil {
		log.Println(err) //
	}

	return tables
}

//--------------------------------------------------------------------------------

func db_tables_exist(db *sql.DB, table_1, table_2 string) bool {

	// Checks if both Tables exist.

	var table_1_exists, table_2_exists bool

	table_1_exists = false
	table_1_exists = db_table_exists(db, table_1)

	table_2_exists = false
	table_2_exists = db_table_exists(db, table_2)

	return (table_1_exists && table_2_exists)
}

//--------------------------------------------------------------------------------

func db_table_value_insert(db *sql.DB, table_name, columns, values string) {

	// Inserts Data into a Table.

	var db_cmd string
	var err error

	db_cmd = fmt.Sprintf("INSERT INTO %s ( %s ) VALUES ( '%s' );",
		table_name, columns, values)

	_, err = db.Exec(db_cmd)
	if err != nil {
		log.Println(err) //
	}
}

//--------------------------------------------------------------------------------

func db_transaction_rollback(db *sql.DB) {

	// Rolls back the Transaction.

	var db_cmd string
	var err error

	db_cmd = "ROLLBACK TRANSACTION;"

	_, err = db.Exec(db_cmd)
	if err != nil {
		log.Println(err) //
	}
}

//--------------------------------------------------------------------------------

func db_action(db *sql.DB, action string) bool {

	// Executes a Command in DataBase.

	var err error

	_, err = db.Exec(action)
	if err != nil {
		log.Println(err) //
		return false
	}

	return true
}

//--------------------------------------------------------------------------------

func db_query_value(db *sql.DB, query string, reply *string) bool {

	// Makes a Query to DataBase and returns true.
	// If Errors occur, returns false.
	// The Query must return only one Value (one Row and one Column).

	var qr_row *sql.Row
	var err error

	qr_row = db.QueryRow(query)
	err = qr_row.Scan(reply)
	if err != nil {
		log.Println(err) //
		return false
	}

	return true
}

//--------------------------------------------------------------------------------
