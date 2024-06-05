package encoder

import (
	"MiniGO/compiler/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"strconv"
)

type Encoder struct {
	*parser.BasegoParserVisitor
	mainModule *ir.Module
	aquiStr    *ir.Global
	printf     *ir.Func
	puts       *ir.Func
	//vars       map[string]value.Value
}

func NewEncoder() *Encoder {
	return &Encoder{
		BasegoParserVisitor: &parser.BasegoParserVisitor{},
		mainModule:          ir.NewModule(),
		aquiStr:             nil,
		printf:              nil,
		puts:                nil,
		//vars:                make(map[string]value.Value),
	}
}

func (v *Encoder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Encoder) VisitChildren(tree antlr.RuleNode) any {
	var result any
	n := tree.GetChildCount()
	for i := 0; i < n; i++ {
		c := tree.GetChild(i)
		val := c.(antlr.ParseTree)
		result2 := result
		result = v.Visit(val)
		if result == nil {
			result = result2
		}
	}
	return result
}

var funcActual *ir.Func //se usa para la función que se esté visitando de turno... solo una a la vez

var blocksFunActual Stack               //se usa almacenar los bloques de la función que se esté visitando de turno
var variablesLocales BlockVariableTable //se usa para almacenar las variables locales del bloque de turno. Variables X Bloque

// var blocksFunActual []blockInfo
var scope int = 0
var vActual value.Value

//esta última tabla de variables locales tiene que pensarse mejor porque funciona bien cuando hablamos de identificadores declarados en el mismo bloque donde se usan
//si estamos hablando de identificadores más globales al bloque donde se usan, entonces no creo que funcione tan bien o al menos debería de pensarse mejor y cambiarse

func (tc *Encoder) VisitRootAST(ctx *parser.RootASTContext) interface{} {
	//se inicializa la estructura de variables locales
	variablesLocales = *NewBlockVariableTable()
	tc.aquiStr = tc.mainModule.NewGlobalDef("aquiStr", constant.NewCharArrayFromString("%d\x00"))
	tc.printf = tc.mainModule.NewFunc("printf", types.I32, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I32))
	tc.puts = tc.mainModule.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
	tc.Visit(ctx.TopDeclarationList())
	return tc.mainModule

}

func (tc *Encoder) VisitTopDeclarationListAST(ctx *parser.TopDeclarationListASTContext) interface{} {

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
	//return tc.VisitChildren(ctx)
}
func (tc *Encoder) VisitSingleVarVDAST(ctx *parser.SingleVarVDASTContext) interface{} {
	tc.Visit(ctx.SingleVarDecl())
	return nil
	//return tc.VisitChildren(ctx)
}
func (tc *Encoder) VisitInnerVarVDAST(ctx *parser.InnerVarVDASTContext) interface{} {
	//return tc.VisitChildren(ctx)
	tc.Visit(ctx.InnerVarDecls())
	return nil
}
func (tc *Encoder) VisitLpVDAST(ctx *parser.LpVDASTContext) interface{} {
	return nil
}
func (tc *Encoder) VisitInnerVarDeclsAST(ctx *parser.InnerVarDeclsASTContext) interface{} {
	for i := 0; i < len(ctx.AllSingleVarDecl()); i++ {
		tc.Visit(ctx.SingleVarDecl(i))
	}
	return nil
}

func (v *Encoder) VisitIdListDTSVDAST(ctx *parser.IdListDTSVDASTContext) interface{} {

	println("\n/*** DECLARACION DE VARIABLES CON ASIGNACIÓN ***/")

	blockactual, _ := blocksFunActual.Peek() //se obtiene el bloque actual para saber donde se va a declarar la variable

	//se obtienen todos los identificadores de la declaración para declararlos todos
	ids := ctx.IdentifierList().(*parser.IdentifierListASTContext).AllID()

	val := ctx.ExpressionList().(*parser.ExpressionListASTContext).AllExpression() //Se obtienen los valores de las expresiones

	if types.IsArray(v.Visit(ctx.DeclType()).(types.Type)) {

		//Se crea el array con los valores de las expresiones
		arrayType := v.Visit(ctx.DeclType()).(types.Type)

		//Se crea un array con los valores de las expresiones segun la cantidad de elementos NewArray(2, constant.NewInt(types.I32, 10), constant.NewInt(types.I32, 20), constant.NewInt(types.I32, 30))
		arrayTypeP := types.NewArray(arrayType.(*types.ArrayType).Len, arrayType.(*types.ArrayType).ElemType)

		//array := constant.NewArray(arrayTypeP, nil) //Se crea un array vacío
		lisConst := make([]constant.Constant, 0)

		for i := 0; i < len(val); i++ {
			valExpression := v.Visit(val[i]).(value.Value) //Esto devuelve un string
			lisConst = append(lisConst, valExpression.(*constant.Int))
		}

		//Se crea el array con los valores de las expresiones segun la cantidad de elementosarr := constant.NewArray(arrayType, constant.NewInt(types.I32, 10), constant.NewInt(types.I32, 20), constant.NewInt(types.I32, 30))
		array := constant.NewArray(arrayTypeP, lisConst...)
		vActual = blockactual.NewAlloca(arrayTypeP)
		blockactual.NewStore(array, vActual)
		variablesLocales.AddVariable(blockactual, ids[0].GetText(), vActual)

		return nil

	}

	for i := 0; i < len(ids); i++ { //Se recorre la lista de identificadores

		idType := v.Visit(ctx.DeclType()).(types.Type) //Se obtiene el tipo de dato de la variable

		//Se recorre la lista de valores de las expresiones y se obtienen los valores para asignarlos a las variables
		valExpression := v.Visit(val[i]).(value.Value) //Esto devuelve un string

		println("\t-- Bloque actual: ", blockactual)
		println("\t-- Variable con asignación: ", ids[i].GetText())
		println("\t-- Tipo de dato: ", idType.String())
		println("\t-- Valor de la expresión: ", valExpression.Ident())

		//Se verifica si el bloque actual es igual a nulo, y si es así, se verifica si es el parametro de una función
		if blockactual == nil {

			if funcActual == nil { //Si no es el parametro de una función, entonces es una variable global

				switch idType {
				case types.I32: //Int
					// Convertir valExpression a int
					intValue, _ := strconv.ParseInt(valExpression.Ident(), 10, 64)
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I32, intValue))
				case types.Float:
					// Convertir valExpression a float
					floatValue, _ := strconv.ParseFloat(valExpression.Ident(), 64)
					println("Valor de la expresión a asigfdgdfgdf: ", floatValue)
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewFloat(types.Float, floatValue))
				case types.I8Ptr: //String
					// Convertir valExpression a string
					println("Valor de la expresión a asignar EN LA VARIABLE: ", valExpression.String())
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewCharArrayFromString(valExpression.Ident()))
				case types.I1: //Bool
					boolValue, _ := strconv.ParseInt(valExpression.Ident(), 10, 64)
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I1, boolValue))
				case types.I8: //Rune
					// Convertir valExpression a int
					runeValue, _ := strconv.ParseInt(valExpression.Ident(), 10, 64)
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I8, runeValue))
				default:
					if types.IsPointer(valExpression.Type()) {
						vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewZeroInitializer(valExpression.Type()))
					} else {
						vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewZeroInitializer(idType))
					}
				}
			}

		} else { //Si no es nulo, entonces es una variable local de una función

			println("Variable local con asignación: ", ids[i].GetText())

			switch idType {
			case types.I32: //Int
				intValue, _ := strconv.ParseInt(valExpression.Ident(), 10, 64)
				vActual = blockactual.NewAlloca(types.I32)
				//Almacenar el valor de la expresión en la dirección de memoria de la variable
				blockactual.NewStore(constant.NewInt(types.I32, intValue), vActual)
				//Se agrega la variable a la tabla de variables locales
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.Float:
				floatValue, _ := strconv.ParseFloat(valExpression.Ident(), 64)
				vActual = blockactual.NewAlloca(types.Float)
				blockactual.NewStore(constant.NewFloat(types.Float, floatValue), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.I8Ptr: //String
				stringValue := valExpression.String()
				println("Valor de la expresión a asignar EN LA VARIABLE: ", stringValue)
				strLen := len(stringValue) // Longitud de la cadena

				// Crear un array local para almacenar la cadena
				strType := types.NewArray(uint64(strLen), types.I8) // Crear un array de bytes (cada byte es un caracter)
				strAlloca := blockactual.NewAlloca(strType)         // Crear un array de bytes en la memoria

				// Inicializar la cadena en la variable local 'strAlloca'
				for i, char := range stringValue {
					elemPtr := blockactual.NewGetElementPtr(strType, strAlloca, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, int64(i)))
					blockactual.NewStore(constant.NewInt(types.I8, int64(char)), elemPtr)
				}

				// Obtener un puntero a la cadena para pasarlo a 'puts'
				ptr := blockactual.NewGetElementPtr(strType, strAlloca, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))

				//Se agrega la variable a la tabla de variables locales
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), ptr)

				//println("ejrfjeoirjoeValor de la expresión a asignar: ", stringValue)
				//vActual = blockactual.NewAlloca(types.I8Ptr)
				//blockactual.NewStore(constant.NewCharArrayFromString(stringValue), vActual)
				//variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.I1: //Bool
				boolValue, _ := strconv.ParseInt(valExpression.Ident(), 10, 64)
				vActual = blockactual.NewAlloca(types.I1)
				blockactual.NewStore(constant.NewInt(types.I1, boolValue), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.I8: //Rune
				runeValue, _ := strconv.ParseInt(valExpression.Ident(), 10, 64)
				vActual = blockactual.NewAlloca(types.I8)
				blockactual.NewStore(constant.NewInt(types.I8, runeValue), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			default:
				if types.IsPointer(idType) { //Si es un puntero
					vActual = blockactual.NewAlloca(valExpression.Type())
					blockactual.NewStore(valExpression, vActual)
					variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
				} else {
					vActual = blockactual.NewAlloca(idType)
					blockactual.NewStore(constant.NewZeroInitializer(idType), vActual)
					variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)

				}
			}
		}
	}

	return nil
}

func (tc *Encoder) VisitIdListSVDAST(ctx *parser.IdListSVDASTContext) interface{} {
	idlist := tc.Visit(ctx.IdentifierList()).([]string)
	explist := tc.Visit(ctx.ExpressionList()).([]value.Value)

	blockactual, err := blocksFunActual.Peek()
	if err != nil { //Si el bloque actual es nulo y la función actual no lo es, entonces se guarda en la lista de parametros de la función
		fmt.Println("No existe el bloque")
	} else {
		for i, _ := range idlist {
			variable := blockactual.NewAlloca(explist[i].Type())
			blockactual.NewStore(explist[i], variable)
			variablesLocales.AddVariable(blockactual, idlist[i], variable)
		}
	}
	return nil
}

func (tc *Encoder) VisitSingleVarDAST(ctx *parser.SingleVarDASTContext) interface{} {
	tc.Visit(ctx.SingleVarDeclNoExps())
	return nil
}

func (v *Encoder) VisitSingleVarDeclNoExpsAST(ctx *parser.SingleVarDeclNoExpsASTContext) interface{} {

	// Aqui se debe de evaluar donde se realiza la declaración de las variables
	// puede ser una declaración de variables globales, locales y como argumentos de funciones

	blockactual, _ := blocksFunActual.Peek() //se obtiene el bloque actual para saber donde se va a declarar la variable

	//se obtienen todos los identificadores de la declaración para declararlos todos
	ids := ctx.IdentifierList().(*parser.IdentifierListASTContext).AllID()

	mapDecl := make(map[types.Type][]string) //se crea un mapa para guardar los identificadores y su tipo de dato

	for i := 0; i < len(ids); i++ {

		idType := v.Visit(ctx.DeclType()).(types.Type)

		//Se verifica si el bloque actual es igual a nulo, y si es así, se verifica si es el parametro de una función
		if blockactual == nil {

			if funcActual == nil { //Si no es el parametro de una función, entonces es una variable global

				println("Variable global: ", ids[i].GetText())

				switch idType {
				case types.I32: //Int
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I32, 0))
				case types.Float:
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewFloat(types.Float, 0))
				case types.I8Ptr: //String
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewCharArrayFromString(""))
				case types.I1: //Bool
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I1, 0))
				case types.I8: //Rune
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewInt(types.I8, 0))
				default:
					vActual = v.mainModule.NewGlobalDef(ids[i].GetText(), constant.NewZeroInitializer(idType))
				}

			} else { //Si es el parametro de una función, entonces se guarda en la lista de parametros de la función

				println("Parametro de función: ", ids[i].GetText())
				mapDecl[idType] = append(mapDecl[idType], ids[i].GetText())
			}

		} else { //Si no es nulo, entonces es una variable local de una función

			println("Variable local: ", ids[i].GetText())

			switch idType {
			case types.I32: //Int
				vActual = blockactual.NewAlloca(types.I32)
				blockactual.NewStore(constant.NewInt(types.I32, 0), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.Float:
				vActual = blockactual.NewAlloca(types.Float)
				blockactual.NewStore(constant.NewFloat(types.Float, 0), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.I8Ptr: //String
				vActual = blockactual.NewAlloca(types.I8Ptr)
				blockactual.NewStore(constant.NewCharArrayFromString(""), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.I1: //Bool
				vActual = blockactual.NewAlloca(types.I1)
				blockactual.NewStore(constant.NewInt(types.I1, 0), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			case types.I8: //Rune
				vActual = blockactual.NewAlloca(types.I8)
				blockactual.NewStore(constant.NewInt(types.I8, 0), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			default:
				vActual = blockactual.NewAlloca(idType)
				blockactual.NewStore(constant.NewZeroInitializer(idType), vActual)
				variablesLocales.AddVariable(blockactual, ids[i].GetText(), vActual)
			}

		}
	}

	if blockactual == nil && funcActual != nil { //Si el bloque actual es nulo y la función actual no lo es, entonces se guarda en la lista de parametros de la función
		fmt.Println("este es el mapa de parametros: ", mapDecl)
		return mapDecl
	} else {
		return nil
	}
}

func (tc *Encoder) VisitSingleTypeDeclTDAST(ctx *parser.SingleTypeDeclTDASTContext) interface{} {
	//tc.Visit(ctx.SingleTypeDecl())
	//return nil
	return tc.VisitChildren(ctx)

}

func (tc *Encoder) VisitInnerTypeDeclsTDAST(ctx *parser.InnerTypeDeclsTDASTContext) interface{} {
	//tc.Visit(ctx.InnerTypeDecls())
	//return nil
	return tc.VisitChildren(ctx)

}
func (tc *Encoder) VisitLPTypeDeclTDAST(ctx *parser.LPTypeDeclTDASTContext) interface{} {
	return nil
}
func (tc *Encoder) VisitInnerTypeDeclsAST(ctx *parser.InnerTypeDeclsASTContext) interface{} {

	return tc.VisitChildren(ctx)
}

func (tc *Encoder) VisitIdSYDAST(ctx *parser.IdSYDASTContext) interface{} {

	idFunc := tc.Visit(ctx.DeclType()).(string)
	nameFunc := ctx.ID().GetText()

	switch idFunc {
	case "int":
		funcActual = tc.mainModule.NewFunc(nameFunc, types.I32)
	case "float":
		funcActual = tc.mainModule.NewFunc(nameFunc, types.Float)
	case "string":
		funcActual = tc.mainModule.NewFunc(nameFunc, types.I8Ptr)
	case "bool":
		funcActual = tc.mainModule.NewFunc(nameFunc, types.I1)
	default:
		funcActual = tc.mainModule.NewFunc(nameFunc, types.Void)
	}

	return tc.VisitChildren(ctx)
}

func (tc *Encoder) VisitFuncDeclAST(ctx *parser.FuncDeclASTContext) interface{} {

	tc.Visit(ctx.FuncFrontDecl())
	tc.Visit(ctx.Block())

	blocksFunActual.Pop()
	funcActual = nil

	return nil
}

func (tc *Encoder) VisitFuncFFDAST(ctx *parser.FuncFFDASTContext) interface{} {
	var idFunc types.Type
	if ctx.DeclType() != nil {
		idFunc = tc.Visit(ctx.DeclType()).(types.Type)
	} else {
		idFunc = types.Void

	}

	nameFunc := ctx.ID().GetText()

	switch idFunc {
	case types.I32:
		funcActual = tc.mainModule.NewFunc(nameFunc, types.I32)
	case types.Float:
		funcActual = tc.mainModule.NewFunc(nameFunc, types.Float)
	case types.I8Ptr:
		funcActual = tc.mainModule.NewFunc(nameFunc, types.I8Ptr)
	case types.I1:
		funcActual = tc.mainModule.NewFunc(nameFunc, types.I1)
	default:
		funcActual = tc.mainModule.NewFunc(nameFunc, types.Void)
	}

	//Se verifica si la función tiene argumentos
	if ctx.FuncArgDecls() != nil {

		args := tc.Visit(ctx.FuncArgDecls()).([]map[types.Type][]string) //Se obtienen los argumentos de la función
		//Recorrer la lista de mapas de argumentos
		for _, arg := range args {
			//Recorrer el mapa de argumentos
			for key, value := range arg {
				//Recorrer la lista de argumentos
				for _, val := range value {
					//Se crea un nuevo argumento
					funcActual.Params = append(funcActual.Params, ir.NewParam(val, key))
				}
			}
		}
	} else {
		fmt.Println("No vienen Argumentos")
	}
	return nil

}

func (tc *Encoder) VisitFuncArgDeclsAST(ctx *parser.FuncArgDeclsASTContext) interface{} {

	var args []map[types.Type][]string //Se crea una lista de mapas de argumentos

	//Se obtiene la lista de argumentos
	for i := 0; i < len(ctx.AllSingleVarDeclNoExps()); i++ {
		args = append(args, tc.Visit(ctx.SingleVarDeclNoExps(i)).(map[types.Type][]string))
	}
	return args

}

func (tc *Encoder) VisitLpDTAST(ctx *parser.LpDTASTContext) interface{} {
	//return tc.Visit(ctx.DeclType()).(antlr.TerminalNode)
	return tc.VisitChildren(ctx)
}

func (tc *Encoder) VisitIdDTAST(ctx *parser.IdDTASTContext) interface{} {
	//Se retorna el tipo de dato del identificador
	var idType types.Type

	switch ctx.ID().GetText() {
	case "int":
		idType = types.I32
	case "float":
		idType = types.Float
	case "string":
		idType = types.I8Ptr
	case "bool":
		idType = types.I1
	case "rune":
		idType = types.I8
	default:
		idType = types.Void
	}

	return idType
}

func (tc *Encoder) VisitSliceDeclTypeDTAST(ctx *parser.SliceDeclTypeDTASTContext) interface{} {

	//return tc.Visit(ctx.SliceDeclType()).(antlr.TerminalNode)
	return tc.VisitChildren(ctx)
}
func (tc *Encoder) VisitArrayDeclTypeDTAST(ctx *parser.ArrayDeclTypeDTASTContext) interface{} {

	return tc.Visit(ctx.ArrayDeclType())

}

func (tc *Encoder) VisitStructDeclTypeDTAST(ctx *parser.StructDeclTypeDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
	//return nil
}

func (tc *Encoder) VisitLsbSDTAST(ctx *parser.LsbSDTASTContext) interface{} {
	//return tc.Visit(ctx.DeclType()).(antlr.TerminalNode)
	return tc.VisitChildren(ctx)

}
func (tc *Encoder) VisitLsbADTAST(ctx *parser.LsbADTASTContext) interface{} {

	idType := tc.Visit(ctx.DeclType()).(types.Type)
	size := ctx.INTLITERAL()

	//convertir a entero
	sizeInt, _ := strconv.Atoi(size.GetText())

	//se crea un nuevo tipo de array
	arrayType := types.NewArray(uint64(sizeInt), idType)
	return arrayType

}
func (tc *Encoder) VisitStructSDTAST(ctx *parser.StructSDTASTContext) interface{} {
	return tc.VisitChildren(ctx)
	//return ctx.STRUCT().(antlr.TerminalNode)
}

func (tc *Encoder) VisitStructMemDeclsAST(ctx *parser.StructMemDeclsASTContext) interface{} {

	return nil
}

func (tc *Encoder) VisitIdentifierListAST(ctx *parser.IdentifierListASTContext) interface{} {
	var list []string
	for i := 0; i < len(ctx.AllID()); i++ {
		list = append(list, ctx.ID(i).GetText())
	}
	return list

}
func (tc *Encoder) VisitPrimaryExpressionEAST(ctx *parser.PrimaryExpressionEASTContext) interface{} {
	//fmt.Println("Primary Expression: ", tc.Visit(ctx.PrimaryExpression()))
	return tc.Visit(ctx.PrimaryExpression())
}
func (tc *Encoder) VisitExpressionEAST(ctx *parser.ExpressionEASTContext) interface{} {
	exp1 := tc.Visit(ctx.Expression(0)).(value.Value)
	exp2 := tc.Visit(ctx.Expression(1)).(value.Value)

	blockActual, _ := blocksFunActual.Peek()

	if types.IsPointer(exp1.Type()) {
		exp1 = blockActual.NewLoad(exp1.Type().(*types.PointerType).ElemType, exp1)
	}
	if types.IsPointer(exp2.Type()) {
		exp2 = blockActual.NewLoad(exp2.Type().(*types.PointerType).ElemType, exp2)
	}
	//preguntar si son punteros y si lo son cargarlos
	/*if types.IsPointer(exp1.Type()) {
		exp1 = blockActual.NewLoad(exp1.Type().(*types.PointerType).ElemType, exp1)
	}
	if types.IsPointer(exp2.Type()) {
		exp2 = blockActual.NewLoad(exp2.Type().(*types.PointerType).ElemType, exp2)
	}*/

	if ctx.PL() != nil {
		fmt.Println("Valores: ", exp1, exp2)
		suma := blockActual.NewAdd(exp1, exp2)
		fmt.Println("Suma: ", suma)
		return suma
	} else if ctx.MIN() != nil {
		resta := blockActual.NewSub(exp1, exp2)
		return resta
	} else if ctx.MUL() != nil {
		mul := blockActual.NewMul(exp1, exp2)
		return mul
	} else if ctx.DIV() != nil {
		div := blockActual.NewSDiv(exp1, exp2)
		return div
	} else if ctx.MOD() != nil {
		porc := blockActual.NewSRem(exp1, exp2)
		return porc
	} else if ctx.EQ() != nil {
		eq := blockActual.NewICmp(enum.IPredEQ, exp1, exp2)
		return eq
	} else if ctx.GT() != nil {
		gt := blockActual.NewICmp(enum.IPredSGT, exp1, exp2)
		return gt
	} else if ctx.LT() != nil {
		lt := blockActual.NewICmp(enum.IPredSLT, exp1, exp2)
		return lt
	} else if ctx.GTE() != nil {
		gte := blockActual.NewICmp(enum.IPredSGE, exp1, exp2)
		return gte
	} else if ctx.LTE() != nil {
		lte := blockActual.NewICmp(enum.IPredSLE, exp1, exp2)
		return lte
	}

	return nil
}
func (tc *Encoder) VisitOperatorEAST(ctx *parser.OperatorEASTContext) interface{} {

	tc.Visit(ctx.Expression())
	return nil
}
func (tc *Encoder) VisitExpressionListAST(ctx *parser.ExpressionListASTContext) interface{} {
	//Lista de valores
	var valList []value.Value

	//Se recorre la lista de expresiones
	for i := 0; i < len(ctx.AllExpression()); i++ {
		//Se obtiene el valor de la expresión
		valList = append(valList, tc.Visit(ctx.Expression(i)).(value.Value))
	}
	fmt.Println("ValList en ExpressionList: ", valList)
	return valList //Se retorna la lista de valores

}
func (tc *Encoder) VisitOperandPEAST(ctx *parser.OperandPEASTContext) interface{} {
	//fmt.Println("visita en operand: ", tc.Visit(ctx.Operand()))
	return tc.Visit(ctx.Operand())
}
func (tc *Encoder) VisitPrimaryExpSelectorPEAST(ctx *parser.PrimaryExpSelectorPEASTContext) interface{} {
	tc.Visit(ctx.PrimaryExpression())
	tc.Visit(ctx.Selector())
	return nil
}
func (tc *Encoder) VisitPrimaryExpIndexPEAST(ctx *parser.PrimaryExpIndexPEASTContext) interface{} {
	tc.Visit(ctx.PrimaryExpression())
	tc.Visit(ctx.Index())
	return nil
}
func (tc *Encoder) VisitPrimaryExpArgumentsPEAST(ctx *parser.PrimaryExpArgumentsPEASTContext) interface{} {
	tc.Visit(ctx.PrimaryExpression())
	tc.Visit(ctx.Arguments())
	return nil
}
func (tc *Encoder) VisitAppendExpPEAST(ctx *parser.AppendExpPEASTContext) interface{} {
	tc.Visit(ctx.AppendExpression())
	return nil
}
func (tc *Encoder) VisitLengthExpPEAST(ctx *parser.LengthExpPEASTContext) interface{} {

	return tc.Visit(ctx.LengthExpression())
}
func (tc *Encoder) VisitCapExpPEAST(ctx *parser.CapExpPEASTContext) interface{} {
	tc.Visit(ctx.CapExpression())
	return nil
}
func (tc *Encoder) VisitLiteralOAST(ctx *parser.LiteralOASTContext) interface{} {
	//fmt.Println("Visit de literalOAST: ", tc.Visit(ctx.Literal()))
	return tc.Visit(ctx.Literal())

}

func (tc *Encoder) VisitIdOAST(ctx *parser.IdOASTContext) interface{} {

	//Siempre va a retornar un nombre de variable (ya sea que esté declarado o no)
	nameVar := ctx.ID().GetText()

	blockActual, err := blocksFunActual.Peek()
	if err != nil {
		fmt.Print("No se encuentra el bloque en IdOAst")
	} else {
		if val, found := variablesLocales.GetVariable(blockActual, nameVar); found { //Si la variable está en la tabla de variables locales
			fmt.Println("IDOAST: ", val)
			//fmt.Println("IDOSAST", nameVar)
			return val
			//return nameVar
		} else if val, found := variablesLocales.GetVariableAllBlock(nameVar); found { //Si la variable está en la tabla de variables locales
			fmt.Println("IDOAST: ", val)
			//fmt.Println("IDOSAST", nameVar)
			return val
			//return nameVar
		} else {
			for _, param := range funcActual.Params {
				//fmt.Print("Está buscando las variables", param.Name())
				if param.Name() == nameVar {
					return param
				}
			}
		}
	}
	return nil
}
func (tc *Encoder) VisitExpressionOAST(ctx *parser.ExpressionOASTContext) interface{} {
	return tc.Visit(ctx.Expression())
}

func (tc *Encoder) VisitIntliteralAST(ctx *parser.IntliteralASTContext) interface{} {
	valReturn, _ := constant.NewIntFromString(types.I32, ctx.INTLITERAL().GetText())
	fmt.Println("Valreturn de intliteral: ", valReturn)
	return valReturn
}
func (tc *Encoder) VisitFloatliteralAST(ctx *parser.FloatliteralASTContext) interface{} {
	//return ctx.FLOATLITERAL().GetText()
	valReturn, _ := constant.NewFloatFromString(types.Float, ctx.FLOATLITERAL().GetText())
	return valReturn
}
func (tc *Encoder) VisitRunliteralAST(ctx *parser.RunliteralASTContext) interface{} {
	//return ctx.RUNELITERAL().GetText()
	valReturn, _ := constant.NewIntFromString(types.I8, ctx.RUNELITERAL().GetText())
	return valReturn
}
func (tc *Encoder) VisitRawsliteralAST(ctx *parser.RawsliteralASTContext) interface{} {
	//return ctx.RAWSTRINGLITERAL().GetText()
	valReturn := constant.NewCharArrayFromString(ctx.RAWSTRINGLITERAL().GetText())
	return valReturn
}
func (tc *Encoder) VisitInterpretedliteralAST(ctx *parser.InterpretedliteralASTContext) interface{} {
	//return ctx.INTERPRETEDSTRINGLITERAL().GetText()
	valReturn := constant.NewCharArrayFromString(ctx.INTERPRETEDSTRINGLITERAL().GetText())
	return valReturn
}
func (tc *Encoder) VisitIdexAST(ctx *parser.IdexASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *Encoder) VisitArgumentsAST(ctx *parser.ArgumentsASTContext) interface{} {
	if (ctx.ExpressionList()) != nil {
		return tc.Visit(ctx.ExpressionList())
	}

	return nil
}
func (tc *Encoder) VisitSelectorAST(ctx *parser.SelectorASTContext) interface{} {
	return ctx.ID().(antlr.TerminalNode)
}
func (tc *Encoder) VisitAppendExpressionAST(ctx *parser.AppendExpressionASTContext) interface{} {
	tc.Visit(ctx.Expression(0))
	tc.Visit(ctx.Expression(1))
	return nil
}
func (tc *Encoder) VisitLengthExpressionAST(ctx *parser.LengthExpressionASTContext) interface{} {
	/*tc.Visit(ctx.Expression())
	return nil*/

	val := tc.Visit(ctx.Expression()).(value.Value)
	blockActual, _ := blocksFunActual.Peek()

	ptrType := val.Type().(*types.PointerType).ElemType

	if types.IsArray(ptrType) {
		//Se obtiene el tamaño del array
		size := ptrType.(*types.ArrayType).Len
		//Se crea una constante con el tamaño del array
		return constant.NewInt(types.I32, int64(size))
	}

	strLen := blockActual.NewCall(tc.puts, val)
	return strLen
}

func (tc *Encoder) VisitCapExpressionAST(ctx *parser.CapExpressionASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *Encoder) VisitStatementListAST(ctx *parser.StatementListASTContext) interface{} {
	for i := 0; i < len(ctx.AllStatement()); i++ {
		tc.Visit(ctx.Statement(i))
	}

	return nil
}
func (tc *Encoder) VisitBlockAST(ctx *parser.BlockASTContext) interface{} {
	//se agrega el bloque actual creado a la lista de bloques.. sin nombre
	fmt.Printf("Entree a block")
	blocksFunActual.Push(funcActual.NewBlock(""))

	if funcActual.Sig.RetType == types.Void { //si la función no tiene un bloque de retorno
		//se agrega un bloque de retorno
		funcActual.Blocks[len(funcActual.Blocks)-1].NewRet(nil)
	}

	tc.Visit(ctx.StatementList())

	//que se devuelva el bloque creado para conocerlo a la vuelta del visit
	returnBlock, _ := blocksFunActual.Peek() //se obtiene el bloque actual para saber donde se va a declarar la variable
	return returnBlock

}
func (v *Encoder) VisitPrintSAST(ctx *parser.PrintSASTContext) interface{} {
	if (ctx.ExpressionList()) != nil {
		//tc.Visit(ctx.ExpressionList())

		val := v.Visit(ctx.ExpressionList()).([]value.Value)

		//Se obtiene el bloque actual de la pila de bloques
		blockActual, _ := blocksFunActual.Peek()

		for i := 0; i < len(val); i++ {

			//Se extra el tipo apuntado por el puntero i32* i32
			ptrType := val[i].Type().(*types.PointerType).ElemType
			//Se verifica el tipo de dato de la variable
			switch ptrType {
			case types.I32: //Int
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.I32, val[i])
				blockActual.NewCall(v.printf, v.aquiStr, valReturn)
			case types.Float: //Float
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.Float, val[i])
				blockActual.NewCall(v.printf, v.aquiStr, valReturn)
			case types.I8Ptr: //String
				println("Esta es una variable de tipo string")
				valReturn := blockActual.NewLoad(types.I8Ptr, val[i])
				blockActual.NewCall(v.printf, v.aquiStr, valReturn)
			case types.I1: //Bool
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.I1, val[i])
				blockActual.NewCall(v.printf, v.aquiStr, valReturn)
			case types.I8:
				println("Esta es una variable de tipo rune")
				println("Valor de la variable: ", val[i])
				//ptr:= blockActual.NewLoad(types.I8, val[i])
				blockActual.NewCall(v.puts, val[i])
			}
		}
	}
	return nil
}

func (v *Encoder) VisitPrintlnSAST(ctx *parser.PrintlnSASTContext) interface{} {
	println("\n/******* IMPRESIÓN DE VARIABLES con PRINTLN *******/")
	//Antes de imprimir se debe de cargar el valor de la variable en el bloque actual (load)

	//Se obtiene el valor de la expresión
	val := v.Visit(ctx.ExpressionList()).([]value.Value)
	fmt.Println("Valores de las expresiones PRINTLN: ", val)
	//Se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()

	for i := 0; i < len(val); i++ {

		var ptrType types.Type //Se guarda el tipo de dato del puntero

		if types.IsPointer(val[i].Type()) {
			ptrType = val[i].Type().(*types.PointerType).ElemType
		} else if types.IsArray(val[i].Type()) {
			ptrType = val[i].Type().(*types.ArrayType).ElemType
		} else {
			ptrType = val[i].Type() // Si no es puntero ni array, usar el tipo directamente
		}

		//Se verifica el tipo de dato de la variable
		switch ptrType {
		case types.I32: //Int
			println("Esta es una variable de tipo int", val[i].Ident())
			//Se carga el valor de la variable
			valReturn := blockActual.NewLoad(types.I32, val[i])
			blockActual.NewCall(v.printf, v.aquiStr, valReturn)
		case types.Float: //Float
			//Se carga el valor de la variable
			valReturn := blockActual.NewLoad(types.Float, val[i])
			blockActual.NewCall(v.printf, v.aquiStr, valReturn)
		case types.I8Ptr: //String
			println("Esta es una variable de tipo string")
			valReturn := blockActual.NewLoad(types.I8Ptr, val[i])
			blockActual.NewCall(v.printf, v.aquiStr, valReturn)
		case types.I1: //Bool
			//Se carga el valor de la variable
			valReturn := blockActual.NewLoad(types.I1, val[i])
			blockActual.NewCall(v.printf, v.aquiStr, valReturn)
		case types.I8:

			println("Esta es una variable de tipo rune")
			println("Valor de la variable: ", val[i].Ident())

			// Crear la cadena local "Hello, World!\0A"
			localStr := blockActual.NewAlloca(types.NewArray(uint64(len(val[i].Ident())+1), types.I8))
			valStr := val[i].Ident() + "\\n\\x00"
			for i, c := range []byte(valStr) {
				ptr := blockActual.NewGetElementPtr(types.NewArray(uint64(len(valStr)), types.I8), localStr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, int64(i)))
				blockActual.NewStore(constant.NewInt(types.I8, int64(c)), ptr)
			}

			// Obtener un puntero a la cadena local
			ptr := blockActual.NewGetElementPtr(types.NewArray(uint64(len(valStr)), types.I8), localStr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))

			blockActual.NewCall(v.puts, ptr)

			/*println("Esta es una variable de tipo rune")
			println("Valor de la variable: ", val[i])
			blockActual.NewCall(v.puts, val[i])*/
		}

	}
	return nil
}
func (tc *Encoder) VisitReturnSAST(ctx *parser.ReturnSASTContext) interface{} {

	//Se obtiene el bloque actual de la pila de bloques
	blockActual, _ := blocksFunActual.Peek()

	var tipoRetorno types.Type
	if funcActual.Sig.RetType != nil {
		tipoRetorno = funcActual.Sig.RetType
	} else {
		tipoRetorno = types.Void
	}

	//Se obtiene el valor de la expresión
	val := tc.Visit(ctx.Expression()).(string)

	//Se obtiene el tipo de dato de la variable
	if valLoad, found := variablesLocales.GetVariable(blockActual, val); found {

		//Se extrae el tipo apuntado por el puntero i32* i32
		ptrType := valLoad.Type().(*types.PointerType).ElemType

		if tipoRetorno == types.Void {
			blockActual.NewRet(nil)
		} else {
			//Se verifica el tipo de dato de la variable
			switch ptrType {
			case types.I32: //Int
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.I32, valLoad)
				blockActual.NewRet(valReturn)
			case types.Float: //Float
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.Float, valLoad)
				blockActual.NewRet(valReturn)
			case types.I8Ptr: //String
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.I8Ptr, valLoad)
				blockActual.NewRet(valReturn)
			case types.I1: //Bool
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.I1, valLoad)
				blockActual.NewRet(valReturn)
			case types.I8: //Rune
				//Se carga el valor de la variable
				valReturn := blockActual.NewLoad(types.I8, valLoad)
				blockActual.NewRet(valReturn)
			}
		}

	} else {
		println("Variable no encontrada")
	}
	return nil

}
func (tc *Encoder) VisitBreakSAST(ctx *parser.BreakSASTContext) interface{} {
	return nil
}
func (tc *Encoder) VisitContinueSAST(ctx *parser.ContinueSASTContext) interface{} {
	return nil
}
func (tc *Encoder) VisitSimpleStatementSAST(ctx *parser.SimpleStatementSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	return nil
}
func (tc *Encoder) VisitBlockSAST(ctx *parser.BlockSASTContext) interface{} {
	tc.Visit(ctx.Block())
	return nil
}
func (tc *Encoder) VisitSwitchSAST(ctx *parser.SwitchSASTContext) interface{} {
	tc.Visit(ctx.Switch_())
	return nil
}
func (tc *Encoder) VisitIfStatementSAST(ctx *parser.IfStatementSASTContext) interface{} {
	tc.Visit(ctx.IfStatement())
	return nil
}
func (tc *Encoder) VisitLoopSAST(ctx *parser.LoopSASTContext) interface{} {
	tc.Visit(ctx.Loop())
	return nil
}
func (tc *Encoder) VisitTypeDeclSAST(ctx *parser.TypeDeclSASTContext) interface{} {
	tc.Visit(ctx.TypeDecl())
	return nil
}
func (tc *Encoder) VisitVariableDeclSAST(ctx *parser.VariableDeclSASTContext) interface{} {
	tc.Visit(ctx.VariableDecl())
	return nil
}
func (tc *Encoder) VisitEpsilonSSAST(ctx *parser.EpsilonSSASTContext) interface{} {
	return nil
}
func (tc *Encoder) VisitExpressionINCSSAST(ctx *parser.ExpressionINCSSASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *Encoder) VisitExpressionDECSSAST(ctx *parser.ExpressionDECSSASTContext) interface{} {
	tc.Visit(ctx.Expression())
	return nil
}
func (tc *Encoder) VisitExpressionEpsilonSSAST(ctx *parser.ExpressionEpsilonSSASTContext) interface{} {

	tc.Visit(ctx.Expression())
	return nil
}
func (tc *Encoder) VisitAssigmentStatementSSAST(ctx *parser.AssigmentStatementSSASTContext) interface{} {
	tc.Visit(ctx.AssignmentStatement())
	return nil
}
func (tc *Encoder) VisitExpListSSAST(ctx *parser.ExpListSSASTContext) interface{} {

	tc.Visit(ctx.ExpressionList(0))
	tc.Visit(ctx.ExpressionList(1))
	return nil
}
func (tc *Encoder) VisitExpListASAST(ctx *parser.ExpListASASTContext) interface{} {
	exp1 := tc.Visit(ctx.ExpressionList(0)).([]value.Value)
	exp2 := tc.Visit(ctx.ExpressionList(1)).([]value.Value)

	blockActual, _ := blocksFunActual.Peek()
	for i := 0; i < len(exp1); i++ {
		blockActual.NewStore(exp2[i], exp1[i])
		variablesLocales.EditValue(blockActual, exp1[i].Ident(), exp2[i])

	}
	return nil

}
func (tc *Encoder) VisitIsExpressionBlockISAST(ctx *parser.IsExpressionBlockISASTContext) interface{} {
	//se respada el bloque que precede al IF
	blockAnteriorIf, _ := blocksFunActual.Peek()
	//se visita la expresión para generar el código de la comparación
	valCond := tc.Visit(ctx.Expression()).(value.Value)

	//se crea el cloque para el cuerpo del if...
	//luego se debe volver a poner las instrucciones del comparación y saldo al bloque Padre
	tc.Visit(ctx.Block())

	//se respalda el bloque del cuerpo del IF
	blockCuerpoIf, _ := blocksFunActual.Peek()

	//se crea el bloque para el cierre del if
	blocksFunActual.Push(funcActual.NewBlock(""))
	blockEndIf, _ := blocksFunActual.Peek()
	//una vez creado el bloque de cierre, se debe agregar la instrucción de salto al final del bloque del cuerpo del if para terminar dicho bloque
	blockCuerpoIf.NewBr(blockEndIf)

	//se debe agregar la instrucción de salto al final del bloque anterior al if para que se ejecute el bloque de cierre
	blockAnteriorIf.NewCondBr(valCond, blockCuerpoIf, blockEndIf)

	return nil
}
func (tc *Encoder) VisitIsExpBlockIsISAST(ctx *parser.IsExpBlockIsISASTContext) interface{} {
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	tc.Visit(ctx.IfStatement())
	return nil
}
func (tc *Encoder) VisitIsExpBlockISAST(ctx *parser.IsExpBlockISASTContext) interface{} {

	blockAnteriorIf, _ := blocksFunActual.Peek()

	valCond := tc.Visit(ctx.Expression()).(value.Value)

	tc.Visit(ctx.Block(0))
	blockIf, _ := blocksFunActual.Peek()

	tc.Visit(ctx.Block(1))
	blockElse, _ := blocksFunActual.Peek()

	blocksFunActual.Push(funcActual.NewBlock(""))
	blockOutIfElse, _ := blocksFunActual.Peek()

	blockIf.NewBr(blockOutIfElse)
	blockElse.NewBr(blockOutIfElse)

	blockAnteriorIf.NewCondBr(valCond, blockIf, blockElse)

	return nil
}

func (tc *Encoder) VisitIsSSExpBlockifSAST(ctx *parser.IsSSExpBlockifSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	return nil
}
func (tc *Encoder) VisitIsSSExpBlockIfSAST(ctx *parser.IsSSExpBlockifSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.Block())
	tc.Visit(ctx.IfStatement())
	return nil
}
func (tc *Encoder) VisitIsSSExpBlockBlockAST(ctx *parser.IsSSExpBlockBlockASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())

	tc.Visit(ctx.Block(0))
	tc.Visit(ctx.Block(1))
	return nil
}
func (tc *Encoder) VisitFBlockLAST(ctx *parser.FBlockLASTContext) interface{} {
	tc.Visit(ctx.Block())
	return nil
}
func (tc *Encoder) VisitFExpBlockLAST(ctx *parser.FExpBlockLASTContext) interface{} {

	blockActual, _ := blocksFunActual.Peek()

	blocksFunActual.Push(funcActual.NewBlock(""))
	blockCond, _ := blocksFunActual.Peek()
	blockActual.NewBr(blockCond)

	valCond := tc.Visit(ctx.Expression()).(value.Value)

	tc.Visit(ctx.Block())

	blockLoop, _ := blocksFunActual.Peek()
	blocksFunActual.Push(funcActual.NewBlock(""))
	blockAfterLoop, _ := blocksFunActual.Peek()

	blockLoop.NewBr(blockCond)
	blockCond.NewCondBr(valCond, blockLoop, blockAfterLoop)

	return nil

}
func (tc *Encoder) VisitFSimpStateExpSBlockLAST(ctx *parser.FSimpStateExpSBlockLASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement(0))
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.SimpleStatement(1))
	tc.Visit(ctx.Block())

	return nil
}
func (tc *Encoder) VisitFSimpStateSimpStateBlockLAST(ctx *parser.FSimpStateSimpStateBlockLASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement(0))
	tc.Visit(ctx.SimpleStatement(1))
	tc.Visit(ctx.Block())

	return nil
}

func (tc *Encoder) VisitSSimpleStatSAST(ctx *parser.SSimpleStatSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *Encoder) VisitSExpressionSAST(ctx *parser.SExpressionSASTContext) interface{} {
	tc.Visit(ctx.Expression())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *Encoder) VisitSSimpleStatExpCCListSAST(ctx *parser.SSimpleStatExpCCListSASTContext) interface{} {
	tc.Visit(ctx.SimpleStatement())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *Encoder) VisitSBlockSAST(ctx *parser.SBlockSASTContext) interface{} {
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *Encoder) VisitEpsilonECCLAST(ctx *parser.EpsilonECCLASTContext) interface{} {
	//tc.Visit(ctx.Epsilon())
	return nil
}
func (tc *Encoder) VisitExpCaseClauseECCLAST(ctx *parser.ExpCaseClauseECCLASTContext) interface{} {
	tc.Visit(ctx.ExpressionCaseClause())
	tc.Visit(ctx.ExpressionCaseClauseList())
	return nil
}
func (tc *Encoder) VisitExpSwitchCaseECCAST(ctx *parser.ExpSwitchCaseECCASTContext) interface{} {
	tc.Visit(ctx.ExpressionSwitchCase())
	tc.Visit(ctx.StatementList())
	return nil
}
func (tc *Encoder) VisitCaseESCAST(ctx *parser.CaseESCASTContext) interface{} {
	tc.Visit(ctx.ExpressionList())
	return nil
}
func (tc *Encoder) VisitDefaultESCAST(ctx *parser.DefaultESCASTContext) interface{} {
	return ctx.DEFAULT()
}

func (tc *Encoder) VisitEpsilonAST(ctx *parser.EpsilonASTContext) interface{} {
	return nil
}
