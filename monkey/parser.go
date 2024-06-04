package monkey

import (
	"strconv"
)

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
		current = lexer.Next()
		tokens = append(tokens, current)
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

func (parser *Parser) parseLetStatement() AstStatement {
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

	if parser.current.Type == TOKEN_SEMICOLON {
		parser.advance()
	}

	return letStatement
}

func (parser *Parser) parseReturnStatement() AstStatement {
	returnStatement := &AstReturnStatement{Token: parser.current}
	parser.advance()

	// TODO: parse the return value
	for parser.current.Type != TOKEN_SEMICOLON &&
		parser.current.Type != TOKEN_EOF {
		parser.advance()
	}

	if parser.current.Type == TOKEN_SEMICOLON {
		parser.advance()
	}

	return returnStatement
}

func (parser *Parser) parseIntegerLiteral() AstExpression {
	value, err := strconv.ParseInt(parser.current.Literal, 10, 64)
	if err != nil {
		return nil
	}

	integerLiteral := &AstIntegerLiteral{
		Token: parser.current,
		Value: value,
	}

	parser.advance()

	return integerLiteral
}

func (parser *Parser) parseBooleanLiteral() AstExpression {
	booleanLiteral := &AstBooleanLiteral{
		Token: parser.current,
		Value: parser.current.Type == TOKEN_TRUE,
	}

	parser.advance()

	return booleanLiteral
}

func (parser *Parser) parsePrefixExpression() AstExpression {
	prefixExpression := &AstPrefixExpression{Token: parser.current}

	switch parser.current.Type {
	case TOKEN_MINUS:
		prefixExpression.Operator = parser.current.Literal
		parser.advance()
		prefixExpression.Right = parser.parseExpression()
	case TOKEN_BANG:
		prefixExpression.Operator = parser.current.Literal
		parser.advance()
		prefixExpression.Right = parser.parseExpression()
	default:
		return nil
	}

	return prefixExpression

}

func (parser *Parser) parseExpression() AstExpression {
	switch parser.current.Type {
	case TOKEN_INTEGER:
		return parser.parseIntegerLiteral()
	case TOKEN_TRUE, TOKEN_FALSE:
		return parser.parseBooleanLiteral()
	default:
		return parser.parsePrefixExpression()
	}
}

func (parser *Parser) parseExpressionStatement() AstStatement {
	expressionStatement := &AstExpressionStatement{Token: parser.current}

	expressionStatement.Expression = parser.parseExpression()

	if parser.current.Type == TOKEN_SEMICOLON {
		parser.advance()
	}

	return expressionStatement
}

func (parser *Parser) parseStatement() AstStatement {
	switch parser.current.Type {
	case TOKEN_LET:
		return parser.parseLetStatement()
	case TOKEN_RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) Parse() *AstCompound {
	compound := &AstCompound{Statements: []AstStatement{}}

	for parser.current.Type != TOKEN_EOF {
		statement := parser.parseStatement()
		compound.Statements = append(compound.Statements, statement)
	}

	return compound
}
