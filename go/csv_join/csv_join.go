// csv_join.go.

package csv_join

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type CsvJoin struct {

	// Settings.
	fileWithName       string
	fileWithPhone      string
	fileWithNamePhone  string
	firstLineIsSkipped bool

	// Temporary Data...

	// 1. File Paths.
	tmpFileWithName  string
	tmpFileWithPhone string

	// 2. Maps.
	// Record Information.
	// Key is ID,
	// Value is an Offset in the Temporary File.
	nameRecordsInfo        map[uint]uint
	nameRecordsInfoSorted  []IdOffset
	phoneRecordsInfo       map[uint]uint
	phoneRecordsInfoSorted []IdOffset
}

func New(
	fileWithName string,
	fileWithPhone string,
	fileWithNamePhone string,
	firstLineIsSkipped bool,
) CsvJoin {

	var cj CsvJoin

	cj.fileWithName = fileWithName
	cj.fileWithPhone = fileWithPhone
	cj.fileWithNamePhone = fileWithNamePhone
	cj.firstLineIsSkipped = firstLineIsSkipped

	return cj
}

func (this *CsvJoin) Process() error {

	var err error

	err = this.readNames()
	if err != nil {
		return err
	}

	err = this.readPhones()
	if err != nil {
		return err
	}

	err = this.combineNamePhone()
	if err != nil {
		return err
	}

	return nil
}

func (this *CsvJoin) combineNamePhone() error {

	var csvWriter *csv.Writer
	var curIdxOfNameId uint
	var curIdxOfPhoneId uint
	var curNameId uint
	var curPhoneId uint
	var endIdxOfNameId uint  // Last Index + 1, "End Marker".
	var endIdxOfPhoneId uint // Last Index + 1, "End Marker".
	var err error
	var name string
	var outputFile *os.File
	var phone string
	var tmpFileNames *os.File
	var tmpFilePhones *os.File

	// Input Data Check.
	if len(this.nameRecordsInfoSorted) == 0 {
		err = fmt.Errorf(
			"Name List is empty",
		)
		return err
	}
	if len(this.phoneRecordsInfoSorted) == 0 {
		err = fmt.Errorf(
			"Phone List is empty",
		)
	}

	// Initial Cursors.
	curIdxOfNameId = 0
	endIdxOfNameId = uint(len(this.nameRecordsInfoSorted))
	curIdxOfPhoneId = 0
	endIdxOfPhoneId = uint(len(this.phoneRecordsInfoSorted))

	// Open Files...

	// 1. Names.
	tmpFileNames, err = os.Open(this.tmpFileWithName)
	if err != nil {
		return err
	}
	defer func() {
		var deferredErr error
		deferredErr = tmpFileNames.Close()
		if deferredErr != nil {
			log.Println(deferredErr)
		}
	}()

	// 2. Phones.
	tmpFilePhones, err = os.Open(this.tmpFileWithPhone)
	if err != nil {
		return err
	}
	defer func() {
		var deferredErr error
		deferredErr = tmpFilePhones.Close()
		if deferredErr != nil {
			log.Println(deferredErr)
		}
	}()

	// 3. Output File.
	outputFile, err = os.OpenFile(
		this.fileWithNamePhone,
		os.O_APPEND|os.O_CREATE,
		0644,
	)
	if err != nil {
		return err
	}
	defer func() {
		var deferredErr error
		deferredErr = outputFile.Close()
		if deferredErr != nil {
			log.Println(deferredErr)
		}
	}()
	csvWriter = csv.NewWriter(outputFile)

	// Combine Records.
	for true {

		// Select current Source Table.
		if curIdxOfNameId >= endIdxOfNameId {
			// Names Table is done.
			if curIdxOfPhoneId >= endIdxOfPhoneId {
				// Phones Table is also done.
				// Use no Table.
				break

			} else {
				// Phones Table is not done.
				// Use Phone Table.
				curPhoneId = this.phoneRecordsInfoSorted[curIdxOfPhoneId].Id
				name, err = this.getNameRecordText(tmpFileNames, curPhoneId)
				if err != nil {
					return err
				}
				phone, err = this.getPhoneRecordText(tmpFilePhones, curPhoneId)
				if err != nil {
					return err
				}

				// Move Cursor.
				curIdxOfPhoneId++

				// Save Data.
				err = this.saveIdNamePhoneRecord(csvWriter, curPhoneId, name, phone)
				if err != nil {
					return err
				}
			}

		} else {

			// Names Table is not done.
			if curIdxOfPhoneId >= endIdxOfPhoneId {
				// Phones Table is done.
				// Use Name Table.
				curNameId = this.nameRecordsInfoSorted[curIdxOfNameId].Id
				name, err = this.getNameRecordText(tmpFileNames, curNameId)
				if err != nil {
					return err
				}
				phone, err = this.getPhoneRecordText(tmpFilePhones, curNameId)
				if err != nil {
					return err
				}

				// Move Cursor.
				curIdxOfNameId++

				// Save Data.
				err = this.saveIdNamePhoneRecord(csvWriter, curNameId, name, phone)
				if err != nil {
					return err
				}

			} else {
				// Phones Table is also not done.
				// Select the least ID.
				curNameId = this.nameRecordsInfoSorted[curIdxOfNameId].Id
				curPhoneId = this.phoneRecordsInfoSorted[curIdxOfPhoneId].Id
				if curNameId < curPhoneId {
					// Use Name Table.
					name, err = this.getNameRecordText(tmpFileNames, curNameId)
					if err != nil {
						return err
					}
					phone, err = this.getPhoneRecordText(tmpFilePhones, curNameId)
					if err != nil {
						return err
					}

					// Move Cursor.
					curIdxOfNameId++

					// Save Data.
					err = this.saveIdNamePhoneRecord(csvWriter, curNameId, name, phone)
					if err != nil {
						return err
					}

				} else if curPhoneId < curNameId {
					// Use Phone Table.
					name, err = this.getNameRecordText(tmpFileNames, curPhoneId)
					if err != nil {
						return err
					}
					phone, err = this.getPhoneRecordText(tmpFilePhones, curPhoneId)
					if err != nil {
						return err
					}

					// Move Cursor.
					curIdxOfPhoneId++

					// Save Data.
					err = this.saveIdNamePhoneRecord(csvWriter, curPhoneId, name, phone)
					if err != nil {
						return err
					}

				} else {
					// Use both Tables.
					name, err = this.getNameRecordText(tmpFileNames, curNameId)
					if err != nil {
						return err
					}
					phone, err = this.getPhoneRecordText(tmpFilePhones, curPhoneId)
					if err != nil {
						return err
					}

					// Move Cursors.
					curIdxOfNameId++
					curIdxOfPhoneId++

					// Save Data.
					err = this.saveIdNamePhoneRecord(csvWriter, curNameId, name, phone)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	csvWriter.Flush()

	return nil
}
