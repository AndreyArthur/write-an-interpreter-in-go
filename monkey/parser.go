package monkey

import (
	"strconv"
)

const (
	PRECEDENCE_LOWEST = iota
	PRECEDENCE_EQUALS
	PRECEDENCE_LESS_GREATER
	PRECEDENCE_SUM
	PRECEDENCE_PRODUCT
	PRECEDENCE_PREFIX
)

var precedences = map[TokenType]int{
	TOKEN_EQUALS:       PRECEDENCE_EQUALS,
	TOKEN_NOT_EQUALS:   PRECEDENCE_EQUALS,
	TOKEN_LESS_THAN:    PRECEDENCE_LESS_GREATER,
	TOKEN_GREATER_THAN: PRECEDENCE_LESS_GREATER,
	TOKEN_PLUS:         PRECEDENCE_SUM,
	TOKEN_MINUS:        PRECEDENCE_SUM,
	TOKEN_ASTERISK:     PRECEDENCE_PRODUCT,
	TOKEN_SLASH:        PRECEDENCE_PRODUCT,
}

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
		prefixExpression.Right = parser.parseExpression(PRECEDENCE_PREFIX)
	case TOKEN_BANG:
		prefixExpression.Operator = parser.current.Literal
		parser.advance()
		prefixExpression.Right = parser.parseExpression(PRECEDENCE_PREFIX)
	default:
		// TODO: handle errors
		return nil
	}

	return prefixExpression

}

func (parser *Parser) parseInfixExpression(left AstExpression) AstExpression {
	infixExpression := &AstInfixExpression{
		Token:    parser.current,
		Left:     left,
		Operator: parser.current.Literal,
	}

	precedence := precedences[parser.current.Type]
	parser.advance()
	infixExpression.Right = parser.parseExpression(precedence)

	return infixExpression
}

func (parser *Parser) parseExpression(precedence int) AstExpression {
	var left AstExpression

	switch parser.current.Type {
	case TOKEN_INTEGER:
		left = parser.parseIntegerLiteral()
	case TOKEN_TRUE, TOKEN_FALSE:
		left = parser.parseBooleanLiteral()
	case TOKEN_MINUS, TOKEN_BANG:
		left = parser.parsePrefixExpression()
	default:
		// TODO: handle errors
		return nil
	}

	for parser.current.Type != TOKEN_SEMICOLON &&
		parser.current.Type != TOKEN_EOF &&
		precedence < precedences[parser.current.Type] {
		left = parser.parseInfixExpression(left)
	}

	return left
}

func (parser *Parser) parseExpressionStatement() AstStatement {
	expressionStatement := &AstExpressionStatement{Token: parser.current}

	expressionStatement.Expression = parser.parseExpression(PRECEDENCE_LOWEST)

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
