package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *************************************
	// *** TestInstruction *** 'TestCaseSetUp'
	TestInstructionUUID_CA_TestCaseSetUp               TypeAndStructs.OriginalElementUUIDType = "3a32edd4-7c73-4d05-9971-9a306dce8df7"
	TestInstructionName_CA_TestCaseSetUp               TypeAndStructs.TestInstructionNameType = "TestCaseSetUp"
	TestInstructionTypeUUID_CA_TestCaseSetUp                                                  = TestInstructionTypeUUID_CA_GeneralSetUpTearDown
	TestInstructionTypeName_CA_TestCaseSetUp                                                  = TestInstructionTypeName_CA_GeneralSetUpTearDown
	TestInstructionDescription_CA_TestCaseSetUp        string                                 = "Initiate Custody Arrangements execution engine to be able to execute TestInstructions"
	TestInstructionMouseOverText_CA_TestCaseSetUp      string                                 = "Initiate Custody Arrangements execution engine to be able to execute TestInstructions"
	TestInstructionDeprecated_CA_TestCaseSetUp         bool                                   = false
	TestInstructionEnabled_CA_TestCaseSetUp            bool                                   = true
	TestInstructionMajorVersionNumber_CA_TestCaseSetUp int                                    = 0
	TestInstructionMinorVersionNumber_CA_TestCaseSetUp int                                    = 1
	TestInstructionColor_CA_TestCaseSetUp              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_CA_TestCaseSetUp                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_CA_TestCaseSetUp                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"

	// *** DropZone *** 'TestCaseSetUp_ExpectsToSucceed'
	TestInstructionDropZoneUUID_CA_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "941e4fd9-5f25-4d9c-a39a-07caa5c68c4e"
	TestInstructionDropZoneName_CA_TestCaseSetUp_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseSetUp_ExpectsToSucceed"
	TestInstructionDropZoneDescription_CA_TestCaseSetUp_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_CA_TestCaseSetUp_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_CA_TestCaseSetUp_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_CA_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "c66aac87-6583-490d-bea9-832406babb73" //TestInstructionAttributeUUID_CA_ExpectedToBePassed
	TestInstructionAttributeName_CA_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_CA_ExpectedToBePassed
	TestInstructionAttributeType_CA_TestCaseSetUp_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_CA_ExpectedToBePassed
	TestInstructionAttributeActionCommand_CA_TestCaseSetUp_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_CA_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_CA_TestCaseSetUp_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_CA_TestCaseSetUp_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_CA_TestCaseSetUp_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// TestInstruction_CA_TestCaseSetUp
// Variable that holds the data for the TestInstruction
var TestInstruction_CA_TestCaseSetUp TestInstruction_CA_TestCaseSetUpStruct

// Initate_TestInstruction_CA_TestCaseSetUp
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_TestCaseSetUp() TestInstruction_CA_TestCaseSetUpStruct {

	// Initiate variable to store all TestInstruction data
	TestInstruction_CA_TestCaseSetUp = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_CA_TestCaseSetUp,
			TestInstructionName:         TestInstructionName_CA_TestCaseSetUp,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralSetupTearDown,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralSetupTearDown,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_GeneralSetupTearDown_Setup,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_GeneralSetupTearDown_Setup,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - TestCaseSetUp
	TestInstruction_CA_TestCaseSetUp.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_CA_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_CA_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_CA_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_CA_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_CA_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_TestCaseSetUp,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseSetUp
	TestInstruction_CA_TestCaseSetUp.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_CA_TestCaseSetUp,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_TestCaseSetUp,
		TestInstructionTypeName:      TestInstructionTypeName_CA_TestCaseSetUp,
		Deprecated:                   TestInstructionDeprecated_CA_TestCaseSetUp,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_TestCaseSetUp,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_TestCaseSetUp,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_TestCaseSetUp,
		TCRuleDeletion:               TCRuleDeletion_CA_TestCaseSetUp,
		TCRuleSwap:                   TCRuleSwap_CA_TestCaseSetUp,
		TestInstructionDescription:   TestInstructionDescription_CA_TestCaseSetUp,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_TestCaseSetUp,
		Enabled:                      TestInstructionEnabled_CA_TestCaseSetUp,
	}

	// DropZone 'TestCaseSetUp_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseSetUp_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_CA_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_TestCaseSetUp_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_TestCaseSetUp,
		TestInstructionName:          TestInstructionName_CA_TestCaseSetUp,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_TestCaseSetUp_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_CA_TestCaseSetUp_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_TestCaseSetUp_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_TestCaseSetUp_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_CA_TestCaseSetUp_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_CA_TestCaseSetUp_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_CA_TestCaseSetUp_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_CA_TestCaseSetUp,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_CA_TestCaseSetUp.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_TestCaseSetUp.ImmatureTestInstructionInformation,
		TestInstruction_CA_TestCaseSetUp_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_CA_TestCaseSetUp_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_TestCaseSetUp_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_TestCaseSetUp,
		TestInstructionName:                           TestInstructionName_CA_TestCaseSetUp,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID:              TestInstructionAttributeTypeUUID_CA_ExpectedToPass,
		TestInstructionAttributeTypeName:              TestInstructionAttributeTypeName_CA_ExpectedToPass,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_TestCaseSetUp_ExpectedToBePassed,
	}
	TestInstruction_CA_TestCaseSetUp.TestInstructionAttribute = append(
		TestInstruction_CA_TestCaseSetUp.TestInstructionAttribute,
		TestInstructionAttribute_CA_TestCaseSetUp_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_TestCaseSetUp_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_CA_TestCaseSetUp.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_TestCaseSetUp_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseSetUp
	var TestInstructionImmatureElementModel_CA_TestCaseSetUp TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_TestCaseSetUp = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_TestCaseSetUp,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_TestCaseSetUp),
		PreviousElementUUID:      TestInstructionUUID_CA_TestCaseSetUp,
		NextElementUUID:          TestInstructionUUID_CA_TestCaseSetUp,
		FirstChildElementUUID:    TestInstructionUUID_CA_TestCaseSetUp,
		ParentElementUUID:        TestInstructionUUID_CA_TestCaseSetUp,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_TestCaseSetUp,
		TopImmatureElementUUID:   TestInstructionUUID_CA_TestCaseSetUp,
		IsTopElement:             true,
	}
	TestInstruction_CA_TestCaseSetUp.ImmatureElementModel = append(
		TestInstruction_CA_TestCaseSetUp.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_TestCaseSetUp)

	return TestInstruction_CA_TestCaseSetUp
}
