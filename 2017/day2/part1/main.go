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
        minV := min(row)
        maxV := max(row)
        
        diff := maxV - minV
        
        checksum += diff
    }
    
    return checksum
}

func min(vs []int) int {
    m := vs[0]
    for i := 1; i < len(vs); i++ {
        v := vs[i]
        if v < m {
            m = v
        }
    }
    return m
}

func max(vs []int) int {
    m := vs[0]
    for i := 1; i < len(vs); i++ {
        v := vs[i]
        if v > m {
            m = v
        }
    }
    return m
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