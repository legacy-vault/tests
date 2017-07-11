// db.go

// Version: 0.3.
// Date: 2017-07-07.
// Author: McArcher.

package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

//------------------------------------------------------------------------------
// Types
//------------------------------------------------------------------------------

type RecordsMapType map[uint64]uint64 // Map of UserRecords in File
// Key is UUID.
// Value is Record's Index (Position) in the File.
// Index (Position) starts from 0.

type AddedUsersMapType map[uint64]bool // Map of added Users.
// Key is UUID.
// Value is a Boolean Flag, which must be TRUE.
// Is used to store "Delta" (Changes) to prevent rewriting the whole DB File.

type ModifiedUsersMapType map[uint64]bool // Map of added Users.
// Key is UUID.
// Value is a Boolean Flag, which must be TRUE.
// Is used to store "Delta" (Changes) to prevent rewriting the whole DB File.

//------------------------------------------------------------------------------
// Variables
//------------------------------------------------------------------------------

var db_path = "db/db.dat"
var db_createMissingFile bool = true

var RecordsMap RecordsMapType
var AddedUsersMap AddedUsersMapType
var ModifiedUsersMap ModifiedUsersMapType

//------------------------------------------------------------------------------
// Data Format
//------------------------------------------------------------------------------
/*
	[Record]

	Field:	[uuid]	[regDate]	[login_len]	[login]
	Type:	uint64	int64		uint8		string
	Size:	8		8			1			239

	Size in Bytes.
	Total Size of one Record is 256 Bytes.
*/
//------------------------------------------------------------------------------
// Constants
//------------------------------------------------------------------------------

const db_field_uuid_maxlen = 8
const db_field_regdate_maxlen = 8
const db_field_loginlen_maxlen = 1
const db_field_login_maxlen = 239
const db_record_len = 256
const db_field_loginlen_maxval = 255

const db_byte_background byte = byte(' ')
const db_backup_postfix = ".bak"

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func db_analyze_read_error(err error, bytesRead int) (ok bool, eof bool) {

	// Analyzes the Error during Read Process.
	//
	// Returned Values:
	// 'ok', if TRUE, shows that Read Process goes normally.
	// 'eof', if TRUE, shows that EOF is reached.

	// No Error?
	if err == nil {

		// No error
		// How many Bytes we have read?
		if bytesRead != db_record_len {

			log.Println(err.Error()) //

			return false, false
		}

		return true, false
	}

	// Error exists
	// EOF?
	if err == io.EOF {

		if bytesRead == 0 {

			return true, true

		} else {

			log.Println(err.Error()) //

			return false, true
		}
	}

	log.Println(err.Error()) //

	return false, false
}

//------------------------------------------------------------------------------

func db_find_file(path string) (bool, error) {

	// Searches for the File.
	//
	// Returns TRUE either if File exists or have been created.
	// Otherwise returns FALSE.

	var file *os.File
	var exists bool

	var err error
	var error_msg string
	var ret_err error

	// File exists?
	exists = db_path_exists(path)

	if !exists {

		// File does not exist

		// Do we need to create a new File?
		if !db_createMissingFile {

			error_msg = fmt.Sprintf("File does not exist and will not be created! [%s].", path)
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			return false, ret_err
		}

		// Create a new File
		file, err = os.OpenFile(path, os.O_CREATE, 0755)
		if err != nil {

			error_msg = fmt.Sprintf("Error. Can not create File [%s]. %s", path, err.Error())
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			return false, ret_err
		}
		file.Close()
	}

	// File exists or have been created

	return true, nil
}

//------------------------------------------------------------------------------

func db_init() {

	// Initializes the DB.

	RecordsMap = make(RecordsMapType)
	AddedUsersMap = make(AddedUsersMapType)
	ModifiedUsersMap = make(ModifiedUsersMapType)
}

//------------------------------------------------------------------------------

func db_load_list(list *UsersModel) (bool, error) {

	// Loads the List of Users from DataBase.

	var file *os.File
	var buffer []byte
	var bytesRead int
	var uuid uint64
	var regDate int64
	var login string
	var user UserModel
	var cur_recnum uint64 // Number (Index) of the current Record in File
	var file_info os.FileInfo
	var file_size int64
	var file_recnum int64
	var rec_interval uint64
	var msg_welcome, msg_done, msg_delta string

	var err error
	var ok, eof bool
	var error_msg string
	var ret_err error

	// Prepare Read Buffer
	buffer = make([]byte, db_record_len)

	// Open File for Reading
	file, err = os.OpenFile(db_path, os.O_RDONLY, 0755)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not open File [%s]. %s", db_path, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}
	defer file.Close()

	// Read File Size
	file_info, err = file.Stat()
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not get File Statistics [%s]. %s", db_path, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}
	file_size = file_info.Size()
	file_recnum = file_size / db_record_len

	// Show welcome Message
	msg_welcome = "Loading DataBase (" + strconv.FormatInt(file_recnum, 10) + " Records)."
	msg_delta = "."
	msg_done = "[DONE]"
	fmt.Print(msg_welcome)
	rec_interval = 10 * 1000 // Information Update Interval (Number of Records)

	// Read Records from File
	cur_recnum = 0
	for {

		// Show Progress
		if cur_recnum%rec_interval == 0 {
			fmt.Print(msg_delta) //
		}

		// Read one Record
		bytesRead, err = io.ReadFull(file, buffer)

		// Analyze Read Results
		ok, eof = db_analyze_read_error(err, bytesRead)
		if !ok {
			// Something bad has happened
			ret_err = err

			return false, ret_err
		}
		if eof {
			// EOF is a normal Thing
			break
		}

		// Parse Data
		db_parse_buffer(buffer, &uuid, &regDate, &login)

		// Add Data to List in Memory
		user.uuid = uuid
		user.regDate = regDate
		user.login = login
		ok, err = list.Add(user, true)
		if !ok {

			error_msg = "Error. Can not Add an Element to List. " + err.Error()
			ret_err = errors.New(error_msg)
			log.Println(error_msg) //

			return false, ret_err
		}

		// Save Record's Index to RecordsMap
		RecordsMap[uuid] = cur_recnum

		// -> Next Record
		cur_recnum++
	}

	fmt.Println(msg_done) //

	return true, nil
}

//------------------------------------------------------------------------------

func db_modify_record(uuid uint64, regDate int64, login *string, file *os.File) (bool, error) {

	// Modifies a Record in the File.

	var record_index uint64
	var offset int64

	var ok bool
	var err error
	var error_msg string
	var ret_err error

	// Get Index of a Record
	record_index = RecordsMap[uuid]

	// Seek Record
	offset = int64(record_index) * db_record_len
	file.Seek(offset, 0)

	// Modify Record
	ok, err = db_write_record(uuid, regDate, login, file)

	if !ok {

		error_msg = "Record Modification failed. " + err.Error()
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	return true, nil
}

//------------------------------------------------------------------------------

func db_parse_buffer(buffer []byte, uuid *uint64, regDate *int64, login *string) {

	// Reads Data from Buffer and puts it into Variables.

	var login_len uint8
	var login_bytes []byte
	var i, j uint8

	// 1. UUID
	*uuid = uint64(buffer[0])<<56 + uint64(buffer[1])<<48 +
		uint64(buffer[2])<<40 + uint64(buffer[3])<<32 +
		uint64(buffer[4])<<24 + uint64(buffer[5])<<16 +
		uint64(buffer[6])<<8 + uint64(buffer[7])
	// 2. regDate
	*regDate = int64(buffer[8])<<56 + int64(buffer[9])<<48 +
		int64(buffer[10])<<40 + int64(buffer[11])<<32 +
		int64(buffer[12])<<24 + int64(buffer[13])<<16 +
		int64(buffer[14])<<8 + int64(buffer[15])
	// 3. login_len
	login_len = uint8(buffer[16])
	// 4. login
	login_bytes = make([]byte, login_len)
	j = 17
	for i = 0; i < login_len; i++ {
		login_bytes[i] = buffer[j]
		j++
	}
	*login = string(login_bytes)
}

//------------------------------------------------------------------------------

func db_path_exists(path string) bool {

	// Checks if the Path exists in the FileSystem.

	var err error

	_, err = os.Stat(path)

	if err != nil {

		// Error
		if os.IsNotExist(err) {

			return false
		}
	}

	// No Error or Error other than 'File does not exist'

	return true
}

//------------------------------------------------------------------------------

func db_reset_records_map() {

	// Resets the RecordsMap.
	// Is used when the DB File has been read and happened to be corrupted.

	RecordsMap = nil // Free Memory
	RecordsMap = make(RecordsMapType)
}

//------------------------------------------------------------------------------

func db_save_added_users(users *AddedUsersMapType) (bool, error) {

	// Saves the List of Users newly added to DataBase.

	var file *os.File
	var cur_enabled bool
	var cur_uuid uint64
	var cur_recnum uint64
	var regDate int64
	var login string

	var err error
	var ok bool
	var error_msg string
	var ret_err error

	// Open File for Appending
	file, err = os.OpenFile(db_path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not open File [%s]. %s.", db_path, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}
	defer file.Close()

	// Append Records to File
	cur_recnum = uint64(len(RecordsMap))

	for cur_uuid, cur_enabled = range *users {

		if !cur_enabled {

			// Delete already added User from AddedUsersMap
			delete(AddedUsersMap, cur_uuid)
			continue
		}

		regDate = Users[cur_uuid].regDate
		login = Users[cur_uuid].login

		// Append one Record to File
		ok, err = db_write_record(cur_uuid, regDate, &login, file)
		if !ok {

			ret_err = err

			return false, ret_err
		}

		// Append one Record to RecordsMap
		RecordsMap[cur_uuid] = cur_recnum
		cur_recnum++

		// Delete already added User from AddedUsersMap
		delete(AddedUsersMap, cur_uuid)
	}

	return true, nil
}

//------------------------------------------------------------------------------

func db_save_modified_users(users *ModifiedUsersMapType) (bool, error) {

	// Saves the List of Users recently modified.

	var file *os.File
	var cur_enabled bool
	var cur_uuid uint64
	var regDate int64
	var login string

	var err error
	var ok bool
	var error_msg string
	var ret_err error

	// Open File for Writing
	file, err = os.OpenFile(db_path, os.O_RDWR, 0666)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not open File [%s]. %s.", db_path, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}
	defer file.Close()

	// Modify Records in File
	for cur_uuid, cur_enabled = range *users {

		if !cur_enabled {

			// Delete already modified User from ModifiedUsersMap
			delete(ModifiedUsersMap, cur_uuid)
			continue
		}

		regDate = Users[cur_uuid].regDate
		login = Users[cur_uuid].login

		// Modify one Record in File
		ok, err = db_modify_record(cur_uuid, regDate, &login, file)
		if !ok {

			ret_err = err

			return false, ret_err
		}

		// Delete already modified User from ModifiedUsersMap
		delete(ModifiedUsersMap, cur_uuid)
	}

	return true, nil
}

//------------------------------------------------------------------------------

func db_write_record(uuid uint64, regDate int64, login *string, file *os.File) (bool, error) {

	// Writes a Record to the File.

	var login_len_int int
	var login_len uint8
	var login_bytes, login_buf []byte
	var i, bytesWritten int

	var err error
	var error_msg string
	var ret_err error

	// Length of login
	login_len_int = len(*login)

	if login_len_int > db_field_loginlen_maxval {

		error_msg = fmt.Sprintf("Error. Login string is too long! [%s].", login)
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	login_len = uint8(login_len_int)

	// Prepare login Byte Array
	login_bytes = []byte(*login)
	login_buf = make([]byte, db_field_login_maxlen)
	for i, _ = range login_buf {
		login_buf[i] = db_byte_background
	}
	for i, _ = range login_bytes {
		login_buf[i] = login_bytes[i]
	}

	// 1. Write UUID
	err = binary.Write(file, binary.BigEndian, uuid)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not write to File [%s]. %s.", file, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// 2. Write regDate
	err = binary.Write(file, binary.BigEndian, regDate)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not write to File [%s]. %s.", file, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// 3. Write login_len
	err = binary.Write(file, binary.BigEndian, login_len)
	if err != nil {

		error_msg = fmt.Sprintf("Error. Can not write to File [%s]. %s.", file, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	// 4. Write login Byte Array
	bytesWritten, err = file.Write(login_buf)
	if (err != nil) || (bytesWritten != len(login_buf)) {

		error_msg = fmt.Sprintf("Error. Can not write to File [%s]. %s.", file, err.Error())
		ret_err = errors.New(error_msg)
		log.Println(error_msg) //

		return false, ret_err
	}

	return true, nil
}

//------------------------------------------------------------------------------
