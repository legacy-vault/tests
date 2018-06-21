// a.go.

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

import (
	"fmt"
	"log"
	"database/sql"
)

type Record struct {
	TheID   uint64 `db:"id"`
	TheName string `db:"name"`
	TheAge  uint8  `db:"age"`
}

func main() {

	var cmd string
	var db *sql.DB
	var dbx *sqlx.DB
	var db_base string
	var db_DSN string
	var db_password string
	var db_table string
	var db_user string
	var err error
	var record Record
	var records []Record
	var result sql.Result

	// Connection Parameters.
	db_base = "test"
	db_user = "test"
	db_password = "test"
	db_DSN = db_user + ":" + db_password + "@/" + db_base

	// Connect.
	db, err = sql.Open("mysql", db_DSN)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Table.
	db_table = "test_table"

	cmd = "CREATE TABLE IF NOT EXISTS `" + db_base + "`.`" + db_table + "` " +
		"( " +
		"`id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT, " +
		"`name` VARCHAR(256) NOT NULL, " +
		"`age` TINYINT UNSIGNED NOT NULL" +
		");"

	fmt.Println(cmd)
	result, err = db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %+v.\r\n", result)

	// Write to the Table.
	record = Record{TheID: 125, TheAge: 21, TheName: "Телепузик"}
	cmd = "INSERT INTO `" + db_base + "`.`" + db_table + "` " +
		"(name, age) VALUES (?, ?);"

	fmt.Println(cmd)
	result, err = db.Exec(cmd, record.TheName, record.TheAge)
	//result, err = db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %+v.\r\n", result)

	// DisConnect.
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Connect.
	dbx, err = sqlx.Connect("mysql", db_DSN)
	if err != nil {
		log.Fatal(err)
	}

	/*
	// Write to the Table.
	record = Record{theID: 125, theAge: 20, theName: "Телепузик"}
	cmd = "INSERT INTO `" + db_base + "`.`" + db_table + "` " +
		"(id, name, age) VALUES (:id, :name, :age);"

	fmt.Println(cmd)
	result, err = dbx.NamedExec(cmd, &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %+v.\r\n", result)
	*/

	// Read from the Table.
	records = []Record{}
	cmd = "SELECT * FROM `" + db_base + "`.`" + db_table + "` " +
		"ORDER BY `age` ASC;"

	fmt.Println(cmd)
	err = dbx.Select(&records, cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Records: %+v.\r\n", records)

	// DisConnect.
	err = dbx.Close()
	if err != nil {
		log.Fatal(err)
	}
}
