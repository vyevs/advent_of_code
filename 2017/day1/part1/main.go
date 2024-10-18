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
    
    for i := 0; i < len(digits) - 1; i++ {
        c := digits[i]
        nc := digits[i+1]
        
        if c == nc {
            sum += int(c - '0');
        }
    }
    
    if digits[0] == digits[len(digits)-1] {
        sum += int(digits[0] - '0');
    }

    return sum
}


func getInput() string {
    bs, _ := io.ReadAll(os.Stdin)
    return string(bs)
}