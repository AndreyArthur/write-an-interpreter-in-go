package monkey

type Lexer struct {
	content  string
	current  byte
	position int
}

func NewLexer(content string) *Lexer {
	lexer := &Lexer{
		content:  content,
		position: 0,
	}
	lexer.current = lexer.content[lexer.position]
	return lexer
}

func (lexer *Lexer) skipWhitespaces() {
	for lexer.current == ' ' ||
		lexer.current == '\t' ||
		lexer.current == '\n' ||
		lexer.current == '\r' {
		lexer.advance()
	}
}

func isAlphanumeric(character byte) bool {
	return (character >= 'a' && character <= 'z') ||
		(character >= 'A' && character <= 'Z') ||
		(character >= '0' && character <= '9') ||
		character == '_'
}

func isNumeric(character byte) bool {
	return character >= '0' && character <= '9'
}

func (lexer *Lexer) collectIdentifierOrKeyword() *Token {
	identifier := ""

	for isAlphanumeric(lexer.current) {
		identifier = identifier + string(lexer.current)
		lexer.advance()
	}

	switch identifier {
	case "fn":
		return NewToken(TOKEN_FUNCTION, identifier)
	case "let":
		return NewToken(TOKEN_LET, identifier)
	case "true":
		return NewToken(TOKEN_TRUE, identifier)
	case "false":
		return NewToken(TOKEN_FALSE, identifier)
	case "if":
		return NewToken(TOKEN_IF, identifier)
	case "else":
		return NewToken(TOKEN_ELSE, identifier)
	case "return":
		return NewToken(TOKEN_RETURN, identifier)
	default:
		return NewToken(TOKEN_IDENTIFIER, identifier)
	}
}

func (lexer *Lexer) collectIntegerLiteral() *Token {
	literal := ""

	for isNumeric(lexer.current) {
		literal = literal + string(lexer.current)
		lexer.advance()
	}

	return NewToken(TOKEN_INTEGER, literal)
}

func (lexer *Lexer) advance() {
	if lexer.position+1 >= len(lexer.content) {
		lexer.current = 0
	} else {
		lexer.position += 1
		lexer.current = lexer.content[lexer.position]
	}
}

func (lexer *Lexer) peek() byte {
	if lexer.position+1 >= len(lexer.content) {
		return 0
	}
	return lexer.content[lexer.position+1]
}

func (lexer *Lexer) Next() *Token {
	lexer.skipWhitespaces()

	switch lexer.current {
	case 0:
		return NewToken(TOKEN_EOF, string(lexer.current))
	case '=':
		current := string(lexer.current)
		lexer.advance()
		if lexer.current == '=' {
			token := NewToken(TOKEN_EQUALS, current+string(lexer.current))
			lexer.advance()
			return token
		}
		return NewToken(TOKEN_ASSIGNMENT, current)
	case '!':
		current := string(lexer.current)
		lexer.advance()
		if lexer.current == '=' {
			token := NewToken(TOKEN_NOT_EQUALS, current+string(lexer.current))
			lexer.advance()
			return token
		}
		return NewToken(TOKEN_BANG, current)
	case '+':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_PLUS, current)
	case '-':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_MINUS, current)
	case '*':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_ASTERISK, current)
	case '/':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_SLASH, current)
	case '<':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_LESS_THAN, current)
	case '>':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_GREATER_THAN, current)
	case ',':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_COMMA, current)
	case ';':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_SEMICOLON, current)
	case '(':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_OPEN_PAREN, current)
	case ')':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_CLOSE_PAREN, current)
	case '{':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_OPEN_BRACE, current)
	case '}':
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_CLOSE_BRACE, current)
	default:
		if isNumeric(lexer.current) {
			return lexer.collectIntegerLiteral()
		}
		if isAlphanumeric(lexer.current) {
			return lexer.collectIdentifierOrKeyword()
		}
		current := string(lexer.current)
		lexer.advance()
		return NewToken(TOKEN_ILLEGAL, current)
	}
}
