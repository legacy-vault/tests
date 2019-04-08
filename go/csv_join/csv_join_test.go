// csv_join_test.go.

package csv_join

import (
	"os"
	"path"
	"testing"
)

func Test_readNames(t *testing.T) {

	const TestFolder = "test"
	const TestNamesFile = "names.csv"

	var cj CsvJoin
	var err error
	var name string
	var namesFile *os.File
	var tmpFileNames *os.File

	// Create the Test Environment.
	err = os.Mkdir(TestFolder, 0755)
	mustBeNoError(t, err)
	namesFile, err = os.Create(
		path.Join(TestFolder, TestNamesFile),
	)
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("ID,Name\r\n")
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("101,John\r\n")
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("103,Джон\r\n")
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("102,Jack\r\n")
	mustBeNoError(t, err)
	err = namesFile.Close()
	mustBeNoError(t, err)

	// Prepare the Object.
	cj = New(
		path.Join(TestFolder, TestNamesFile),
		"none",
		"none",
		true,
	)

	// Run the tested Function.
	err = cj.readNames()
	mustBeNoError(t, err)

	// Open temporary File.
	tmpFileNames, err = os.Open(cj.tmpFileWithName)
	mustBeNoError(t, err)

	// Read some Values.
	name, err = cj.getNameRecordText(tmpFileNames, 102)
	mustBeNoError(t, err)
	if name != "Jack" {
		t.Errorf("102 -> Name")
		t.FailNow()
	}
	name, err = cj.getNameRecordText(tmpFileNames, 103)
	mustBeNoError(t, err)
	if name != "Джон" {
		t.Errorf("103 -> Name")
		t.FailNow()
	}

	// Close temporary File.
	err = tmpFileNames.Close()
	mustBeNoError(t, err)

	// Clean the File System.
	err = os.RemoveAll(TestFolder)
	mustBeNoError(t, err)
}

func Test_readPhones(t *testing.T) {

	const TestFolder = "test"
	const TestPhonesFile = "phones.csv"

	var cj CsvJoin
	var err error
	var phone string
	var phonesFile *os.File
	var tmpFilePhones *os.File

	// Create the Test Environment.
	err = os.Mkdir(TestFolder, 0755)
	mustBeNoError(t, err)
	phonesFile, err = os.Create(
		path.Join(TestFolder, TestPhonesFile),
	)
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("ID,Phone\r\n")
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("101,+7 (901) 101-0000\r\n")
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("103,8-903-333-00-33\r\n")
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("102,8 /902/ 102 0202\r\n")
	mustBeNoError(t, err)
	err = phonesFile.Close()
	mustBeNoError(t, err)

	// Prepare the Object.
	cj = New(
		"none",
		path.Join(TestFolder, TestPhonesFile),
		"none",
		true,
	)

	// Run the tested Function.
	err = cj.readPhones()
	mustBeNoError(t, err)

	// Open temporary File.
	tmpFilePhones, err = os.Open(cj.tmpFileWithPhone)
	mustBeNoError(t, err)

	// Read some Values.
	phone, err = cj.getPhoneRecordText(tmpFilePhones, 102)
	mustBeNoError(t, err)
	if phone != "8 /902/ 102 0202" {
		t.Errorf("102 -> Phone")
		t.FailNow()
	}
	phone, err = cj.getPhoneRecordText(tmpFilePhones, 103)
	mustBeNoError(t, err)
	if phone != "8-903-333-00-33" {
		t.Errorf("103 -> Phone")
		t.FailNow()
	}

	// Close temporary File.
	err = tmpFilePhones.Close()
	mustBeNoError(t, err)

	// Clean the File System.
	err = os.RemoveAll(TestFolder)
	mustBeNoError(t, err)
}

// Manual and partial Test. //TODO.
func Test_Process(t *testing.T) {

	const TestFolder = "test"
	const TestNamesFile = "names.csv"
	const TestPhonesFile = "phones.csv"
	const TestNamesPhonesFile = "names-phones.csv"

	var cj CsvJoin
	var err error
	var namesFile *os.File
	var phonesFile *os.File

	// Create the Test Environment.
	err = os.Mkdir(TestFolder, 0755)
	mustBeNoError(t, err)
	namesFile, err = os.Create(
		path.Join(TestFolder, TestNamesFile),
	)
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("ID,Name\r\n")
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("101,John\r\n")
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("103,Джон\r\n")
	mustBeNoError(t, err)
	_, err = namesFile.WriteString("105,Alice\r\n")
	mustBeNoError(t, err)
	err = namesFile.Close()
	mustBeNoError(t, err)
	phonesFile, err = os.Create(
		path.Join(TestFolder, TestPhonesFile),
	)
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("ID,Phone\r\n")
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("102,8 /902/ 102 0202\r\n")
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("103,8-903-333-00-33\r\n")
	mustBeNoError(t, err)
	_, err = phonesFile.WriteString("104,4444\r\n")
	mustBeNoError(t, err)
	err = phonesFile.Close()
	mustBeNoError(t, err)

	// Run the tested Function.
	cj = New(
		path.Join(TestFolder, TestNamesFile),
		path.Join(TestFolder, TestPhonesFile),
		path.Join(TestFolder, TestNamesPhonesFile),
		true,
	)
	err = cj.Process()
	mustBeNoError(t, err)

	// Clean the File System.
	err = os.RemoveAll(TestFolder)
	mustBeNoError(t, err)
}

func mustBeNoError(
	t *testing.T,
	e error,
) {

	if e != nil {
		t.Error(e)
		t.FailNow()
	}
}
