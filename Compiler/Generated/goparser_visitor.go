// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/Compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Generated // goParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by goParser.
type goParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by goParser#rootAST.
	VisitRootAST(ctx *RootASTContext) interface{}

	// Visit a parse tree produced by goParser#topDeclarationListAST.
	VisitTopDeclarationListAST(ctx *TopDeclarationListASTContext) interface{}

	// Visit a parse tree produced by goParser#varVDAST.
	VisitVarVDAST(ctx *VarVDASTContext) interface{}

	// Visit a parse tree produced by goParser#innerVarDeclsAST.
	VisitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) interface{}

	// Visit a parse tree produced by goParser#identifierListDTSVDAST.
	VisitIdentifierListDTSVDAST(ctx *IdentifierListDTSVDASTContext) interface{}

	// Visit a parse tree produced by goParser#identifierListSVDAST.
	VisitIdentifierListSVDAST(ctx *IdentifierListSVDASTContext) interface{}

	// Visit a parse tree produced by goParser#singleVarDAST.
	VisitSingleVarDAST(ctx *SingleVarDASTContext) interface{}

	// Visit a parse tree produced by goParser#singleVarDeclNoExpsAST.
	VisitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) interface{}

	// Visit a parse tree produced by goParser#typeTDAST.
	VisitTypeTDAST(ctx *TypeTDASTContext) interface{}

	// Visit a parse tree produced by goParser#innerTypeDeclsASt.
	VisitInnerTypeDeclsASt(ctx *InnerTypeDeclsAStContext) interface{}

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

	// Visit a parse tree produced by goParser#declTypeAST.
	VisitDeclTypeAST(ctx *DeclTypeASTContext) interface{}

	// Visit a parse tree produced by goParser#lsbSDTAST.
	VisitLsbSDTAST(ctx *LsbSDTASTContext) interface{}

	// Visit a parse tree produced by goParser#lsbADTAST.
	VisitLsbADTAST(ctx *LsbADTASTContext) interface{}

	// Visit a parse tree produced by goParser#strctSDTAST.
	VisitStrctSDTAST(ctx *StrctSDTASTContext) interface{}

	// Visit a parse tree produced by goParser#structMemDeclsAST.
	VisitStructMemDeclsAST(ctx *StructMemDeclsASTContext) interface{}

	// Visit a parse tree produced by goParser#identifierListAST.
	VisitIdentifierListAST(ctx *IdentifierListASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionAST.
	VisitExpressionAST(ctx *ExpressionASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionEAST.
	VisitExpressionEAST(ctx *ExpressionEASTContext) interface{}

	// Visit a parse tree produced by goParser#primaryExpressionEAST.
	VisitPrimaryExpressionEAST(ctx *PrimaryExpressionEASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionListASt.
	VisitExpressionListASt(ctx *ExpressionListAStContext) interface{}

	// Visit a parse tree produced by goParser#capExpressionPEAST.
	VisitCapExpressionPEAST(ctx *CapExpressionPEASTContext) interface{}

	// Visit a parse tree produced by goParser#primaryExpressionPEAST.
	VisitPrimaryExpressionPEAST(ctx *PrimaryExpressionPEASTContext) interface{}

	// Visit a parse tree produced by goParser#lengthExpressionPEAST.
	VisitLengthExpressionPEAST(ctx *LengthExpressionPEASTContext) interface{}

	// Visit a parse tree produced by goParser#appendExpressionPEAST.
	VisitAppendExpressionPEAST(ctx *AppendExpressionPEASTContext) interface{}

	// Visit a parse tree produced by goParser#operandPEAST.
	VisitOperandPEAST(ctx *OperandPEASTContext) interface{}

	// Visit a parse tree produced by goParser#literalOAST.
	VisitLiteralOAST(ctx *LiteralOASTContext) interface{}

	// Visit a parse tree produced by goParser#idOAST.
	VisitIdOAST(ctx *IdOASTContext) interface{}

	// Visit a parse tree produced by goParser#lpOAST.
	VisitLpOAST(ctx *LpOASTContext) interface{}

	// Visit a parse tree produced by goParser#intliteralAST.
	VisitIntliteralAST(ctx *IntliteralASTContext) interface{}

	// Visit a parse tree produced by goParser#floatliteralAST.
	VisitFloatliteralAST(ctx *FloatliteralASTContext) interface{}

	// Visit a parse tree produced by goParser#runliteral.
	VisitRunliteral(ctx *RunliteralContext) interface{}

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

	// Visit a parse tree produced by goParser#expressionSSAST.
	VisitExpressionSSAST(ctx *ExpressionSSASTContext) interface{}

	// Visit a parse tree produced by goParser#assignmentStatementSS.
	VisitAssignmentStatementSS(ctx *AssignmentStatementSSContext) interface{}

	// Visit a parse tree produced by goParser#expressionListSS.
	VisitExpressionListSS(ctx *ExpressionListSSContext) interface{}

	// Visit a parse tree produced by goParser#expressionListASAST.
	VisitExpressionListASAST(ctx *ExpressionListASASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionASAST.
	VisitExpressionASAST(ctx *ExpressionASASTContext) interface{}

	// Visit a parse tree produced by goParser#isExpressionBlockISAST.
	VisitIsExpressionBlockISAST(ctx *IsExpressionBlockISASTContext) interface{}

	// Visit a parse tree produced by goParser#isExpressionBlockIsISAST.
	VisitIsExpressionBlockIsISAST(ctx *IsExpressionBlockIsISASTContext) interface{}

	// Visit a parse tree produced by goParser#isSimpleStamentBlockISAST.
	VisitIsSimpleStamentBlockISAST(ctx *IsSimpleStamentBlockISASTContext) interface{}

	// Visit a parse tree produced by goParser#isSimpleStamentExpressionBlockISAST.
	VisitIsSimpleStamentExpressionBlockISAST(ctx *IsSimpleStamentExpressionBlockISASTContext) interface{}

	// Visit a parse tree produced by goParser#isSimpleStamentExpressionBlockifSAST.
	VisitIsSimpleStamentExpressionBlockifSAST(ctx *IsSimpleStamentExpressionBlockifSASTContext) interface{}

	// Visit a parse tree produced by goParser#isSimpleStamentExpressionBlockBlockAST.
	VisitIsSimpleStamentExpressionBlockBlockAST(ctx *IsSimpleStamentExpressionBlockBlockASTContext) interface{}

	// Visit a parse tree produced by goParser#fBlockLAST.
	VisitFBlockLAST(ctx *FBlockLASTContext) interface{}

	// Visit a parse tree produced by goParser#fExpressionBlockLAST.
	VisitFExpressionBlockLAST(ctx *FExpressionBlockLASTContext) interface{}

	// Visit a parse tree produced by goParser#fSimpleStatementLAST.
	VisitFSimpleStatementLAST(ctx *FSimpleStatementLASTContext) interface{}

	// Visit a parse tree produced by goParser#sSimpleStatSAST.
	VisitSSimpleStatSAST(ctx *SSimpleStatSASTContext) interface{}

	// Visit a parse tree produced by goParser#sExpressionSAST.
	VisitSExpressionSAST(ctx *SExpressionSASTContext) interface{}

	// Visit a parse tree produced by goParser#sBlockSAST.
	VisitSBlockSAST(ctx *SBlockSASTContext) interface{}

	// Visit a parse tree produced by goParser#epsilonECCLAST.
	VisitEpsilonECCLAST(ctx *EpsilonECCLASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionCaseClauseECCLAST.
	VisitExpressionCaseClauseECCLAST(ctx *ExpressionCaseClauseECCLASTContext) interface{}

	// Visit a parse tree produced by goParser#expressionSwitchCaseECCAST.
	VisitExpressionSwitchCaseECCAST(ctx *ExpressionSwitchCaseECCASTContext) interface{}

	// Visit a parse tree produced by goParser#caseESCAST.
	VisitCaseESCAST(ctx *CaseESCASTContext) interface{}

	// Visit a parse tree produced by goParser#defaultESCAST.
	VisitDefaultESCAST(ctx *DefaultESCASTContext) interface{}

	// Visit a parse tree produced by goParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by goParser#epsilonAST.
	VisitEpsilonAST(ctx *EpsilonASTContext) interface{}
}
