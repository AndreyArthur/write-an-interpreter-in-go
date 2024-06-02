package monkey

type Parser struct {
	tokens   []*Token
	position int
	current  *Token
}

func NewParser(lexer *Lexer) *Parser {
	tokens := []*Token{}

	current := lexer.Next()
	tokens = append(tokens, current)
	for current.Type != TOKEN_EOF {
		current := lexer.Next()
		tokens = append(tokens, current)
		// Don't know why but the condition on the top is never satisfied (???)
		if current.Type == TOKEN_EOF {
			break
		}
	}

	parser := &Parser{
		tokens:   tokens,
		position: 0,
	}
	parser.current = parser.tokens[parser.position]

	return parser
}

func (parser *Parser) advance() {
	if parser.current.Type == TOKEN_EOF {
		return
	}

	parser.position += 1
	parser.current = parser.tokens[parser.position]
}

func (parser *Parser) parseLetStatement() *AstLetStatement {
	letStatement := &AstLetStatement{Token: parser.current}
	parser.advance()

	identifier := &AstIdentifier{
		Token: parser.current,
		Value: parser.current.Literal,
	}
	letStatement.Identifier = identifier

	// TODO: parse the expression after assignment
	for parser.current.Type != TOKEN_SEMICOLON &&
		parser.current.Type != TOKEN_EOF {
		parser.advance()
	}

	return letStatement
}

func (parser *Parser) parseStatement() AstStatement {
	switch parser.current.Type {
	case TOKEN_LET:
		return parser.parseLetStatement()
	default:
		// TODO: implement other cases and handle errors
		return nil
	}
}

func (parser *Parser) Parse() *AstCompound {
	compound := &AstCompound{Statements: []AstStatement{}}

	for parser.current.Type != TOKEN_EOF {
		statement := parser.parseStatement()
		compound.Statements = append(compound.Statements, statement)
		parser.advance()
	}

	return compound
}
