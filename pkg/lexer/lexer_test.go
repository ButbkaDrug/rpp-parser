package lexer

import (
	"log"
	"rpp-parser/pkg/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `<REAPER_PROJECT 0.1 "7.07/linux-x86_64" 1715322766
  AUTHOR PROJECT_AUTHOR
  CURSOR 1234
  TIMEMODE 1 5 -1 30 0 0 -1
  <RECORD_CFG
    ZXZhdxgAAA==
  >
  <RENDER_CFG
    bDNwbaAAAAAAAAAAAgAAAP////8EAAAAoAAAAAAAAAA=
  >
  MARKER 1 0 "" 0 0 1 B {5B24D1B6-3896-B803-0145-386DA9EF6B5B} 0
  MARKER 2 3.33333333333333 Intro 1 0 1 B {35835D1C-B5DB-C77F-84E2-D46EE9131075} 0
  MARKER 2 3.33333333333333 "Pre Chorus" 1 0 1 B {35835D1C-B5DB-C77F-84E2-D46EE9131075} 0
    AUXRECV 0 0 1 0 0 0 0 2 0 -1:U 0 -1 ''
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LT, "<"},
		{token.IDENT, "REAPER_PROJECT"},
		{token.FLOAT, "0.1"},
		{token.STRING, "7.07/linux-x86_64"},
		{token.INT, "1715322766"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "AUTHOR"},
		{token.STRING, "PROJECT_AUTHOR"},
		{token.NEWLINE, "\n"},
		{token.CURSOR, "CURSOR"},
		{token.INT, "1234"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "TIMEMODE"},
		{token.INT, "1"},
		{token.INT, "5"},
		{token.INT, "-1"},
		{token.INT, "30"},
		{token.INT, "0"},
		{token.INT, "0"},
		{token.INT, "-1"},
		{token.NEWLINE, "\n"},
		{token.LT, "<"},
		{token.IDENT, "RECORD_CFG"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "ZXZhdxgAAA=="},
		{token.NEWLINE, "\n"},
		{token.GT, ">"},
		{token.NEWLINE, "\n"},
		{token.LT, "<"},
		{token.IDENT, "RENDER_CFG"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "bDNwbaAAAAAAAAAAAgAAAP////8EAAAAoAAAAAAAAAA="},
		{token.NEWLINE, "\n"},
		{token.GT, ">"},
		{token.NEWLINE, "\n"},
		{token.MARKER, "MARKER"},
		{token.INT, "1"},
		{token.INT, "0"},
		{token.STRING, ""},
		{token.INT, "0"},
		{token.INT, "0"},
		{token.INT, "1"},
		{token.STRING, "B"},
		{token.STRING, "5B24D1B6-3896-B803-0145-386DA9EF6B5B"},
		{token.INT, "0"},
		{token.NEWLINE, "\n"},
		{token.MARKER, "MARKER"},
		{token.INT, "2"},
		{token.FLOAT, "3.33333333333333"},
		{token.STRING, "Intro"},
		{token.INT, "1"},
		{token.INT, "0"},
		{token.INT, "1"},
		{token.STRING, "B"},
		{token.STRING, "35835D1C-B5DB-C77F-84E2-D46EE9131075"},
		{token.INT, "0"},
		{token.NEWLINE, "\n"},
		{token.MARKER, "MARKER"},
		{token.INT, "2"},
		{token.FLOAT, "3.33333333333333"},
		{token.STRING, "Pre Chorus"},
		{token.INT, "1"},
		{token.INT, "0"},
		{token.INT, "1"},
		{token.STRING, "B"},
		{token.STRING, "35835D1C-B5DB-C77F-84E2-D46EE9131075"},
		{token.INT, "0"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "AUXRECV"},
		{token.INT, "0"},
		{token.INT, "0"},
		{token.INT, "1"},
		{token.INT, "0"},
		{token.INT, "0"},
		{token.INT, "0"},
		{token.INT, "0"},
		{token.INT, "2"},
		{token.INT, "0"},
		{token.INT, "-1"},
		{token.COL, ":"},
		{token.STRING, "U"},
		{token.INT, "0"},
		{token.INT, "-1"},
		{token.STRING, ""},
		{token.NEWLINE, "\n"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			log.Fatalf("test[%d] - wrong tokentype. expected %q got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Value != tt.expectedLiteral {
			log.Fatalf("test[%d] - wrog literal. expected %q got %q", i, tt.expectedLiteral, tok.Value)
		}
	}

}
