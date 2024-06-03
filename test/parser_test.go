package test

import (
	"monkey/monkey"
	"testing"
)

type parserHelpers struct{}

func (*parserHelpers) letStatementIsOk(
	t *testing.T,
	statement monkey.AstStatement,
	name string,
) bool {
	letStatement, ok := statement.(*monkey.AstLetStatement)
	if !ok {
		t.Fatal("Given statement is not a let statement.")
		return false
	}

	if letStatement.TokenLiteral() != "let" {
		t.Fatalf(
			"Expected token literal to be \"let\", got %q.",
			letStatement.TokenLiteral(),
		)
		return false
	}

	if letStatement.Identifier.String() != name {
		t.Fatalf(
			"Expected let identifier name to be %q, got %q.",
			name,
			letStatement.Value.String(),
		)
		return false
	}

	return true
}

func (*parserHelpers) returnStatementIsOk(
	t *testing.T,
	statement monkey.AstStatement,
) bool {
	returnStatement, ok := statement.(*monkey.AstReturnStatement)
	if !ok {
		t.Fatal("Given statement is not a return statement.")
		return false
	}

	if returnStatement.TokenLiteral() != "return" {
		t.Fatalf(
			"Expected token literal to be \"return\", got %q.",
			returnStatement.TokenLiteral(),
		)
		return false
	}

	return true
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
		if !helpers.letStatementIsOk(
			t,
			compound.Statements[index],
			expectation.identifier,
		) {
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
		if !helpers.returnStatementIsOk(
			t,
			statement,
		) {
			return
		}
	}
}
