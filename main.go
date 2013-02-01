package main

import (
//    "fmt"
)

func main() {
    //fmt.Println(Interpret("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.", 3000))
    exp := CreateExperiment("hi", 100, 256, 512)
    exp.Start()
}
