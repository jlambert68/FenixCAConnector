package restCallsToCAEngine

import (
	"FenixCAConnector/common_config"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/TestInstructions"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/sirupsen/logrus"
)

// InitiateRestCallsToCAEngine
// Do all initiation to have restEngine be able to do RestCalls to Custody Arrangements FangEngine
func InitiateRestCallsToCAEngine() (err error)  {


	// Load all TestInstruction-data for 'Custody Arrangement'
	allTestInstructions_CA = TestInstructions.InitiateAllTestInstructionsForCA()

	// Initiate map-objects
	testInstructionAttributesMap = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]*TypeAndStructs.TestInstructionAttributeStruct)
	fangEngineClassesMethodsAttributesMap = make(map[TypeAndStructs.OriginalElementUUIDType]*FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct)

	// Convert TestInstruction-data for 'Custody Arrangement' into map-objects

	if err != nil {
		return err
	}

	return err
}

// Convert TestInstruction-data for 'Custody Arrangement' into map-objects
func convertTestInstructionDataIntoMapStructures() (err error) {

	// Loop TestInstructionsAttributes and create Map
	for _, testInstructionsAttribute := range 	allTestInstructions_CA.TestInstructionAttribute {
		var tempTestInstructionsAttribute TypeAndStructs.TestInstructionAttributeStruct

		tempTestInstructionsAttribute = testInstructionsAttribute
		testInstructionAttributesMap[tempTestInstructionsAttribute.TestInstructionAttributeUUID] = &tempTestInstructionsAttribute
	}

	// Loop FangEngineClassesMethodsAttributes and create Map
	for _, fangEngineClassesMethodsAttribute := range 	allTestInstructions_CA.FangEngineClassesMethodsAttributes {
		var tempFangEngineClassesMethodsAttribute FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct

		tempFangEngineClassesMethodsAttribute = fangEngineClassesMethodsAttribute
		fangEngineClassesMethodsAttributesMap[tempFangEngineClassesMethodsAttribute.TestInstructionAttributeUUID] = &tempFangEngineClassesMethodsAttribute
	}

	fangEngineClassesMethodsAttributesMap map[TypeAndStructs.OriginalElementUUIDType]*FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct

	common_config.Logger.WithFields(logrus.Fields{
		"ID":    "eee9f6d1-d773-422c-ac2c-1bdb1b5518b7",
		"error": err,
	}).Error("Couldn't convert ")


	return err
}

