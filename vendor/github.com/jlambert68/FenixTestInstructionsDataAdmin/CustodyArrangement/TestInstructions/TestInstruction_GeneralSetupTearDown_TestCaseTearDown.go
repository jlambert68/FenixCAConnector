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
	// *** TestInstruction *** 'TestCaseTearDown'
	TestInstructionUUID_CA_TestCaseTearDown               TypeAndStructs.OriginalElementUUIDType = "874a91df-e5c0-4369-ad3f-ce8aad8bc317"
	TestInstructionName_CA_TestCaseTearDown               TypeAndStructs.TestInstructionNameType = "TestCaseTearDown"
	TestInstructionTypeUUID_CA_TestCaseTearDown                                                  = TestInstructionTypeUUID_CA_GeneralSetUpTearDown
	TestInstructionTypeName_CA_TestCaseTearDown                                                  = TestInstructionTypeName_CA_GeneralSetUpTearDown
	TestInstructionDescription_CA_TestCaseTearDown        string                                 = "TearDown the Custody Arrangements execution engine after executing TestInstructions"
	TestInstructionMouseOverText_CA_TestCaseTearDown      string                                 = "TearDown the Custody Arrangements execution engine after executing TestInstructions"
	TestInstructionDeprecated_CA_TestCaseTearDown         bool                                   = false
	TestInstructionEnabled_CA_TestCaseTearDown            bool                                   = true
	TestInstructionMajorVersionNumber_CA_TestCaseTearDown int                                    = 0
	TestInstructionMinorVersionNumber_CA_TestCaseTearDown int                                    = 1
	TestInstructionColor_CA_TestCaseTearDown              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_CA_TestCaseTearDown                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion020"
	TCRuleSwap_CA_TestCaseTearDown                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap020"

	// *** DropZone *** 'TestCaseTearDown_ExpectsToSucceed'
	TestInstructionDropZoneUUID_CA_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "1e8a5ba3-ae79-425a-a701-acad5a18cd89"
	TestInstructionDropZoneName_CA_TestCaseTearDown_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "TestCaseTearDown_ExpectsToSucceed"
	TestInstructionDropZoneDescription_CA_TestCaseTearDown_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_CA_TestCaseTearDown_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_CA_TestCaseTearDown_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_CA_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "8261c605-1be0-44ca-8b2f-415342a810bd" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_CA_ExpectedToBePassed
	TestInstructionAttributeName_CA_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_CA_ExpectedToBePassed
	TestInstructionAttributeType_CA_TestCaseTearDown_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_CA_ExpectedToBePassed
	TestInstructionAttributeActionCommand_CA_TestCaseTearDown_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_CA_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_CA_TestCaseTearDown_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_CA_TestCaseTearDown_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_CA_TestCaseTearDown_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	FangEngine_Class_Name_CA_TestCaseTearDown = "GeneralSetupTearDown"
)

// TestInstruction_CA_TestCaseTearDown
// Variable that holds the data for the TestInstruction
var TestInstruction_CA_TestCaseTearDown TestInstruction_CA_TestCaseSetUpStruct

// Initate_TestInstruction_CA_TestCaseTearDown
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_TestCaseTearDown() TestInstruction_CA_TestCaseSetUpStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_TestCaseTearDown = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_CA_TestCaseTearDown,
			TestInstructionName:         TestInstructionName_CA_TestCaseTearDown,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralSetupTearDown,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralSetupTearDown,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_GeneralSetupTearDown_TearDown,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_GeneralSetupTearDown_TearDown,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - TestCaseTearDown
	TestInstruction_CA_TestCaseTearDown.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_CA_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_CA_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_CA_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_CA_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_CA_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_TestCaseTearDown,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - TestCaseTearDown
	TestInstruction_CA_TestCaseTearDown.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_CA_TestCaseTearDown,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_TestCaseTearDown,
		TestInstructionTypeName:      TestInstructionTypeName_CA_TestCaseTearDown,
		Deprecated:                   TestInstructionDeprecated_CA_TestCaseTearDown,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_TestCaseTearDown,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_TestCaseTearDown,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_TestCaseTearDown,
		TCRuleDeletion:               TCRuleDeletion_CA_TestCaseTearDown,
		TCRuleSwap:                   TCRuleSwap_CA_TestCaseTearDown,
		TestInstructionDescription:   TestInstructionDescription_CA_TestCaseTearDown,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_TestCaseTearDown,
		Enabled:                      TestInstructionEnabled_CA_TestCaseTearDown,
	}

	// DropZone 'TestCaseTearDown_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: TestCaseTearDown_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_CA_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_TestCaseTearDown_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_TestCaseTearDown,
		TestInstructionName:          TestInstructionName_CA_TestCaseTearDown,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_TestCaseTearDown_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_CA_TestCaseTearDown_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_TestCaseTearDown_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_TestCaseTearDown_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_CA_TestCaseTearDown_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_CA_TestCaseTearDown_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_CA_TestCaseTearDown_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_CA_TestCaseTearDown,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_CA_TestCaseTearDown.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_TestCaseTearDown.ImmatureTestInstructionInformation,
		TestInstruction_CA_TestCaseTearDown_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_CA_TestCaseTearDown_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_TestCaseTearDown_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_TestCaseTearDown,
		TestInstructionName:                           TestInstructionName_CA_TestCaseTearDown,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_TestCaseTearDown_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_TestCaseTearDown_ExpectedToBePassed,
	}
	TestInstruction_CA_TestCaseTearDown.TestInstructionAttribute = append(
		TestInstruction_CA_TestCaseTearDown.TestInstructionAttribute,
		TestInstructionAttribute_CA_TestCaseTearDown_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_TestCaseTearDown_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_CA_TestCaseTearDown.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_TestCaseTearDown_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - TestCaseTearDown
	var TestInstructionImmatureElementModel_CA_TestCaseTearDown TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_TestCaseTearDown = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_TestCaseTearDown,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_TestCaseTearDown),
		PreviousElementUUID:      TestInstructionUUID_CA_TestCaseTearDown,
		NextElementUUID:          TestInstructionUUID_CA_TestCaseTearDown,
		FirstChildElementUUID:    TestInstructionUUID_CA_TestCaseTearDown,
		ParentElementUUID:        TestInstructionUUID_CA_TestCaseTearDown,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_TestCaseTearDown,
		TopImmatureElementUUID:   TestInstructionUUID_CA_TestCaseTearDown,
		IsTopElement:             true,
	}
	TestInstruction_CA_TestCaseTearDown.ImmatureElementModel = append(
		TestInstruction_CA_TestCaseTearDown.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_TestCaseTearDown)

	return TestInstruction_CA_TestCaseTearDown
}
