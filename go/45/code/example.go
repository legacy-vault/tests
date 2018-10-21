// ...

package main

import (
	"fmt"
	"github.com/legacy-vault/library/go/fixed_size_registry"
	"time"
)

const FormatShowRecord = "Last Record of the %d Records contains '%v'" +
	" and was created on %s.\r\n"
const FormatShowNoRecord = "Last Record does not exist.\r\n"
const FormatShowRegistry = "Registry: [%s] (%d Records). "

func main() {

	var capacity uint64
	var registry *fsregistry.Registry
	var size uint64

	capacity = 5

	// Create a Registry.
	registry = fsregistry.New(capacity)
	size = registry.GetSize()
	fmt.Println(size)
	ShowLastRecordAndRegistry(registry)

	// Add a single Record.
	registry.AddARecord(&fsregistry.Record{Data: "A"})
	ShowLastRecordAndRegistry(registry)

	// Add a few Records.
	time.Sleep(time.Second * 2)
	registry.AddARecord(&fsregistry.Record{Data: "B"})
	time.Sleep(time.Second * 2)
	registry.AddARecord(&fsregistry.Record{Data: "C"})
	ShowLastRecordAndRegistry(registry)

	// Add more Records (more than Capacity).
	time.Sleep(time.Second * 1)
	registry.AddARecord(&fsregistry.Record{Data: "D"})
	time.Sleep(time.Second * 1)
	registry.AddARecord(&fsregistry.Record{Data: "E"})
	time.Sleep(time.Second * 1)
	registry.AddARecord(&fsregistry.Record{Data: "F"})
	time.Sleep(time.Second * 1)
	registry.AddARecord(&fsregistry.Record{Data: "G"})
	ShowLastRecordAndRegistry(registry)

	// Reset (clear) the Registry.
	fsregistry.Clear(registry)
	ShowLastRecordAndRegistry(registry)

	// Add a single Record.
	registry.AddARecord(&fsregistry.Record{Data: "H"})
	ShowLastRecordAndRegistry(registry)
}

func RecordsAsString(records []*fsregistry.Record) string {

	var itemStr string
	var record *fsregistry.Record
	var result string

	// Convert empty Interfaces to Strings.
	for _, record = range records {
		itemStr = fmt.Sprintf("(%v)", record.Data)
		result = result + itemStr
	}

	return result
}

// Shows the last Record and the whole Registry.
func ShowLastRecordAndRegistry(registry *fsregistry.Registry) {

	var record *fsregistry.Record
	var records []*fsregistry.Record
	var recordsText string
	var size uint64
	var toc int64
	var tocStr string

	// Show all Records.
	records, size = registry.GetStoredRecords()
	recordsText = RecordsAsString(records)
	fmt.Printf(FormatShowRegistry, recordsText, size)

	// Show the last Record.
	record, size = registry.GetLastRecord()
	if record != nil {

		// Record exists.
		toc = record.GetTimeOfCreation()
		tocStr = UnixTimestampToText(toc)
		fmt.Printf(FormatShowRecord, size, record.Data, tocStr)

	} else {

		// Record does not exist.
		fmt.Printf(FormatShowNoRecord)
	}
}

// Converts UNIX TimeStamp to Human-readable Text.
func UnixTimestampToText(unixTS int64) string {

	var tocText string
	var toc time.Time

	toc = time.Unix(unixTS, 0)
	tocText = toc.Format(time.RFC3339)

	return tocText
}
