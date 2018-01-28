// binary_frame.go

/*

	Binary Frame Tool.

	Version: 0.1.1.
	Date of Creation: 2018-01-28.
	Author: McArcher.

	This is a simple Tool which draws a binary Frame around the Content.
	A Frame consists of logical Ones (1) and has a Spacer of Zeroes (0).

	So,
	---
	XXX
	XXX
	---

	Becomes something with a binary Frame with a Spacer:
	-------
	1111111
	1000001
	10XXX01
	10XXX01
	1000001
	1111111
	-------

	This Technique is a great Thing to enclose Dimensions of the "X" into the
	File. It may be useful for Transmission of Signals in Space or other Places
	with a great Chances for Signal Corruption. Even when damaged, a Frame or
	Remains of a Frame may help a lot in the Process of Data Recovery. This
	Frame may be used as a Transport Package when there is no Package of the
	Transmission Protocol, and may be used in a universal Range.

	To comply with this-day Computers, the Algorithm makes Sizes compatible
	with 8-Bit Bytes, but it is also able to use any Size you want.

*/

//==============================================================================

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

//==============================================================================

type bit bool
type row_of_bits []bit
type field_of_bits []row_of_bits

//==============================================================================

const bits_in_byte = 8
const FILLED = true
const EMPTY = false

const ACTION_NONE = 0
const ACTION_ENCODE_F1 = 1
const ACTION_ENCODE_F2 = 2
const ACTION_DECODE_F1 = 3
const ACTION_DECODE_F2 = 4

const ERROR_1 = 1

//==============================================================================

var cla_file_in *string
var cla_file_out *string
var cla_action_type *string
var cla_x *string
var cla_y *string

var action_type uint8
var file_input_path string
var file_output_path string
var file_input_content []byte
var file_output_content []byte
var file_input_x uint64
var file_input_y uint64
var file_input_size uint64   // Number of Bytes.
var input_field_size uint64  // Number of Bits.
var output_field_size uint64 // Number of Bits.
var file_input_cols uint64
var file_input_rows uint64
var file_output_cols uint64
var file_output_rows uint64

var field field_of_bits
var field_framed field_of_bits

//==============================================================================

func main() {

	var err error
	var err_code uint8
	var ok bool

	// Command Line Arguments.
	read_cla()

	// Read Input File.
	file_input_content, err = ioutil.ReadFile(file_input_path)
	check_error(err)
	file_input_size = uint64(len(file_input_content))
	input_field_size = file_input_size * bits_in_byte

	// Check X & Y.
	if file_input_x <= 0 {
		fmt.Println("Bad X Size.")
		os.Exit(ERROR_1)
	}
	file_input_cols = file_input_x
	if (input_field_size % file_input_x) != 0 {
		fmt.Println("Bad X Size.")
		os.Exit(ERROR_1)
	}
	file_input_rows = input_field_size / file_input_x

	// Do Action.
	if action_type == ACTION_ENCODE_F1 {
		fmt.Println("Encoding (F1) \"" + file_input_path + "\"...") //

		// Output Size.
		file_output_cols = file_input_cols + 4
		file_output_rows = file_input_rows + 4
		output_field_size = file_output_cols * file_output_rows

		// Report.
		fmt.Println("Input Data (WxH):", file_input_cols, "x", file_input_rows,
			".") ///
		fmt.Println("Output Data (WxH):", file_output_cols, "x",
			file_output_rows, ".") ///

		// Bytes -> Field.
		field, ok = bytes_to_field(input_field_size, file_input_cols,
			file_input_rows, file_input_content)
		check_ok(ok)
		// Field -> Frame.
		field_framed, err_code = pack_data_f1(input_field_size,
			file_input_cols, file_input_rows, field)
		check_err_code(err_code)
		// Frame -> Bytes.
		file_output_content, ok = field_to_bytes(output_field_size,
			file_output_cols, file_output_rows, field_framed)
		check_ok(ok)
		// Bytes -> File.
		fmt.Println("Writing \"" + file_output_path + "\"...") //
		err = ioutil.WriteFile(file_output_path, file_output_content, 0644)
		check_error(err)
	}

	if action_type == ACTION_ENCODE_F2 {
		fmt.Println("Encoding (F2) \"" + file_input_path + "\"...") //

		// Output Size.
		file_output_cols = file_input_cols + 8
		file_output_rows = file_input_rows + 8
		output_field_size = file_output_cols * file_output_rows

		// Report.
		fmt.Println("Input Data (WxH):", file_input_cols, "x", file_input_rows,
			".") ///
		fmt.Println("Output Data (WxH):", file_output_cols, "x",
			file_output_rows, ".") ///

		// Bytes -> Field.
		field, ok = bytes_to_field(input_field_size, file_input_cols,
			file_input_rows, file_input_content)
		check_ok(ok)
		// Field -> Frame.
		field_framed, err_code = pack_data_f2(input_field_size,
			file_input_cols, file_input_rows, field)
		check_err_code(err_code)
		// Frame -> Bytes.
		file_output_content, ok = field_to_bytes(output_field_size,
			file_output_cols, file_output_rows,
			field_framed)
		check_ok(ok)
		// Bytes -> File.
		fmt.Println("Writing \"" + file_output_path + "\"...") //
		err = ioutil.WriteFile(file_output_path, file_output_content, 0644)
		check_error(err)
	}

	if action_type == ACTION_DECODE_F1 {
		fmt.Println("Decoding (F1) \"" + file_input_path + "\"...") //

		// Output Size.
		file_output_cols = file_input_cols - 4
		file_output_rows = file_input_rows - 4
		output_field_size = file_output_cols * file_output_rows

		// Report.
		fmt.Println("Input Data (WxH):", file_input_cols, "x", file_input_rows,
			".") ///
		fmt.Println("Output Data (WxH):", file_output_cols, "x",
			file_output_rows, ".") ///

		// Bytes -> Frame.
		field_framed, ok = bytes_to_field(input_field_size, file_input_cols,
			file_input_rows, file_input_content)
		check_ok(ok)
		// Frame -> Field.
		field, ok = get_data_f1(input_field_size, file_input_cols,
			file_input_rows, field_framed)
		check_ok(ok)
		// Field -> Bytes.
		file_output_content, ok = field_to_bytes(output_field_size,
			file_output_cols, file_output_rows, field)
		check_ok(ok)
		// Bytes -> File.
		fmt.Println("Writing \"" + file_output_path + "\"...") //
		err = ioutil.WriteFile(file_output_path, file_output_content, 0644)
		check_error(err)
	}

	if action_type == ACTION_DECODE_F2 {
		fmt.Println("Decoding (F2) \"" + file_input_path + "\"...") //

		// Output Size.
		file_output_cols = file_input_cols - 8
		file_output_rows = file_input_rows - 8
		output_field_size = file_output_cols * file_output_rows

		// Report.
		fmt.Println("Input Data (WxH):", file_input_cols, "x", file_input_rows,
			".") ///
		fmt.Println("Output Data (WxH):", file_output_cols, "x",
			file_output_rows, ".") ///

		// Bytes -> Frame.
		field_framed, ok = bytes_to_field(input_field_size, file_input_cols,
			file_input_rows, file_input_content)
		check_ok(ok)
		// Frame -> Field.
		field, ok = get_data_f2(input_field_size, file_input_cols,
			file_input_rows, field_framed)
		check_ok(ok)
		// Field -> Bytes.
		file_output_content, ok = field_to_bytes(output_field_size,
			file_output_cols, file_output_rows, field)
		check_ok(ok)
		// Bytes -> File.
		fmt.Println("Writing \"" + file_output_path + "\"...") //
		err = ioutil.WriteFile(file_output_path, file_output_content, 0644)
		check_error(err)
	}

	if action_type == ACTION_NONE {
		fmt.Println("Idle...") //
	}

}

//==============================================================================

// Packs useful Data into Message and surrounds it with a Frame I.

func pack_data_f1(
	data_bits_count uint64,
	data_columns_count uint64,
	data_rows_count uint64,
	data field_of_bits) (field_of_bits, uint8) {

	const DS = 4
	const DO = DS / 2
	const data_columns_count_limit = math.MaxUint64 - DS
	const data_rows_count_limit = math.MaxUint64 - DS

	const ERROR_ALL_CLEAR = 0     // No Error.
	const ERROR_BAD_SIZE = 1      // (Colums * Rows) ≠ (Bit Count).
	const ERROR_COLUMNS_ERROR = 2 // Too many Columns in Data.
	const ERROR_ROWS_ERROR = 3    // Too many Rows in Data.

	var data_bits_count_required uint64
	var result field_of_bits

	// Cursors in Result.
	var i uint64     // Current Row #.
	var i_max uint64 //
	var i_min uint64 //
	var j uint64     // Current Column #.
	var j_max uint64 //
	var j_min uint64 //

	// Cursors in Data.
	var y uint64 // Current Row #.
	var x uint64 // Current Column #.

	var result_columns_count uint64
	var result_rows_count uint64

	var data_first_column_index uint64
	var data_first_row_index uint64
	//var data_last_column_index uint64
	//var data_last_row_index uint64

	var result_first_column_index uint64
	var result_first_row_index uint64
	var result_last_column_index uint64
	var result_last_row_index uint64

	// Check Input Data.
	data_bits_count_required = data_columns_count * data_rows_count
	if data_bits_count != data_bits_count_required {
		return nil, ERROR_BAD_SIZE
	}
	if data_columns_count > data_columns_count_limit {
		return nil, ERROR_COLUMNS_ERROR
	}
	if data_rows_count > data_rows_count_limit {
		return nil, ERROR_ROWS_ERROR
	}

	// Indices & Sizes.
	result_columns_count = data_columns_count + DS
	result_rows_count = data_rows_count + DS

	data_first_column_index = 0
	data_first_row_index = 0
	//data_last_column_index = data_columns_count - 1
	//data_last_row_index = data_rows_count - 1

	result_first_column_index = 0
	result_first_row_index = 0
	result_last_column_index = result_columns_count - 1
	result_last_row_index = result_rows_count - 1

	// Create an empty Field.
	result = make(field_of_bits, result_rows_count)
	for i = result_first_row_index; i <= result_last_row_index; i++ {
		result[i] = make(row_of_bits, result_columns_count)
		for j = result_first_column_index; j <= result_last_column_index; j++ {
			result[i][j] = EMPTY
		}
	}

	// Draw the Frame I.
	for j = result_first_column_index; j <= result_last_column_index; j++ {
		result[result_first_row_index][j] = FILLED
		result[result_last_row_index][j] = FILLED
	}
	for i = result_first_row_index; i <= result_last_row_index; i++ {
		result[i][result_first_column_index] = FILLED
		result[i][result_last_column_index] = FILLED
	}
	// Draw Frame's Spacer.
	i_min = result_first_row_index + 1
	i_max = result_last_row_index - 1
	j_min = result_first_column_index + 1
	j_max = result_last_column_index - 1
	for j = j_min; j <= j_max; j++ {
		result[i_min][j] = EMPTY
		result[i_max][j] = EMPTY
	}
	for i = i_min; i <= i_max; i++ {
		result[i][j_min] = EMPTY
		result[i][j_max] = EMPTY
	}

	// Draw Data.
	i_min = result_first_row_index + DO
	i_max = result_last_row_index - DO
	j_min = result_first_column_index + DO
	j_max = result_last_column_index - DO
	y = data_first_row_index
	for i = i_min; i <= i_max; i++ {
		x = data_first_column_index
		for j = j_min; j <= j_max; j++ {
			result[i][j] = data[y][x]
			x++
		}
		y++
	}

	return result, ERROR_ALL_CLEAR
}

//==============================================================================

// Packs useful Data into Message and surrounds it with a Frame II.

func pack_data_f2(
	data_bits_count uint64,
	data_columns_count uint64,
	data_rows_count uint64,
	data field_of_bits) (field_of_bits, uint8) {

	const DS = 8
	const DO = DS / 2
	const data_columns_count_limit = math.MaxUint64 - DS
	const data_rows_count_limit = math.MaxUint64 - DS

	const ERROR_ALL_CLEAR = 0     // No Error.
	const ERROR_BAD_SIZE = 1      // (Colums * Rows) ≠ (Bit Count).
	const ERROR_COLUMNS_ERROR = 2 // Too many Columns in Data.
	const ERROR_ROWS_ERROR = 3    // Too many Rows in Data.

	var data_bits_count_required uint64
	var result field_of_bits

	// Cursors in Result.
	var i uint64     // Current Row #.
	var i_max uint64 //
	var i_min uint64 //
	var j uint64     // Current Column #.
	var j_max uint64 //
	var j_min uint64 //

	// Cursors in Data.
	var y uint64 // Current Row #.
	var x uint64 // Current Column #.

	var result_columns_count uint64
	var result_rows_count uint64

	var data_first_column_index uint64
	var data_first_row_index uint64
	//var data_last_column_index uint64
	//var data_last_row_index uint64

	var result_first_column_index uint64
	var result_first_row_index uint64
	var result_last_column_index uint64
	var result_last_row_index uint64

	// Check Input Data.
	data_bits_count_required = data_columns_count * data_rows_count
	if data_bits_count != data_bits_count_required {
		return nil, ERROR_BAD_SIZE
	}
	if data_columns_count > data_columns_count_limit {
		return nil, ERROR_COLUMNS_ERROR
	}
	if data_rows_count > data_rows_count_limit {
		return nil, ERROR_ROWS_ERROR
	}

	// Indices & Sizes.
	result_columns_count = data_columns_count + DS
	result_rows_count = data_rows_count + DS

	data_first_column_index = 0
	data_first_row_index = 0
	//data_last_column_index = data_columns_count - 1
	//data_last_row_index = data_rows_count - 1

	result_first_column_index = 0
	result_first_row_index = 0
	result_last_column_index = result_columns_count - 1
	result_last_row_index = result_rows_count - 1

	// Create an empty Field.
	result = make(field_of_bits, result_rows_count)
	for i = result_first_row_index; i <= result_last_row_index; i++ {
		result[i] = make(row_of_bits, result_columns_count)
		for j = result_first_column_index; j <= result_last_column_index; j++ {
			result[i][j] = EMPTY
		}
	}

	// Draw the Frame I.
	for j = result_first_column_index; j <= result_last_column_index; j++ {
		result[result_first_row_index][j] = FILLED
		result[result_last_row_index][j] = FILLED
	}
	for i = result_first_row_index; i <= result_last_row_index; i++ {
		result[i][result_first_column_index] = FILLED
		result[i][result_last_column_index] = FILLED
	}
	// Draw Frame's Spacer.
	i_min = result_first_row_index + 1
	i_max = result_last_row_index - 1
	j_min = result_first_column_index + 1
	j_max = result_last_column_index - 1
	for j = j_min; j <= j_max; j++ {
		result[i_min][j] = EMPTY
		result[i_max][j] = EMPTY
	}
	for i = i_min; i <= i_max; i++ {
		result[i][j_min] = EMPTY
		result[i][j_max] = EMPTY
	}
	// Draw the Frame II.
	i_min = result_first_row_index + 2
	i_max = result_last_row_index - 2
	j_min = result_first_column_index + 2
	j_max = result_last_column_index - 2
	for j = j_min; j <= j_max; j++ {
		result[i_min][j] = FILLED
		result[i_max][j] = FILLED
	}
	for i = i_min; i <= i_max; i++ {
		result[i][j_min] = FILLED
		result[i][j_max] = FILLED
	}
	// Draw Frame's Spacer.
	i_min = result_first_row_index + 3
	i_max = result_last_row_index - 3
	j_min = result_first_column_index + 3
	j_max = result_last_column_index - 3
	for j = j_min; j <= j_max; j++ {
		result[i_min][j] = EMPTY
		result[i_max][j] = EMPTY
	}
	for i = i_min; i <= i_max; i++ {
		result[i][j_min] = EMPTY
		result[i][j_max] = EMPTY
	}

	// Draw Data.
	i_min = result_first_row_index + DO
	i_max = result_last_row_index - DO
	j_min = result_first_column_index + DO
	j_max = result_last_column_index - DO
	y = data_first_row_index
	for i = i_min; i <= i_max; i++ {
		x = data_first_column_index
		for j = j_min; j <= j_max; j++ {
			result[i][j] = data[y][x]
			x++
		}
		y++
	}

	return result, ERROR_ALL_CLEAR
}

//==============================================================================

// Checks Integrity of a Frame I of the Message.

func check_frame_f1(
	message_bits_count uint64,
	message_columns_count uint64,
	message_rows_count uint64,
	message field_of_bits) bool {

	const message_columns_count_limit = math.MaxUint64
	const message_rows_count_limit = math.MaxUint64

	const message_rows_count_min = 4 + 1    // Rows in empty Message.
	const message_columns_count_min = 4 + 1 // Columns in empty Message.

	const ERROR_ALL_CLEAR = true // No Error.
	const ERROR = false

	var data_bits_count_required uint64

	// Cursors in Message.
	var i uint64     // Current Row #.
	var i_max uint64 //
	var i_min uint64 //
	var j uint64     // Current Column #.
	var j_max uint64 //
	var j_min uint64 //

	// Check Input Data.
	data_bits_count_required = message_columns_count * message_rows_count
	if message_bits_count != data_bits_count_required {
		return ERROR
	}
	if message_columns_count > message_columns_count_limit {
		return ERROR
	}
	if message_rows_count > message_rows_count_limit {
		return ERROR
	}

	// Check Minimum Sizes.
	if message_rows_count < message_rows_count_min {
		return ERROR
	}
	if message_columns_count < message_columns_count_min {
		return ERROR
	}

	// Check Dimensions of Array.
	if uint64(len(message)) != message_rows_count {
		return ERROR
	}
	i_min = 0
	i_max = message_rows_count - 1
	for i = i_min; i <= i_max; i++ {
		if uint64(len(message[i])) != message_columns_count {
			return ERROR
		}
	}

	// Check Frame I.
	j_min = 0
	j_max = message_columns_count - 1
	for j = j_min; j <= j_max; j++ {
		if message[i_min][j] != FILLED {
			return ERROR
		}
		if message[i_max][j] != FILLED {
			return ERROR
		}
	}
	for i = i_min; i <= i_max; i++ {
		if message[i][j_min] != FILLED {
			return ERROR
		}
		if message[i][j_max] != FILLED {
			return ERROR
		}
	}
	// Check Frame's Spacer.
	i_min = 1
	i_max = message_rows_count - 2
	j_min = 1
	j_max = message_columns_count - 2
	for j = j_min; j <= j_max; j++ {
		if message[i_min][j] != EMPTY {
			return ERROR
		}
		if message[i_max][j] != EMPTY {
			return ERROR
		}
	}
	for i = i_min; i <= i_max; i++ {
		if message[i][j_min] != EMPTY {
			return ERROR
		}
		if message[i][j_max] != EMPTY {
			return ERROR
		}
	}

	return ERROR_ALL_CLEAR
}

//==============================================================================

// Checks Integrity of a Frame II of the Message.

func check_frame_f2(
	message_bits_count uint64,
	message_columns_count uint64,
	message_rows_count uint64,
	message field_of_bits) bool {

	const message_columns_count_limit = math.MaxUint64
	const message_rows_count_limit = math.MaxUint64

	const message_rows_count_min = 8 + 1    // Rows in empty Message.
	const message_columns_count_min = 8 + 1 // Columns in empty Message.

	const ERROR_ALL_CLEAR = true // No Error.
	const ERROR = false

	var data_bits_count_required uint64

	// Cursors in Message.
	var i uint64     // Current Row #.
	var i_max uint64 //
	var i_min uint64 //
	var j uint64     // Current Column #.
	var j_max uint64 //
	var j_min uint64 //

	// Check Input Data.
	data_bits_count_required = message_columns_count * message_rows_count
	if message_bits_count != data_bits_count_required {
		return ERROR
	}
	if message_columns_count > message_columns_count_limit {
		return ERROR
	}
	if message_rows_count > message_rows_count_limit {
		return ERROR
	}

	// Check Minimum Sizes.
	if message_rows_count < message_rows_count_min {
		return ERROR
	}
	if message_columns_count < message_columns_count_min {
		return ERROR
	}

	// Check Dimensions of Array.
	if uint64(len(message)) != message_rows_count {
		return ERROR
	}
	i_min = 0
	i_max = message_rows_count - 1
	for i = i_min; i <= i_max; i++ {
		if uint64(len(message[i])) != message_columns_count {
			return ERROR
		}
	}

	// Check Frame I.
	j_min = 0
	j_max = message_columns_count - 1
	for j = j_min; j <= j_max; j++ {
		if message[i_min][j] != FILLED {
			return ERROR
		}
		if message[i_max][j] != FILLED {
			return ERROR
		}
	}
	for i = i_min; i <= i_max; i++ {
		if message[i][j_min] != FILLED {
			return ERROR
		}
		if message[i][j_max] != FILLED {
			return ERROR
		}
	}
	// Check Frame's Spacer.
	i_min = 1
	i_max = message_rows_count - 2
	j_min = 1
	j_max = message_columns_count - 2
	for j = j_min; j <= j_max; j++ {
		if message[i_min][j] != EMPTY {
			return ERROR
		}
		if message[i_max][j] != EMPTY {
			return ERROR
		}
	}
	for i = i_min; i <= i_max; i++ {
		if message[i][j_min] != EMPTY {
			return ERROR
		}
		if message[i][j_max] != EMPTY {
			return ERROR
		}
	}

	// Check Frame II.
	i_min = 2
	i_max = message_rows_count - 3
	j_min = 2
	j_max = message_columns_count - 3
	for j = j_min; j <= j_max; j++ {
		if message[i_min][j] != FILLED {
			return ERROR
		}
		if message[i_max][j] != FILLED {
			return ERROR
		}
	}
	for i = i_min; i <= i_max; i++ {
		if message[i][j_min] != FILLED {
			return ERROR
		}
		if message[i][j_max] != FILLED {
			return ERROR
		}
	}
	// Check Frame's Spacer.
	i_min = 3
	i_max = message_rows_count - 4
	j_min = 3
	j_max = message_columns_count - 4
	for j = j_min; j <= j_max; j++ {
		if message[i_min][j] != EMPTY {
			return ERROR
		}
		if message[i_max][j] != EMPTY {
			return ERROR
		}
	}
	for i = i_min; i <= i_max; i++ {
		if message[i][j_min] != EMPTY {
			return ERROR
		}
		if message[i][j_max] != EMPTY {
			return ERROR
		}
	}

	return ERROR_ALL_CLEAR
}

//==============================================================================

// Gets Data from Message with Frame I.

func get_data_f1(
	message_bits_count uint64,
	message_columns_count uint64,
	message_rows_count uint64,
	message field_of_bits) (field_of_bits, bool) {

	const DS = 4
	const DO = DS / 2

	const ERROR_ALL_CLEAR = true // No Error.
	const ERROR = false

	var data field_of_bits
	var data_rows_count uint64
	var data_columns_count uint64
	var cf bool // Result of Frame Check.

	// Cursors in Message.
	var i uint64     // Current Row #.
	var i_min uint64 //
	var j uint64     // Current Column #.
	var j_min uint64 //

	// Cursors in Data.
	var y uint64 // Current Row #.
	var x uint64 // Current Column #.

	// Check Frame.
	cf = check_frame_f1(message_bits_count, message_columns_count,
		message_rows_count, message)
	if cf == ERROR {
		return nil, ERROR
	}

	// Prepare Data.
	data_rows_count = message_rows_count - DS
	data_columns_count = message_columns_count - DS
	//
	data = make(field_of_bits, data_rows_count)
	for y = 0; y < data_rows_count; y++ {
		data[y] = make(row_of_bits, data_columns_count)
		for x = 0; x < data_columns_count; x++ {
			data[y][x] = EMPTY
		}
	}

	// Get Data.
	i_min = DO
	j_min = DO
	i = i_min
	for y = 0; y < data_rows_count; y++ {
		j = j_min
		for x = 0; x < data_columns_count; x++ {
			data[y][x] = message[i][j]
			j++
		}
		i++
	}

	return data, ERROR_ALL_CLEAR
}

//==============================================================================

// Gets Data from Message with Frame II.

func get_data_f2(
	message_bits_count uint64,
	message_columns_count uint64,
	message_rows_count uint64,
	message field_of_bits) (field_of_bits, bool) {

	const DS = 8
	const DO = DS / 2

	const ERROR_ALL_CLEAR = true // No Error.
	const ERROR = false

	var data field_of_bits
	var data_rows_count uint64
	var data_columns_count uint64
	var cf bool // Result of Frame Check.

	// Cursors in Message.
	var i uint64     // Current Row #.
	var i_min uint64 //
	var j uint64     // Current Column #.
	var j_min uint64 //

	// Cursors in Data.
	var y uint64 // Current Row #.
	var x uint64 // Current Column #.

	// Check Frame.
	cf = check_frame_f2(message_bits_count, message_columns_count,
		message_rows_count, message)
	if cf == ERROR {
		return nil, ERROR
	}

	// Prepare Data.
	data_rows_count = message_rows_count - DS
	data_columns_count = message_columns_count - DS
	//
	data = make(field_of_bits, data_rows_count)
	for y = 0; y < data_rows_count; y++ {
		data[y] = make(row_of_bits, data_columns_count)
		for x = 0; x < data_columns_count; x++ {
			data[y][x] = EMPTY
		}
	}

	// Get Data.
	i_min = DO
	j_min = DO
	i = i_min
	for y = 0; y < data_rows_count; y++ {
		j = j_min
		for x = 0; x < data_columns_count; x++ {
			data[y][x] = message[i][j]
			j++
		}
		i++
	}

	return data, ERROR_ALL_CLEAR
}

//==============================================================================

// Converts Field into Array of Bytes.

func field_to_bytes(
	field_bits_count uint64,
	field_columns_count uint64,
	field_rows_count uint64,
	field field_of_bits) ([]byte, bool) {

	const field_columns_count_limit = math.MaxUint64
	const field_rows_count_limit = math.MaxUint64

	const ERROR_ALL_CLEAR = true // No Error.
	const ERROR = false

	const MSG_1 = "Warning ! The Size of the Output Data can not be stored " +
		"using 8-Bit Bytes ! The Size is not a Multiple of 8 !"

	var i uint64
	var j uint64

	// Cursors in Field.
	var y uint64
	var x uint64

	var array []byte
	var current_bit bit
	var current_byte byte
	var bytes_count uint64
	var field_bits_count_required uint64
	var field_column_first uint64
	var field_column_last uint64

	field_column_first = 0
	field_column_last = field_columns_count - 1

	// Check Input Data.
	field_bits_count_required = field_columns_count * field_rows_count
	if field_bits_count != field_bits_count_required {
		log.Println("1")
		return nil, ERROR
	}
	if field_columns_count > field_columns_count_limit {
		log.Println("2")
		return nil, ERROR
	}
	if field_rows_count > field_rows_count_limit {
		log.Println("3")
		return nil, ERROR
	}

	// Can be converted to Bytes ?
	if (field_bits_count % bits_in_byte) != 0 {
		fmt.Println(MSG_1)
		return nil, ERROR
	}

	bytes_count = field_bits_count / bits_in_byte
	array = make([]byte, bytes_count)

	x = 0
	y = 0
	for i = 0; i < bytes_count; i++ {

		current_byte = 0

		// Read 8 Bits.
		for j = 0; j < bits_in_byte; j++ {

			current_bit = field[y][x]

			// Save Bit in Byte.
			if current_bit == FILLED {
				current_byte = current_byte | (128 >> j)
			}

			// Next Element in Field.
			if x == field_column_last {
				y++
				x = field_column_first
			} else {
				x++
			}
		}

		// Save to Array.
		array[i] = current_byte
	}

	return array, ERROR_ALL_CLEAR
}

//==============================================================================

// Converts Array of Bytes into Field.

func bytes_to_field(
	field_bits_count uint64,
	field_columns_count uint64,
	field_rows_count uint64,
	array []byte) (field_of_bits, bool) {

	const field_columns_count_limit = math.MaxUint64
	const field_rows_count_limit = math.MaxUint64

	const ERROR_ALL_CLEAR = true // No Error.
	const ERROR = false

	var i uint64
	var j uint64

	// Cursors in Field.
	var y uint64
	var x uint64

	var field field_of_bits
	var current_bit bit
	var current_byte byte
	var current_byte_tmp byte
	var bytes_count uint64
	var field_bits_count_required uint64
	var field_column_first uint64
	var field_column_last uint64

	field_column_first = 0
	field_column_last = field_columns_count - 1

	// Check Input Data.
	field_bits_count_required = field_columns_count * field_rows_count
	if field_bits_count != field_bits_count_required {
		return nil, ERROR
	}
	if field_columns_count > field_columns_count_limit {
		return nil, ERROR
	}
	if field_rows_count > field_rows_count_limit {
		return nil, ERROR
	}

	// Can be converted to Bytes ?
	if (field_bits_count % bits_in_byte) != 0 {
		return nil, ERROR
	}
	bytes_count = uint64(len(array))
	if bytes_count*bits_in_byte != field_bits_count {
		return nil, ERROR
	}

	// Create an empty Field.
	field = make(field_of_bits, field_rows_count)
	for y = 0; y < field_rows_count; y++ {
		field[y] = make(row_of_bits, field_columns_count)
		for x = 0; x < field_columns_count; x++ {
			field[y][x] = EMPTY
		}
	}

	x = 0
	y = 0
	for i = 0; i < bytes_count; i++ {

		current_byte = array[i]

		// Read 8 Bits.
		for j = 0; j < bits_in_byte; j++ {

			current_byte_tmp = (current_byte >> (7 - j)) & 1
			if current_byte_tmp == 1 {
				current_bit = FILLED
			} else {
				current_bit = EMPTY
			}

			// Save Bit in Field.
			field[y][x] = current_bit

			// Next Element in Field.
			if x == field_column_last {
				y++
				x = field_column_first
			} else {
				x++
			}
		}
	}

	return field, ERROR_ALL_CLEAR
}

//==============================================================================

// Read Command Line Arguments (Keys, Flags, Switches).

func read_cla() {

	var err error

	// Set Rules.
	cla_file_in = flag.String("fi", "input", "File Input.")
	cla_file_out = flag.String("fo", "output", "File Output.")
	cla_action_type = flag.String("a", "", "Action Type.")
	cla_x = flag.String("x", "0", "Columns.")
	cla_y = flag.String("y", "0", "Rows.")

	// Read C.L.A.
	flag.Parse()

	// Files.
	file_input_path = *cla_file_in
	file_output_path = *cla_file_out

	// Action Type.
	if *cla_action_type == "e1" {
		action_type = ACTION_ENCODE_F1
	} else if *cla_action_type == "e2" {
		action_type = ACTION_ENCODE_F2
	} else if *cla_action_type == "d1" {
		action_type = ACTION_DECODE_F1
	} else if *cla_action_type == "d2" {
		action_type = ACTION_DECODE_F2
	} else {
		action_type = ACTION_NONE
	}

	// X, Y.
	file_input_x, err = strconv.ParseUint(*cla_x, 10, 64)
	check_error(err)
	file_input_y, err = strconv.ParseUint(*cla_y, 10, 64)
	check_error(err)
}

//==============================================================================

func check_error(err error) {

	if err != nil {
		log.Println(err)
		os.Exit(ERROR_1)
	}
}

//==============================================================================

func check_ok(ok bool) {

	if !ok {
		log.Println("Error.")
		os.Exit(ERROR_1)
	} else {
		//fmt.Println("OK.")
	}
}

//==============================================================================

func check_err_code(err_code uint8) {

	if err_code == 0 {
		//fmt.Println("OK.")
	} else {
		log.Println("Error.")
		os.Exit(ERROR_1)
	}
}

//==============================================================================
