// model_user.go

// Version: 0.4.
// Date: 2017-07-08.
// Author: McArcher.

package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
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

type LoginsMapType map[string]uint64 // Map of 'login' Fields
// Key is 'login'.
// Value is UUID.

type UsersIntegrityCheckType struct {
	StageOneEnabled   bool
	StageTwoEnabled   bool
	StageThreeEnabled bool
	ShowStages        bool
}

type UserTask struct {
	user   UserModel
	sender chan UserTask
	result bool
	err    error
}

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var Users UsersModel
var LoginsMap LoginsMapType

var UsersIntegrityCheck UsersIntegrityCheckType

// Channels
var AddedUsersChan chan UserTask
var ModifiedUsersChan chan UserTask
var GotUsersChan chan UserTask

var AddedUsersChanSize int
var ModifiedUsersChanSize int
var GotUsersChanSize int

//------------------------------------------------------------------------------
// Methods
//------------------------------------------------------------------------------

func (users *UsersModel) Add(user UserModel, silent bool) (bool, error) {

	// Adds a User to the List.

	// 'silent' Flag is used to disable Registration of added Users in
	// AddedUsersMap. When User is added to AddedUsersMap, User is then
	// appended to the File. When we add Users from File, we will get the File
	// Contents doubled. So, to add Users at start-up (during File Read),
	// Addition to AddedUsersMap must be disabled (silent = true).

	var uuid uint64
	var uuidIsBusy bool
	var login string
	var loginIsBusy bool

	var error_msg string
	var ret_err error

	// Accept a Request
	RequestsWaitGroup.Add(1)

	uuid = user.uuid
	login = user.login

	// 1. Check if UUID is already taken.
	uuidIsBusy = isUUIDbusy(uuid)
	if uuidIsBusy {

		error_msg = "UUID is busy."
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		// Request is Done
		RequestsWaitGroup.Done()

		return false, ret_err
	}

	// 2. Check if login is already taken.
	loginIsBusy = isLoginBusy(login)
	if loginIsBusy {

		error_msg = "Login is busy."
		ret_err = errors.New(error_msg)
		log.Printf(error_msg) //

		// Request is Done
		RequestsWaitGroup.Done()

		return false, ret_err
	}

	// Add User to Users List
	(*users)[uuid] = user

	// Add Login to LoginsMap
	LoginsMap[login] = uuid

	// Save Changes in Delta Map
	if !silent {
		AddedUsersMap[uuid] = true
	}

	// Request is Done
	RequestsWaitGroup.Done()

	return true, nil
}

//------------------------------------------------------------------------------

func (users *UsersModel) CheckIntegrity() (bool, error) {

	// Checks the Integrity of all Elements in the List.
	// This Function is used for testing Purposes.
	// Returns TRUE if Check is successful.

	var cur_user, cur_user_2 UserModel
	var cur_index uint64
	var cur_uuid, cur_uuid_2 uint64
	var cur_login, cur_login_2 string
	var cur_regDate int64

	var indexIsBad bool
	var uuidIsBad bool
	var loginIsBad bool
	var regDateIsBad bool
	var loginIsTooLong bool
	var errorIsFound bool
	var count uint64
	var msg_welcome, msg_stage_1, msg_stage_2, msg_stage_3, msg_success string
	var msg_done, msg_skipped string
	var recnum, cur_rec int // Number of all Records (Users)  & current one
	var rec_interval int

	var error_msg string
	var ret_err error
	var err_log string

	errorIsFound = false
	err_log = ""

	recnum = len(*users)
	msg_welcome = "Integrity Check DataBase (" + strconv.Itoa(recnum) + " Records)."
	msg_stage_1 = "Integrity Check : Stage I. "
	msg_stage_2 = "Integrity Check : Stage II. "
	msg_stage_3 = "Integrity Check : Stage III. "
	msg_success = "No Errors were found."
	msg_done = "[DONE]"
	msg_skipped = "[SKIPPED]"
	fmt.Println(msg_welcome) //
	rec_interval = 1 * 1000  // Information Update Interval (Number of Records)

	// 1. Basic Checks.

	if UsersIntegrityCheck.ShowStages {
		fmt.Print(msg_stage_1) //
	}
	if UsersIntegrityCheck.StageOneEnabled {

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
			loginIsTooLong = !isLoginLengthGood(&cur_login)
			if loginIsTooLong {

				err_log = err_log + fmt.Sprintf("Too long login [%s].", cur_login)
				errorIsFound = true
			}
		}
		if UsersIntegrityCheck.ShowStages {
			fmt.Println(msg_done) //
		}

	} else {

		if UsersIntegrityCheck.ShowStages {
			fmt.Println(msg_skipped) //
		}
	}

	// 2. Check duplicate UUIDs.

	// This Check is needed only to test Data in Memory.
	// When the List is loaded from File at Application Start, the Duplicates
	// in the List are impossible while the 'Add' Function already checks for
	// Duplicates (existing Elements). So, at Start (Load) this Test has no
	// Sense! It is useful only for finding broken memory Cells or some other
	// parasitic Activity in Memory of the Server. In most Cases this Test is
	// just of the 'for any Case' Type and should be disabled (skipped).

	if UsersIntegrityCheck.ShowStages {
		fmt.Print(msg_stage_2) //
	}
	if UsersIntegrityCheck.StageTwoEnabled {

		cur_rec = 1
		for _, cur_user = range *users {

			// Show Progress
			if cur_rec%rec_interval == 0 {
				fmt.Print(".") //
			}

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

			cur_rec++
		}
		if UsersIntegrityCheck.ShowStages {
			fmt.Println(msg_done) //
		}

	} else {

		if UsersIntegrityCheck.ShowStages {
			fmt.Println(msg_skipped) //
		}
	}

	// 3. Check duplicate logins.

	// This Check is needed only to test Data in Memory.
	// When the List is loaded from File at Application Start, the Duplicates
	// in the List are impossible while the 'Add' Function already checks for
	// Duplicates (existing Elements). So, at Start (Load) this Test has no
	// Sense! It is useful only for finding broken memory Cells or some other
	// parasitic Activity in Memory of the Server. In most Cases this Test is
	// just of the 'for any Case' Type and should be disabled (skipped).

	if UsersIntegrityCheck.ShowStages {
		fmt.Print(msg_stage_3) //
	}
	if UsersIntegrityCheck.StageThreeEnabled {

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
		if UsersIntegrityCheck.ShowStages {
			fmt.Println(msg_done) //
		}

	} else {

		if UsersIntegrityCheck.ShowStages {
			fmt.Println(msg_skipped) //
		}
	}

	// Summary.
	if errorIsFound {

		error_msg = "Integrity is BAD." + err_log
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	fmt.Println(msg_success) //

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
	var uuidByLogin uint64
	var uuidIsSet bool
	var uuidIsFree bool
	var login string
	var loginIsFree bool
	var loginIsSet bool

	var error_msg string
	var ret_err error
	var ok bool
	var err error

	// Accept a Request
	RequestsWaitGroup.Add(1)

	uuid = user.uuid
	login = user.login

	// 1. Get User by UUID if it is specified.

	// Check for empty UUID.
	uuidIsSet = (uuid != 0)
	if uuidIsSet {

		// Check if UUID is free.
		uuidIsFree = isUUIDfree(uuid)
		if uuidIsFree {

			error_msg = "UUID does not exist."
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			// Request is Done
			RequestsWaitGroup.Done()

			return false, ret_err
		}

		// UUID is set & exists.
		*reply = (*users)[uuid]

		// Request is Done
		RequestsWaitGroup.Done()

		return true, nil
	}

	// 2. Get User by login if it is specified.

	// Check for empty login.
	loginIsSet = (login != "")
	if loginIsSet {

		// Check if login is already taken.
		loginIsFree = isLoginFree(login)
		if loginIsFree {

			error_msg = "Login does not exist"
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			// Request is Done
			RequestsWaitGroup.Done()

			return false, ret_err
		}

		// login is set & exists.
		ok, err = getUUIDbyLogin(login, &uuidByLogin)
		if !ok {

			error_msg = fmt.Sprintf("Can not get UUID by login. %s", err.Error())
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			// Request is Done
			RequestsWaitGroup.Done()

			return false, ret_err
		}

		*reply = (*users)[uuidByLogin]

		// Request is Done
		RequestsWaitGroup.Done()

		return true, nil
	}

	// 3. Both 'UUID' & 'login' are not specified.

	error_msg = "Both UUID & login are empty."
	ret_err = errors.New(error_msg)
	log.Println(error_msg) //

	// Request is Done
	RequestsWaitGroup.Done()

	return false, ret_err
}

//------------------------------------------------------------------------------

func getUUIDbyLogin(login string, uuid *uint64) (bool, error) {

	// Gets an 'UUID' of a User with a specified 'login'.
	//
	// Returns FALSE if:
	//		- 'login' is empty;
	//		- 'login' is not found.

	var loginIsEmpty bool
	var loginExists bool

	var error_msg string
	var ret_err error

	// Fool Check.
	loginIsEmpty = (login == "")
	if loginIsEmpty {

		error_msg = "Login is empty."
		ret_err = errors.New(error_msg)
		//log.Println(error_msg) // Do NOT log this Spam as it is very often used

		return false, ret_err
	}

	_, loginExists = LoginsMap[login]

	if loginExists {

		*uuid = LoginsMap[login]

		return true, nil
	}

	// 'login' is not found.
	error_msg = "Login is not found."
	ret_err = errors.New(error_msg)
	//log.Println(error_msg) // Do NOT log this Spam as it is very often used

	return false, ret_err
}

//------------------------------------------------------------------------------

func (users *UsersModel) Init() {

	// Initializes the Model.

	AddedUsersChanSize = 1000
	ModifiedUsersChanSize = 1000
	GotUsersChanSize = 1000

	AddedUsersChan = make(chan UserTask, AddedUsersChanSize)
	ModifiedUsersChan = make(chan UserTask, ModifiedUsersChanSize)
	GotUsersChan = make(chan UserTask, GotUsersChanSize)
	ManagerAliveChan = make(chan bool)

	*users = make(UsersModel)

	LoginsMap = make(LoginsMapType)

	UsersIntegrityCheck.StageOneEnabled = true
	UsersIntegrityCheck.StageTwoEnabled = false   // Read Comments in 'CheckIntegrity' !
	UsersIntegrityCheck.StageThreeEnabled = false // Read Comments in 'CheckIntegrity' !
	UsersIntegrityCheck.ShowStages = true
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

func isLoginBusy(login string) bool {

	// Checks if the specified 'login' is already taken.

	var loginExists bool

	_, loginExists = LoginsMap[login]

	return loginExists
}

//------------------------------------------------------------------------------

func isLoginFree(login string) bool {

	// Checks if the specified 'login' is not used in the List.

	var loginExists bool

	_, loginExists = LoginsMap[login]

	if loginExists {

		return false
	}

	return true
}

//------------------------------------------------------------------------------

func isLoginLengthGood(login *string) bool {

	// Checks if the Length of 'login' is good enough to be stored into the DB.

	var len_is_ok bool

	len_is_ok = (len(*login) <= db_field_login_maxlen)

	return len_is_ok
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
	var uuidIsFree bool
	var uuidByLogin uint64

	var login string
	var login_old string
	var loginIsBusy bool
	var loginOwnerIsNotUs bool
	var loginHasChanged bool

	var error_msg string
	var ret_err error

	// Accept a Request
	RequestsWaitGroup.Add(1)

	uuid = user.uuid
	login = user.login

	// 1. Fake UUID?
	uuidIsFree = isUUIDfree(uuid)
	if uuidIsFree {

		error_msg = "UUID does not exist"
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		// Request is Done
		RequestsWaitGroup.Done()

		return false, ret_err
	}

	// 2. Login Collision Check.
	// 'login' is already taken by someone else?
	loginIsBusy, _ = getUUIDbyLogin(login, &uuidByLogin)
	if loginIsBusy {

		loginOwnerIsNotUs = (uuidByLogin != uuid)
		if loginOwnerIsNotUs {

			error_msg = fmt.Sprintf("An Attempt to take existing foreign login [%s].", login)
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			// Request is Done
			RequestsWaitGroup.Done()

			return false, ret_err
		}
	}

	// Save previous Login & analyze it
	login_old = (*users)[uuid].login
	loginHasChanged = (login_old != login)

	(*users)[uuid] = user

	// Modify Login in LoginsMap if it has changed
	if loginHasChanged {
		// Delete old Mapping
		delete(LoginsMap, login_old)
		// Create new Mapping
		LoginsMap[login] = uuid
	}

	// Save changes in Delta Map
	ModifiedUsersMap[uuid] = true

	// Request is Done
	RequestsWaitGroup.Done()

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

	// Save modified Users to DB
	ok, err = db_save_modified_users(&ModifiedUsersMap)
	if !ok {

		error_msg = "Error. Can not save modified Users. " + err.Error()
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	//log.Printf("The DataBase has been updated.\r\n") // dbg

	return true, nil
}

//------------------------------------------------------------------------------

func isUUIDbusy(uuid uint64) bool {

	// Checks if the specified UUID is used in the List.

	var exists bool

	_, exists = Users[uuid]

	return exists
}

//------------------------------------------------------------------------------

func isUUIDfree(uuid uint64) bool {

	// Checks if the specified UUID is used in the List.

	var exists bool

	_, exists = Users[uuid]

	return !exists
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

func debug_PrintLoginsMap() {

	// Prints LoginsMap.
	// This Function is used for testing Purposes.

	var cur_uuid uint64
	var cur_login string

	fmt.Println("LoginsMap")                             //dbg
	fmt.Println("-------------------------------------") //dbg
	fmt.Println(" [login] [UUID]")                       //dbg
	fmt.Println("-------------------------------------") //dbg

	for cur_login, cur_uuid = range LoginsMap {

		fmt.Printf("[%s][%d].\r\n", cur_login, cur_uuid) //dbg
	}

	fmt.Println("-------------------------------------") //dbg
}

//------------------------------------------------------------------------------
