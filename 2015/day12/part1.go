package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	var m map[string]interface{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		log.Fatalf("Decode: %v", err)
	}
	f.Close()

	var sumf func(interface{}) float64
	sumf = func(i interface{}) float64 {
		var sum float64

		switch concrete := i.(type) {
		case float64:
			sum += concrete

		case []interface{}:
			for _, v := range concrete {
				sum += sumf(v)
			}

		case map[string]interface{}:
			if !containsRed(concrete) {
				for _, v := range concrete {
					sum += sumf(v)
				}
			}
		}

		return sum
	}

	sum := sumf(m)

	fmt.Println(sum)
}

func containsRed(m map[string]interface{}) bool {
	for _, v := range m {
		switch concrete := v.(type) {
		case string:
			if concrete == "red" {
				return true
			}
		}
	}

	return false
}
