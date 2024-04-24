package compiler

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type ErrorType int

const (
	PARSER_ERROR ErrorType = iota
	SCANNER_ERROR
	OTHER_ERROR
)

type MyErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorMsgs []string
}

func NewMyErrorListener() *MyErrorListener {
	return &MyErrorListener{
		ErrorMsgs: make([]string, 0),
	}
}

func (er *MyErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	var errorType ErrorType

	switch recognizer.(type) {
	case antlr.Parser:
		errorType = PARSER_ERROR
	case antlr.Lexer:
		errorType = SCANNER_ERROR
	default:
		errorType = OTHER_ERROR
	}

	errorMsg := fmt.Sprintf("%s ERROR - line %d:%d %s", errorType, line, column, msg)
	er.ErrorMsgs = append(er.ErrorMsgs, errorMsg)
}

func (e *MyErrorListener) HasErrors() bool {
	return len(e.ErrorMsgs) > 0
}

func (e *MyErrorListener) String() string {
	if !e.HasErrors() {
		return "0 errors"
	}
	builder := ""
	for _, s := range e.ErrorMsgs {
		builder += s + "\n"
	}
	return builder
}
