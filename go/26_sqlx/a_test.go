// a_test.go.

package main

import (
	"strings"
	"testing"
)

func BenchmarkA1(b *testing.B) {

	b.StartTimer()

	var db_base string
	var db_DSN strings.Builder
	var db_password string
	var db_user string
	var i int

	for i = 0; i < b.N; i++ {

		db_base = "test_base"
		db_user = "test_user"
		db_password = "test_password"

		db_DSN.WriteString(db_user)
		db_DSN.WriteString(":")
		db_DSN.WriteString(db_password)
		db_DSN.WriteString("@/")
		db_DSN.WriteString(db_base)
	}

	b.StopTimer()
}

func BenchmarkA2(b *testing.B) {

	b.StartTimer()

	var db_base string
	var db_DSN string
	var db_password string
	var db_user string
	var i int

	for i = 0; i < b.N; i++ {

		db_base = "test_base"
		db_user = "test_user"
		db_password = "test_password"

		db_DSN = db_user + ":" + db_password + "@/" + db_base
	}

	b.StopTimer()

	db_DSN = db_DSN //!
}
