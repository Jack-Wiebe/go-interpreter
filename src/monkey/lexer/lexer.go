package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input string
	position int
	readPosition int
	char byte 
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if(lexer.readPosition >= len(lexer.input)){
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.char){
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]

}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.char){
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) NextToken() (tok token.Token)  {
	//tok token.Token

	lexer.skipWitespace()

	switch lexer.char {
	case '=':
		if(lexer.peekChar() == '='){
			tok = makeTwoCharToken(lexer, token.EQUALS)
		}else{
			tok = newToken(token.ASSIGN, lexer.char)
		}
	case '!':
		if(lexer.peekChar() == '='){
			tok = makeTwoCharToken(lexer, token.NOT_EQUAL)
		}else{
			tok = newToken(token.BANG, lexer.char)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = newToken(token.LPAREN, lexer.char)
	case ')':
		tok = newToken(token.RPAREN, lexer.char)
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case '+':
		tok = newToken(token.PLUS, lexer.char)
	case '-':
		tok = newToken(token.MINUS, lexer.char)
	case '/':
		tok = newToken(token.SLASH, lexer.char)
	case '{':
		tok = newToken(token.LBRACE, lexer.char)
	case '}':
		tok = newToken(token.RBRACE, lexer.char)
	case '>':
		tok = newToken(token.GT, lexer.char)
	case '<':
		tok = newToken(token.LT, lexer.char)
	case '*':
		tok = newToken(token.STAR, lexer.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.char){
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok 
		} else if isDigit(lexer.char){
			tok.Literal = lexer.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.char)
		}
		
	}

	lexer.readChar()
	return tok
	
}

func (lexer *Lexer) skipWitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	} 
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func makeTwoCharToken(lexer *Lexer, tokenType token.TokenType) token.Token {
	lastChar := lexer.char
	lexer.readChar()
	literal := string(lastChar) + string(lexer.char)
	return token.Token{Type: tokenType, Literal: literal}
}