// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // goParser
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

	// EnterIdListDTSVDAST is called when entering the idListDTSVDAST production.
	EnterIdListDTSVDAST(c *IdListDTSVDASTContext)

	// EnterIdListSVDAST is called when entering the idListSVDAST production.
	EnterIdListSVDAST(c *IdListSVDASTContext)

	// EnterSingleVarDAST is called when entering the singleVarDAST production.
	EnterSingleVarDAST(c *SingleVarDASTContext)

	// EnterSingleVarDeclNoExpsAST is called when entering the singleVarDeclNoExpsAST production.
	EnterSingleVarDeclNoExpsAST(c *SingleVarDeclNoExpsASTContext)

	// EnterTypeTDAST is called when entering the typeTDAST production.
	EnterTypeTDAST(c *TypeTDASTContext)

	// EnterInnerTypeDeclsAST is called when entering the innerTypeDeclsAST production.
	EnterInnerTypeDeclsAST(c *InnerTypeDeclsASTContext)

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

	// EnterStructDeclTypeDTAST is called when entering the structDeclTypeDTAST production.
	EnterStructDeclTypeDTAST(c *StructDeclTypeDTASTContext)

	// EnterLsbSDTAST is called when entering the lsbSDTAST production.
	EnterLsbSDTAST(c *LsbSDTASTContext)

	// EnterLsbADTAST is called when entering the lsbADTAST production.
	EnterLsbADTAST(c *LsbADTASTContext)

	// EnterStructSDTAST is called when entering the structSDTAST production.
	EnterStructSDTAST(c *StructSDTASTContext)

	// EnterStructMemDeclsAST is called when entering the structMemDeclsAST production.
	EnterStructMemDeclsAST(c *StructMemDeclsASTContext)

	// EnterIdentifierListAST is called when entering the identifierListAST production.
	EnterIdentifierListAST(c *IdentifierListASTContext)

	// EnterExpressionEAST is called when entering the expressionEAST production.
	EnterExpressionEAST(c *ExpressionEASTContext)

	// EnterPrimaryExpressionEAST is called when entering the primaryExpressionEAST production.
	EnterPrimaryExpressionEAST(c *PrimaryExpressionEASTContext)

	// EnterOperatorEAST is called when entering the operatorEAST production.
	EnterOperatorEAST(c *OperatorEASTContext)

	// EnterExpressionListAST is called when entering the expressionListAST production.
	EnterExpressionListAST(c *ExpressionListASTContext)

	// EnterLengthExpPEAST is called when entering the lengthExpPEAST production.
	EnterLengthExpPEAST(c *LengthExpPEASTContext)

	// EnterAppendExpPEAST is called when entering the appendExpPEAST production.
	EnterAppendExpPEAST(c *AppendExpPEASTContext)

	// EnterOperandPEAST is called when entering the operandPEAST production.
	EnterOperandPEAST(c *OperandPEASTContext)

	// EnterPrimaryExpPEAST is called when entering the primaryExpPEAST production.
	EnterPrimaryExpPEAST(c *PrimaryExpPEASTContext)

	// EnterCapExpPEAST is called when entering the capExpPEAST production.
	EnterCapExpPEAST(c *CapExpPEASTContext)

	// EnterLiteralOAST is called when entering the literalOAST production.
	EnterLiteralOAST(c *LiteralOASTContext)

	// EnterIdOAST is called when entering the idOAST production.
	EnterIdOAST(c *IdOASTContext)

	// EnterExpressionOAST is called when entering the expressionOAST production.
	EnterExpressionOAST(c *ExpressionOASTContext)

	// EnterIntliteralAST is called when entering the intliteralAST production.
	EnterIntliteralAST(c *IntliteralASTContext)

	// EnterFloatliteralAST is called when entering the floatliteralAST production.
	EnterFloatliteralAST(c *FloatliteralASTContext)

	// EnterRunliteralAST is called when entering the runliteralAST production.
	EnterRunliteralAST(c *RunliteralASTContext)

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

	// EnterAssigmentStatementSSAST is called when entering the assigmentStatementSSAST production.
	EnterAssigmentStatementSSAST(c *AssigmentStatementSSASTContext)

	// EnterExpListSSAST is called when entering the expListSSAST production.
	EnterExpListSSAST(c *ExpListSSASTContext)

	// EnterExpListASAST is called when entering the expListASAST production.
	EnterExpListASAST(c *ExpListASASTContext)

	// EnterExpASAST is called when entering the expASAST production.
	EnterExpASAST(c *ExpASASTContext)

	// EnterIsExpressionBlockISAST is called when entering the isExpressionBlockISAST production.
	EnterIsExpressionBlockISAST(c *IsExpressionBlockISASTContext)

	// EnterIsExpBlockIsISAST is called when entering the isExpBlockIsISAST production.
	EnterIsExpBlockIsISAST(c *IsExpBlockIsISASTContext)

	// EnterIsExpBlockISAST is called when entering the isExpBlockISAST production.
	EnterIsExpBlockISAST(c *IsExpBlockISASTContext)

	// EnterIsSSExpBlockISAST is called when entering the isSSExpBlockISAST production.
	EnterIsSSExpBlockISAST(c *IsSSExpBlockISASTContext)

	// EnterIsSSExpBlockifSAST is called when entering the isSSExpBlockifSAST production.
	EnterIsSSExpBlockifSAST(c *IsSSExpBlockifSASTContext)

	// EnterIsSSExpBlockBlockAST is called when entering the isSSExpBlockBlockAST production.
	EnterIsSSExpBlockBlockAST(c *IsSSExpBlockBlockASTContext)

	// EnterFBlockLAST is called when entering the fBlockLAST production.
	EnterFBlockLAST(c *FBlockLASTContext)

	// EnterFExpBlockLAST is called when entering the fExpBlockLAST production.
	EnterFExpBlockLAST(c *FExpBlockLASTContext)

	// EnterFSimpStateExpSBlockLAST is called when entering the fSimpStateExpSBlockLAST production.
	EnterFSimpStateExpSBlockLAST(c *FSimpStateExpSBlockLASTContext)

	// EnterFSimpStateSimpStateBlockLAST is called when entering the fSimpStateSimpStateBlockLAST production.
	EnterFSimpStateSimpStateBlockLAST(c *FSimpStateSimpStateBlockLASTContext)

	// EnterSSimpleStatSAST is called when entering the sSimpleStatSAST production.
	EnterSSimpleStatSAST(c *SSimpleStatSASTContext)

	// EnterSExpressionSAST is called when entering the sExpressionSAST production.
	EnterSExpressionSAST(c *SExpressionSASTContext)

	// EnterSSimpleStatExpCCListSAST is called when entering the sSimpleStatExpCCListSAST production.
	EnterSSimpleStatExpCCListSAST(c *SSimpleStatExpCCListSASTContext)

	// EnterSBlockSAST is called when entering the sBlockSAST production.
	EnterSBlockSAST(c *SBlockSASTContext)

	// EnterEpsilonECCLAST is called when entering the epsilonECCLAST production.
	EnterEpsilonECCLAST(c *EpsilonECCLASTContext)

	// EnterExpCaseClauseECCLAST is called when entering the expCaseClauseECCLAST production.
	EnterExpCaseClauseECCLAST(c *ExpCaseClauseECCLASTContext)

	// EnterExpSwitchCaseECCAST is called when entering the expSwitchCaseECCAST production.
	EnterExpSwitchCaseECCAST(c *ExpSwitchCaseECCASTContext)

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

	// ExitIdListDTSVDAST is called when exiting the idListDTSVDAST production.
	ExitIdListDTSVDAST(c *IdListDTSVDASTContext)

	// ExitIdListSVDAST is called when exiting the idListSVDAST production.
	ExitIdListSVDAST(c *IdListSVDASTContext)

	// ExitSingleVarDAST is called when exiting the singleVarDAST production.
	ExitSingleVarDAST(c *SingleVarDASTContext)

	// ExitSingleVarDeclNoExpsAST is called when exiting the singleVarDeclNoExpsAST production.
	ExitSingleVarDeclNoExpsAST(c *SingleVarDeclNoExpsASTContext)

	// ExitTypeTDAST is called when exiting the typeTDAST production.
	ExitTypeTDAST(c *TypeTDASTContext)

	// ExitInnerTypeDeclsAST is called when exiting the innerTypeDeclsAST production.
	ExitInnerTypeDeclsAST(c *InnerTypeDeclsASTContext)

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

	// ExitStructDeclTypeDTAST is called when exiting the structDeclTypeDTAST production.
	ExitStructDeclTypeDTAST(c *StructDeclTypeDTASTContext)

	// ExitLsbSDTAST is called when exiting the lsbSDTAST production.
	ExitLsbSDTAST(c *LsbSDTASTContext)

	// ExitLsbADTAST is called when exiting the lsbADTAST production.
	ExitLsbADTAST(c *LsbADTASTContext)

	// ExitStructSDTAST is called when exiting the structSDTAST production.
	ExitStructSDTAST(c *StructSDTASTContext)

	// ExitStructMemDeclsAST is called when exiting the structMemDeclsAST production.
	ExitStructMemDeclsAST(c *StructMemDeclsASTContext)

	// ExitIdentifierListAST is called when exiting the identifierListAST production.
	ExitIdentifierListAST(c *IdentifierListASTContext)

	// ExitExpressionEAST is called when exiting the expressionEAST production.
	ExitExpressionEAST(c *ExpressionEASTContext)

	// ExitPrimaryExpressionEAST is called when exiting the primaryExpressionEAST production.
	ExitPrimaryExpressionEAST(c *PrimaryExpressionEASTContext)

	// ExitOperatorEAST is called when exiting the operatorEAST production.
	ExitOperatorEAST(c *OperatorEASTContext)

	// ExitExpressionListAST is called when exiting the expressionListAST production.
	ExitExpressionListAST(c *ExpressionListASTContext)

	// ExitLengthExpPEAST is called when exiting the lengthExpPEAST production.
	ExitLengthExpPEAST(c *LengthExpPEASTContext)

	// ExitAppendExpPEAST is called when exiting the appendExpPEAST production.
	ExitAppendExpPEAST(c *AppendExpPEASTContext)

	// ExitOperandPEAST is called when exiting the operandPEAST production.
	ExitOperandPEAST(c *OperandPEASTContext)

	// ExitPrimaryExpPEAST is called when exiting the primaryExpPEAST production.
	ExitPrimaryExpPEAST(c *PrimaryExpPEASTContext)

	// ExitCapExpPEAST is called when exiting the capExpPEAST production.
	ExitCapExpPEAST(c *CapExpPEASTContext)

	// ExitLiteralOAST is called when exiting the literalOAST production.
	ExitLiteralOAST(c *LiteralOASTContext)

	// ExitIdOAST is called when exiting the idOAST production.
	ExitIdOAST(c *IdOASTContext)

	// ExitExpressionOAST is called when exiting the expressionOAST production.
	ExitExpressionOAST(c *ExpressionOASTContext)

	// ExitIntliteralAST is called when exiting the intliteralAST production.
	ExitIntliteralAST(c *IntliteralASTContext)

	// ExitFloatliteralAST is called when exiting the floatliteralAST production.
	ExitFloatliteralAST(c *FloatliteralASTContext)

	// ExitRunliteralAST is called when exiting the runliteralAST production.
	ExitRunliteralAST(c *RunliteralASTContext)

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

	// ExitAssigmentStatementSSAST is called when exiting the assigmentStatementSSAST production.
	ExitAssigmentStatementSSAST(c *AssigmentStatementSSASTContext)

	// ExitExpListSSAST is called when exiting the expListSSAST production.
	ExitExpListSSAST(c *ExpListSSASTContext)

	// ExitExpListASAST is called when exiting the expListASAST production.
	ExitExpListASAST(c *ExpListASASTContext)

	// ExitExpASAST is called when exiting the expASAST production.
	ExitExpASAST(c *ExpASASTContext)

	// ExitIsExpressionBlockISAST is called when exiting the isExpressionBlockISAST production.
	ExitIsExpressionBlockISAST(c *IsExpressionBlockISASTContext)

	// ExitIsExpBlockIsISAST is called when exiting the isExpBlockIsISAST production.
	ExitIsExpBlockIsISAST(c *IsExpBlockIsISASTContext)

	// ExitIsExpBlockISAST is called when exiting the isExpBlockISAST production.
	ExitIsExpBlockISAST(c *IsExpBlockISASTContext)

	// ExitIsSSExpBlockISAST is called when exiting the isSSExpBlockISAST production.
	ExitIsSSExpBlockISAST(c *IsSSExpBlockISASTContext)

	// ExitIsSSExpBlockifSAST is called when exiting the isSSExpBlockifSAST production.
	ExitIsSSExpBlockifSAST(c *IsSSExpBlockifSASTContext)

	// ExitIsSSExpBlockBlockAST is called when exiting the isSSExpBlockBlockAST production.
	ExitIsSSExpBlockBlockAST(c *IsSSExpBlockBlockASTContext)

	// ExitFBlockLAST is called when exiting the fBlockLAST production.
	ExitFBlockLAST(c *FBlockLASTContext)

	// ExitFExpBlockLAST is called when exiting the fExpBlockLAST production.
	ExitFExpBlockLAST(c *FExpBlockLASTContext)

	// ExitFSimpStateExpSBlockLAST is called when exiting the fSimpStateExpSBlockLAST production.
	ExitFSimpStateExpSBlockLAST(c *FSimpStateExpSBlockLASTContext)

	// ExitFSimpStateSimpStateBlockLAST is called when exiting the fSimpStateSimpStateBlockLAST production.
	ExitFSimpStateSimpStateBlockLAST(c *FSimpStateSimpStateBlockLASTContext)

	// ExitSSimpleStatSAST is called when exiting the sSimpleStatSAST production.
	ExitSSimpleStatSAST(c *SSimpleStatSASTContext)

	// ExitSExpressionSAST is called when exiting the sExpressionSAST production.
	ExitSExpressionSAST(c *SExpressionSASTContext)

	// ExitSSimpleStatExpCCListSAST is called when exiting the sSimpleStatExpCCListSAST production.
	ExitSSimpleStatExpCCListSAST(c *SSimpleStatExpCCListSASTContext)

	// ExitSBlockSAST is called when exiting the sBlockSAST production.
	ExitSBlockSAST(c *SBlockSASTContext)

	// ExitEpsilonECCLAST is called when exiting the epsilonECCLAST production.
	ExitEpsilonECCLAST(c *EpsilonECCLASTContext)

	// ExitExpCaseClauseECCLAST is called when exiting the expCaseClauseECCLAST production.
	ExitExpCaseClauseECCLAST(c *ExpCaseClauseECCLASTContext)

	// ExitExpSwitchCaseECCAST is called when exiting the expSwitchCaseECCAST production.
	ExitExpSwitchCaseECCAST(c *ExpSwitchCaseECCASTContext)

	// ExitCaseESCAST is called when exiting the caseESCAST production.
	ExitCaseESCAST(c *CaseESCASTContext)

	// ExitDefaultESCAST is called when exiting the defaultESCAST production.
	ExitDefaultESCAST(c *DefaultESCASTContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitEpsilonAST is called when exiting the epsilonAST production.
	ExitEpsilonAST(c *EpsilonASTContext)
}
