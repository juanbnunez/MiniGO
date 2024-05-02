// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // goParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by goParser.
type goParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by goParser#rootAST.
	VisitRootAST(ctx *RootASTContext) interface{}

	// Visit a parse tree produced by goParser#topDeclarationListAST.
	VisitTopDeclarationListAST(ctx *TopDeclarationListASTContext) interface{}

	// Visit a parse tree produced by goParser#singleVarVDAST.
	VisitSingleVarVDAST(ctx *SingleVarVDASTContext) interface{}

	// Visit a parse tree produced by goParser#innerVarVDAST.
	VisitInnerVarVDAST(ctx *InnerVarVDASTContext) interface{}

	// Visit a parse tree produced by goParser#lpVDAST.
	VisitLpVDAST(ctx *LpVDASTContext) interface{}

	// Visit a parse tree produced by goParser#innerVarDeclsAST.
	VisitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) interface{}

	// Visit a parse tree produced by goParser#idListDTSVDAST.
	VisitIdListDTSVDAST(ctx *IdListDTSVDASTContext) interface{}

	// Visit a parse tree produced by goParser#idListSVDAST.
	VisitIdListSVDAST(ctx *IdListSVDASTContext) interface{}

	// Visit a parse tree produced by goParser#singleVarDAST.
	VisitSingleVarDAST(ctx *SingleVarDASTContext) interface{}

	// Visit a parse tree produced by goParser#singleVarDeclNoExpsAST.
	VisitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) interface{}

	// Visit a parse tree produced by goParser#singleTypeDeclTDAST.
	VisitSingleTypeDeclTDAST(ctx *SingleTypeDeclTDASTContext) interface{}

	// Visit a parse tree produced by goParser#innerTypeDeclsTDAST.
	VisitInnerTypeDeclsTDAST(ctx *InnerTypeDeclsTDASTContext) interface{}

	// Visit a parse tree produced by goParser#lPTypeDeclTDAST.
	VisitLPTypeDeclTDAST(ctx *LPTypeDeclTDASTContext) interface{}

	// Visit a parse tree produced by goParser#innerTypeDeclsAST.
	VisitInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) interface{}

	// Visit a parse tree produced by goParser#idSYDAST.
	VisitIdSYDAST(ctx *IdSYDASTContext) interface{}

	// Visit a parse tree produced by goParser#funcDeclAST.
	VisitFuncDeclAST(ctx *FuncDeclASTContext) interface{}

	// Visit a parse tree produced by goParser#funcFFDAST.
	VisitFuncFFDAST(ctx *FuncFFDASTContext) interface{}

	// Visit a parse tree produced by goParser#funcArgDeclsAST.
	VisitFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) interface{}

	// Visit a parse tree produced by goParser#lpDTAST.
	VisitLpDTAST(ctx *LpDTASTContext) interface{}

	// Visit a parse tree produced by goParser#idDTAST.
	VisitIdDTAST(ctx *IdDTASTContext) interface{}

	// Visit a parse tree produced by goParser#sliceDeclTypeDTAST.
	VisitSliceDeclTypeDTAST(ctx *SliceDeclTypeDTASTContext) interface{}

	// Visit a parse tree produced by goParser#arrayDeclTypeDTAST.
	VisitArrayDeclTypeDTAST(ctx *ArrayDeclTypeDTASTContext) interface{}

	// Visit a parse tree produced by goParser#structDeclTypeDTAST.
	VisitStructDeclTypeDTAST(ctx *StructDeclTypeDTASTContext) interface{}

	// Visit a parse tree produced by goParser#lsbSDTAST.
	VisitLsbSDTAST(ctx *LsbSDTASTContext) interface{}

	// Visit a parse tree produced by goParser#lsbADTAST.
	VisitLsbADTAST(ctx *LsbADTASTContext) interface{}

	// Visit a parse tree produced by goParser#structSDTAST.
	VisitStructSDTAST(ctx *StructSDTASTContext) interface{}

	// Visit a parse tree produced by goParser#structMemDeclsAST.
	VisitStructMemDeclsAST(ctx *StructMemDeclsASTContext) interface{}

	// Visit a parse tree produced by goParser#identifierListAST.
	VisitIdentifierListAST(ctx *IdentifierListASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionEAST.
	VisitExpressionEAST(ctx *ExpressionEASTContext) interface{}

	// Visit a parse tree produced by goParser#primaryExpressionEAST.
	VisitPrimaryExpressionEAST(ctx *PrimaryExpressionEASTContext) interface{}

	// Visit a parse tree produced by goParser#operatorEAST.
	VisitOperatorEAST(ctx *OperatorEASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionListAST.
	VisitExpressionListAST(ctx *ExpressionListASTContext) interface{}

	// Visit a parse tree produced by goParser#primaryExpArgumentsPEAST.
	VisitPrimaryExpArgumentsPEAST(ctx *PrimaryExpArgumentsPEASTContext) interface{}

	// Visit a parse tree produced by goParser#primaryExpSelectorPEAST.
	VisitPrimaryExpSelectorPEAST(ctx *PrimaryExpSelectorPEASTContext) interface{}

	// Visit a parse tree produced by goParser#lengthExpPEAST.
	VisitLengthExpPEAST(ctx *LengthExpPEASTContext) interface{}

	// Visit a parse tree produced by goParser#appendExpPEAST.
	VisitAppendExpPEAST(ctx *AppendExpPEASTContext) interface{}

	// Visit a parse tree produced by goParser#operandPEAST.
	VisitOperandPEAST(ctx *OperandPEASTContext) interface{}

	// Visit a parse tree produced by goParser#capExpPEAST.
	VisitCapExpPEAST(ctx *CapExpPEASTContext) interface{}

	// Visit a parse tree produced by goParser#primaryExpIndexPEAST.
	VisitPrimaryExpIndexPEAST(ctx *PrimaryExpIndexPEASTContext) interface{}

	// Visit a parse tree produced by goParser#literalOAST.
	VisitLiteralOAST(ctx *LiteralOASTContext) interface{}

	// Visit a parse tree produced by goParser#idOAST.
	VisitIdOAST(ctx *IdOASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionOAST.
	VisitExpressionOAST(ctx *ExpressionOASTContext) interface{}

	// Visit a parse tree produced by goParser#intliteralAST.
	VisitIntliteralAST(ctx *IntliteralASTContext) interface{}

	// Visit a parse tree produced by goParser#floatliteralAST.
	VisitFloatliteralAST(ctx *FloatliteralASTContext) interface{}

	// Visit a parse tree produced by goParser#runliteralAST.
	VisitRunliteralAST(ctx *RunliteralASTContext) interface{}

	// Visit a parse tree produced by goParser#rawsliteralAST.
	VisitRawsliteralAST(ctx *RawsliteralASTContext) interface{}

	// Visit a parse tree produced by goParser#interpretedliteralAST.
	VisitInterpretedliteralAST(ctx *InterpretedliteralASTContext) interface{}

	// Visit a parse tree produced by goParser#idexAST.
	VisitIdexAST(ctx *IdexASTContext) interface{}

	// Visit a parse tree produced by goParser#argumentsAST.
	VisitArgumentsAST(ctx *ArgumentsASTContext) interface{}

	// Visit a parse tree produced by goParser#selectorAST.
	VisitSelectorAST(ctx *SelectorASTContext) interface{}

	// Visit a parse tree produced by goParser#appendExpressionAST.
	VisitAppendExpressionAST(ctx *AppendExpressionASTContext) interface{}

	// Visit a parse tree produced by goParser#lengthExpressionAST.
	VisitLengthExpressionAST(ctx *LengthExpressionASTContext) interface{}

	// Visit a parse tree produced by goParser#capExpressionAST.
	VisitCapExpressionAST(ctx *CapExpressionASTContext) interface{}

	// Visit a parse tree produced by goParser#statementListAST.
	VisitStatementListAST(ctx *StatementListASTContext) interface{}

	// Visit a parse tree produced by goParser#blockAST.
	VisitBlockAST(ctx *BlockASTContext) interface{}

	// Visit a parse tree produced by goParser#printSAST.
	VisitPrintSAST(ctx *PrintSASTContext) interface{}

	// Visit a parse tree produced by goParser#printlnSAST.
	VisitPrintlnSAST(ctx *PrintlnSASTContext) interface{}

	// Visit a parse tree produced by goParser#returnSAST.
	VisitReturnSAST(ctx *ReturnSASTContext) interface{}

	// Visit a parse tree produced by goParser#breakSAST.
	VisitBreakSAST(ctx *BreakSASTContext) interface{}

	// Visit a parse tree produced by goParser#continueSAST.
	VisitContinueSAST(ctx *ContinueSASTContext) interface{}

	// Visit a parse tree produced by goParser#simpleStatementSAST.
	VisitSimpleStatementSAST(ctx *SimpleStatementSASTContext) interface{}

	// Visit a parse tree produced by goParser#blockSAST.
	VisitBlockSAST(ctx *BlockSASTContext) interface{}

	// Visit a parse tree produced by goParser#switchSAST.
	VisitSwitchSAST(ctx *SwitchSASTContext) interface{}

	// Visit a parse tree produced by goParser#ifStatementSAST.
	VisitIfStatementSAST(ctx *IfStatementSASTContext) interface{}

	// Visit a parse tree produced by goParser#loopSAST.
	VisitLoopSAST(ctx *LoopSASTContext) interface{}

	// Visit a parse tree produced by goParser#typeDeclSAST.
	VisitTypeDeclSAST(ctx *TypeDeclSASTContext) interface{}

	// Visit a parse tree produced by goParser#variableDeclSAST.
	VisitVariableDeclSAST(ctx *VariableDeclSASTContext) interface{}

	// Visit a parse tree produced by goParser#epsilonSSAST.
	VisitEpsilonSSAST(ctx *EpsilonSSASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionINCSSAST.
	VisitExpressionINCSSAST(ctx *ExpressionINCSSASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionDECSSAST.
	VisitExpressionDECSSAST(ctx *ExpressionDECSSASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionEpsilonSSAST.
	VisitExpressionEpsilonSSAST(ctx *ExpressionEpsilonSSASTContext) interface{}

	// Visit a parse tree produced by goParser#assigmentStatementSSAST.
	VisitAssigmentStatementSSAST(ctx *AssigmentStatementSSASTContext) interface{}

	// Visit a parse tree produced by goParser#expListSSAST.
	VisitExpListSSAST(ctx *ExpListSSASTContext) interface{}

	// Visit a parse tree produced by goParser#expListASAST.
	VisitExpListASAST(ctx *ExpListASASTContext) interface{}

	// Visit a parse tree produced by goParser#expASAST.
	VisitExpASAST(ctx *ExpASASTContext) interface{}

	// Visit a parse tree produced by goParser#isExpressionBlockISAST.
	VisitIsExpressionBlockISAST(ctx *IsExpressionBlockISASTContext) interface{}

	// Visit a parse tree produced by goParser#isExpBlockIsISAST.
	VisitIsExpBlockIsISAST(ctx *IsExpBlockIsISASTContext) interface{}

	// Visit a parse tree produced by goParser#isExpBlockISAST.
	VisitIsExpBlockISAST(ctx *IsExpBlockISASTContext) interface{}

	// Visit a parse tree produced by goParser#isSSExpBlockISAST.
	VisitIsSSExpBlockISAST(ctx *IsSSExpBlockISASTContext) interface{}

	// Visit a parse tree produced by goParser#isSSExpBlockifSAST.
	VisitIsSSExpBlockifSAST(ctx *IsSSExpBlockifSASTContext) interface{}

	// Visit a parse tree produced by goParser#isSSExpBlockBlockAST.
	VisitIsSSExpBlockBlockAST(ctx *IsSSExpBlockBlockASTContext) interface{}

	// Visit a parse tree produced by goParser#fBlockLAST.
	VisitFBlockLAST(ctx *FBlockLASTContext) interface{}

	// Visit a parse tree produced by goParser#fExpBlockLAST.
	VisitFExpBlockLAST(ctx *FExpBlockLASTContext) interface{}

	// Visit a parse tree produced by goParser#fSimpStateExpSBlockLAST.
	VisitFSimpStateExpSBlockLAST(ctx *FSimpStateExpSBlockLASTContext) interface{}

	// Visit a parse tree produced by goParser#fSimpStateSimpStateBlockLAST.
	VisitFSimpStateSimpStateBlockLAST(ctx *FSimpStateSimpStateBlockLASTContext) interface{}

	// Visit a parse tree produced by goParser#sSimpleStatSAST.
	VisitSSimpleStatSAST(ctx *SSimpleStatSASTContext) interface{}

	// Visit a parse tree produced by goParser#sExpressionSAST.
	VisitSExpressionSAST(ctx *SExpressionSASTContext) interface{}

	// Visit a parse tree produced by goParser#sSimpleStatExpCCListSAST.
	VisitSSimpleStatExpCCListSAST(ctx *SSimpleStatExpCCListSASTContext) interface{}

	// Visit a parse tree produced by goParser#sBlockSAST.
	VisitSBlockSAST(ctx *SBlockSASTContext) interface{}

	// Visit a parse tree produced by goParser#epsilonECCLAST.
	VisitEpsilonECCLAST(ctx *EpsilonECCLASTContext) interface{}

	// Visit a parse tree produced by goParser#expCaseClauseECCLAST.
	VisitExpCaseClauseECCLAST(ctx *ExpCaseClauseECCLASTContext) interface{}

	// Visit a parse tree produced by goParser#expSwitchCaseECCAST.
	VisitExpSwitchCaseECCAST(ctx *ExpSwitchCaseECCASTContext) interface{}

	// Visit a parse tree produced by goParser#caseESCAST.
	VisitCaseESCAST(ctx *CaseESCASTContext) interface{}

	// Visit a parse tree produced by goParser#defaultESCAST.
	VisitDefaultESCAST(ctx *DefaultESCASTContext) interface{}

	// Visit a parse tree produced by goParser#epsilonAST.
	VisitEpsilonAST(ctx *EpsilonASTContext) interface{}
}
