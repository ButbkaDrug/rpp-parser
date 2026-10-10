package lexer

import (
	"rpp-parser/pkg/token"
	"strings"
	"unicode"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
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

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) prevChar() byte {
	if l.position < 1 {
		return 0
	}

	position := l.position - 1

	for l.input[position] == ' ' && position > 0 {
		position--
	}

	return l.input[position]
}

func (l *Lexer) readIdentifyer() string {
	position := l.position

	for unicode.IsLetter(rune(l.ch)) || l.ch == '_' {
		l.readChar()
	}

	return l.input[position:l.position]
}

// Will read any chars until encounters stop char
func (l *Lexer) readAny(stop byte) string {
	position := l.position

	for l.ch != stop && l.ch != 0 && l.ch != '\n' && l.ch != '\r' {

		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readLine() string {
	position := l.position

	for l.ch != '\n' {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for unicode.IsNumber(rune(l.ch)) || l.ch == '.' || l.ch == '-' {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) isFloat(s string) bool {
	return strings.Index(s, ".") != -1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '"':
		l.readChar()
		tok = token.Token{Type: token.STRING, Value: l.readAny('"')}
	case '\'':
		l.readChar()
		tok = token.Token{Type: token.STRING, Value: l.readAny('\'')}
	case '/':
		tok = token.Token{Type: token.STRING, Value: l.readLine()}
	case '{':
		l.readChar()
		tok = token.Token{Type: token.STRING, Value: l.readAny('}')}
	case '}':
		tok = token.Token{Type: token.RCUR, Value: "}"}
	case '<':
		tok = token.Token{Type: token.LT, Value: "<"}
	case '>':
		tok = token.Token{Type: token.GT, Value: ">"}
	case ':':
		tok = token.Token{Type: token.COL, Value: ":"}
	case '|':
		tok = token.Token{Type: token.PIPE, Value: "|"}
		l.readChar()
		tok = token.Token{Type: token.STRING, Value: l.readLine()}
	case '\n':
		tok = token.Token{Type: token.NEWLINE, Value: "\n"}

	case 0:
		return token.Token{Type: token.EOF, Value: ""}
	default:
		if unicode.IsDigit(rune(l.ch)) || l.ch == '-' {
			n := l.readNumber()
			tok = token.Token{Value: n}

			if l.isFloat(n) {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}
			return tok
		} else {

			if c := l.prevChar(); c == '\n' || c == '<' {
				tok.Value = l.readAny(' ')
				tok.Type = token.LookupIdent(tok.Value)
				return tok
			}
			tok = token.Token{Type: token.STRING, Value: l.readAny(' ')}
			return tok

		}
	}

	l.readChar()

	return tok
}
