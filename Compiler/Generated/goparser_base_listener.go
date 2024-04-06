// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/Compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Generated // goParser
import "github.com/antlr4-go/antlr/v4"

// BasegoParserListener is a complete listener for a parse tree produced by goParser.
type BasegoParserListener struct{}

var _ goParserListener = &BasegoParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegoParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegoParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegoParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegoParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRootAST is called when production rootAST is entered.
func (s *BasegoParserListener) EnterRootAST(ctx *RootASTContext) {}

// ExitRootAST is called when production rootAST is exited.
func (s *BasegoParserListener) ExitRootAST(ctx *RootASTContext) {}

// EnterTopDeclarationListAST is called when production topDeclarationListAST is entered.
func (s *BasegoParserListener) EnterTopDeclarationListAST(ctx *TopDeclarationListASTContext) {}

// ExitTopDeclarationListAST is called when production topDeclarationListAST is exited.
func (s *BasegoParserListener) ExitTopDeclarationListAST(ctx *TopDeclarationListASTContext) {}

// EnterVarVDAST is called when production varVDAST is entered.
func (s *BasegoParserListener) EnterVarVDAST(ctx *VarVDASTContext) {}

// ExitVarVDAST is called when production varVDAST is exited.
func (s *BasegoParserListener) ExitVarVDAST(ctx *VarVDASTContext) {}

// EnterInnerVarDeclsAST is called when production innerVarDeclsAST is entered.
func (s *BasegoParserListener) EnterInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) {}

// ExitInnerVarDeclsAST is called when production innerVarDeclsAST is exited.
func (s *BasegoParserListener) ExitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) {}

// EnterIdentifierListDTSVDAST is called when production identifierListDTSVDAST is entered.
func (s *BasegoParserListener) EnterIdentifierListDTSVDAST(ctx *IdentifierListDTSVDASTContext) {}

// ExitIdentifierListDTSVDAST is called when production identifierListDTSVDAST is exited.
func (s *BasegoParserListener) ExitIdentifierListDTSVDAST(ctx *IdentifierListDTSVDASTContext) {}

// EnterIdentifierListSVDAST is called when production identifierListSVDAST is entered.
func (s *BasegoParserListener) EnterIdentifierListSVDAST(ctx *IdentifierListSVDASTContext) {}

// ExitIdentifierListSVDAST is called when production identifierListSVDAST is exited.
func (s *BasegoParserListener) ExitIdentifierListSVDAST(ctx *IdentifierListSVDASTContext) {}

// EnterSingleVarDAST is called when production singleVarDAST is entered.
func (s *BasegoParserListener) EnterSingleVarDAST(ctx *SingleVarDASTContext) {}

// ExitSingleVarDAST is called when production singleVarDAST is exited.
func (s *BasegoParserListener) ExitSingleVarDAST(ctx *SingleVarDASTContext) {}

// EnterSingleVarDeclNoExpsAST is called when production singleVarDeclNoExpsAST is entered.
func (s *BasegoParserListener) EnterSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) {}

// ExitSingleVarDeclNoExpsAST is called when production singleVarDeclNoExpsAST is exited.
func (s *BasegoParserListener) ExitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) {}

// EnterTypeTDAST is called when production typeTDAST is entered.
func (s *BasegoParserListener) EnterTypeTDAST(ctx *TypeTDASTContext) {}

// ExitTypeTDAST is called when production typeTDAST is exited.
func (s *BasegoParserListener) ExitTypeTDAST(ctx *TypeTDASTContext) {}

// EnterInnerTypeDeclsASt is called when production innerTypeDeclsASt is entered.
func (s *BasegoParserListener) EnterInnerTypeDeclsASt(ctx *InnerTypeDeclsAStContext) {}

// ExitInnerTypeDeclsASt is called when production innerTypeDeclsASt is exited.
func (s *BasegoParserListener) ExitInnerTypeDeclsASt(ctx *InnerTypeDeclsAStContext) {}

// EnterIdSYDAST is called when production idSYDAST is entered.
func (s *BasegoParserListener) EnterIdSYDAST(ctx *IdSYDASTContext) {}

// ExitIdSYDAST is called when production idSYDAST is exited.
func (s *BasegoParserListener) ExitIdSYDAST(ctx *IdSYDASTContext) {}

// EnterFuncDeclAST is called when production funcDeclAST is entered.
func (s *BasegoParserListener) EnterFuncDeclAST(ctx *FuncDeclASTContext) {}

// ExitFuncDeclAST is called when production funcDeclAST is exited.
func (s *BasegoParserListener) ExitFuncDeclAST(ctx *FuncDeclASTContext) {}

// EnterFuncFFDAST is called when production funcFFDAST is entered.
func (s *BasegoParserListener) EnterFuncFFDAST(ctx *FuncFFDASTContext) {}

// ExitFuncFFDAST is called when production funcFFDAST is exited.
func (s *BasegoParserListener) ExitFuncFFDAST(ctx *FuncFFDASTContext) {}

// EnterFuncArgDeclsAST is called when production funcArgDeclsAST is entered.
func (s *BasegoParserListener) EnterFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) {}

// ExitFuncArgDeclsAST is called when production funcArgDeclsAST is exited.
func (s *BasegoParserListener) ExitFuncArgDeclsAST(ctx *FuncArgDeclsASTContext) {}

// EnterLpDTAST is called when production lpDTAST is entered.
func (s *BasegoParserListener) EnterLpDTAST(ctx *LpDTASTContext) {}

// ExitLpDTAST is called when production lpDTAST is exited.
func (s *BasegoParserListener) ExitLpDTAST(ctx *LpDTASTContext) {}

// EnterIdDTAST is called when production idDTAST is entered.
func (s *BasegoParserListener) EnterIdDTAST(ctx *IdDTASTContext) {}

// ExitIdDTAST is called when production idDTAST is exited.
func (s *BasegoParserListener) ExitIdDTAST(ctx *IdDTASTContext) {}

// EnterSliceDeclTypeDTAST is called when production sliceDeclTypeDTAST is entered.
func (s *BasegoParserListener) EnterSliceDeclTypeDTAST(ctx *SliceDeclTypeDTASTContext) {}

// ExitSliceDeclTypeDTAST is called when production sliceDeclTypeDTAST is exited.
func (s *BasegoParserListener) ExitSliceDeclTypeDTAST(ctx *SliceDeclTypeDTASTContext) {}

// EnterArrayDeclTypeDTAST is called when production arrayDeclTypeDTAST is entered.
func (s *BasegoParserListener) EnterArrayDeclTypeDTAST(ctx *ArrayDeclTypeDTASTContext) {}

// ExitArrayDeclTypeDTAST is called when production arrayDeclTypeDTAST is exited.
func (s *BasegoParserListener) ExitArrayDeclTypeDTAST(ctx *ArrayDeclTypeDTASTContext) {}

// EnterDeclTypeAST is called when production declTypeAST is entered.
func (s *BasegoParserListener) EnterDeclTypeAST(ctx *DeclTypeASTContext) {}

// ExitDeclTypeAST is called when production declTypeAST is exited.
func (s *BasegoParserListener) ExitDeclTypeAST(ctx *DeclTypeASTContext) {}

// EnterLsbSDTAST is called when production lsbSDTAST is entered.
func (s *BasegoParserListener) EnterLsbSDTAST(ctx *LsbSDTASTContext) {}

// ExitLsbSDTAST is called when production lsbSDTAST is exited.
func (s *BasegoParserListener) ExitLsbSDTAST(ctx *LsbSDTASTContext) {}

// EnterLsbADTAST is called when production lsbADTAST is entered.
func (s *BasegoParserListener) EnterLsbADTAST(ctx *LsbADTASTContext) {}

// ExitLsbADTAST is called when production lsbADTAST is exited.
func (s *BasegoParserListener) ExitLsbADTAST(ctx *LsbADTASTContext) {}

// EnterStrctSDTAST is called when production strctSDTAST is entered.
func (s *BasegoParserListener) EnterStrctSDTAST(ctx *StrctSDTASTContext) {}

// ExitStrctSDTAST is called when production strctSDTAST is exited.
func (s *BasegoParserListener) ExitStrctSDTAST(ctx *StrctSDTASTContext) {}

// EnterStructMemDeclsAST is called when production structMemDeclsAST is entered.
func (s *BasegoParserListener) EnterStructMemDeclsAST(ctx *StructMemDeclsASTContext) {}

// ExitStructMemDeclsAST is called when production structMemDeclsAST is exited.
func (s *BasegoParserListener) ExitStructMemDeclsAST(ctx *StructMemDeclsASTContext) {}

// EnterIdentifierListAST is called when production identifierListAST is entered.
func (s *BasegoParserListener) EnterIdentifierListAST(ctx *IdentifierListASTContext) {}

// ExitIdentifierListAST is called when production identifierListAST is exited.
func (s *BasegoParserListener) ExitIdentifierListAST(ctx *IdentifierListASTContext) {}

// EnterExpressionAST is called when production expressionAST is entered.
func (s *BasegoParserListener) EnterExpressionAST(ctx *ExpressionASTContext) {}

// ExitExpressionAST is called when production expressionAST is exited.
func (s *BasegoParserListener) ExitExpressionAST(ctx *ExpressionASTContext) {}

// EnterExpressionEAST is called when production expressionEAST is entered.
func (s *BasegoParserListener) EnterExpressionEAST(ctx *ExpressionEASTContext) {}

// ExitExpressionEAST is called when production expressionEAST is exited.
func (s *BasegoParserListener) ExitExpressionEAST(ctx *ExpressionEASTContext) {}

// EnterPrimaryExpressionEAST is called when production primaryExpressionEAST is entered.
func (s *BasegoParserListener) EnterPrimaryExpressionEAST(ctx *PrimaryExpressionEASTContext) {}

// ExitPrimaryExpressionEAST is called when production primaryExpressionEAST is exited.
func (s *BasegoParserListener) ExitPrimaryExpressionEAST(ctx *PrimaryExpressionEASTContext) {}

// EnterExpressionListASt is called when production expressionListASt is entered.
func (s *BasegoParserListener) EnterExpressionListASt(ctx *ExpressionListAStContext) {}

// ExitExpressionListASt is called when production expressionListASt is exited.
func (s *BasegoParserListener) ExitExpressionListASt(ctx *ExpressionListAStContext) {}

// EnterCapExpressionPEAST is called when production capExpressionPEAST is entered.
func (s *BasegoParserListener) EnterCapExpressionPEAST(ctx *CapExpressionPEASTContext) {}

// ExitCapExpressionPEAST is called when production capExpressionPEAST is exited.
func (s *BasegoParserListener) ExitCapExpressionPEAST(ctx *CapExpressionPEASTContext) {}

// EnterPrimaryExpressionPEAST is called when production primaryExpressionPEAST is entered.
func (s *BasegoParserListener) EnterPrimaryExpressionPEAST(ctx *PrimaryExpressionPEASTContext) {}

// ExitPrimaryExpressionPEAST is called when production primaryExpressionPEAST is exited.
func (s *BasegoParserListener) ExitPrimaryExpressionPEAST(ctx *PrimaryExpressionPEASTContext) {}

// EnterLengthExpressionPEAST is called when production lengthExpressionPEAST is entered.
func (s *BasegoParserListener) EnterLengthExpressionPEAST(ctx *LengthExpressionPEASTContext) {}

// ExitLengthExpressionPEAST is called when production lengthExpressionPEAST is exited.
func (s *BasegoParserListener) ExitLengthExpressionPEAST(ctx *LengthExpressionPEASTContext) {}

// EnterAppendExpressionPEAST is called when production appendExpressionPEAST is entered.
func (s *BasegoParserListener) EnterAppendExpressionPEAST(ctx *AppendExpressionPEASTContext) {}

// ExitAppendExpressionPEAST is called when production appendExpressionPEAST is exited.
func (s *BasegoParserListener) ExitAppendExpressionPEAST(ctx *AppendExpressionPEASTContext) {}

// EnterOperandPEAST is called when production operandPEAST is entered.
func (s *BasegoParserListener) EnterOperandPEAST(ctx *OperandPEASTContext) {}

// ExitOperandPEAST is called when production operandPEAST is exited.
func (s *BasegoParserListener) ExitOperandPEAST(ctx *OperandPEASTContext) {}

// EnterLiteralOAST is called when production literalOAST is entered.
func (s *BasegoParserListener) EnterLiteralOAST(ctx *LiteralOASTContext) {}

// ExitLiteralOAST is called when production literalOAST is exited.
func (s *BasegoParserListener) ExitLiteralOAST(ctx *LiteralOASTContext) {}

// EnterIdOAST is called when production idOAST is entered.
func (s *BasegoParserListener) EnterIdOAST(ctx *IdOASTContext) {}

// ExitIdOAST is called when production idOAST is exited.
func (s *BasegoParserListener) ExitIdOAST(ctx *IdOASTContext) {}

// EnterLpOAST is called when production lpOAST is entered.
func (s *BasegoParserListener) EnterLpOAST(ctx *LpOASTContext) {}

// ExitLpOAST is called when production lpOAST is exited.
func (s *BasegoParserListener) ExitLpOAST(ctx *LpOASTContext) {}

// EnterIntliteralAST is called when production intliteralAST is entered.
func (s *BasegoParserListener) EnterIntliteralAST(ctx *IntliteralASTContext) {}

// ExitIntliteralAST is called when production intliteralAST is exited.
func (s *BasegoParserListener) ExitIntliteralAST(ctx *IntliteralASTContext) {}

// EnterFloatliteralAST is called when production floatliteralAST is entered.
func (s *BasegoParserListener) EnterFloatliteralAST(ctx *FloatliteralASTContext) {}

// ExitFloatliteralAST is called when production floatliteralAST is exited.
func (s *BasegoParserListener) ExitFloatliteralAST(ctx *FloatliteralASTContext) {}

// EnterRunliteral is called when production runliteral is entered.
func (s *BasegoParserListener) EnterRunliteral(ctx *RunliteralContext) {}

// ExitRunliteral is called when production runliteral is exited.
func (s *BasegoParserListener) ExitRunliteral(ctx *RunliteralContext) {}

// EnterRawsliteralAST is called when production rawsliteralAST is entered.
func (s *BasegoParserListener) EnterRawsliteralAST(ctx *RawsliteralASTContext) {}

// ExitRawsliteralAST is called when production rawsliteralAST is exited.
func (s *BasegoParserListener) ExitRawsliteralAST(ctx *RawsliteralASTContext) {}

// EnterInterpretedliteralAST is called when production interpretedliteralAST is entered.
func (s *BasegoParserListener) EnterInterpretedliteralAST(ctx *InterpretedliteralASTContext) {}

// ExitInterpretedliteralAST is called when production interpretedliteralAST is exited.
func (s *BasegoParserListener) ExitInterpretedliteralAST(ctx *InterpretedliteralASTContext) {}

// EnterIdexAST is called when production idexAST is entered.
func (s *BasegoParserListener) EnterIdexAST(ctx *IdexASTContext) {}

// ExitIdexAST is called when production idexAST is exited.
func (s *BasegoParserListener) ExitIdexAST(ctx *IdexASTContext) {}

// EnterArgumentsAST is called when production argumentsAST is entered.
func (s *BasegoParserListener) EnterArgumentsAST(ctx *ArgumentsASTContext) {}

// ExitArgumentsAST is called when production argumentsAST is exited.
func (s *BasegoParserListener) ExitArgumentsAST(ctx *ArgumentsASTContext) {}

// EnterSelectorAST is called when production selectorAST is entered.
func (s *BasegoParserListener) EnterSelectorAST(ctx *SelectorASTContext) {}

// ExitSelectorAST is called when production selectorAST is exited.
func (s *BasegoParserListener) ExitSelectorAST(ctx *SelectorASTContext) {}

// EnterAppendExpressionAST is called when production appendExpressionAST is entered.
func (s *BasegoParserListener) EnterAppendExpressionAST(ctx *AppendExpressionASTContext) {}

// ExitAppendExpressionAST is called when production appendExpressionAST is exited.
func (s *BasegoParserListener) ExitAppendExpressionAST(ctx *AppendExpressionASTContext) {}

// EnterLengthExpressionAST is called when production lengthExpressionAST is entered.
func (s *BasegoParserListener) EnterLengthExpressionAST(ctx *LengthExpressionASTContext) {}

// ExitLengthExpressionAST is called when production lengthExpressionAST is exited.
func (s *BasegoParserListener) ExitLengthExpressionAST(ctx *LengthExpressionASTContext) {}

// EnterCapExpressionAST is called when production capExpressionAST is entered.
func (s *BasegoParserListener) EnterCapExpressionAST(ctx *CapExpressionASTContext) {}

// ExitCapExpressionAST is called when production capExpressionAST is exited.
func (s *BasegoParserListener) ExitCapExpressionAST(ctx *CapExpressionASTContext) {}

// EnterStatementListAST is called when production statementListAST is entered.
func (s *BasegoParserListener) EnterStatementListAST(ctx *StatementListASTContext) {}

// ExitStatementListAST is called when production statementListAST is exited.
func (s *BasegoParserListener) ExitStatementListAST(ctx *StatementListASTContext) {}

// EnterBlockAST is called when production blockAST is entered.
func (s *BasegoParserListener) EnterBlockAST(ctx *BlockASTContext) {}

// ExitBlockAST is called when production blockAST is exited.
func (s *BasegoParserListener) ExitBlockAST(ctx *BlockASTContext) {}

// EnterPrintSAST is called when production printSAST is entered.
func (s *BasegoParserListener) EnterPrintSAST(ctx *PrintSASTContext) {}

// ExitPrintSAST is called when production printSAST is exited.
func (s *BasegoParserListener) ExitPrintSAST(ctx *PrintSASTContext) {}

// EnterPrintlnSAST is called when production printlnSAST is entered.
func (s *BasegoParserListener) EnterPrintlnSAST(ctx *PrintlnSASTContext) {}

// ExitPrintlnSAST is called when production printlnSAST is exited.
func (s *BasegoParserListener) ExitPrintlnSAST(ctx *PrintlnSASTContext) {}

// EnterReturnSAST is called when production returnSAST is entered.
func (s *BasegoParserListener) EnterReturnSAST(ctx *ReturnSASTContext) {}

// ExitReturnSAST is called when production returnSAST is exited.
func (s *BasegoParserListener) ExitReturnSAST(ctx *ReturnSASTContext) {}

// EnterBreakSAST is called when production breakSAST is entered.
func (s *BasegoParserListener) EnterBreakSAST(ctx *BreakSASTContext) {}

// ExitBreakSAST is called when production breakSAST is exited.
func (s *BasegoParserListener) ExitBreakSAST(ctx *BreakSASTContext) {}

// EnterContinueSAST is called when production continueSAST is entered.
func (s *BasegoParserListener) EnterContinueSAST(ctx *ContinueSASTContext) {}

// ExitContinueSAST is called when production continueSAST is exited.
func (s *BasegoParserListener) ExitContinueSAST(ctx *ContinueSASTContext) {}

// EnterSimpleStatementSAST is called when production simpleStatementSAST is entered.
func (s *BasegoParserListener) EnterSimpleStatementSAST(ctx *SimpleStatementSASTContext) {}

// ExitSimpleStatementSAST is called when production simpleStatementSAST is exited.
func (s *BasegoParserListener) ExitSimpleStatementSAST(ctx *SimpleStatementSASTContext) {}

// EnterBlockSAST is called when production blockSAST is entered.
func (s *BasegoParserListener) EnterBlockSAST(ctx *BlockSASTContext) {}

// ExitBlockSAST is called when production blockSAST is exited.
func (s *BasegoParserListener) ExitBlockSAST(ctx *BlockSASTContext) {}

// EnterSwitchSAST is called when production switchSAST is entered.
func (s *BasegoParserListener) EnterSwitchSAST(ctx *SwitchSASTContext) {}

// ExitSwitchSAST is called when production switchSAST is exited.
func (s *BasegoParserListener) ExitSwitchSAST(ctx *SwitchSASTContext) {}

// EnterIfStatementSAST is called when production ifStatementSAST is entered.
func (s *BasegoParserListener) EnterIfStatementSAST(ctx *IfStatementSASTContext) {}

// ExitIfStatementSAST is called when production ifStatementSAST is exited.
func (s *BasegoParserListener) ExitIfStatementSAST(ctx *IfStatementSASTContext) {}

// EnterLoopSAST is called when production loopSAST is entered.
func (s *BasegoParserListener) EnterLoopSAST(ctx *LoopSASTContext) {}

// ExitLoopSAST is called when production loopSAST is exited.
func (s *BasegoParserListener) ExitLoopSAST(ctx *LoopSASTContext) {}

// EnterTypeDeclSAST is called when production typeDeclSAST is entered.
func (s *BasegoParserListener) EnterTypeDeclSAST(ctx *TypeDeclSASTContext) {}

// ExitTypeDeclSAST is called when production typeDeclSAST is exited.
func (s *BasegoParserListener) ExitTypeDeclSAST(ctx *TypeDeclSASTContext) {}

// EnterVariableDeclSAST is called when production variableDeclSAST is entered.
func (s *BasegoParserListener) EnterVariableDeclSAST(ctx *VariableDeclSASTContext) {}

// ExitVariableDeclSAST is called when production variableDeclSAST is exited.
func (s *BasegoParserListener) ExitVariableDeclSAST(ctx *VariableDeclSASTContext) {}

// EnterEpsilonSSAST is called when production epsilonSSAST is entered.
func (s *BasegoParserListener) EnterEpsilonSSAST(ctx *EpsilonSSASTContext) {}

// ExitEpsilonSSAST is called when production epsilonSSAST is exited.
func (s *BasegoParserListener) ExitEpsilonSSAST(ctx *EpsilonSSASTContext) {}

// EnterExpressionSSAST is called when production expressionSSAST is entered.
func (s *BasegoParserListener) EnterExpressionSSAST(ctx *ExpressionSSASTContext) {}

// ExitExpressionSSAST is called when production expressionSSAST is exited.
func (s *BasegoParserListener) ExitExpressionSSAST(ctx *ExpressionSSASTContext) {}

// EnterAssignmentStatementSS is called when production assignmentStatementSS is entered.
func (s *BasegoParserListener) EnterAssignmentStatementSS(ctx *AssignmentStatementSSContext) {}

// ExitAssignmentStatementSS is called when production assignmentStatementSS is exited.
func (s *BasegoParserListener) ExitAssignmentStatementSS(ctx *AssignmentStatementSSContext) {}

// EnterExpressionListSS is called when production expressionListSS is entered.
func (s *BasegoParserListener) EnterExpressionListSS(ctx *ExpressionListSSContext) {}

// ExitExpressionListSS is called when production expressionListSS is exited.
func (s *BasegoParserListener) ExitExpressionListSS(ctx *ExpressionListSSContext) {}

// EnterExpressionListASAST is called when production expressionListASAST is entered.
func (s *BasegoParserListener) EnterExpressionListASAST(ctx *ExpressionListASASTContext) {}

// ExitExpressionListASAST is called when production expressionListASAST is exited.
func (s *BasegoParserListener) ExitExpressionListASAST(ctx *ExpressionListASASTContext) {}

// EnterExpressionASAST is called when production expressionASAST is entered.
func (s *BasegoParserListener) EnterExpressionASAST(ctx *ExpressionASASTContext) {}

// ExitExpressionASAST is called when production expressionASAST is exited.
func (s *BasegoParserListener) ExitExpressionASAST(ctx *ExpressionASASTContext) {}

// EnterIsExpressionBlockISAST is called when production isExpressionBlockISAST is entered.
func (s *BasegoParserListener) EnterIsExpressionBlockISAST(ctx *IsExpressionBlockISASTContext) {}

// ExitIsExpressionBlockISAST is called when production isExpressionBlockISAST is exited.
func (s *BasegoParserListener) ExitIsExpressionBlockISAST(ctx *IsExpressionBlockISASTContext) {}

// EnterIsExpressionBlockIsISAST is called when production isExpressionBlockIsISAST is entered.
func (s *BasegoParserListener) EnterIsExpressionBlockIsISAST(ctx *IsExpressionBlockIsISASTContext) {}

// ExitIsExpressionBlockIsISAST is called when production isExpressionBlockIsISAST is exited.
func (s *BasegoParserListener) ExitIsExpressionBlockIsISAST(ctx *IsExpressionBlockIsISASTContext) {}

// EnterIsSimpleStamentBlockISAST is called when production isSimpleStamentBlockISAST is entered.
func (s *BasegoParserListener) EnterIsSimpleStamentBlockISAST(ctx *IsSimpleStamentBlockISASTContext) {
}

// ExitIsSimpleStamentBlockISAST is called when production isSimpleStamentBlockISAST is exited.
func (s *BasegoParserListener) ExitIsSimpleStamentBlockISAST(ctx *IsSimpleStamentBlockISASTContext) {}

// EnterIsSimpleStamentExpressionBlockISAST is called when production isSimpleStamentExpressionBlockISAST is entered.
func (s *BasegoParserListener) EnterIsSimpleStamentExpressionBlockISAST(ctx *IsSimpleStamentExpressionBlockISASTContext) {
}

// ExitIsSimpleStamentExpressionBlockISAST is called when production isSimpleStamentExpressionBlockISAST is exited.
func (s *BasegoParserListener) ExitIsSimpleStamentExpressionBlockISAST(ctx *IsSimpleStamentExpressionBlockISASTContext) {
}

// EnterIsSimpleStamentExpressionBlockifSAST is called when production isSimpleStamentExpressionBlockifSAST is entered.
func (s *BasegoParserListener) EnterIsSimpleStamentExpressionBlockifSAST(ctx *IsSimpleStamentExpressionBlockifSASTContext) {
}

// ExitIsSimpleStamentExpressionBlockifSAST is called when production isSimpleStamentExpressionBlockifSAST is exited.
func (s *BasegoParserListener) ExitIsSimpleStamentExpressionBlockifSAST(ctx *IsSimpleStamentExpressionBlockifSASTContext) {
}

// EnterIsSimpleStamentExpressionBlockBlockAST is called when production isSimpleStamentExpressionBlockBlockAST is entered.
func (s *BasegoParserListener) EnterIsSimpleStamentExpressionBlockBlockAST(ctx *IsSimpleStamentExpressionBlockBlockASTContext) {
}

// ExitIsSimpleStamentExpressionBlockBlockAST is called when production isSimpleStamentExpressionBlockBlockAST is exited.
func (s *BasegoParserListener) ExitIsSimpleStamentExpressionBlockBlockAST(ctx *IsSimpleStamentExpressionBlockBlockASTContext) {
}

// EnterFBlockLAST is called when production fBlockLAST is entered.
func (s *BasegoParserListener) EnterFBlockLAST(ctx *FBlockLASTContext) {}

// ExitFBlockLAST is called when production fBlockLAST is exited.
func (s *BasegoParserListener) ExitFBlockLAST(ctx *FBlockLASTContext) {}

// EnterFExpressionBlockLAST is called when production fExpressionBlockLAST is entered.
func (s *BasegoParserListener) EnterFExpressionBlockLAST(ctx *FExpressionBlockLASTContext) {}

// ExitFExpressionBlockLAST is called when production fExpressionBlockLAST is exited.
func (s *BasegoParserListener) ExitFExpressionBlockLAST(ctx *FExpressionBlockLASTContext) {}

// EnterFSimpleStatementLAST is called when production fSimpleStatementLAST is entered.
func (s *BasegoParserListener) EnterFSimpleStatementLAST(ctx *FSimpleStatementLASTContext) {}

// ExitFSimpleStatementLAST is called when production fSimpleStatementLAST is exited.
func (s *BasegoParserListener) ExitFSimpleStatementLAST(ctx *FSimpleStatementLASTContext) {}

// EnterSSimpleStatSAST is called when production sSimpleStatSAST is entered.
func (s *BasegoParserListener) EnterSSimpleStatSAST(ctx *SSimpleStatSASTContext) {}

// ExitSSimpleStatSAST is called when production sSimpleStatSAST is exited.
func (s *BasegoParserListener) ExitSSimpleStatSAST(ctx *SSimpleStatSASTContext) {}

// EnterSExpressionSAST is called when production sExpressionSAST is entered.
func (s *BasegoParserListener) EnterSExpressionSAST(ctx *SExpressionSASTContext) {}

// ExitSExpressionSAST is called when production sExpressionSAST is exited.
func (s *BasegoParserListener) ExitSExpressionSAST(ctx *SExpressionSASTContext) {}

// EnterSBlockSAST is called when production sBlockSAST is entered.
func (s *BasegoParserListener) EnterSBlockSAST(ctx *SBlockSASTContext) {}

// ExitSBlockSAST is called when production sBlockSAST is exited.
func (s *BasegoParserListener) ExitSBlockSAST(ctx *SBlockSASTContext) {}

// EnterEpsilonECCLAST is called when production epsilonECCLAST is entered.
func (s *BasegoParserListener) EnterEpsilonECCLAST(ctx *EpsilonECCLASTContext) {}

// ExitEpsilonECCLAST is called when production epsilonECCLAST is exited.
func (s *BasegoParserListener) ExitEpsilonECCLAST(ctx *EpsilonECCLASTContext) {}

// EnterExpressionCaseClauseECCLAST is called when production expressionCaseClauseECCLAST is entered.
func (s *BasegoParserListener) EnterExpressionCaseClauseECCLAST(ctx *ExpressionCaseClauseECCLASTContext) {
}

// ExitExpressionCaseClauseECCLAST is called when production expressionCaseClauseECCLAST is exited.
func (s *BasegoParserListener) ExitExpressionCaseClauseECCLAST(ctx *ExpressionCaseClauseECCLASTContext) {
}

// EnterExpressionSwitchCaseECCAST is called when production expressionSwitchCaseECCAST is entered.
func (s *BasegoParserListener) EnterExpressionSwitchCaseECCAST(ctx *ExpressionSwitchCaseECCASTContext) {
}

// ExitExpressionSwitchCaseECCAST is called when production expressionSwitchCaseECCAST is exited.
func (s *BasegoParserListener) ExitExpressionSwitchCaseECCAST(ctx *ExpressionSwitchCaseECCASTContext) {
}

// EnterCaseESCAST is called when production caseESCAST is entered.
func (s *BasegoParserListener) EnterCaseESCAST(ctx *CaseESCASTContext) {}

// ExitCaseESCAST is called when production caseESCAST is exited.
func (s *BasegoParserListener) ExitCaseESCAST(ctx *CaseESCASTContext) {}

// EnterDefaultESCAST is called when production defaultESCAST is entered.
func (s *BasegoParserListener) EnterDefaultESCAST(ctx *DefaultESCASTContext) {}

// ExitDefaultESCAST is called when production defaultESCAST is exited.
func (s *BasegoParserListener) ExitDefaultESCAST(ctx *DefaultESCASTContext) {}

// EnterOperator is called when production operator is entered.
func (s *BasegoParserListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BasegoParserListener) ExitOperator(ctx *OperatorContext) {}

// EnterEpsilonAST is called when production epsilonAST is entered.
func (s *BasegoParserListener) EnterEpsilonAST(ctx *EpsilonASTContext) {}

// ExitEpsilonAST is called when production epsilonAST is exited.
func (s *BasegoParserListener) ExitEpsilonAST(ctx *EpsilonASTContext) {}
