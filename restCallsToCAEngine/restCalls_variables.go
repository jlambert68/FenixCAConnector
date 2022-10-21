package restCallsToCAEngine

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
)

// FangEngineRestApiMessageStruct
// Used when converting a TestInstruction into data to be used in RestCall to FangEngine
type FangEngineRestApiMessageStruct struct {
	FangEngineClassNameNAME           FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_Type
	FangEngineMethodNameNAME          FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_Type
	FangEngineExpectedToBePassedValue TypeAndStructs.AttributeValueAsStringType
	TestInstructionAttribute          []TestInstructionAttributesUuidAndValueStruct
	FangAttributes                    map[TypeAndStructs.TestInstructionAttributeUUIDType]*FangEngineClassesAndMethods.FangEngineAttributesStruct

	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TestInstructionAttributesUuidAndValueStruct
// Holds one Attribute UUID, Name and Value
type TestInstructionAttributesUuidAndValueStruct struct {
	TestInstructionAttributeUUID          TypeAndStructs.TestInstructionAttributeUUIDType
	TestInstructionAttributeName          TypeAndStructs.TestInstructionAttributeNameType
	TestInstructionAttributeValueAsString TypeAndStructs.AttributeValueAsStringType
}

var (
	// All TestInstruction-data for 'Custody Arrangement'
	allTestInstructions_CA TestInstructions.AllTestInstructions_CA_TestCaseSetUpStruct

	// All Attributes-data for 'Custody Arrangement' as map
	testInstructionAttributesMap map[TypeAndStructs.TestInstructionAttributeUUIDType]*TypeAndStructs.TestInstructionAttributeStruct

	// All FangEngineData for 'Custody Arrangement' as map
	fangEngineClassesMethodsAttributesMap map[TypeAndStructs.OriginalElementUUIDType]*FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
)
