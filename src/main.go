package main

import (
    "fmt"
    "scanner"
)

func main() {
    source := "/* comment*/def foo:int[]..[](x1, x2ńńńńńńń) { return =  0xff + 077 + 12.34 + 56 + \"hello \n\rw\torld\" + 'hello world' + x1 * y2 + x3 / y3 && y || !y == y != {m, n} ? c)//hello world"
    scanner := scanner.CreateScanner(source)
    token := scanner.Scan()
    for !scanner.EOF() {
        fmt.Println(token.Value())
        token = scanner.Scan()
    }
}
