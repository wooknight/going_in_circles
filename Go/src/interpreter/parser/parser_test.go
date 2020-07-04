package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y  = 10;
	let foobar=   838383;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
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

func testReturnStatement(t *testing.T, stmt ast.Statement, nameIdentifier string) {
	input := `
	return 5;
	return x;
	return (x+5);
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got = %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got = %T", stmt)
			continue
		}
		if retStmt.TokenLiteral() != "return" {
			t.Errorf("statement.TokenLiteral not return. got = %q", retStmt.TokenLiteral())
			continue
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, nameIdentifier string) bool {
	// fmt.Printf("statement  got = %#v", stmt)
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error :%q", msg)
	}
	t.FailNow()
}
func TestIdentifierExpression(t *testing.T) {
	input := "foobar"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements.got = %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement.got = %T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier .got = %T", stmt.Expression)
	}
	if ident.Value != "foobar"{
		t.Errorf("ident.Value is not foobar. got = %s",ident.Value)
	}
	if ident.TokenLiteral()!="foobar"{
		t.Errorf("ident.TokenLiteral not foobar. got = %s",ident.TokenLiteral())
	}
}
