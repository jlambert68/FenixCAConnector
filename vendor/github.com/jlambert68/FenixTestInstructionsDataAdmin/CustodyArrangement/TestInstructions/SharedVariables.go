package TestInstructions

import (
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/CustodyArrangement/FangEngineClassesAndMethods"
	"github.com/jlambert68/FenixTestInstructionsDataAdmin/TypeAndStructs"
)

// TestInstruction_CA_TestCaseSetUpStruct
// Struct for holding all data for the TestInstruction
type TestInstruction_CA_TestCaseSetUpStruct struct {
	TestInstruction                    TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}

// AllTestInstructions_CA_TestCaseSetUpStruct
// Struct for holding all data for all TestInstructions
type AllTestInstructions_CA_TestCaseSetUpStruct struct {
	TestInstruction                    []TypeAndStructs.TestInstructionStruct
	BasicTestInstructionInformation    []TypeAndStructs.BasicTestInstructionInformationStruct
	ImmatureTestInstructionInformation []TypeAndStructs.ImmatureTestInstructionInformationStruct
	TestInstructionAttribute           []TypeAndStructs.TestInstructionAttributeStruct
	ImmatureElementModel               []TypeAndStructs.ImmatureElementModelMessageStruct
	FangEngineClassesMethodsAttributes []FangEngineClassesAndMethods.FangEngineClassesMethodsAttributesStruct
}
