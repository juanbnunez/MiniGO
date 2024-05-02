// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // goParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type goParser struct {
	*antlr.BaseParser
}

var GoParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goparserParserInit() {
	staticData := &GoParserParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "':='", "'('", "')'", "'~'", "':'", "'+'", "'-'", "'*'",
		"'/'", "','", "'['", "']'", "'{'", "'}'", "'++'", "'--'", "'%'", "'<<'",
		"'>>'", "'&'", "'&^'", "'|'", "'^'", "'=='", "'!='", "'<'", "'>'", "'<='",
		"'>='", "'&&'", "'||'", "'!'", "'='", "'+='", "'&='", "'-='", "'|='",
		"'*='", "'^='", "'<<='", "'>>='", "'&^='", "'%='", "'/='", "'.'", "'\\'",
		"'''", "'package'", "'if'", "'while'", "'let'", "'then'", "'else'",
		"'in'", "'do'", "'begin'", "'end'", "'const'", "'var'", "'type'", "'func'",
		"'struct'", "'append'", "'len'", "'cap'", "'print'", "'println'", "'return'",
		"'break'", "'continue'", "'for'", "'switch'", "'case'", "'default'",
	}
	staticData.SymbolicNames = []string{
		"", "SEMI", "SVD", "LP", "RP", "TIL", "COL", "PL", "MIN", "MUL", "DIV",
		"CM", "LSB", "RSB", "LCB", "RCB", "INC", "DEC", "MOD", "LSH", "RSH",
		"AMPER", "BC", "VB", "CARET", "EQ", "NEQ", "LT", "GT", "LTE", "GTE",
		"LAND", "LOR", "NEG", "ASOP", "AAOP", "BAOP", "SAOP", "BOAOP", "MAOP",
		"BXOOP", "LSAOP", "RSAOP", "BCAOP", "RAOP", "DAOP", "DOT", "DBS", "BS",
		"PACKAGE", "IF", "WHILE", "LET", "THEN", "ELSE", "IN", "DO", "BEGIN",
		"END", "CONST", "VAR", "TYPE", "FUNC", "STRUCT", "APPEND", "LEN", "CAP",
		"PRINT", "PRINTLN", "RETURN", "BREAK", "CONTINUE", "FOR", "SWITCH",
		"CASE", "DEFAULT", "ID", "INTLITERAL", "FLOATLITERAL", "RUNELITERAL",
		"RAWSTRINGLITERAL", "INTERPRETEDSTRINGLITERAL", "LETTER", "DIGIT", "LINE_COMMENT",
		"SPACES",
	}
	staticData.RuleNames = []string{
		"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl",
		"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl",
		"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType",
		"arrayDeclType", "structDeclType", "structMemDecls", "identifierList",
		"expression", "expressionList", "primaryExpression", "operand", "literal",
		"index", "arguments", "selector", "appendExpression", "lengthExpression",
		"capExpression", "statementList", "block", "statement", "simpleStatement",
		"assignmentStatement", "ifStatement", "loop", "switch", "expressionCaseClauseList",
		"expressionCaseClause", "expressionSwitchCase", "epsilon",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 85, 514, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 91, 8, 1, 10, 1, 12, 1, 94, 9, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 117, 8,
		3, 10, 3, 12, 3, 120, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 151,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 158, 8, 7, 10, 7, 12, 7, 161,
		9, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 175, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 180, 8, 10, 1,
		11, 1, 11, 1, 11, 5, 11, 185, 8, 11, 10, 11, 12, 11, 188, 9, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 198, 8, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 3, 15, 213, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 5, 16, 222, 8, 16, 10, 16, 12, 16, 225, 9, 16, 1, 17, 1, 17,
		1, 17, 5, 17, 230, 8, 17, 10, 17, 12, 17, 233, 9, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 239, 8, 18, 1, 18, 1, 18, 1, 18, 5, 18, 244, 8, 18, 10,
		18, 12, 18, 247, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 252, 8, 19, 10, 19,
		12, 19, 255, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 262, 8, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 270, 8, 20, 10, 20, 12,
		20, 273, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 281, 8,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 288, 8, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 3, 24, 297, 8, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 5,
		29, 322, 8, 29, 10, 29, 12, 29, 325, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 335, 8, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 3, 31, 344, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 3, 31, 352, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 377, 8, 31, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 394, 8, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 404, 8, 33, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 3, 34, 444, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 467, 8, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 495, 8, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 3, 37, 501, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		3, 39, 510, 8, 39, 1, 40, 1, 40, 1, 40, 0, 2, 36, 40, 41, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
		0, 3, 3, 0, 7, 8, 24, 24, 33, 33, 2, 0, 7, 10, 18, 32, 1, 0, 35, 45, 543,
		0, 82, 1, 0, 0, 0, 2, 92, 1, 0, 0, 0, 4, 109, 1, 0, 0, 0, 6, 111, 1, 0,
		0, 0, 8, 131, 1, 0, 0, 0, 10, 133, 1, 0, 0, 0, 12, 150, 1, 0, 0, 0, 14,
		152, 1, 0, 0, 0, 16, 162, 1, 0, 0, 0, 18, 165, 1, 0, 0, 0, 20, 169, 1,
		0, 0, 0, 22, 181, 1, 0, 0, 0, 24, 197, 1, 0, 0, 0, 26, 199, 1, 0, 0, 0,
		28, 203, 1, 0, 0, 0, 30, 208, 1, 0, 0, 0, 32, 216, 1, 0, 0, 0, 34, 226,
		1, 0, 0, 0, 36, 238, 1, 0, 0, 0, 38, 248, 1, 0, 0, 0, 40, 261, 1, 0, 0,
		0, 42, 280, 1, 0, 0, 0, 44, 287, 1, 0, 0, 0, 46, 289, 1, 0, 0, 0, 48, 293,
		1, 0, 0, 0, 50, 300, 1, 0, 0, 0, 52, 303, 1, 0, 0, 0, 54, 310, 1, 0, 0,
		0, 56, 315, 1, 0, 0, 0, 58, 323, 1, 0, 0, 0, 60, 326, 1, 0, 0, 0, 62, 376,
		1, 0, 0, 0, 64, 393, 1, 0, 0, 0, 66, 403, 1, 0, 0, 0, 68, 443, 1, 0, 0,
		0, 70, 466, 1, 0, 0, 0, 72, 494, 1, 0, 0, 0, 74, 500, 1, 0, 0, 0, 76, 502,
		1, 0, 0, 0, 78, 509, 1, 0, 0, 0, 80, 511, 1, 0, 0, 0, 82, 83, 5, 49, 0,
		0, 83, 84, 5, 76, 0, 0, 84, 85, 5, 1, 0, 0, 85, 86, 3, 2, 1, 0, 86, 1,
		1, 0, 0, 0, 87, 91, 3, 4, 2, 0, 88, 91, 3, 12, 6, 0, 89, 91, 3, 18, 9,
		0, 90, 87, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94,
		1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 3, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 95, 96, 5, 60, 0, 0, 96, 97, 3, 8, 4, 0, 97, 98, 5,
		1, 0, 0, 98, 110, 1, 0, 0, 0, 99, 100, 5, 60, 0, 0, 100, 101, 5, 3, 0,
		0, 101, 102, 3, 6, 3, 0, 102, 103, 5, 4, 0, 0, 103, 104, 5, 1, 0, 0, 104,
		110, 1, 0, 0, 0, 105, 106, 5, 60, 0, 0, 106, 107, 5, 3, 0, 0, 107, 108,
		5, 4, 0, 0, 108, 110, 5, 1, 0, 0, 109, 95, 1, 0, 0, 0, 109, 99, 1, 0, 0,
		0, 109, 105, 1, 0, 0, 0, 110, 5, 1, 0, 0, 0, 111, 112, 3, 8, 4, 0, 112,
		118, 5, 1, 0, 0, 113, 114, 3, 8, 4, 0, 114, 115, 5, 1, 0, 0, 115, 117,
		1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0,
		0, 0, 118, 119, 1, 0, 0, 0, 119, 7, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121,
		122, 3, 34, 17, 0, 122, 123, 3, 24, 12, 0, 123, 124, 5, 34, 0, 0, 124,
		125, 3, 38, 19, 0, 125, 132, 1, 0, 0, 0, 126, 127, 3, 34, 17, 0, 127, 128,
		5, 34, 0, 0, 128, 129, 3, 38, 19, 0, 129, 132, 1, 0, 0, 0, 130, 132, 3,
		10, 5, 0, 131, 121, 1, 0, 0, 0, 131, 126, 1, 0, 0, 0, 131, 130, 1, 0, 0,
		0, 132, 9, 1, 0, 0, 0, 133, 134, 3, 34, 17, 0, 134, 135, 3, 24, 12, 0,
		135, 11, 1, 0, 0, 0, 136, 137, 5, 61, 0, 0, 137, 138, 3, 16, 8, 0, 138,
		139, 5, 1, 0, 0, 139, 151, 1, 0, 0, 0, 140, 141, 5, 61, 0, 0, 141, 142,
		5, 3, 0, 0, 142, 143, 3, 14, 7, 0, 143, 144, 5, 4, 0, 0, 144, 145, 5, 1,
		0, 0, 145, 151, 1, 0, 0, 0, 146, 147, 5, 61, 0, 0, 147, 148, 5, 3, 0, 0,
		148, 149, 5, 4, 0, 0, 149, 151, 5, 1, 0, 0, 150, 136, 1, 0, 0, 0, 150,
		140, 1, 0, 0, 0, 150, 146, 1, 0, 0, 0, 151, 13, 1, 0, 0, 0, 152, 153, 3,
		16, 8, 0, 153, 159, 5, 1, 0, 0, 154, 155, 3, 16, 8, 0, 155, 156, 5, 1,
		0, 0, 156, 158, 1, 0, 0, 0, 157, 154, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0,
		159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 15, 1, 0, 0, 0, 161, 159,
		1, 0, 0, 0, 162, 163, 5, 76, 0, 0, 163, 164, 3, 24, 12, 0, 164, 17, 1,
		0, 0, 0, 165, 166, 3, 20, 10, 0, 166, 167, 3, 60, 30, 0, 167, 168, 5, 1,
		0, 0, 168, 19, 1, 0, 0, 0, 169, 170, 5, 62, 0, 0, 170, 171, 5, 76, 0, 0,
		171, 174, 5, 3, 0, 0, 172, 175, 3, 22, 11, 0, 173, 175, 3, 80, 40, 0, 174,
		172, 1, 0, 0, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 179,
		5, 4, 0, 0, 177, 180, 3, 24, 12, 0, 178, 180, 3, 80, 40, 0, 179, 177, 1,
		0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 21, 1, 0, 0, 0, 181, 186, 3, 10, 5,
		0, 182, 183, 5, 11, 0, 0, 183, 185, 3, 10, 5, 0, 184, 182, 1, 0, 0, 0,
		185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		23, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 190, 5, 3, 0, 0, 190, 191, 3,
		24, 12, 0, 191, 192, 5, 4, 0, 0, 192, 198, 1, 0, 0, 0, 193, 198, 5, 76,
		0, 0, 194, 198, 3, 26, 13, 0, 195, 198, 3, 28, 14, 0, 196, 198, 3, 30,
		15, 0, 197, 189, 1, 0, 0, 0, 197, 193, 1, 0, 0, 0, 197, 194, 1, 0, 0, 0,
		197, 195, 1, 0, 0, 0, 197, 196, 1, 0, 0, 0, 198, 25, 1, 0, 0, 0, 199, 200,
		5, 12, 0, 0, 200, 201, 5, 13, 0, 0, 201, 202, 3, 24, 12, 0, 202, 27, 1,
		0, 0, 0, 203, 204, 5, 12, 0, 0, 204, 205, 5, 77, 0, 0, 205, 206, 5, 13,
		0, 0, 206, 207, 3, 24, 12, 0, 207, 29, 1, 0, 0, 0, 208, 209, 5, 63, 0,
		0, 209, 212, 5, 14, 0, 0, 210, 213, 3, 32, 16, 0, 211, 213, 3, 80, 40,
		0, 212, 210, 1, 0, 0, 0, 212, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214,
		215, 5, 15, 0, 0, 215, 31, 1, 0, 0, 0, 216, 217, 3, 10, 5, 0, 217, 223,
		5, 1, 0, 0, 218, 219, 3, 10, 5, 0, 219, 220, 5, 1, 0, 0, 220, 222, 1, 0,
		0, 0, 221, 218, 1, 0, 0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0,
		223, 224, 1, 0, 0, 0, 224, 33, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 231,
		5, 76, 0, 0, 227, 228, 5, 11, 0, 0, 228, 230, 5, 76, 0, 0, 229, 227, 1,
		0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0,
		0, 232, 35, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 235, 6, 18, -1, 0, 235,
		239, 3, 40, 20, 0, 236, 237, 7, 0, 0, 0, 237, 239, 3, 36, 18, 1, 238, 234,
		1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 245, 1, 0, 0, 0, 240, 241, 10, 2,
		0, 0, 241, 242, 7, 1, 0, 0, 242, 244, 3, 36, 18, 3, 243, 240, 1, 0, 0,
		0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246,
		37, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 253, 3, 36, 18, 0, 249, 250,
		5, 11, 0, 0, 250, 252, 3, 36, 18, 0, 251, 249, 1, 0, 0, 0, 252, 255, 1,
		0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 39, 1, 0, 0,
		0, 255, 253, 1, 0, 0, 0, 256, 257, 6, 20, -1, 0, 257, 262, 3, 42, 21, 0,
		258, 262, 3, 52, 26, 0, 259, 262, 3, 54, 27, 0, 260, 262, 3, 56, 28, 0,
		261, 256, 1, 0, 0, 0, 261, 258, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261,
		260, 1, 0, 0, 0, 262, 271, 1, 0, 0, 0, 263, 264, 10, 6, 0, 0, 264, 270,
		3, 50, 25, 0, 265, 266, 10, 5, 0, 0, 266, 270, 3, 46, 23, 0, 267, 268,
		10, 4, 0, 0, 268, 270, 3, 48, 24, 0, 269, 263, 1, 0, 0, 0, 269, 265, 1,
		0, 0, 0, 269, 267, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0,
		0, 271, 272, 1, 0, 0, 0, 272, 41, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274,
		281, 3, 44, 22, 0, 275, 281, 5, 76, 0, 0, 276, 277, 5, 3, 0, 0, 277, 278,
		3, 36, 18, 0, 278, 279, 5, 4, 0, 0, 279, 281, 1, 0, 0, 0, 280, 274, 1,
		0, 0, 0, 280, 275, 1, 0, 0, 0, 280, 276, 1, 0, 0, 0, 281, 43, 1, 0, 0,
		0, 282, 288, 5, 77, 0, 0, 283, 288, 5, 78, 0, 0, 284, 288, 5, 79, 0, 0,
		285, 288, 5, 80, 0, 0, 286, 288, 5, 81, 0, 0, 287, 282, 1, 0, 0, 0, 287,
		283, 1, 0, 0, 0, 287, 284, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 286,
		1, 0, 0, 0, 288, 45, 1, 0, 0, 0, 289, 290, 5, 12, 0, 0, 290, 291, 3, 36,
		18, 0, 291, 292, 5, 13, 0, 0, 292, 47, 1, 0, 0, 0, 293, 296, 5, 3, 0, 0,
		294, 297, 3, 38, 19, 0, 295, 297, 3, 80, 40, 0, 296, 294, 1, 0, 0, 0, 296,
		295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299, 5, 4, 0, 0, 299, 49, 1,
		0, 0, 0, 300, 301, 5, 46, 0, 0, 301, 302, 5, 76, 0, 0, 302, 51, 1, 0, 0,
		0, 303, 304, 5, 64, 0, 0, 304, 305, 5, 3, 0, 0, 305, 306, 3, 36, 18, 0,
		306, 307, 5, 11, 0, 0, 307, 308, 3, 36, 18, 0, 308, 309, 5, 4, 0, 0, 309,
		53, 1, 0, 0, 0, 310, 311, 5, 65, 0, 0, 311, 312, 5, 3, 0, 0, 312, 313,
		3, 36, 18, 0, 313, 314, 5, 4, 0, 0, 314, 55, 1, 0, 0, 0, 315, 316, 5, 66,
		0, 0, 316, 317, 5, 3, 0, 0, 317, 318, 3, 36, 18, 0, 318, 319, 5, 4, 0,
		0, 319, 57, 1, 0, 0, 0, 320, 322, 3, 62, 31, 0, 321, 320, 1, 0, 0, 0, 322,
		325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 59, 1,
		0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 327, 5, 14, 0, 0, 327, 328, 3, 58,
		29, 0, 328, 329, 5, 15, 0, 0, 329, 61, 1, 0, 0, 0, 330, 331, 5, 67, 0,
		0, 331, 334, 5, 3, 0, 0, 332, 335, 3, 38, 19, 0, 333, 335, 3, 80, 40, 0,
		334, 332, 1, 0, 0, 0, 334, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336,
		337, 5, 4, 0, 0, 337, 338, 5, 1, 0, 0, 338, 377, 1, 0, 0, 0, 339, 340,
		5, 68, 0, 0, 340, 343, 5, 3, 0, 0, 341, 344, 3, 38, 19, 0, 342, 344, 3,
		80, 40, 0, 343, 341, 1, 0, 0, 0, 343, 342, 1, 0, 0, 0, 344, 345, 1, 0,
		0, 0, 345, 346, 5, 4, 0, 0, 346, 347, 5, 1, 0, 0, 347, 377, 1, 0, 0, 0,
		348, 351, 5, 69, 0, 0, 349, 352, 3, 36, 18, 0, 350, 352, 3, 80, 40, 0,
		351, 349, 1, 0, 0, 0, 351, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353,
		354, 5, 1, 0, 0, 354, 377, 1, 0, 0, 0, 355, 356, 5, 70, 0, 0, 356, 377,
		5, 1, 0, 0, 357, 358, 5, 71, 0, 0, 358, 377, 5, 1, 0, 0, 359, 360, 3, 64,
		32, 0, 360, 361, 5, 1, 0, 0, 361, 377, 1, 0, 0, 0, 362, 363, 3, 60, 30,
		0, 363, 364, 5, 1, 0, 0, 364, 377, 1, 0, 0, 0, 365, 366, 3, 72, 36, 0,
		366, 367, 5, 1, 0, 0, 367, 377, 1, 0, 0, 0, 368, 369, 3, 68, 34, 0, 369,
		370, 5, 1, 0, 0, 370, 377, 1, 0, 0, 0, 371, 372, 3, 70, 35, 0, 372, 373,
		5, 1, 0, 0, 373, 377, 1, 0, 0, 0, 374, 377, 3, 12, 6, 0, 375, 377, 3, 4,
		2, 0, 376, 330, 1, 0, 0, 0, 376, 339, 1, 0, 0, 0, 376, 348, 1, 0, 0, 0,
		376, 355, 1, 0, 0, 0, 376, 357, 1, 0, 0, 0, 376, 359, 1, 0, 0, 0, 376,
		362, 1, 0, 0, 0, 376, 365, 1, 0, 0, 0, 376, 368, 1, 0, 0, 0, 376, 371,
		1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1, 0, 0, 0, 377, 63, 1, 0,
		0, 0, 378, 394, 3, 80, 40, 0, 379, 380, 3, 36, 18, 0, 380, 381, 5, 16,
		0, 0, 381, 394, 1, 0, 0, 0, 382, 383, 3, 36, 18, 0, 383, 384, 5, 17, 0,
		0, 384, 394, 1, 0, 0, 0, 385, 386, 3, 36, 18, 0, 386, 387, 3, 80, 40, 0,
		387, 394, 1, 0, 0, 0, 388, 394, 3, 66, 33, 0, 389, 390, 3, 38, 19, 0, 390,
		391, 5, 2, 0, 0, 391, 392, 3, 38, 19, 0, 392, 394, 1, 0, 0, 0, 393, 378,
		1, 0, 0, 0, 393, 379, 1, 0, 0, 0, 393, 382, 1, 0, 0, 0, 393, 385, 1, 0,
		0, 0, 393, 388, 1, 0, 0, 0, 393, 389, 1, 0, 0, 0, 394, 65, 1, 0, 0, 0,
		395, 396, 3, 38, 19, 0, 396, 397, 5, 34, 0, 0, 397, 398, 3, 38, 19, 0,
		398, 404, 1, 0, 0, 0, 399, 400, 3, 36, 18, 0, 400, 401, 7, 2, 0, 0, 401,
		402, 3, 36, 18, 0, 402, 404, 1, 0, 0, 0, 403, 395, 1, 0, 0, 0, 403, 399,
		1, 0, 0, 0, 404, 67, 1, 0, 0, 0, 405, 406, 5, 50, 0, 0, 406, 407, 3, 36,
		18, 0, 407, 408, 3, 60, 30, 0, 408, 444, 1, 0, 0, 0, 409, 410, 5, 50, 0,
		0, 410, 411, 3, 36, 18, 0, 411, 412, 3, 60, 30, 0, 412, 413, 5, 54, 0,
		0, 413, 414, 3, 68, 34, 0, 414, 444, 1, 0, 0, 0, 415, 416, 5, 50, 0, 0,
		416, 417, 3, 36, 18, 0, 417, 418, 3, 60, 30, 0, 418, 419, 5, 54, 0, 0,
		419, 420, 3, 60, 30, 0, 420, 444, 1, 0, 0, 0, 421, 422, 5, 50, 0, 0, 422,
		423, 3, 64, 32, 0, 423, 424, 5, 1, 0, 0, 424, 425, 3, 36, 18, 0, 425, 426,
		3, 60, 30, 0, 426, 444, 1, 0, 0, 0, 427, 428, 5, 50, 0, 0, 428, 429, 3,
		64, 32, 0, 429, 430, 5, 1, 0, 0, 430, 431, 3, 36, 18, 0, 431, 432, 3, 60,
		30, 0, 432, 433, 5, 54, 0, 0, 433, 434, 3, 68, 34, 0, 434, 444, 1, 0, 0,
		0, 435, 436, 5, 50, 0, 0, 436, 437, 3, 64, 32, 0, 437, 438, 5, 1, 0, 0,
		438, 439, 3, 36, 18, 0, 439, 440, 3, 60, 30, 0, 440, 441, 5, 54, 0, 0,
		441, 442, 3, 60, 30, 0, 442, 444, 1, 0, 0, 0, 443, 405, 1, 0, 0, 0, 443,
		409, 1, 0, 0, 0, 443, 415, 1, 0, 0, 0, 443, 421, 1, 0, 0, 0, 443, 427,
		1, 0, 0, 0, 443, 435, 1, 0, 0, 0, 444, 69, 1, 0, 0, 0, 445, 446, 5, 72,
		0, 0, 446, 467, 3, 60, 30, 0, 447, 448, 5, 72, 0, 0, 448, 449, 3, 36, 18,
		0, 449, 450, 3, 60, 30, 0, 450, 467, 1, 0, 0, 0, 451, 452, 5, 72, 0, 0,
		452, 453, 3, 64, 32, 0, 453, 454, 5, 1, 0, 0, 454, 455, 3, 36, 18, 0, 455,
		456, 5, 1, 0, 0, 456, 457, 3, 64, 32, 0, 457, 458, 3, 60, 30, 0, 458, 467,
		1, 0, 0, 0, 459, 460, 5, 72, 0, 0, 460, 461, 3, 64, 32, 0, 461, 462, 5,
		1, 0, 0, 462, 463, 5, 1, 0, 0, 463, 464, 3, 64, 32, 0, 464, 465, 3, 60,
		30, 0, 465, 467, 1, 0, 0, 0, 466, 445, 1, 0, 0, 0, 466, 447, 1, 0, 0, 0,
		466, 451, 1, 0, 0, 0, 466, 459, 1, 0, 0, 0, 467, 71, 1, 0, 0, 0, 468, 469,
		5, 73, 0, 0, 469, 470, 3, 64, 32, 0, 470, 471, 5, 1, 0, 0, 471, 472, 3,
		36, 18, 0, 472, 473, 5, 14, 0, 0, 473, 474, 3, 74, 37, 0, 474, 475, 5,
		15, 0, 0, 475, 495, 1, 0, 0, 0, 476, 477, 5, 73, 0, 0, 477, 478, 3, 36,
		18, 0, 478, 479, 5, 14, 0, 0, 479, 480, 3, 74, 37, 0, 480, 481, 5, 15,
		0, 0, 481, 495, 1, 0, 0, 0, 482, 483, 5, 73, 0, 0, 483, 484, 3, 64, 32,
		0, 484, 485, 5, 1, 0, 0, 485, 486, 5, 14, 0, 0, 486, 487, 3, 74, 37, 0,
		487, 488, 5, 15, 0, 0, 488, 495, 1, 0, 0, 0, 489, 490, 5, 73, 0, 0, 490,
		491, 5, 14, 0, 0, 491, 492, 3, 74, 37, 0, 492, 493, 5, 15, 0, 0, 493, 495,
		1, 0, 0, 0, 494, 468, 1, 0, 0, 0, 494, 476, 1, 0, 0, 0, 494, 482, 1, 0,
		0, 0, 494, 489, 1, 0, 0, 0, 495, 73, 1, 0, 0, 0, 496, 501, 3, 80, 40, 0,
		497, 498, 3, 76, 38, 0, 498, 499, 3, 74, 37, 0, 499, 501, 1, 0, 0, 0, 500,
		496, 1, 0, 0, 0, 500, 497, 1, 0, 0, 0, 501, 75, 1, 0, 0, 0, 502, 503, 3,
		78, 39, 0, 503, 504, 5, 6, 0, 0, 504, 505, 3, 58, 29, 0, 505, 77, 1, 0,
		0, 0, 506, 507, 5, 74, 0, 0, 507, 510, 3, 38, 19, 0, 508, 510, 5, 75, 0,
		0, 509, 506, 1, 0, 0, 0, 509, 508, 1, 0, 0, 0, 510, 79, 1, 0, 0, 0, 511,
		512, 1, 0, 0, 0, 512, 81, 1, 0, 0, 0, 35, 90, 92, 109, 118, 131, 150, 159,
		174, 179, 186, 197, 212, 223, 231, 238, 245, 253, 261, 269, 271, 280, 287,
		296, 323, 334, 343, 351, 376, 393, 403, 443, 466, 494, 500, 509,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// goParserInit initializes any static state used to implement goParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoParserInit() {
	staticData := &GoParserParserStaticData
	staticData.once.Do(goparserParserInit)
}

// NewgoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgoParser(input antlr.TokenStream) *goParser {
	GoParserInit()
	this := new(goParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "goParser.g4"

	return this
}

// goParser tokens.
const (
	goParserEOF                      = antlr.TokenEOF
	goParserSEMI                     = 1
	goParserSVD                      = 2
	goParserLP                       = 3
	goParserRP                       = 4
	goParserTIL                      = 5
	goParserCOL                      = 6
	goParserPL                       = 7
	goParserMIN                      = 8
	goParserMUL                      = 9
	goParserDIV                      = 10
	goParserCM                       = 11
	goParserLSB                      = 12
	goParserRSB                      = 13
	goParserLCB                      = 14
	goParserRCB                      = 15
	goParserINC                      = 16
	goParserDEC                      = 17
	goParserMOD                      = 18
	goParserLSH                      = 19
	goParserRSH                      = 20
	goParserAMPER                    = 21
	goParserBC                       = 22
	goParserVB                       = 23
	goParserCARET                    = 24
	goParserEQ                       = 25
	goParserNEQ                      = 26
	goParserLT                       = 27
	goParserGT                       = 28
	goParserLTE                      = 29
	goParserGTE                      = 30
	goParserLAND                     = 31
	goParserLOR                      = 32
	goParserNEG                      = 33
	goParserASOP                     = 34
	goParserAAOP                     = 35
	goParserBAOP                     = 36
	goParserSAOP                     = 37
	goParserBOAOP                    = 38
	goParserMAOP                     = 39
	goParserBXOOP                    = 40
	goParserLSAOP                    = 41
	goParserRSAOP                    = 42
	goParserBCAOP                    = 43
	goParserRAOP                     = 44
	goParserDAOP                     = 45
	goParserDOT                      = 46
	goParserDBS                      = 47
	goParserBS                       = 48
	goParserPACKAGE                  = 49
	goParserIF                       = 50
	goParserWHILE                    = 51
	goParserLET                      = 52
	goParserTHEN                     = 53
	goParserELSE                     = 54
	goParserIN                       = 55
	goParserDO                       = 56
	goParserBEGIN                    = 57
	goParserEND                      = 58
	goParserCONST                    = 59
	goParserVAR                      = 60
	goParserTYPE                     = 61
	goParserFUNC                     = 62
	goParserSTRUCT                   = 63
	goParserAPPEND                   = 64
	goParserLEN                      = 65
	goParserCAP                      = 66
	goParserPRINT                    = 67
	goParserPRINTLN                  = 68
	goParserRETURN                   = 69
	goParserBREAK                    = 70
	goParserCONTINUE                 = 71
	goParserFOR                      = 72
	goParserSWITCH                   = 73
	goParserCASE                     = 74
	goParserDEFAULT                  = 75
	goParserID                       = 76
	goParserINTLITERAL               = 77
	goParserFLOATLITERAL             = 78
	goParserRUNELITERAL              = 79
	goParserRAWSTRINGLITERAL         = 80
	goParserINTERPRETEDSTRINGLITERAL = 81
	goParserLETTER                   = 82
	goParserDIGIT                    = 83
	goParserLINE_COMMENT             = 84
	goParserSPACES                   = 85
)

// goParser rules.
const (
	goParserRULE_root                     = 0
	goParserRULE_topDeclarationList       = 1
	goParserRULE_variableDecl             = 2
	goParserRULE_innerVarDecls            = 3
	goParserRULE_singleVarDecl            = 4
	goParserRULE_singleVarDeclNoExps      = 5
	goParserRULE_typeDecl                 = 6
	goParserRULE_innerTypeDecls           = 7
	goParserRULE_singleTypeDecl           = 8
	goParserRULE_funcDecl                 = 9
	goParserRULE_funcFrontDecl            = 10
	goParserRULE_funcArgDecls             = 11
	goParserRULE_declType                 = 12
	goParserRULE_sliceDeclType            = 13
	goParserRULE_arrayDeclType            = 14
	goParserRULE_structDeclType           = 15
	goParserRULE_structMemDecls           = 16
	goParserRULE_identifierList           = 17
	goParserRULE_expression               = 18
	goParserRULE_expressionList           = 19
	goParserRULE_primaryExpression        = 20
	goParserRULE_operand                  = 21
	goParserRULE_literal                  = 22
	goParserRULE_index                    = 23
	goParserRULE_arguments                = 24
	goParserRULE_selector                 = 25
	goParserRULE_appendExpression         = 26
	goParserRULE_lengthExpression         = 27
	goParserRULE_capExpression            = 28
	goParserRULE_statementList            = 29
	goParserRULE_block                    = 30
	goParserRULE_statement                = 31
	goParserRULE_simpleStatement          = 32
	goParserRULE_assignmentStatement      = 33
	goParserRULE_ifStatement              = 34
	goParserRULE_loop                     = 35
	goParserRULE_switch                   = 36
	goParserRULE_expressionCaseClauseList = 37
	goParserRULE_expressionCaseClause     = 38
	goParserRULE_expressionSwitchCase     = 39
	goParserRULE_epsilon                  = 40
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) CopyAll(ctx *RootContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RootASTContext struct {
	RootContext
}

func NewRootASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RootASTContext {
	var p = new(RootASTContext)

	InitEmptyRootContext(&p.RootContext)
	p.parser = parser
	p.CopyAll(ctx.(*RootContext))

	return p
}

func (s *RootASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootASTContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(goParserPACKAGE, 0)
}

func (s *RootASTContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *RootASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *RootASTContext) TopDeclarationList() ITopDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDeclarationListContext)
}

func (s *RootASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterRootAST(s)
	}
}

func (s *RootASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitRootAST(s)
	}
}

func (s *RootASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitRootAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, goParserRULE_root)
	localctx = NewRootASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(goParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.TopDeclarationList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopDeclarationListContext is an interface to support dynamic dispatch.
type ITopDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTopDeclarationListContext differentiates from other interfaces.
	IsTopDeclarationListContext()
}

type TopDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDeclarationListContext() *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_topDeclarationList
	return p
}

func InitEmptyTopDeclarationListContext(p *TopDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_topDeclarationList
}

func (*TopDeclarationListContext) IsTopDeclarationListContext() {}

func NewTopDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_topDeclarationList

	return p
}

func (s *TopDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclarationListContext) CopyAll(ctx *TopDeclarationListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TopDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TopDeclarationListASTContext struct {
	TopDeclarationListContext
}

func NewTopDeclarationListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TopDeclarationListASTContext {
	var p = new(TopDeclarationListASTContext)

	InitEmptyTopDeclarationListContext(&p.TopDeclarationListContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopDeclarationListContext))

	return p
}

func (s *TopDeclarationListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListASTContext) AllVariableDecl() []IVariableDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclContext); ok {
			tst[i] = t.(IVariableDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) VariableDecl(i int) IVariableDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *TopDeclarationListASTContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TopDeclarationListASTContext) AllFuncDecl() []IFuncDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclContext); ok {
			tst[i] = t.(IFuncDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListASTContext) FuncDecl(i int) IFuncDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *TopDeclarationListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterTopDeclarationListAST(s)
	}
}

func (s *TopDeclarationListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitTopDeclarationListAST(s)
	}
}

func (s *TopDeclarationListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitTopDeclarationListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) TopDeclarationList() (localctx ITopDeclarationListContext) {
	localctx = NewTopDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, goParserRULE_topDeclarationList)
	var _la int

	localctx = NewTopDeclarationListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8070450532247928832) != 0 {
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case goParserVAR:
			{
				p.SetState(87)
				p.VariableDecl()
			}

		case goParserTYPE:
			{
				p.SetState(88)
				p.TypeDecl()
			}

		case goParserFUNC:
			{
				p.SetState(89)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) CopyAll(ctx *VariableDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerVarVDASTContext struct {
	VariableDeclContext
}

func NewInnerVarVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerVarVDASTContext {
	var p = new(InnerVarVDASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *InnerVarVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarVDASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(goParserVAR, 0)
}

func (s *InnerVarVDASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *InnerVarVDASTContext) InnerVarDecls() IInnerVarDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerVarDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerVarDeclsContext)
}

func (s *InnerVarVDASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *InnerVarVDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *InnerVarVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterInnerVarVDAST(s)
	}
}

func (s *InnerVarVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitInnerVarVDAST(s)
	}
}

func (s *InnerVarVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitInnerVarVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LpVDASTContext struct {
	VariableDeclContext
}

func NewLpVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LpVDASTContext {
	var p = new(LpVDASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *LpVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LpVDASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(goParserVAR, 0)
}

func (s *LpVDASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *LpVDASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *LpVDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *LpVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLpVDAST(s)
	}
}

func (s *LpVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLpVDAST(s)
	}
}

func (s *LpVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLpVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleVarVDASTContext struct {
	VariableDeclContext
}

func NewSingleVarVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarVDASTContext {
	var p = new(SingleVarVDASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *SingleVarVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarVDASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(goParserVAR, 0)
}

func (s *SingleVarVDASTContext) SingleVarDecl() ISingleVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *SingleVarVDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *SingleVarVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSingleVarVDAST(s)
	}
}

func (s *SingleVarVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSingleVarVDAST(s)
	}
}

func (s *SingleVarVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSingleVarVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, goParserRULE_variableDecl)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleVarVDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(goParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.SingleVarDecl()
		}
		{
			p.SetState(97)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewInnerVarVDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(goParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.InnerVarDecls()
		}
		{
			p.SetState(102)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewLpVDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(goParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerVarDeclsContext is an interface to support dynamic dispatch.
type IInnerVarDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInnerVarDeclsContext differentiates from other interfaces.
	IsInnerVarDeclsContext()
}

type InnerVarDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerVarDeclsContext() *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_innerVarDecls
	return p
}

func InitEmptyInnerVarDeclsContext(p *InnerVarDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_innerVarDecls
}

func (*InnerVarDeclsContext) IsInnerVarDeclsContext() {}

func NewInnerVarDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_innerVarDecls

	return p
}

func (s *InnerVarDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerVarDeclsContext) CopyAll(ctx *InnerVarDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InnerVarDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerVarDeclsASTContext struct {
	InnerVarDeclsContext
}

func NewInnerVarDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerVarDeclsASTContext {
	var p = new(InnerVarDeclsASTContext)

	InitEmptyInnerVarDeclsContext(&p.InnerVarDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InnerVarDeclsContext))

	return p
}

func (s *InnerVarDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsASTContext) AllSingleVarDecl() []ISingleVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclContext); ok {
			tst[i] = t.(ISingleVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerVarDeclsASTContext) SingleVarDecl(i int) ISingleVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *InnerVarDeclsASTContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(goParserSEMI)
}

func (s *InnerVarDeclsASTContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(goParserSEMI, i)
}

func (s *InnerVarDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterInnerVarDeclsAST(s)
	}
}

func (s *InnerVarDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitInnerVarDeclsAST(s)
	}
}

func (s *InnerVarDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitInnerVarDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) InnerVarDecls() (localctx IInnerVarDeclsContext) {
	localctx = NewInnerVarDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, goParserRULE_innerVarDecls)
	var _la int

	localctx = NewInnerVarDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.SingleVarDecl()
	}
	{
		p.SetState(112)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserID {
		{
			p.SetState(113)
			p.SingleVarDecl()
		}
		{
			p.SetState(114)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclContext is an interface to support dynamic dispatch.
type ISingleVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleVarDeclContext differentiates from other interfaces.
	IsSingleVarDeclContext()
}

type SingleVarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclContext() *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleVarDecl
	return p
}

func InitEmptySingleVarDeclContext(p *SingleVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleVarDecl
}

func (*SingleVarDeclContext) IsSingleVarDeclContext() {}

func NewSingleVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_singleVarDecl

	return p
}

func (s *SingleVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclContext) CopyAll(ctx *SingleVarDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdListDTSVDASTContext struct {
	SingleVarDeclContext
}

func NewIdListDTSVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdListDTSVDASTContext {
	var p = new(IdListDTSVDASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *IdListDTSVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListDTSVDASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *IdListDTSVDASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *IdListDTSVDASTContext) ASOP() antlr.TerminalNode {
	return s.GetToken(goParserASOP, 0)
}

func (s *IdListDTSVDASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *IdListDTSVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdListDTSVDAST(s)
	}
}

func (s *IdListDTSVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdListDTSVDAST(s)
	}
}

func (s *IdListDTSVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdListDTSVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdListSVDASTContext struct {
	SingleVarDeclContext
}

func NewIdListSVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdListSVDASTContext {
	var p = new(IdListSVDASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *IdListSVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListSVDASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *IdListSVDASTContext) ASOP() antlr.TerminalNode {
	return s.GetToken(goParserASOP, 0)
}

func (s *IdListSVDASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *IdListSVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdListSVDAST(s)
	}
}

func (s *IdListSVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdListSVDAST(s)
	}
}

func (s *IdListSVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdListSVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleVarDASTContext struct {
	SingleVarDeclContext
}

func NewSingleVarDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDASTContext {
	var p = new(SingleVarDASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *SingleVarDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDASTContext) SingleVarDeclNoExps() ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *SingleVarDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSingleVarDAST(s)
	}
}

func (s *SingleVarDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSingleVarDAST(s)
	}
}

func (s *SingleVarDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSingleVarDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) SingleVarDecl() (localctx ISingleVarDeclContext) {
	localctx = NewSingleVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, goParserRULE_singleVarDecl)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdListDTSVDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.IdentifierList()
		}
		{
			p.SetState(122)
			p.DeclType()
		}
		{
			p.SetState(123)
			p.Match(goParserASOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.ExpressionList()
		}

	case 2:
		localctx = NewIdListSVDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.IdentifierList()
		}
		{
			p.SetState(127)
			p.Match(goParserASOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.ExpressionList()
		}

	case 3:
		localctx = NewSingleVarDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.SingleVarDeclNoExps()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclNoExpsContext is an interface to support dynamic dispatch.
type ISingleVarDeclNoExpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleVarDeclNoExpsContext differentiates from other interfaces.
	IsSingleVarDeclNoExpsContext()
}

type SingleVarDeclNoExpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclNoExpsContext() *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleVarDeclNoExps
	return p
}

func InitEmptySingleVarDeclNoExpsContext(p *SingleVarDeclNoExpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleVarDeclNoExps
}

func (*SingleVarDeclNoExpsContext) IsSingleVarDeclNoExpsContext() {}

func NewSingleVarDeclNoExpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_singleVarDeclNoExps

	return p
}

func (s *SingleVarDeclNoExpsContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclNoExpsContext) CopyAll(ctx *SingleVarDeclNoExpsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclNoExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleVarDeclNoExpsASTContext struct {
	SingleVarDeclNoExpsContext
}

func NewSingleVarDeclNoExpsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleVarDeclNoExpsASTContext {
	var p = new(SingleVarDeclNoExpsASTContext)

	InitEmptySingleVarDeclNoExpsContext(&p.SingleVarDeclNoExpsContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclNoExpsContext))

	return p
}

func (s *SingleVarDeclNoExpsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsASTContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclNoExpsASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclNoExpsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSingleVarDeclNoExpsAST(s)
	}
}

func (s *SingleVarDeclNoExpsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSingleVarDeclNoExpsAST(s)
	}
}

func (s *SingleVarDeclNoExpsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSingleVarDeclNoExpsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) SingleVarDeclNoExps() (localctx ISingleVarDeclNoExpsContext) {
	localctx = NewSingleVarDeclNoExpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, goParserRULE_singleVarDeclNoExps)
	localctx = NewSingleVarDeclNoExpsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.IdentifierList()
	}
	{
		p.SetState(134)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) CopyAll(ctx *TypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerTypeDeclsTDASTContext struct {
	TypeDeclContext
}

func NewInnerTypeDeclsTDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerTypeDeclsTDASTContext {
	var p = new(InnerTypeDeclsTDASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *InnerTypeDeclsTDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsTDASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(goParserTYPE, 0)
}

func (s *InnerTypeDeclsTDASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *InnerTypeDeclsTDASTContext) InnerTypeDecls() IInnerTypeDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerTypeDeclsContext)
}

func (s *InnerTypeDeclsTDASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *InnerTypeDeclsTDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *InnerTypeDeclsTDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterInnerTypeDeclsTDAST(s)
	}
}

func (s *InnerTypeDeclsTDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitInnerTypeDeclsTDAST(s)
	}
}

func (s *InnerTypeDeclsTDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitInnerTypeDeclsTDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SingleTypeDeclTDASTContext struct {
	TypeDeclContext
}

func NewSingleTypeDeclTDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleTypeDeclTDASTContext {
	var p = new(SingleTypeDeclTDASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *SingleTypeDeclTDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclTDASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(goParserTYPE, 0)
}

func (s *SingleTypeDeclTDASTContext) SingleTypeDecl() ISingleTypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *SingleTypeDeclTDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *SingleTypeDeclTDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSingleTypeDeclTDAST(s)
	}
}

func (s *SingleTypeDeclTDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSingleTypeDeclTDAST(s)
	}
}

func (s *SingleTypeDeclTDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSingleTypeDeclTDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LPTypeDeclTDASTContext struct {
	TypeDeclContext
}

func NewLPTypeDeclTDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LPTypeDeclTDASTContext {
	var p = new(LPTypeDeclTDASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *LPTypeDeclTDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LPTypeDeclTDASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(goParserTYPE, 0)
}

func (s *LPTypeDeclTDASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *LPTypeDeclTDASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *LPTypeDeclTDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *LPTypeDeclTDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLPTypeDeclTDAST(s)
	}
}

func (s *LPTypeDeclTDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLPTypeDeclTDAST(s)
	}
}

func (s *LPTypeDeclTDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLPTypeDeclTDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, goParserRULE_typeDecl)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleTypeDeclTDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Match(goParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.SingleTypeDecl()
		}
		{
			p.SetState(138)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewInnerTypeDeclsTDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(goParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.InnerTypeDecls()
		}
		{
			p.SetState(143)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewLPTypeDeclTDASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Match(goParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerTypeDeclsContext is an interface to support dynamic dispatch.
type IInnerTypeDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInnerTypeDeclsContext differentiates from other interfaces.
	IsInnerTypeDeclsContext()
}

type InnerTypeDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerTypeDeclsContext() *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_innerTypeDecls
	return p
}

func InitEmptyInnerTypeDeclsContext(p *InnerTypeDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_innerTypeDecls
}

func (*InnerTypeDeclsContext) IsInnerTypeDeclsContext() {}

func NewInnerTypeDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_innerTypeDecls

	return p
}

func (s *InnerTypeDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeDeclsContext) CopyAll(ctx *InnerTypeDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InnerTypeDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InnerTypeDeclsASTContext struct {
	InnerTypeDeclsContext
}

func NewInnerTypeDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerTypeDeclsASTContext {
	var p = new(InnerTypeDeclsASTContext)

	InitEmptyInnerTypeDeclsContext(&p.InnerTypeDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InnerTypeDeclsContext))

	return p
}

func (s *InnerTypeDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsASTContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleTypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleTypeDeclContext); ok {
			tst[i] = t.(ISingleTypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerTypeDeclsASTContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *InnerTypeDeclsASTContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(goParserSEMI)
}

func (s *InnerTypeDeclsASTContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(goParserSEMI, i)
}

func (s *InnerTypeDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterInnerTypeDeclsAST(s)
	}
}

func (s *InnerTypeDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitInnerTypeDeclsAST(s)
	}
}

func (s *InnerTypeDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitInnerTypeDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, goParserRULE_innerTypeDecls)
	var _la int

	localctx = NewInnerTypeDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.SingleTypeDecl()
	}
	{
		p.SetState(153)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserID {
		{
			p.SetState(154)
			p.SingleTypeDecl()
		}
		{
			p.SetState(155)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleTypeDeclContext is an interface to support dynamic dispatch.
type ISingleTypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleTypeDeclContext differentiates from other interfaces.
	IsSingleTypeDeclContext()
}

type SingleTypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTypeDeclContext() *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleTypeDecl
	return p
}

func InitEmptySingleTypeDeclContext(p *SingleTypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_singleTypeDecl
}

func (*SingleTypeDeclContext) IsSingleTypeDeclContext() {}

func NewSingleTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_singleTypeDecl

	return p
}

func (s *SingleTypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTypeDeclContext) CopyAll(ctx *SingleTypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdSYDASTContext struct {
	SingleTypeDeclContext
}

func NewIdSYDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdSYDASTContext {
	var p = new(IdSYDASTContext)

	InitEmptySingleTypeDeclContext(&p.SingleTypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleTypeDeclContext))

	return p
}

func (s *IdSYDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdSYDASTContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *IdSYDASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *IdSYDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdSYDAST(s)
	}
}

func (s *IdSYDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdSYDAST(s)
	}
}

func (s *IdSYDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdSYDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) SingleTypeDecl() (localctx ISingleTypeDeclContext) {
	localctx = NewSingleTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, goParserRULE_singleTypeDecl)
	localctx = NewIdSYDASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) CopyAll(ctx *FuncDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncDeclASTContext struct {
	FuncDeclContext
}

func NewFuncDeclASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncDeclASTContext {
	var p = new(FuncDeclASTContext)

	InitEmptyFuncDeclContext(&p.FuncDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncDeclContext))

	return p
}

func (s *FuncDeclASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclASTContext) FuncFrontDecl() IFuncFrontDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFrontDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFrontDeclContext)
}

func (s *FuncDeclASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *FuncDeclASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFuncDeclAST(s)
	}
}

func (s *FuncDeclASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFuncDeclAST(s)
	}
}

func (s *FuncDeclASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFuncDeclAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, goParserRULE_funcDecl)
	localctx = NewFuncDeclASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.FuncFrontDecl()
	}
	{
		p.SetState(166)
		p.Block()
	}
	{
		p.SetState(167)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncFrontDeclContext is an interface to support dynamic dispatch.
type IFuncFrontDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncFrontDeclContext differentiates from other interfaces.
	IsFuncFrontDeclContext()
}

type FuncFrontDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFrontDeclContext() *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_funcFrontDecl
	return p
}

func InitEmptyFuncFrontDeclContext(p *FuncFrontDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_funcFrontDecl
}

func (*FuncFrontDeclContext) IsFuncFrontDeclContext() {}

func NewFuncFrontDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_funcFrontDecl

	return p
}

func (s *FuncFrontDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFrontDeclContext) CopyAll(ctx *FuncFrontDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncFrontDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncFFDASTContext struct {
	FuncFrontDeclContext
}

func NewFuncFFDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncFFDASTContext {
	var p = new(FuncFFDASTContext)

	InitEmptyFuncFrontDeclContext(&p.FuncFrontDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncFrontDeclContext))

	return p
}

func (s *FuncFFDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFFDASTContext) FUNC() antlr.TerminalNode {
	return s.GetToken(goParserFUNC, 0)
}

func (s *FuncFFDASTContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *FuncFFDASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *FuncFFDASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *FuncFFDASTContext) FuncArgDecls() IFuncArgDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgDeclsContext)
}

func (s *FuncFFDASTContext) AllEpsilon() []IEpsilonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEpsilonContext); ok {
			len++
		}
	}

	tst := make([]IEpsilonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEpsilonContext); ok {
			tst[i] = t.(IEpsilonContext)
			i++
		}
	}

	return tst
}

func (s *FuncFFDASTContext) Epsilon(i int) IEpsilonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *FuncFFDASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *FuncFFDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFuncFFDAST(s)
	}
}

func (s *FuncFFDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFuncFFDAST(s)
	}
}

func (s *FuncFFDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFuncFFDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) FuncFrontDecl() (localctx IFuncFrontDeclContext) {
	localctx = NewFuncFrontDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, goParserRULE_funcFrontDecl)
	localctx = NewFuncFFDASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(goParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(goParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserID:
		{
			p.SetState(172)
			p.FuncArgDecls()
		}

	case goParserRP:
		{
			p.SetState(173)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(176)
		p.Match(goParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserLP, goParserLSB, goParserSTRUCT, goParserID:
		{
			p.SetState(177)
			p.DeclType()
		}

	case goParserLCB:
		{
			p.SetState(178)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncArgDeclsContext is an interface to support dynamic dispatch.
type IFuncArgDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFuncArgDeclsContext differentiates from other interfaces.
	IsFuncArgDeclsContext()
}

type FuncArgDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgDeclsContext() *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_funcArgDecls
	return p
}

func InitEmptyFuncArgDeclsContext(p *FuncArgDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_funcArgDecls
}

func (*FuncArgDeclsContext) IsFuncArgDeclsContext() {}

func NewFuncArgDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_funcArgDecls

	return p
}

func (s *FuncArgDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgDeclsContext) CopyAll(ctx *FuncArgDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FuncArgDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncArgDeclsASTContext struct {
	FuncArgDeclsContext
}

func NewFuncArgDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArgDeclsASTContext {
	var p = new(FuncArgDeclsASTContext)

	InitEmptyFuncArgDeclsContext(&p.FuncArgDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FuncArgDeclsContext))

	return p
}

func (s *FuncArgDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsASTContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgDeclsASTContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *FuncArgDeclsASTContext) AllCM() []antlr.TerminalNode {
	return s.GetTokens(goParserCM)
}

func (s *FuncArgDeclsASTContext) CM(i int) antlr.TerminalNode {
	return s.GetToken(goParserCM, i)
}

func (s *FuncArgDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFuncArgDeclsAST(s)
	}
}

func (s *FuncArgDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFuncArgDeclsAST(s)
	}
}

func (s *FuncArgDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFuncArgDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) FuncArgDecls() (localctx IFuncArgDeclsContext) {
	localctx = NewFuncArgDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, goParserRULE_funcArgDecls)
	var _la int

	localctx = NewFuncArgDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.SingleVarDeclNoExps()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserCM {
		{
			p.SetState(182)
			p.Match(goParserCM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.SingleVarDeclNoExps()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclTypeContext is an interface to support dynamic dispatch.
type IDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclTypeContext differentiates from other interfaces.
	IsDeclTypeContext()
}

type DeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclTypeContext() *DeclTypeContext {
	var p = new(DeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_declType
	return p
}

func InitEmptyDeclTypeContext(p *DeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_declType
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_declType

	return p
}

func (s *DeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclTypeContext) CopyAll(ctx *DeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDeclTypeDTASTContext struct {
	DeclTypeContext
}

func NewStructDeclTypeDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDeclTypeDTASTContext {
	var p = new(StructDeclTypeDTASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *StructDeclTypeDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeDTASTContext) StructDeclType() IStructDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclTypeContext)
}

func (s *StructDeclTypeDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterStructDeclTypeDTAST(s)
	}
}

func (s *StructDeclTypeDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitStructDeclTypeDTAST(s)
	}
}

func (s *StructDeclTypeDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitStructDeclTypeDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceDeclTypeDTASTContext struct {
	DeclTypeContext
}

func NewSliceDeclTypeDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceDeclTypeDTASTContext {
	var p = new(SliceDeclTypeDTASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *SliceDeclTypeDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeDTASTContext) SliceDeclType() ISliceDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDeclTypeContext)
}

func (s *SliceDeclTypeDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSliceDeclTypeDTAST(s)
	}
}

func (s *SliceDeclTypeDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSliceDeclTypeDTAST(s)
	}
}

func (s *SliceDeclTypeDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSliceDeclTypeDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayDeclTypeDTASTContext struct {
	DeclTypeContext
}

func NewArrayDeclTypeDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayDeclTypeDTASTContext {
	var p = new(ArrayDeclTypeDTASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *ArrayDeclTypeDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeDTASTContext) ArrayDeclType() IArrayDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclTypeContext)
}

func (s *ArrayDeclTypeDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterArrayDeclTypeDTAST(s)
	}
}

func (s *ArrayDeclTypeDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitArrayDeclTypeDTAST(s)
	}
}

func (s *ArrayDeclTypeDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitArrayDeclTypeDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LpDTASTContext struct {
	DeclTypeContext
}

func NewLpDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LpDTASTContext {
	var p = new(LpDTASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *LpDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LpDTASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *LpDTASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *LpDTASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *LpDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLpDTAST(s)
	}
}

func (s *LpDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLpDTAST(s)
	}
}

func (s *LpDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLpDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdDTASTContext struct {
	DeclTypeContext
}

func NewIdDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdDTASTContext {
	var p = new(IdDTASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *IdDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdDTASTContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *IdDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdDTAST(s)
	}
}

func (s *IdDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdDTAST(s)
	}
}

func (s *IdDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, goParserRULE_declType)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLpDTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.DeclType()
		}
		{
			p.SetState(191)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewIdDTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(goParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSliceDeclTypeDTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.SliceDeclType()
		}

	case 4:
		localctx = NewArrayDeclTypeDTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.ArrayDeclType()
		}

	case 5:
		localctx = NewStructDeclTypeDTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(196)
			p.StructDeclType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceDeclTypeContext is an interface to support dynamic dispatch.
type ISliceDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSliceDeclTypeContext differentiates from other interfaces.
	IsSliceDeclTypeContext()
}

type SliceDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDeclTypeContext() *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_sliceDeclType
	return p
}

func InitEmptySliceDeclTypeContext(p *SliceDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_sliceDeclType
}

func (*SliceDeclTypeContext) IsSliceDeclTypeContext() {}

func NewSliceDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_sliceDeclType

	return p
}

func (s *SliceDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDeclTypeContext) CopyAll(ctx *SliceDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SliceDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LsbSDTASTContext struct {
	SliceDeclTypeContext
}

func NewLsbSDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LsbSDTASTContext {
	var p = new(LsbSDTASTContext)

	InitEmptySliceDeclTypeContext(&p.SliceDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SliceDeclTypeContext))

	return p
}

func (s *LsbSDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LsbSDTASTContext) LSB() antlr.TerminalNode {
	return s.GetToken(goParserLSB, 0)
}

func (s *LsbSDTASTContext) RSB() antlr.TerminalNode {
	return s.GetToken(goParserRSB, 0)
}

func (s *LsbSDTASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *LsbSDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLsbSDTAST(s)
	}
}

func (s *LsbSDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLsbSDTAST(s)
	}
}

func (s *LsbSDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLsbSDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) SliceDeclType() (localctx ISliceDeclTypeContext) {
	localctx = NewSliceDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, goParserRULE_sliceDeclType)
	localctx = NewLsbSDTASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(goParserLSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Match(goParserRSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayDeclTypeContext is an interface to support dynamic dispatch.
type IArrayDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArrayDeclTypeContext differentiates from other interfaces.
	IsArrayDeclTypeContext()
}

type ArrayDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclTypeContext() *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_arrayDeclType
	return p
}

func InitEmptyArrayDeclTypeContext(p *ArrayDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_arrayDeclType
}

func (*ArrayDeclTypeContext) IsArrayDeclTypeContext() {}

func NewArrayDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_arrayDeclType

	return p
}

func (s *ArrayDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclTypeContext) CopyAll(ctx *ArrayDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LsbADTASTContext struct {
	ArrayDeclTypeContext
}

func NewLsbADTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LsbADTASTContext {
	var p = new(LsbADTASTContext)

	InitEmptyArrayDeclTypeContext(&p.ArrayDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayDeclTypeContext))

	return p
}

func (s *LsbADTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LsbADTASTContext) LSB() antlr.TerminalNode {
	return s.GetToken(goParserLSB, 0)
}

func (s *LsbADTASTContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(goParserINTLITERAL, 0)
}

func (s *LsbADTASTContext) RSB() antlr.TerminalNode {
	return s.GetToken(goParserRSB, 0)
}

func (s *LsbADTASTContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *LsbADTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLsbADTAST(s)
	}
}

func (s *LsbADTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLsbADTAST(s)
	}
}

func (s *LsbADTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLsbADTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) ArrayDeclType() (localctx IArrayDeclTypeContext) {
	localctx = NewArrayDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, goParserRULE_arrayDeclType)
	localctx = NewLsbADTASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(goParserLSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(goParserINTLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(goParserRSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclTypeContext is an interface to support dynamic dispatch.
type IStructDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructDeclTypeContext differentiates from other interfaces.
	IsStructDeclTypeContext()
}

type StructDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclTypeContext() *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_structDeclType
	return p
}

func InitEmptyStructDeclTypeContext(p *StructDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_structDeclType
}

func (*StructDeclTypeContext) IsStructDeclTypeContext() {}

func NewStructDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_structDeclType

	return p
}

func (s *StructDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclTypeContext) CopyAll(ctx *StructDeclTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructSDTASTContext struct {
	StructDeclTypeContext
}

func NewStructSDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructSDTASTContext {
	var p = new(StructSDTASTContext)

	InitEmptyStructDeclTypeContext(&p.StructDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructDeclTypeContext))

	return p
}

func (s *StructSDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructSDTASTContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(goParserSTRUCT, 0)
}

func (s *StructSDTASTContext) LCB() antlr.TerminalNode {
	return s.GetToken(goParserLCB, 0)
}

func (s *StructSDTASTContext) RCB() antlr.TerminalNode {
	return s.GetToken(goParserRCB, 0)
}

func (s *StructSDTASTContext) StructMemDecls() IStructMemDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructMemDeclsContext)
}

func (s *StructSDTASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *StructSDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterStructSDTAST(s)
	}
}

func (s *StructSDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitStructSDTAST(s)
	}
}

func (s *StructSDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitStructSDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, goParserRULE_structDeclType)
	localctx = NewStructSDTASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(goParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(goParserLCB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserID:
		{
			p.SetState(210)
			p.StructMemDecls()
		}

	case goParserRCB:
		{
			p.SetState(211)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(214)
		p.Match(goParserRCB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructMemDeclsContext is an interface to support dynamic dispatch.
type IStructMemDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructMemDeclsContext differentiates from other interfaces.
	IsStructMemDeclsContext()
}

type StructMemDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemDeclsContext() *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_structMemDecls
	return p
}

func InitEmptyStructMemDeclsContext(p *StructMemDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_structMemDecls
}

func (*StructMemDeclsContext) IsStructMemDeclsContext() {}

func NewStructMemDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_structMemDecls

	return p
}

func (s *StructMemDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemDeclsContext) CopyAll(ctx *StructMemDeclsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructMemDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructMemDeclsASTContext struct {
	StructMemDeclsContext
}

func NewStructMemDeclsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructMemDeclsASTContext {
	var p = new(StructMemDeclsASTContext)

	InitEmptyStructMemDeclsContext(&p.StructMemDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructMemDeclsContext))

	return p
}

func (s *StructMemDeclsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsASTContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *StructMemDeclsASTContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *StructMemDeclsASTContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(goParserSEMI)
}

func (s *StructMemDeclsASTContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(goParserSEMI, i)
}

func (s *StructMemDeclsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterStructMemDeclsAST(s)
	}
}

func (s *StructMemDeclsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitStructMemDeclsAST(s)
	}
}

func (s *StructMemDeclsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitStructMemDeclsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) StructMemDecls() (localctx IStructMemDeclsContext) {
	localctx = NewStructMemDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, goParserRULE_structMemDecls)
	var _la int

	localctx = NewStructMemDeclsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(217)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserID {
		{
			p.SetState(218)
			p.SingleVarDeclNoExps()
		}
		{
			p.SetState(219)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) CopyAll(ctx *IdentifierListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierListASTContext struct {
	IdentifierListContext
}

func NewIdentifierListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierListASTContext {
	var p = new(IdentifierListASTContext)

	InitEmptyIdentifierListContext(&p.IdentifierListContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentifierListContext))

	return p
}

func (s *IdentifierListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListASTContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(goParserID)
}

func (s *IdentifierListASTContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(goParserID, i)
}

func (s *IdentifierListASTContext) AllCM() []antlr.TerminalNode {
	return s.GetTokens(goParserCM)
}

func (s *IdentifierListASTContext) CM(i int) antlr.TerminalNode {
	return s.GetToken(goParserCM, i)
}

func (s *IdentifierListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdentifierListAST(s)
	}
}

func (s *IdentifierListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdentifierListAST(s)
	}
}

func (s *IdentifierListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdentifierListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, goParserRULE_identifierList)
	var _la int

	localctx = NewIdentifierListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserCM {
		{
			p.SetState(227)
			p.Match(goParserCM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(goParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionEASTContext struct {
	ExpressionContext
}

func NewExpressionEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionEASTContext {
	var p = new(ExpressionEASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionEASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionEASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionEASTContext) MUL() antlr.TerminalNode {
	return s.GetToken(goParserMUL, 0)
}

func (s *ExpressionEASTContext) DIV() antlr.TerminalNode {
	return s.GetToken(goParserDIV, 0)
}

func (s *ExpressionEASTContext) MOD() antlr.TerminalNode {
	return s.GetToken(goParserMOD, 0)
}

func (s *ExpressionEASTContext) LSH() antlr.TerminalNode {
	return s.GetToken(goParserLSH, 0)
}

func (s *ExpressionEASTContext) RSH() antlr.TerminalNode {
	return s.GetToken(goParserRSH, 0)
}

func (s *ExpressionEASTContext) AMPER() antlr.TerminalNode {
	return s.GetToken(goParserAMPER, 0)
}

func (s *ExpressionEASTContext) BC() antlr.TerminalNode {
	return s.GetToken(goParserBC, 0)
}

func (s *ExpressionEASTContext) PL() antlr.TerminalNode {
	return s.GetToken(goParserPL, 0)
}

func (s *ExpressionEASTContext) MIN() antlr.TerminalNode {
	return s.GetToken(goParserMIN, 0)
}

func (s *ExpressionEASTContext) VB() antlr.TerminalNode {
	return s.GetToken(goParserVB, 0)
}

func (s *ExpressionEASTContext) CARET() antlr.TerminalNode {
	return s.GetToken(goParserCARET, 0)
}

func (s *ExpressionEASTContext) EQ() antlr.TerminalNode {
	return s.GetToken(goParserEQ, 0)
}

func (s *ExpressionEASTContext) NEQ() antlr.TerminalNode {
	return s.GetToken(goParserNEQ, 0)
}

func (s *ExpressionEASTContext) LT() antlr.TerminalNode {
	return s.GetToken(goParserLT, 0)
}

func (s *ExpressionEASTContext) LTE() antlr.TerminalNode {
	return s.GetToken(goParserLTE, 0)
}

func (s *ExpressionEASTContext) GT() antlr.TerminalNode {
	return s.GetToken(goParserGT, 0)
}

func (s *ExpressionEASTContext) GTE() antlr.TerminalNode {
	return s.GetToken(goParserGTE, 0)
}

func (s *ExpressionEASTContext) LAND() antlr.TerminalNode {
	return s.GetToken(goParserLAND, 0)
}

func (s *ExpressionEASTContext) LOR() antlr.TerminalNode {
	return s.GetToken(goParserLOR, 0)
}

func (s *ExpressionEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionEAST(s)
	}
}

func (s *ExpressionEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionEAST(s)
	}
}

func (s *ExpressionEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionEASTContext struct {
	ExpressionContext
}

func NewPrimaryExpressionEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionEASTContext {
	var p = new(PrimaryExpressionEASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PrimaryExpressionEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionEASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterPrimaryExpressionEAST(s)
	}
}

func (s *PrimaryExpressionEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitPrimaryExpressionEAST(s)
	}
}

func (s *PrimaryExpressionEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitPrimaryExpressionEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperatorEASTContext struct {
	ExpressionContext
}

func NewOperatorEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperatorEASTContext {
	var p = new(OperatorEASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OperatorEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorEASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperatorEASTContext) PL() antlr.TerminalNode {
	return s.GetToken(goParserPL, 0)
}

func (s *OperatorEASTContext) MIN() antlr.TerminalNode {
	return s.GetToken(goParserMIN, 0)
}

func (s *OperatorEASTContext) NEG() antlr.TerminalNode {
	return s.GetToken(goParserNEG, 0)
}

func (s *OperatorEASTContext) CARET() antlr.TerminalNode {
	return s.GetToken(goParserCARET, 0)
}

func (s *OperatorEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterOperatorEAST(s)
	}
}

func (s *OperatorEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitOperatorEAST(s)
	}
}

func (s *OperatorEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitOperatorEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *goParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, goParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserLP, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
		localctx = NewPrimaryExpressionEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(235)
			p.primaryExpression(0)
		}

	case goParserPL, goParserMIN, goParserCARET, goParserNEG:
		localctx = NewOperatorEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(236)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8606712192) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(237)
			p.expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionEASTContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, goParserRULE_expression)
			p.SetState(240)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(241)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589674368) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(242)
				p.expression(3)
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) CopyAll(ctx *ExpressionListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionListASTContext struct {
	ExpressionListContext
}

func NewExpressionListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListASTContext {
	var p = new(ExpressionListASTContext)

	InitEmptyExpressionListContext(&p.ExpressionListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionListContext))

	return p
}

func (s *ExpressionListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListASTContext) AllCM() []antlr.TerminalNode {
	return s.GetTokens(goParserCM)
}

func (s *ExpressionListASTContext) CM(i int) antlr.TerminalNode {
	return s.GetToken(goParserCM, i)
}

func (s *ExpressionListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionListAST(s)
	}
}

func (s *ExpressionListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionListAST(s)
	}
}

func (s *ExpressionListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, goParserRULE_expressionList)
	var _la int

	localctx = NewExpressionListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.expression(0)
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserCM {
		{
			p.SetState(249)
			p.Match(goParserCM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.expression(0)
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyAll(ctx *PrimaryExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryExpArgumentsPEASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpArgumentsPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpArgumentsPEASTContext {
	var p = new(PrimaryExpArgumentsPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpArgumentsPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpArgumentsPEASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpArgumentsPEASTContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExpArgumentsPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterPrimaryExpArgumentsPEAST(s)
	}
}

func (s *PrimaryExpArgumentsPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitPrimaryExpArgumentsPEAST(s)
	}
}

func (s *PrimaryExpArgumentsPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitPrimaryExpArgumentsPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpSelectorPEASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpSelectorPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpSelectorPEASTContext {
	var p = new(PrimaryExpSelectorPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpSelectorPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpSelectorPEASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpSelectorPEASTContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PrimaryExpSelectorPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterPrimaryExpSelectorPEAST(s)
	}
}

func (s *PrimaryExpSelectorPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitPrimaryExpSelectorPEAST(s)
	}
}

func (s *PrimaryExpSelectorPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitPrimaryExpSelectorPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LengthExpPEASTContext struct {
	PrimaryExpressionContext
}

func NewLengthExpPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthExpPEASTContext {
	var p = new(LengthExpPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *LengthExpPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpPEASTContext) LengthExpression() ILengthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthExpressionContext)
}

func (s *LengthExpPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLengthExpPEAST(s)
	}
}

func (s *LengthExpPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLengthExpPEAST(s)
	}
}

func (s *LengthExpPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLengthExpPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendExpPEASTContext struct {
	PrimaryExpressionContext
}

func NewAppendExpPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExpPEASTContext {
	var p = new(AppendExpPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *AppendExpPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpPEASTContext) AppendExpression() IAppendExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendExpressionContext)
}

func (s *AppendExpPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterAppendExpPEAST(s)
	}
}

func (s *AppendExpPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitAppendExpPEAST(s)
	}
}

func (s *AppendExpPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitAppendExpPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type OperandPEASTContext struct {
	PrimaryExpressionContext
}

func NewOperandPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandPEASTContext {
	var p = new(OperandPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *OperandPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandPEASTContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *OperandPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterOperandPEAST(s)
	}
}

func (s *OperandPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitOperandPEAST(s)
	}
}

func (s *OperandPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitOperandPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type CapExpPEASTContext struct {
	PrimaryExpressionContext
}

func NewCapExpPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapExpPEASTContext {
	var p = new(CapExpPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *CapExpPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpPEASTContext) CapExpression() ICapExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapExpressionContext)
}

func (s *CapExpPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterCapExpPEAST(s)
	}
}

func (s *CapExpPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitCapExpPEAST(s)
	}
}

func (s *CapExpPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitCapExpPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpIndexPEASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpIndexPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpIndexPEASTContext {
	var p = new(PrimaryExpIndexPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpIndexPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpIndexPEASTContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpIndexPEASTContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExpIndexPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterPrimaryExpIndexPEAST(s)
	}
}

func (s *PrimaryExpIndexPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitPrimaryExpIndexPEAST(s)
	}
}

func (s *PrimaryExpIndexPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitPrimaryExpIndexPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *goParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, goParserRULE_primaryExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserLP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
		localctx = NewOperandPEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(257)
			p.Operand()
		}

	case goParserAPPEND:
		localctx = NewAppendExpPEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.AppendExpression()
		}

	case goParserLEN:
		localctx = NewLengthExpPEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.LengthExpression()
		}

	case goParserCAP:
		localctx = NewCapExpPEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExpSelectorPEASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goParserRULE_primaryExpression)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(264)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExpIndexPEASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goParserRULE_primaryExpression)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(266)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExpArgumentsPEASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, goParserRULE_primaryExpression)
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(268)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) CopyAll(ctx *OperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralOASTContext struct {
	OperandContext
}

func NewLiteralOASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralOASTContext {
	var p = new(LiteralOASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *LiteralOASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOASTContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralOASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLiteralOAST(s)
	}
}

func (s *LiteralOASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLiteralOAST(s)
	}
}

func (s *LiteralOASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLiteralOAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdOASTContext struct {
	OperandContext
}

func NewIdOASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdOASTContext {
	var p = new(IdOASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *IdOASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdOASTContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *IdOASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdOAST(s)
	}
}

func (s *IdOASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdOAST(s)
	}
}

func (s *IdOASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdOAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionOASTContext struct {
	OperandContext
}

func NewExpressionOASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionOASTContext {
	var p = new(ExpressionOASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *ExpressionOASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionOASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *ExpressionOASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionOASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *ExpressionOASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionOAST(s)
	}
}

func (s *ExpressionOASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionOAST(s)
	}
}

func (s *ExpressionOASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionOAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, goParserRULE_operand)
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
		localctx = NewLiteralOASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Literal()
		}

	case goParserID:
		localctx = NewIdOASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.Match(goParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserLP:
		localctx = NewExpressionOASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.expression(0)
		}
		{
			p.SetState(278)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntliteralASTContext struct {
	LiteralContext
}

func NewIntliteralASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntliteralASTContext {
	var p = new(IntliteralASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntliteralASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntliteralASTContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(goParserINTLITERAL, 0)
}

func (s *IntliteralASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIntliteralAST(s)
	}
}

func (s *IntliteralASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIntliteralAST(s)
	}
}

func (s *IntliteralASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIntliteralAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type InterpretedliteralASTContext struct {
	LiteralContext
}

func NewInterpretedliteralASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InterpretedliteralASTContext {
	var p = new(InterpretedliteralASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *InterpretedliteralASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpretedliteralASTContext) INTERPRETEDSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(goParserINTERPRETEDSTRINGLITERAL, 0)
}

func (s *InterpretedliteralASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterInterpretedliteralAST(s)
	}
}

func (s *InterpretedliteralASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitInterpretedliteralAST(s)
	}
}

func (s *InterpretedliteralASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitInterpretedliteralAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type RunliteralASTContext struct {
	LiteralContext
}

func NewRunliteralASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunliteralASTContext {
	var p = new(RunliteralASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RunliteralASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunliteralASTContext) RUNELITERAL() antlr.TerminalNode {
	return s.GetToken(goParserRUNELITERAL, 0)
}

func (s *RunliteralASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterRunliteralAST(s)
	}
}

func (s *RunliteralASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitRunliteralAST(s)
	}
}

func (s *RunliteralASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitRunliteralAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatliteralASTContext struct {
	LiteralContext
}

func NewFloatliteralASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatliteralASTContext {
	var p = new(FloatliteralASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatliteralASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatliteralASTContext) FLOATLITERAL() antlr.TerminalNode {
	return s.GetToken(goParserFLOATLITERAL, 0)
}

func (s *FloatliteralASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFloatliteralAST(s)
	}
}

func (s *FloatliteralASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFloatliteralAST(s)
	}
}

func (s *FloatliteralASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFloatliteralAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type RawsliteralASTContext struct {
	LiteralContext
}

func NewRawsliteralASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RawsliteralASTContext {
	var p = new(RawsliteralASTContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RawsliteralASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawsliteralASTContext) RAWSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(goParserRAWSTRINGLITERAL, 0)
}

func (s *RawsliteralASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterRawsliteralAST(s)
	}
}

func (s *RawsliteralASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitRawsliteralAST(s)
	}
}

func (s *RawsliteralASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitRawsliteralAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, goParserRULE_literal)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserINTLITERAL:
		localctx = NewIntliteralASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.Match(goParserINTLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserFLOATLITERAL:
		localctx = NewFloatliteralASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)
			p.Match(goParserFLOATLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserRUNELITERAL:
		localctx = NewRunliteralASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.Match(goParserRUNELITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserRAWSTRINGLITERAL:
		localctx = NewRawsliteralASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(285)
			p.Match(goParserRAWSTRINGLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserINTERPRETEDSTRINGLITERAL:
		localctx = NewInterpretedliteralASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(286)
			p.Match(goParserINTERPRETEDSTRINGLITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) CopyAll(ctx *IndexContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdexASTContext struct {
	IndexContext
}

func NewIdexASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdexASTContext {
	var p = new(IdexASTContext)

	InitEmptyIndexContext(&p.IndexContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexContext))

	return p
}

func (s *IdexASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdexASTContext) LSB() antlr.TerminalNode {
	return s.GetToken(goParserLSB, 0)
}

func (s *IdexASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IdexASTContext) RSB() antlr.TerminalNode {
	return s.GetToken(goParserRSB, 0)
}

func (s *IdexASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdexAST(s)
	}
}

func (s *IdexASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdexAST(s)
	}
}

func (s *IdexASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdexAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, goParserRULE_index)
	localctx = NewIdexASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(goParserLSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.expression(0)
	}
	{
		p.SetState(291)
		p.Match(goParserRSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) CopyAll(ctx *ArgumentsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgumentsASTContext struct {
	ArgumentsContext
}

func NewArgumentsASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsASTContext {
	var p = new(ArgumentsASTContext)

	InitEmptyArgumentsContext(&p.ArgumentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArgumentsContext))

	return p
}

func (s *ArgumentsASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *ArgumentsASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *ArgumentsASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *ArgumentsASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterArgumentsAST(s)
	}
}

func (s *ArgumentsASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitArgumentsAST(s)
	}
}

func (s *ArgumentsASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitArgumentsAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, goParserRULE_arguments)
	localctx = NewArgumentsASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(goParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserLP, goParserPL, goParserMIN, goParserCARET, goParserNEG, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
		{
			p.SetState(294)
			p.ExpressionList()
		}

	case goParserRP:
		{
			p.SetState(295)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(298)
		p.Match(goParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) CopyAll(ctx *SelectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectorASTContext struct {
	SelectorContext
}

func NewSelectorASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorASTContext {
	var p = new(SelectorASTContext)

	InitEmptySelectorContext(&p.SelectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*SelectorContext))

	return p
}

func (s *SelectorASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorASTContext) DOT() antlr.TerminalNode {
	return s.GetToken(goParserDOT, 0)
}

func (s *SelectorASTContext) ID() antlr.TerminalNode {
	return s.GetToken(goParserID, 0)
}

func (s *SelectorASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSelectorAST(s)
	}
}

func (s *SelectorASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSelectorAST(s)
	}
}

func (s *SelectorASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSelectorAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, goParserRULE_selector)
	localctx = NewSelectorASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(goParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAppendExpressionContext is an interface to support dynamic dispatch.
type IAppendExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAppendExpressionContext differentiates from other interfaces.
	IsAppendExpressionContext()
}

type AppendExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendExpressionContext() *AppendExpressionContext {
	var p = new(AppendExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_appendExpression
	return p
}

func InitEmptyAppendExpressionContext(p *AppendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_appendExpression
}

func (*AppendExpressionContext) IsAppendExpressionContext() {}

func NewAppendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendExpressionContext {
	var p = new(AppendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_appendExpression

	return p
}

func (s *AppendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendExpressionContext) CopyAll(ctx *AppendExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AppendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AppendExpressionASTContext struct {
	AppendExpressionContext
}

func NewAppendExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExpressionASTContext {
	var p = new(AppendExpressionASTContext)

	InitEmptyAppendExpressionContext(&p.AppendExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AppendExpressionContext))

	return p
}

func (s *AppendExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionASTContext) APPEND() antlr.TerminalNode {
	return s.GetToken(goParserAPPEND, 0)
}

func (s *AppendExpressionASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *AppendExpressionASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AppendExpressionASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AppendExpressionASTContext) CM() antlr.TerminalNode {
	return s.GetToken(goParserCM, 0)
}

func (s *AppendExpressionASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *AppendExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterAppendExpressionAST(s)
	}
}

func (s *AppendExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitAppendExpressionAST(s)
	}
}

func (s *AppendExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitAppendExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) AppendExpression() (localctx IAppendExpressionContext) {
	localctx = NewAppendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, goParserRULE_appendExpression)
	localctx = NewAppendExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(goParserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Match(goParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.expression(0)
	}
	{
		p.SetState(306)
		p.Match(goParserCM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.expression(0)
	}
	{
		p.SetState(308)
		p.Match(goParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILengthExpressionContext is an interface to support dynamic dispatch.
type ILengthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLengthExpressionContext differentiates from other interfaces.
	IsLengthExpressionContext()
}

type LengthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthExpressionContext() *LengthExpressionContext {
	var p = new(LengthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_lengthExpression
	return p
}

func InitEmptyLengthExpressionContext(p *LengthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_lengthExpression
}

func (*LengthExpressionContext) IsLengthExpressionContext() {}

func NewLengthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthExpressionContext {
	var p = new(LengthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_lengthExpression

	return p
}

func (s *LengthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthExpressionContext) CopyAll(ctx *LengthExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LengthExpressionASTContext struct {
	LengthExpressionContext
}

func NewLengthExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthExpressionASTContext {
	var p = new(LengthExpressionASTContext)

	InitEmptyLengthExpressionContext(&p.LengthExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*LengthExpressionContext))

	return p
}

func (s *LengthExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionASTContext) LEN() antlr.TerminalNode {
	return s.GetToken(goParserLEN, 0)
}

func (s *LengthExpressionASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *LengthExpressionASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LengthExpressionASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *LengthExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLengthExpressionAST(s)
	}
}

func (s *LengthExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLengthExpressionAST(s)
	}
}

func (s *LengthExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLengthExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) LengthExpression() (localctx ILengthExpressionContext) {
	localctx = NewLengthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, goParserRULE_lengthExpression)
	localctx = NewLengthExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(goParserLEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(goParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.expression(0)
	}
	{
		p.SetState(313)
		p.Match(goParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICapExpressionContext is an interface to support dynamic dispatch.
type ICapExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCapExpressionContext differentiates from other interfaces.
	IsCapExpressionContext()
}

type CapExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapExpressionContext() *CapExpressionContext {
	var p = new(CapExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_capExpression
	return p
}

func InitEmptyCapExpressionContext(p *CapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_capExpression
}

func (*CapExpressionContext) IsCapExpressionContext() {}

func NewCapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapExpressionContext {
	var p = new(CapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_capExpression

	return p
}

func (s *CapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CapExpressionContext) CopyAll(ctx *CapExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CapExpressionASTContext struct {
	CapExpressionContext
}

func NewCapExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapExpressionASTContext {
	var p = new(CapExpressionASTContext)

	InitEmptyCapExpressionContext(&p.CapExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CapExpressionContext))

	return p
}

func (s *CapExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionASTContext) CAP() antlr.TerminalNode {
	return s.GetToken(goParserCAP, 0)
}

func (s *CapExpressionASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *CapExpressionASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CapExpressionASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *CapExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterCapExpressionAST(s)
	}
}

func (s *CapExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitCapExpressionAST(s)
	}
}

func (s *CapExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitCapExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) CapExpression() (localctx ICapExpressionContext) {
	localctx = NewCapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, goParserRULE_capExpression)
	localctx = NewCapExpressionASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(goParserCAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Match(goParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.expression(0)
	}
	{
		p.SetState(318)
		p.Match(goParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) CopyAll(ctx *StatementListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementListASTContext struct {
	StatementListContext
}

func NewStatementListASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementListASTContext {
	var p = new(StatementListASTContext)

	InitEmptyStatementListContext(&p.StatementListContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementListContext))

	return p
}

func (s *StatementListASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListASTContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListASTContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterStatementListAST(s)
	}
}

func (s *StatementListASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitStatementListAST(s)
	}
}

func (s *StatementListASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitStatementListAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, goParserRULE_statementList)
	var _la int

	localctx = NewStatementListASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3459890422334112138) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&259071) != 0) {
		{
			p.SetState(320)
			p.Statement()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) CopyAll(ctx *BlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockASTContext struct {
	BlockContext
}

func NewBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockASTContext {
	var p = new(BlockASTContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *BlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockASTContext) LCB() antlr.TerminalNode {
	return s.GetToken(goParserLCB, 0)
}

func (s *BlockASTContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockASTContext) RCB() antlr.TerminalNode {
	return s.GetToken(goParserRCB, 0)
}

func (s *BlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterBlockAST(s)
	}
}

func (s *BlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitBlockAST(s)
	}
}

func (s *BlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, goParserRULE_block)
	localctx = NewBlockASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(goParserLCB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.StatementList()
	}
	{
		p.SetState(328)
		p.Match(goParserRCB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakSASTContext struct {
	StatementContext
}

func NewBreakSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakSASTContext {
	var p = new(BreakSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakSASTContext) BREAK() antlr.TerminalNode {
	return s.GetToken(goParserBREAK, 0)
}

func (s *BreakSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *BreakSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterBreakSAST(s)
	}
}

func (s *BreakSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitBreakSAST(s)
	}
}

func (s *BreakSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitBreakSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockSASTContext struct {
	StatementContext
}

func NewBlockSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockSASTContext {
	var p = new(BlockSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BlockSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockSASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *BlockSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterBlockSAST(s)
	}
}

func (s *BlockSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitBlockSAST(s)
	}
}

func (s *BlockSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitBlockSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclSASTContext struct {
	StatementContext
}

func NewVariableDeclSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclSASTContext {
	var p = new(VariableDeclSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *VariableDeclSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclSASTContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *VariableDeclSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterVariableDeclSAST(s)
	}
}

func (s *VariableDeclSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitVariableDeclSAST(s)
	}
}

func (s *VariableDeclSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitVariableDeclSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueSASTContext struct {
	StatementContext
}

func NewContinueSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueSASTContext {
	var p = new(ContinueSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueSASTContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(goParserCONTINUE, 0)
}

func (s *ContinueSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *ContinueSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterContinueSAST(s)
	}
}

func (s *ContinueSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitContinueSAST(s)
	}
}

func (s *ContinueSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitContinueSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclSASTContext struct {
	StatementContext
}

func NewTypeDeclSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclSASTContext {
	var p = new(TypeDeclSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *TypeDeclSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclSASTContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TypeDeclSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterTypeDeclSAST(s)
	}
}

func (s *TypeDeclSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitTypeDeclSAST(s)
	}
}

func (s *TypeDeclSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitTypeDeclSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlnSASTContext struct {
	StatementContext
}

func NewPrintlnSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnSASTContext {
	var p = new(PrintlnSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintlnSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnSASTContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(goParserPRINTLN, 0)
}

func (s *PrintlnSASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *PrintlnSASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *PrintlnSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *PrintlnSASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PrintlnSASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *PrintlnSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterPrintlnSAST(s)
	}
}

func (s *PrintlnSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitPrintlnSAST(s)
	}
}

func (s *PrintlnSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitPrintlnSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchSASTContext struct {
	StatementContext
}

func NewSwitchSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchSASTContext {
	var p = new(SwitchSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SwitchSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchSASTContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *SwitchSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *SwitchSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSwitchSAST(s)
	}
}

func (s *SwitchSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSwitchSAST(s)
	}
}

func (s *SwitchSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSwitchSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementSASTContext struct {
	StatementContext
}

func NewIfStatementSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementSASTContext {
	var p = new(IfStatementSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStatementSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementSASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *IfStatementSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIfStatementSAST(s)
	}
}

func (s *IfStatementSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIfStatementSAST(s)
	}
}

func (s *IfStatementSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIfStatementSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnSASTContext struct {
	StatementContext
}

func NewReturnSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnSASTContext {
	var p = new(ReturnSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnSASTContext) RETURN() antlr.TerminalNode {
	return s.GetToken(goParserRETURN, 0)
}

func (s *ReturnSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *ReturnSASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnSASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *ReturnSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterReturnSAST(s)
	}
}

func (s *ReturnSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitReturnSAST(s)
	}
}

func (s *ReturnSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitReturnSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStatementSASTContext struct {
	StatementContext
}

func NewSimpleStatementSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStatementSASTContext {
	var p = new(SimpleStatementSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SimpleStatementSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementSASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SimpleStatementSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *SimpleStatementSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSimpleStatementSAST(s)
	}
}

func (s *SimpleStatementSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSimpleStatementSAST(s)
	}
}

func (s *SimpleStatementSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSimpleStatementSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopSASTContext struct {
	StatementContext
}

func NewLoopSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopSASTContext {
	var p = new(LoopSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *LoopSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopSASTContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *LoopSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *LoopSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLoopSAST(s)
	}
}

func (s *LoopSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLoopSAST(s)
	}
}

func (s *LoopSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLoopSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintSASTContext struct {
	StatementContext
}

func NewPrintSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintSASTContext {
	var p = new(PrintSASTContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintSASTContext) PRINT() antlr.TerminalNode {
	return s.GetToken(goParserPRINT, 0)
}

func (s *PrintSASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *PrintSASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *PrintSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *PrintSASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PrintSASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *PrintSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterPrintSAST(s)
	}
}

func (s *PrintSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitPrintSAST(s)
	}
}

func (s *PrintSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitPrintSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, goParserRULE_statement)
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserPRINT:
		localctx = NewPrintSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(goParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case goParserLP, goParserPL, goParserMIN, goParserCARET, goParserNEG, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
			{
				p.SetState(332)
				p.ExpressionList()
			}

		case goParserRP:
			{
				p.SetState(333)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(336)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserPRINTLN:
		localctx = NewPrintlnSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.Match(goParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case goParserLP, goParserPL, goParserMIN, goParserCARET, goParserNEG, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
			{
				p.SetState(341)
				p.ExpressionList()
			}

		case goParserRP:
			{
				p.SetState(342)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(345)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserRETURN:
		localctx = NewReturnSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(348)
			p.Match(goParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case goParserLP, goParserPL, goParserMIN, goParserCARET, goParserNEG, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
			{
				p.SetState(349)
				p.expression(0)
			}

		case goParserSEMI:
			{
				p.SetState(350)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(353)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserBREAK:
		localctx = NewBreakSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(355)
			p.Match(goParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserCONTINUE:
		localctx = NewContinueSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(357)
			p.Match(goParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserSEMI, goParserLP, goParserPL, goParserMIN, goParserCARET, goParserNEG, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
		localctx = NewSimpleStatementSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(359)
			p.SimpleStatement()
		}
		{
			p.SetState(360)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserLCB:
		localctx = NewBlockSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(362)
			p.Block()
		}
		{
			p.SetState(363)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserSWITCH:
		localctx = NewSwitchSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(365)
			p.Switch_()
		}
		{
			p.SetState(366)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserIF:
		localctx = NewIfStatementSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(368)
			p.IfStatement()
		}
		{
			p.SetState(369)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserFOR:
		localctx = NewLoopSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(371)
			p.Loop()
		}
		{
			p.SetState(372)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case goParserTYPE:
		localctx = NewTypeDeclSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(374)
			p.TypeDecl()
		}

	case goParserVAR:
		localctx = NewVariableDeclSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(375)
			p.VariableDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) CopyAll(ctx *SimpleStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssigmentStatementSSASTContext struct {
	SimpleStatementContext
}

func NewAssigmentStatementSSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssigmentStatementSSASTContext {
	var p = new(AssigmentStatementSSASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *AssigmentStatementSSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssigmentStatementSSASTContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *AssigmentStatementSSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterAssigmentStatementSSAST(s)
	}
}

func (s *AssigmentStatementSSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitAssigmentStatementSSAST(s)
	}
}

func (s *AssigmentStatementSSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitAssigmentStatementSSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionINCSSASTContext struct {
	SimpleStatementContext
}

func NewExpressionINCSSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionINCSSASTContext {
	var p = new(ExpressionINCSSASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpressionINCSSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionINCSSASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionINCSSASTContext) INC() antlr.TerminalNode {
	return s.GetToken(goParserINC, 0)
}

func (s *ExpressionINCSSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionINCSSAST(s)
	}
}

func (s *ExpressionINCSSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionINCSSAST(s)
	}
}

func (s *ExpressionINCSSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionINCSSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpListSSASTContext struct {
	SimpleStatementContext
}

func NewExpListSSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpListSSASTContext {
	var p = new(ExpListSSASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpListSSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpListSSASTContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *ExpListSSASTContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpListSSASTContext) SVD() antlr.TerminalNode {
	return s.GetToken(goParserSVD, 0)
}

func (s *ExpListSSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpListSSAST(s)
	}
}

func (s *ExpListSSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpListSSAST(s)
	}
}

func (s *ExpListSSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpListSSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type EpsilonSSASTContext struct {
	SimpleStatementContext
}

func NewEpsilonSSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpsilonSSASTContext {
	var p = new(EpsilonSSASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *EpsilonSSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsilonSSASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *EpsilonSSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterEpsilonSSAST(s)
	}
}

func (s *EpsilonSSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitEpsilonSSAST(s)
	}
}

func (s *EpsilonSSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitEpsilonSSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionEpsilonSSASTContext struct {
	SimpleStatementContext
}

func NewExpressionEpsilonSSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionEpsilonSSASTContext {
	var p = new(ExpressionEpsilonSSASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpressionEpsilonSSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionEpsilonSSASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionEpsilonSSASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *ExpressionEpsilonSSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionEpsilonSSAST(s)
	}
}

func (s *ExpressionEpsilonSSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionEpsilonSSAST(s)
	}
}

func (s *ExpressionEpsilonSSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionEpsilonSSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionDECSSASTContext struct {
	SimpleStatementContext
}

func NewExpressionDECSSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionDECSSASTContext {
	var p = new(ExpressionDECSSASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpressionDECSSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionDECSSASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionDECSSASTContext) DEC() antlr.TerminalNode {
	return s.GetToken(goParserDEC, 0)
}

func (s *ExpressionDECSSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionDECSSAST(s)
	}
}

func (s *ExpressionDECSSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionDECSSAST(s)
	}
}

func (s *ExpressionDECSSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionDECSSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, goParserRULE_simpleStatement)
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEpsilonSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.Epsilon()
		}

	case 2:
		localctx = NewExpressionINCSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.expression(0)
		}
		{
			p.SetState(380)
			p.Match(goParserINC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExpressionDECSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(382)
			p.expression(0)
		}
		{
			p.SetState(383)
			p.Match(goParserDEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpressionEpsilonSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(385)
			p.expression(0)
		}
		{
			p.SetState(386)
			p.Epsilon()
		}

	case 5:
		localctx = NewAssigmentStatementSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(388)
			p.AssignmentStatement()
		}

	case 6:
		localctx = NewExpListSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(389)
			p.ExpressionList()
		}
		{
			p.SetState(390)
			p.Match(goParserSVD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.ExpressionList()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) CopyAll(ctx *AssignmentStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpListASASTContext struct {
	AssignmentStatementContext
}

func NewExpListASASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpListASASTContext {
	var p = new(ExpListASASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *ExpListASASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpListASASTContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *ExpListASASTContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpListASASTContext) ASOP() antlr.TerminalNode {
	return s.GetToken(goParserASOP, 0)
}

func (s *ExpListASASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpListASAST(s)
	}
}

func (s *ExpListASASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpListASAST(s)
	}
}

func (s *ExpListASASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpListASAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpASASTContext struct {
	AssignmentStatementContext
}

func NewExpASASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpASASTContext {
	var p = new(ExpASASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *ExpASASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpASASTContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpASASTContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpASASTContext) AAOP() antlr.TerminalNode {
	return s.GetToken(goParserAAOP, 0)
}

func (s *ExpASASTContext) BAOP() antlr.TerminalNode {
	return s.GetToken(goParserBAOP, 0)
}

func (s *ExpASASTContext) SAOP() antlr.TerminalNode {
	return s.GetToken(goParserSAOP, 0)
}

func (s *ExpASASTContext) BOAOP() antlr.TerminalNode {
	return s.GetToken(goParserBOAOP, 0)
}

func (s *ExpASASTContext) MAOP() antlr.TerminalNode {
	return s.GetToken(goParserMAOP, 0)
}

func (s *ExpASASTContext) BXOOP() antlr.TerminalNode {
	return s.GetToken(goParserBXOOP, 0)
}

func (s *ExpASASTContext) LSAOP() antlr.TerminalNode {
	return s.GetToken(goParserLSAOP, 0)
}

func (s *ExpASASTContext) RSAOP() antlr.TerminalNode {
	return s.GetToken(goParserRSAOP, 0)
}

func (s *ExpASASTContext) BCAOP() antlr.TerminalNode {
	return s.GetToken(goParserBCAOP, 0)
}

func (s *ExpASASTContext) RAOP() antlr.TerminalNode {
	return s.GetToken(goParserRAOP, 0)
}

func (s *ExpASASTContext) DAOP() antlr.TerminalNode {
	return s.GetToken(goParserDAOP, 0)
}

func (s *ExpASASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpASAST(s)
	}
}

func (s *ExpASASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpASAST(s)
	}
}

func (s *ExpASASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpASAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, goParserRULE_assignmentStatement)
	var _la int

	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpListASASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.ExpressionList()
		}
		{
			p.SetState(396)
			p.Match(goParserASOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.ExpressionList()
		}

	case 2:
		localctx = NewExpASASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.expression(0)
		}
		{
			p.SetState(400)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70334384439296) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(401)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) CopyAll(ctx *IfStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IsExpBlockIsISASTContext struct {
	IfStatementContext
}

func NewIsExpBlockIsISASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsExpBlockIsISASTContext {
	var p = new(IsExpBlockIsISASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsExpBlockIsISASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsExpBlockIsISASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsExpBlockIsISASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsExpBlockIsISASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IsExpBlockIsISASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsExpBlockIsISASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IsExpBlockIsISASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsExpBlockIsISAST(s)
	}
}

func (s *IsExpBlockIsISASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsExpBlockIsISAST(s)
	}
}

func (s *IsExpBlockIsISASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsExpBlockIsISAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsSSExpBlockISASTContext struct {
	IfStatementContext
}

func NewIsSSExpBlockISASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsSSExpBlockISASTContext {
	var p = new(IsSSExpBlockISASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsSSExpBlockISASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsSSExpBlockISASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsSSExpBlockISASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IsSSExpBlockISASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *IsSSExpBlockISASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsSSExpBlockISASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IsSSExpBlockISASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsSSExpBlockISAST(s)
	}
}

func (s *IsSSExpBlockISASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsSSExpBlockISAST(s)
	}
}

func (s *IsSSExpBlockISASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsSSExpBlockISAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsExpBlockISASTContext struct {
	IfStatementContext
}

func NewIsExpBlockISASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsExpBlockISASTContext {
	var p = new(IsExpBlockISASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsExpBlockISASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsExpBlockISASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsExpBlockISASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsExpBlockISASTContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IsExpBlockISASTContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IsExpBlockISASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsExpBlockISASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsExpBlockISAST(s)
	}
}

func (s *IsExpBlockISASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsExpBlockISAST(s)
	}
}

func (s *IsExpBlockISASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsExpBlockISAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsSSExpBlockBlockASTContext struct {
	IfStatementContext
}

func NewIsSSExpBlockBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsSSExpBlockBlockASTContext {
	var p = new(IsSSExpBlockBlockASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsSSExpBlockBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsSSExpBlockBlockASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsSSExpBlockBlockASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IsSSExpBlockBlockASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *IsSSExpBlockBlockASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsSSExpBlockBlockASTContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IsSSExpBlockBlockASTContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IsSSExpBlockBlockASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsSSExpBlockBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsSSExpBlockBlockAST(s)
	}
}

func (s *IsSSExpBlockBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsSSExpBlockBlockAST(s)
	}
}

func (s *IsSSExpBlockBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsSSExpBlockBlockAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsExpressionBlockISASTContext struct {
	IfStatementContext
}

func NewIsExpressionBlockISASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsExpressionBlockISASTContext {
	var p = new(IsExpressionBlockISASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsExpressionBlockISASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsExpressionBlockISASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsExpressionBlockISASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsExpressionBlockISASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IsExpressionBlockISASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsExpressionBlockISAST(s)
	}
}

func (s *IsExpressionBlockISASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsExpressionBlockISAST(s)
	}
}

func (s *IsExpressionBlockISASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsExpressionBlockISAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsSSExpBlockifSASTContext struct {
	IfStatementContext
}

func NewIsSSExpBlockifSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsSSExpBlockifSASTContext {
	var p = new(IsSSExpBlockifSASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsSSExpBlockifSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsSSExpBlockifSASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsSSExpBlockifSASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IsSSExpBlockifSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *IsSSExpBlockifSASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsSSExpBlockifSASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IsSSExpBlockifSASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsSSExpBlockifSASTContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IsSSExpBlockifSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsSSExpBlockifSAST(s)
	}
}

func (s *IsSSExpBlockifSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsSSExpBlockifSAST(s)
	}
}

func (s *IsSSExpBlockifSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsSSExpBlockifSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, goParserRULE_ifStatement)
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIsExpressionBlockISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.expression(0)
		}
		{
			p.SetState(407)
			p.Block()
		}

	case 2:
		localctx = NewIsExpBlockIsISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(409)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.expression(0)
		}
		{
			p.SetState(411)
			p.Block()
		}
		{
			p.SetState(412)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.IfStatement()
		}

	case 3:
		localctx = NewIsExpBlockISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(415)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.expression(0)
		}
		{
			p.SetState(417)
			p.Block()
		}
		{
			p.SetState(418)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Block()
		}

	case 4:
		localctx = NewIsSSExpBlockISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(421)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.SimpleStatement()
		}
		{
			p.SetState(423)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.expression(0)
		}
		{
			p.SetState(425)
			p.Block()
		}

	case 5:
		localctx = NewIsSSExpBlockifSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(427)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.SimpleStatement()
		}
		{
			p.SetState(429)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.expression(0)
		}
		{
			p.SetState(431)
			p.Block()
		}
		{
			p.SetState(432)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.IfStatement()
		}

	case 6:
		localctx = NewIsSSExpBlockBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(435)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.SimpleStatement()
		}
		{
			p.SetState(437)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.expression(0)
		}
		{
			p.SetState(439)
			p.Block()
		}
		{
			p.SetState(440)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) CopyAll(ctx *LoopContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FExpBlockLASTContext struct {
	LoopContext
}

func NewFExpBlockLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FExpBlockLASTContext {
	var p = new(FExpBlockLASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *FExpBlockLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FExpBlockLASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(goParserFOR, 0)
}

func (s *FExpBlockLASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FExpBlockLASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FExpBlockLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFExpBlockLAST(s)
	}
}

func (s *FExpBlockLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFExpBlockLAST(s)
	}
}

func (s *FExpBlockLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFExpBlockLAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type FSimpStateSimpStateBlockLASTContext struct {
	LoopContext
}

func NewFSimpStateSimpStateBlockLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FSimpStateSimpStateBlockLASTContext {
	var p = new(FSimpStateSimpStateBlockLASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *FSimpStateSimpStateBlockLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FSimpStateSimpStateBlockLASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(goParserFOR, 0)
}

func (s *FSimpStateSimpStateBlockLASTContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *FSimpStateSimpStateBlockLASTContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *FSimpStateSimpStateBlockLASTContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(goParserSEMI)
}

func (s *FSimpStateSimpStateBlockLASTContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(goParserSEMI, i)
}

func (s *FSimpStateSimpStateBlockLASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FSimpStateSimpStateBlockLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFSimpStateSimpStateBlockLAST(s)
	}
}

func (s *FSimpStateSimpStateBlockLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFSimpStateSimpStateBlockLAST(s)
	}
}

func (s *FSimpStateSimpStateBlockLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFSimpStateSimpStateBlockLAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type FBlockLASTContext struct {
	LoopContext
}

func NewFBlockLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FBlockLASTContext {
	var p = new(FBlockLASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *FBlockLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FBlockLASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(goParserFOR, 0)
}

func (s *FBlockLASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FBlockLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFBlockLAST(s)
	}
}

func (s *FBlockLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFBlockLAST(s)
	}
}

func (s *FBlockLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFBlockLAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type FSimpStateExpSBlockLASTContext struct {
	LoopContext
}

func NewFSimpStateExpSBlockLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FSimpStateExpSBlockLASTContext {
	var p = new(FSimpStateExpSBlockLASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *FSimpStateExpSBlockLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FSimpStateExpSBlockLASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(goParserFOR, 0)
}

func (s *FSimpStateExpSBlockLASTContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *FSimpStateExpSBlockLASTContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *FSimpStateExpSBlockLASTContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(goParserSEMI)
}

func (s *FSimpStateExpSBlockLASTContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(goParserSEMI, i)
}

func (s *FSimpStateExpSBlockLASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FSimpStateExpSBlockLASTContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FSimpStateExpSBlockLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFSimpStateExpSBlockLAST(s)
	}
}

func (s *FSimpStateExpSBlockLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFSimpStateExpSBlockLAST(s)
	}
}

func (s *FSimpStateExpSBlockLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFSimpStateExpSBlockLAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, goParserRULE_loop)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFBlockLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Block()
		}

	case 2:
		localctx = NewFExpBlockLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(447)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.expression(0)
		}
		{
			p.SetState(449)
			p.Block()
		}

	case 3:
		localctx = NewFSimpStateExpSBlockLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(451)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.SimpleStatement()
		}
		{
			p.SetState(453)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.expression(0)
		}
		{
			p.SetState(455)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.SimpleStatement()
		}
		{
			p.SetState(457)
			p.Block()
		}

	case 4:
		localctx = NewFSimpStateSimpStateBlockLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(459)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.SimpleStatement()
		}
		{
			p.SetState(461)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.SimpleStatement()
		}
		{
			p.SetState(464)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) CopyAll(ctx *SwitchContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SSimpleStatExpCCListSASTContext struct {
	SwitchContext
}

func NewSSimpleStatExpCCListSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SSimpleStatExpCCListSASTContext {
	var p = new(SSimpleStatExpCCListSASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SSimpleStatExpCCListSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SSimpleStatExpCCListSASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(goParserSWITCH, 0)
}

func (s *SSimpleStatExpCCListSASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SSimpleStatExpCCListSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *SSimpleStatExpCCListSASTContext) LCB() antlr.TerminalNode {
	return s.GetToken(goParserLCB, 0)
}

func (s *SSimpleStatExpCCListSASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SSimpleStatExpCCListSASTContext) RCB() antlr.TerminalNode {
	return s.GetToken(goParserRCB, 0)
}

func (s *SSimpleStatExpCCListSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSSimpleStatExpCCListSAST(s)
	}
}

func (s *SSimpleStatExpCCListSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSSimpleStatExpCCListSAST(s)
	}
}

func (s *SSimpleStatExpCCListSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSSimpleStatExpCCListSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SExpressionSASTContext struct {
	SwitchContext
}

func NewSExpressionSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SExpressionSASTContext {
	var p = new(SExpressionSASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SExpressionSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SExpressionSASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(goParserSWITCH, 0)
}

func (s *SExpressionSASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SExpressionSASTContext) LCB() antlr.TerminalNode {
	return s.GetToken(goParserLCB, 0)
}

func (s *SExpressionSASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SExpressionSASTContext) RCB() antlr.TerminalNode {
	return s.GetToken(goParserRCB, 0)
}

func (s *SExpressionSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSExpressionSAST(s)
	}
}

func (s *SExpressionSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSExpressionSAST(s)
	}
}

func (s *SExpressionSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSExpressionSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SBlockSASTContext struct {
	SwitchContext
}

func NewSBlockSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SBlockSASTContext {
	var p = new(SBlockSASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SBlockSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SBlockSASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(goParserSWITCH, 0)
}

func (s *SBlockSASTContext) LCB() antlr.TerminalNode {
	return s.GetToken(goParserLCB, 0)
}

func (s *SBlockSASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SBlockSASTContext) RCB() antlr.TerminalNode {
	return s.GetToken(goParserRCB, 0)
}

func (s *SBlockSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSBlockSAST(s)
	}
}

func (s *SBlockSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSBlockSAST(s)
	}
}

func (s *SBlockSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSBlockSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type SSimpleStatSASTContext struct {
	SwitchContext
}

func NewSSimpleStatSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SSimpleStatSASTContext {
	var p = new(SSimpleStatSASTContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SSimpleStatSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SSimpleStatSASTContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(goParserSWITCH, 0)
}

func (s *SSimpleStatSASTContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SSimpleStatSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *SSimpleStatSASTContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SSimpleStatSASTContext) LCB() antlr.TerminalNode {
	return s.GetToken(goParserLCB, 0)
}

func (s *SSimpleStatSASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SSimpleStatSASTContext) RCB() antlr.TerminalNode {
	return s.GetToken(goParserRCB, 0)
}

func (s *SSimpleStatSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterSSimpleStatSAST(s)
	}
}

func (s *SSimpleStatSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitSSimpleStatSAST(s)
	}
}

func (s *SSimpleStatSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitSSimpleStatSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, goParserRULE_switch)
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSSimpleStatSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(468)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.SimpleStatement()
		}
		{
			p.SetState(470)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.expression(0)
		}
		{
			p.SetState(472)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(474)
			p.Match(goParserRCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSExpressionSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)
			p.expression(0)
		}
		{
			p.SetState(478)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(480)
			p.Match(goParserRCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSSimpleStatExpCCListSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(482)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.SimpleStatement()
		}
		{
			p.SetState(484)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(486)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(487)
			p.Match(goParserRCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSBlockSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(489)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(492)
			p.Match(goParserRCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseListContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionCaseClauseListContext differentiates from other interfaces.
	IsExpressionCaseClauseListContext()
}

type ExpressionCaseClauseListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseListContext() *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionCaseClauseList
	return p
}

func InitEmptyExpressionCaseClauseListContext(p *ExpressionCaseClauseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionCaseClauseList
}

func (*ExpressionCaseClauseListContext) IsExpressionCaseClauseListContext() {}

func NewExpressionCaseClauseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_expressionCaseClauseList

	return p
}

func (s *ExpressionCaseClauseListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseListContext) CopyAll(ctx *ExpressionCaseClauseListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionCaseClauseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpCaseClauseECCLASTContext struct {
	ExpressionCaseClauseListContext
}

func NewExpCaseClauseECCLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpCaseClauseECCLASTContext {
	var p = new(ExpCaseClauseECCLASTContext)

	InitEmptyExpressionCaseClauseListContext(&p.ExpressionCaseClauseListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseListContext))

	return p
}

func (s *ExpCaseClauseECCLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpCaseClauseECCLASTContext) ExpressionCaseClause() IExpressionCaseClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseContext)
}

func (s *ExpCaseClauseECCLASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *ExpCaseClauseECCLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpCaseClauseECCLAST(s)
	}
}

func (s *ExpCaseClauseECCLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpCaseClauseECCLAST(s)
	}
}

func (s *ExpCaseClauseECCLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpCaseClauseECCLAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type EpsilonECCLASTContext struct {
	ExpressionCaseClauseListContext
}

func NewEpsilonECCLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpsilonECCLASTContext {
	var p = new(EpsilonECCLASTContext)

	InitEmptyExpressionCaseClauseListContext(&p.ExpressionCaseClauseListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseListContext))

	return p
}

func (s *EpsilonECCLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsilonECCLASTContext) Epsilon() IEpsilonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpsilonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpsilonContext)
}

func (s *EpsilonECCLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterEpsilonECCLAST(s)
	}
}

func (s *EpsilonECCLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitEpsilonECCLAST(s)
	}
}

func (s *EpsilonECCLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitEpsilonECCLAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) ExpressionCaseClauseList() (localctx IExpressionCaseClauseListContext) {
	localctx = NewExpressionCaseClauseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, goParserRULE_expressionCaseClauseList)
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserRCB:
		localctx = NewEpsilonECCLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.Epsilon()
		}

	case goParserCASE, goParserDEFAULT:
		localctx = NewExpCaseClauseECCLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(497)
			p.ExpressionCaseClause()
		}
		{
			p.SetState(498)
			p.ExpressionCaseClauseList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionCaseClauseContext differentiates from other interfaces.
	IsExpressionCaseClauseContext()
}

type ExpressionCaseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseContext() *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionCaseClause
	return p
}

func InitEmptyExpressionCaseClauseContext(p *ExpressionCaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionCaseClause
}

func (*ExpressionCaseClauseContext) IsExpressionCaseClauseContext() {}

func NewExpressionCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_expressionCaseClause

	return p
}

func (s *ExpressionCaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseContext) CopyAll(ctx *ExpressionCaseClauseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionCaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpSwitchCaseECCASTContext struct {
	ExpressionCaseClauseContext
}

func NewExpSwitchCaseECCASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpSwitchCaseECCASTContext {
	var p = new(ExpSwitchCaseECCASTContext)

	InitEmptyExpressionCaseClauseContext(&p.ExpressionCaseClauseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseContext))

	return p
}

func (s *ExpSwitchCaseECCASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpSwitchCaseECCASTContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSwitchCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSwitchCaseContext)
}

func (s *ExpSwitchCaseECCASTContext) COL() antlr.TerminalNode {
	return s.GetToken(goParserCOL, 0)
}

func (s *ExpSwitchCaseECCASTContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ExpSwitchCaseECCASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpSwitchCaseECCAST(s)
	}
}

func (s *ExpSwitchCaseECCASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpSwitchCaseECCAST(s)
	}
}

func (s *ExpSwitchCaseECCASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpSwitchCaseECCAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, goParserRULE_expressionCaseClause)
	localctx = NewExpSwitchCaseECCASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(503)
		p.Match(goParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(504)
		p.StatementList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionSwitchCaseContext is an interface to support dynamic dispatch.
type IExpressionSwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionSwitchCaseContext differentiates from other interfaces.
	IsExpressionSwitchCaseContext()
}

type ExpressionSwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSwitchCaseContext() *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionSwitchCase
	return p
}

func InitEmptyExpressionSwitchCaseContext(p *ExpressionSwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_expressionSwitchCase
}

func (*ExpressionSwitchCaseContext) IsExpressionSwitchCaseContext() {}

func NewExpressionSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_expressionSwitchCase

	return p
}

func (s *ExpressionSwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSwitchCaseContext) CopyAll(ctx *ExpressionSwitchCaseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionSwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultESCASTContext struct {
	ExpressionSwitchCaseContext
}

func NewDefaultESCASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultESCASTContext {
	var p = new(DefaultESCASTContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *DefaultESCASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultESCASTContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(goParserDEFAULT, 0)
}

func (s *DefaultESCASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterDefaultESCAST(s)
	}
}

func (s *DefaultESCASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitDefaultESCAST(s)
	}
}

func (s *DefaultESCASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitDefaultESCAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type CaseESCASTContext struct {
	ExpressionSwitchCaseContext
}

func NewCaseESCASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseESCASTContext {
	var p = new(CaseESCASTContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *CaseESCASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseESCASTContext) CASE() antlr.TerminalNode {
	return s.GetToken(goParserCASE, 0)
}

func (s *CaseESCASTContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CaseESCASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterCaseESCAST(s)
	}
}

func (s *CaseESCASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitCaseESCAST(s)
	}
}

func (s *CaseESCASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitCaseESCAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, goParserRULE_expressionSwitchCase)
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserCASE:
		localctx = NewCaseESCASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(506)
			p.Match(goParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(507)
			p.ExpressionList()
		}

	case goParserDEFAULT:
		localctx = NewDefaultESCASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(508)
			p.Match(goParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEpsilonContext is an interface to support dynamic dispatch.
type IEpsilonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEpsilonContext differentiates from other interfaces.
	IsEpsilonContext()
}

type EpsilonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpsilonContext() *EpsilonContext {
	var p = new(EpsilonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_epsilon
	return p
}

func InitEmptyEpsilonContext(p *EpsilonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_epsilon
}

func (*EpsilonContext) IsEpsilonContext() {}

func NewEpsilonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpsilonContext {
	var p = new(EpsilonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_epsilon

	return p
}

func (s *EpsilonContext) GetParser() antlr.Parser { return s.parser }

func (s *EpsilonContext) CopyAll(ctx *EpsilonContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EpsilonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsilonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EpsilonASTContext struct {
	EpsilonContext
}

func NewEpsilonASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpsilonASTContext {
	var p = new(EpsilonASTContext)

	InitEmptyEpsilonContext(&p.EpsilonContext)
	p.parser = parser
	p.CopyAll(ctx.(*EpsilonContext))

	return p
}

func (s *EpsilonASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsilonASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterEpsilonAST(s)
	}
}

func (s *EpsilonASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitEpsilonAST(s)
	}
}

func (s *EpsilonASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitEpsilonAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Epsilon() (localctx IEpsilonContext) {
	localctx = NewEpsilonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, goParserRULE_epsilon)
	localctx = NewEpsilonASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *goParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 20:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *goParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *goParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
