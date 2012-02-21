package parser

import (
    "bytes"
    "fmt"
)

func is_ws(char byte) bool {
    return char == ' ' || char == '\n' || char == '\t'
}

type Parser struct {
    code   string
    cursor int
}

func New(code string) *Parser {
    return &Parser{code, 0}
}

func (p *Parser) Get_token() (string) {
    if p.cursor >= len(p.code) {
        p.cursor = -1
        return ""
    }
    for is_ws(p.code[p.cursor]) {
        p.cursor++
    }
    look := p.code[p.cursor]
    if look == '(' || look == ')' {
        p.cursor++
        return string(look)
    }
    ret := bytes.NewBufferString("")
    for look != '(' && look != ')' && !is_ws(look) {
        fmt.Fprint(ret, string(look))
        p.cursor++
        look = p.code[p.cursor]
    }
    return string(ret.Bytes())
}

func (p *Parser) Lookahead() string {
    cur := p.cursor
    for is_ws(p.code[cur]) {
        cur++
    }
    return string(p.code[cur])
}
