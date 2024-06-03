package test

import (
	"monkey/monkey"
	"testing"
)

type parserHelpers struct{}

func (*parserHelpers) expectLetStatement(
	t *testing.T,
	statement monkey.AstStatement,
	name string,
) *monkey.AstLetStatement {
	// TODO: We're not verfying the value

	letStatement, ok := statement.(*monkey.AstLetStatement)
	if !ok {
		t.Fatal("Given statement is not a let statement.")
		return nil
	}

	if letStatement.TokenLiteral() != "let" {
		t.Fatalf(
			"Expected token literal to be \"let\", got %q.",
			letStatement.TokenLiteral(),
		)
		return nil
	}

	if letStatement.Identifier.String() != name {
		t.Fatalf(
			"Expected let identifier name to be %q, got %q.",
			name,
			letStatement.Value.String(),
		)
		return nil
	}

	return letStatement
}

func (*parserHelpers) expectReturnStatement(
	t *testing.T,
	statement monkey.AstStatement,
) *monkey.AstReturnStatement {
	// TODO: We're not verfying the value

	returnStatement, ok := statement.(*monkey.AstReturnStatement)
	if !ok {
		t.Fatal("Given statement is not a return statement.")
		return nil
	}

	if returnStatement.TokenLiteral() != "return" {
		t.Fatalf(
			"Expected token literal to be \"return\", got %q.",
			returnStatement.TokenLiteral(),
		)
		return nil
	}

	return returnStatement
}

func (*parserHelpers) expectExpressionStatement(
	t *testing.T,
	statement monkey.AstStatement,
) *monkey.AstExpressionStatement {
	expressionStatement, ok := statement.(*monkey.AstExpressionStatement)
	if !ok {
		t.Fatal("Given statement is not an expression statement.")
		return nil
	}

	return expressionStatement
}

func (*parserHelpers) expectIntegerLiteral(
	t *testing.T,
	expression monkey.AstExpression,
	value int64,
) *monkey.AstIntegerLiteral {
	integerLiteral, ok := expression.(*monkey.AstIntegerLiteral)
	if !ok {
		t.Fatal("Given expression is not an integer literal.")
		return nil
	}

	if integerLiteral.Token.Type != monkey.TOKEN_INTEGER {
		t.Fatalf("Expected integer token type to be %s, got %s.",
			monkey.GetTokenTypeString(monkey.TOKEN_INTEGER),
			monkey.GetTokenTypeString(integerLiteral.Token.Type),
		)
		return nil
	}

	if integerLiteral.Value != value {
		t.Fatalf("Expected integer literal value to be %d, got %d.",
			value,
			integerLiteral.Value,
		)
		return nil
	}

	return integerLiteral
}

func TestLetStatements(t *testing.T) {
	input := `let a = 5;
let b = true;
let foo = 10;
`
	lexer := monkey.NewLexer(input)
	parser := monkey.NewParser(lexer)

	compound := parser.Parse()

	if len(compound.Statements) < 3 {
		t.Fatalf("Expected 3 statements, got %d.", len(compound.Statements))
	}

	expectations := []struct {
		identifier string
	}{
		{"a"},
		{"b"},
		{"foo"},
	}

	helpers := &parserHelpers{}

	for index, expectation := range expectations {
		if helpers.expectLetStatement(
			t,
			compound.Statements[index],
			expectation.identifier,
		) == nil {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `return 5;
return true;
return 10;
`
	lexer := monkey.NewLexer(input)
	parser := monkey.NewParser(lexer)

	compound := parser.Parse()

	if len(compound.Statements) < 3 {
		t.Fatalf("Expected 3 statements, got %d.", len(compound.Statements))
	}

	helpers := &parserHelpers{}

	for _, statement := range compound.Statements {
		if helpers.expectReturnStatement(
			t,
			statement,
		) == nil {
			return
		}
	}
}

func TestIntegerLiterals(t *testing.T) {
	input := `5;
20;
10;
`
	lexer := monkey.NewLexer(input)
	parser := monkey.NewParser(lexer)

	compound := parser.Parse()

	if len(compound.Statements) < 3 {
		t.Fatalf("Expected 3 statements, got %d.", len(compound.Statements))
	}

	expectations := []struct {
		value int64
	}{
		{5},
		{20},
		{10},
	}

	helpers := &parserHelpers{}

	for index, expectation := range expectations {
		expressionStatement := helpers.expectExpressionStatement(
			t,
			compound.Statements[index],
		)
		if expressionStatement == nil {
			return
		}

		integerLiteral := helpers.expectIntegerLiteral(
			t,
			expressionStatement.Expression,
			expectation.value,
		)
		if integerLiteral == nil {
			return
		}
	}
}
