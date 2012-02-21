package main

import (
    "fmt"
    "strings"
    "strconv"
    "./parser"
)

func execute(op string, args []string) string {
    if op == "join" {
        return strings.Join(args, " ")
    }
    if op == "print" {
        fmt.Println(args[0])
        return ""
    }
    if op == "sum" {
        sum := 0
        for _, x := range args {
            num, err := strconv.Atoi(x)
            if err != nil {
                panic(fmt.Sprintf("Argument to sum non-numeric: '%s'\n", x))
            }
            sum += num
        }
        return strconv.Itoa(sum)
    }
    if op == ">" {
        num1, err := strconv.Atoi(args[0])
        if err != nil {
            panic(fmt.Sprintf("Argument to > non-numeric: '%s'\n", args[0]))
        }
        num2, err := strconv.Atoi(args[1])
        if err != nil {
            panic(fmt.Sprintf("Argument to > non-numeric: '%s'\n", args[1]))
        }
        if num1 > num2 {
            return "true"
        }
        return ""
    }
    if op == "not" {
        if args[0] == "" {
            return "true"
        }
        return ""
    }
    if op == "if" {
        if args[0] != "" {
            return args[1]
        }
        return args[2]
    }
    panic(fmt.Sprintf("Unknown op: %s", op))
}

func eval(p *parser.Parser) string {
    if p.Lookahead() != "(" {
        panic("Syntax error: expected '('")
    }
    var tok string
    tok = p.Get_token() // eat "("
    tok = p.Get_token() // proceed

    op   := tok
    args := make([]string, 0)
    for p.Lookahead() != ")" {
        if p.Lookahead() == "(" {
            var res string
            res  = eval(p)
            args = append(args, res)
        } else {
            tok  = p.Get_token()
            args = append(args, tok)
        }
    }
    tok = p.Get_token()
    if tok != ")" {
        panic("Syntax error: expected ')'")
    }
    return execute(op, args)
}

func main() {
    code := `(print (join Hello World
                     (if (not (> 2 5))
                         (sum 40 2)
                         (sum 40 11)
                     )
             ))`
    p := parser.New(code)
    eval(p)
}
