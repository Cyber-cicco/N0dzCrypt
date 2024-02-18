package lexer

import (
	"fr/nzc/thymeleaf/token"
	"unicode"
)

type HTMLLexer struct {
    input string
    position int //current position in input
    readPosition int // current reading position in input 
    ch byte // current char under examination
}

func New(input string) *HTMLLexer {
    l := & HTMLLexer{input : input}
    l.readChar()
    return l
}

/**
l : mutable
l.ch
l.position
l.readPosition
*/
func (l *HTMLLexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

/**
l : immutable
*/
func (l *HTMLLexer) nextCharDontBreakWord() bool {
    return unicode.IsLetter(rune(l.input[l.readPosition])) || unicode.IsDigit(rune(l.input[l.readPosition])) || l.input[l.readPosition] == '-' || l.input[l.readPosition] == '_'
}

/**
l : mutable
l.position
l.readPosition
*/
func (l *HTMLLexer) resetPos() {
    l.position = l.readPosition
    l.readPosition++
}

/**
l : mutable
l.ch
l.position
l.readPosition
*/
func (l *HTMLLexer) NextToken() token.HTMLToken {
    var tok token.HTMLToken
    switch l.ch {
    case '<' :
        tok = newToken(token.TAG_OPENER, l.ch) 
    case '>' :
        tok = newToken(token.TAG_CLOSER, l.ch)
    case '=' :
        tok = newToken(token.EQUAL_SIGN, l.ch)
    case '"' :
        tok = newToken(token.QUOTATION_MARK, l.ch)
    case '-' :
        tok = newToken(token.MINUS_SIGN, l.ch)
    case '!' :
        tok = newToken(token.EXCLAMATION_MARK, l.ch)
    case '/' :
        tok = newToken(token.SLASH, l.ch)
    case ':' :
        tok = newToken(token.DOULBE_DOT, l.ch)
    case ' ' :
        tok = newToken(token.WHITESPACE, l.ch)
    }
    if unicode.IsLetter(rune(l.ch)) {
        for l.nextCharDontBreakWord() {
            l.readPosition++
        }
        tok = newTokenFromString(token.WORD, l.input[l.position:l.readPosition])
        l.position = l.readPosition - 1
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.HTMLToken {
    return token.HTMLToken{Type: tokenType, Literal: string(ch)}
}

func newTokenFromString(tokenType token.TokenType, chars string) token.HTMLToken {
    return token.HTMLToken{Type: tokenType, Literal: chars}
}
