// model_user_test.go

// Version: 0.1.
// Date: 2017-07-07.
// Author: McArcher.

package main

import (
	"fmt"
	"testing"
	"time"
)

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var tUsers UsersModel
var tUser, tUserReply, tmp_user UserModel

var uuid, JohnUUID, KateUUID uint64
var JohnLogin, KateLogin string
var johnDate, KateDate int64

var loginIsWrong, regDateIsWrong, UUIDisWrong, err, ok bool
var e error

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func test_add_user_normal(t *testing.T) {

	fmt.Println("\r\nTest #1. Add Normal User.") //

	// Add [5][John][Now] & check
	JohnLogin = "John"
	tUser.login = JohnLogin
	johnDate = time.Now().Unix()
	tUser.regDate = johnDate
	JohnUUID = 5
	tUser.uuid = JohnUUID
	tUsers.Add(tUser)

	tUser.login = ""
	tUser.regDate = 0
	tUser.uuid = JohnUUID
	ok, e = tUsers.Get(tUser, &tUserReply)
	err = !ok
	if err {
		fmt.Printf(e.Error()) //
		t.FailNow()
	}

	loginIsWrong = tUserReply.login != JohnLogin
	UUIDisWrong = tUserReply.uuid != JohnUUID
	regDateIsWrong = tUserReply.regDate != johnDate
	err = loginIsWrong || UUIDisWrong || regDateIsWrong
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func prepare_data_1() {

	fmt.Println("\r\nData Preparation #1.") //

	time.Sleep(time.Second * 1)

	// Add [6][Kate][Now]
	KateLogin = "Kate"
	tUser.login = KateLogin
	KateDate = time.Now().Unix()
	tUser.regDate = KateDate
	KateUUID = 6
	tUser.uuid = KateUUID
	tUsers.Add(tUser)

	time.Sleep(time.Second * 1)

	// Add [7][Иннокентий][Now]
	tUser.login = "Иннокентий"
	tUser.regDate = time.Now().Unix()
	tUser.uuid = 7
	tUsers.Add(tUser)
}

//------------------------------------------------------------------------------

func test_add_user_existing_uuid(t *testing.T) {

	fmt.Println("\r\nTest #2. Add User with already existing UUID.") //

	// Add fake User with existing UUID & check
	tUser.login = "Existing UUID"
	tUser.regDate = time.Now().Unix()
	tUser.uuid = JohnUUID
	err, e = tUsers.Add(tUser) // must be FALSE
	if err {
		fmt.Printf(e.Error()) //
		t.FailNow()
	}

	tUser.login = ""
	tUser.regDate = 0
	tUser.uuid = JohnUUID
	ok, e = tUsers.Get(tUser, &tUserReply)
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
	err = (tUserReply.login != JohnLogin)
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_add_user_existing_login(t *testing.T) {

	fmt.Println("\r\nTest #3. Add User with already existing login.") //

	// Add fake User with existing login & check
	tUser.login = "John"
	tUser.regDate = time.Now().Unix()
	tUser.uuid = 101
	tUsers.Add(tUser)

	ok, e = tUsers.GetUUIDbyLogin("John", &uuid)
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
	err = (uuid == 101)
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_add_user_empty_uuid(t *testing.T) {

	fmt.Println("\r\nTest #4. Add User with empty UUID.") //

	// Add fake User with empty UUID & check
	tUser.login = "Bad UUID"
	tUser.regDate = time.Now().Unix()
	tUser.uuid = 0
	tUsers.Add(tUser)

	err, e = tUsers.GetUUIDbyLogin("Bad UUID", &uuid) // must be FALSE
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_add_user_empty_login(t *testing.T) {

	fmt.Println("\r\nTest #5. Add User with empty login.") //

	// Add fake User with empty login & check
	tUser.login = ""
	tUser.regDate = time.Now().Unix()
	tUser.uuid = 102
	err, e = tUsers.Add(tUser) // must be FALSE
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
	err, e = tUsers.Get(tUser, &tUserReply) // must be FALSE
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_add_user_empty_regdate(t *testing.T) {

	fmt.Println("\r\nTest #6. Add User with empty regDate.") //

	// Add User with empty regDate & check
	tUser.login = "Empty-1"
	tUser.regDate = 0
	tUser.uuid = 8
	tUsers.Add(tUser)

	tUsers.Get(tUser, &tUserReply)
	tUser.debug_Print()      //
	tUserReply.debug_Print() //

	loginIsWrong = tUserReply.login != "Empty-1"
	UUIDisWrong = tUserReply.uuid != 8
	regDateIsWrong = tUserReply.regDate == 0
	err = loginIsWrong || UUIDisWrong || regDateIsWrong
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_get_user_by_uuid(t *testing.T) {

	fmt.Println("\r\nTest #7. Get User by UUID.") //

	// Get User by UUID
	tUser.login = ""
	tUser.regDate = 0
	tUser.uuid = JohnUUID
	tUsers.Get(tUser, &tUserReply)
	tUserReply.debug_Print() //

	// Check Get
	UUIDisWrong = tUserReply.uuid != JohnUUID
	loginIsWrong = tUserReply.login != JohnLogin
	regDateIsWrong = tUserReply.regDate != johnDate
	err = loginIsWrong || regDateIsWrong || UUIDisWrong
	if err {
		fmt.Printf("Error with %d.\r\n", JohnLogin) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_get_user_by_login(t *testing.T) {

	fmt.Println("\r\nTest #8. Get User by login.") //

	// Get User by login
	tUser.login = KateLogin
	tUser.regDate = 0
	tUser.uuid = 0
	tUsers.Get(tUser, &tUserReply)
	tUserReply.debug_Print() //

	// Check Get
	loginIsWrong = tUserReply.login != KateLogin
	regDateIsWrong = tUserReply.regDate != KateDate
	UUIDisWrong = tUserReply.uuid != KateUUID
	err = loginIsWrong || regDateIsWrong || UUIDisWrong
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_get_uuid_by_login(t *testing.T) {

	fmt.Println("\r\nTest #9. Get UUID by login.") //

	// Get UUID by login & check
	ok, e = tUsers.GetUUIDbyLogin(KateLogin, &uuid)
	UUIDisWrong = (!ok) || (uuid != KateUUID)
	fmt.Printf("Kate's UUID is [%d]. Must be [%d].\r\n", uuid, KateUUID) //
	if UUIDisWrong {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_get_user_by_nonexistent_uuid(t *testing.T) {

	fmt.Println("\r\nTest #10. Get User by non-existent UUID.") //

	// Get User by non-existent ID
	tUser.login = ""
	tUser.regDate = 0
	tUser.uuid = 103
	err, e = tUsers.Get(tUser, &tUserReply)

	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_get_user_by_nonexistent_login(t *testing.T) {

	fmt.Println("\r\nTest #11. Get User by non-existent login.") //

	// Get User by non-existent login
	tUser.login = "Non-Existent Login"
	tUser.regDate = 0
	tUser.uuid = 0
	err, e = tUsers.Get(tUser, &tUserReply)

	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_get_user_empty_uuid_and_login(t *testing.T) {

	fmt.Println("\r\nTest #12. Get User by empty UUID & login.") //

	// Get User when both UUID & login are empty
	tUser.login = ""
	tUser.regDate = 0
	tUser.uuid = 0
	err, e = tUsers.Get(tUser, &tUserReply)

	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_loginisfree(t *testing.T) {

	fmt.Println("\r\nTest #13. Test 'LoginIsFree' Method.") //

	// Test 'LoginIsFree' Method
	err = tUsers.LoginIsFree(KateLogin)
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}

	ok = tUsers.LoginIsFree("Non-Existent Login")
	err = !ok
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_uuidexists(t *testing.T) {

	fmt.Println("\r\nTest #14. Test 'UUIDexists' Method.") //

	// Test 'UUIDexists' Method
	ok = tUsers.UUIDexists(KateUUID)
	err = !ok
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}

	err = tUsers.UUIDexists(999)
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func show_list() {

	// Show List.
	tUsers.debug_Print() //
}

//------------------------------------------------------------------------------

func test_modify_user_empty_uuid(t *testing.T) {

	fmt.Println("\r\nTest #15. Modify a User, empty UUID.") //

	// Modify, empty UUID
	tUser.login = "Kate"
	tUser.regDate = 0
	tUser.uuid = 0

	err, e = tUsers.Modify(tUser)
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	ok, e = tUsers.Get(tUser, &tUserReply)
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
	loginIsWrong = tUserReply.login != KateLogin
	regDateIsWrong = tUserReply.regDate != KateDate
	UUIDisWrong = tUserReply.uuid != KateUUID
	err = loginIsWrong || regDateIsWrong || UUIDisWrong
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_modify_user_fake_uuid(t *testing.T) {

	fmt.Println("\r\nTest #16. Modify a User, fake UUID.") //

	// Modify, fake UUID
	tUser.login = "Kate"
	tUser.regDate = 0
	tUser.uuid = 999

	err, e = tUsers.Modify(tUser)
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	err, e = tUsers.Get(tUser, &tUserReply)
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_modify_user_empty_login(t *testing.T) {

	fmt.Println("\r\nTest #17. Modify a User, empty login.") //

	// Modify, empty login
	tUser.login = ""
	tUser.regDate = 0
	tUser.uuid = KateUUID // 6

	err, e = tUsers.Modify(tUser)
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	ok, e = tUsers.Get(tUser, &tUserReply)
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
	loginIsWrong = tUserReply.login != KateLogin
	regDateIsWrong = tUserReply.regDate != KateDate
	UUIDisWrong = tUserReply.uuid != KateUUID
	err = loginIsWrong || regDateIsWrong || UUIDisWrong
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_modify_user_foreign_existing_login(t *testing.T) {

	fmt.Println("\r\nTest #18. Modify a User, foreign existing login.") //

	// Modify, login is already taken by other UUID.
	// Setting John's login for Kate must return FALSE.
	tUser.login = JohnLogin
	tUser.regDate = 0
	tUser.uuid = KateUUID

	err, e = tUsers.Modify(tUser)
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	ok, e = tUsers.Get(tUser, &tUserReply)
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
	loginIsWrong = tUserReply.login != KateLogin
	regDateIsWrong = tUserReply.regDate != KateDate
	UUIDisWrong = tUserReply.uuid != KateUUID
	err = loginIsWrong || regDateIsWrong || UUIDisWrong
	if err {
		fmt.Println("Error.") //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_integrity_check(t *testing.T) {

	fmt.Println("\r\nTest #19. Integrity Check.") //

	// Integrity Check.
	ok, e = tUsers.CheckIntegrity()
	if !ok {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_integrity_check_broken_index(t *testing.T) {

	fmt.Println("\r\nTest #20. Integrity Check, broken Index.") //

	// 'CheckIntegrity' Check: bad UUID (Mismatch with Map Index)
	uuid = 100 // hacked UUID
	tmp_user = tUsers[KateUUID]
	tmp_user.uuid = uuid
	tUsers[KateUUID] = tmp_user

	tUser.login = KateLogin
	tUser.regDate = 0
	tUser.uuid = uuid
	ok, e = tUsers.Get(tUser, &tUserReply)
	fmt.Printf("tUsers[%d] is:\r\n", uuid) //
	tUserReply.debug_Print()               //

	// Start Integrity Check & check its Result
	err, e = tUsers.CheckIntegrity() // must be FASLE as we broke the List
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	// Revert changes to normal
	tmp_user = tUsers[KateUUID]
	tmp_user.uuid = KateUUID
	tUsers[KateUUID] = tmp_user
	ok, e = tUsers.CheckIntegrity()
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_integrity_check_empty_uuid(t *testing.T) {

	fmt.Println("\r\nTest #21. Integrity Check, empty UUID.") //

	// 'CheckIntegrity' Check: empty UUID.
	tmp_user = tUsers[KateUUID]
	tmp_user.uuid = 0
	tUsers[KateUUID] = tmp_user

	// Start Integrity Check & check its Result
	err, e = tUsers.CheckIntegrity() // must be FASLE as we broke the List
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	// Revert changes to normal
	tmp_user = tUsers[KateUUID]
	tmp_user.uuid = KateUUID
	tUsers[KateUUID] = tmp_user
	ok, e = tUsers.CheckIntegrity()
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_integrity_check_empty_login(t *testing.T) {

	fmt.Println("\r\nTest #22. Integrity Check, empty login.") //

	// 'CheckIntegrity' Check: empty login.
	tmp_user = tUsers[KateUUID]
	tmp_user.login = ""
	tUsers[KateUUID] = tmp_user

	// Start Integrity Check & check its Result
	err, e = tUsers.CheckIntegrity() // must be FASLE as we broke the List
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	// Revert changes to normal
	tmp_user = tUsers[KateUUID]
	tmp_user.login = KateLogin
	tUsers[KateUUID] = tmp_user
	ok, e = tUsers.CheckIntegrity()
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_integrity_check_empty_regdate(t *testing.T) {

	fmt.Println("\r\nTest #23. Integrity Check, empty regDate.") //

	// 'CheckIntegrity' Check: empty regDate.
	tmp_user = tUsers[KateUUID]
	tmp_user.regDate = 0
	tUsers[KateUUID] = tmp_user

	// Start Integrity Check & check its Result
	err, e = tUsers.CheckIntegrity() // must be FASLE as we broke the List
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	// Revert changes to normal
	tmp_user = tUsers[KateUUID]
	tmp_user.regDate = KateDate
	tUsers[KateUUID] = tmp_user
	ok, e = tUsers.CheckIntegrity()
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_integrity_check_duplicate_uuid(t *testing.T) {

	fmt.Println("\r\nTest #24. Integrity Check, duplicate UUID.") //

	// 'CheckIntegrity' Check: duplicate UUID.
	// Set Kate's UUID the same as John's UUID & check.
	tmp_user = tUsers[KateUUID]
	tmp_user.uuid = JohnUUID
	tUsers[KateUUID] = tmp_user

	// Start Integrity Check & check its Result
	err, e = tUsers.CheckIntegrity() // must be FASLE as we broke the List
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}

	// Revert changes to normal
	tmp_user = tUsers[KateUUID]
	tmp_user.uuid = KateUUID
	tUsers[KateUUID] = tmp_user
	ok, e = tUsers.CheckIntegrity()
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_integrity_check_duplicate_login(t *testing.T) {

	fmt.Println("\r\nTest #25. Integrity Check, duplicate login.") //

	// 'CheckIntegrity' Check: duplicate login.
	// Set Kate's login the same as John's login & check.
	tmp_user = tUsers[KateUUID]
	tmp_user.login = JohnLogin
	tUsers[KateUUID] = tmp_user

	// Start Integrity Check & check its Result
	ok, e = tUsers.CheckIntegrity() // must be FASLE as we broke the List
	err = ok
	if err {
		t.FailNow()
	}

	// Revert changes to normal
	tmp_user = tUsers[KateUUID]
	tmp_user.login = KateLogin
	tUsers[KateUUID] = tmp_user
	ok, e = tUsers.CheckIntegrity()
	err = !ok
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------

func test_add_user_long_login(t *testing.T) {

	fmt.Println("\r\nTest #26. Add User with too long login.") //

	// Add User with too long login & check
	tUser.login = "LOGIN_" + // +6 S=6
		"1234567890" + "1234567890" + "1234567890" + "1234567890" + "1234567890" + // +50 S=56
		"1234567890" + "1234567890" + "1234567890" + "1234567890" + "1234567890" + // +50 S=106
		"1234567890" + "1234567890" + "1234567890" + "1234567890" + "1234567890" + // +50 S=156
		"1234567890" + "1234567890" + "1234567890" + "1234567890" + "1234567890" + // +50 S=206
		"1234567890" + "1234567890" + "1234567890" + "1234567890" + "_______XXX" // +50 S=256
	tUser.regDate = time.Now().Unix()
	tUser.uuid = 102
	err, e = tUsers.Add(tUser) // must be FALSE
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
	err, e = tUsers.Get(tUser, &tUserReply) // must be FALSE
	if err {
		fmt.Println(e.Error()) //
		t.FailNow()
	}
}

//------------------------------------------------------------------------------
