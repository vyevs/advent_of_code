package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    digits := getInput()
    
    sum := solveCaptcha(digits)
    
    fmt.Println(sum)
}

func solveCaptcha(digits string) int {
    var sum int
    
    dLen := len(digits)
    offset := dLen / 2
    
    
    for i := 0; i < len(digits); i++ {
        c := digits[i]
        
        ncIdx := (i + offset) % dLen
        nc := digits[ncIdx]
        
        if c == nc {
            sum += int(c - '0');
        }
    }

    return sum
}


func getInput() string {
    bs, _ := io.ReadAll(os.Stdin)
    return string(bs)
}