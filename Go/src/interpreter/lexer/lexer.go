package lexer

import (
	"interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = newToken(token.NOT_EQ, literal)
		} else {
			tok = newToken(token.BANG, string(l.ch))
		}
	case '<':
		tok = newToken(token.LT, string(l.ch))
	case '>':
		tok = newToken(token.GT, string(l.ch))
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = newToken(token.EQ, literal)
		} else {
			tok = newToken(token.ASSIGN, string(l.ch))
		}
	case '/':
		tok = newToken(token.SLASH, string(l.ch))
	case '*':
		tok = newToken(token.ASTERISK, string(l.ch))
	case '+':
		tok = newToken(token.PLUS, string(l.ch))
	case '-':
		tok = newToken(token.MINUS, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.ch))
	case '{':
		tok = newToken(token.LBRACE, string(l.ch))
	case '}':
		tok = newToken(token.RBRACE, string(l.ch))
	case '(':
		tok = newToken(token.LPAREN, string(l.ch))
	case ')':
		tok = newToken(token.RPAREN, string(l.ch))
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.ch))
		}
	}
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || ch == '_'
}

func isDigit(ch byte) bool {

	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\b' || l.ch == '\r' {
		l.readChar()
	}

}
