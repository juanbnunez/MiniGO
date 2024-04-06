package Checker

import "fmt"

type Token struct {
	Text string
	// Type int  // If needed
}

type Ident struct {
	Tok   Token
	Type  int // this may change to a more structured type
	Level int
	Value int
}

type SymbolTable struct {
	Table        []*Ident
	CurrentLevel int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Table:        make([]*Ident, 0),
		CurrentLevel: -1,
	}
}

func (st *SymbolTable) Insert(id Token, typ int) {
	i := &Ident{
		Tok:   id,
		Type:  typ,
		Level: st.CurrentLevel,
		Value: 0,
	}
	st.Table = append([]*Ident{i}, st.Table...)
}

func (st *SymbolTable) Search(name string) *Ident {
	for _, id := range st.Table {
		if id.Tok.Text == name {
			return id
		}
	}
	return nil
}

func (st *SymbolTable) SearchOnCurrentScope(name string) *Ident {
	var temp *Ident
	for _, id := range st.Table {
		if id.Tok.Text == name && id.Level == st.CurrentLevel {
			temp = id
		}
	}
	return temp
}

func (st *SymbolTable) OpenScope() {
	st.CurrentLevel++
}

func (st *SymbolTable) CloseScope() {
	var newTable []*Ident
	for _, ident := range st.Table {
		if ident.Level != st.CurrentLevel {
			newTable = append(newTable, ident)
		}
	}
	st.Table = newTable
	st.CurrentLevel--
}

func (st *SymbolTable) Print() {
	fmt.Println("----- SYMBOL TABLE START ------")
	for _, ident := range st.Table {
		fmt.Printf("Name: %s - %d - %d\n", ident.Tok.Text, ident.Level, ident.Type)
	}
	fmt.Println("----- SYMBOL TABLE END ------")
}
