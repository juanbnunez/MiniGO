// Code generated from C:/Users/nunez/OneDrive - Estudiantes ITCR/Documentos/TEC/2024/I Semestre 2024/Compiladores e interpretes/Proyectos/MiniGO/MiniGO/Compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Generated // goParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by goParser.
type goParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by goParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by goParser#topDeclarationList.
	VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{}

	// Visit a parse tree produced by goParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by goParser#innerVarDecls.
	VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{}

	// Visit a parse tree produced by goParser#singleVarDecl.
	VisitSingleVarDecl(ctx *SingleVarDeclContext) interface{}

	// Visit a parse tree produced by goParser#singleVarDeclNoExps.
	VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{}

	// Visit a parse tree produced by goParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by goParser#innerTypeDecls.
	VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{}

	// Visit a parse tree produced by goParser#singleTypeDecl.
	VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{}

	// Visit a parse tree produced by goParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by goParser#funcFrontDecl.
	VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{}

	// Visit a parse tree produced by goParser#funcArgDecls.
	VisitFuncArgDecls(ctx *FuncArgDeclsContext) interface{}

	// Visit a parse tree produced by goParser#declType.
	VisitDeclType(ctx *DeclTypeContext) interface{}

	// Visit a parse tree produced by goParser#sliceDeclType.
	VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{}

	// Visit a parse tree produced by goParser#arrayDeclType.
	VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{}

	// Visit a parse tree produced by goParser#structDeclType.
	VisitStructDeclType(ctx *StructDeclTypeContext) interface{}

	// Visit a parse tree produced by goParser#structMemDecls.
	VisitStructMemDecls(ctx *StructMemDeclsContext) interface{}

	// Visit a parse tree produced by goParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by goParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by goParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by goParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by goParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by goParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by goParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by goParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by goParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by goParser#appendExpression.
	VisitAppendExpression(ctx *AppendExpressionContext) interface{}

	// Visit a parse tree produced by goParser#lengthExpression.
	VisitLengthExpression(ctx *LengthExpressionContext) interface{}

	// Visit a parse tree produced by goParser#capExpression.
	VisitCapExpression(ctx *CapExpressionContext) interface{}

	// Visit a parse tree produced by goParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by goParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by goParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by goParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by goParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by goParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by goParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by goParser#switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by goParser#expressionCaseClauseList.
	VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{}

	// Visit a parse tree produced by goParser#expressionCaseClause.
	VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{}

	// Visit a parse tree produced by goParser#expressionSwitchCase.
	VisitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) interface{}

	// Visit a parse tree produced by goParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}
}
