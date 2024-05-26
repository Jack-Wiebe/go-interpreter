package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	BANG = "!"

	SLASH = "/"
	STAR  = "*"

	COMMA     = ","
	SEMICOLON = ";"

	GT     = ">"
	LT     = "<"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION  = "FUNCTION"
	RETURN    = "RETURN"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	EQUALS    = "EQUALS"
	NOT_EQUAL = "NOTEQUAL"
	IF        = "IF"
	ELSE      = "ELSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"return": RETURN,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"==":     EQUALS,
	"!=":     NOT_EQUAL,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
