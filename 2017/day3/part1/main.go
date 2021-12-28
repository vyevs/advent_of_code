package main

import (
    "fmt"
)

func main() {
    input := getInput()
    
    
    nSteps := determineNumberOfSteps(input)
    
    fmt.Printf("it takes %d steps to get from %d to 1\n", nSteps, input)
}

func determineNumberOfSteps(v int) int {  
    for i := 1; ; i += 2 {
        sq := i * i
        if v <= sq {
            return i/2
        }
    }
}

func getInput() int {
    var v int
    _, _ = fmt.Scan(&v)
    return v
}