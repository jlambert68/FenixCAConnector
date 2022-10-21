package restCallsToCAEngine

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
)

// InitiateRestCallsToCAEngine
// Do all initiation to have restEngine be able to do RestCalls to Custody Arrangements FangEngine
func InitiateRestCallsToCAEngine() {

	// Load all TestInstruction-data for 'Custody Arrangement'
	allTestInstructions_CA = TestInstructions.InitiateAllTestInstructionsForCA()

	// Initiate map-objects
	testInstructionAttributesMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*TypeAndStructs.TestInstructionAttributeStruct)
	fangEngineClassesMethodsAttributesMap = make(map[TypeAndStructs.OriginalElementUUIDType]*FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct)

	// Convert TestInstruction-data for 'Custody Arrangement' into map-objects
	convertTestInstructionDataIntoMapStructures()

}

// Convert TestInstruction-data for 'Custody Arrangement' into map-objects
func convertTestInstructionDataIntoMapStructures() {

	// Loop TestInstructionsAttributes and create Map
	for _, testInstructionsAttribute := range allTestInstructions_CA.TestInstructionAttribute {
		var tempTestInstructionsAttribute TypeAndStructs.TestInstructionAttributeStruct

		tempTestInstructionsAttribute = testInstructionsAttribute
		testInstructionAttributesMap[tempTestInstructionsAttribute.TestInstructionAttributeUUID] = &tempTestInstructionsAttribute
	}

	// Loop FangEngineClassesMethodsAttributes and create Map
	for _, fangEngineClassesMethodsAttribute := range allTestInstructions_CA.FangEngineClassesMethodsAttributes {
		var tempFangEngineClassesMethodsAttribute FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct

		tempFangEngineClassesMethodsAttribute = fangEngineClassesMethodsAttribute
		fangEngineClassesMethodsAttributesMap[tempFangEngineClassesMethodsAttribute.TestInstructionOriginalUUID] = &tempFangEngineClassesMethodsAttribute
	}

}
