package monkey

import "bytes"

type AstNode interface {
	TokenLiteral() string
	String() string
}

type AstStatement interface {
	AstNode
	statement()
}

type AstExpression interface {
	AstNode
	expression()
}

type AstCompound struct {
	Statements []AstStatement
}

func (compound *AstCompound) TokenLiteral() string {
	if len(compound.Statements) > 0 {
		return compound.Statements[0].TokenLiteral()
	}
	return ""
}
func (compound *AstCompound) String() string {
	var out bytes.Buffer

	for _, statement := range compound.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}

type AstIdentifier struct {
	Token *Token // identifier name
	Value string
}

func (identifier *AstIdentifier) expression() {}
func (identifier *AstIdentifier) TokenLiteral() string {
	return identifier.Token.Literal
}
func (identifier *AstIdentifier) String() string {
	return identifier.TokenLiteral()
}

type AstLetStatement struct {
	Token      *Token // "let"
	Identifier *AstIdentifier
	Value      AstExpression
}

func (let *AstLetStatement) statement()           {}
func (let *AstLetStatement) TokenLiteral() string { return let.Token.Literal }
func (let *AstLetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(let.TokenLiteral() + " ")
	out.WriteString(let.Identifier.String())
	out.WriteString(" = ")
	out.WriteString(let.Value.String())
	out.WriteString(";")

	return out.String()
}

type AstReturnStatement struct {
	Token *Token // "return"
	Value AstExpression
}

func (returnStatement *AstReturnStatement) statement() {}
func (returnStatement *AstReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}
func (returnStatement *AstReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStatement.TokenLiteral() + " ")
	out.WriteString(returnStatement.Value.String())
	out.WriteString(";")

	return out.String()
}

type AstExpressionStatement struct {
	Token      *Token // first token
	Expression AstExpression
}

func (expression *AstExpressionStatement) statement() {}
func (expression *AstExpressionStatement) TokenLiteral() string {
	return expression.Expression.TokenLiteral()
}
func (expression *AstExpressionStatement) String() string {
	return expression.Expression.String() + ";"
}

type AstIntegerLiteral struct {
	Token *Token // the integer string
	Value int64
}

func (integer *AstIntegerLiteral) expression() {}
func (integer *AstIntegerLiteral) TokenLiteral() string {
	return integer.Token.Literal
}
func (integer *AstIntegerLiteral) String() string {
	return integer.TokenLiteral()
}

type AstBooleanLiteral struct {
	Token *Token // "true" or "false"
	Value bool
}

func (boolean *AstBooleanLiteral) expression() {}
func (boolean *AstBooleanLiteral) TokenLiteral() string {
	return boolean.Token.Literal
}
func (boolean *AstBooleanLiteral) String() string {
	return boolean.TokenLiteral()
}

type AstPrefixExpression struct {
	Token    *Token // Operator token
	Operator string
	Right    AstExpression
}

func (prefix *AstPrefixExpression) expression() {}
func (prefix *AstPrefixExpression) TokenLiteral() string {
	return prefix.Token.Literal
}
func (prefix *AstPrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(prefix.Operator)
	out.WriteString(prefix.Right.String())
	out.WriteString(")")

	return out.String()
}
