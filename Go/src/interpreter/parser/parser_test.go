package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y= 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Program.Statements does not contain 3 statements. got = %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, nameIdentifier string) bool {
	fmt.Printf("statement  got = %#v", stmt)
	if stmt.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral not let. got = %q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement not *ast.LetStatement.got =%T", stmt)
		return false
	}
	if letStmt.Name.Value != nameIdentifier {
		t.Errorf("letStmt.Name.Value not '%s' . got = %s", nameIdentifier, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != nameIdentifier {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s' . got = %s", nameIdentifier, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}
