// main_test.go

// Version: 0.3.
// Date: 2017-07-08.
// Author: McArcher.

package main

import (
	"testing"
)

//------------------------------------------------------------------------------

func TestMain(t *testing.T) {

	tUsers.Init()

	UsersIntegrityCheck.StageOneEnabled = true   // Must be TRUE !
	UsersIntegrityCheck.StageTwoEnabled = true   // Must be TRUE !
	UsersIntegrityCheck.StageThreeEnabled = true // Must be TRUE !
	UsersIntegrityCheck.ShowStages = false

	test_add_user_normal(t)
	prepare_data_1() // aux
	test_add_user_existing_uuid(t)
	test_add_user_existing_login(t)
	test_add_user_empty_uuid(t)
	test_add_user_empty_login(t)
	test_add_user_empty_regdate(t)
	test_get_user_by_uuid(t)
	test_get_user_by_login(t)
	test_get_uuid_by_login(t)
	test_get_user_by_nonexistent_uuid(t)
	test_get_user_by_nonexistent_login(t)
	test_get_user_empty_uuid_and_login(t)
	test_loginisfree(t)
	test_uuidexists(t)
	show_list() // aux
	test_modify_user_empty_uuid(t)
	test_modify_user_fake_uuid(t)
	test_modify_user_empty_login(t)
	test_modify_user_foreign_existing_login(t)
	test_integrity_check(t)
	test_integrity_check_broken_index(t)
	test_integrity_check_empty_uuid(t)
	test_integrity_check_empty_login(t)
	test_integrity_check_empty_regdate(t)
	test_integrity_check_duplicate_uuid(t)
	test_integrity_check_duplicate_login(t)
	test_add_user_long_login(t)

	//test_server_1(t)
	//test_server_2(t)
}

//------------------------------------------------------------------------------
