package parser

import (
	"fmt"
	"strconv"
)

type MyErrorListener struct {
	errorMsgs []string
}

func NewMyErrorListener() *MyErrorListener {
	return &MyErrorListener{
		errorMsgs: make([]string, 0),
	}
}

func (m *MyErrorListener) syntaxError(recognizer interface{},
	offendingSymbol interface{},
	line int,
	charPositionInLine int,
	msg string,
	e interface{}) {

	switch recognizer.(type) {
	case *goParser:
		m.errorMsgs = append(m.errorMsgs, "PARSER ERROR - line "+strconv.Itoa(line)+":"+strconv.Itoa(charPositionInLine)+" "+msg)
	case *goScanner:
		m.errorMsgs = append(m.errorMsgs, "SCANNER ERROR - line "+strconv.Itoa(line)+":"+strconv.Itoa(charPositionInLine)+" "+msg)
	default:
		m.errorMsgs = append(m.errorMsgs, "Other Error")
	}
}

func (m *MyErrorListener) HasErrors() bool {
	return len(m.errorMsgs) > 0
}

func (m *MyErrorListener) String() string {
	if !m.HasErrors() {
		return "0 errors"
	}
	var builder string
	for _, s := range m.errorMsgs {
		builder += fmt.Sprintf("%s\n", s)
	}
	return builder
}
