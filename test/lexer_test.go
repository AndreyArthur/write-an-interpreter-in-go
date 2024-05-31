package test

import (
	"monkey/monkey"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		tokenType monkey.TokenType
		literal   string
	}{
		{monkey.TOKEN_LET, "let"},
		{monkey.TOKEN_IDENTIFIER, "five"},
		{monkey.TOKEN_ASSIGNMENT, "="},
		{monkey.TOKEN_INTEGER, "5"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_LET, "let"},
		{monkey.TOKEN_IDENTIFIER, "ten"},
		{monkey.TOKEN_ASSIGNMENT, "="},
		{monkey.TOKEN_INTEGER, "10"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_LET, "let"},
		{monkey.TOKEN_IDENTIFIER, "add"},
		{monkey.TOKEN_ASSIGNMENT, "="},
		{monkey.TOKEN_FUNCTION, "fn"},
		{monkey.TOKEN_OPEN_PAREN, "("},
		{monkey.TOKEN_IDENTIFIER, "x"},
		{monkey.TOKEN_COMMA, ","},
		{monkey.TOKEN_IDENTIFIER, "y"},
		{monkey.TOKEN_CLOSE_PAREN, ")"},
		{monkey.TOKEN_OPEN_BRACE, "{"},
		{monkey.TOKEN_IDENTIFIER, "x"},
		{monkey.TOKEN_PLUS, "+"},
		{monkey.TOKEN_IDENTIFIER, "y"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_CLOSE_BRACE, "}"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_LET, "let"},
		{monkey.TOKEN_IDENTIFIER, "result"},
		{monkey.TOKEN_ASSIGNMENT, "="},
		{monkey.TOKEN_IDENTIFIER, "add"},
		{monkey.TOKEN_OPEN_PAREN, "("},
		{monkey.TOKEN_IDENTIFIER, "five"},
		{monkey.TOKEN_COMMA, ","},
		{monkey.TOKEN_IDENTIFIER, "ten"},
		{monkey.TOKEN_CLOSE_PAREN, ")"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_BANG, "!"},
		{monkey.TOKEN_MINUS, "-"},
		{monkey.TOKEN_SLASH, "/"},
		{monkey.TOKEN_ASTERISK, "*"},
		{monkey.TOKEN_INTEGER, "5"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_INTEGER, "5"},
		{monkey.TOKEN_LESS_THAN, "<"},
		{monkey.TOKEN_INTEGER, "10"},
		{monkey.TOKEN_GREATER_THAN, ">"},
		{monkey.TOKEN_INTEGER, "5"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_IF, "if"},
		{monkey.TOKEN_OPEN_PAREN, "("},
		{monkey.TOKEN_INTEGER, "5"},
		{monkey.TOKEN_LESS_THAN, "<"},
		{monkey.TOKEN_INTEGER, "10"},
		{monkey.TOKEN_CLOSE_PAREN, ")"},
		{monkey.TOKEN_OPEN_BRACE, "{"},
		{monkey.TOKEN_RETURN, "return"},
		{monkey.TOKEN_TRUE, "true"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_CLOSE_BRACE, "}"},
		{monkey.TOKEN_ELSE, "else"},
		{monkey.TOKEN_OPEN_BRACE, "{"},
		{monkey.TOKEN_RETURN, "return"},
		{monkey.TOKEN_FALSE, "false"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_CLOSE_BRACE, "}"},
		{monkey.TOKEN_INTEGER, "10"},
		{monkey.TOKEN_EQUALS, "=="},
		{monkey.TOKEN_INTEGER, "10"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_INTEGER, "10"},
		{monkey.TOKEN_NOT_EQUALS, "!="},
		{monkey.TOKEN_INTEGER, "9"},
		{monkey.TOKEN_SEMICOLON, ";"},
		{monkey.TOKEN_EOF, "\x00"},
	}

	lexer := monkey.NewLexer(input)

	for index, expected := range tests {
		token := lexer.Next()

		if token.Type != expected.tokenType {
			t.Fatalf(
				"tests[%d] - TokenType wrong. expect=%q, got=%q,",
				index,
				monkey.GetTokenTypeString(expected.tokenType),
				monkey.GetTokenTypeString(token.Type),
			)
		}

		if token.Literal != expected.literal {
			t.Fatalf(
				"tests[%d] - TokenLiteral wrong. expect=%q, got=%q,",
				index,
				expected.literal,
				token.Literal,
			)
		}
	}
}
