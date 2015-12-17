package main

import (
    "fmt"
    "scanner"
)

func main() {
    source :=
`
/*
multiline comment 1
multiline comment 2
*/
import ("test_library")

x = 1 + 2 * 3 / 5;
y = true && false || true;
xń = 0xff + 077 + 12.34;  // single line comment 1
// single line comment 2
def foo:int[]..[]()
{
    return = "hello\nworld\n!";
}
`
    lexer := scanner.CreateScanner(source)
    for !lexer.EOF() {
        token := lexer.Scan()
        fmt.Printf("(%2d, %2d) %s\n", token.Line, token.Col, token.StringValue())
    }
}
