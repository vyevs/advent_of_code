package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)


func main() {
	start := time.Now()
	if err := doIt(); err != nil {
		log.Printf("something went wrong: %v\n", err)
	} else {
		fmt.Println("success :)")
	}
	fmt.Printf("that took %f seconds\n", time.Since(start).Seconds())
}

func doIt() error {
	ps, err := getCoordinates()
	if err != nil {
		return err
	}
	fmt.Println(ps)
	
	return nil
}

type p struct {
	x, y int
}

func getCoordinates() ([]p, error) {
	s := bufio.NewScanner(os.Stdin)
	
	
	ps := make([]p, 0, 1024)
	for s.Scan() {
		line := s.Text()
		
		var p p
		_, err := fmt.Sscanf(line, "%d, %d", &p.x, &p.y)
		if err != nil {
			return nil, fmt.Errorf("error parsing line: %v", err)
		}
		
		ps = append(ps, p)
	}
	
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %v", err)
	}
	
	return ps, nil
}