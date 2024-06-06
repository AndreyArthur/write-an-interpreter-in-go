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

type AstInfixExpression struct {
	Token    *Token // Operator token
	Left     AstExpression
	Operator string
	Right    AstExpression
}

func (infix *AstInfixExpression) expression() {}
func (infix *AstInfixExpression) TokenLiteral() string {
	return infix.Token.Literal
}
func (infix *AstInfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infix.Left.String())
	out.WriteString(" " + infix.Operator + " ")
	out.WriteString(infix.Right.String())
	out.WriteString(")")

	return out.String()
}

type AstFunctionCall struct {
	Token      *Token // the identifier token
	Identifier *AstIdentifier
	Arguments  []AstExpression
}

func (functionCall *AstFunctionCall) expression() {}
func (functionCall *AstFunctionCall) TokenLiteral() string {
	return functionCall.Token.Literal
}
func (functionCall *AstFunctionCall) String() string {
	var out bytes.Buffer

	out.WriteString(functionCall.TokenLiteral())
	out.WriteString("(")
	for index, argument := range functionCall.Arguments {
		out.WriteString(argument.String())
		if index < len(functionCall.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")

	return out.String()
}

type AstFunctionDefinition struct {
	Token  *Token // "fn"
	Params []*AstIdentifier
	Body   *AstCompound
}

func (functionDefinition *AstFunctionDefinition) expression() {}
func (functionDefinition *AstFunctionDefinition) TokenLiteral() string {
	return functionDefinition.Token.Literal
}
func (functionDefinition *AstFunctionDefinition) String() string {
	var out bytes.Buffer

	out.WriteString(functionDefinition.TokenLiteral() + " (")
	for index, param := range functionDefinition.Params {
		out.WriteString(param.Value)
		if index < len(functionDefinition.Params)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(") { ")
	out.WriteString(functionDefinition.Body.String())
	out.WriteString(" }")

	return out.String()
}
