package main

import (
//    "fmt"
)

func main() {
    //fmt.Println(Interpret("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.", 3000))
    exp := CreateExperiment("hei", 100, 256, 768)
    exp.Start()
}
