package token

type system int
type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	ILLIGAL = "ILLIGAL"
	EOF     = "EOF"

	// identifyers and literals
	IDENT   = "IDENTIFIER"
	NEWLINE = "NEWLINE"
	INT     = "INT"
	FLOAT   = "FLOAT"
	STRING  = "STRING"

	PROJECT  = "PROJECT"
	MARKER   = "MARKER"
	TRACK    = "TRACK"
	NAME     = "NAME"
	ITEM     = "ITEM"
	SOURCE   = "SOURCE"
	POSITION = "POSITION"
	LENGTH   = "LENGTH"
	FILE     = "FILE"
	TRACKID  = "TRACKID"
	CURSOR   = "CURSOR"

	LT    = "LESSTHEN"
	GT    = "GREATERTHEN"
	MINUS = "MINUS"
	EQ    = "EQUAL"
	COL   = "COLON"
	QOTE  = "QOTE"
	AMP   = "AMPERSAND"
	LCUR  = "LEFTCURLY"
	RCUR  = "RIGHTCURLY"
	PIPE  = "PIPE"

	LINUX system = iota
	WIN
	UNKNOWN
)

var keywords = map[string]TokenType{
	"MARKER":   MARKER,
	"TRACK":    TRACK,
	"NAME":     NAME,
	"ITEM":     ITEM,
	"SOURCE":   SOURCE,
	"POSITION": POSITION,
	"LENGTH":   LENGTH,
	"FILE":     FILE,
	"TRACKID":  TRACKID,
	"CURSOR":   CURSOR,
}

func LookupIdent(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	}

	return IDENT
}
