package FangEngineClassesAndMethods

import "github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"

// Types for FangEngine
type FangEngine_ClassName_UUID_CA_Type string
type FangEngine_ClassName_Name_CA_Type string
type FangEngine_MethodName_UUID_CA_Type string
type FangEngine_MethodName_Name_CA_Type string
type FangEngine_AttributeName_UUID_CA_Type string
type FangEngine_AttributeName_Name_CA_Type string

// Type this is used for specifying Classes, Methods and Attributes for FangEngine which is used by Custody Arrangements TestAutomation
type FangEngineClassesMethodsAttributesStruct struct {
	TestInstructionOriginalUUID TypeAndStructs.OriginalElementUUIDType
	TestInstructionName         TypeAndStructs.TestInstructionNameType
	FangEngineClassNameUUID     FangEngine_ClassName_UUID_CA_Type
	FangEngineClassNameNAME     FangEngine_ClassName_Name_CA_Type
	FangEngineMethodNameUUID    FangEngine_MethodName_UUID_CA_Type
	FangEngineMethodNameNAME    FangEngine_MethodName_Name_CA_Type
	Attributes                  map[TypeAndStructs.TestInstructionAttributeUUIDType]FangEngineAttributesStruct
}

type FangEngineAttributesStruct struct {
	TestInstructionAttributeUUID     TypeAndStructs.TestInstructionAttributeUUIDType
	TestInstructionAttributeName     TypeAndStructs.TestInstructionAttributeNameType
	TestInstructionAttributeTypeUUID TypeAndStructs.TestInstructionAttributeTypeUUIDType
	FangEngineAttributeNameUUID      FangEngine_AttributeName_UUID_CA_Type
	FangEngineAttributeNameName      FangEngine_AttributeName_Name_CA_Type
}

// Classes, Methods and their Parameters in FangEngine for Custody Arrangement
const (

	// General Attribute - ''
	FangEngine_ClassName_UUID_CA_GeneralAttribute_ExpectedToBePassed FangEngine_AttributeName_UUID_CA_Type = "2e78043c-00bf-4a4a-8690-9981425ca239"
	FangEngine_ClassName_Name_CA_GeneralAttribute_ExpectedToBePassed FangEngine_AttributeName_Name_CA_Type = "expectedToBePassed"

	// ClassName - ***** 'GeneralSetupTearDown' *****
	FangEngine_ClassName_UUID_CA_GeneralSetupTearDown FangEngine_ClassName_UUID_CA_Type = "998b8ec9-6dc8-477a-a9a6-453417279537"
	FangEngine_ClassName_Name_CA_GeneralSetupTearDown FangEngine_ClassName_Name_CA_Type = "GeneralSetupTearDown"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'Setup'
	FangEngine_MethodName_UUID_CA_GeneralSetupTearDown_Setup FangEngine_MethodName_UUID_CA_Type = "4d5ee1e6-858c-4b9f-8e15-3cb350b6367d"
	FangEngine_MethodName_Name_CA_GeneralSetupTearDown_Setup FangEngine_MethodName_Name_CA_Type = "Setup"

	// ClassName: 'GeneralSetupTearDown' - MethodName: 'TearDown'
	FangEngine_MethodName_UUID_CA_GeneralSetupTearDown_TearDown FangEngine_MethodName_UUID_CA_Type = "e58dd894-1d22-47ea-96d8-d2696eb347a0"
	FangEngine_MethodName_Name_CA_GeneralSetupTearDown_TearDown FangEngine_MethodName_Name_CA_Type = "TearDown"

	// ClassName - ****** 'CustodyAccount' ******
	FangEngine_ClassName_UUID_CA_CustodyAccount FangEngine_ClassName_UUID_CA_Type = "d0623650-a2ea-422b-af6a-6114d5197c4a"
	FangEngine_ClassName_Name_CA_CustodyAccount FangEngine_ClassName_Name_CA_Type = "CustodyAccount"

	// ClassName: 'CustodyAccount' - MethodName: 'Search'
	FangEngine_MethodName_UUID_CA_CustodyAccount_Search FangEngine_MethodName_UUID_CA_Type = "024cd19a-d41d-4abf-95e6-207544fa41c3"
	FangEngine_MethodName_Name_CA_CustodyAccount_Search FangEngine_MethodName_Name_CA_Type = "Search"

	// ClassName: 'CustodyAccount' - MethodName: 'Search' - Attributes
	FangEngine_AttributeName_UUID_CA_CustodyAccount_Search_CustodyAccountId FangEngine_AttributeName_UUID_CA_Type = "5b64b7b0-d30b-4fd3-b742-9c4b95784dbe"
	FangEngine_AttributeName_Name_CA_CustodyAccount_Search_CustodyAccountId FangEngine_AttributeName_Name_CA_Type = "CustodyAccountId"

	// ClassName: ****** 'SettlementAgreement' ******
	FangEngine_ClassName_UUID_CA_SettlementAgreement FangEngine_ClassName_UUID_CA_Type = "8e48aefd-e847-4046-bc36-6749a24d7321"
	FangEngine_ClassName_Name_CA_SettlementAgreement FangEngine_ClassName_Name_CA_Type = "SettlementAgreement"

	// ClassName: 'SettlementAgreement' - MethodName: 'Edit'
	FangEngine_MethodName_UUID_CA_SettlementAgreement_Edit FangEngine_MethodName_UUID_CA_Type = "bb8af241-ba7c-4eba-87dc-bb1bf52be822"
	FangEngine_MethodName_Name_CA_SettlementAgreement_Edit FangEngine_MethodName_Name_CA_Type = "Edit"

	// ClassName: 'SettlementAgreement' - MethodName: 'Edit' - Attributes
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataDuplicateCheck        FangEngine_AttributeName_UUID_CA_Type = "0c94a266-2415-4765-9d13-33dfcfc81952"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataDuplicateCheck        FangEngine_AttributeName_Name_CA_Type = "TestDataDuplicateCheck"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataContractualRule       FangEngine_AttributeName_UUID_CA_Type = "732273f5-22d9-4fa4-9480-e149a087789f"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataContractualRule       FangEngine_AttributeName_Name_CA_Type = "TestDataContractualRule"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible   FangEngine_AttributeName_UUID_CA_Type = "99417b13-b943-4007-816e-a628dde293dd"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataMatchAffirmEligible   FangEngine_AttributeName_Name_CA_Type = "TestDataMatchAffirmEligible"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd       FangEngine_AttributeName_UUID_CA_Type = "ee811e9b-6129-48fa-b1ef-ee40b1b5c371"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataDirectAffirmInd       FangEngine_AttributeName_Name_CA_Type = "TestDataDirectAffirmInd"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck  FangEngine_AttributeName_UUID_CA_Type = "11bb4baa-305e-4136-a9a0-6e1a86199627"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataAffirmDublicateCheck  FangEngine_AttributeName_Name_CA_Type = "TestDataAffirmDublicateCheck"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd    FangEngine_AttributeName_UUID_CA_Type = "ef69e703-38fa-4ad0-8d2d-97b37252154e"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataAffirmCancelledInd    FangEngine_AttributeName_Name_CA_Type = "TestDataAffirmCancelledInd"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataNearMatch             FangEngine_AttributeName_UUID_CA_Type = "5a57b7d0-89eb-40b7-9bed-66c4572481a7"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataNearMatch             FangEngine_AttributeName_Name_CA_Type = "TestDataNearMatch"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService FangEngine_AttributeName_UUID_CA_Type = "c68c40cd-1431-470b-b63d-1bc86ccce2c6"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_Edit_TestDataReturnDeliveryService FangEngine_AttributeName_Name_CA_Type = "TestDataReturnDeliveryService"

	// ClassName:'SettlementAgreement' - MethodName: 'AddSelectedSwift'
	FangEngine_MethodName_UUID_CA_SettlementAgreement_AddSelectedSwift FangEngine_MethodName_UUID_CA_Type = "15474aa0-42b0-4f2a-93da-eb8243ccaec3"
	FangEngine_MethodName_Name_CA_SettlementAgreement_AddSelectedSwift FangEngine_MethodName_Name_CA_Type = "AddSelectedSwift"

	// ClassName:'SettlementAgreement' - MethodName: 'AddSelectedSwift' - Attributes
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataType     FangEngine_AttributeName_UUID_CA_Type = "db66ad38-ea09-4b15-af00-9bc18b8bef1d"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataType     FangEngine_AttributeName_Name_CA_Type = "TestDataType"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel  FangEngine_AttributeName_UUID_CA_Type = "e1219a09-a5d6-4d66-810a-46f3818edc41"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataChannel  FangEngine_AttributeName_Name_CA_Type = "TestDataChannel"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataBic      FangEngine_AttributeName_UUID_CA_Type = "5ede7885-2e2b-4e7b-852e-9c83f1373312"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataBic      FangEngine_AttributeName_Name_CA_Type = "TestDataBic"
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval FangEngine_AttributeName_UUID_CA_Type = "c8154d0c-ca2d-42cd-964e-1162e37a41b4"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedSwift_TestDataInterval FangEngine_AttributeName_Name_CA_Type = "TestDataInterval"

	// ClassName:'SettlementAgreement' - MethodName: 'DeleteSelectedSwift'
	FangEngine_MethodName_UUID_CA_SettlementAgreement_DeleteSelectedSwift FangEngine_MethodName_UUID_CA_Type = "771f6735-1fad-4f67-917a-a468936a59bc"
	FangEngine_MethodName_Name_CA_SettlementAgreement_DeleteSelectedSwift FangEngine_MethodName_Name_CA_Type = "DeleteSelectedSwift"

	// ClassName:'SettlementAgreement' - MethodName: 'DeleteSelectedSwift' - Attributes
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress FangEngine_AttributeName_UUID_CA_Type = "1d319ccf-5584-45a3-862a-1bbc57f83e32"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_DeleteSelectedSwift_bicAddress FangEngine_AttributeName_Name_CA_Type = "bicAddress"

	// ClassName:'SettlementAgreement' - MethodName: 'AddSelectedInstructedParties'
	FangEngine_MethodName_UUID_CA_SettlementAgreement_AddSelectedInstructedParties FangEngine_MethodName_UUID_CA_Type = "17629ae9-db7b-40e6-ab3d-7d8646b93a09"
	FangEngine_MethodName_Name_CA_SettlementAgreement_AddSelectedInstructedParties FangEngine_MethodName_Name_CA_Type = "AddSelectedInstructedParties"

	// ClassName:'SettlementAgreement' - MethodName: 'AddSelectedInstructedParties' - Attributes
	FangEngine_AttributeName_UUID_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress FangEngine_AttributeName_UUID_CA_Type = "7855d6fe-aaad-4b32-b168-8af5c5b1de3b"
	FangEngine_AttributeName_Name_CA_SettlementAgreement_AddSelectedInstructedParties_bicAddress FangEngine_AttributeName_Name_CA_Type = "bicAddress"

	// ClassName:'SettlementAgreement' - MethodName: 'DeleteSelectedInstructedParties'
	FangEngine_MethodName_UUID_CA_SettlementAgreement_DeleteSelectedInstructedParties FangEngine_MethodName_UUID_CA_Type = "656493c9-eb59-4c55-a5a1-f74ed9df91b2"
	FangEngine_MethodName_Name_CA_SettlementAgreement_DeleteSelectedInstructedParties FangEngine_MethodName_Name_CA_Type = "DeleteSelectedInstructedParties"

	// ClassName:'SettlementAgreement' - MethodName: 'DeleteSelectedInstructedParties' - Attributes
	// No attributes
)

/*


 */
