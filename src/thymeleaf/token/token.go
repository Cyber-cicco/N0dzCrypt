package token

type TokenType string

const (
    TAG_OPENER = "TAG_OPENER"
    TAG_CLOSER = "TAG_CLOSER"
    SLASH = "SLASH"
    EXCLAMATION_MARK = "EXCLAMATION_MARK"
    EQUAL_SIGN = "EQUAL_SIGN"
    MINUS_SIGN = "MINUS_SIGN"
    QUOTATION_MARK = "QUOTATION_MARK"
    DOULBE_DOT = "DOULBE_DOT"
    WORD = "WORD"
    WHITESPACE = "WHITESPACE"
    EOF = "EOF"
)

type HTMLToken struct {
    Type TokenType
    Literal string
}
