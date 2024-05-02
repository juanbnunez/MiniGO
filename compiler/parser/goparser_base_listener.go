// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // goParser
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

// EnterSingleVarVDAST is called when production singleVarVDAST is entered.
func (s *BasegoParserListener) EnterSingleVarVDAST(ctx *SingleVarVDASTContext) {}

// ExitSingleVarVDAST is called when production singleVarVDAST is exited.
func (s *BasegoParserListener) ExitSingleVarVDAST(ctx *SingleVarVDASTContext) {}

// EnterInnerVarVDAST is called when production innerVarVDAST is entered.
func (s *BasegoParserListener) EnterInnerVarVDAST(ctx *InnerVarVDASTContext) {}

// ExitInnerVarVDAST is called when production innerVarVDAST is exited.
func (s *BasegoParserListener) ExitInnerVarVDAST(ctx *InnerVarVDASTContext) {}

// EnterLpVDAST is called when production lpVDAST is entered.
func (s *BasegoParserListener) EnterLpVDAST(ctx *LpVDASTContext) {}

// ExitLpVDAST is called when production lpVDAST is exited.
func (s *BasegoParserListener) ExitLpVDAST(ctx *LpVDASTContext) {}

// EnterInnerVarDeclsAST is called when production innerVarDeclsAST is entered.
func (s *BasegoParserListener) EnterInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) {}

// ExitInnerVarDeclsAST is called when production innerVarDeclsAST is exited.
func (s *BasegoParserListener) ExitInnerVarDeclsAST(ctx *InnerVarDeclsASTContext) {}

// EnterIdListDTSVDAST is called when production idListDTSVDAST is entered.
func (s *BasegoParserListener) EnterIdListDTSVDAST(ctx *IdListDTSVDASTContext) {}

// ExitIdListDTSVDAST is called when production idListDTSVDAST is exited.
func (s *BasegoParserListener) ExitIdListDTSVDAST(ctx *IdListDTSVDASTContext) {}

// EnterIdListSVDAST is called when production idListSVDAST is entered.
func (s *BasegoParserListener) EnterIdListSVDAST(ctx *IdListSVDASTContext) {}

// ExitIdListSVDAST is called when production idListSVDAST is exited.
func (s *BasegoParserListener) ExitIdListSVDAST(ctx *IdListSVDASTContext) {}

// EnterSingleVarDAST is called when production singleVarDAST is entered.
func (s *BasegoParserListener) EnterSingleVarDAST(ctx *SingleVarDASTContext) {}

// ExitSingleVarDAST is called when production singleVarDAST is exited.
func (s *BasegoParserListener) ExitSingleVarDAST(ctx *SingleVarDASTContext) {}

// EnterSingleVarDeclNoExpsAST is called when production singleVarDeclNoExpsAST is entered.
func (s *BasegoParserListener) EnterSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) {}

// ExitSingleVarDeclNoExpsAST is called when production singleVarDeclNoExpsAST is exited.
func (s *BasegoParserListener) ExitSingleVarDeclNoExpsAST(ctx *SingleVarDeclNoExpsASTContext) {}

// EnterSingleTypeDeclTDAST is called when production singleTypeDeclTDAST is entered.
func (s *BasegoParserListener) EnterSingleTypeDeclTDAST(ctx *SingleTypeDeclTDASTContext) {}

// ExitSingleTypeDeclTDAST is called when production singleTypeDeclTDAST is exited.
func (s *BasegoParserListener) ExitSingleTypeDeclTDAST(ctx *SingleTypeDeclTDASTContext) {}

// EnterInnerTypeDeclsTDAST is called when production innerTypeDeclsTDAST is entered.
func (s *BasegoParserListener) EnterInnerTypeDeclsTDAST(ctx *InnerTypeDeclsTDASTContext) {}

// ExitInnerTypeDeclsTDAST is called when production innerTypeDeclsTDAST is exited.
func (s *BasegoParserListener) ExitInnerTypeDeclsTDAST(ctx *InnerTypeDeclsTDASTContext) {}

// EnterLPTypeDeclTDAST is called when production lPTypeDeclTDAST is entered.
func (s *BasegoParserListener) EnterLPTypeDeclTDAST(ctx *LPTypeDeclTDASTContext) {}

// ExitLPTypeDeclTDAST is called when production lPTypeDeclTDAST is exited.
func (s *BasegoParserListener) ExitLPTypeDeclTDAST(ctx *LPTypeDeclTDASTContext) {}

// EnterInnerTypeDeclsAST is called when production innerTypeDeclsAST is entered.
func (s *BasegoParserListener) EnterInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) {}

// ExitInnerTypeDeclsAST is called when production innerTypeDeclsAST is exited.
func (s *BasegoParserListener) ExitInnerTypeDeclsAST(ctx *InnerTypeDeclsASTContext) {}

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

// EnterStructDeclTypeDTAST is called when production structDeclTypeDTAST is entered.
func (s *BasegoParserListener) EnterStructDeclTypeDTAST(ctx *StructDeclTypeDTASTContext) {}

// ExitStructDeclTypeDTAST is called when production structDeclTypeDTAST is exited.
func (s *BasegoParserListener) ExitStructDeclTypeDTAST(ctx *StructDeclTypeDTASTContext) {}

// EnterLsbSDTAST is called when production lsbSDTAST is entered.
func (s *BasegoParserListener) EnterLsbSDTAST(ctx *LsbSDTASTContext) {}

// ExitLsbSDTAST is called when production lsbSDTAST is exited.
func (s *BasegoParserListener) ExitLsbSDTAST(ctx *LsbSDTASTContext) {}

// EnterLsbADTAST is called when production lsbADTAST is entered.
func (s *BasegoParserListener) EnterLsbADTAST(ctx *LsbADTASTContext) {}

// ExitLsbADTAST is called when production lsbADTAST is exited.
func (s *BasegoParserListener) ExitLsbADTAST(ctx *LsbADTASTContext) {}

// EnterStructSDTAST is called when production structSDTAST is entered.
func (s *BasegoParserListener) EnterStructSDTAST(ctx *StructSDTASTContext) {}

// ExitStructSDTAST is called when production structSDTAST is exited.
func (s *BasegoParserListener) ExitStructSDTAST(ctx *StructSDTASTContext) {}

// EnterStructMemDeclsAST is called when production structMemDeclsAST is entered.
func (s *BasegoParserListener) EnterStructMemDeclsAST(ctx *StructMemDeclsASTContext) {}

// ExitStructMemDeclsAST is called when production structMemDeclsAST is exited.
func (s *BasegoParserListener) ExitStructMemDeclsAST(ctx *StructMemDeclsASTContext) {}

// EnterIdentifierListAST is called when production identifierListAST is entered.
func (s *BasegoParserListener) EnterIdentifierListAST(ctx *IdentifierListASTContext) {}

// ExitIdentifierListAST is called when production identifierListAST is exited.
func (s *BasegoParserListener) ExitIdentifierListAST(ctx *IdentifierListASTContext) {}

// EnterExpressionEAST is called when production expressionEAST is entered.
func (s *BasegoParserListener) EnterExpressionEAST(ctx *ExpressionEASTContext) {}

// ExitExpressionEAST is called when production expressionEAST is exited.
func (s *BasegoParserListener) ExitExpressionEAST(ctx *ExpressionEASTContext) {}

// EnterPrimaryExpressionEAST is called when production primaryExpressionEAST is entered.
func (s *BasegoParserListener) EnterPrimaryExpressionEAST(ctx *PrimaryExpressionEASTContext) {}

// ExitPrimaryExpressionEAST is called when production primaryExpressionEAST is exited.
func (s *BasegoParserListener) ExitPrimaryExpressionEAST(ctx *PrimaryExpressionEASTContext) {}

// EnterOperatorEAST is called when production operatorEAST is entered.
func (s *BasegoParserListener) EnterOperatorEAST(ctx *OperatorEASTContext) {}

// ExitOperatorEAST is called when production operatorEAST is exited.
func (s *BasegoParserListener) ExitOperatorEAST(ctx *OperatorEASTContext) {}

// EnterExpressionListAST is called when production expressionListAST is entered.
func (s *BasegoParserListener) EnterExpressionListAST(ctx *ExpressionListASTContext) {}

// ExitExpressionListAST is called when production expressionListAST is exited.
func (s *BasegoParserListener) ExitExpressionListAST(ctx *ExpressionListASTContext) {}

// EnterPrimaryExpArgumentsPEAST is called when production primaryExpArgumentsPEAST is entered.
func (s *BasegoParserListener) EnterPrimaryExpArgumentsPEAST(ctx *PrimaryExpArgumentsPEASTContext) {}

// ExitPrimaryExpArgumentsPEAST is called when production primaryExpArgumentsPEAST is exited.
func (s *BasegoParserListener) ExitPrimaryExpArgumentsPEAST(ctx *PrimaryExpArgumentsPEASTContext) {}

// EnterPrimaryExpSelectorPEAST is called when production primaryExpSelectorPEAST is entered.
func (s *BasegoParserListener) EnterPrimaryExpSelectorPEAST(ctx *PrimaryExpSelectorPEASTContext) {}

// ExitPrimaryExpSelectorPEAST is called when production primaryExpSelectorPEAST is exited.
func (s *BasegoParserListener) ExitPrimaryExpSelectorPEAST(ctx *PrimaryExpSelectorPEASTContext) {}

// EnterLengthExpPEAST is called when production lengthExpPEAST is entered.
func (s *BasegoParserListener) EnterLengthExpPEAST(ctx *LengthExpPEASTContext) {}

// ExitLengthExpPEAST is called when production lengthExpPEAST is exited.
func (s *BasegoParserListener) ExitLengthExpPEAST(ctx *LengthExpPEASTContext) {}

// EnterAppendExpPEAST is called when production appendExpPEAST is entered.
func (s *BasegoParserListener) EnterAppendExpPEAST(ctx *AppendExpPEASTContext) {}

// ExitAppendExpPEAST is called when production appendExpPEAST is exited.
func (s *BasegoParserListener) ExitAppendExpPEAST(ctx *AppendExpPEASTContext) {}

// EnterOperandPEAST is called when production operandPEAST is entered.
func (s *BasegoParserListener) EnterOperandPEAST(ctx *OperandPEASTContext) {}

// ExitOperandPEAST is called when production operandPEAST is exited.
func (s *BasegoParserListener) ExitOperandPEAST(ctx *OperandPEASTContext) {}

// EnterCapExpPEAST is called when production capExpPEAST is entered.
func (s *BasegoParserListener) EnterCapExpPEAST(ctx *CapExpPEASTContext) {}

// ExitCapExpPEAST is called when production capExpPEAST is exited.
func (s *BasegoParserListener) ExitCapExpPEAST(ctx *CapExpPEASTContext) {}

// EnterPrimaryExpIndexPEAST is called when production primaryExpIndexPEAST is entered.
func (s *BasegoParserListener) EnterPrimaryExpIndexPEAST(ctx *PrimaryExpIndexPEASTContext) {}

// ExitPrimaryExpIndexPEAST is called when production primaryExpIndexPEAST is exited.
func (s *BasegoParserListener) ExitPrimaryExpIndexPEAST(ctx *PrimaryExpIndexPEASTContext) {}

// EnterLiteralOAST is called when production literalOAST is entered.
func (s *BasegoParserListener) EnterLiteralOAST(ctx *LiteralOASTContext) {}

// ExitLiteralOAST is called when production literalOAST is exited.
func (s *BasegoParserListener) ExitLiteralOAST(ctx *LiteralOASTContext) {}

// EnterIdOAST is called when production idOAST is entered.
func (s *BasegoParserListener) EnterIdOAST(ctx *IdOASTContext) {}

// ExitIdOAST is called when production idOAST is exited.
func (s *BasegoParserListener) ExitIdOAST(ctx *IdOASTContext) {}

// EnterExpressionOAST is called when production expressionOAST is entered.
func (s *BasegoParserListener) EnterExpressionOAST(ctx *ExpressionOASTContext) {}

// ExitExpressionOAST is called when production expressionOAST is exited.
func (s *BasegoParserListener) ExitExpressionOAST(ctx *ExpressionOASTContext) {}

// EnterIntliteralAST is called when production intliteralAST is entered.
func (s *BasegoParserListener) EnterIntliteralAST(ctx *IntliteralASTContext) {}

// ExitIntliteralAST is called when production intliteralAST is exited.
func (s *BasegoParserListener) ExitIntliteralAST(ctx *IntliteralASTContext) {}

// EnterFloatliteralAST is called when production floatliteralAST is entered.
func (s *BasegoParserListener) EnterFloatliteralAST(ctx *FloatliteralASTContext) {}

// ExitFloatliteralAST is called when production floatliteralAST is exited.
func (s *BasegoParserListener) ExitFloatliteralAST(ctx *FloatliteralASTContext) {}

// EnterRunliteralAST is called when production runliteralAST is entered.
func (s *BasegoParserListener) EnterRunliteralAST(ctx *RunliteralASTContext) {}

// ExitRunliteralAST is called when production runliteralAST is exited.
func (s *BasegoParserListener) ExitRunliteralAST(ctx *RunliteralASTContext) {}

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

// EnterExpressionINCSSAST is called when production expressionINCSSAST is entered.
func (s *BasegoParserListener) EnterExpressionINCSSAST(ctx *ExpressionINCSSASTContext) {}

// ExitExpressionINCSSAST is called when production expressionINCSSAST is exited.
func (s *BasegoParserListener) ExitExpressionINCSSAST(ctx *ExpressionINCSSASTContext) {}

// EnterExpressionDECSSAST is called when production expressionDECSSAST is entered.
func (s *BasegoParserListener) EnterExpressionDECSSAST(ctx *ExpressionDECSSASTContext) {}

// ExitExpressionDECSSAST is called when production expressionDECSSAST is exited.
func (s *BasegoParserListener) ExitExpressionDECSSAST(ctx *ExpressionDECSSASTContext) {}

// EnterExpressionEpsilonSSAST is called when production expressionEpsilonSSAST is entered.
func (s *BasegoParserListener) EnterExpressionEpsilonSSAST(ctx *ExpressionEpsilonSSASTContext) {}

// ExitExpressionEpsilonSSAST is called when production expressionEpsilonSSAST is exited.
func (s *BasegoParserListener) ExitExpressionEpsilonSSAST(ctx *ExpressionEpsilonSSASTContext) {}

// EnterAssigmentStatementSSAST is called when production assigmentStatementSSAST is entered.
func (s *BasegoParserListener) EnterAssigmentStatementSSAST(ctx *AssigmentStatementSSASTContext) {}

// ExitAssigmentStatementSSAST is called when production assigmentStatementSSAST is exited.
func (s *BasegoParserListener) ExitAssigmentStatementSSAST(ctx *AssigmentStatementSSASTContext) {}

// EnterExpListSSAST is called when production expListSSAST is entered.
func (s *BasegoParserListener) EnterExpListSSAST(ctx *ExpListSSASTContext) {}

// ExitExpListSSAST is called when production expListSSAST is exited.
func (s *BasegoParserListener) ExitExpListSSAST(ctx *ExpListSSASTContext) {}

// EnterExpListASAST is called when production expListASAST is entered.
func (s *BasegoParserListener) EnterExpListASAST(ctx *ExpListASASTContext) {}

// ExitExpListASAST is called when production expListASAST is exited.
func (s *BasegoParserListener) ExitExpListASAST(ctx *ExpListASASTContext) {}

// EnterExpASAST is called when production expASAST is entered.
func (s *BasegoParserListener) EnterExpASAST(ctx *ExpASASTContext) {}

// ExitExpASAST is called when production expASAST is exited.
func (s *BasegoParserListener) ExitExpASAST(ctx *ExpASASTContext) {}

// EnterIsExpressionBlockISAST is called when production isExpressionBlockISAST is entered.
func (s *BasegoParserListener) EnterIsExpressionBlockISAST(ctx *IsExpressionBlockISASTContext) {}

// ExitIsExpressionBlockISAST is called when production isExpressionBlockISAST is exited.
func (s *BasegoParserListener) ExitIsExpressionBlockISAST(ctx *IsExpressionBlockISASTContext) {}

// EnterIsExpBlockIsISAST is called when production isExpBlockIsISAST is entered.
func (s *BasegoParserListener) EnterIsExpBlockIsISAST(ctx *IsExpBlockIsISASTContext) {}

// ExitIsExpBlockIsISAST is called when production isExpBlockIsISAST is exited.
func (s *BasegoParserListener) ExitIsExpBlockIsISAST(ctx *IsExpBlockIsISASTContext) {}

// EnterIsExpBlockISAST is called when production isExpBlockISAST is entered.
func (s *BasegoParserListener) EnterIsExpBlockISAST(ctx *IsExpBlockISASTContext) {}

// ExitIsExpBlockISAST is called when production isExpBlockISAST is exited.
func (s *BasegoParserListener) ExitIsExpBlockISAST(ctx *IsExpBlockISASTContext) {}

// EnterIsSSExpBlockISAST is called when production isSSExpBlockISAST is entered.
func (s *BasegoParserListener) EnterIsSSExpBlockISAST(ctx *IsSSExpBlockISASTContext) {}

// ExitIsSSExpBlockISAST is called when production isSSExpBlockISAST is exited.
func (s *BasegoParserListener) ExitIsSSExpBlockISAST(ctx *IsSSExpBlockISASTContext) {}

// EnterIsSSExpBlockifSAST is called when production isSSExpBlockifSAST is entered.
func (s *BasegoParserListener) EnterIsSSExpBlockifSAST(ctx *IsSSExpBlockifSASTContext) {}

// ExitIsSSExpBlockifSAST is called when production isSSExpBlockifSAST is exited.
func (s *BasegoParserListener) ExitIsSSExpBlockifSAST(ctx *IsSSExpBlockifSASTContext) {}

// EnterIsSSExpBlockBlockAST is called when production isSSExpBlockBlockAST is entered.
func (s *BasegoParserListener) EnterIsSSExpBlockBlockAST(ctx *IsSSExpBlockBlockASTContext) {}

// ExitIsSSExpBlockBlockAST is called when production isSSExpBlockBlockAST is exited.
func (s *BasegoParserListener) ExitIsSSExpBlockBlockAST(ctx *IsSSExpBlockBlockASTContext) {}

// EnterFBlockLAST is called when production fBlockLAST is entered.
func (s *BasegoParserListener) EnterFBlockLAST(ctx *FBlockLASTContext) {}

// ExitFBlockLAST is called when production fBlockLAST is exited.
func (s *BasegoParserListener) ExitFBlockLAST(ctx *FBlockLASTContext) {}

// EnterFExpBlockLAST is called when production fExpBlockLAST is entered.
func (s *BasegoParserListener) EnterFExpBlockLAST(ctx *FExpBlockLASTContext) {}

// ExitFExpBlockLAST is called when production fExpBlockLAST is exited.
func (s *BasegoParserListener) ExitFExpBlockLAST(ctx *FExpBlockLASTContext) {}

// EnterFSimpStateExpSBlockLAST is called when production fSimpStateExpSBlockLAST is entered.
func (s *BasegoParserListener) EnterFSimpStateExpSBlockLAST(ctx *FSimpStateExpSBlockLASTContext) {}

// ExitFSimpStateExpSBlockLAST is called when production fSimpStateExpSBlockLAST is exited.
func (s *BasegoParserListener) ExitFSimpStateExpSBlockLAST(ctx *FSimpStateExpSBlockLASTContext) {}

// EnterFSimpStateSimpStateBlockLAST is called when production fSimpStateSimpStateBlockLAST is entered.
func (s *BasegoParserListener) EnterFSimpStateSimpStateBlockLAST(ctx *FSimpStateSimpStateBlockLASTContext) {
}

// ExitFSimpStateSimpStateBlockLAST is called when production fSimpStateSimpStateBlockLAST is exited.
func (s *BasegoParserListener) ExitFSimpStateSimpStateBlockLAST(ctx *FSimpStateSimpStateBlockLASTContext) {
}

// EnterSSimpleStatSAST is called when production sSimpleStatSAST is entered.
func (s *BasegoParserListener) EnterSSimpleStatSAST(ctx *SSimpleStatSASTContext) {}

// ExitSSimpleStatSAST is called when production sSimpleStatSAST is exited.
func (s *BasegoParserListener) ExitSSimpleStatSAST(ctx *SSimpleStatSASTContext) {}

// EnterSExpressionSAST is called when production sExpressionSAST is entered.
func (s *BasegoParserListener) EnterSExpressionSAST(ctx *SExpressionSASTContext) {}

// ExitSExpressionSAST is called when production sExpressionSAST is exited.
func (s *BasegoParserListener) ExitSExpressionSAST(ctx *SExpressionSASTContext) {}

// EnterSSimpleStatExpCCListSAST is called when production sSimpleStatExpCCListSAST is entered.
func (s *BasegoParserListener) EnterSSimpleStatExpCCListSAST(ctx *SSimpleStatExpCCListSASTContext) {}

// ExitSSimpleStatExpCCListSAST is called when production sSimpleStatExpCCListSAST is exited.
func (s *BasegoParserListener) ExitSSimpleStatExpCCListSAST(ctx *SSimpleStatExpCCListSASTContext) {}

// EnterSBlockSAST is called when production sBlockSAST is entered.
func (s *BasegoParserListener) EnterSBlockSAST(ctx *SBlockSASTContext) {}

// ExitSBlockSAST is called when production sBlockSAST is exited.
func (s *BasegoParserListener) ExitSBlockSAST(ctx *SBlockSASTContext) {}

// EnterEpsilonECCLAST is called when production epsilonECCLAST is entered.
func (s *BasegoParserListener) EnterEpsilonECCLAST(ctx *EpsilonECCLASTContext) {}

// ExitEpsilonECCLAST is called when production epsilonECCLAST is exited.
func (s *BasegoParserListener) ExitEpsilonECCLAST(ctx *EpsilonECCLASTContext) {}

// EnterExpCaseClauseECCLAST is called when production expCaseClauseECCLAST is entered.
func (s *BasegoParserListener) EnterExpCaseClauseECCLAST(ctx *ExpCaseClauseECCLASTContext) {}

// ExitExpCaseClauseECCLAST is called when production expCaseClauseECCLAST is exited.
func (s *BasegoParserListener) ExitExpCaseClauseECCLAST(ctx *ExpCaseClauseECCLASTContext) {}

// EnterExpSwitchCaseECCAST is called when production expSwitchCaseECCAST is entered.
func (s *BasegoParserListener) EnterExpSwitchCaseECCAST(ctx *ExpSwitchCaseECCASTContext) {}

// ExitExpSwitchCaseECCAST is called when production expSwitchCaseECCAST is exited.
func (s *BasegoParserListener) ExitExpSwitchCaseECCAST(ctx *ExpSwitchCaseECCASTContext) {}

// EnterCaseESCAST is called when production caseESCAST is entered.
func (s *BasegoParserListener) EnterCaseESCAST(ctx *CaseESCASTContext) {}

// ExitCaseESCAST is called when production caseESCAST is exited.
func (s *BasegoParserListener) ExitCaseESCAST(ctx *CaseESCASTContext) {}

// EnterDefaultESCAST is called when production defaultESCAST is entered.
func (s *BasegoParserListener) EnterDefaultESCAST(ctx *DefaultESCASTContext) {}

// ExitDefaultESCAST is called when production defaultESCAST is exited.
func (s *BasegoParserListener) ExitDefaultESCAST(ctx *DefaultESCASTContext) {}

// EnterEpsilonAST is called when production epsilonAST is entered.
func (s *BasegoParserListener) EnterEpsilonAST(ctx *EpsilonASTContext) {}

// ExitEpsilonAST is called when production epsilonAST is exited.
func (s *BasegoParserListener) ExitEpsilonAST(ctx *EpsilonASTContext) {}
