// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/Compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Generated // goParser
import "github.com/antlr4-go/antlr/v4"

// goParserListener is a complete listener for a parse tree produced by goParser.
type goParserListener interface {
	antlr.ParseTreeListener

	// EnterRootAST is called when entering the rootAST production.
	EnterRootAST(c *RootASTContext)

	// EnterTopDeclarationListAST is called when entering the topDeclarationListAST production.
	EnterTopDeclarationListAST(c *TopDeclarationListASTContext)

	// EnterVarVDAST is called when entering the varVDAST production.
	EnterVarVDAST(c *VarVDASTContext)

	// EnterInnerVarDeclsAST is called when entering the innerVarDeclsAST production.
	EnterInnerVarDeclsAST(c *InnerVarDeclsASTContext)

	// EnterIdentifierListDTSVDAST is called when entering the identifierListDTSVDAST production.
	EnterIdentifierListDTSVDAST(c *IdentifierListDTSVDASTContext)

	// EnterIdentifierListSVDAST is called when entering the identifierListSVDAST production.
	EnterIdentifierListSVDAST(c *IdentifierListSVDASTContext)

	// EnterSingleVarDAST is called when entering the singleVarDAST production.
	EnterSingleVarDAST(c *SingleVarDASTContext)

	// EnterSingleVarDeclNoExpsAST is called when entering the singleVarDeclNoExpsAST production.
	EnterSingleVarDeclNoExpsAST(c *SingleVarDeclNoExpsASTContext)

	// EnterTypeTDAST is called when entering the typeTDAST production.
	EnterTypeTDAST(c *TypeTDASTContext)

	// EnterInnerTypeDeclsASt is called when entering the innerTypeDeclsASt production.
	EnterInnerTypeDeclsASt(c *InnerTypeDeclsAStContext)

	// EnterIdSYDAST is called when entering the idSYDAST production.
	EnterIdSYDAST(c *IdSYDASTContext)

	// EnterFuncDeclAST is called when entering the funcDeclAST production.
	EnterFuncDeclAST(c *FuncDeclASTContext)

	// EnterFuncFFDAST is called when entering the funcFFDAST production.
	EnterFuncFFDAST(c *FuncFFDASTContext)

	// EnterFuncArgDeclsAST is called when entering the funcArgDeclsAST production.
	EnterFuncArgDeclsAST(c *FuncArgDeclsASTContext)

	// EnterLpDTAST is called when entering the lpDTAST production.
	EnterLpDTAST(c *LpDTASTContext)

	// EnterIdDTAST is called when entering the idDTAST production.
	EnterIdDTAST(c *IdDTASTContext)

	// EnterSliceDeclTypeDTAST is called when entering the sliceDeclTypeDTAST production.
	EnterSliceDeclTypeDTAST(c *SliceDeclTypeDTASTContext)

	// EnterArrayDeclTypeDTAST is called when entering the arrayDeclTypeDTAST production.
	EnterArrayDeclTypeDTAST(c *ArrayDeclTypeDTASTContext)

	// EnterDeclTypeAST is called when entering the declTypeAST production.
	EnterDeclTypeAST(c *DeclTypeASTContext)

	// EnterLsbSDTAST is called when entering the lsbSDTAST production.
	EnterLsbSDTAST(c *LsbSDTASTContext)

	// EnterLsbADTAST is called when entering the lsbADTAST production.
	EnterLsbADTAST(c *LsbADTASTContext)

	// EnterStrctSDTAST is called when entering the strctSDTAST production.
	EnterStrctSDTAST(c *StrctSDTASTContext)

	// EnterStructMemDeclsAST is called when entering the structMemDeclsAST production.
	EnterStructMemDeclsAST(c *StructMemDeclsASTContext)

	// EnterIdentifierListAST is called when entering the identifierListAST production.
	EnterIdentifierListAST(c *IdentifierListASTContext)

	// EnterExpressionAST is called when entering the expressionAST production.
	EnterExpressionAST(c *ExpressionASTContext)

	// EnterExpressionEAST is called when entering the expressionEAST production.
	EnterExpressionEAST(c *ExpressionEASTContext)

	// EnterPrimaryExpressionEAST is called when entering the primaryExpressionEAST production.
	EnterPrimaryExpressionEAST(c *PrimaryExpressionEASTContext)

	// EnterExpressionListASt is called when entering the expressionListASt production.
	EnterExpressionListASt(c *ExpressionListAStContext)

	// EnterCapExpressionPEAST is called when entering the capExpressionPEAST production.
	EnterCapExpressionPEAST(c *CapExpressionPEASTContext)

	// EnterPrimaryExpressionPEAST is called when entering the primaryExpressionPEAST production.
	EnterPrimaryExpressionPEAST(c *PrimaryExpressionPEASTContext)

	// EnterLengthExpressionPEAST is called when entering the lengthExpressionPEAST production.
	EnterLengthExpressionPEAST(c *LengthExpressionPEASTContext)

	// EnterAppendExpressionPEAST is called when entering the appendExpressionPEAST production.
	EnterAppendExpressionPEAST(c *AppendExpressionPEASTContext)

	// EnterOperandPEAST is called when entering the operandPEAST production.
	EnterOperandPEAST(c *OperandPEASTContext)

	// EnterLiteralOAST is called when entering the literalOAST production.
	EnterLiteralOAST(c *LiteralOASTContext)

	// EnterIdOAST is called when entering the idOAST production.
	EnterIdOAST(c *IdOASTContext)

	// EnterLpOAST is called when entering the lpOAST production.
	EnterLpOAST(c *LpOASTContext)

	// EnterIntliteralAST is called when entering the intliteralAST production.
	EnterIntliteralAST(c *IntliteralASTContext)

	// EnterFloatliteralAST is called when entering the floatliteralAST production.
	EnterFloatliteralAST(c *FloatliteralASTContext)

	// EnterRunliteral is called when entering the runliteral production.
	EnterRunliteral(c *RunliteralContext)

	// EnterRawsliteralAST is called when entering the rawsliteralAST production.
	EnterRawsliteralAST(c *RawsliteralASTContext)

	// EnterInterpretedliteralAST is called when entering the interpretedliteralAST production.
	EnterInterpretedliteralAST(c *InterpretedliteralASTContext)

	// EnterIdexAST is called when entering the idexAST production.
	EnterIdexAST(c *IdexASTContext)

	// EnterArgumentsAST is called when entering the argumentsAST production.
	EnterArgumentsAST(c *ArgumentsASTContext)

	// EnterSelectorAST is called when entering the selectorAST production.
	EnterSelectorAST(c *SelectorASTContext)

	// EnterAppendExpressionAST is called when entering the appendExpressionAST production.
	EnterAppendExpressionAST(c *AppendExpressionASTContext)

	// EnterLengthExpressionAST is called when entering the lengthExpressionAST production.
	EnterLengthExpressionAST(c *LengthExpressionASTContext)

	// EnterCapExpressionAST is called when entering the capExpressionAST production.
	EnterCapExpressionAST(c *CapExpressionASTContext)

	// EnterStatementListAST is called when entering the statementListAST production.
	EnterStatementListAST(c *StatementListASTContext)

	// EnterBlockAST is called when entering the blockAST production.
	EnterBlockAST(c *BlockASTContext)

	// EnterPrintSAST is called when entering the printSAST production.
	EnterPrintSAST(c *PrintSASTContext)

	// EnterPrintlnSAST is called when entering the printlnSAST production.
	EnterPrintlnSAST(c *PrintlnSASTContext)

	// EnterReturnSAST is called when entering the returnSAST production.
	EnterReturnSAST(c *ReturnSASTContext)

	// EnterBreakSAST is called when entering the breakSAST production.
	EnterBreakSAST(c *BreakSASTContext)

	// EnterContinueSAST is called when entering the continueSAST production.
	EnterContinueSAST(c *ContinueSASTContext)

	// EnterSimpleStatementSAST is called when entering the simpleStatementSAST production.
	EnterSimpleStatementSAST(c *SimpleStatementSASTContext)

	// EnterBlockSAST is called when entering the blockSAST production.
	EnterBlockSAST(c *BlockSASTContext)

	// EnterSwitchSAST is called when entering the switchSAST production.
	EnterSwitchSAST(c *SwitchSASTContext)

	// EnterIfStatementSAST is called when entering the ifStatementSAST production.
	EnterIfStatementSAST(c *IfStatementSASTContext)

	// EnterLoopSAST is called when entering the loopSAST production.
	EnterLoopSAST(c *LoopSASTContext)

	// EnterTypeDeclSAST is called when entering the typeDeclSAST production.
	EnterTypeDeclSAST(c *TypeDeclSASTContext)

	// EnterVariableDeclSAST is called when entering the variableDeclSAST production.
	EnterVariableDeclSAST(c *VariableDeclSASTContext)

	// EnterEpsilonSSAST is called when entering the epsilonSSAST production.
	EnterEpsilonSSAST(c *EpsilonSSASTContext)

	// EnterExpressionSSAST is called when entering the expressionSSAST production.
	EnterExpressionSSAST(c *ExpressionSSASTContext)

	// EnterAssignmentStatementSS is called when entering the assignmentStatementSS production.
	EnterAssignmentStatementSS(c *AssignmentStatementSSContext)

	// EnterExpressionListSS is called when entering the expressionListSS production.
	EnterExpressionListSS(c *ExpressionListSSContext)

	// EnterExpressionListASAST is called when entering the expressionListASAST production.
	EnterExpressionListASAST(c *ExpressionListASASTContext)

	// EnterExpressionASAST is called when entering the expressionASAST production.
	EnterExpressionASAST(c *ExpressionASASTContext)

	// EnterIsExpressionBlockISAST is called when entering the isExpressionBlockISAST production.
	EnterIsExpressionBlockISAST(c *IsExpressionBlockISASTContext)

	// EnterIsExpressionBlockIsISAST is called when entering the isExpressionBlockIsISAST production.
	EnterIsExpressionBlockIsISAST(c *IsExpressionBlockIsISASTContext)

	// EnterIsSimpleStamentBlockISAST is called when entering the isSimpleStamentBlockISAST production.
	EnterIsSimpleStamentBlockISAST(c *IsSimpleStamentBlockISASTContext)

	// EnterIsSimpleStamentExpressionBlockISAST is called when entering the isSimpleStamentExpressionBlockISAST production.
	EnterIsSimpleStamentExpressionBlockISAST(c *IsSimpleStamentExpressionBlockISASTContext)

	// EnterIsSimpleStamentExpressionBlockifSAST is called when entering the isSimpleStamentExpressionBlockifSAST production.
	EnterIsSimpleStamentExpressionBlockifSAST(c *IsSimpleStamentExpressionBlockifSASTContext)

	// EnterIsSimpleStamentExpressionBlockBlockAST is called when entering the isSimpleStamentExpressionBlockBlockAST production.
	EnterIsSimpleStamentExpressionBlockBlockAST(c *IsSimpleStamentExpressionBlockBlockASTContext)

	// EnterFBlockLAST is called when entering the fBlockLAST production.
	EnterFBlockLAST(c *FBlockLASTContext)

	// EnterFExpressionBlockLAST is called when entering the fExpressionBlockLAST production.
	EnterFExpressionBlockLAST(c *FExpressionBlockLASTContext)

	// EnterFSimpleStatementLAST is called when entering the fSimpleStatementLAST production.
	EnterFSimpleStatementLAST(c *FSimpleStatementLASTContext)

	// EnterSSimpleStatSAST is called when entering the sSimpleStatSAST production.
	EnterSSimpleStatSAST(c *SSimpleStatSASTContext)

	// EnterSExpressionSAST is called when entering the sExpressionSAST production.
	EnterSExpressionSAST(c *SExpressionSASTContext)

	// EnterSBlockSAST is called when entering the sBlockSAST production.
	EnterSBlockSAST(c *SBlockSASTContext)

	// EnterEpsilonECCLAST is called when entering the epsilonECCLAST production.
	EnterEpsilonECCLAST(c *EpsilonECCLASTContext)

	// EnterExpressionCaseClauseECCLAST is called when entering the expressionCaseClauseECCLAST production.
	EnterExpressionCaseClauseECCLAST(c *ExpressionCaseClauseECCLASTContext)

	// EnterExpressionSwitchCaseECCAST is called when entering the expressionSwitchCaseECCAST production.
	EnterExpressionSwitchCaseECCAST(c *ExpressionSwitchCaseECCASTContext)

	// EnterCaseESCAST is called when entering the caseESCAST production.
	EnterCaseESCAST(c *CaseESCASTContext)

	// EnterDefaultESCAST is called when entering the defaultESCAST production.
	EnterDefaultESCAST(c *DefaultESCASTContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterEpsilonAST is called when entering the epsilonAST production.
	EnterEpsilonAST(c *EpsilonASTContext)

	// ExitRootAST is called when exiting the rootAST production.
	ExitRootAST(c *RootASTContext)

	// ExitTopDeclarationListAST is called when exiting the topDeclarationListAST production.
	ExitTopDeclarationListAST(c *TopDeclarationListASTContext)

	// ExitVarVDAST is called when exiting the varVDAST production.
	ExitVarVDAST(c *VarVDASTContext)

	// ExitInnerVarDeclsAST is called when exiting the innerVarDeclsAST production.
	ExitInnerVarDeclsAST(c *InnerVarDeclsASTContext)

	// ExitIdentifierListDTSVDAST is called when exiting the identifierListDTSVDAST production.
	ExitIdentifierListDTSVDAST(c *IdentifierListDTSVDASTContext)

	// ExitIdentifierListSVDAST is called when exiting the identifierListSVDAST production.
	ExitIdentifierListSVDAST(c *IdentifierListSVDASTContext)

	// ExitSingleVarDAST is called when exiting the singleVarDAST production.
	ExitSingleVarDAST(c *SingleVarDASTContext)

	// ExitSingleVarDeclNoExpsAST is called when exiting the singleVarDeclNoExpsAST production.
	ExitSingleVarDeclNoExpsAST(c *SingleVarDeclNoExpsASTContext)

	// ExitTypeTDAST is called when exiting the typeTDAST production.
	ExitTypeTDAST(c *TypeTDASTContext)

	// ExitInnerTypeDeclsASt is called when exiting the innerTypeDeclsASt production.
	ExitInnerTypeDeclsASt(c *InnerTypeDeclsAStContext)

	// ExitIdSYDAST is called when exiting the idSYDAST production.
	ExitIdSYDAST(c *IdSYDASTContext)

	// ExitFuncDeclAST is called when exiting the funcDeclAST production.
	ExitFuncDeclAST(c *FuncDeclASTContext)

	// ExitFuncFFDAST is called when exiting the funcFFDAST production.
	ExitFuncFFDAST(c *FuncFFDASTContext)

	// ExitFuncArgDeclsAST is called when exiting the funcArgDeclsAST production.
	ExitFuncArgDeclsAST(c *FuncArgDeclsASTContext)

	// ExitLpDTAST is called when exiting the lpDTAST production.
	ExitLpDTAST(c *LpDTASTContext)

	// ExitIdDTAST is called when exiting the idDTAST production.
	ExitIdDTAST(c *IdDTASTContext)

	// ExitSliceDeclTypeDTAST is called when exiting the sliceDeclTypeDTAST production.
	ExitSliceDeclTypeDTAST(c *SliceDeclTypeDTASTContext)

	// ExitArrayDeclTypeDTAST is called when exiting the arrayDeclTypeDTAST production.
	ExitArrayDeclTypeDTAST(c *ArrayDeclTypeDTASTContext)

	// ExitDeclTypeAST is called when exiting the declTypeAST production.
	ExitDeclTypeAST(c *DeclTypeASTContext)

	// ExitLsbSDTAST is called when exiting the lsbSDTAST production.
	ExitLsbSDTAST(c *LsbSDTASTContext)

	// ExitLsbADTAST is called when exiting the lsbADTAST production.
	ExitLsbADTAST(c *LsbADTASTContext)

	// ExitStrctSDTAST is called when exiting the strctSDTAST production.
	ExitStrctSDTAST(c *StrctSDTASTContext)

	// ExitStructMemDeclsAST is called when exiting the structMemDeclsAST production.
	ExitStructMemDeclsAST(c *StructMemDeclsASTContext)

	// ExitIdentifierListAST is called when exiting the identifierListAST production.
	ExitIdentifierListAST(c *IdentifierListASTContext)

	// ExitExpressionAST is called when exiting the expressionAST production.
	ExitExpressionAST(c *ExpressionASTContext)

	// ExitExpressionEAST is called when exiting the expressionEAST production.
	ExitExpressionEAST(c *ExpressionEASTContext)

	// ExitPrimaryExpressionEAST is called when exiting the primaryExpressionEAST production.
	ExitPrimaryExpressionEAST(c *PrimaryExpressionEASTContext)

	// ExitExpressionListASt is called when exiting the expressionListASt production.
	ExitExpressionListASt(c *ExpressionListAStContext)

	// ExitCapExpressionPEAST is called when exiting the capExpressionPEAST production.
	ExitCapExpressionPEAST(c *CapExpressionPEASTContext)

	// ExitPrimaryExpressionPEAST is called when exiting the primaryExpressionPEAST production.
	ExitPrimaryExpressionPEAST(c *PrimaryExpressionPEASTContext)

	// ExitLengthExpressionPEAST is called when exiting the lengthExpressionPEAST production.
	ExitLengthExpressionPEAST(c *LengthExpressionPEASTContext)

	// ExitAppendExpressionPEAST is called when exiting the appendExpressionPEAST production.
	ExitAppendExpressionPEAST(c *AppendExpressionPEASTContext)

	// ExitOperandPEAST is called when exiting the operandPEAST production.
	ExitOperandPEAST(c *OperandPEASTContext)

	// ExitLiteralOAST is called when exiting the literalOAST production.
	ExitLiteralOAST(c *LiteralOASTContext)

	// ExitIdOAST is called when exiting the idOAST production.
	ExitIdOAST(c *IdOASTContext)

	// ExitLpOAST is called when exiting the lpOAST production.
	ExitLpOAST(c *LpOASTContext)

	// ExitIntliteralAST is called when exiting the intliteralAST production.
	ExitIntliteralAST(c *IntliteralASTContext)

	// ExitFloatliteralAST is called when exiting the floatliteralAST production.
	ExitFloatliteralAST(c *FloatliteralASTContext)

	// ExitRunliteral is called when exiting the runliteral production.
	ExitRunliteral(c *RunliteralContext)

	// ExitRawsliteralAST is called when exiting the rawsliteralAST production.
	ExitRawsliteralAST(c *RawsliteralASTContext)

	// ExitInterpretedliteralAST is called when exiting the interpretedliteralAST production.
	ExitInterpretedliteralAST(c *InterpretedliteralASTContext)

	// ExitIdexAST is called when exiting the idexAST production.
	ExitIdexAST(c *IdexASTContext)

	// ExitArgumentsAST is called when exiting the argumentsAST production.
	ExitArgumentsAST(c *ArgumentsASTContext)

	// ExitSelectorAST is called when exiting the selectorAST production.
	ExitSelectorAST(c *SelectorASTContext)

	// ExitAppendExpressionAST is called when exiting the appendExpressionAST production.
	ExitAppendExpressionAST(c *AppendExpressionASTContext)

	// ExitLengthExpressionAST is called when exiting the lengthExpressionAST production.
	ExitLengthExpressionAST(c *LengthExpressionASTContext)

	// ExitCapExpressionAST is called when exiting the capExpressionAST production.
	ExitCapExpressionAST(c *CapExpressionASTContext)

	// ExitStatementListAST is called when exiting the statementListAST production.
	ExitStatementListAST(c *StatementListASTContext)

	// ExitBlockAST is called when exiting the blockAST production.
	ExitBlockAST(c *BlockASTContext)

	// ExitPrintSAST is called when exiting the printSAST production.
	ExitPrintSAST(c *PrintSASTContext)

	// ExitPrintlnSAST is called when exiting the printlnSAST production.
	ExitPrintlnSAST(c *PrintlnSASTContext)

	// ExitReturnSAST is called when exiting the returnSAST production.
	ExitReturnSAST(c *ReturnSASTContext)

	// ExitBreakSAST is called when exiting the breakSAST production.
	ExitBreakSAST(c *BreakSASTContext)

	// ExitContinueSAST is called when exiting the continueSAST production.
	ExitContinueSAST(c *ContinueSASTContext)

	// ExitSimpleStatementSAST is called when exiting the simpleStatementSAST production.
	ExitSimpleStatementSAST(c *SimpleStatementSASTContext)

	// ExitBlockSAST is called when exiting the blockSAST production.
	ExitBlockSAST(c *BlockSASTContext)

	// ExitSwitchSAST is called when exiting the switchSAST production.
	ExitSwitchSAST(c *SwitchSASTContext)

	// ExitIfStatementSAST is called when exiting the ifStatementSAST production.
	ExitIfStatementSAST(c *IfStatementSASTContext)

	// ExitLoopSAST is called when exiting the loopSAST production.
	ExitLoopSAST(c *LoopSASTContext)

	// ExitTypeDeclSAST is called when exiting the typeDeclSAST production.
	ExitTypeDeclSAST(c *TypeDeclSASTContext)

	// ExitVariableDeclSAST is called when exiting the variableDeclSAST production.
	ExitVariableDeclSAST(c *VariableDeclSASTContext)

	// ExitEpsilonSSAST is called when exiting the epsilonSSAST production.
	ExitEpsilonSSAST(c *EpsilonSSASTContext)

	// ExitExpressionSSAST is called when exiting the expressionSSAST production.
	ExitExpressionSSAST(c *ExpressionSSASTContext)

	// ExitAssignmentStatementSS is called when exiting the assignmentStatementSS production.
	ExitAssignmentStatementSS(c *AssignmentStatementSSContext)

	// ExitExpressionListSS is called when exiting the expressionListSS production.
	ExitExpressionListSS(c *ExpressionListSSContext)

	// ExitExpressionListASAST is called when exiting the expressionListASAST production.
	ExitExpressionListASAST(c *ExpressionListASASTContext)

	// ExitExpressionASAST is called when exiting the expressionASAST production.
	ExitExpressionASAST(c *ExpressionASASTContext)

	// ExitIsExpressionBlockISAST is called when exiting the isExpressionBlockISAST production.
	ExitIsExpressionBlockISAST(c *IsExpressionBlockISASTContext)

	// ExitIsExpressionBlockIsISAST is called when exiting the isExpressionBlockIsISAST production.
	ExitIsExpressionBlockIsISAST(c *IsExpressionBlockIsISASTContext)

	// ExitIsSimpleStamentBlockISAST is called when exiting the isSimpleStamentBlockISAST production.
	ExitIsSimpleStamentBlockISAST(c *IsSimpleStamentBlockISASTContext)

	// ExitIsSimpleStamentExpressionBlockISAST is called when exiting the isSimpleStamentExpressionBlockISAST production.
	ExitIsSimpleStamentExpressionBlockISAST(c *IsSimpleStamentExpressionBlockISASTContext)

	// ExitIsSimpleStamentExpressionBlockifSAST is called when exiting the isSimpleStamentExpressionBlockifSAST production.
	ExitIsSimpleStamentExpressionBlockifSAST(c *IsSimpleStamentExpressionBlockifSASTContext)

	// ExitIsSimpleStamentExpressionBlockBlockAST is called when exiting the isSimpleStamentExpressionBlockBlockAST production.
	ExitIsSimpleStamentExpressionBlockBlockAST(c *IsSimpleStamentExpressionBlockBlockASTContext)

	// ExitFBlockLAST is called when exiting the fBlockLAST production.
	ExitFBlockLAST(c *FBlockLASTContext)

	// ExitFExpressionBlockLAST is called when exiting the fExpressionBlockLAST production.
	ExitFExpressionBlockLAST(c *FExpressionBlockLASTContext)

	// ExitFSimpleStatementLAST is called when exiting the fSimpleStatementLAST production.
	ExitFSimpleStatementLAST(c *FSimpleStatementLASTContext)

	// ExitSSimpleStatSAST is called when exiting the sSimpleStatSAST production.
	ExitSSimpleStatSAST(c *SSimpleStatSASTContext)

	// ExitSExpressionSAST is called when exiting the sExpressionSAST production.
	ExitSExpressionSAST(c *SExpressionSASTContext)

	// ExitSBlockSAST is called when exiting the sBlockSAST production.
	ExitSBlockSAST(c *SBlockSASTContext)

	// ExitEpsilonECCLAST is called when exiting the epsilonECCLAST production.
	ExitEpsilonECCLAST(c *EpsilonECCLASTContext)

	// ExitExpressionCaseClauseECCLAST is called when exiting the expressionCaseClauseECCLAST production.
	ExitExpressionCaseClauseECCLAST(c *ExpressionCaseClauseECCLASTContext)

	// ExitExpressionSwitchCaseECCAST is called when exiting the expressionSwitchCaseECCAST production.
	ExitExpressionSwitchCaseECCAST(c *ExpressionSwitchCaseECCASTContext)

	// ExitCaseESCAST is called when exiting the caseESCAST production.
	ExitCaseESCAST(c *CaseESCASTContext)

	// ExitDefaultESCAST is called when exiting the defaultESCAST production.
	ExitDefaultESCAST(c *DefaultESCASTContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitEpsilonAST is called when exiting the epsilonAST production.
	ExitEpsilonAST(c *EpsilonASTContext)
}
