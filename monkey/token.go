package monkey

const (
	TOKEN_ILLEGAL = iota
	TOKEN_EOF
	TOKEN_IDENTIFIER
	TOKEN_ASSIGNMENT
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_BANG
	TOKEN_ASTERISK
	TOKEN_SLASH
	TOKEN_LESS_THAN
	TOKEN_GREATER_THAN
	TOKEN_EQUALS
	TOKEN_NOT_EQUALS
	TOKEN_COMMA
	TOKEN_SEMICOLON
	TOKEN_OPEN_PAREN
	TOKEN_CLOSE_PAREN
	TOKEN_OPEN_BRACE
	TOKEN_CLOSE_BRACE
	TOKEN_INTEGER
	TOKEN_FUNCTION
	TOKEN_LET
	TOKEN_TRUE
	TOKEN_FALSE
	TOKEN_IF
	TOKEN_ELSE
	TOKEN_RETURN
)

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

func GetTokenTypeString(tokenType TokenType) string {
	types := map[TokenType]string{
		TOKEN_ILLEGAL:      "Illegal",
		TOKEN_EOF:          "Eof",
		TOKEN_IDENTIFIER:   "Identifier",
		TOKEN_ASSIGNMENT:   "Assignment",
		TOKEN_PLUS:         "Plus",
		TOKEN_MINUS:        "Minus",
		TOKEN_BANG:         "Bang",
		TOKEN_ASTERISK:     "Asterisk",
		TOKEN_SLASH:        "Slash",
		TOKEN_LESS_THAN:    "Less Than",
		TOKEN_GREATER_THAN: "Greater Than",
		TOKEN_EQUALS:       "Equals",
		TOKEN_NOT_EQUALS:   "Not Equals",
		TOKEN_COMMA:        "Comma",
		TOKEN_SEMICOLON:    "Semicolon",
		TOKEN_OPEN_PAREN:   "Open Paren",
		TOKEN_CLOSE_PAREN:  "Close Paren",
		TOKEN_OPEN_BRACE:   "Open Brace",
		TOKEN_CLOSE_BRACE:  "Close Brace",
		TOKEN_INTEGER:      "Integer",
		TOKEN_FUNCTION:     "Function",
		TOKEN_LET:          "Let",
		TOKEN_TRUE:         "True",
		TOKEN_FALSE:        "False",
		TOKEN_IF:           "If",
		TOKEN_ELSE:         "Else",
		TOKEN_RETURN:       "Return",
	}
	return types[tokenType]
}

func NewToken(tokenType TokenType, literal string) *Token {
	return &Token{Type: tokenType, Literal: literal}
}
