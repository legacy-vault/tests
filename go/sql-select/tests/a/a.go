package a

import (
	"database/sql"
	"fmt"
	"log"
	"test/sql-select/text"

	_ "github.com/lib/pq"

	"test/sql-select/errorz"
)

const ErrfTypeUnrecognized = "Type is unrecognized: %v."

const (
	TypeCreateTable      = 1
	TypeFillTable        = 2
	TypeSelectAll        = 4
	TypeCreateFunction   = 8
	TypeSelectDumbNormal = 16
	TypeSelectDumbTest   = 32
	TypeCreateCreator    = 64
	TypeFinalization     = 128
)

func Demo(
	testName string,
	testDescription string,
	testDatabaseDsn string,
	testType int,
) (err error) {

	const TestName = `Demo`

	// Name Pattern.
	const (
		NamePatternAll    = ``
		NamePatternNormal = `Jo`
		NamePatternTest   = `Jo%') AND
	("Id" >= 0) AND
	("Id" <> (SELECT public."add_name_clear"('` + TestName + `'))) AND
	("Name" ILIKE '%Jo`
	)

	var connection *sql.DB
	var namePattern string
	var query string

	fmt.Println("[ " + testName + " ]")
	fmt.Println(testDescription)
	fmt.Println("DSN: " + testDatabaseDsn)

	connection, err = sql.Open("postgres", testDatabaseDsn)
	if err != nil {
		return
	}
	defer func() {
		var derr error
		derr = connection.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	err = connection.Ping()
	if err != nil {
		return
	}

	switch testType {

	case TypeCreateTable:
		query = QueryCreateTable
		err = queryExecute(connection, query)
		if err != nil {
			return
		}
		return

	case TypeFillTable:
		query = QueryFillTable
		err = queryExecute(connection, query)
		if err != nil {
			return
		}
		return

	case TypeSelectAll:
		namePattern = NamePatternAll
		err = querySelectPrintIdName(connection, namePattern)
		if err != nil {
			return
		}
		return

	case TypeCreateFunction:
		query = QueryCreateFunction
		err = queryExecute(connection, query)
		if err != nil {
			return
		}
		return

	case TypeSelectDumbNormal:
		namePattern = NamePatternNormal
		err = querySelectPrintIdName(connection, namePattern)
		if err != nil {
			return
		}
		return

	case TypeSelectDumbTest:
		namePattern = NamePatternTest
		err = querySelectPrintIdName(connection, namePattern)
		if err != nil {
			return
		}
		return

	case TypeCreateCreator:
		query = QueryCreateCreator
		err = queryExecute(connection, query)
		if err != nil {
			return
		}
		return

	case TypeFinalization:
		query = QueryFinalization
		err = queryExecute(connection, query)
		if err != nil {
			return
		}
		return

	default:
		err = fmt.Errorf(
			ErrfTypeUnrecognized,
			testType,
		)
		return
	}
}

func queryExecute(
	connection *sql.DB,
	query string,
) (err error) {

	log.Println("SQL:" + text.NewLine + query)

	_, err = connection.Exec(query)
	if err != nil {
		return
	}

	return nil
}

func querySelectPrintIdName(
	connection *sql.DB,
	namePattern string,
) (err error) {

	var query string
	var rows *sql.Rows

	query = fmt.Sprintf(
		QueryFormatDumb,
		namePattern,
	)
	log.Println("SQL:" + text.NewLine + query)

	rows, err = connection.Query(query)
	if err != nil {
		return
	}
	defer func() {
		var derr error
		derr = rows.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	err = scanPrintIdName(rows)
	if err != nil {
		return
	}

	return nil
}

func scanPrintIdName(
	rows *sql.Rows,
) (err error) {

	var id int
	var name string

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return
		}
		fmt.Printf(
			"Id: %v, Name: %v."+text.NewLine,
			id,
			name,
		)
	}

	return nil
}
