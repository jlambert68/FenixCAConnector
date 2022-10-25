package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** DeleteSelectedInstructedParties
	TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties               TypeAndStructs.OriginalElementUUIDType = "e25e07ef-473b-42be-8a2e-5e1eb3c76cc1"
	TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties               TypeAndStructs.TestInstructionNameType = "DeleteSelectedInstructedParties"
	TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties                                                  = TestInstructionTypeUUID_CA_SettlementAgreementManagement
	TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedInstructedParties                                                  = TestInstructionTypeName_CA_SettlementAgreementManagement
	TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties        string                                 = "DeleteSelectedInstructedParties for an existing Settlement Agreement"
	TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties      string                                 = "DeleteSelectedInstructedParties for an existing Settlement Agreement"
	TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedInstructedParties         bool                                   = false
	TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedInstructedParties            bool                                   = true
	TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties int                                    = 0
	TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties int                                    = 1
	TestInstructionColor_CA_SettlementAgreement_DeleteSelectedInstructedParties              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_SettlementAgreement_DeleteSelectedInstructedParties                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_SettlementAgreement_DeleteSelectedInstructedParties                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// *** DropZone *** 'SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed'
	TestInstructionDropZoneUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "1f9b546b-9d83-4fb3-aa59-be0ed15593b7"
	TestInstructionDropZoneName_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed"
	TestInstructionDropZoneDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "897e2250-e524-4bf6-ab63-619b1c7dd157" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_CA_ExpectedToBePassed
	TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_CA_ExpectedToBePassed
	TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_CA_ExpectedToBePassed
	TestInstructionAttributeActionCommand_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"
)

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties TestInstruction_CA_TestCaseSetUpStruct

// Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties() TestInstruction_CA_TestCaseSetUpStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
			TestInstructionName:         TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_SettlementAgreement,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_SettlementAgreement,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_SettlementAgreement_DeleteSelectedInstructedParties,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - SettlementAgreement_DeleteSelectedInstructedParties
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - SettlementAgreement_DeleteSelectedInstructedParties
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TCRuleDeletion:               TCRuleDeletion_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TCRuleSwap:                   TCRuleSwap_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedInstructedParties,
	}

	// DropZone 'SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.ImmatureTestInstructionInformation,
		TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// ImmatureElementModel - 'SettlementAgreement' - MethodName: 'DeleteSelectedInstructedParties'
	var TestInstructionImmatureElementModel_CA_SettlementAgreement_DeleteSelectedInstructedParties TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_SettlementAgreement_DeleteSelectedInstructedParties = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_SettlementAgreement_DeleteSelectedInstructedParties),
		PreviousElementUUID:      TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		NextElementUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		FirstChildElementUUID:    TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		ParentElementUUID:        TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		TopImmatureElementUUID:   TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedInstructedParties,
		IsTopElement:             true,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.ImmatureElementModel = append(
		TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_SettlementAgreement_DeleteSelectedInstructedParties)

	return TestInstruction_CA_SettlementAgreement_DeleteSelectedInstructedParties
}
