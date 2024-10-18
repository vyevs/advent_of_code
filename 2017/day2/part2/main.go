package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    rows := readSpreadsheet()
    
    checksum := calculateChecksum(rows)
    
    fmt.Printf("the spreadsheet's checksum is %d\n", checksum)
}

func calculateChecksum(rows [][]int) int {
    var checksum int
    
    for _, row := range rows {
        checksum += divisionResult(row)
    }
    
    return checksum
}

func divisionResult(row []int) int {
    for i := 0; i < len(row); i++ {
        for j := 0; j < len(row); j++ {
            if i == j {
                continue
            }
            
            v1 := row[i]
            v2 := row[j]
            
            quo := v1 / v2
            rem := v1 % v2
            
            if rem == 0 && quo != 0 {
                return quo
            }
        }
    }
    
    
    panic(fmt.Sprintf("row %v does not have an evenly divisible pair", row))
}


func readSpreadsheet() [][]int {
    scanner := bufio.NewScanner(os.Stdin)
    
    
    
    rows := make([][]int, 0, 64)
    
    for scanner.Scan() {
        line := scanner.Text()
                
        row := make([]int, 0, 64)
        
        lineR := strings.NewReader(line)
        
        for {
            var v int
            n, _ := fmt.Fscan(lineR, &v)
            
            if n != 1 {
                break
            }
            
            row = append(row, v)
        }
        
        
        rows = append(rows, row)
    }
    
    return rows
}