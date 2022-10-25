package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/Domains"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TestCaseModelElementTypes"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/shared_code"
)

const (

	// *** TestInstruction *** DeleteSelectedSwift
	TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift               TypeAndStructs.OriginalElementUUIDType = "889b3337-fd6d-4bdb-9dc6-0fd6ed04a43b"
	TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift               TypeAndStructs.TestInstructionNameType = "DeleteSelectedSwift"
	TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedSwift                                                  = TestInstructionTypeUUID_CA_SettlementAgreementManagement
	TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedSwift                                                  = TestInstructionTypeName_CA_SettlementAgreementManagement
	TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedSwift        string                                 = "DeleteSelectedSwift for an existing Settlement Agreement"
	TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedSwift      string                                 = "DeleteSelectedSwift for an existing Settlement Agreement"
	TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedSwift         bool                                   = false
	TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedSwift            bool                                   = true
	TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedSwift int                                    = 0
	TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedSwift int                                    = 1
	TestInstructionColor_CA_SettlementAgreement_DeleteSelectedSwift              TypeAndStructs.ColorType               = "xxxxx"
	TCRuleDeletion_CA_SettlementAgreement_DeleteSelectedSwift                    TypeAndStructs.TCRuleDeletionType      = "TCRuleDeletion011"
	TCRuleSwap_CA_SettlementAgreement_DeleteSelectedSwift                        TypeAndStructs.TCRuleSwapType          = "TCRuleSwap011"

	// *** DropZone *** 'SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed'
	TestInstructionDropZoneUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed        TypeAndStructs.DropZoneUUIDType = "39dfff94-2f5a-47ae-a3f5-29556e65cd84"
	TestInstructionDropZoneName_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed        TypeAndStructs.DropZoneNameType = "SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed"
	TestInstructionDropZoneDescription_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneMouseOver_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed   string                          = "Presets attribute that TestInstruction expects to succeed in its execution"
	TestInstructionDropZoneColor_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed       TypeAndStructs.ColorType        = "#00000000"

	// Attribute - 'ExpectedToBePassed'
	TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeUUIDType = "26c83dd2-ac59-4c6b-aa55-6da27033d1c5" // TODO fix so they use the same UUID, Can't bu done now because UUID is key in Attrubutes-table in DB .TestInstructionAttributeUUID_CA_ExpectedToBePassed
	TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeNameType = TestInstructionAttributeName_CA_ExpectedToBePassed
	TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed               TypeAndStructs.TestInstructionAttributeTypeType = TestInstructionAttributeType_CA_ExpectedToBePassed
	TestInstructionAttributeActionCommand_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed      TypeAndStructs.AttributeActionCommandType       = Domains.AttributeActionCommand_USE_DROPZONE_VALUE_FOR_ATTRIBUTE
	TestInstructionAttributeValueAsStringValue_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed TypeAndStructs.AttributeValueAsStringType       = Domains.TestInstructionAttributeValueAsStringValue_TRUE
	TestInstructionAttributeValueUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed          TypeAndStructs.AttributeValueUUIDType           = Domains.TestInstructionAttributeValueUUID_TRUE
	TestInstructionAttributeDescription_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed        string                                          = "Should the TestInstruction execution to be expected to succeed or not"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed      string                                          = "Should the TestInstruction execution to be expected to succeed or not"

	// Attribute - 'bicAddress'
	TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress          TypeAndStructs.TestInstructionAttributeUUIDType = "8d73b3c3-59f1-48b9-88bb-cf2f5085ca0d"
	TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress          TypeAndStructs.TestInstructionAttributeNameType = "bicAddress"
	TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress          TypeAndStructs.TestInstructionAttributeTypeType = "TEXTBOX"
	TestInstructionAttributeDescription_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress   string                                          = "bicAddress is xxxxx"
	TestInstructionAttributeMouseOverText_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress string                                          = "bicAddress is xxxxx"
)

// Variable that holds the data for the TestInstruction
var TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift TestInstruction_CA_TestCaseSetUpStruct

// Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift
// Function that creates all data for the TestInstruction
func Initate_TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift() TestInstruction_CA_TestCaseSetUpStruct {

	// Initiate variable to be able to store all TestInstruction data
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift = TestInstruction_CA_TestCaseSetUpStruct{
		TestInstruction:                    TypeAndStructs.TestInstructionStruct{},
		BasicTestInstructionInformation:    TypeAndStructs.BasicTestInstructionInformationStruct{},
		ImmatureTestInstructionInformation: nil,
		TestInstructionAttribute:           nil,
		ImmatureElementModel:               nil,
		FangEngineClassesMethodsAttributes: FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct{
			TestInstructionOriginalUUID: TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
			TestInstructionName:         TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift,
			FangEngineClassNameUUID:     FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_SettlementAgreement,
			FangEngineClassNameNAME:     FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_SettlementAgreement,
			FangEngineMethodNameUUID:    FangEngineClassesAndMethods.FangEngine_MethodName_UUID_CA_SettlementAgreement_DeleteSelectedSwift,
			FangEngineMethodNameNAME:    FangEngineClassesAndMethods.FangEngine_MethodName_Name_CA_SettlementAgreement_DeleteSelectedSwift,
			Attributes:                  make(map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineClassesAndMethods.FangEngineAttributesStruct),
		},
	}

	// Create TimeStamp to be used on all places where creation/update timestamp is needed
	updatedTimeStamp := TypeAndStructs.UpdatedTimeStampType(shared_code.GenerateDatetimeTimeStampForDB())

	// TestInstruction - SettlementAgreement_DeleteSelectedSwift
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.TestInstruction = TypeAndStructs.TestInstructionStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedSwift,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedSwift,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedSwift,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedSwift,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedSwift,
		UpdatedTimeStamp:             updatedTimeStamp,
	}

	// BasicTestInstructionInformation - SettlementAgreement_DeleteSelectedSwift
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.BasicTestInstructionInformation = TypeAndStructs.BasicTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionTypeUUID:      TestInstructionTypeUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionTypeName:      TestInstructionTypeName_CA_SettlementAgreement_DeleteSelectedSwift,
		Deprecated:                   TestInstructionDeprecated_CA_SettlementAgreement_DeleteSelectedSwift,
		MajorVersionNumber:           TestInstructionMajorVersionNumber_CA_SettlementAgreement_DeleteSelectedSwift,
		MinorVersionNumber:           TestInstructionMinorVersionNumber_CA_SettlementAgreement_DeleteSelectedSwift,
		UpdatedTimeStamp:             updatedTimeStamp,
		TestInstructionColor:         TestInstructionColor_CA_SettlementAgreement_DeleteSelectedSwift,
		TCRuleDeletion:               TCRuleDeletion_CA_SettlementAgreement_DeleteSelectedSwift,
		TCRuleSwap:                   TCRuleSwap_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionDescription:   TestInstructionDescription_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionMouseOverText: TestInstructionMouseOverText_CA_SettlementAgreement_DeleteSelectedSwift,
		Enabled:                      TestInstructionEnabled_CA_SettlementAgreement_DeleteSelectedSwift,
	}

	// DropZone 'SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed'
	// ImmatureTestInstructionInformation  - DropZone: SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed, Attr: ExpectedToBePassed
	var TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed = TypeAndStructs.ImmatureTestInstructionInformationStruct{
		DomainUUID:                   Domains.DomainUUID_CA,
		DomainName:                   Domains.DomainName_CA,
		TestInstructionUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionName:          TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift,
		DropZoneUUID:                 TestInstructionDropZoneUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed,
		DropZoneName:                 TestInstructionDropZoneName_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed,
		DropZoneDescription:          TestInstructionDropZoneDescription_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed,
		DropZoneMouseOver:            TestInstructionDropZoneMouseOver_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed,
		DropZoneColor:                TestInstructionDropZoneColor_CA_SettlementAgreement_DeleteSelectedSwift_ExpectsToSucceed,
		TestInstructionAttributeType: TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		TestInstructionAttributeUUID: TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		TestInstructionAttributeName: TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		AttributeValueAsString:       TestInstructionAttributeValueAsStringValue_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		AttributeValueUUID:           TestInstructionAttributeValueUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		FirstImmatureElementUUID:     TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		AttributeActionCommand:       TestInstructionAttributeActionCommand_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.ImmatureTestInstructionInformation = append(
		TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.ImmatureTestInstructionInformation,
		TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed)

	// TestInstruction Attribute - 'ExpectedToBePassed'
	var TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed)

	// Add FangEngine relation for Attribute - 'ExpectedToBePassed'
	var tempFangEngineAttributeExpectedToBePassed FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributeExpectedToBePassed = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_ExpectedToPass,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_ClassName_UUID_CA_GeneralAttribute_ExpectedToBePassed,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_ClassName_Name_CA_GeneralAttribute_ExpectedToBePassed,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_ExpectedToBePassed] = tempFangEngineAttributeExpectedToBePassed

	// TestInstruction Attribute - 'bicAddress'
	var TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress TypeAndStructs.TestInstructionAttributeStruct
	TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress = TypeAndStructs.TestInstructionAttributeStruct{
		DomainUUID:                                    Domains.DomainUUID_CA,
		DomainName:                                    Domains.DomainName_CA,
		TestInstructionUUID:                           TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionName:                           TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift,
		TestInstructionAttributeUUID:                  TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
		TestInstructionAttributeName:                  TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
		TestInstructionAttributeDescription:           TestInstructionAttributeDescription_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
		TestInstructionAttributeMouseOver:             TestInstructionAttributeMouseOverText_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
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
		TestInstructionAttributeType:                  TestInstructionAttributeType_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.TestInstructionAttribute = append(
		TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.TestInstructionAttribute,
		TestInstructionAttribute_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress)

	// Add FangEngine relation for Attribute - 'bicAddress'
	var tempFangEngineAttributebicAddress FangEngineClassesAndMethods.FangEngineAttributesStruct
	tempFangEngineAttributebicAddress = FangEngineClassesAndMethods.FangEngineAttributesStruct{
		TestInstructionAttributeUUID:     TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
		TestInstructionAttributeName:     TestInstructionAttributeName_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
		TestInstructionAttributeTypeUUID: TestInstructionAttributeTypeUUID_CA_Standard,
		FangEngineAttributeNameUUID:      FangEngineClassesAndMethods.FangEngine_AttributeName_UUID_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
		FangEngineAttributeNameName:      FangEngineClassesAndMethods.FangEngine_AttributeName_Name_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.FangEngineClassesMethodsAttributes.Attributes[TestInstructionAttributeUUID_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress] = tempFangEngineAttributebicAddress

	// ImmatureElementModel - 'SettlementAgreement' - MethodName: 'DeleteSelectedSwift'
	var TestInstructionImmatureElementModel_CA_SettlementAgreement_DeleteSelectedSwift TypeAndStructs.ImmatureElementModelMessageStruct
	TestInstructionImmatureElementModel_CA_SettlementAgreement_DeleteSelectedSwift = TypeAndStructs.ImmatureElementModelMessageStruct{
		DomainUUID:               Domains.DomainUUID_CA,
		DomainName:               Domains.DomainName_CA,
		ImmatureElementUUID:      TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		ImmatureElementName:      TypeAndStructs.OriginalElementNameType(TestInstructionName_CA_SettlementAgreement_DeleteSelectedSwift),
		PreviousElementUUID:      TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		NextElementUUID:          TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		FirstChildElementUUID:    TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		ParentElementUUID:        TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TestCaseModelElementType: TestCaseModelElementTypes.TestCaseModelElementType_TI,
		OriginalElementUUID:      TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		TopImmatureElementUUID:   TestInstructionUUID_CA_SettlementAgreement_DeleteSelectedSwift,
		IsTopElement:             true,
	}
	TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.ImmatureElementModel = append(
		TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift.ImmatureElementModel,
		TestInstructionImmatureElementModel_CA_SettlementAgreement_DeleteSelectedSwift)

	return TestInstruction_CA_SettlementAgreement_DeleteSelectedSwift
}
