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
	// *** TestInstruction *** 'Search'
	TestInstructionUUID_CA_CustodyAccount_Search               TypeAndStructs.OriginalElementUUIDType = "f3f7723e-d117-434a-9583-30db87d198ce"
	TestInstructionName_CA_CustodyAccount_Search               TypeAndStructs.TestInstructionNameType = "Search"
	TestInstructionTypeUUID_CA_CustodyAccount_Search                                                  = TestInstructionTypeUUID_CA_CustodyAccount
	TestInstructionTypeName_CA_CustodyAccount_Search                                                  = TestInstructionTypeName_CA_CustodyAccount
	TestInstructionDescription_CA_CustodyAccount_Search        string                                 = "Search for a specific Custody Account"
	TestInstructionMouseOverText_CA_CustodyAccount_Search      string                                 = "Search for a specific Custody Account"
	TestInstructionDeprecated_CA_CustodyAccount_Search         bool                                   = false
	TestInstructionEnabled_CA_CustodyAccount_Search            bool                                   = true
	TestInstructionMajorVersionNumber_CA_CustodyAccount_Search int                                    = 0
	TestInstructionMinorVersionNumber_CA_CustodyAccount_Search int                                    = 1
	TestInstructionColor_CA_CustodyAccount_Search              TypeAndStructs.ColorType               = "#00ff00AA"
	TCRuleDeletion_CA_CustodyAccount_Search                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_CustodyAccount_Search                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// *** DropZone *** 'SearchExpectsToSucceed_ExpectsToSucceed'
	TestInstructionDropZoneUUID_CA_CustodyAccount_Search_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "96d9babf-e6ff-49ef-bf69-4ae1c34cf672"
	TestInstructionDropZoneName_CA_CustodyAccount_Search_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "CustodyAccount_Search_ExpectsToSucceed"
	TestInstructionDropZoneDescription_CA_CustodyAccount_Search_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_CA_CustodyAccount_Search_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_CA_CustodyAccount_Search_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_CA_CustodyAccount_Search_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "a2950f1d-abfd-4452-94b6-e77ddfbb8f51" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_CA_ExpectedToBePassed
	TestInstructionAttributeName_CA_CustodyAccount_Search_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_CA_ExpectedToBePassed
	TestInstructionAttributeType_CA_CustodyAccount_Search_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_CA_ExpectedToBePassed
	TestInstructionAttributeActionCommand_CA_CustodyAccount_Search_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_CA_CustodyAccount_Search_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_CA_CustodyAccount_Search_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_CA_CustodyAccount_Search_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_CA_CustodyAccount_Search_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'CustodyAccountId'
	TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId          TypeAndStructs.TestInstructionAttributeUUIDType = "8f7c1ad6-2891-4e28-9b96-567d21385916"
	TestInstructionAttributeName_CA_CustodyAccount_Search_CustodyAccountId          TypeAndStructs.TestInstructionAttributeNameType = "CustodyAccountId"
	TestInstructionAttributeType_CA_CustodyAccount_Search_CustodyAccountId          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_CustodyAccount_Search_CustodyAccountId   string                                          = "The Custody Account to search for"
	TestInstructionAttributeMouseOverText_CA_CustodyAccount_Search_CustodyAccountId string                                          = "The Custody Account to search for"
)

// TestInstruction_CA_CustodyAccount_Search
// Variable that holds the data for the TestInstruction
var TestInstruction_CA_CustodyAccount_Search TestInstruction_CA_TestCaseSetUpStruct

// Initate_TestInstruction_CA_CustodyAccount_Search
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_CustodyAccount_Search() TestInstruction_CA_TestCaseSetUpStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_CustodyAccount_Search = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_CA_CustodyAccount_Search,
			TestInstructionName:         TestInstructionName_CA_CustodyAccount_Search,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_CustodyAccount,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_CustodyAccount,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_CustodyAccount_Search,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_CustodyAccount_Search,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - Search
	TestInstruction_CA_CustodyAccount_Search.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:          TestInstructionName_CA_CustodyAccount_Search,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CustodyAccount_Search,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CustodyAccount_Search,
		TestInstructionDescription:   TestInstructionDescription_CA_CustodyAccount_Search,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_CustodyAccount_Search,
		Deprecated:                   TestInstructionDeprecated_CA_CustodyAccount_Search,
		Enabled:                      TestInstructionEnabled_CA_CustodyAccount_Search,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_CustodyAccount_Search,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_CustodyAccount_Search,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - CustodyAccount_Search
	TestInstruction_CA_CustodyAccount_Search.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:          TestInstructionName_CA_CustodyAccount_Search,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_CustodyAccount_Search,
		TestInstructionTypeName:      TestInstructionTypeName_CA_CustodyAccount_Search,
		Deprecated:                   TestInstructionDeprecated_CA_CustodyAccount_Search,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_CustodyAccount_Search,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_CustodyAccount_Search,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_CustodyAccount_Search,
		TCRuleDeletion:               TCRuleDeletion_CA_CustodyAccount_Search,
		TCRuleSwap:                   TCRuleSwap_CA_CustodyAccount_Search,
		TestInstructionDescription:   TestInstructionDescription_CA_CustodyAccount_Search,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_CustodyAccount_Search,
		Enabled:                      TestInstructionEnabled_CA_CustodyAccount_Search,
	}

	// DropZone 'SearchExpectsToSucceed'
	// DropZone 'CustodyAccount_Search_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: CustodyAccount_Search_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_CA_CustodyAccount_Search_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_CustodyAccount_Search_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:          TestInstructionName_CA_CustodyAccount_Search,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_CustodyAccount_Search_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_CA_CustodyAccount_Search_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_CustodyAccount_Search_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_CustodyAccount_Search_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_CA_CustodyAccount_Search_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_CustodyAccount_Search_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_CustodyAccount_Search_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_CustodyAccount_Search_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_CA_CustodyAccount_Search_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_CA_CustodyAccount_Search_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_CA_CustodyAccount_Search,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_CustodyAccount_Search_ExpectedToBePassed,
	}
	TestInstruction_CA_CustodyAccount_Search.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_CustodyAccount_Search.ImmatureTestInstructionInformation,
		TestInstruction_CA_CustodyAccount_Search_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_CA_CustodyAccount_Search_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_CustodyAccount_Search_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:                           TestInstructionName_CA_CustodyAccount_Search,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_CustodyAccount_Search_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_CustodyAccount_Search_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_CustodyAccount_Search_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_CustodyAccount_Search_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_CustodyAccount_Search_ExpectedToBePassed,
	}
	TestInstruction_CA_CustodyAccount_Search.TestInstructionAttribute = append(
		TestInstruction_CA_CustodyAccount_Search.TestInstructionAttribute,
		TestInstructionAttribute_CA_CustodyAccount_Search_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_CustodyAccount_Search_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_CustodyAccount_Search_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_CA_CustodyAccount_Search.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_CustodyAccount_Search_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// TestInstruction Attribute - 'CustodyAccountId'
	var TestInstructionAttribute_CA_CustodyAccount_Search_CustodyAccountId TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_CustodyAccount_Search_CustodyAccountId = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_CustodyAccount_Search,
		TestInstructionName:                           TestInstructionName_CA_CustodyAccount_Search,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeTypeUUID:              TestInstructionAttributeTypeUUID_CA_Standard,
		TestInstructionAttributeTypeName:              TestInstructionAttributeTypeName_CA_Standard,
		TestInstructionAttributeValueAsString:         Domains.TestInstructionAttributeValueAsStringValue_NO_VALUE,
		TestInstructionAttributeValueUUID:             Domains.TestInstructionAttributeValueUUID_NO_VALUE,
		TestInstructionAttributeVisible:               true,
		TestInstructionAttributeEnabled:               true,
		TestInstructionAttributeMandatory:             true,
		TestInstructionAttributeVisibleInTestCaseArea: false,
		TestInstructionAttributeIsDeprecated:          false,
		TestInstructionAttributeInputMask:             ".",
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_CustodyAccount_Search_CustodyAccountId,
	}
	TestInstruction_CA_CustodyAccount_Search.TestInstructionAttribute = append(
		TestInstruction_CA_CustodyAccount_Search.TestInstructionAttribute,
		TestInstructionAttribute_CA_CustodyAccount_Search_CustodyAccountId)

	// Add FangEngine relation for Attribute - 'CustodyAccountId'
	var tempFangEngineAttributeCustodyAccountId FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeCustodyAccountId = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_CustodyAccount_Search_CustodyAccountId,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_CustodyAccount_Search_CustodyAccountId,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_CustodyAccount_Search_CustodyAccountId,
	}
	//TestInstruction_CA_CustodyAccount_Search.FangEngineClassesMethodsAttributes.Attributes = make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct)
	TestInstruction_CA_CustodyAccount_Search.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_CustodyAccount_Search_CustodyAccountId] = tempFangEngineAttributeCustodyAccountId

	// ImmatureElementModel - Search
	var TestInstructionImmatureElementModel_CA_CustodyAccount_Search TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_CustodyAccount_Search = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_CustodyAccount_Search,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_CustodyAccount_Search),
		PreviousElementUUID:      TestInstructionUUID_CA_CustodyAccount_Search,
		NextElementUUID:          TestInstructionUUID_CA_CustodyAccount_Search,
		FirstChildElementUUID:    TestInstructionUUID_CA_CustodyAccount_Search,
		ParentElementUUID:        TestInstructionUUID_CA_CustodyAccount_Search,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_CustodyAccount_Search,
		TopImmatureElementUUID:   TestInstructionUUID_CA_CustodyAccount_Search,
		IsTopElement:             true,
	}
	TestInstruction_CA_CustodyAccount_Search.ImmatureElementModel = append(
		TestInstruction_CA_CustodyAccount_Search.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_CustodyAccount_Search)

	return TestInstruction_CA_CustodyAccount_Search
}
