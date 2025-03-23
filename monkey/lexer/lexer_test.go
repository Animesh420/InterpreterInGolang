package lexer

import (
	"fmt"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
	}

	fmt.Printf(input)
	fmt.Printf(tests)
}
