package main

import (
    "fmt"
    "brainfuck"
)

func main() {
    fmt.Println(Interpret("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.", 3000))
}
