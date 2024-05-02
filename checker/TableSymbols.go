package checker

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type Ident interface {
}

type VarIdent struct {
	tok        antlr.TerminalNode // token
	typ        int                // tipo
	level      int                // nivel
	isConstant bool
}

type MethodIdent struct {
	tok    antlr.TerminalNode
	typ    int
	level  int
	params []int
}

type SymbolTable struct {
	table        []Ident // lista de identificadores
	currentLevel int     // nivel actual
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		table:        []Ident{},
		currentLevel: -1,
	}
}

type StructIdent struct {
	token  antlr.TerminalNode
	typ    int
	nivel  int
	fields []int
}

func (t *SymbolTable) insertStructIdent(id antlr.TerminalNode, fields []int) {
	s := &StructIdent{
		token:  id,
		typ:    50,
		nivel:  t.currentLevel,
		fields: fields,
	}
	t.table = append([]Ident{s}, t.table...)

}

func (t *SymbolTable) insert(id antlr.TerminalNode, tp int) {
	v := &VarIdent{
		tok:        id,
		typ:        tp,
		level:      t.currentLevel,
		isConstant: false,
	}
	t.table = append([]Ident{v}, t.table...) // se inserta al inicio
}

func (t *SymbolTable) insertWithParams(id antlr.TerminalNode, tipo int, p []int) {
	m := &MethodIdent{
		tok:    id,
		typ:    tipo,
		level:  t.currentLevel,
		params: p,
	}
	t.table = append([]Ident{m}, t.table...) // se inserta al inicio
}

func (t *SymbolTable) search(name string) Ident {
	for _, id := range t.table {
		if v, ok := id.(VarIdent); ok {
			if v.tok.GetText() == name {
				return v
			}
		} else if m, ok := id.(MethodIdent); ok {
			if m.tok.GetText() == name {
				return m
			}
		}

	}
	return nil
}

func (t *SymbolTable) searchMethodInCurrentLevel() Ident {
	var temp Ident
	tempLevel := t.currentLevel // se guarda el nivel actual para comparar
	for _, id := range t.table {
		if m, ok := id.(*MethodIdent); ok {
			if m.level == tempLevel {
				return m
			}
		}
	}
	return temp
}

func (t *SymbolTable) searchInCurrentLevel(name string) Ident {
	var temp Ident
	tempLevel := t.currentLevel // se guarda el nivel actual para comparar
	for _, id := range t.table {
		if v, ok := id.(*VarIdent); ok {
			if v.tok.GetText() == name && v.level == tempLevel {
				return v
			}
		} else if m, ok := id.(*MethodIdent); ok {
			if m.tok.GetText() == name && m.level == tempLevel {
				return m
			}
		} else if s, ok := id.(*StructIdent); ok {
			if s.token.GetText() == name && s.nivel == tempLevel {
				return s
			}
		}
	}
	return temp
}

func (t *SymbolTable) openScope() { // abre un nuevo nivel en la tabla
	t.currentLevel++
}

func (t *SymbolTable) closeScope() {
	var newTabla []Ident

	for _, ident := range t.table {
		switch ident.(type) {
		case *VarIdent:
			if ident.(*VarIdent).level != t.currentLevel {
				newTabla = append(newTabla, ident) // se copian los identificadores
			}
		case *MethodIdent:
			if ident.(*MethodIdent).level != t.currentLevel {
				newTabla = append(newTabla, ident) // se copian los identificadores
			}
		}

		t.table = newTabla // se actualiza la tabla
		t.currentLevel--   // se disminuye el nivel
	}
}

func (t *SymbolTable) printTable() {
	fmt.Println("---------- START TABLE ----------")
	for _, ident := range t.table {
		if v, ok := ident.(*VarIdent); ok {
			fmt.Printf("NAME: %s - %d - %d\n", v.tok, v.level, v.typ)
		} else if m, ok := ident.(*MethodIdent); ok {
			fmt.Printf("NAME: %s - %d - %d\n", m.tok, m.level, m.typ)
		}
	}
	fmt.Println("---------- FIN TABLA ----------")
}
