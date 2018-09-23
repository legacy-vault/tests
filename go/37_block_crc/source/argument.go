// argument.go.

// O.S. Command Line Arguments Handling.

// Date: 2018-09-22.

package main

import (
	"flag"
	"fmt"
)

// Help Hint Descriptions.
const CLA_Usage_Action = "Action"
const CLA_Usage_BlockSize = "Block Size"
const CLA_Usage_DataFile = "Data File Name"
const CLA_Usage_SumFile = "CheckSum File Name"
const CLA_Usage_VerboseMode = "Verbose Mode"

// Parameter Names.
const CLA_ParamName_Action = "action"
const CLA_ParamName_BlockSize = "block_size"
const CLA_ParamName_DataFile = "data_file"
const CLA_ParamName_SumFile = "sum_file"
const CLA_ParamName_VerboseMode = "verbose"

// Default Values.
const CLA_DefaultValue_Action = ActionName_Unknown
const CLA_DefaultValue_BlockSize = 4096
const CLA_DefaultValue_DataFile = "input.txt"
const CLA_DefaultValue_SumFile = ""
const CLA_DefaultValue_VerboseMode = false

const ReportFormatA = "<%s> = [%v]." + NL

var claAction *string
var claBlockSize *uint64
var claDataFile *string
var claSumFile *string
var claSumFile2 string
var claVerboseMode *bool

// Reads O.S. Command Line Arguments.
func claRead() error {

	var reportFormat string

	// Read Verbose Mode.
	claVerboseMode = flag.Bool(
		CLA_ParamName_VerboseMode,
		CLA_DefaultValue_VerboseMode,
		CLA_Usage_VerboseMode,
	)

	// Read Input File's Name.
	claDataFile = flag.String(
		CLA_ParamName_DataFile,
		CLA_DefaultValue_DataFile,
		CLA_Usage_DataFile,
	)

	// Read Action.
	claAction = flag.String(
		CLA_ParamName_Action,
		CLA_DefaultValue_Action,
		CLA_Usage_Action,
	)

	// Read Block Size.
	claBlockSize = flag.Uint64(
		CLA_ParamName_BlockSize,
		CLA_DefaultValue_BlockSize,
		CLA_Usage_BlockSize,
	)

	// Read Output File's Name.
	claSumFile = flag.String(
		CLA_ParamName_SumFile,
		CLA_DefaultValue_SumFile,
		CLA_Usage_SumFile,
	)

	// Parse Flags.
	flag.Parse()

	// Set Values according to Flag States...

	// 1. Verbose Mode.
	verboseMode = *claVerboseMode

	// 2. Output File Name.
	if *claSumFile == CLA_DefaultValue_SumFile {
		claSumFile2 = *claDataFile +
			FileExtensionSeparator + FileOutputExt
	} else {
		claSumFile2 = *claSumFile
	}

	// 3. Action.
	parseAction(*claAction)

	// Report.
	if verboseMode {

		reportFormat = ReportFormatA

		// Input File.
		fmt.Printf(
			reportFormat,
			CLA_Usage_DataFile,
			*claDataFile,
		)

		// Block Size.
		fmt.Printf(
			reportFormat,
			CLA_Usage_BlockSize,
			*claBlockSize,
		)

		// Output File.
		fmt.Printf(
			reportFormat,
			CLA_Usage_SumFile,
			claSumFile2,
		)
	}

	return nil
}
