package main

func Interpret(input string, cellcount int) (output string) {
    cell_ptr := 0
    cells := make([]uint8, cellcount)
    input_ptr := 0
    loop_depth := 0

    for ; input_ptr < len(input); input_ptr++ {
        switch input[input_ptr] {
            case '>':
                cell_ptr++
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
                    input_ptr++
                    for ; loop_depth > 0 || input[input_ptr] != ']'; input_ptr++ {
                        if input[input_ptr] == '[' {
                            loop_depth++
                        } else if input[input_ptr] == ']' {
                            loop_depth--
                        }

                        //FIXME oob quick fix, do better.
                        if input_ptr-1 == len(input) {
                            return ""
                        }
                    }
                }
            case ']':
                if cells[cell_ptr] != 0 {
                    input_ptr--
                    for ; loop_depth > 0 || input[input_ptr] != '['; input_ptr-- {
                        if input[input_ptr] == ']' {
                            loop_depth++
                        } else if input[input_ptr] == '[' {
                            loop_depth--
                        }

                        //FIXME oob quick fix, do better.
                        if input_ptr-1 < 0 {
                            return ""
                        }
                    }
                }
            default:
                // TODO
                // At some point, handle illegal instructions properly.
        }
    }
    return
}
