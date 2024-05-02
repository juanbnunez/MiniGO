// Code generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goScanner.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type goScanner struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GoScannerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goscannerLexerInit() {
	staticData := &GoScannerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"SEMI", "SVD", "LP", "RP", "TIL", "COL", "PL", "MIN", "MUL", "DIV",
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
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 85, 536, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57,
		1, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1,
		59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61,
		1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1,
		63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65,
		1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1,
		67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68,
		1, 68, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1, 70, 1,
		70, 1, 70, 1, 70, 1, 70, 1, 70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 71, 1, 72,
		1, 72, 1, 72, 1, 72, 1, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1, 73, 1,
		73, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 74, 1, 75, 1, 75,
		3, 75, 445, 8, 75, 1, 75, 1, 75, 1, 75, 1, 75, 5, 75, 451, 8, 75, 10, 75,
		12, 75, 454, 9, 75, 1, 76, 4, 76, 457, 8, 76, 11, 76, 12, 76, 458, 1, 77,
		4, 77, 462, 8, 77, 11, 77, 12, 77, 463, 1, 77, 1, 77, 5, 77, 468, 8, 77,
		10, 77, 12, 77, 471, 9, 77, 1, 77, 1, 77, 3, 77, 475, 8, 77, 1, 77, 4,
		77, 478, 8, 77, 11, 77, 12, 77, 479, 3, 77, 482, 8, 77, 1, 78, 1, 78, 1,
		78, 1, 78, 1, 78, 1, 78, 1, 78, 3, 78, 491, 8, 78, 1, 79, 1, 79, 1, 79,
		1, 79, 5, 79, 497, 8, 79, 10, 79, 12, 79, 500, 9, 79, 1, 79, 1, 79, 1,
		80, 1, 80, 1, 80, 1, 80, 5, 80, 508, 8, 80, 10, 80, 12, 80, 511, 9, 80,
		1, 80, 1, 80, 1, 81, 1, 81, 1, 82, 1, 82, 1, 83, 1, 83, 1, 83, 1, 83, 5,
		83, 523, 8, 83, 10, 83, 12, 83, 526, 9, 83, 1, 83, 1, 83, 1, 84, 4, 84,
		531, 8, 84, 11, 84, 12, 84, 532, 1, 84, 1, 84, 0, 0, 85, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97,
		49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113,
		57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 64, 129,
		65, 131, 66, 133, 67, 135, 68, 137, 69, 139, 70, 141, 71, 143, 72, 145,
		73, 147, 74, 149, 75, 151, 76, 153, 77, 155, 78, 157, 79, 159, 80, 161,
		81, 163, 82, 165, 83, 167, 84, 169, 85, 1, 0, 8, 2, 0, 69, 69, 101, 101,
		2, 0, 43, 43, 45, 45, 6, 0, 34, 34, 39, 39, 92, 92, 110, 110, 114, 114,
		116, 116, 2, 0, 92, 92, 96, 96, 2, 0, 34, 34, 92, 92, 2, 0, 65, 90, 97,
		122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 552, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1,
		0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71,
		1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0,
		79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0,
		0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0,
		0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1,
		0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0,
		109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0,
		0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123,
		1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0,
		0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1,
		0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0,
		145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 151, 1, 0,
		0, 0, 0, 153, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 0, 157, 1, 0, 0, 0, 0, 159,
		1, 0, 0, 0, 0, 161, 1, 0, 0, 0, 0, 163, 1, 0, 0, 0, 0, 165, 1, 0, 0, 0,
		0, 167, 1, 0, 0, 0, 0, 169, 1, 0, 0, 0, 1, 171, 1, 0, 0, 0, 3, 173, 1,
		0, 0, 0, 5, 176, 1, 0, 0, 0, 7, 178, 1, 0, 0, 0, 9, 180, 1, 0, 0, 0, 11,
		182, 1, 0, 0, 0, 13, 184, 1, 0, 0, 0, 15, 186, 1, 0, 0, 0, 17, 188, 1,
		0, 0, 0, 19, 190, 1, 0, 0, 0, 21, 192, 1, 0, 0, 0, 23, 194, 1, 0, 0, 0,
		25, 196, 1, 0, 0, 0, 27, 198, 1, 0, 0, 0, 29, 200, 1, 0, 0, 0, 31, 202,
		1, 0, 0, 0, 33, 205, 1, 0, 0, 0, 35, 208, 1, 0, 0, 0, 37, 210, 1, 0, 0,
		0, 39, 213, 1, 0, 0, 0, 41, 216, 1, 0, 0, 0, 43, 218, 1, 0, 0, 0, 45, 221,
		1, 0, 0, 0, 47, 223, 1, 0, 0, 0, 49, 225, 1, 0, 0, 0, 51, 228, 1, 0, 0,
		0, 53, 231, 1, 0, 0, 0, 55, 233, 1, 0, 0, 0, 57, 235, 1, 0, 0, 0, 59, 238,
		1, 0, 0, 0, 61, 241, 1, 0, 0, 0, 63, 244, 1, 0, 0, 0, 65, 247, 1, 0, 0,
		0, 67, 249, 1, 0, 0, 0, 69, 251, 1, 0, 0, 0, 71, 254, 1, 0, 0, 0, 73, 257,
		1, 0, 0, 0, 75, 260, 1, 0, 0, 0, 77, 263, 1, 0, 0, 0, 79, 266, 1, 0, 0,
		0, 81, 269, 1, 0, 0, 0, 83, 273, 1, 0, 0, 0, 85, 277, 1, 0, 0, 0, 87, 281,
		1, 0, 0, 0, 89, 284, 1, 0, 0, 0, 91, 287, 1, 0, 0, 0, 93, 289, 1, 0, 0,
		0, 95, 291, 1, 0, 0, 0, 97, 293, 1, 0, 0, 0, 99, 301, 1, 0, 0, 0, 101,
		304, 1, 0, 0, 0, 103, 310, 1, 0, 0, 0, 105, 314, 1, 0, 0, 0, 107, 319,
		1, 0, 0, 0, 109, 324, 1, 0, 0, 0, 111, 327, 1, 0, 0, 0, 113, 330, 1, 0,
		0, 0, 115, 336, 1, 0, 0, 0, 117, 340, 1, 0, 0, 0, 119, 346, 1, 0, 0, 0,
		121, 350, 1, 0, 0, 0, 123, 355, 1, 0, 0, 0, 125, 360, 1, 0, 0, 0, 127,
		367, 1, 0, 0, 0, 129, 374, 1, 0, 0, 0, 131, 378, 1, 0, 0, 0, 133, 382,
		1, 0, 0, 0, 135, 388, 1, 0, 0, 0, 137, 396, 1, 0, 0, 0, 139, 403, 1, 0,
		0, 0, 141, 409, 1, 0, 0, 0, 143, 418, 1, 0, 0, 0, 145, 422, 1, 0, 0, 0,
		147, 429, 1, 0, 0, 0, 149, 434, 1, 0, 0, 0, 151, 444, 1, 0, 0, 0, 153,
		456, 1, 0, 0, 0, 155, 461, 1, 0, 0, 0, 157, 490, 1, 0, 0, 0, 159, 492,
		1, 0, 0, 0, 161, 503, 1, 0, 0, 0, 163, 514, 1, 0, 0, 0, 165, 516, 1, 0,
		0, 0, 167, 518, 1, 0, 0, 0, 169, 530, 1, 0, 0, 0, 171, 172, 5, 59, 0, 0,
		172, 2, 1, 0, 0, 0, 173, 174, 5, 58, 0, 0, 174, 175, 5, 61, 0, 0, 175,
		4, 1, 0, 0, 0, 176, 177, 5, 40, 0, 0, 177, 6, 1, 0, 0, 0, 178, 179, 5,
		41, 0, 0, 179, 8, 1, 0, 0, 0, 180, 181, 5, 126, 0, 0, 181, 10, 1, 0, 0,
		0, 182, 183, 5, 58, 0, 0, 183, 12, 1, 0, 0, 0, 184, 185, 5, 43, 0, 0, 185,
		14, 1, 0, 0, 0, 186, 187, 5, 45, 0, 0, 187, 16, 1, 0, 0, 0, 188, 189, 5,
		42, 0, 0, 189, 18, 1, 0, 0, 0, 190, 191, 5, 47, 0, 0, 191, 20, 1, 0, 0,
		0, 192, 193, 5, 44, 0, 0, 193, 22, 1, 0, 0, 0, 194, 195, 5, 91, 0, 0, 195,
		24, 1, 0, 0, 0, 196, 197, 5, 93, 0, 0, 197, 26, 1, 0, 0, 0, 198, 199, 5,
		123, 0, 0, 199, 28, 1, 0, 0, 0, 200, 201, 5, 125, 0, 0, 201, 30, 1, 0,
		0, 0, 202, 203, 5, 43, 0, 0, 203, 204, 5, 43, 0, 0, 204, 32, 1, 0, 0, 0,
		205, 206, 5, 45, 0, 0, 206, 207, 5, 45, 0, 0, 207, 34, 1, 0, 0, 0, 208,
		209, 5, 37, 0, 0, 209, 36, 1, 0, 0, 0, 210, 211, 5, 60, 0, 0, 211, 212,
		5, 60, 0, 0, 212, 38, 1, 0, 0, 0, 213, 214, 5, 62, 0, 0, 214, 215, 5, 62,
		0, 0, 215, 40, 1, 0, 0, 0, 216, 217, 5, 38, 0, 0, 217, 42, 1, 0, 0, 0,
		218, 219, 5, 38, 0, 0, 219, 220, 5, 94, 0, 0, 220, 44, 1, 0, 0, 0, 221,
		222, 5, 124, 0, 0, 222, 46, 1, 0, 0, 0, 223, 224, 5, 94, 0, 0, 224, 48,
		1, 0, 0, 0, 225, 226, 5, 61, 0, 0, 226, 227, 5, 61, 0, 0, 227, 50, 1, 0,
		0, 0, 228, 229, 5, 33, 0, 0, 229, 230, 5, 61, 0, 0, 230, 52, 1, 0, 0, 0,
		231, 232, 5, 60, 0, 0, 232, 54, 1, 0, 0, 0, 233, 234, 5, 62, 0, 0, 234,
		56, 1, 0, 0, 0, 235, 236, 5, 60, 0, 0, 236, 237, 5, 61, 0, 0, 237, 58,
		1, 0, 0, 0, 238, 239, 5, 62, 0, 0, 239, 240, 5, 61, 0, 0, 240, 60, 1, 0,
		0, 0, 241, 242, 5, 38, 0, 0, 242, 243, 5, 38, 0, 0, 243, 62, 1, 0, 0, 0,
		244, 245, 5, 124, 0, 0, 245, 246, 5, 124, 0, 0, 246, 64, 1, 0, 0, 0, 247,
		248, 5, 33, 0, 0, 248, 66, 1, 0, 0, 0, 249, 250, 5, 61, 0, 0, 250, 68,
		1, 0, 0, 0, 251, 252, 5, 43, 0, 0, 252, 253, 5, 61, 0, 0, 253, 70, 1, 0,
		0, 0, 254, 255, 5, 38, 0, 0, 255, 256, 5, 61, 0, 0, 256, 72, 1, 0, 0, 0,
		257, 258, 5, 45, 0, 0, 258, 259, 5, 61, 0, 0, 259, 74, 1, 0, 0, 0, 260,
		261, 5, 124, 0, 0, 261, 262, 5, 61, 0, 0, 262, 76, 1, 0, 0, 0, 263, 264,
		5, 42, 0, 0, 264, 265, 5, 61, 0, 0, 265, 78, 1, 0, 0, 0, 266, 267, 5, 94,
		0, 0, 267, 268, 5, 61, 0, 0, 268, 80, 1, 0, 0, 0, 269, 270, 5, 60, 0, 0,
		270, 271, 5, 60, 0, 0, 271, 272, 5, 61, 0, 0, 272, 82, 1, 0, 0, 0, 273,
		274, 5, 62, 0, 0, 274, 275, 5, 62, 0, 0, 275, 276, 5, 61, 0, 0, 276, 84,
		1, 0, 0, 0, 277, 278, 5, 38, 0, 0, 278, 279, 5, 94, 0, 0, 279, 280, 5,
		61, 0, 0, 280, 86, 1, 0, 0, 0, 281, 282, 5, 37, 0, 0, 282, 283, 5, 61,
		0, 0, 283, 88, 1, 0, 0, 0, 284, 285, 5, 47, 0, 0, 285, 286, 5, 61, 0, 0,
		286, 90, 1, 0, 0, 0, 287, 288, 5, 46, 0, 0, 288, 92, 1, 0, 0, 0, 289, 290,
		5, 92, 0, 0, 290, 94, 1, 0, 0, 0, 291, 292, 5, 39, 0, 0, 292, 96, 1, 0,
		0, 0, 293, 294, 5, 112, 0, 0, 294, 295, 5, 97, 0, 0, 295, 296, 5, 99, 0,
		0, 296, 297, 5, 107, 0, 0, 297, 298, 5, 97, 0, 0, 298, 299, 5, 103, 0,
		0, 299, 300, 5, 101, 0, 0, 300, 98, 1, 0, 0, 0, 301, 302, 5, 105, 0, 0,
		302, 303, 5, 102, 0, 0, 303, 100, 1, 0, 0, 0, 304, 305, 5, 119, 0, 0, 305,
		306, 5, 104, 0, 0, 306, 307, 5, 105, 0, 0, 307, 308, 5, 108, 0, 0, 308,
		309, 5, 101, 0, 0, 309, 102, 1, 0, 0, 0, 310, 311, 5, 108, 0, 0, 311, 312,
		5, 101, 0, 0, 312, 313, 5, 116, 0, 0, 313, 104, 1, 0, 0, 0, 314, 315, 5,
		116, 0, 0, 315, 316, 5, 104, 0, 0, 316, 317, 5, 101, 0, 0, 317, 318, 5,
		110, 0, 0, 318, 106, 1, 0, 0, 0, 319, 320, 5, 101, 0, 0, 320, 321, 5, 108,
		0, 0, 321, 322, 5, 115, 0, 0, 322, 323, 5, 101, 0, 0, 323, 108, 1, 0, 0,
		0, 324, 325, 5, 105, 0, 0, 325, 326, 5, 110, 0, 0, 326, 110, 1, 0, 0, 0,
		327, 328, 5, 100, 0, 0, 328, 329, 5, 111, 0, 0, 329, 112, 1, 0, 0, 0, 330,
		331, 5, 98, 0, 0, 331, 332, 5, 101, 0, 0, 332, 333, 5, 103, 0, 0, 333,
		334, 5, 105, 0, 0, 334, 335, 5, 110, 0, 0, 335, 114, 1, 0, 0, 0, 336, 337,
		5, 101, 0, 0, 337, 338, 5, 110, 0, 0, 338, 339, 5, 100, 0, 0, 339, 116,
		1, 0, 0, 0, 340, 341, 5, 99, 0, 0, 341, 342, 5, 111, 0, 0, 342, 343, 5,
		110, 0, 0, 343, 344, 5, 115, 0, 0, 344, 345, 5, 116, 0, 0, 345, 118, 1,
		0, 0, 0, 346, 347, 5, 118, 0, 0, 347, 348, 5, 97, 0, 0, 348, 349, 5, 114,
		0, 0, 349, 120, 1, 0, 0, 0, 350, 351, 5, 116, 0, 0, 351, 352, 5, 121, 0,
		0, 352, 353, 5, 112, 0, 0, 353, 354, 5, 101, 0, 0, 354, 122, 1, 0, 0, 0,
		355, 356, 5, 102, 0, 0, 356, 357, 5, 117, 0, 0, 357, 358, 5, 110, 0, 0,
		358, 359, 5, 99, 0, 0, 359, 124, 1, 0, 0, 0, 360, 361, 5, 115, 0, 0, 361,
		362, 5, 116, 0, 0, 362, 363, 5, 114, 0, 0, 363, 364, 5, 117, 0, 0, 364,
		365, 5, 99, 0, 0, 365, 366, 5, 116, 0, 0, 366, 126, 1, 0, 0, 0, 367, 368,
		5, 97, 0, 0, 368, 369, 5, 112, 0, 0, 369, 370, 5, 112, 0, 0, 370, 371,
		5, 101, 0, 0, 371, 372, 5, 110, 0, 0, 372, 373, 5, 100, 0, 0, 373, 128,
		1, 0, 0, 0, 374, 375, 5, 108, 0, 0, 375, 376, 5, 101, 0, 0, 376, 377, 5,
		110, 0, 0, 377, 130, 1, 0, 0, 0, 378, 379, 5, 99, 0, 0, 379, 380, 5, 97,
		0, 0, 380, 381, 5, 112, 0, 0, 381, 132, 1, 0, 0, 0, 382, 383, 5, 112, 0,
		0, 383, 384, 5, 114, 0, 0, 384, 385, 5, 105, 0, 0, 385, 386, 5, 110, 0,
		0, 386, 387, 5, 116, 0, 0, 387, 134, 1, 0, 0, 0, 388, 389, 5, 112, 0, 0,
		389, 390, 5, 114, 0, 0, 390, 391, 5, 105, 0, 0, 391, 392, 5, 110, 0, 0,
		392, 393, 5, 116, 0, 0, 393, 394, 5, 108, 0, 0, 394, 395, 5, 110, 0, 0,
		395, 136, 1, 0, 0, 0, 396, 397, 5, 114, 0, 0, 397, 398, 5, 101, 0, 0, 398,
		399, 5, 116, 0, 0, 399, 400, 5, 117, 0, 0, 400, 401, 5, 114, 0, 0, 401,
		402, 5, 110, 0, 0, 402, 138, 1, 0, 0, 0, 403, 404, 5, 98, 0, 0, 404, 405,
		5, 114, 0, 0, 405, 406, 5, 101, 0, 0, 406, 407, 5, 97, 0, 0, 407, 408,
		5, 107, 0, 0, 408, 140, 1, 0, 0, 0, 409, 410, 5, 99, 0, 0, 410, 411, 5,
		111, 0, 0, 411, 412, 5, 110, 0, 0, 412, 413, 5, 116, 0, 0, 413, 414, 5,
		105, 0, 0, 414, 415, 5, 110, 0, 0, 415, 416, 5, 117, 0, 0, 416, 417, 5,
		101, 0, 0, 417, 142, 1, 0, 0, 0, 418, 419, 5, 102, 0, 0, 419, 420, 5, 111,
		0, 0, 420, 421, 5, 114, 0, 0, 421, 144, 1, 0, 0, 0, 422, 423, 5, 115, 0,
		0, 423, 424, 5, 119, 0, 0, 424, 425, 5, 105, 0, 0, 425, 426, 5, 116, 0,
		0, 426, 427, 5, 99, 0, 0, 427, 428, 5, 104, 0, 0, 428, 146, 1, 0, 0, 0,
		429, 430, 5, 99, 0, 0, 430, 431, 5, 97, 0, 0, 431, 432, 5, 115, 0, 0, 432,
		433, 5, 101, 0, 0, 433, 148, 1, 0, 0, 0, 434, 435, 5, 100, 0, 0, 435, 436,
		5, 101, 0, 0, 436, 437, 5, 102, 0, 0, 437, 438, 5, 97, 0, 0, 438, 439,
		5, 117, 0, 0, 439, 440, 5, 108, 0, 0, 440, 441, 5, 116, 0, 0, 441, 150,
		1, 0, 0, 0, 442, 445, 5, 95, 0, 0, 443, 445, 1, 0, 0, 0, 444, 442, 1, 0,
		0, 0, 444, 443, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 452, 3, 163, 81,
		0, 447, 451, 3, 163, 81, 0, 448, 451, 3, 165, 82, 0, 449, 451, 5, 95, 0,
		0, 450, 447, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 450, 449, 1, 0, 0, 0, 451,
		454, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 152,
		1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 455, 457, 3, 165, 82, 0, 456, 455, 1,
		0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0,
		0, 459, 154, 1, 0, 0, 0, 460, 462, 3, 165, 82, 0, 461, 460, 1, 0, 0, 0,
		462, 463, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464,
		465, 1, 0, 0, 0, 465, 469, 5, 46, 0, 0, 466, 468, 3, 165, 82, 0, 467, 466,
		1, 0, 0, 0, 468, 471, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469, 470, 1, 0,
		0, 0, 470, 481, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 472, 474, 7, 0, 0, 0,
		473, 475, 7, 1, 0, 0, 474, 473, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475,
		477, 1, 0, 0, 0, 476, 478, 3, 165, 82, 0, 477, 476, 1, 0, 0, 0, 478, 479,
		1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 482, 1, 0,
		0, 0, 481, 472, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 156, 1, 0, 0, 0,
		483, 484, 5, 39, 0, 0, 484, 485, 9, 0, 0, 0, 485, 491, 5, 39, 0, 0, 486,
		487, 5, 39, 0, 0, 487, 488, 5, 92, 0, 0, 488, 489, 7, 2, 0, 0, 489, 491,
		5, 39, 0, 0, 490, 483, 1, 0, 0, 0, 490, 486, 1, 0, 0, 0, 491, 158, 1, 0,
		0, 0, 492, 498, 5, 96, 0, 0, 493, 494, 5, 92, 0, 0, 494, 497, 9, 0, 0,
		0, 495, 497, 8, 3, 0, 0, 496, 493, 1, 0, 0, 0, 496, 495, 1, 0, 0, 0, 497,
		500, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0, 498, 499, 1, 0, 0, 0, 499, 501,
		1, 0, 0, 0, 500, 498, 1, 0, 0, 0, 501, 502, 5, 96, 0, 0, 502, 160, 1, 0,
		0, 0, 503, 509, 5, 34, 0, 0, 504, 505, 5, 92, 0, 0, 505, 508, 9, 0, 0,
		0, 506, 508, 8, 4, 0, 0, 507, 504, 1, 0, 0, 0, 507, 506, 1, 0, 0, 0, 508,
		511, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 512,
		1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 512, 513, 5, 34, 0, 0, 513, 162, 1, 0,
		0, 0, 514, 515, 7, 5, 0, 0, 515, 164, 1, 0, 0, 0, 516, 517, 2, 48, 57,
		0, 517, 166, 1, 0, 0, 0, 518, 519, 5, 47, 0, 0, 519, 520, 5, 47, 0, 0,
		520, 524, 1, 0, 0, 0, 521, 523, 8, 6, 0, 0, 522, 521, 1, 0, 0, 0, 523,
		526, 1, 0, 0, 0, 524, 522, 1, 0, 0, 0, 524, 525, 1, 0, 0, 0, 525, 527,
		1, 0, 0, 0, 526, 524, 1, 0, 0, 0, 527, 528, 6, 83, 0, 0, 528, 168, 1, 0,
		0, 0, 529, 531, 7, 7, 0, 0, 530, 529, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0,
		532, 530, 1, 0, 0, 0, 532, 533, 1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534,
		535, 6, 84, 0, 0, 535, 170, 1, 0, 0, 0, 17, 0, 444, 450, 452, 458, 463,
		469, 474, 479, 481, 490, 496, 498, 507, 509, 524, 532, 1, 6, 0, 0,
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

// goScannerInit initializes any static state used to implement goScanner. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewgoScanner(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoScannerInit() {
	staticData := &GoScannerLexerStaticData
	staticData.once.Do(goscannerLexerInit)
}

// NewgoScanner produces a new lexer instance for the optional input antlr.CharStream.
func NewgoScanner(input antlr.CharStream) *goScanner {
	GoScannerInit()
	l := new(goScanner)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GoScannerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "goScanner.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// goScanner tokens.
const (
	goScannerSEMI                     = 1
	goScannerSVD                      = 2
	goScannerLP                       = 3
	goScannerRP                       = 4
	goScannerTIL                      = 5
	goScannerCOL                      = 6
	goScannerPL                       = 7
	goScannerMIN                      = 8
	goScannerMUL                      = 9
	goScannerDIV                      = 10
	goScannerCM                       = 11
	goScannerLSB                      = 12
	goScannerRSB                      = 13
	goScannerLCB                      = 14
	goScannerRCB                      = 15
	goScannerINC                      = 16
	goScannerDEC                      = 17
	goScannerMOD                      = 18
	goScannerLSH                      = 19
	goScannerRSH                      = 20
	goScannerAMPER                    = 21
	goScannerBC                       = 22
	goScannerVB                       = 23
	goScannerCARET                    = 24
	goScannerEQ                       = 25
	goScannerNEQ                      = 26
	goScannerLT                       = 27
	goScannerGT                       = 28
	goScannerLTE                      = 29
	goScannerGTE                      = 30
	goScannerLAND                     = 31
	goScannerLOR                      = 32
	goScannerNEG                      = 33
	goScannerASOP                     = 34
	goScannerAAOP                     = 35
	goScannerBAOP                     = 36
	goScannerSAOP                     = 37
	goScannerBOAOP                    = 38
	goScannerMAOP                     = 39
	goScannerBXOOP                    = 40
	goScannerLSAOP                    = 41
	goScannerRSAOP                    = 42
	goScannerBCAOP                    = 43
	goScannerRAOP                     = 44
	goScannerDAOP                     = 45
	goScannerDOT                      = 46
	goScannerDBS                      = 47
	goScannerBS                       = 48
	goScannerPACKAGE                  = 49
	goScannerIF                       = 50
	goScannerWHILE                    = 51
	goScannerLET                      = 52
	goScannerTHEN                     = 53
	goScannerELSE                     = 54
	goScannerIN                       = 55
	goScannerDO                       = 56
	goScannerBEGIN                    = 57
	goScannerEND                      = 58
	goScannerCONST                    = 59
	goScannerVAR                      = 60
	goScannerTYPE                     = 61
	goScannerFUNC                     = 62
	goScannerSTRUCT                   = 63
	goScannerAPPEND                   = 64
	goScannerLEN                      = 65
	goScannerCAP                      = 66
	goScannerPRINT                    = 67
	goScannerPRINTLN                  = 68
	goScannerRETURN                   = 69
	goScannerBREAK                    = 70
	goScannerCONTINUE                 = 71
	goScannerFOR                      = 72
	goScannerSWITCH                   = 73
	goScannerCASE                     = 74
	goScannerDEFAULT                  = 75
	goScannerID                       = 76
	goScannerINTLITERAL               = 77
	goScannerFLOATLITERAL             = 78
	goScannerRUNELITERAL              = 79
	goScannerRAWSTRINGLITERAL         = 80
	goScannerINTERPRETEDSTRINGLITERAL = 81
	goScannerLETTER                   = 82
	goScannerDIGIT                    = 83
	goScannerLINE_COMMENT             = 84
	goScannerSPACES                   = 85
)
