package main

import (
	"test/sql-select/errorz"
	"test/sql-select/tests/a"
	"time"
)

func main() {

	const DatabaseDsn = `postgres://test:test@localhost/test?sslmode=disable`
	const SleepIntervalInterTestSec = 5

	type Test struct {
		Name        string
		Description string
		DatabaseDsn string
		Type        int
	}

	var err error
	var tests []Test

	tests = []Test{
		{
			"Table Creation",
			"...",
			DatabaseDsn,
			a.TypeCreateTable,
		},
		{
			"Table Filling",
			"...",
			DatabaseDsn,
			a.TypeFillTable,
		},
		{
			"Selection of all Records",
			"...",
			DatabaseDsn,
			a.TypeSelectAll,
		},
		{
			"Function Creation",
			"...",
			DatabaseDsn,
			a.TypeCreateFunction,
		},
		{
			"Dumb Selection with normal Pattern",
			"...",
			DatabaseDsn,
			a.TypeSelectDumbNormal,
		},
		{
			"Dumb Selection with test Pattern",
			"...",
			DatabaseDsn,
			a.TypeSelectDumbTest,
		},
		{
			"Procedure Creation",
			"...",
			DatabaseDsn,
			a.TypeCreateCreator,
		},
		{
			"Finalization",
			"...",
			DatabaseDsn,
			a.TypeFinalization,
		},
	}

	testIdxLast := len(tests) - 1
	for testIdx, test := range tests {
		err = a.Demo(
			test.Name,
			test.Description,
			test.DatabaseDsn,
			test.Type,
		)
		errorz.MustBeNoError(err)
		if testIdx != testIdxLast {
			time.Sleep(time.Second * SleepIntervalInterTestSec)
		}
	}
}
