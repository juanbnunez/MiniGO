// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/Compiler/goParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Generated // goParser
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
		"expressionCaseClause", "expressionSwitchCase", "operator", "epsilon",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 85, 513, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 93, 8, 1, 10, 1, 12,
		1, 96, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 117, 8,
		3, 10, 3, 12, 3, 120, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 132, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 149, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 156, 8, 7, 10, 7, 12, 7, 159, 9, 7, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 173, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 178, 8, 10, 1, 11, 1, 11,
		1, 11, 5, 11, 183, 8, 11, 10, 11, 12, 11, 186, 9, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 196, 8, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 3, 15, 211, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		5, 16, 220, 8, 16, 10, 16, 12, 16, 223, 9, 16, 1, 17, 1, 17, 1, 17, 5,
		17, 228, 8, 17, 10, 17, 12, 17, 231, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 3, 18, 238, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 244, 8, 18,
		10, 18, 12, 18, 247, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 252, 8, 19, 10,
		19, 12, 19, 255, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 262,
		8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 268, 8, 20, 5, 20, 270, 8, 20,
		10, 20, 12, 20, 273, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 281, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 288, 8, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 3, 24, 297, 8, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 5, 29, 322, 8, 29, 10, 29, 12, 29, 325, 9, 29, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 335, 8, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 344, 8, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 3, 31, 352, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 377,
		8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 384, 8, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 3, 32, 391, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 401, 8, 33, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 3, 34, 441, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 464, 8, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 3, 36, 492, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		3, 37, 498, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 3,
		39, 507, 8, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 0, 2, 36, 40, 42, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 0, 1, 2, 0, 7, 10, 18, 45, 541, 0, 84, 1, 0, 0, 0, 2, 94,
		1, 0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 111, 1, 0, 0, 0, 8, 131, 1, 0, 0, 0,
		10, 133, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 150, 1, 0, 0, 0, 16, 160,
		1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 167, 1, 0, 0, 0, 22, 179, 1, 0, 0,
		0, 24, 195, 1, 0, 0, 0, 26, 197, 1, 0, 0, 0, 28, 201, 1, 0, 0, 0, 30, 206,
		1, 0, 0, 0, 32, 214, 1, 0, 0, 0, 34, 224, 1, 0, 0, 0, 36, 237, 1, 0, 0,
		0, 38, 248, 1, 0, 0, 0, 40, 261, 1, 0, 0, 0, 42, 280, 1, 0, 0, 0, 44, 287,
		1, 0, 0, 0, 46, 289, 1, 0, 0, 0, 48, 293, 1, 0, 0, 0, 50, 300, 1, 0, 0,
		0, 52, 303, 1, 0, 0, 0, 54, 310, 1, 0, 0, 0, 56, 315, 1, 0, 0, 0, 58, 323,
		1, 0, 0, 0, 60, 326, 1, 0, 0, 0, 62, 376, 1, 0, 0, 0, 64, 390, 1, 0, 0,
		0, 66, 400, 1, 0, 0, 0, 68, 440, 1, 0, 0, 0, 70, 463, 1, 0, 0, 0, 72, 491,
		1, 0, 0, 0, 74, 497, 1, 0, 0, 0, 76, 499, 1, 0, 0, 0, 78, 506, 1, 0, 0,
		0, 80, 508, 1, 0, 0, 0, 82, 510, 1, 0, 0, 0, 84, 85, 5, 49, 0, 0, 85, 86,
		5, 76, 0, 0, 86, 87, 5, 1, 0, 0, 87, 88, 3, 2, 1, 0, 88, 1, 1, 0, 0, 0,
		89, 93, 3, 4, 2, 0, 90, 93, 3, 12, 6, 0, 91, 93, 3, 18, 9, 0, 92, 89, 1,
		0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94,
		92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 3, 1, 0, 0, 0, 96, 94, 1, 0, 0,
		0, 97, 109, 5, 60, 0, 0, 98, 99, 3, 8, 4, 0, 99, 100, 5, 1, 0, 0, 100,
		110, 1, 0, 0, 0, 101, 102, 5, 3, 0, 0, 102, 103, 3, 6, 3, 0, 103, 104,
		5, 4, 0, 0, 104, 105, 5, 1, 0, 0, 105, 110, 1, 0, 0, 0, 106, 107, 5, 3,
		0, 0, 107, 108, 5, 4, 0, 0, 108, 110, 5, 1, 0, 0, 109, 98, 1, 0, 0, 0,
		109, 101, 1, 0, 0, 0, 109, 106, 1, 0, 0, 0, 110, 5, 1, 0, 0, 0, 111, 112,
		3, 8, 4, 0, 112, 118, 5, 1, 0, 0, 113, 114, 3, 8, 4, 0, 114, 115, 5, 1,
		0, 0, 115, 117, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 7, 1, 0, 0, 0, 120, 118,
		1, 0, 0, 0, 121, 122, 3, 34, 17, 0, 122, 123, 3, 24, 12, 0, 123, 124, 5,
		34, 0, 0, 124, 125, 3, 38, 19, 0, 125, 132, 1, 0, 0, 0, 126, 127, 3, 34,
		17, 0, 127, 128, 5, 34, 0, 0, 128, 129, 3, 38, 19, 0, 129, 132, 1, 0, 0,
		0, 130, 132, 3, 10, 5, 0, 131, 121, 1, 0, 0, 0, 131, 126, 1, 0, 0, 0, 131,
		130, 1, 0, 0, 0, 132, 9, 1, 0, 0, 0, 133, 134, 3, 34, 17, 0, 134, 135,
		3, 24, 12, 0, 135, 11, 1, 0, 0, 0, 136, 148, 5, 61, 0, 0, 137, 138, 3,
		16, 8, 0, 138, 139, 5, 1, 0, 0, 139, 149, 1, 0, 0, 0, 140, 141, 5, 3, 0,
		0, 141, 142, 3, 14, 7, 0, 142, 143, 5, 4, 0, 0, 143, 144, 5, 1, 0, 0, 144,
		149, 1, 0, 0, 0, 145, 146, 5, 3, 0, 0, 146, 147, 5, 4, 0, 0, 147, 149,
		5, 1, 0, 0, 148, 137, 1, 0, 0, 0, 148, 140, 1, 0, 0, 0, 148, 145, 1, 0,
		0, 0, 149, 13, 1, 0, 0, 0, 150, 151, 3, 16, 8, 0, 151, 157, 5, 1, 0, 0,
		152, 153, 3, 16, 8, 0, 153, 154, 5, 1, 0, 0, 154, 156, 1, 0, 0, 0, 155,
		152, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 15, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 76,
		0, 0, 161, 162, 3, 24, 12, 0, 162, 17, 1, 0, 0, 0, 163, 164, 3, 20, 10,
		0, 164, 165, 3, 60, 30, 0, 165, 166, 5, 1, 0, 0, 166, 19, 1, 0, 0, 0, 167,
		168, 5, 62, 0, 0, 168, 169, 5, 76, 0, 0, 169, 172, 5, 3, 0, 0, 170, 173,
		3, 22, 11, 0, 171, 173, 3, 82, 41, 0, 172, 170, 1, 0, 0, 0, 172, 171, 1,
		0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 177, 5, 4, 0, 0, 175, 178, 3, 24, 12,
		0, 176, 178, 3, 82, 41, 0, 177, 175, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0,
		178, 21, 1, 0, 0, 0, 179, 184, 3, 10, 5, 0, 180, 181, 5, 11, 0, 0, 181,
		183, 3, 10, 5, 0, 182, 180, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 182,
		1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 23, 1, 0, 0, 0, 186, 184, 1, 0,
		0, 0, 187, 188, 5, 3, 0, 0, 188, 189, 3, 24, 12, 0, 189, 190, 5, 4, 0,
		0, 190, 196, 1, 0, 0, 0, 191, 196, 5, 76, 0, 0, 192, 196, 3, 26, 13, 0,
		193, 196, 3, 28, 14, 0, 194, 196, 3, 30, 15, 0, 195, 187, 1, 0, 0, 0, 195,
		191, 1, 0, 0, 0, 195, 192, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 194,
		1, 0, 0, 0, 196, 25, 1, 0, 0, 0, 197, 198, 5, 12, 0, 0, 198, 199, 5, 13,
		0, 0, 199, 200, 3, 24, 12, 0, 200, 27, 1, 0, 0, 0, 201, 202, 5, 12, 0,
		0, 202, 203, 5, 77, 0, 0, 203, 204, 5, 13, 0, 0, 204, 205, 3, 24, 12, 0,
		205, 29, 1, 0, 0, 0, 206, 207, 5, 63, 0, 0, 207, 210, 5, 14, 0, 0, 208,
		211, 3, 32, 16, 0, 209, 211, 3, 82, 41, 0, 210, 208, 1, 0, 0, 0, 210, 209,
		1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 5, 15, 0, 0, 213, 31, 1, 0,
		0, 0, 214, 215, 3, 10, 5, 0, 215, 221, 5, 1, 0, 0, 216, 217, 3, 10, 5,
		0, 217, 218, 5, 1, 0, 0, 218, 220, 1, 0, 0, 0, 219, 216, 1, 0, 0, 0, 220,
		223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 33, 1,
		0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 229, 5, 76, 0, 0, 225, 226, 5, 11,
		0, 0, 226, 228, 5, 76, 0, 0, 227, 225, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0,
		229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 35, 1, 0, 0, 0, 231, 229,
		1, 0, 0, 0, 232, 233, 6, 18, -1, 0, 233, 238, 3, 40, 20, 0, 234, 235, 3,
		80, 40, 0, 235, 236, 3, 36, 18, 1, 236, 238, 1, 0, 0, 0, 237, 232, 1, 0,
		0, 0, 237, 234, 1, 0, 0, 0, 238, 245, 1, 0, 0, 0, 239, 240, 10, 2, 0, 0,
		240, 241, 3, 80, 40, 0, 241, 242, 3, 36, 18, 3, 242, 244, 1, 0, 0, 0, 243,
		239, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246,
		1, 0, 0, 0, 246, 37, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 253, 3, 36,
		18, 0, 249, 250, 5, 11, 0, 0, 250, 252, 3, 36, 18, 0, 251, 249, 1, 0, 0,
		0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		39, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 257, 6, 20, -1, 0, 257, 262,
		3, 42, 21, 0, 258, 262, 3, 52, 26, 0, 259, 262, 3, 54, 27, 0, 260, 262,
		3, 56, 28, 0, 261, 256, 1, 0, 0, 0, 261, 258, 1, 0, 0, 0, 261, 259, 1,
		0, 0, 0, 261, 260, 1, 0, 0, 0, 262, 271, 1, 0, 0, 0, 263, 267, 10, 4, 0,
		0, 264, 268, 3, 50, 25, 0, 265, 268, 3, 46, 23, 0, 266, 268, 3, 48, 24,
		0, 267, 264, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 266, 1, 0, 0, 0, 268,
		270, 1, 0, 0, 0, 269, 263, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269,
		1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 41, 1, 0, 0, 0, 273, 271, 1, 0,
		0, 0, 274, 281, 3, 44, 22, 0, 275, 281, 5, 76, 0, 0, 276, 277, 5, 3, 0,
		0, 277, 278, 3, 36, 18, 0, 278, 279, 5, 4, 0, 0, 279, 281, 1, 0, 0, 0,
		280, 274, 1, 0, 0, 0, 280, 275, 1, 0, 0, 0, 280, 276, 1, 0, 0, 0, 281,
		43, 1, 0, 0, 0, 282, 288, 5, 77, 0, 0, 283, 288, 5, 78, 0, 0, 284, 288,
		5, 79, 0, 0, 285, 288, 5, 80, 0, 0, 286, 288, 5, 81, 0, 0, 287, 282, 1,
		0, 0, 0, 287, 283, 1, 0, 0, 0, 287, 284, 1, 0, 0, 0, 287, 285, 1, 0, 0,
		0, 287, 286, 1, 0, 0, 0, 288, 45, 1, 0, 0, 0, 289, 290, 5, 12, 0, 0, 290,
		291, 3, 36, 18, 0, 291, 292, 5, 13, 0, 0, 292, 47, 1, 0, 0, 0, 293, 296,
		5, 3, 0, 0, 294, 297, 3, 38, 19, 0, 295, 297, 3, 82, 41, 0, 296, 294, 1,
		0, 0, 0, 296, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299, 5, 4, 0,
		0, 299, 49, 1, 0, 0, 0, 300, 301, 5, 46, 0, 0, 301, 302, 5, 76, 0, 0, 302,
		51, 1, 0, 0, 0, 303, 304, 5, 64, 0, 0, 304, 305, 5, 3, 0, 0, 305, 306,
		3, 36, 18, 0, 306, 307, 5, 11, 0, 0, 307, 308, 3, 36, 18, 0, 308, 309,
		5, 4, 0, 0, 309, 53, 1, 0, 0, 0, 310, 311, 5, 65, 0, 0, 311, 312, 5, 3,
		0, 0, 312, 313, 3, 36, 18, 0, 313, 314, 5, 4, 0, 0, 314, 55, 1, 0, 0, 0,
		315, 316, 5, 66, 0, 0, 316, 317, 5, 3, 0, 0, 317, 318, 3, 36, 18, 0, 318,
		319, 5, 4, 0, 0, 319, 57, 1, 0, 0, 0, 320, 322, 3, 62, 31, 0, 321, 320,
		1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0,
		0, 0, 324, 59, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 327, 5, 14, 0, 0,
		327, 328, 3, 58, 29, 0, 328, 329, 5, 15, 0, 0, 329, 61, 1, 0, 0, 0, 330,
		331, 5, 67, 0, 0, 331, 334, 5, 3, 0, 0, 332, 335, 3, 38, 19, 0, 333, 335,
		3, 82, 41, 0, 334, 332, 1, 0, 0, 0, 334, 333, 1, 0, 0, 0, 335, 336, 1,
		0, 0, 0, 336, 337, 5, 4, 0, 0, 337, 338, 5, 1, 0, 0, 338, 377, 1, 0, 0,
		0, 339, 340, 5, 68, 0, 0, 340, 343, 5, 3, 0, 0, 341, 344, 3, 38, 19, 0,
		342, 344, 3, 82, 41, 0, 343, 341, 1, 0, 0, 0, 343, 342, 1, 0, 0, 0, 344,
		345, 1, 0, 0, 0, 345, 346, 5, 4, 0, 0, 346, 347, 5, 1, 0, 0, 347, 377,
		1, 0, 0, 0, 348, 351, 5, 69, 0, 0, 349, 352, 3, 36, 18, 0, 350, 352, 3,
		82, 41, 0, 351, 349, 1, 0, 0, 0, 351, 350, 1, 0, 0, 0, 352, 353, 1, 0,
		0, 0, 353, 354, 5, 1, 0, 0, 354, 377, 1, 0, 0, 0, 355, 356, 5, 70, 0, 0,
		356, 377, 5, 1, 0, 0, 357, 358, 5, 71, 0, 0, 358, 377, 5, 1, 0, 0, 359,
		360, 3, 64, 32, 0, 360, 361, 5, 1, 0, 0, 361, 377, 1, 0, 0, 0, 362, 363,
		3, 60, 30, 0, 363, 364, 5, 1, 0, 0, 364, 377, 1, 0, 0, 0, 365, 366, 3,
		72, 36, 0, 366, 367, 5, 1, 0, 0, 367, 377, 1, 0, 0, 0, 368, 369, 3, 68,
		34, 0, 369, 370, 5, 1, 0, 0, 370, 377, 1, 0, 0, 0, 371, 372, 3, 70, 35,
		0, 372, 373, 5, 1, 0, 0, 373, 377, 1, 0, 0, 0, 374, 377, 3, 12, 6, 0, 375,
		377, 3, 4, 2, 0, 376, 330, 1, 0, 0, 0, 376, 339, 1, 0, 0, 0, 376, 348,
		1, 0, 0, 0, 376, 355, 1, 0, 0, 0, 376, 357, 1, 0, 0, 0, 376, 359, 1, 0,
		0, 0, 376, 362, 1, 0, 0, 0, 376, 365, 1, 0, 0, 0, 376, 368, 1, 0, 0, 0,
		376, 371, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1, 0, 0, 0, 377,
		63, 1, 0, 0, 0, 378, 391, 3, 82, 41, 0, 379, 383, 3, 36, 18, 0, 380, 384,
		5, 16, 0, 0, 381, 384, 5, 17, 0, 0, 382, 384, 3, 82, 41, 0, 383, 380, 1,
		0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 382, 1, 0, 0, 0, 384, 391, 1, 0, 0,
		0, 385, 391, 3, 66, 33, 0, 386, 387, 3, 38, 19, 0, 387, 388, 5, 2, 0, 0,
		388, 389, 3, 38, 19, 0, 389, 391, 1, 0, 0, 0, 390, 378, 1, 0, 0, 0, 390,
		379, 1, 0, 0, 0, 390, 385, 1, 0, 0, 0, 390, 386, 1, 0, 0, 0, 391, 65, 1,
		0, 0, 0, 392, 393, 3, 38, 19, 0, 393, 394, 3, 80, 40, 0, 394, 395, 3, 38,
		19, 0, 395, 401, 1, 0, 0, 0, 396, 397, 3, 36, 18, 0, 397, 398, 3, 80, 40,
		0, 398, 399, 3, 36, 18, 0, 399, 401, 1, 0, 0, 0, 400, 392, 1, 0, 0, 0,
		400, 396, 1, 0, 0, 0, 401, 67, 1, 0, 0, 0, 402, 403, 5, 50, 0, 0, 403,
		404, 3, 36, 18, 0, 404, 405, 3, 60, 30, 0, 405, 441, 1, 0, 0, 0, 406, 407,
		5, 50, 0, 0, 407, 408, 3, 36, 18, 0, 408, 409, 3, 60, 30, 0, 409, 410,
		5, 54, 0, 0, 410, 411, 3, 68, 34, 0, 411, 441, 1, 0, 0, 0, 412, 413, 5,
		50, 0, 0, 413, 414, 3, 36, 18, 0, 414, 415, 3, 60, 30, 0, 415, 416, 5,
		54, 0, 0, 416, 417, 3, 60, 30, 0, 417, 441, 1, 0, 0, 0, 418, 419, 5, 50,
		0, 0, 419, 420, 3, 64, 32, 0, 420, 421, 5, 1, 0, 0, 421, 422, 3, 36, 18,
		0, 422, 423, 3, 60, 30, 0, 423, 441, 1, 0, 0, 0, 424, 425, 5, 50, 0, 0,
		425, 426, 3, 64, 32, 0, 426, 427, 5, 1, 0, 0, 427, 428, 3, 36, 18, 0, 428,
		429, 3, 60, 30, 0, 429, 430, 5, 54, 0, 0, 430, 431, 3, 68, 34, 0, 431,
		441, 1, 0, 0, 0, 432, 433, 5, 50, 0, 0, 433, 434, 3, 64, 32, 0, 434, 435,
		5, 1, 0, 0, 435, 436, 3, 36, 18, 0, 436, 437, 3, 60, 30, 0, 437, 438, 5,
		54, 0, 0, 438, 439, 3, 60, 30, 0, 439, 441, 1, 0, 0, 0, 440, 402, 1, 0,
		0, 0, 440, 406, 1, 0, 0, 0, 440, 412, 1, 0, 0, 0, 440, 418, 1, 0, 0, 0,
		440, 424, 1, 0, 0, 0, 440, 432, 1, 0, 0, 0, 441, 69, 1, 0, 0, 0, 442, 443,
		5, 72, 0, 0, 443, 464, 3, 60, 30, 0, 444, 445, 5, 72, 0, 0, 445, 446, 3,
		36, 18, 0, 446, 447, 3, 60, 30, 0, 447, 464, 1, 0, 0, 0, 448, 449, 5, 72,
		0, 0, 449, 450, 3, 64, 32, 0, 450, 451, 5, 1, 0, 0, 451, 452, 3, 36, 18,
		0, 452, 453, 5, 1, 0, 0, 453, 454, 3, 64, 32, 0, 454, 455, 3, 60, 30, 0,
		455, 464, 1, 0, 0, 0, 456, 457, 5, 72, 0, 0, 457, 458, 3, 64, 32, 0, 458,
		459, 5, 1, 0, 0, 459, 460, 5, 1, 0, 0, 460, 461, 3, 64, 32, 0, 461, 462,
		3, 60, 30, 0, 462, 464, 1, 0, 0, 0, 463, 442, 1, 0, 0, 0, 463, 444, 1,
		0, 0, 0, 463, 448, 1, 0, 0, 0, 463, 456, 1, 0, 0, 0, 464, 71, 1, 0, 0,
		0, 465, 466, 5, 73, 0, 0, 466, 467, 3, 64, 32, 0, 467, 468, 5, 1, 0, 0,
		468, 469, 3, 36, 18, 0, 469, 470, 5, 14, 0, 0, 470, 471, 3, 74, 37, 0,
		471, 472, 5, 15, 0, 0, 472, 492, 1, 0, 0, 0, 473, 474, 5, 73, 0, 0, 474,
		475, 3, 36, 18, 0, 475, 476, 5, 14, 0, 0, 476, 477, 3, 74, 37, 0, 477,
		478, 5, 15, 0, 0, 478, 492, 1, 0, 0, 0, 479, 480, 5, 73, 0, 0, 480, 481,
		3, 64, 32, 0, 481, 482, 5, 1, 0, 0, 482, 483, 5, 14, 0, 0, 483, 484, 3,
		74, 37, 0, 484, 485, 5, 15, 0, 0, 485, 492, 1, 0, 0, 0, 486, 487, 5, 73,
		0, 0, 487, 488, 5, 14, 0, 0, 488, 489, 3, 74, 37, 0, 489, 490, 5, 15, 0,
		0, 490, 492, 1, 0, 0, 0, 491, 465, 1, 0, 0, 0, 491, 473, 1, 0, 0, 0, 491,
		479, 1, 0, 0, 0, 491, 486, 1, 0, 0, 0, 492, 73, 1, 0, 0, 0, 493, 498, 3,
		82, 41, 0, 494, 495, 3, 76, 38, 0, 495, 496, 3, 74, 37, 0, 496, 498, 1,
		0, 0, 0, 497, 493, 1, 0, 0, 0, 497, 494, 1, 0, 0, 0, 498, 75, 1, 0, 0,
		0, 499, 500, 3, 78, 39, 0, 500, 501, 5, 6, 0, 0, 501, 502, 3, 58, 29, 0,
		502, 77, 1, 0, 0, 0, 503, 504, 5, 74, 0, 0, 504, 507, 3, 38, 19, 0, 505,
		507, 5, 75, 0, 0, 506, 503, 1, 0, 0, 0, 506, 505, 1, 0, 0, 0, 507, 79,
		1, 0, 0, 0, 508, 509, 7, 0, 0, 0, 509, 81, 1, 0, 0, 0, 510, 511, 1, 0,
		0, 0, 511, 83, 1, 0, 0, 0, 36, 92, 94, 109, 118, 131, 148, 157, 172, 177,
		184, 195, 210, 221, 229, 237, 245, 253, 261, 267, 271, 280, 287, 296, 323,
		334, 343, 351, 376, 383, 390, 400, 440, 463, 491, 497, 506,
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
	goParserRULE_operator                 = 40
	goParserRULE_epsilon                  = 41
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
		p.SetState(84)
		p.Match(goParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8070450532247928832) != 0 {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case goParserVAR:
			{
				p.SetState(89)
				p.VariableDecl()
			}

		case goParserTYPE:
			{
				p.SetState(90)
				p.TypeDecl()
			}

		case goParserFUNC:
			{
				p.SetState(91)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(96)
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

type VarVDASTContext struct {
	VariableDeclContext
}

func NewVarVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarVDASTContext {
	var p = new(VarVDASTContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VarVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarVDASTContext) VAR() antlr.TerminalNode {
	return s.GetToken(goParserVAR, 0)
}

func (s *VarVDASTContext) SingleVarDecl() ISingleVarDeclContext {
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

func (s *VarVDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *VarVDASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *VarVDASTContext) InnerVarDecls() IInnerVarDeclsContext {
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

func (s *VarVDASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *VarVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterVarVDAST(s)
	}
}

func (s *VarVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitVarVDAST(s)
	}
}

func (s *VarVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitVarVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, goParserRULE_variableDecl)
	localctx = NewVarVDASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(goParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(98)
			p.SingleVarDecl()
		}
		{
			p.SetState(99)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(101)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.InnerVarDecls()
		}
		{
			p.SetState(103)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
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

type IdentifierListDTSVDASTContext struct {
	SingleVarDeclContext
}

func NewIdentifierListDTSVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierListDTSVDASTContext {
	var p = new(IdentifierListDTSVDASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *IdentifierListDTSVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListDTSVDASTContext) IdentifierList() IIdentifierListContext {
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

func (s *IdentifierListDTSVDASTContext) DeclType() IDeclTypeContext {
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

func (s *IdentifierListDTSVDASTContext) ASOP() antlr.TerminalNode {
	return s.GetToken(goParserASOP, 0)
}

func (s *IdentifierListDTSVDASTContext) ExpressionList() IExpressionListContext {
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

func (s *IdentifierListDTSVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdentifierListDTSVDAST(s)
	}
}

func (s *IdentifierListDTSVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdentifierListDTSVDAST(s)
	}
}

func (s *IdentifierListDTSVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdentifierListDTSVDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierListSVDASTContext struct {
	SingleVarDeclContext
}

func NewIdentifierListSVDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierListSVDASTContext {
	var p = new(IdentifierListSVDASTContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *IdentifierListSVDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListSVDASTContext) IdentifierList() IIdentifierListContext {
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

func (s *IdentifierListSVDASTContext) ASOP() antlr.TerminalNode {
	return s.GetToken(goParserASOP, 0)
}

func (s *IdentifierListSVDASTContext) ExpressionList() IExpressionListContext {
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

func (s *IdentifierListSVDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIdentifierListSVDAST(s)
	}
}

func (s *IdentifierListSVDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIdentifierListSVDAST(s)
	}
}

func (s *IdentifierListSVDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIdentifierListSVDAST(s)

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
		localctx = NewIdentifierListDTSVDASTContext(p, localctx)
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
		localctx = NewIdentifierListSVDASTContext(p, localctx)
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

type TypeTDASTContext struct {
	TypeDeclContext
}

func NewTypeTDASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeTDASTContext {
	var p = new(TypeTDASTContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeTDASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTDASTContext) TYPE() antlr.TerminalNode {
	return s.GetToken(goParserTYPE, 0)
}

func (s *TypeTDASTContext) SingleTypeDecl() ISingleTypeDeclContext {
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

func (s *TypeTDASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *TypeTDASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *TypeTDASTContext) InnerTypeDecls() IInnerTypeDeclsContext {
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

func (s *TypeTDASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *TypeTDASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterTypeTDAST(s)
	}
}

func (s *TypeTDASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitTypeTDAST(s)
	}
}

func (s *TypeTDASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitTypeTDAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, goParserRULE_typeDecl)
	localctx = NewTypeTDASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(goParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
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
		{
			p.SetState(140)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.InnerTypeDecls()
		}
		{
			p.SetState(142)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(145)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(goParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
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

type InnerTypeDeclsAStContext struct {
	InnerTypeDeclsContext
}

func NewInnerTypeDeclsAStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerTypeDeclsAStContext {
	var p = new(InnerTypeDeclsAStContext)

	InitEmptyInnerTypeDeclsContext(&p.InnerTypeDeclsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InnerTypeDeclsContext))

	return p
}

func (s *InnerTypeDeclsAStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsAStContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
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

func (s *InnerTypeDeclsAStContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
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

func (s *InnerTypeDeclsAStContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(goParserSEMI)
}

func (s *InnerTypeDeclsAStContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(goParserSEMI, i)
}

func (s *InnerTypeDeclsAStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterInnerTypeDeclsASt(s)
	}
}

func (s *InnerTypeDeclsAStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitInnerTypeDeclsASt(s)
	}
}

func (s *InnerTypeDeclsAStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitInnerTypeDeclsASt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, goParserRULE_innerTypeDecls)
	var _la int

	localctx = NewInnerTypeDeclsAStContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.SingleTypeDecl()
	}
	{
		p.SetState(151)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserID {
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
		p.SetState(160)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
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
		p.SetState(163)
		p.FuncFrontDecl()
	}
	{
		p.SetState(164)
		p.Block()
	}
	{
		p.SetState(165)
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
		p.SetState(167)
		p.Match(goParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(goParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserID:
		{
			p.SetState(170)
			p.FuncArgDecls()
		}

	case goParserRP:
		{
			p.SetState(171)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(174)
		p.Match(goParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserLP, goParserLSB, goParserSTRUCT, goParserID:
		{
			p.SetState(175)
			p.DeclType()
		}

	case goParserLCB:
		{
			p.SetState(176)
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
		p.SetState(179)
		p.SingleVarDeclNoExps()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserCM {
		{
			p.SetState(180)
			p.Match(goParserCM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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

type DeclTypeASTContext struct {
	DeclTypeContext
}

func NewDeclTypeASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclTypeASTContext {
	var p = new(DeclTypeASTContext)

	InitEmptyDeclTypeContext(&p.DeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclTypeContext))

	return p
}

func (s *DeclTypeASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeASTContext) StructDeclType() IStructDeclTypeContext {
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

func (s *DeclTypeASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterDeclTypeAST(s)
	}
}

func (s *DeclTypeASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitDeclTypeAST(s)
	}
}

func (s *DeclTypeASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitDeclTypeAST(s)

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
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLpDTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(goParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.DeclType()
		}
		{
			p.SetState(189)
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
			p.SetState(191)
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
			p.SetState(192)
			p.SliceDeclType()
		}

	case 4:
		localctx = NewArrayDeclTypeDTASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(193)
			p.ArrayDeclType()
		}

	case 5:
		localctx = NewDeclTypeASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(194)
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
		p.SetState(197)
		p.Match(goParserLSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(goParserRSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
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
		p.SetState(201)
		p.Match(goParserLSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(goParserINTLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(goParserRSB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
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

type StrctSDTASTContext struct {
	StructDeclTypeContext
}

func NewStrctSDTASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrctSDTASTContext {
	var p = new(StrctSDTASTContext)

	InitEmptyStructDeclTypeContext(&p.StructDeclTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructDeclTypeContext))

	return p
}

func (s *StrctSDTASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrctSDTASTContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(goParserSTRUCT, 0)
}

func (s *StrctSDTASTContext) LCB() antlr.TerminalNode {
	return s.GetToken(goParserLCB, 0)
}

func (s *StrctSDTASTContext) RCB() antlr.TerminalNode {
	return s.GetToken(goParserRCB, 0)
}

func (s *StrctSDTASTContext) StructMemDecls() IStructMemDeclsContext {
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

func (s *StrctSDTASTContext) Epsilon() IEpsilonContext {
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

func (s *StrctSDTASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterStrctSDTAST(s)
	}
}

func (s *StrctSDTASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitStrctSDTAST(s)
	}
}

func (s *StrctSDTASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitStrctSDTAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, goParserRULE_structDeclType)
	localctx = NewStrctSDTASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(goParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Match(goParserLCB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserID:
		{
			p.SetState(208)
			p.StructMemDecls()
		}

	case goParserRCB:
		{
			p.SetState(209)
			p.Epsilon()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(212)
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
		p.SetState(214)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(215)
		p.Match(goParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserID {
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
		p.SetState(224)
		p.Match(goParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == goParserCM {
		{
			p.SetState(225)
			p.Match(goParserCM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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

type ExpressionASTContext struct {
	ExpressionContext
}

func NewExpressionASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionASTContext {
	var p = new(ExpressionASTContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionASTContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExpressionASTContext) Expression() IExpressionContext {
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

func (s *ExpressionASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionAST(s)
	}
}

func (s *ExpressionASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionAST(s)
	}
}

func (s *ExpressionASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionAST(s)

	default:
		return t.VisitChildren(s)
	}
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

func (s *ExpressionEASTContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(237)
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
			p.SetState(233)
			p.primaryExpression(0)
		}

	case goParserPL, goParserMIN, goParserMUL, goParserDIV, goParserMOD, goParserLSH, goParserRSH, goParserAMPER, goParserBC, goParserVB, goParserCARET, goParserEQ, goParserNEQ, goParserLT, goParserGT, goParserLTE, goParserGTE, goParserLAND, goParserLOR, goParserNEG, goParserASOP, goParserAAOP, goParserBAOP, goParserSAOP, goParserBOAOP, goParserMAOP, goParserBXOOP, goParserLSAOP, goParserRSAOP, goParserBCAOP, goParserRAOP, goParserDAOP:
		localctx = NewExpressionASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Operator()
		}
		{
			p.SetState(235)
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
			p.SetState(239)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(240)
				p.Operator()
			}
			{
				p.SetState(241)
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

type ExpressionListAStContext struct {
	ExpressionListContext
}

func NewExpressionListAStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListAStContext {
	var p = new(ExpressionListAStContext)

	InitEmptyExpressionListContext(&p.ExpressionListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionListContext))

	return p
}

func (s *ExpressionListAStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListAStContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionListAStContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionListAStContext) AllCM() []antlr.TerminalNode {
	return s.GetTokens(goParserCM)
}

func (s *ExpressionListAStContext) CM(i int) antlr.TerminalNode {
	return s.GetToken(goParserCM, i)
}

func (s *ExpressionListAStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionListASt(s)
	}
}

func (s *ExpressionListAStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionListASt(s)
	}
}

func (s *ExpressionListAStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionListASt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, goParserRULE_expressionList)
	var _la int

	localctx = NewExpressionListAStContext(p, localctx)
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

type CapExpressionPEASTContext struct {
	PrimaryExpressionContext
}

func NewCapExpressionPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapExpressionPEASTContext {
	var p = new(CapExpressionPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *CapExpressionPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionPEASTContext) CapExpression() ICapExpressionContext {
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

func (s *CapExpressionPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterCapExpressionPEAST(s)
	}
}

func (s *CapExpressionPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitCapExpressionPEAST(s)
	}
}

func (s *CapExpressionPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitCapExpressionPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExpressionPEASTContext struct {
	PrimaryExpressionContext
}

func NewPrimaryExpressionPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExpressionPEASTContext {
	var p = new(PrimaryExpressionPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *PrimaryExpressionPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionPEASTContext) PrimaryExpression() IPrimaryExpressionContext {
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

func (s *PrimaryExpressionPEASTContext) Selector() ISelectorContext {
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

func (s *PrimaryExpressionPEASTContext) Index() IIndexContext {
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

func (s *PrimaryExpressionPEASTContext) Arguments() IArgumentsContext {
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

func (s *PrimaryExpressionPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterPrimaryExpressionPEAST(s)
	}
}

func (s *PrimaryExpressionPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitPrimaryExpressionPEAST(s)
	}
}

func (s *PrimaryExpressionPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitPrimaryExpressionPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type LengthExpressionPEASTContext struct {
	PrimaryExpressionContext
}

func NewLengthExpressionPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthExpressionPEASTContext {
	var p = new(LengthExpressionPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *LengthExpressionPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionPEASTContext) LengthExpression() ILengthExpressionContext {
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

func (s *LengthExpressionPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLengthExpressionPEAST(s)
	}
}

func (s *LengthExpressionPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLengthExpressionPEAST(s)
	}
}

func (s *LengthExpressionPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLengthExpressionPEAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendExpressionPEASTContext struct {
	PrimaryExpressionContext
}

func NewAppendExpressionPEASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExpressionPEASTContext {
	var p = new(AppendExpressionPEASTContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *AppendExpressionPEASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionPEASTContext) AppendExpression() IAppendExpressionContext {
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

func (s *AppendExpressionPEASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterAppendExpressionPEAST(s)
	}
}

func (s *AppendExpressionPEASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitAppendExpressionPEAST(s)
	}
}

func (s *AppendExpressionPEASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitAppendExpressionPEAST(s)

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
		localctx = NewAppendExpressionPEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.AppendExpression()
		}

	case goParserLEN:
		localctx = NewLengthExpressionPEASTContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.LengthExpression()
		}

	case goParserCAP:
		localctx = NewCapExpressionPEASTContext(p, localctx)
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
			localctx = NewPrimaryExpressionPEASTContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, goParserRULE_primaryExpression)
			p.SetState(263)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				goto errorExit
			}
			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case goParserDOT:
				{
					p.SetState(264)
					p.Selector()
				}

			case goParserLSB:
				{
					p.SetState(265)
					p.Index()
				}

			case goParserLP:
				{
					p.SetState(266)
					p.Arguments()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

type LpOASTContext struct {
	OperandContext
}

func NewLpOASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LpOASTContext {
	var p = new(LpOASTContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *LpOASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LpOASTContext) LP() antlr.TerminalNode {
	return s.GetToken(goParserLP, 0)
}

func (s *LpOASTContext) Expression() IExpressionContext {
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

func (s *LpOASTContext) RP() antlr.TerminalNode {
	return s.GetToken(goParserRP, 0)
}

func (s *LpOASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterLpOAST(s)
	}
}

func (s *LpOASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitLpOAST(s)
	}
}

func (s *LpOASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitLpOAST(s)

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
		localctx = NewLpOASTContext(p, localctx)
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

type RunliteralContext struct {
	LiteralContext
}

func NewRunliteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunliteralContext {
	var p = new(RunliteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RunliteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunliteralContext) RUNELITERAL() antlr.TerminalNode {
	return s.GetToken(goParserRUNELITERAL, 0)
}

func (s *RunliteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterRunliteral(s)
	}
}

func (s *RunliteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitRunliteral(s)
	}
}

func (s *RunliteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitRunliteral(s)

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
		localctx = NewRunliteralContext(p, localctx)
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
	case goParserLP, goParserPL, goParserMIN, goParserMUL, goParserDIV, goParserMOD, goParserLSH, goParserRSH, goParserAMPER, goParserBC, goParserVB, goParserCARET, goParserEQ, goParserNEQ, goParserLT, goParserGT, goParserLTE, goParserGTE, goParserLAND, goParserLOR, goParserNEG, goParserASOP, goParserAAOP, goParserBAOP, goParserSAOP, goParserBOAOP, goParserMAOP, goParserBXOOP, goParserLSAOP, goParserRSAOP, goParserBCAOP, goParserRAOP, goParserDAOP, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
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

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3459960782471317386) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&259071) != 0) {
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
		case goParserLP, goParserPL, goParserMIN, goParserMUL, goParserDIV, goParserMOD, goParserLSH, goParserRSH, goParserAMPER, goParserBC, goParserVB, goParserCARET, goParserEQ, goParserNEQ, goParserLT, goParserGT, goParserLTE, goParserGTE, goParserLAND, goParserLOR, goParserNEG, goParserASOP, goParserAAOP, goParserBAOP, goParserSAOP, goParserBOAOP, goParserMAOP, goParserBXOOP, goParserLSAOP, goParserRSAOP, goParserBCAOP, goParserRAOP, goParserDAOP, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
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
		case goParserLP, goParserPL, goParserMIN, goParserMUL, goParserDIV, goParserMOD, goParserLSH, goParserRSH, goParserAMPER, goParserBC, goParserVB, goParserCARET, goParserEQ, goParserNEQ, goParserLT, goParserGT, goParserLTE, goParserGTE, goParserLAND, goParserLOR, goParserNEG, goParserASOP, goParserAAOP, goParserBAOP, goParserSAOP, goParserBOAOP, goParserMAOP, goParserBXOOP, goParserLSAOP, goParserRSAOP, goParserBCAOP, goParserRAOP, goParserDAOP, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
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
		case goParserLP, goParserPL, goParserMIN, goParserMUL, goParserDIV, goParserMOD, goParserLSH, goParserRSH, goParserAMPER, goParserBC, goParserVB, goParserCARET, goParserEQ, goParserNEQ, goParserLT, goParserGT, goParserLTE, goParserGTE, goParserLAND, goParserLOR, goParserNEG, goParserASOP, goParserAAOP, goParserBAOP, goParserSAOP, goParserBOAOP, goParserMAOP, goParserBXOOP, goParserLSAOP, goParserRSAOP, goParserBCAOP, goParserRAOP, goParserDAOP, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
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

	case goParserSEMI, goParserLP, goParserPL, goParserMIN, goParserMUL, goParserDIV, goParserMOD, goParserLSH, goParserRSH, goParserAMPER, goParserBC, goParserVB, goParserCARET, goParserEQ, goParserNEQ, goParserLT, goParserGT, goParserLTE, goParserGTE, goParserLAND, goParserLOR, goParserNEG, goParserASOP, goParserAAOP, goParserBAOP, goParserSAOP, goParserBOAOP, goParserMAOP, goParserBXOOP, goParserLSAOP, goParserRSAOP, goParserBCAOP, goParserRAOP, goParserDAOP, goParserAPPEND, goParserLEN, goParserCAP, goParserID, goParserINTLITERAL, goParserFLOATLITERAL, goParserRUNELITERAL, goParserRAWSTRINGLITERAL, goParserINTERPRETEDSTRINGLITERAL:
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

type AssignmentStatementSSContext struct {
	SimpleStatementContext
}

func NewAssignmentStatementSSContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementSSContext {
	var p = new(AssignmentStatementSSContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *AssignmentStatementSSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementSSContext) AssignmentStatement() IAssignmentStatementContext {
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

func (s *AssignmentStatementSSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterAssignmentStatementSS(s)
	}
}

func (s *AssignmentStatementSSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitAssignmentStatementSS(s)
	}
}

func (s *AssignmentStatementSSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitAssignmentStatementSS(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionListSSContext struct {
	SimpleStatementContext
}

func NewExpressionListSSContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListSSContext {
	var p = new(ExpressionListSSContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpressionListSSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListSSContext) AllExpressionList() []IExpressionListContext {
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

func (s *ExpressionListSSContext) ExpressionList(i int) IExpressionListContext {
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

func (s *ExpressionListSSContext) SVD() antlr.TerminalNode {
	return s.GetToken(goParserSVD, 0)
}

func (s *ExpressionListSSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionListSS(s)
	}
}

func (s *ExpressionListSSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionListSS(s)
	}
}

func (s *ExpressionListSSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionListSS(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionSSASTContext struct {
	SimpleStatementContext
}

func NewExpressionSSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSSASTContext {
	var p = new(ExpressionSSASTContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpressionSSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSSASTContext) Expression() IExpressionContext {
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

func (s *ExpressionSSASTContext) INC() antlr.TerminalNode {
	return s.GetToken(goParserINC, 0)
}

func (s *ExpressionSSASTContext) DEC() antlr.TerminalNode {
	return s.GetToken(goParserDEC, 0)
}

func (s *ExpressionSSASTContext) Epsilon() IEpsilonContext {
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

func (s *ExpressionSSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionSSAST(s)
	}
}

func (s *ExpressionSSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionSSAST(s)
	}
}

func (s *ExpressionSSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionSSAST(s)

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

func (p *goParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, goParserRULE_simpleStatement)
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEpsilonSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.Epsilon()
		}

	case 2:
		localctx = NewExpressionSSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.expression(0)
		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case goParserINC:
			{
				p.SetState(380)
				p.Match(goParserINC)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case goParserDEC:
			{
				p.SetState(381)
				p.Match(goParserDEC)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case goParserSEMI, goParserLCB:
			{
				p.SetState(382)
				p.Epsilon()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case 3:
		localctx = NewAssignmentStatementSSContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(385)
			p.AssignmentStatement()
		}

	case 4:
		localctx = NewExpressionListSSContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(386)
			p.ExpressionList()
		}
		{
			p.SetState(387)
			p.Match(goParserSVD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
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

type ExpressionASASTContext struct {
	AssignmentStatementContext
}

func NewExpressionASASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionASASTContext {
	var p = new(ExpressionASASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *ExpressionASASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionASASTContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionASASTContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionASASTContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExpressionASASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionASAST(s)
	}
}

func (s *ExpressionASASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionASAST(s)
	}
}

func (s *ExpressionASASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionASAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionListASASTContext struct {
	AssignmentStatementContext
}

func NewExpressionListASASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListASASTContext {
	var p = new(ExpressionListASASTContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *ExpressionListASASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListASASTContext) AllExpressionList() []IExpressionListContext {
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

func (s *ExpressionListASASTContext) ExpressionList(i int) IExpressionListContext {
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

func (s *ExpressionListASASTContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExpressionListASASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionListASAST(s)
	}
}

func (s *ExpressionListASASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionListASAST(s)
	}
}

func (s *ExpressionListASASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionListASAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, goParserRULE_assignmentStatement)
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionListASASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.ExpressionList()
		}
		{
			p.SetState(393)
			p.Operator()
		}
		{
			p.SetState(394)
			p.ExpressionList()
		}

	case 2:
		localctx = NewExpressionASASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.expression(0)
		}
		{
			p.SetState(397)
			p.Operator()
		}
		{
			p.SetState(398)
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

type IsSimpleStamentExpressionBlockISASTContext struct {
	IfStatementContext
}

func NewIsSimpleStamentExpressionBlockISASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsSimpleStamentExpressionBlockISASTContext {
	var p = new(IsSimpleStamentExpressionBlockISASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsSimpleStamentExpressionBlockISASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsSimpleStamentExpressionBlockISASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsSimpleStamentExpressionBlockISASTContext) SimpleStatement() ISimpleStatementContext {
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

func (s *IsSimpleStamentExpressionBlockISASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *IsSimpleStamentExpressionBlockISASTContext) Expression() IExpressionContext {
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

func (s *IsSimpleStamentExpressionBlockISASTContext) Block() IBlockContext {
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

func (s *IsSimpleStamentExpressionBlockISASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsSimpleStamentExpressionBlockISAST(s)
	}
}

func (s *IsSimpleStamentExpressionBlockISASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsSimpleStamentExpressionBlockISAST(s)
	}
}

func (s *IsSimpleStamentExpressionBlockISASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsSimpleStamentExpressionBlockISAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsSimpleStamentBlockISASTContext struct {
	IfStatementContext
}

func NewIsSimpleStamentBlockISASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsSimpleStamentBlockISASTContext {
	var p = new(IsSimpleStamentBlockISASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsSimpleStamentBlockISASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsSimpleStamentBlockISASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsSimpleStamentBlockISASTContext) Expression() IExpressionContext {
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

func (s *IsSimpleStamentBlockISASTContext) AllBlock() []IBlockContext {
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

func (s *IsSimpleStamentBlockISASTContext) Block(i int) IBlockContext {
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

func (s *IsSimpleStamentBlockISASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsSimpleStamentBlockISASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsSimpleStamentBlockISAST(s)
	}
}

func (s *IsSimpleStamentBlockISASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsSimpleStamentBlockISAST(s)
	}
}

func (s *IsSimpleStamentBlockISASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsSimpleStamentBlockISAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsExpressionBlockIsISASTContext struct {
	IfStatementContext
}

func NewIsExpressionBlockIsISASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsExpressionBlockIsISASTContext {
	var p = new(IsExpressionBlockIsISASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsExpressionBlockIsISASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsExpressionBlockIsISASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsExpressionBlockIsISASTContext) Expression() IExpressionContext {
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

func (s *IsExpressionBlockIsISASTContext) Block() IBlockContext {
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

func (s *IsExpressionBlockIsISASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsExpressionBlockIsISASTContext) IfStatement() IIfStatementContext {
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

func (s *IsExpressionBlockIsISASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsExpressionBlockIsISAST(s)
	}
}

func (s *IsExpressionBlockIsISASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsExpressionBlockIsISAST(s)
	}
}

func (s *IsExpressionBlockIsISASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsExpressionBlockIsISAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsSimpleStamentExpressionBlockBlockASTContext struct {
	IfStatementContext
}

func NewIsSimpleStamentExpressionBlockBlockASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsSimpleStamentExpressionBlockBlockASTContext {
	var p = new(IsSimpleStamentExpressionBlockBlockASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsSimpleStamentExpressionBlockBlockASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsSimpleStamentExpressionBlockBlockASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsSimpleStamentExpressionBlockBlockASTContext) SimpleStatement() ISimpleStatementContext {
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

func (s *IsSimpleStamentExpressionBlockBlockASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *IsSimpleStamentExpressionBlockBlockASTContext) Expression() IExpressionContext {
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

func (s *IsSimpleStamentExpressionBlockBlockASTContext) AllBlock() []IBlockContext {
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

func (s *IsSimpleStamentExpressionBlockBlockASTContext) Block(i int) IBlockContext {
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

func (s *IsSimpleStamentExpressionBlockBlockASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsSimpleStamentExpressionBlockBlockASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsSimpleStamentExpressionBlockBlockAST(s)
	}
}

func (s *IsSimpleStamentExpressionBlockBlockASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsSimpleStamentExpressionBlockBlockAST(s)
	}
}

func (s *IsSimpleStamentExpressionBlockBlockASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsSimpleStamentExpressionBlockBlockAST(s)

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

type IsSimpleStamentExpressionBlockifSASTContext struct {
	IfStatementContext
}

func NewIsSimpleStamentExpressionBlockifSASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsSimpleStamentExpressionBlockifSASTContext {
	var p = new(IsSimpleStamentExpressionBlockifSASTContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IsSimpleStamentExpressionBlockifSASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsSimpleStamentExpressionBlockifSASTContext) IF() antlr.TerminalNode {
	return s.GetToken(goParserIF, 0)
}

func (s *IsSimpleStamentExpressionBlockifSASTContext) SimpleStatement() ISimpleStatementContext {
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

func (s *IsSimpleStamentExpressionBlockifSASTContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goParserSEMI, 0)
}

func (s *IsSimpleStamentExpressionBlockifSASTContext) Expression() IExpressionContext {
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

func (s *IsSimpleStamentExpressionBlockifSASTContext) Block() IBlockContext {
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

func (s *IsSimpleStamentExpressionBlockifSASTContext) ELSE() antlr.TerminalNode {
	return s.GetToken(goParserELSE, 0)
}

func (s *IsSimpleStamentExpressionBlockifSASTContext) IfStatement() IIfStatementContext {
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

func (s *IsSimpleStamentExpressionBlockifSASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterIsSimpleStamentExpressionBlockifSAST(s)
	}
}

func (s *IsSimpleStamentExpressionBlockifSASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitIsSimpleStamentExpressionBlockifSAST(s)
	}
}

func (s *IsSimpleStamentExpressionBlockifSASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitIsSimpleStamentExpressionBlockifSAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, goParserRULE_ifStatement)
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIsExpressionBlockISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.expression(0)
		}
		{
			p.SetState(404)
			p.Block()
		}

	case 2:
		localctx = NewIsExpressionBlockIsISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.expression(0)
		}
		{
			p.SetState(408)
			p.Block()
		}
		{
			p.SetState(409)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.IfStatement()
		}

	case 3:
		localctx = NewIsSimpleStamentBlockISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(412)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.expression(0)
		}
		{
			p.SetState(414)
			p.Block()
		}
		{
			p.SetState(415)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Block()
		}

	case 4:
		localctx = NewIsSimpleStamentExpressionBlockISASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(418)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.SimpleStatement()
		}
		{
			p.SetState(420)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.expression(0)
		}
		{
			p.SetState(422)
			p.Block()
		}

	case 5:
		localctx = NewIsSimpleStamentExpressionBlockifSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(424)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.SimpleStatement()
		}
		{
			p.SetState(426)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.expression(0)
		}
		{
			p.SetState(428)
			p.Block()
		}
		{
			p.SetState(429)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.IfStatement()
		}

	case 6:
		localctx = NewIsSimpleStamentExpressionBlockBlockASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(432)
			p.Match(goParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.SimpleStatement()
		}
		{
			p.SetState(434)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.expression(0)
		}
		{
			p.SetState(436)
			p.Block()
		}
		{
			p.SetState(437)
			p.Match(goParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
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

type FSimpleStatementLASTContext struct {
	LoopContext
}

func NewFSimpleStatementLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FSimpleStatementLASTContext {
	var p = new(FSimpleStatementLASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *FSimpleStatementLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FSimpleStatementLASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(goParserFOR, 0)
}

func (s *FSimpleStatementLASTContext) AllSimpleStatement() []ISimpleStatementContext {
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

func (s *FSimpleStatementLASTContext) SimpleStatement(i int) ISimpleStatementContext {
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

func (s *FSimpleStatementLASTContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(goParserSEMI)
}

func (s *FSimpleStatementLASTContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(goParserSEMI, i)
}

func (s *FSimpleStatementLASTContext) Expression() IExpressionContext {
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

func (s *FSimpleStatementLASTContext) Block() IBlockContext {
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

func (s *FSimpleStatementLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFSimpleStatementLAST(s)
	}
}

func (s *FSimpleStatementLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFSimpleStatementLAST(s)
	}
}

func (s *FSimpleStatementLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFSimpleStatementLAST(s)

	default:
		return t.VisitChildren(s)
	}
}

type FExpressionBlockLASTContext struct {
	LoopContext
}

func NewFExpressionBlockLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FExpressionBlockLASTContext {
	var p = new(FExpressionBlockLASTContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *FExpressionBlockLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FExpressionBlockLASTContext) FOR() antlr.TerminalNode {
	return s.GetToken(goParserFOR, 0)
}

func (s *FExpressionBlockLASTContext) Expression() IExpressionContext {
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

func (s *FExpressionBlockLASTContext) Block() IBlockContext {
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

func (s *FExpressionBlockLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterFExpressionBlockLAST(s)
	}
}

func (s *FExpressionBlockLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitFExpressionBlockLAST(s)
	}
}

func (s *FExpressionBlockLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitFExpressionBlockLAST(s)

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

func (p *goParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, goParserRULE_loop)
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFBlockLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(443)
			p.Block()
		}

	case 2:
		localctx = NewFExpressionBlockLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.expression(0)
		}
		{
			p.SetState(446)
			p.Block()
		}

	case 3:
		localctx = NewFSimpleStatementLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(448)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.SimpleStatement()
		}
		{
			p.SetState(450)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.expression(0)
		}
		{
			p.SetState(452)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.SimpleStatement()
		}
		{
			p.SetState(454)
			p.Block()
		}

	case 4:
		localctx = NewFSimpleStatementLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(456)
			p.Match(goParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.SimpleStatement()
		}
		{
			p.SetState(458)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.Match(goParserSEMI)
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
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSSimpleStatSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.SimpleStatement()
		}
		{
			p.SetState(467)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.expression(0)
		}
		{
			p.SetState(469)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(471)
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
			p.SetState(473)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)
			p.expression(0)
		}
		{
			p.SetState(475)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(477)
			p.Match(goParserRCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSSimpleStatSASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(479)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.SimpleStatement()
		}
		{
			p.SetState(481)
			p.Match(goParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(484)
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
			p.SetState(486)
			p.Match(goParserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(487)
			p.Match(goParserLCB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(488)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(489)
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

type ExpressionCaseClauseECCLASTContext struct {
	ExpressionCaseClauseListContext
}

func NewExpressionCaseClauseECCLASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionCaseClauseECCLASTContext {
	var p = new(ExpressionCaseClauseECCLASTContext)

	InitEmptyExpressionCaseClauseListContext(&p.ExpressionCaseClauseListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseListContext))

	return p
}

func (s *ExpressionCaseClauseECCLASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseECCLASTContext) ExpressionCaseClause() IExpressionCaseClauseContext {
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

func (s *ExpressionCaseClauseECCLASTContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
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

func (s *ExpressionCaseClauseECCLASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionCaseClauseECCLAST(s)
	}
}

func (s *ExpressionCaseClauseECCLASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionCaseClauseECCLAST(s)
	}
}

func (s *ExpressionCaseClauseECCLASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionCaseClauseECCLAST(s)

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
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserRCB:
		localctx = NewEpsilonECCLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(493)
			p.Epsilon()
		}

	case goParserCASE, goParserDEFAULT:
		localctx = NewExpressionCaseClauseECCLASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(494)
			p.ExpressionCaseClause()
		}
		{
			p.SetState(495)
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

type ExpressionSwitchCaseECCASTContext struct {
	ExpressionCaseClauseContext
}

func NewExpressionSwitchCaseECCASTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSwitchCaseECCASTContext {
	var p = new(ExpressionSwitchCaseECCASTContext)

	InitEmptyExpressionCaseClauseContext(&p.ExpressionCaseClauseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionCaseClauseContext))

	return p
}

func (s *ExpressionSwitchCaseECCASTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseECCASTContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
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

func (s *ExpressionSwitchCaseECCASTContext) COL() antlr.TerminalNode {
	return s.GetToken(goParserCOL, 0)
}

func (s *ExpressionSwitchCaseECCASTContext) StatementList() IStatementListContext {
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

func (s *ExpressionSwitchCaseECCASTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterExpressionSwitchCaseECCAST(s)
	}
}

func (s *ExpressionSwitchCaseECCASTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitExpressionSwitchCaseECCAST(s)
	}
}

func (s *ExpressionSwitchCaseECCASTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitExpressionSwitchCaseECCAST(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, goParserRULE_expressionCaseClause)
	localctx = NewExpressionSwitchCaseECCASTContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(500)
		p.Match(goParserCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(501)
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
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case goParserCASE:
		localctx = NewCaseESCASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(503)
			p.Match(goParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.ExpressionList()
		}

	case goParserDEFAULT:
		localctx = NewDefaultESCASTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(505)
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

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LSH() antlr.TerminalNode
	RSH() antlr.TerminalNode
	AMPER() antlr.TerminalNode
	BC() antlr.TerminalNode
	PL() antlr.TerminalNode
	MIN() antlr.TerminalNode
	VB() antlr.TerminalNode
	CARET() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LAND() antlr.TerminalNode
	LOR() antlr.TerminalNode
	NEG() antlr.TerminalNode
	ASOP() antlr.TerminalNode
	AAOP() antlr.TerminalNode
	BAOP() antlr.TerminalNode
	SAOP() antlr.TerminalNode
	BOAOP() antlr.TerminalNode
	MAOP() antlr.TerminalNode
	BXOOP() antlr.TerminalNode
	LSAOP() antlr.TerminalNode
	RSAOP() antlr.TerminalNode
	BCAOP() antlr.TerminalNode
	RAOP() antlr.TerminalNode
	DAOP() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = goParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = goParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(goParserMUL, 0)
}

func (s *OperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(goParserDIV, 0)
}

func (s *OperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(goParserMOD, 0)
}

func (s *OperatorContext) LSH() antlr.TerminalNode {
	return s.GetToken(goParserLSH, 0)
}

func (s *OperatorContext) RSH() antlr.TerminalNode {
	return s.GetToken(goParserRSH, 0)
}

func (s *OperatorContext) AMPER() antlr.TerminalNode {
	return s.GetToken(goParserAMPER, 0)
}

func (s *OperatorContext) BC() antlr.TerminalNode {
	return s.GetToken(goParserBC, 0)
}

func (s *OperatorContext) PL() antlr.TerminalNode {
	return s.GetToken(goParserPL, 0)
}

func (s *OperatorContext) MIN() antlr.TerminalNode {
	return s.GetToken(goParserMIN, 0)
}

func (s *OperatorContext) VB() antlr.TerminalNode {
	return s.GetToken(goParserVB, 0)
}

func (s *OperatorContext) CARET() antlr.TerminalNode {
	return s.GetToken(goParserCARET, 0)
}

func (s *OperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(goParserEQ, 0)
}

func (s *OperatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(goParserNEQ, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(goParserLT, 0)
}

func (s *OperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(goParserLTE, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(goParserGT, 0)
}

func (s *OperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(goParserGTE, 0)
}

func (s *OperatorContext) LAND() antlr.TerminalNode {
	return s.GetToken(goParserLAND, 0)
}

func (s *OperatorContext) LOR() antlr.TerminalNode {
	return s.GetToken(goParserLOR, 0)
}

func (s *OperatorContext) NEG() antlr.TerminalNode {
	return s.GetToken(goParserNEG, 0)
}

func (s *OperatorContext) ASOP() antlr.TerminalNode {
	return s.GetToken(goParserASOP, 0)
}

func (s *OperatorContext) AAOP() antlr.TerminalNode {
	return s.GetToken(goParserAAOP, 0)
}

func (s *OperatorContext) BAOP() antlr.TerminalNode {
	return s.GetToken(goParserBAOP, 0)
}

func (s *OperatorContext) SAOP() antlr.TerminalNode {
	return s.GetToken(goParserSAOP, 0)
}

func (s *OperatorContext) BOAOP() antlr.TerminalNode {
	return s.GetToken(goParserBOAOP, 0)
}

func (s *OperatorContext) MAOP() antlr.TerminalNode {
	return s.GetToken(goParserMAOP, 0)
}

func (s *OperatorContext) BXOOP() antlr.TerminalNode {
	return s.GetToken(goParserBXOOP, 0)
}

func (s *OperatorContext) LSAOP() antlr.TerminalNode {
	return s.GetToken(goParserLSAOP, 0)
}

func (s *OperatorContext) RSAOP() antlr.TerminalNode {
	return s.GetToken(goParserRSAOP, 0)
}

func (s *OperatorContext) BCAOP() antlr.TerminalNode {
	return s.GetToken(goParserBCAOP, 0)
}

func (s *OperatorContext) RAOP() antlr.TerminalNode {
	return s.GetToken(goParserRAOP, 0)
}

func (s *OperatorContext) DAOP() antlr.TerminalNode {
	return s.GetToken(goParserDAOP, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goParserListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (s *OperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goParserVisitor:
		return t.VisitOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, goParserRULE_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70368743917440) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.EnterRule(localctx, 82, goParserRULE_epsilon)
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
