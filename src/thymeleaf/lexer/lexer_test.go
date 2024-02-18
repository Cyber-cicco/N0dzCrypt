package lexer

import (
	"fr/nzc/thymeleaf/token"
	"testing"
)

func TestNextToken(t *testing.T) {
    input := `<button type="submit" class="bg-primary">submit</button>`
    l := New(input)
    tests := []struct {
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.TAG_OPENER, "<"},
        {token.WORD, "button"},
        {token.WHITESPACE, " "},
        {token.WORD, "type"},
        {token.EQUAL_SIGN, "="},
        {token.QUOTATION_MARK, "\""},
        {token.WORD, "submit"},
        {token.QUOTATION_MARK, "\""},
        {token.WHITESPACE, " "},
        {token.WORD, "class"},
        {token.EQUAL_SIGN, "="},
        {token.QUOTATION_MARK, "\""},
        {token.WORD, "bg-primary"},
        {token.QUOTATION_MARK, "\""},
        {token.TAG_CLOSER, ">"},
        {token.WORD, "submit"},
        {token.TAG_OPENER, "<"},
        {token.SLASH, "/"},
        {token.WORD, "button"},
        {token.TAG_CLOSER, ">"},
    }
    for i, tt := range tests {
        tok := l.NextToken()
        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
            i, tt.expectedType, tok.Type)
        }
        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
            i, tt.expectedLiteral, tok.Literal)
        }
    }
}
