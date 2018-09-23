// stage.go.

// Application Stages.

// Date: 2018-09-22.

package main

// Known Stages List.
const StagesCount = 5
const Stage_Index_Min = 1
const Stage_Index_Max = 5

// Stage Name Indices.
const Stage_Unknown = 0
const Stage_Initialization = 1
const Stage_CommandLineArguments = 2
const Stage_Finalization = 3
const Stage_BlockCRCCreation = 4
const Stage_BlockCRCCheck = 5

// Stage Names.
const Stage_Name_Unknown = "Unknown"                             // 0.
const Stage_Name_Initialization = "Program Initialization"       // 1.
const Stage_Name_CommandLineArguments = "Command Line Arguments" // 2.
const Stage_Name_Finalization = "Program Finalization"           // 3.
const Stage_Name_BlockCRCCreation = "Block CRC Creation"         // 4.
const Stage_Name_BlockCRCCheck = "Block CRC Check"               // 5.

// Known Statuses List.
const StatusesCount = 3
const Status_Index_Min = 1
const Status_Index_Max = 3

// Status Indices.
const Status_Unknown = 0
const Status_Started = 1
const Status_Finished = 2
const Status_Failed = 3

// Status Names.
//const Status_Name_Unknown = "has unknown Status" // 0.
//const Status_Name_Started = "has started"        // 1.
//const Status_Name_Finished = "has finished"      // 2.
//const Status_Name_Failed = "has failed"          // 3.
const Status_Name_Unknown = "???"    // 0.
const Status_Name_Started = "..."    // 1.
const Status_Name_Finished = "OK"    // 2.
const Status_Name_Failed = "Failure" // 3.

// Delimiters.
const Stage_Delimiter_NamePrefix = "["
const Stage_Delimiter_NamePostfix = "]"
const Stage_Delimiter_StatusPrefix = " "
const Stage_Delimiter_StatusPostfix = "."

type StageFunction func() error

var stageFucntions []StageFunction
var stageNames []string
var stageStatusNames []string

// Initializes Stages Names.
func stagesInit() error {

	// Cache Stage Names.
	stageNames = make([]string, StagesCount+1)

	stageNames[Stage_Unknown] = Stage_Name_Unknown
	stageNames[Stage_Initialization] = Stage_Name_Initialization
	stageNames[Stage_CommandLineArguments] = Stage_Name_CommandLineArguments
	stageNames[Stage_Finalization] = Stage_Name_Finalization
	stageNames[Stage_BlockCRCCreation] = Stage_Name_BlockCRCCreation
	stageNames[Stage_BlockCRCCheck] = Stage_Name_BlockCRCCheck

	// Cache Stage Status Names.
	stageStatusNames = make([]string, StatusesCount+1)

	stageStatusNames[Status_Unknown] = Status_Name_Unknown
	stageStatusNames[Status_Started] = Status_Name_Started
	stageStatusNames[Status_Finished] = Status_Name_Finished
	stageStatusNames[Status_Failed] = Status_Name_Failed

	// Cache Stage Functions.
	stageFucntions = make([]StageFunction, StagesCount+1)

	stageFucntions[Stage_Unknown] = void
	stageFucntions[Stage_Initialization] = Init
	stageFucntions[Stage_CommandLineArguments] = claRead
	stageFucntions[Stage_Finalization] = Fin
	stageFucntions[Stage_BlockCRCCreation] = bcrcCreate
	stageFucntions[Stage_BlockCRCCheck] = bcrcCheck

	return nil
}

// Returns the Text containing the Stage Name and its Status.
func stageStatusText(stage int, status int) string {

	var s string

	// Check Stage Index.
	if (stage < Stage_Index_Min) || (stage > Stage_Index_Max) {
		stage = Stage_Unknown
	}

	// Check Stage Status.
	if (status < Status_Index_Min) || (status > Status_Index_Max) {
		status = Status_Unknown
	}

	// Create a String.
	s = Stage_Delimiter_NamePrefix +
		stageNames[stage] +
		Stage_Delimiter_NamePostfix +
		Stage_Delimiter_StatusPrefix +
		stageStatusNames[status] +
		Stage_Delimiter_StatusPostfix

	return s
}
