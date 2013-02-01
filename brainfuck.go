package main

import (
//    "fmt"
    "time"
)

func Interpret(input string, cellcount int) (output string) {
    cell_ptr := 0
    cells := make([]uint8, cellcount)
    input_ptr := 0
    loop_depth := 0

    // Used for tracking infinite loops.
    start_time := time.Now()

    for ; input_ptr < len(input); input_ptr++ {
        elapsed_time := time.Since(start_time)

        // We're probably in an infinite loop, time to go home.
        // TODO Find sweetspot in time.
        if elapsed_time.Nanoseconds() > 5000000 {
            //fmt.Println("Timed out!", output)
            return
        }

        switch input[input_ptr] {
            case '>':
                // Loop around.
                if cell_ptr < len(cells) - 2 {
                    cell_ptr++
                } else {
                    cell_ptr = 0
                }
            case '<':
                // Loop around if below zero.
                if cell_ptr > 0 {
                    cell_ptr--
                } else {
                    cell_ptr = len(cells) - 1
                }
            case '+':
                cells[cell_ptr]++
            case '-':
                cells[cell_ptr]--
            case '.':
                output += string(cells[cell_ptr])
            case ',':
                // TODO
                // Input not needed yet, so not supported yet either.
            case '[':
                if cells[cell_ptr] == 0 {
                    //FIXME oob quick fix, do better.
                    //if input_ptr == len(input) - 1 {
                    //    return
                    //}

                    input_ptr++
                    for ; input_ptr < len(input) && (loop_depth > 0 || input[input_ptr] != ']'); input_ptr++ {
                        if input[input_ptr] == '[' {
                            loop_depth++
                        } else if input[input_ptr] == ']' {
                            loop_depth--
                        }

                        //FIXME oob quick fix, do better.
                        //if input_ptr == len(input) - 1 {
                        //    return
                        //}
                    }
                }
            case ']':
                if cells[cell_ptr] != 0 {
                    input_ptr--
                    for ; input_ptr > 0 && (loop_depth > 0 || input[input_ptr] != '['); input_ptr-- {
                        if input[input_ptr] == ']' {
                            loop_depth++
                        } else if input[input_ptr] == '[' {
                            loop_depth--
                        }

                        //FIXME oob quick fix, do better.
                        //if input_ptr-1 < 0 {
                        //    return
                        //}
                    }
                }
            default:
                // TODO
                // At some point, handle illegal instructions properly.
        }
    }
    //fmt.Println("Did not time out! Output: ", output)
    return
}
