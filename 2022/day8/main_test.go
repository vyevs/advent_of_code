package main

import (
	"os"
	"testing"
)

func BenchmarkCountVisibleTrees(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		b.Fatalf("error opening file: %v", err)
	}
	
	g := getGrid(f)
	_ = f.Close()
	
	want := 1835
	
	b.Run("naive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			got := countVisibleTrees(g)
			if got != want {
				b.Fatalf("incorrect answer %d, want %d", got, want)
			}	
		}
	})
}