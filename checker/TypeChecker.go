package checker

import (
	generated "MiniGO/compiler/generated"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

const (
	PACKAGE = 0
	INT     = 1
	CHAR    = 2
	STRING  = 4
	BOOL    = 5
	FLOAT   = 6
	RUNE    = 7
)

type TypeChecker struct {
	*generated.BasegoParserVisitor
	symbolTable *SymbolTable
	errorList   []string
}

func NewTypeChecker() *TypeChecker {
	return &TypeChecker{
		BasegoParserVisitor: &generated.BasegoParserVisitor{},
		symbolTable:         NewSymbolTable(),
		errorList:           []string{},
	}
}
func (v *TypeChecker) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (tc *TypeChecker) HasErrors() bool {
	return len(tc.errorList) != 0
}

func (tc *TypeChecker) GetErrors() string {
	if !tc.HasErrors() {
		return "0 errors"
	}
	var builder strings.Builder
	for _, err := range tc.errorList {
		builder.WriteString(fmt.Sprintf("%s\n", err))
	}
	return builder.String()
}

func (tc *TypeChecker) VisitRootAST(ctx *generated.RootASTContext) interface{} {
	fmt.Printf("entrooo")
	packageNameSymbol := ctx.ID() // Obtiene el nombre del paquete del contexto

	if tc.symbolTable.search(packageNameSymbol.GetText()) != nil { // Verifica si el paquete ya está definido
		tc.errorList = append(tc.errorList, "Error: Package '"+packageNameSymbol.GetText()+"' already defined")
	} else {
		tc.symbolTable.insert(packageNameSymbol, PACKAGE) // Inserta el paquete en la tabla de símbolos
	}
	//t.Visit(ctx.TopDeclarationList())
	ctx.TopDeclarationList().Accept(tc)
	tc.symbolTable.printTable()

	return nil
}

func (tc *TypeChecker) VisitTopDeclarationListAst(ctx *generated.TopDeclarationListASTContext) interface{} {
	fmt.Println("holaaa")
	return nil
}

func (tc *TypeChecker) VisitVarVDAST(ctx *generated.VarVDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitInnerVarDeclAST(ctx *generated.InnerVarDeclsASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIdListDTSVDAST(ctx *generated.IdListDTSVDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIdListSVDAST(ctx *generated.IdListSVDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSingleVarDAST(ctx *generated.SingleVarDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSingleVarDeclNoExpsAST(ctx *generated.SingleVarDeclNoExpsASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitTypeTDAST(ctx *generated.TypeTDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitInnerTypeDeclsAST(ctx *generated.IdListSVDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIdSYDAST(ctx *generated.IdSYDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFuncDeclASt(ctx *generated.FuncDeclASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFuncFFDAST(ctx *generated.FuncFFDASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFuncArgDeclsAST(ctx *generated.FuncArgDeclsASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitLpDTAST(ctx *generated.LpDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIdDTAST(ctx *generated.IdDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSliceDeclTypeDTAST(ctx *generated.SliceDeclTypeDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitArrayDeclTypeDTAST(ctx *generated.ArrayDeclTypeDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitStructDeclTypeDTAST(ctx *generated.StructDeclTypeDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitLsbSDTAST(ctx *generated.LsbSDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitLsbADTAST(ctx *generated.LsbADTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitStructSDTAST(ctx *generated.StructSDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitStructMemDeclsAST(ctx *generated.StructMemDeclsASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIdentifierListAST(ctx *generated.IdentifierListASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitPrimaryExpressionEAST(ctx *generated.PrimaryExpressionEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpressionEAST(ctx *generated.ExpressionEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitOperationEAST(ctx *generated.OperatorEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpressionListAST(ctx *generated.ExpressionListASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitOperandPEAST(ctx *generated.OperandPEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitPrimaryExpPEAST(ctx *generated.PrimaryExpPEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitAppendExpPEAST(ctx *generated.AppendExpPEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitLengthExpPEAST(ctx *generated.LengthExpPEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitCapExpPEAST(ctx *generated.CapExpPEASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitLiteralOAST(ctx *generated.LiteralOASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIdOAST(ctx *generated.IdOASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpressionOAST(ctx *generated.ExpressionOASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIntLiteralAST(ctx *generated.IntliteralASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFloatLiteralAST(ctx *generated.FloatliteralASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitRunLiteralAST(ctx *generated.RunliteralASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitRawsLiteralAST(ctx *generated.RawsliteralASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitInterpretedLiteralAST(ctx *generated.InterpretedliteralASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIdexAST(ctx *generated.IdexASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitArgumentsAST(ctx *generated.ArgumentsASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSelectorAST(ctx *generated.SelectorASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitAppendExpressionAST(ctx *generated.AppendExpressionASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitLengthExpressionAST(ctx *generated.LengthExpressionASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitCapExpressionAST(ctx *generated.CapExpressionASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitStatementListAST(ctx *generated.StatementListASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitBlockAST(ctx *generated.BlockASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitPrintSAST(ctx *generated.PrintSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitPrintlnSAST(ctx *generated.PrintlnSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitReturnSAST(ctx *generated.ReturnSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitBreakSAST(ctx *generated.BreakSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitContinueSAST(ctx *generated.ContinueSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSimpleStatementSAST(ctx *generated.SimpleStatementSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitBlockSAST(ctx *generated.BlockSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSwitchSAST(ctx *generated.SwitchSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIfStatementSAST(ctx *generated.IfStatementSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitTypeDeclSAST(ctx *generated.TypeDeclSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitVariableDeclSAST(ctx *generated.VariableDeclSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitEpsilonSSAST(ctx *generated.EpsilonSSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpressionSSAST(ctx *generated.ExpressionSSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitAssigmentStatement(ctx *generated.AssigmentStatementSSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpListSSAST(ctx *generated.ExpListSSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpListASAST(ctx *generated.ExpListASASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpASAST(ctx *generated.ExpASASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIsExpressionBlockISAST(ctx *generated.IsExpressionBlockISASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIsExpBlockIsISAST(ctx *generated.IsExpBlockIsISASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIsSSExpBlockISAST(ctx *generated.IsSSExpBlockISASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIsSSExpBlockIfSAST(ctx *generated.IsSSExpBlockifSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitIsSSExpBlockBlockAST(ctx *generated.IsSSExpBlockBlockASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFBlockLAST(ctx *generated.FBlockLASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFExpBlockLAST(ctx *generated.FExpBlockLASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFSimpStateExpSBlockLAST(ctx *generated.FSimpStateExpSBlockLASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitFSimpStateSimpStateBlockLAST(ctx *generated.FSimpStateSimpStateBlockLASTContext) interface{} {
	return tc.VisitChildren(ctx)
}

func (tc *TypeChecker) VisitSSimpleStatSAST(ctx *generated.SSimpleStatSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitISExpressionSAST(ctx *generated.SExpressionSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSSimpleStatExpCCListSAST(ctx *generated.SSimpleStatExpCCListSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitSBlockSAST(ctx *generated.SBlockSASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitEpsilonECCLAST(ctx *generated.EpsilonECCLASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpCaseClauseECCLAST(ctx *generated.ExpCaseClauseECCLASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitExpSwitchCaseECCAST(ctx *generated.ExpSwitchCaseECCASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitCaseESCAST(ctx *generated.CaseESCASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitDefaultESCAST(ctx *generated.DefaultESCASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitOperatorAST(ctx *generated.OperatorContext) interface{} {
	return tc.VisitChildren(ctx)
}
func (tc *TypeChecker) VisitEpsilonAST(ctx *generated.EpsilonASTContext) interface{} {
	return tc.VisitChildren(ctx)
}
