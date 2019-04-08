// file.go.

package csv_join

import (
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

const TextEmpty = ""

// Temporary File is used to store Data with manageable Offset.
// Standard Golang's CSV Reader does not return an Offset of a Record in a File.
// Each Record in the temporary File consists of:
//	1.	uint (8 Bytes) containing the ID;
//	2.	uint32 (4 Bytes) containing the Length of Text;
//	3.	string	(Length is set in a previous Field) with Text.

func (this *CsvJoin) readNames() error {

	var bytesWritten int
	var csvReader *csv.Reader
	var duplicate bool
	var err error
	var fileWithName *os.File
	var idOffset IdOffset
	var line []string
	var tmpFile *os.File
	var tmpFileDataLength uint
	var tmpFileExists bool
	var tmpTextBA []byte
	var tmpTextLen uint32
	var tmpUint64 uint64

	// Preparations.
	this.tmpFileWithName = this.fileWithName + TmpFileNamePostfix
	this.nameRecordsInfo = make(map[uint]uint)

	// Check Existence of temporary File.
	tmpFileExists, err = FileExists(this.tmpFileWithName)
	if err != nil {
		return err
	}
	if tmpFileExists {
		err = fmt.Errorf(
			"Temporary File '%v' already exists",
			this.tmpFileWithName,
		)
		return err
	}

	// Open the input File.
	fileWithName, err = os.Open(this.fileWithName)
	if err != nil {
		return err
	}
	defer func() {
		var deferredErr error
		deferredErr = fileWithName.Close()
		if deferredErr != nil {
			log.Println(deferredErr)
		}
	}()
	csvReader = csv.NewReader(fileWithName)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = 2

	// Create the temporary File.
	tmpFile, err = os.Create(this.tmpFileWithName)
	if err != nil {
		return err
	}
	defer func() {
		var deferredErr error
		deferredErr = tmpFile.Close()
		if deferredErr != nil {
			log.Println(deferredErr)
		}
	}()

	// Skip the first Line if needed.
	if this.firstLineIsSkipped {
		_, err = csvReader.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
	}

	// Process each Line.
	tmpFileDataLength = 0
	for {
		// Read the Line and parse it.
		line, err = csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		tmpUint64, err = strconv.ParseUint(line[0], 10, 64)
		if err != nil {
			return err
		}
		idOffset.Id = uint(tmpUint64)
		idOffset.Offset = tmpFileDataLength

		// Save Information about a Record.
		_, duplicate = this.nameRecordsInfo[idOffset.Id]
		if duplicate {
			err = fmt.Errorf(
				"Duplicate ID '%v' in Record",
				idOffset.Id,
			)
			return err
		}
		this.nameRecordsInfo[idOffset.Id] = idOffset.Offset

		// Write Data to the temporary File...

		// 1. ID (8 Bytes).
		err = binary.Write(tmpFile, binary.BigEndian, uint64(idOffset.Id))
		if err != nil {
			return err
		}

		// 2. Text Size (4 Bytes).
		tmpTextBA = []byte(line[1])
		if len(tmpTextBA) > math.MaxUint32 {
			err = fmt.Errorf(
				"Text Field is too long (%v Bytes) to store its Size in uint32",
				len(tmpTextBA),
			)
			return err
		}
		tmpTextLen = uint32(len(tmpTextBA))
		err = binary.Write(tmpFile, binary.BigEndian, tmpTextLen)
		if err != nil {
			return err
		}

		// 3. Text (Variable Size).
		bytesWritten, err = tmpFile.Write(tmpTextBA)
		if err != nil {
			return err
		}
		if bytesWritten != int(tmpTextLen) {
			err = fmt.Errorf(
				"Temporary File Write Error: Size Mismatch (%v vs %v)",
				bytesWritten,
				tmpTextLen,
			)
			return err
		}

		// Remember the Data Size of the Temporary File.
		tmpFileDataLength += (8 + 4 + uint(tmpTextLen))
	}

	// Convert the Map into an Array (Slice) and sort it.
	this.nameRecordsInfoSorted = make([]IdOffset, 0, len(this.nameRecordsInfo))
	for id, offset := range this.nameRecordsInfo {
		this.nameRecordsInfoSorted = append(
			this.nameRecordsInfoSorted,
			IdOffset{Id: id, Offset: offset},
		)
	}
	lessFunction := func(i, j int) bool {
		return this.nameRecordsInfoSorted[i].Id < this.nameRecordsInfoSorted[j].Id
	}
	sort.Slice(this.nameRecordsInfoSorted, lessFunction)

	return nil
}

func (this *CsvJoin) readPhones() error {

	var bytesWritten int
	var csvReader *csv.Reader
	var duplicate bool
	var err error
	var fileWithPhone *os.File
	var idOffset IdOffset
	var line []string
	var tmpFile *os.File
	var tmpFileDataLength uint
	var tmpFileExists bool
	var tmpTextBA []byte
	var tmpTextLen uint32
	var tmpUint64 uint64

	// Preparations.
	this.tmpFileWithPhone = this.fileWithPhone + TmpFileNamePostfix
	this.phoneRecordsInfo = make(map[uint]uint)

	// Check Existence of temporary File.
	tmpFileExists, err = FileExists(this.tmpFileWithPhone)
	if err != nil {
		return err
	}
	if tmpFileExists {
		err = fmt.Errorf(
			"Temporary File '%v' already exists",
			this.tmpFileWithPhone,
		)
		return err
	}

	// Open the input File.
	fileWithPhone, err = os.Open(this.fileWithPhone)
	if err != nil {
		return err
	}
	defer func() {
		var deferredErr error
		deferredErr = fileWithPhone.Close()
		if deferredErr != nil {
			log.Println(deferredErr)
		}
	}()
	csvReader = csv.NewReader(fileWithPhone)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = 2

	// Create the temporary File.
	tmpFile, err = os.Create(this.tmpFileWithPhone)
	if err != nil {
		return err
	}
	defer func() {
		var deferredErr error
		deferredErr = tmpFile.Close()
		if deferredErr != nil {
			log.Println(deferredErr)
		}
	}()

	// Skip the first Line if needed.
	if this.firstLineIsSkipped {
		_, err = csvReader.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
	}

	// Process each Line.
	tmpFileDataLength = 0
	for {
		// Read the Line and parse it.
		line, err = csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		tmpUint64, err = strconv.ParseUint(line[0], 10, 64)
		if err != nil {
			return err
		}
		idOffset.Id = uint(tmpUint64)
		idOffset.Offset = tmpFileDataLength

		// Save Information about a Record.
		_, duplicate = this.phoneRecordsInfo[idOffset.Id]
		if duplicate {
			err = fmt.Errorf(
				"Duplicate ID '%v' in Record",
				idOffset.Id,
			)
			return err
		}
		this.phoneRecordsInfo[idOffset.Id] = idOffset.Offset

		// Write Data to the temporary File...

		// 1. ID (8 Bytes).
		err = binary.Write(tmpFile, binary.BigEndian, uint64(idOffset.Id))
		if err != nil {
			return err
		}

		// 2. Text Size (4 Bytes).
		tmpTextBA = []byte(line[1])
		if len(tmpTextBA) > math.MaxUint32 {
			err = fmt.Errorf(
				"Text Field is too long (%v Bytes) to store its Size in uint32",
				len(tmpTextBA),
			)
			return err
		}
		tmpTextLen = uint32(len(tmpTextBA))
		err = binary.Write(tmpFile, binary.BigEndian, tmpTextLen)
		if err != nil {
			return err
		}

		// 3. Text (Variable Size).
		bytesWritten, err = tmpFile.Write(tmpTextBA)
		if err != nil {
			return err
		}
		if bytesWritten != int(tmpTextLen) {
			err = fmt.Errorf(
				"Temporary File Write Error: Size Mismatch (%v vs %v)",
				bytesWritten,
				tmpTextLen,
			)
			return err
		}

		// Remember the Data Size of the Temporary File.
		tmpFileDataLength += (8 + 4 + uint(tmpTextLen))
	}

	// Convert the Map into an Array (Slice) and sort it.
	this.phoneRecordsInfoSorted = make([]IdOffset, 0, len(this.phoneRecordsInfo))
	for id, offset := range this.phoneRecordsInfo {
		this.phoneRecordsInfoSorted = append(
			this.phoneRecordsInfoSorted,
			IdOffset{Id: id, Offset: offset},
		)
	}
	lessFunction := func(i, j int) bool {
		return this.phoneRecordsInfoSorted[i].Id < this.phoneRecordsInfoSorted[j].Id
	}
	sort.Slice(this.phoneRecordsInfoSorted, lessFunction)

	return nil
}

// Returns a Name Text if it is set for an ID.
// Returns an empty String if a Name is not set for an ID.
func (this CsvJoin) getNameRecordText(
	tmpFile *os.File,
	id uint,
) (string, error) {

	var err error
	var exists bool
	var idUint64 uint64
	var name string
	var nameBA []byte
	var nameLenUint32 uint32
	var offset uint

	// Get an Offset.
	offset, exists = this.nameRecordsInfo[id]
	if !exists {
		return TextEmpty, nil
	}
	if offset > math.MaxInt64 {
		err = fmt.Errorf(
			"Offset %v is too big to be stored in 'int64' Variable",
			offset,
		)
		return name, err
	}

	// Find Data in the temporary File...

	// 1. Set the File Cursor to the Start of the Record.
	_, err = tmpFile.Seek(int64(offset), io.SeekStart)
	if err != nil {
		return name, err
	}

	// 2. Read the ID.
	err = binary.Read(tmpFile, binary.BigEndian, &idUint64)
	if err != nil {
		return name, err
	}
	if uint(idUint64) != id {
		err = fmt.Errorf(
			"ID Mismatch (%v vs %v)",
			idUint64,
			id,
		)
		return name, err
	}

	// 3. Read the Length of the Name.
	err = binary.Read(tmpFile, binary.BigEndian, &nameLenUint32)
	if err != nil {
		return name, err
	}

	// 4. Read the Name String.
	nameBA = make([]byte, nameLenUint32)
	err = binary.Read(tmpFile, binary.BigEndian, &nameBA)
	if err != nil {
		return name, err
	}
	name = string(nameBA)

	return name, nil
}

// Returns a Phone Text if it is set for an ID.
// Returns an empty String if a Phone is not set for an ID.
func (this CsvJoin) getPhoneRecordText(
	tmpFile *os.File,
	id uint,
) (string, error) {

	var err error
	var exists bool
	var idUint64 uint64
	var offset uint
	var phone string
	var phoneBA []byte
	var phoneLenUint32 uint32

	// Get an Offset.
	offset, exists = this.phoneRecordsInfo[id]
	if !exists {
		return TextEmpty, nil
	}
	if offset > math.MaxInt64 {
		err = fmt.Errorf(
			"Offset %v is too big to be stored in 'int64' Variable",
			offset,
		)
		return phone, err
	}

	// Find Data in the temporary File...

	// 1. Set the File Cursor to the Start of the Record.
	_, err = tmpFile.Seek(int64(offset), io.SeekStart)
	if err != nil {
		return phone, err
	}

	// 2. Read the ID.
	err = binary.Read(tmpFile, binary.BigEndian, &idUint64)
	if err != nil {
		return phone, err
	}
	if uint(idUint64) != id {
		err = fmt.Errorf(
			"ID Mismatch (%v vs %v)",
			idUint64,
			id,
		)
		return phone, err
	}

	// 3. Read the Length of the Phone.
	err = binary.Read(tmpFile, binary.BigEndian, &phoneLenUint32)
	if err != nil {
		return phone, err
	}

	// 4. Read the Phone String.
	phoneBA = make([]byte, phoneLenUint32)
	err = binary.Read(tmpFile, binary.BigEndian, &phoneBA)
	if err != nil {
		return phone, err
	}
	phone = string(phoneBA)

	return phone, nil
}

// Saves ID, Name and Phone Information into the File.
func (this CsvJoin) saveIdNamePhoneRecord(
	csvWriter *csv.Writer,
	id uint,
	name string,
	phone string,
) error {

	var err error
	var text []string

	// Encode Data using the CSV Format.
	text = []string{
		strconv.FormatUint(uint64(id), 10),
		name,
		phone,
	}

	// Write to File.
	err = csvWriter.Write(text)
	if err != nil {
		return err
	}

	return nil
}
