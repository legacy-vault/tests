// model_user.go

// Version: 0.1.
// Date: 2017-07-06.
// Author: McArcher.

package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

//------------------------------------------------------------------------------
// Types
//------------------------------------------------------------------------------

type UserModel struct {
	uuid    uint64
	login   string
	regDate int64 // UNIX Timestamp
}

type UsersModel map[uint64]UserModel

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var Users UsersModel

//------------------------------------------------------------------------------
// Methods
//------------------------------------------------------------------------------

func (users *UsersModel) Add(user UserModel) (bool, error) {

	// Adds a User to the List.
	//
	// Returns FALSE if:
	//		- 'UUID' is empty;
	//		- 'login' is empty;
	//		- 'UUID' already exists;
	//		- 'login' already exists.

	var uuid uint64
	var uuidIsTaken bool
	var uuidIsEmpty bool

	var login string
	var loginIsTaken bool
	var loginIsEmpty bool
	var loginsLengthIsGood bool

	var regDate int64
	var regDateIsEmpty bool

	var tmp_user *UserModel

	var error_msg string
	var warning_msg string
	var ret_err error

	error_msg = ""

	// Check for empty 'UUID'.
	uuid = user.uuid
	uuidIsEmpty = (uuid == 0)
	if uuidIsEmpty {

		error_msg = "Add Error. Empty 'UUID' is given."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Check for empty 'login'.
	login = user.login
	loginIsEmpty = (login == "")
	if loginIsEmpty {

		error_msg = "Add Error. Empty 'login' is given."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Check 'login' Length
	loginsLengthIsGood = users.LoginLengthIsGood(&login)
	if !loginsLengthIsGood {

		error_msg = fmt.Sprintf("Add Error. 'login' is too long [%s].", login)
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Check if UUID is already taken.
	uuidIsTaken = users.UUIDexists(uuid)
	if uuidIsTaken {

		error_msg = fmt.Sprintf("Add Error. 'UUID' already exists [%d].", uuid)
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Check if login is already taken.
	loginIsTaken = !users.LoginIsFree(login)
	if loginIsTaken {

		error_msg = fmt.Sprintf("Add Error. 'login' already exists [%s].", login)
		ret_err = errors.New(error_msg)
		log.Printf(error_msg) //

		return false, ret_err
	}

	// Check for empty regDate.
	regDate = user.regDate
	regDateIsEmpty = (regDate == 0)
	if regDateIsEmpty {

		warning_msg = "Add Warning. 'regDate' Field is not set in the Parameter."
		log.Println(warning_msg) //
		regDate = time.Now().Unix()

		// 'user' Parameter is not correct.
		// Create a correct temporary Parameter.
		tmp_user = new(UserModel)
		tmp_user.login = login
		tmp_user.regDate = regDate
		tmp_user.uuid = uuid

		// Add User from temporary Object.
		(*users)[uuid] = *tmp_user

	} else {

		// 'user' Parameter is good.
		// Use it.

		// Add User from Parameter.
		(*users)[uuid] = user
	}

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) AddAsRequest(user UserModel) (bool, error) {

	// Adds a User, which was given as a Request.

	var ok bool
	var err error
	var ret_err error

	ok, err = users.Add(user)

	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Save changes in Delta Map
	AddedUsersMap[user.uuid] = true

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) CheckIntegrity() (bool, error) {

	// Checks the Integrity of all Elements in the List.
	// This Function is used for testing Purposes.

	var cur_user, cur_user_2 UserModel
	var cur_index uint64
	var cur_uuid, cur_uuid_2 uint64
	var cur_login, cur_login_2 string
	var cur_regDate int64

	var indexIsBad bool
	var uuidIsBad bool
	var loginIsBad bool
	var regDateIsBad bool
	var loginsLengthIsGood bool
	var errorIsFound bool
	var count uint64

	var error_msg string
	var ret_err error
	var err_log string

	errorIsFound = false
	err_log = ""

	// 1. Basic Checks.

	for cur_index, cur_user = range *users {

		// Read current User.
		cur_uuid = cur_user.uuid
		cur_login = cur_user.login
		cur_regDate = cur_user.regDate

		// Check Integrity of Index.
		indexIsBad = (cur_index != cur_uuid)
		if indexIsBad {

			err_log = err_log + fmt.Sprintf("Bad Index [%d].", cur_index)
			errorIsFound = true
		}

		// Check for empty UUID.
		uuidIsBad = (cur_uuid == 0)
		if uuidIsBad {

			err_log = err_log + fmt.Sprintf("Bad UUID [%d].", cur_uuid)
			errorIsFound = true
		}

		// Check for empty login.
		loginIsBad = (cur_login == "")
		if loginIsBad {

			err_log = err_log + fmt.Sprintf("Bad login [%s].", cur_login)
			errorIsFound = true
		}

		// Check for empty regDate.
		regDateIsBad = (cur_regDate == 0)
		if regDateIsBad {

			err_log = err_log + fmt.Sprintf("Bad regDate [%d].", cur_regDate)
			errorIsFound = true
		}

		// Check 'login' Length
		loginsLengthIsGood = users.LoginLengthIsGood(&cur_login)
		if !loginsLengthIsGood {

			err_log = err_log + fmt.Sprintf("Too long login [%s].", cur_login)
			errorIsFound = true
		}
	}

	// 2. Check duplicate UUIDs.
	for _, cur_user = range *users {

		cur_uuid = cur_user.uuid
		count = 0

		for _, cur_user_2 = range *users {

			cur_uuid_2 = cur_user_2.uuid
			if cur_uuid_2 == cur_uuid {
				count++
			}
		}

		if count != 1 {

			err_log = err_log + fmt.Sprintf("UUID [%d] is found [%d] Times in List.", cur_uuid, count)
			errorIsFound = true
		}
	}

	// 3. Check duplicate logins.
	for _, cur_user = range *users {

		cur_login = cur_user.login
		count = 0

		for _, cur_user_2 = range *users {

			cur_login_2 = cur_user_2.login
			if cur_login_2 == cur_login {
				count++
			}
		}

		if count != 1 {

			err_log = err_log + fmt.Sprintf("login [%s] is found [%d] Times in List.", cur_login, count)
			errorIsFound = true
		}
	}

	// Summary.
	if errorIsFound {

		error_msg = "Integrity is BAD." + err_log
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	return true, nil

}

//------------------------------------------------------------------------------

func (users *UsersModel) Get(user UserModel, reply *UserModel) (bool, error) {

	// Gets the User from the List.
	//
	// Returns FALSE if:
	//		- 'UUID' does not exist.
	//		- 'login' does not exist.
	//		- both 'UUID' & 'login' are empty.

	var uuid uint64
	var uuidIsSet bool
	var uuidIsFree bool
	var uuidByLogin uint64

	var login string
	var loginIsSet bool
	var loginIsFree bool

	var ok bool
	var error_msg string
	var err error
	var ret_err error

	// 1. Get User by UUID if it is specified.

	// Check for empty UUID.
	uuid = user.uuid
	uuidIsSet = (uuid != 0)
	if uuidIsSet {

		// Check if UUID is already taken.
		uuidIsFree = !(users.UUIDexists(uuid))
		if uuidIsFree {

			error_msg = fmt.Sprintf("Get Error. 'UUID' does not exist [%d].", uuid)
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			return false, ret_err
		}

		// UUID is set & exists.
		*reply = (*users)[uuid]

		return true, nil
	}

	// 2. Get User by login if it is specified.

	// Check for empty login.
	login = user.login
	loginIsSet = (login != "")
	if loginIsSet {

		// Check if login is already taken.
		loginIsFree = users.LoginIsFree(login)
		if loginIsFree {

			error_msg = fmt.Sprintf("Get Error. 'login' does not exist [%s].", login)
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			return false, ret_err
		}

		// login is set & exists.
		ok, err = users.GetUUIDbyLogin(login, &uuidByLogin)
		if !ok {

			error_msg = fmt.Sprintf("Get Error. Can not get 'UUID' by 'login' [%s]. %s", login, err.Error())
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			return false, ret_err
		}
		*reply = (*users)[uuidByLogin]

		return true, nil
	}

	// 3. Both 'UUID' & 'login' are not specified.

	error_msg = "Get Error. Neither 'UUID' nor 'login' is specified."
	ret_err = errors.New(error_msg)
	log.Println(error_msg) //

	return false, ret_err
}

//------------------------------------------------------------------------------

func (users *UsersModel) GetUUIDbyLogin(login string, uuid *uint64) (bool, error) {

	// Gets an 'UUID' of a User with a specified 'login'.
	//
	// Returns FALSE if:
	//		- 'login' is empty;
	//		- 'login' is not found.

	var loginIsEmpty bool
	var cur_user UserModel
	var cur_login string
	var cur_uuid uint64
	var loginIsFound bool

	var error_msg string
	var ret_err error

	// Fool Check.
	loginIsEmpty = (login == "")
	if loginIsEmpty {

		error_msg = "GetUUIDbyLogin Error. Empty 'login' is given."
		ret_err = errors.New(error_msg)
		//log.Println(error_msg) // Do NOT log this Spam as it is very often used

		return false, ret_err
	}

	// Search for 'login'.
	for cur_uuid, cur_user = range *users {

		cur_login = cur_user.login
		loginIsFound = (cur_login == login)

		// 'login' is found.
		if loginIsFound {

			*uuid = cur_uuid

			return true, nil
		}
	}

	// 'login' is not found.

	error_msg = "GetUUIDbyLogin Error. 'login' is not found."
	ret_err = errors.New(error_msg)
	//log.Println(error_msg) // Do NOT log this Spam as it is very often used

	return false, ret_err
}

//------------------------------------------------------------------------------

func (users *UsersModel) Init() {

	// Initializes the List.

	//*users = make(map[uint64]UserModel)
	*users = make(UsersModel)
}

//------------------------------------------------------------------------------

func (users *UsersModel) Load() (bool, error) {

	// Loads all Users from DataBase & checks Integrity.

	var ok bool
	var err error
	var error_msg string
	var ret_err error

	// Find File
	ok, err = db_find_file(db_path)
	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Load DB
	ok, err = db_load_list(users)
	if !ok {

		error_msg = "Error. Can not load List. " + err.Error()
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		// Reset RecordsMap
		db_reset_records_map()

		return false, ret_err
	}

	// Check Integrity
	ok, err = users.CheckIntegrity()
	if !ok {

		error_msg = "Error. The DataBase is broken." + err.Error()
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		// Reset RecordsMap
		db_reset_records_map()

		return false, ret_err
	}

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) LoginIsFree(login string) bool {

	// Checks if the specified 'login' is not used in the List.

	var cur_user UserModel
	var cur_login string
	var loginExists bool

	for _, cur_user = range *users {

		cur_login = cur_user.login
		loginExists = (cur_login == login)

		if loginExists {

			return false
		}
	}

	return true
}

//------------------------------------------------------------------------------

func (users *UsersModel) LoginLengthIsGood(login *string) bool {

	// Checks if the Length of 'login' is good enough to be stored into the DB.

	var len_is_ok bool

	len_is_ok = len(*login) <= db_field_login_maxlen

	if len_is_ok {
		return true
	}

	return false
}

//------------------------------------------------------------------------------

func (users *UsersModel) Modify(user UserModel) (bool, error) {

	// Modifies a specified User.
	// User is selected by UUID.
	//
	// Returns FALSE if:
	//		- 'UUID' is empty;
	//		- 'UUID' is fake;
	//		- 'login' is empty;
	//		- 'login' is already taken by another User.

	var uuid uint64
	var uuidIsEmpty bool
	var uuidIsFree bool
	var uuidByLogin uint64

	var login string
	var loginIsEmpty bool
	var loginIsBusy bool
	var loginOwnerIsNotUs bool

	var regDate int64
	var regDateIsEmpty bool

	var tmp_user *UserModel
	//var err error

	var warning_msg string
	var error_msg string
	var ret_err error

	uuid = user.uuid
	login = user.login
	regDate = user.regDate

	// 1. Check 'UUID'.

	// 1.1. Empty?
	uuidIsEmpty = (uuid == 0)
	if uuidIsEmpty {

		error_msg = "Modify Error. Empty 'UUID' is given."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// 1.2. Fake?
	uuidIsFree = !(users.UUIDexists(uuid))
	if uuidIsFree {

		error_msg = fmt.Sprintf("Modify Error. 'UUID' does not exist [%d].", uuid)
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// 2. Check 'login'.

	// 2.1. Empty?
	loginIsEmpty = (login == "")
	if loginIsEmpty {

		error_msg = "Modify Error. Empty 'login' is given."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// 2.2. Collision Check.
	// 'login' is already taken by someone else?
	loginIsBusy, _ = users.GetUUIDbyLogin(login, &uuidByLogin)
	if loginIsBusy {

		loginOwnerIsNotUs = (uuidByLogin != uuid)
		if loginOwnerIsNotUs {

			error_msg = fmt.Sprintf("Modify Error. An Attempt to take existing foreign 'login' [%s].", login)
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			return false, ret_err
		}
	}

	// 3. Check for empty regDate.
	regDateIsEmpty = (regDate == 0)
	if regDateIsEmpty {

		warning_msg = "Modify Warning. 'regDate' Field is not set in the Parameter."
		log.Println(warning_msg) //
		regDate = time.Now().Unix()

		// 'user' Parameter is not correct.
		// Create a correct temporary Parameter.
		tmp_user = new(UserModel)
		tmp_user.login = login
		tmp_user.regDate = regDate
		tmp_user.uuid = uuid

		// Modify User from temporary Object.
		(*users)[uuid] = *tmp_user

	} else {

		// 'user' Parameter is good.
		// Use it.

		// Modify User from Parameter.
		(*users)[uuid] = user
	}

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) ModifyAsRequest(user UserModel) (bool, error) {

	// Modifies a User, which was given as a Request.

	var ok bool
	var err error
	var ret_err error

	ok, err = users.Modify(user)

	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Save changes in Delta Map
	ModifiedUsersMap[user.uuid] = true

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) StoreAll() (bool, error) {

	// Saves all Users to DataBase.
	// The DB File is recreated from Zero and is totally rewritten.

	var tmp_path string

	var ok bool
	var err error
	var error_msg string
	var ret_err error

	// Find old File
	ok, err = db_find_file(db_path)
	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Make a back-up Copy
	ok, err = db_copy_file(db_path)
	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Check that Copy exists
	tmp_path = db_path + db_backup_postfix
	// Find a back-up File
	ok, err = db_find_file(tmp_path)
	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Find new File
	// Now it does not exist & will be created
	ok, err = db_find_file(db_path)
	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Save to DB
	ok, err = db_save_list(users)
	if !ok {

		error_msg = "Error. Can not store List. " + err.Error()
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Remove the temporary back-up Copy
	ok, err = db_remove_file(tmp_path)
	if !ok {

		ret_err = err

		return false, ret_err
	}

	//log.Printf("The DataBase has been saved.\r\n") // dbg

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) StoreDelta() (bool, error) {

	// Saves Changes to DataBase.
	// 1. The DB File is appended with newly added Users.
	// 2. The DB File is modified with modified Users.
	// While each User's Record has a constant Size, this Operation is very fast.

	var ok bool
	var err error
	var error_msg string
	var ret_err error

	// Find DB File
	ok, err = db_find_file(db_path)
	if !ok {

		ret_err = err

		return false, ret_err
	}

	// Save added Users to DB
	ok, err = db_save_added_users(&AddedUsersMap)
	if !ok {

		error_msg = "Error. Can not save added Users. " + err.Error()
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Clear List of added Users
	db_reset_added_users_map()

	// Save modified Users to DB
	ok, err = db_save_modified_users(&ModifiedUsersMap)
	if !ok {

		error_msg = "Error. Can not save modified Users. " + err.Error()
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// Clear List of modified Users
	db_reset_modified_users_map()

	//log.Printf("The DataBase has been updated.\r\n") // dbg

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) UUIDexists(uuid uint64) bool {

	// Checks if the specified 'uuid' is used in the List.

	var exists bool

	_, exists = (*users)[uuid]

	return exists
}

//------------------------------------------------------------------------------

func (users *UsersModel) debug_Print() {

	// Prints the Fields of all User Elements of the List.
	// This Function is used for testing Purposes.

	var cur_user UserModel

	fmt.Println("Users List.")                           //dbg
	fmt.Println("-------------------------------------") //dbg

	for _, cur_user = range *users {

		cur_user.debug_Print() //dbg
	}

	fmt.Println("-------------------------------------") //dbg
}

//------------------------------------------------------------------------------

func (user *UserModel) debug_Print() {

	// Prints the Fields of a User Element.
	// This Function is used for testing Purposes.

	fmt.Printf("User [%d][%d][%s].\r\n", user.uuid, user.regDate, user.login) //dbg
}

//------------------------------------------------------------------------------
