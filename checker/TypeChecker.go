package checker

import (
	"MiniGO/compiler/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"reflect"
	"strconv"
	"strings"
)

const (
	PACKAGE = 49
	INT     = 77
	CHAR    = 79
	STRING  = 81
	BOOL    = 76
	FLOAT   = 78
	RUNE    = 79
	STRUCT  = 63
)

type TypeChecker struct {
	*parser.BasegoParserVisitor
	symbolTable *SymbolTable
	errorList   []string
}

func NewTypeChecker() *TypeChecker {
	return &TypeChecker{
		BasegoParserVisitor: &parser.BasegoParserVisitor{},
		symbolTable:         NewSymbolTable(),
		errorList:           []string{},
	}
}

func (v *TypeChecker) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *TypeChecker) VisitChildren(node antlr.RuleNode) interface{} {
	var result any
	children := node.GetChildren()
	for _, child := range children {
		childResult := child.(antlr.ParseTree).Accept(v)
		result = childResult
	}
	return result
}
func (v *TypeChecker) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}
func (v *TypeChecker) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
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

func (tc *TypeChecker) VisitRootAST(ctx *parser.RootASTContext) interface{} {
	//fmt.Printf("\n entrooo Visit Root")
	tc.symbolTable.openScope()
	packageNameSymbol := ctx.ID() // Obtiene el nombre del paquete del contexto

	if tc.symbolTable.search(packageNameSymbol.GetText()) != nil { // Verifica si el paquete ya está definido
		tc.errorList = append(tc.errorList, "Error: Package '"+packageNameSymbol.GetText()+"' already defined")
	} else {
		tc.symbolTable.insert(packageNameSymbol, PACKAGE) // Inserta el paquete en la tabla de símbolos
		tc.symbolTable.printTable()
	}
	//tc.symbolTable.printTable()
	tc.Visit(ctx.TopDeclarationList())
	tc.symbolTable.closeScope()

	return nil
}

func (tc *TypeChecker) VisitTopDeclarationListAST(ctx *parser.TopDeclarationListASTContext) interface{} {
	for _, decl := range ctx.GetChildren() {
		switch d := decl.(type) {
		case *parser.SingleVarVDASTContext:
			tc.Visit(d)
		case *parser.InnerVarVDASTContext:
			tc.Visit(d)
		case *parser.LpVDASTContext:
			tc.Visit(d)
		case *parser.SingleTypeDeclTDASTContext:
			tc.Visit(d)
		case *parser.InnerTypeDeclsTDASTContext:
			tc.Visit(d)
		case *parser.LPTypeDeclTDASTContext:
			tc.Visit(d)
		case *parser.FuncDeclASTContext:
			tc.Visit(d)

		}
	}
	return nil
}
func (tc *TypeChecker) VisitSingleVarVDAST(ctx *parser.SingleVarVDASTContext) interface{} {
	tc.Visit(ctx.SingleVarDecl())
	return nil
}
func (tc *TypeChecker) VisitInnerVarVDAST(ctx *parser.InnerVarVDASTContext) interface{} {
	tc.Visit(ctx.InnerVarDecls())
	return nil
}
func (tc *TypeChecker) VisitLpVDAST(ctx *parser.LpVDASTContext) interface{} {
	return nil
}
func (tc *TypeChecker) VisitInnerVarDeclsAST(ctx *parser.InnerVarDeclsASTContext) interface{} {
	for i := 0; i < len(ctx.AllSingleVarDecl()); i++ {
		tc.Visit(ctx.SingleVarDecl(i))
	}
	return nil
}

// ------------------------------------
func (tc *TypeChecker) VisitIdListDTSVDAST(ctx *parser.IdListDTSVDASTContext) interface{} {
	identifierList := tc.Visit(ctx.IdentifierList()).([]antlr.TerminalNode)
	declType := tc.Visit(ctx.DeclType()).(antlr.TerminalNode)
	expressionList := tc.Visit(ctx.ExpressionList()).([]antlr.TerminalNode)

	var expectedElements int //variable para guardar la cantidad de elementos esperados

	types := ctx.DeclType().GetText()
	// Verificar si es un tipo de array o slice
	isArray := strings.HasPrefix(types, "[")

	if isArray == true {
		endIndex := strings.Index(types, "]") // Remove the square brackets "[" and "]"

		numberString := types[1:endIndex] // Take the substring from the start to the character "]"
		var err error
		if numberString != "" { // If numberString is not empty, a specific number of elements is expected
			expectedElements, err = strconv.Atoi(numberString)
			if err != nil {
				tc.errorList = append(tc.errorList, fmt.Sprintf("Error: se esperaba un tamaño del arreglo de tipo int", err))
				return nil
			}
		} else {
			expectedElements = -1 // If numberString is empty, there are no restrictions on the number of elements
		}
	} else {
		isArray = false
	}
	switch declType.GetText() {
	case "int":
		// Validate that the number of elements in expressionList matches the expected number (if specified)
		if isArray == true && expectedElements != -1 && len(expressionList) > expectedElements {
			tc.errorList = append(tc.errorList, fmt.Sprintf("The number of expressions does not match the specified quantity: %d", expectedElements))
		}
		// Check that all expressions are of type INT
		allInt := true
		for _, expr := range expressionList {
			if expr.GetSymbol().GetTokenType() != INT {
				allInt = false
				break
			}
		}
		// If all expressions are INT, perform insertion into the symbol table
		if allInt {
			for i, _ := range identifierList {
				if tc.symbolTable.searchVarInCurrentLevel(identifierList[i].GetText()) != nil {
					tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", identifierList[i].GetText(), identifierList[i].GetSymbol().GetLine(), identifierList[i].GetSymbol().GetColumn()))
				} else {
					tc.symbolTable.insert(identifierList[i], INT)
				}
			}
		} else {
			tc.errorList = append(tc.errorList, fmt.Sprintf("All expressions must be of type int and match the number of identifiers"))
		}
	case "char":
		// Validate that the number of elements in expressionList matches the expected number (if specified)
		if isArray == true && expectedElements != -1 && len(expressionList) > expectedElements {
			tc.errorList = append(tc.errorList, fmt.Sprintf("The number of expressions does not match the specified quantity: %d", expectedElements))
		}
		// Check that all expressions are of type INT
		allChar := true
		for _, expr := range expressionList {
			if expr.GetSymbol().GetTokenType() != CHAR {
				allChar = false
				break
			}
		}
		// If all expressions are INT, perform insertion into the symbol table
		if allChar {
			for i, _ := range identifierList {
				if tc.symbolTable.searchVarInCurrentLevel(identifierList[i].GetText()) != nil {
					tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", identifierList[i].GetText(), identifierList[i].GetSymbol().GetLine(), identifierList[i].GetSymbol().GetColumn()))
				} else {
					tc.symbolTable.insert(identifierList[i], CHAR)
				}
			}
		} else {
			tc.errorList = append(tc.errorList, fmt.Sprintf("All expressions must be of type char and match the number of identifiers"))
		}
	case "string":
		// Validate that the number of elements in expressionList matches the expected number (if specified)
		if isArray == true && expectedElements != -1 && len(expressionList) > expectedElements {
			tc.errorList = append(tc.errorList, fmt.Sprintf("The number of expressions does not match the specified quantity: %d", expectedElements))
		}
		// Check that all expressions are of type INT
		allString := true
		for _, expr := range expressionList {
			if expr.GetSymbol().GetTokenType() != STRING {
				allString = false
				break
			}
		}
		// If all expressions are INT, perform insertion into the symbol table
		if allString {
			for i, _ := range identifierList {
				if tc.symbolTable.searchVarInCurrentLevel(identifierList[i].GetText()) != nil {
					tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", identifierList[i].GetText(), identifierList[i].GetSymbol().GetLine(), identifierList[i].GetSymbol().GetColumn()))
				} else {
					tc.symbolTable.insert(identifierList[i], STRING)
				}
			}
		} else {
			tc.errorList = append(tc.errorList, fmt.Sprintf("All expressions must be of type string and match the number of identifiers"))
		}

	case "bool":
		// Validate that the number of elements in expressionList matches the expected number (if specified)
		if isArray == true && expectedElements != -1 && len(expressionList) > expectedElements {
			tc.errorList = append(tc.errorList, fmt.Sprintf("The number of expressions does not match the specified quantity: %d", expectedElements))
		}
		// Check that all expressions are of type INT
		allBool := true
		for _, expr := range expressionList {
			if expr.GetSymbol().GetTokenType() != BOOL {
				allBool = false
				break
			}
		}
		// If all expressions are INT, perform insertion into the symbol table
		if allBool {
			for i, _ := range identifierList {
				if tc.symbolTable.searchVarInCurrentLevel(identifierList[i].GetText()) != nil {
					tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", identifierList[i].GetText(), identifierList[i].GetSymbol().GetLine(), identifierList[i].GetSymbol().GetColumn()))
				} else {
					tc.symbolTable.insert(identifierList[i], BOOL)
				}
			}
		} else {
			tc.errorList = append(tc.errorList, fmt.Sprintf("All expressions must be of type bool and match the number of identifiers"))
		}
	case "float":
		// Validate that the number of elements in expressionList matches the expected number (if specified)
		if isArray == true && expectedElements != -1 && len(expressionList) > expectedElements {
			tc.errorList = append(tc.errorList, fmt.Sprintf("The number of expressions does not match the specified quantity: %d", expectedElements))
		}
		// Check that all expressions are of type INT
		allFloat := true
		for _, expr := range expressionList {
			if expr.GetSymbol().GetTokenType() != FLOAT {
				allFloat = false
				break
			}
		}
		// If all expressions are INT, perform insertion into the symbol table
		if allFloat {
			for i, _ := range identifierList {
				if tc.symbolTable.searchVarInCurrentLevel(identifierList[i].GetText()) != nil {
					tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", identifierList[i].GetText(), identifierList[i].GetSymbol().GetLine(), identifierList[i].GetSymbol().GetColumn()))
				} else {
					tc.symbolTable.insert(identifierList[i], FLOAT)
				}
			}
		} else {
			tc.errorList = append(tc.errorList, fmt.Sprintf("All expressions must be of type float and match the number of identifiers"))
		}
	case "rune":
		// Validate that the number of elements in expressionList matches the expected number (if specified)
		if isArray == true && expectedElements != -1 && len(expressionList) > expectedElements {
			tc.errorList = append(tc.errorList, fmt.Sprintf("The number of expressions does not match the specified quantity: %d", expectedElements))
		}
		// Check that all expressions are of type INT
		allRune := true
		for _, expr := range expressionList {
			if expr.GetSymbol().GetTokenType() != RUNE {
				allRune = false
				break
			}
		}
		// If all expressions are INT, perform insertion into the symbol table
		if allRune {
			for i, _ := range identifierList {
				if tc.symbolTable.searchVarInCurrentLevel(identifierList[i].GetText()) != nil {
					tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", identifierList[i].GetText(), identifierList[i].GetSymbol().GetLine(), identifierList[i].GetSymbol().GetColumn()))
				} else {
					tc.symbolTable.insert(identifierList[i], RUNE)
				}
			}
		} else {
			tc.errorList = append(tc.errorList, fmt.Sprintf("All expressions must be of type rune and match the number of identifiers"))
		}

	}

	tc.symbolTable.printTable()
	return nil

}

//------------------------------------

func (tc *TypeChecker) VisitIdListSVDAST(ctx *parser.IdListSVDASTContext) interface{} {
	identifierList := tc.Visit(ctx.IdentifierList()).([]antlr.TerminalNode)
	expresionList := tc.Visit(ctx.ExpressionList()).([]antlr.TerminalNode)

	for i, _ := range identifierList {
		if tc.symbolTable.search(identifierList[i].GetText()) != nil {
			tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", identifierList[i].GetText(), identifierList[i].GetSymbol().GetLine(), identifierList[i].GetSymbol().GetColumn()))
		} else {
			switch expresionList[i].GetSymbol().GetTokenType() {
			case INT:
				tc.symbolTable.insert(identifierList[i], INT)
			case FLOAT:
				tc.symbolTable.insert(identifierList[i], FLOAT)
			case STRING:
				tc.symbolTable.insert(identifierList[i], STRING)
			case RUNE:
				tc.symbolTable.insert(identifierList[i], RUNE)
			case BOOL:
				tc.symbolTable.insert(identifierList[i], BOOL)
			}
		}
	}
	//tc.symbolTable.printTable()
	return nil
}

func (tc *TypeChecker) VisitSingleVarDAST(ctx *parser.SingleVarDASTContext) interface{} {
	tc.Visit(ctx.SingleVarDeclNoExps())
	return nil
}

func (tc *TypeChecker) VisitSingleVarDeclNoExpsAST(ctx *parser.SingleVarDeclNoExpsASTContext) interface{} {
	//fmt.Println("Entre a SingleVarDeclNoExpsAST")
	identificadores := tc.Visit(ctx.IdentifierList()).([]antlr.TerminalNode)
	declType := tc.Visit(ctx.DeclType()).(antlr.TerminalNode)

	for _, id := range identificadores {
		if tc.symbolTable.searchInCurrentLevel(id.GetText()) != nil {
			tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", id.GetText(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn()))
		} else {
			switch declType.GetText() {
			case "int":
				tc.symbolTable.insert(id, INT)
				fmt.Println("(singleVarDeclNoExp) Se insertó em la tabla de simbolos: ", id)
				tc.symbolTable.printTable()
			case "rune":
				tc.symbolTable.insert(id, RUNE)
			case "bool":
				tc.symbolTable.insert(id, BOOL)
			case "string":
				tc.symbolTable.insert(id, STRING)
			case "float":
				tc.symbolTable.insert(id, FLOAT)
			default:
				tc.errorList = append(tc.errorList, fmt.Sprintf("Tipo %s no reconocido", declType))
			}
		}
	}
	tc.symbolTable.printTable()

	switch declType.GetText() {
	case "int":
		return INT
	case "rune":
		return RUNE
	case "bool":
		return BOOL
	case "string":
		return STRING
	case "float":
		return FLOAT
	}

	//tc.symbolTable.printTable()

	return nil
}

func (tc *TypeChecker) VisitSingleTypeDeclTDAST(ctx *parser.SingleTypeDeclTDASTContext) interface{} {
	tc.Visit(ctx.SingleTypeDecl())
	return nil
}

func (tc *TypeChecker) VisitInnerTypeDeclsTDAST(ctx *parser.InnerTypeDeclsTDASTContext) interface{} {
	tc.Visit(ctx.InnerTypeDecls())
	return nil
}
func (tc *TypeChecker) VisitLPTypeDeclTDAST(ctx *parser.LPTypeDeclTDASTContext) interface{} {
	return nil
}
func (tc *TypeChecker) VisitInnerTypeDeclsAST(ctx *parser.InnerTypeDeclsASTContext) interface{} {
	for i := 0; i < len(ctx.AllSingleTypeDecl()); i++ {
		tc.Visit(ctx.SingleTypeDecl(i))
	}
	return nil
}

func (tc *TypeChecker) VisitIdSYDAST(ctx *parser.IdSYDASTContext) interface{} {
	ident := ctx.ID().(antlr.TerminalNode)
	declType := tc.Visit(ctx.DeclType()).(antlr.TerminalNode)

	if tc.symbolTable.search(ident.GetText()) != nil {
		tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", ident.GetText(), ident.GetSymbol().GetLine(), ident.GetSymbol().GetColumn()))

	} else {
		//fmt.Println("Tipo: ", declType.GetText())
		switch declType.GetText() {

		case "int":
			tc.symbolTable.insert(ident, INT)
		case "rune":
			tc.symbolTable.insert(ident, RUNE)
		case "bool":
			tc.symbolTable.insert(ident, BOOL)
		case "string":
			tc.symbolTable.insert(ident, STRING)
		case "float":
			tc.symbolTable.insert(ident, FLOAT)
		case "struct":

			//Verifica si el struct ya está definido
			if tc.symbolTable.searchInCurrentLevel(ident.GetText()) != nil {
				tc.errorList = append(tc.errorList, "Error: Struct '"+ident.GetText()+"' already defined")
			} else {
				tc.symbolTable.insert(ident, STRUCT) // Inserta el identificador en la tabla de símbolos
			}

		default:
			fmt.Println("Es error")
			tc.errorList = append(tc.errorList, fmt.Sprintf("Tipo %s no reconocido", declType))
		}
	}

	tc.symbolTable.printTable()

	return nil
}

func (tc *TypeChecker) VisitFuncDeclAST(ctx *parser.FuncDeclASTContext) interface{} {
	tc.symbolTable.openScope()
	tc.Visit(ctx.FuncFrontDecl())
	tc.Visit(ctx.Block())
	tc.symbolTable.closeScope()

	return nil
}

func (tc *TypeChecker) VisitFuncFFDAST(ctx *parser.FuncFFDASTContext) interface{} {
	ident := ctx.ID().(antlr.TerminalNode)
	var FuncArg []int = nil

	if tc.symbolTable.search(ident.GetText()) != nil {
		tc.errorList = append(tc.errorList, fmt.Sprintf("Existing identifier %s in symbols table, line: %d:%d", ident.GetText(), ident.GetSymbol().GetLine(), ident.GetSymbol().GetColumn()))
	} else {

		if ctx.FuncArgDecls() != nil {
			FuncArg = tc.Visit(ctx.FuncArgDecls()).([]int)

		}

		if ctx.DeclType() != nil {
			switch tc.Visit(ctx.DeclType()).(antlr.TerminalNode).GetSymbol().GetText() {
			case "int":
				tc.symbolTable.insertWithParams(ident, INT, FuncArg)
			case "rune":
				tc.symbolTable.insertWithParams(ident, RUNE, FuncArg)
			case "bool":
				tc.symbolTable.insertWithParams(ident, BOOL, FuncArg)
			case "string":
				tc.symbolTable.insertWithParams(ident, STRING, FuncArg)
			case "float":
				tc.symbolTable.insertWithParams(ident, FLOAT, FuncArg)
			default:
				tc.errorList = append(tc.errorList, fmt.Sprintf("Type %s unrecognize", ctx.DeclType().Accept(tc).(antlr.TerminalNode).GetText()))
			}
		} else {
			tc.symbolTable.insertWithParams(ident, -1, FuncArg)
		}

	}
	tc.symbolTable.printTable()

	return nil
}

func (tc *TypeChecker) VisitFuncArgDeclsAST(ctx *parser.FuncArgDeclsASTContext) interface{} {
	var ListFuncArdDecl []int = nil
	for i := 0; i < len(ctx.AllSingleVarDeclNoExps()); i++ {
		ListFuncArdDecl = append(ListFuncArdDecl, tc.Visit(ctx.SingleVarDeclNoExps(i)).(int))
	}
	return ListFuncArdDecl
}

func (tc *TypeChecker) VisitLpDTAST(ctx *parser.LpDTASTContext) interface{} {
	return tc.Visit(ctx.DeclType()).(antlr.TerminalNode)
}

func (tc *TypeChecker) VisitIdDTAST(ctx *parser.IdDTASTContext) interface{} {
	var list = ctx.ID().(antlr.TerminalNode)
	//fmt.Println("Retorno en idDTAST: ", list)
	return list
}

func (tc *TypeChecker) VisitSliceDeclTypeDTAST(ctx *parser.SliceDeclTypeDTASTContext) interface{} {

	return tc.Visit(ctx.SliceDeclType()).(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitArrayDeclTypeDTAST(ctx *parser.ArrayDeclTypeDTASTContext) interface{} {

	return tc.Visit(ctx.ArrayDeclType()).(antlr.TerminalNode)
}

func (tc *TypeChecker) VisitStructDeclTypeDTAST(ctx *parser.StructDeclTypeDTASTContext) interface{} {
	tc.symbolTable.openScope()
	stru := tc.Visit(ctx.StructDeclType()).(antlr.TerminalNode)
	tc.symbolTable.closeScope()
	return stru
}

func (tc *TypeChecker) VisitLsbSDTAST(ctx *parser.LsbSDTASTContext) interface{} {
	return tc.Visit(ctx.DeclType()).(antlr.TerminalNode)

}
func (tc *TypeChecker) VisitLsbADTAST(ctx *parser.LsbADTASTContext) interface{} {
	return tc.Visit(ctx.DeclType()).(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitStructSDTAST(ctx *parser.StructSDTASTContext) interface{} {
	//Se crea una lista vacia para guardar las declaraciones de los campos del struct
	//var fields []int = nil

	//Se verifica si el struct tiene campos
	if ctx.StructMemDecls() != nil {
		//Se obtiene la lista de campos del struct
		//fields = tc.Visit(ctx.StructMemDecls()).([]int)
	}

	//Imprime la tabla de símbolos
	tc.symbolTable.printTable()
	return ctx.STRUCT().(antlr.TerminalNode)
}

func (tc *TypeChecker) VisitStructMemDeclsAST(ctx *parser.StructMemDeclsASTContext) interface{} {
	var list []int
	for i := 0; i < len(ctx.AllSingleVarDeclNoExps()); i++ {
		list = append(list, tc.Visit(ctx.SingleVarDeclNoExps(i)).(int))

	}
	return list
}

func (tc *TypeChecker) VisitIdentifierListAST(ctx *parser.IdentifierListASTContext) interface{} {
	var list []antlr.TerminalNode
	for i, _ := range ctx.AllID() {
		list = append(list, ctx.ID(i).(antlr.TerminalNode))
	}
	//fmt.Println("Retorno en identifierListAST: ", list)

	return list
}
func (tc *TypeChecker) VisitPrimaryExpressionEAST(ctx *parser.PrimaryExpressionEASTContext) interface{} {

	//fmt.Println("Retorno en PrimaryExpressionEAST: ", expression)
	return tc.Visit(ctx.PrimaryExpression())
}

// ---------------------Funciones para convertir tipos INT, STRING, FLOAT a TerminalNode------------------------------
func CreateTerminalNodeFromInt(value int) antlr.TerminalNode {
	// Convert the integer to a string
	input := fmt.Sprintf("%d", value)

	// Create a CharStream from the string
	charStream := antlr.NewInputStream(input)

	// Create a lexer instance
	lexer := parser.NewgoScanner(charStream)

	// Get the first token (which should be the integer)
	token := lexer.NextToken()

	// Create a TerminalNodeImpl from the token
	return antlr.NewTerminalNodeImpl(token)
}
func CreateTerminalNodeFromString(value string) antlr.TerminalNode {
	// Add double quotes around the string to match the STRING token rule
	input := fmt.Sprintf("\"%s\"", value)

	// Create a CharStream from the string
	charStream := antlr.NewInputStream(input)

	// Create a lexer instance
	lexer := parser.NewgoScanner(charStream)

	// Get the first token (which should be the string)
	token := lexer.NextToken()

	// Create a TerminalNodeImpl from the token
	return antlr.NewTerminalNodeImpl(token)
}
func CreateTerminalNodeFromFloat(value float64) antlr.TerminalNode {
	// Convert the float to a string
	input := fmt.Sprintf("%f", value)

	// Create a CharStream from the string
	charStream := antlr.NewInputStream(input)

	// Create a lexer instance
	lexer := parser.NewgoScanner(charStream)

	// Get the first token (which should be the float)
	token := lexer.NextToken()

	// Create a TerminalNodeImpl from the token
	return antlr.NewTerminalNodeImpl(token)
}

func (tc *TypeChecker) VisitExpressionEAST(ctx *parser.ExpressionEASTContext) interface{} {
	exp1 := tc.Visit(ctx.Expression(0)).(antlr.TerminalNode)
	exp2 := tc.Visit(ctx.Expression(1)).(antlr.TerminalNode)

	tipo1 := exp1.GetSymbol().GetTokenType()
	tipo2 := exp2.GetSymbol().GetTokenType()

	if tc.symbolTable.searchVarInCurrentLevel(exp1.GetText()) == nil {
		fmt.Println("No está en la tabla de simbolos")
	} else {
		fmt.Println("Está en la tabla de simbolos")
		//validar el tipo de la expresion

		//exp := tc.symbolTable.searchVarInCurrentLevel(exp1.GetText()).(*VarIdent)

	}
	if tc.symbolTable.searchVarInCurrentLevel(exp2.GetText()) == nil {
		fmt.Println("No está en la tabla de simbolos")
	} else {
		fmt.Println("Está en la tabla de simbolos")
	}

	if tipo1 == tipo2 {
		switch tipo1 {
		case 77:
			valInt1, err1 := strconv.Atoi(exp1.GetText())
			valInt2, err2 := strconv.Atoi(exp2.GetText())
			if err1 != nil && err2 != nil {
				fmt.Println("Error: Los valores de la expresión no se pudieron convertir a int")
			}
			if ctx.PL() != nil {

				result := valInt1 + valInt2                    //suma las dos expresiones, pero son de tipo int
				resultado := CreateTerminalNodeFromInt(result) //convierte el resultado a tipo antlr.TerminalNode
				fmt.Println("resultado: ", resultado)
				return resultado

			} else if ctx.MIN() != nil {
				result := valInt1 - valInt2
				resultado := CreateTerminalNodeFromInt(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else if ctx.MUL() != nil {
				result := valInt1 * valInt2
				resultado := CreateTerminalNodeFromInt(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else if ctx.DIV() != nil {
				result := valInt1 / valInt2
				resultado := CreateTerminalNodeFromInt(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else if ctx.MOD() != nil {
				result := valInt1 % valInt2
				resultado := CreateTerminalNodeFromInt(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else {
				fmt.Println("La Operación no se puede realizar")
			}
		case 78:
			valFloat1, err1 := strconv.ParseFloat(exp1.GetText(), 78)
			valFloat2, err2 := strconv.ParseFloat(exp2.GetText(), 78)
			if err1 != nil && err2 != nil {
				// Manejo del error
				fmt.Println("Error: Los valores de la expresión no se pudieron convertir a float64")
			}
			if ctx.PL() != nil {
				result := valFloat1 + valFloat2
				fmt.Println("resultado: ", result)
				resultado := CreateTerminalNodeFromFloat(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else if ctx.MIN() != nil {
				result := valFloat1 - valFloat2
				resultado := CreateTerminalNodeFromFloat(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else if ctx.MUL() != nil {
				result := valFloat1 * valFloat2
				resultado := CreateTerminalNodeFromFloat(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else if ctx.DIV() != nil {
				result := valFloat1 / valFloat2
				resultado := CreateTerminalNodeFromFloat(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else {
				fmt.Println("La Operación no se puede realizar")
			}
		case 81:
			//extrae los valores sin dobles comillas dobles
			valString1 := strings.Trim(exp1.GetText(), "\"")
			valString2 := strings.Trim(exp2.GetText(), "\"")
			fmt.Println(valString2, valString1)
			if ctx.PL() != nil {
				result := valString1 + " " + valString2
				resultado := CreateTerminalNodeFromString(result)
				fmt.Println("Resultado de la operación: ", resultado)
				return resultado
			} else {
				fmt.Println("La Operación no se puede realizar")
			}

		case 79:
			fmt.Println("La operación no se puede realizar con el tipo de dato rune o char, debido a que " +
				"sería un tipo string y no coincide con el tipo de la variable")

		}

	} else {
		fmt.Println("La expresion presenta un tipo de dato no permitido para realizar la operación")

	}
	return nil
}

func (tc *TypeChecker) VisitOperatorEAST(ctx *parser.OperatorEASTContext) interface{} {
	//TODO: HACER LA VERIFICACION DE TIPOS
	tc.Visit(ctx.Expression())
	return nil
}

func (tc *TypeChecker) VisitExpressionListAST(ctx *parser.ExpressionListASTContext) interface{} {

	var list []antlr.TerminalNode
	for i := 0; i < len(ctx.AllExpression()); i++ {
		list = append(list, tc.Visit(ctx.Expression(i)).(antlr.TerminalNode))

	}
	fmt.Println("ExpressionListAST (deberia dar el resultado de la operación): ", list)
	return list

}
func (tc *TypeChecker) VisitOperandPEAST(ctx *parser.OperandPEASTContext) interface{} {

	return tc.Visit(ctx.Operand())
}
func (tc *TypeChecker) VisitPrimaryExpSelectorPEAST(ctx *parser.PrimaryExpSelectorPEASTContext) interface{} {
	tc.Visit(ctx.PrimaryExpression())
	tc.Visit(ctx.Selector())
	return nil
}
func (tc *TypeChecker) VisitPrimaryExpIndexPEAST(ctx *parser.PrimaryExpIndexPEASTContext) interface{} {
	tc.Visit(ctx.PrimaryExpression())
	tc.Visit(ctx.Index())
	return nil
}
func (tc *TypeChecker) VisitPrimaryExpArgumentsPEAST(ctx *parser.PrimaryExpArgumentsPEASTContext) interface{} {
	tc.Visit(ctx.PrimaryExpression())
	tc.Visit(ctx.Arguments())
	return nil
}
func (tc *TypeChecker) VisitAppendExpPEAST(ctx *parser.AppendExpPEASTContext) interface{} {
	tc.Visit(ctx.AppendExpression())
	return nil
}
func (tc *TypeChecker) VisitLengthExpPEAST(ctx *parser.LengthExpPEASTContext) interface{} {
	tc.Visit(ctx.LengthExpression())
	return nil
}
func (tc *TypeChecker) VisitCapExpPEAST(ctx *parser.CapExpPEASTContext) interface{} {
	tc.Visit(ctx.CapExpression())
	return nil
}
func (tc *TypeChecker) VisitLiteralOAST(ctx *parser.LiteralOASTContext) interface{} {
	return tc.Visit(ctx.Literal())
}

func (tc *TypeChecker) VisitIdOAST(ctx *parser.IdOASTContext) interface{} {
	return ctx.ID().(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitExpressionOAST(ctx *parser.ExpressionOASTContext) interface{} {
	return tc.Visit(ctx.Expression())
}

func (tc *TypeChecker) VisitIntliteralAST(ctx *parser.IntliteralASTContext) interface{} {
	return ctx.INTLITERAL().(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitFloatliteralAST(ctx *parser.FloatliteralASTContext) interface{} {
	return ctx.FLOATLITERAL().(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitRunliteralAST(ctx *parser.RunliteralASTContext) interface{} {
	return ctx.RUNELITERAL().(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitRawsliteralAST(ctx *parser.RawsliteralASTContext) interface{} {
	return ctx.RAWSTRINGLITERAL().(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitInterpretedliteralAST(ctx *parser.InterpretedliteralASTContext) interface{} {
	return ctx.INTERPRETEDSTRINGLITERAL().(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitIdexAST(ctx *parser.IdexASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *TypeChecker) VisitArgumentsAST(ctx *parser.ArgumentsASTContext) interface{} {
	if (ctx.ExpressionList()) != nil {
		return tc.Visit(ctx.ExpressionList())
	}

	return nil
}
func (tc *TypeChecker) VisitSelectorAST(ctx *parser.SelectorASTContext) interface{} {
	return ctx.ID().(antlr.TerminalNode)
}
func (tc *TypeChecker) VisitAppendExpressionAST(ctx *parser.AppendExpressionASTContext) interface{} {
	tc.Visit(ctx.Expression(0))
	tc.Visit(ctx.Expression(1))
	return nil
}
func (tc *TypeChecker) VisitLengthExpressionAST(ctx *parser.LengthExpressionASTContext) interface{} {
	exp := tc.Visit(ctx.Expression())
	fmt.Println("LengthExpressionAST: ", exp)

	//valida si exp está en la tabla de simbolos
	if tc.symbolTable.searchVarInCurrentLevel(exp.(antlr.TerminalNode).GetText()) == nil {
		fmt.Println("No está en la tabla de simbolos")
	} else {
		fmt.Println("Está en la tabla de simbolos")
		//saca el valor de la expresion de la tabla de simbolos para validar el largo

	}

	return nil
}
func (tc *TypeChecker) VisitCapExpressionAST(ctx *parser.CapExpressionASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *TypeChecker) VisitStatementListAST(ctx *parser.StatementListASTContext) interface{} {
	for i := 0; i < len(ctx.AllStatement()); i++ {
		tc.Visit(ctx.Statement(i))
	}

	return nil
}
func (tc *TypeChecker) VisitBlockAST(ctx *parser.BlockASTContext) interface{} {
	tc.Visit(ctx.StatementList())
	return nil
}
func (tc *TypeChecker) VisitPrintSAST(ctx *parser.PrintSASTContext) interface{} {
	if (ctx.ExpressionList()) != nil {
		tc.Visit(ctx.ExpressionList())
	}
	return nil
}
func (tc *TypeChecker) VisitPrintlnSAST(ctx *parser.PrintlnSASTContext) interface{} {
	if (ctx.ExpressionList()) != nil {
		tc.Visit(ctx.ExpressionList())
	}
	return nil
}
func (tc *TypeChecker) VisitReturnSAST(ctx *parser.ReturnSASTContext) interface{} {
	var methodType *MethodIdent
	if tc.symbolTable.searchMethodInCurrentLevel().(*MethodIdent) != nil {
		methodType = tc.symbolTable.searchMethodInCurrentLevel().(*MethodIdent)
	}

	if methodType.typ == -1 {
		tc.errorList = append(tc.errorList, fmt.Sprintf("La función %s no espera ningún tipo de retorno en la linea %d:%d", methodType.tok.GetText(), methodType.tok.GetSymbol().GetLine(), methodType.tok.GetSymbol().GetColumn()))
	} else {
		if ctx.Expression() != nil {
			if reflect.TypeOf(tc.Visit(ctx.Expression())) != nil && tc.Visit(ctx.Expression()).(int) != methodType.typ {
				tc.errorList = append(tc.errorList, fmt.Sprintf("Tipo de la expresion que quiere retornar no es igual al tipo de retorno del metodo %s, en la linea %d:%d", tc.symbolTable.searchMethodInCurrentLevel().(*MethodIdent).tok.GetText(), tc.symbolTable.searchMethodInCurrentLevel().(*MethodIdent).tok.GetSymbol().GetLine(), tc.symbolTable.searchMethodInCurrentLevel().(*MethodIdent).tok.GetSymbol().GetColumn()))
			}

		}
	}

	return nil
}
func (tc *TypeChecker) VisitBreakSAST(ctx *parser.BreakSASTContext) interface{} {
	return nil
}
func (tc *TypeChecker) VisitContinueSAST(ctx *parser.ContinueSASTContext) interface{} {
	return nil
}
func (tc *TypeChecker) VisitSimpleStatementSAST(ctx *parser.SimpleStatementSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	return nil
}
func (tc *TypeChecker) VisitBlockSAST(ctx *parser.BlockSASTContext) interface{} {
	tc.Visit(ctx.Block())
	return nil
}
func (tc *TypeChecker) VisitSwitchSAST(ctx *parser.SwitchSASTContext) interface{} {
	tc.Visit(ctx.Switch_())
	return nil
}
func (tc *TypeChecker) VisitIfStatementSAST(ctx *parser.IfStatementSASTContext) interface{} {
	tc.Visit(ctx.IfStatement())
	return nil
}
func (tc *TypeChecker) VisitLoopSAST(ctx *parser.LoopSASTContext) interface{} {
	tc.Visit(ctx.Loop())
	return nil
}
func (tc *TypeChecker) VisitTypeDeclSAST(ctx *parser.TypeDeclSASTContext) interface{} {
	tc.Visit(ctx.TypeDecl())
	return nil
}
func (tc *TypeChecker) VisitVariableDeclSAST(ctx *parser.VariableDeclSASTContext) interface{} {
	tc.Visit(ctx.VariableDecl())
	return nil
}
func (tc *TypeChecker) VisitEpsilonSSAST(ctx *parser.EpsilonSSASTContext) interface{} {
	return nil
}
func (tc *TypeChecker) VisitExpressionINCSSAST(ctx *parser.ExpressionINCSSASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *TypeChecker) VisitExpressionDECSSAST(ctx *parser.ExpressionDECSSASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *TypeChecker) VisitExpressionEpsilonSSAST(ctx *parser.ExpressionEpsilonSSASTContext) interface{} {

	tc.Visit(ctx.Expression())
	return nil
}
func (tc *TypeChecker) VisitAssigmentStatementSSAST(ctx *parser.AssigmentStatementSSASTContext) interface{} {
	tc.Visit(ctx.AssignmentStatement())
	return nil
}
func (tc *TypeChecker) VisitExpListSSAST(ctx *parser.ExpListSSASTContext) interface{} {

	tc.Visit(ctx.ExpressionList(0))
	tc.Visit(ctx.ExpressionList(1))
	return nil
}
func (tc *TypeChecker) VisitExpListASAST(ctx *parser.ExpListASASTContext) interface{} {

	var variableList = ctx.ExpressionList(0).Accept(tc).([]antlr.TerminalNode)
	var valorList = ctx.ExpressionList(1).Accept(tc).([]antlr.TerminalNode)
	var variableTemp *VarIdent

	for i := 0; i < len(variableList); i++ {
		if tc.symbolTable.searchInCurrentLevel(variableList[i].(antlr.TerminalNode).GetText()) == nil {
			tc.errorList = append(tc.errorList, fmt.Sprintf("Variable %s no declarada en la linea %d:%d", variableList[i].(antlr.TerminalNode).GetText(), variableList[i].(antlr.TerminalNode).GetSymbol().GetLine(), variableList[i].(antlr.TerminalNode).GetSymbol().GetColumn()))
		} else {
			variableTemp = tc.symbolTable.searchInCurrentLevel(variableList[i].(antlr.TerminalNode).GetText()).(*VarIdent)
			if valorList[i].(antlr.TerminalNode).GetSymbol().GetTokenType() != variableTemp.typ {
				tc.errorList = append(tc.errorList, fmt.Sprintf("Tipo de la expresion %s no es igual al tipo de la variable %s", valorList[i].(antlr.TerminalNode).GetText(), variableList[i].(antlr.TerminalNode).GetText()))
			}
		}
	}
	return nil
}
func (tc *TypeChecker) VisitExpASAST(ctx *parser.ExpASASTContext) interface{} {

	tc.Visit(ctx.Expression(0))
	tc.Visit(ctx.Expression(1))

	return nil
}
func (tc *TypeChecker) VisitIsExpressionBlockISAST(ctx *parser.IsExpressionBlockISASTContext) interface{} {
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	return nil
}
func (tc *TypeChecker) VisitIsExpBlockIsISAST(ctx *parser.IsExpBlockIsISASTContext) interface{} {
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	tc.Visit(ctx.IfStatement())
	return nil
}
func (tc *TypeChecker) VisitIsExpBlockISAST(ctx *parser.IsExpBlockISASTContext) interface{} {
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block(0))
	tc.Visit(ctx.Block(1))
	return nil
}

func (tc *TypeChecker) VisitIsSSExpBlockifSAST(ctx *parser.IsSSExpBlockifSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	return nil
}
func (tc *TypeChecker) VisitIsSSExpBlockIfSAST(ctx *parser.IsSSExpBlockifSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	tc.Visit(ctx.IfStatement())
	return nil
}
func (tc *TypeChecker) VisitIsSSExpBlockBlockAST(ctx *parser.IsSSExpBlockBlockASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())

	tc.Visit(ctx.Block(0))
	tc.Visit(ctx.Block(1))
	return nil
}
func (tc *TypeChecker) VisitFBlockLAST(ctx *parser.FBlockLASTContext) interface{} {
	tc.Visit(ctx.Block())
	return nil
}
func (tc *TypeChecker) VisitFExpBlockLAST(ctx *parser.FExpBlockLASTContext) interface{} {
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	return nil
}
func (tc *TypeChecker) VisitFSimpStateExpSBlockLAST(ctx *parser.FSimpStateExpSBlockLASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement(0))
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.SimpleStatement(1))
	tc.Visit(ctx.Block())

	return nil
}
func (tc *TypeChecker) VisitFSimpStateSimpStateBlockLAST(ctx *parser.FSimpStateSimpStateBlockLASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement(0))
	tc.Visit(ctx.SimpleStatement(1))
	tc.Visit(ctx.Block())

	return nil
}

func (tc *TypeChecker) VisitSSimpleStatSAST(ctx *parser.SSimpleStatSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *TypeChecker) VisitSExpressionSAST(ctx *parser.SExpressionSASTContext) interface{} {
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *TypeChecker) VisitSSimpleStatExpCCListSAST(ctx *parser.SSimpleStatExpCCListSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *TypeChecker) VisitSBlockSAST(ctx *parser.SBlockSASTContext) interface{} {
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *TypeChecker) VisitEpsilonECCLAST(ctx *parser.EpsilonECCLASTContext) interface{} {
	//tc.Visit(ctx.Epsilon())
	return nil
}
func (tc *TypeChecker) VisitExpCaseClauseECCLAST(ctx *parser.ExpCaseClauseECCLASTContext) interface{} {
	tc.Visit(ctx.ExpressionCaseClause())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *TypeChecker) VisitExpSwitchCaseECCAST(ctx *parser.ExpSwitchCaseECCASTContext) interface{} {
	tc.Visit(ctx.ExpressionSwitchCase())
	tc.Visit(ctx.StatementList())
	return nil
}
func (tc *TypeChecker) VisitCaseESCAST(ctx *parser.CaseESCASTContext) interface{} {
	tc.Visit(ctx.ExpressionList())
	return nil
}
func (tc *TypeChecker) VisitDefaultESCAST(ctx *parser.DefaultESCASTContext) interface{} {
	return ctx.DEFAULT()
}

func (tc *TypeChecker) VisitEpsilonAST(ctx *parser.EpsilonASTContext) interface{} {
	return nil
}
