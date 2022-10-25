package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** Edit
	TestInstructionUUID_CA_SettlementAgreement_Edit               TypeAndStructs.OriginalElementUUIDType = "19c2a8b9-c48b-4f5e-9f41-a9b571f77ebb"
	TestInstructionName_CA_SettlementAgreement_Edit               TypeAndStructs.TestInstructionNameType = "Edit"
	TestInstructionTypeUUID_CA_SettlementAgreement_Edit                                                  = TestInstructionTypeUUID_CA_SettlementAgreementManagement
	TestInstructionTypeName_CA_SettlementAgreement_Edit                                                  = TestInstructionTypeName_CA_SettlementAgreementManagement
	TestInstructionDescription_CA_SettlementAgreement_Edit        string                                 = "Edit an existing Settlement Agreement"
	TestInstructionMouseOverText_CA_SettlementAgreement_Edit      string                                 = "Edit an existing Settlement Agreement"
	TestInstructionDeprecated_CA_SettlementAgreement_Edit         bool                                   = false
	TestInstructionEnabled_CA_SettlementAgreement_Edit            bool                                   = true
	TestInstructionMajorVersionNumber_CA_SettlementAgreement_Edit int                                    = 0
	TestInstructionMinorVersionNumber_CA_SettlementAgreement_Edit int                                    = 1
	TestInstructionColor_CA_SettlementAgreement_Edit              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_SettlementAgreement_Edit                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_SettlementAgreement_Edit                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// *** DropZone *** 'SettlementAgreement_Edit_ExpectsToSucceed'
	TestInstructionDropZoneUUID_CA_SettlementAgreement_Edit_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "35b69f2a-4672-4b31-acf5-a01db3be3f70"
	TestInstructionDropZoneName_CA_SettlementAgreement_Edit_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "SettlementAgreement_Edit_ExpectsToSucceed"
	TestInstructionDropZoneDescription_CA_SettlementAgreement_Edit_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_CA_SettlementAgreement_Edit_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_CA_SettlementAgreement_Edit_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "c1374659-3c59-4854-bd8b-623d6a6f59b1"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_CA_ExpectedToBePassed
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_CA_ExpectedToBePassed
	TestInstructionAttributeActionCommand_CA_SettlementAgreement_Edit_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_CA_SettlementAgreement_Edit_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_CA_SettlementAgreement_Edit_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'TestDataDuplicateCheck'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck          TypeAndStructs.TestInstructionAttributeUUIDType = "48763abc-b792-4bff-9132-258ec747f7ea"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDuplicateCheck          TypeAndStructs.TestInstructionAttributeNameType = "TestDataDuplicateCheck"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDuplicateCheck          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDuplicateCheck   string                                          = "TestDataDuplicateCheck is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDuplicateCheck string                                          = "TestDataDuplicateCheck is xxxxx"

	// Attribute - 'TestDataContractualRule'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule          TypeAndStructs.TestInstructionAttributeUUIDType = "ef315036-f4bc-4255-a090-110a4e9ec836"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataContractualRule          TypeAndStructs.TestInstructionAttributeNameType = "TestDataContractualRule"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataContractualRule          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataContractualRule   string                                          = "TestDataContractualRule is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataContractualRule string                                          = "TestDataContractualRule is xxxxx"

	// Attribute - 'TestDataMatchAffirmEligible'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible          TypeAndStructs.TestInstructionAttributeUUIDType = "7894ee2d-c6d5-47a3-97f3-5b75cabe1abb"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible          TypeAndStructs.TestInstructionAttributeNameType = "TestDataMatchAffirmEligible"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible   string                                          = "TestDataMatchAffirmEligible is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible string                                          = "TestDataMatchAffirmEligible is xxxxx"

	// Attribute - 'TestDataDirectAffirmInd'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd          TypeAndStructs.TestInstructionAttributeUUIDType = "0dae0c87-c659-46e1-a03b-15231f91d567"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd          TypeAndStructs.TestInstructionAttributeNameType = "TestDataDirectAffirmInd"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd   string                                          = "TestDataDirectAffirmInd is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd string                                          = "TestDataDirectAffirmInd is xxxxx"

	// Attribute - 'TestDataAffirmDublicateCheck'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck          TypeAndStructs.TestInstructionAttributeUUIDType = "0b22e854-bc7c-4f9b-9964-051ce216f14e"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck          TypeAndStructs.TestInstructionAttributeNameType = "TestDataAffirmDublicateCheck"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck   string                                          = "TestDataAffirmDublicateCheck is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck string                                          = "TestDataAffirmDublicateCheck is xxxxx"

	// Attribute - 'TestDataAffirmCancelledInd'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd          TypeAndStructs.TestInstructionAttributeUUIDType = "36ebedb6-3cdb-437a-84e5-5fa18bd7b383"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd          TypeAndStructs.TestInstructionAttributeNameType = "TestDataAffirmCancelledInd"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd   string                                          = "TestDataAffirmCancelledInd is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd string                                          = "TestDataAffirmCancelledInd is xxxxx"

	// Attribute - 'TestDataNearMatch'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch          TypeAndStructs.TestInstructionAttributeUUIDType = "c08d107f-91f4-42a5-aff7-4e50e66bbfc5"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataNearMatch          TypeAndStructs.TestInstructionAttributeNameType = "TestDataNearMatch"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataNearMatch          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataNearMatch   string                                          = "TestDataNearMatch is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataNearMatch string                                          = "TestDataNearMatch is xxxxx"

	// Attribute - 'TestDataReturnDeliveryService'
	TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService          TypeAndStructs.TestInstructionAttributeUUIDType = "b4addbd9-95d4-46d8-a044-946d9704afc5"
	TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService          TypeAndStructs.TestInstructionAttributeNameType = "TestDataReturnDeliveryService"
	TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService   string                                          = "TestDataReturnDeliveryService is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService string                                          = "TestDataReturnDeliveryService is xxxxx"
)

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_SettlementAgreement_Edit TestInstruction_CA_TestCaseSetUpStruct

// Initate_TestInstruction_CA_SettlementAgreement_Edit
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_SettlementAgreement_Edit() TestInstruction_CA_TestCaseSetUpStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_SettlementAgreement_Edit = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_CA_SettlementAgreement_Edit,
			TestInstructionName:         TestInstructionName_CA_SettlementAgreement_Edit,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_SettlementAgreement,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_SettlementAgreement,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_SettlementAgreement_Edit,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_SettlementAgreement_Edit,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - SettlementAgreement_Edit
	TestInstruction_CA_SettlementAgreement_Edit.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_Edit,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_Edit,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_Edit,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_Edit,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_Edit,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_Edit,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_Edit,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_Edit,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - SettlementAgreement_Edit
	TestInstruction_CA_SettlementAgreement_Edit.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_Edit,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_Edit,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_Edit,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_Edit,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_Edit,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_SettlementAgreement_Edit,
		TCRuleDeletion:               TCRuleDeletion_CA_SettlementAgreement_Edit,
		TCRuleSwap:                   TCRuleSwap_CA_SettlementAgreement_Edit,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_Edit,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_Edit,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_Edit,
	}

	// DropZone 'SettlementAgreement_Edit_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: SettlementAgreement_Edit_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_CA_SettlementAgreement_Edit_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_SettlementAgreement_Edit_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_Edit,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_SettlementAgreement_Edit_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_CA_SettlementAgreement_Edit_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_SettlementAgreement_Edit_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_SettlementAgreement_Edit_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_CA_SettlementAgreement_Edit_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_CA_SettlementAgreement_Edit,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_SettlementAgreement_Edit_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_Edit.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_SettlementAgreement_Edit.ImmatureTestInstructionInformation,
		TestInstruction_CA_SettlementAgreement_Edit_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// TestInstruction Attribute - 'TestDataDuplicateCheck'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDuplicateCheck TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDuplicateCheck = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDuplicateCheck)

	// Add FangEngine relation for Attribute - 'TestDataDuplicateCheck'
	var tempFangEngineAttributeTestDataDuplicateCheck FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataDuplicateCheck = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataDuplicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck] = tempFangEngineAttributeTestDataDuplicateCheck

	// TestInstruction Attribute - 'TestDataContractualRule'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataContractualRule TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataContractualRule = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataContractualRule,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataContractualRule,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataContractualRule)

	// Add FangEngine relation for Attribute - 'TestDataContractualRule'
	var tempFangEngineAttributeTestDataContractualRule FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataContractualRule = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataContractualRule,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataContractualRule,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataContractualRule,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataContractualRule] = tempFangEngineAttributeTestDataContractualRule

	// TestInstruction Attribute - 'TestDataMatchAffirmEligible'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible)

	// Add FangEngine relation for Attribute - 'TestDataMatchAffirmEligible'
	var tempFangEngineAttributeTestDataMatchAffirmEligible FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataMatchAffirmEligible = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible] = tempFangEngineAttributeTestDataMatchAffirmEligible

	// TestInstruction Attribute - 'TestDataDirectAffirmInd'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd)

	// Add FangEngine relation for Attribute - 'TestDataDirectAffirmInd'
	var tempFangEngineAttributeTestDataDirectAffirmInd FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataDirectAffirmInd = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd] = tempFangEngineAttributeTestDataDirectAffirmInd

	// TestInstruction Attribute - 'TestDataAffirmDublicateCheck'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck)

	// Add FangEngine relation for Attribute - 'TestDataAffirmDublicateCheck'
	var tempFangEngineAttributeTestDataAffirmDublicateCheck FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataAffirmDublicateCheck = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck] = tempFangEngineAttributeTestDataAffirmDublicateCheck

	// TestInstruction Attribute - 'TestDataAffirmCancelledInd'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd)

	// Add FangEngine relation for Attribute - 'TestDataAffirmCancelledInd'
	var tempFangEngineAttributeTestDataAffirmCancelledInd FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataAffirmCancelledInd = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd] = tempFangEngineAttributeTestDataAffirmCancelledInd

	// TestInstruction Attribute - 'TestDataNearMatch'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataNearMatch TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataNearMatch = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataNearMatch,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataNearMatch,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataNearMatch)

	// Add FangEngine relation for Attribute - 'TestDataNearMatch'
	var tempFangEngineAttributeTestDataNearMatch FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataNearMatch = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataNearMatch,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataNearMatch,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataNearMatch,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataNearMatch] = tempFangEngineAttributeTestDataNearMatch

	// TestInstruction Attribute - 'TestDataReturnDeliveryService'
	var TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_Edit,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
	}
	TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_Edit.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService)

	// Add FangEngine relation for Attribute - 'TestDataReturnDeliveryService'
	var tempFangEngineAttributeTestDataReturnDeliveryService FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeTestDataReturnDeliveryService = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService,
	}
	TestInstruction_CA_SettlementAgreement_Edit.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService] = tempFangEngineAttributeTestDataReturnDeliveryService

	// ImmatureElementModel - 'SettlementAgreement' - MethodName: 'Edit'
	var TestInstructionImmatureElementModel_CA_SettlementAgreement_Edit TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_SettlementAgreement_Edit = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_SettlementAgreement_Edit,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_SettlementAgreement_Edit),
		PreviousElementUUID:      TestInstructionUUID_CA_SettlementAgreement_Edit,
		NextElementUUID:          TestInstructionUUID_CA_SettlementAgreement_Edit,
		FirstChildElementUUID:    TestInstructionUUID_CA_SettlementAgreement_Edit,
		ParentElementUUID:        TestInstructionUUID_CA_SettlementAgreement_Edit,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_SettlementAgreement_Edit,
		TopImmatureElementUUID:   TestInstructionUUID_CA_SettlementAgreement_Edit,
		IsTopElement:             true,
	}
	TestInstruction_CA_SettlementAgreement_Edit.ImmatureElementModel = append(
		TestInstruction_CA_SettlementAgreement_Edit.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_SettlementAgreement_Edit)

	return TestInstruction_CA_SettlementAgreement_Edit
}
