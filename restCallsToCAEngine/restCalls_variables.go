package restCallsToCAEngine

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
)

// Todo struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	// All TestInstruction-data for 'Custody Arrangement'
	allTestInstructions_CA TestInstructions.AllTestInstructions_CA_TestCaseSetUpStruct

	// All Attributes-data for 'Custody Arrangement' as map
	testInstructionAttributesMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TypeAndStructs.TestInstructionAttributeStruct

	// All FangEngineData for 'Custody Arrangement' as map
	fangEngineClassesMethodsAttributesMap map[TypeAndStructs.OriginalElementUUIDType]*FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
)
