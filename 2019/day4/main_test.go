package main

import (
	"strconv"
	"testing"
)

func TestHasAdjacentSameDigits(t *testing.T) {
	tests := []struct {
		in   int
		want bool
	}{
		{
			in:   111111,
			want: true,
		},
		{
			in:   123789,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			got := hasAdjacentSameDigits(test.in)

			if got != test.want {
				t.Errorf("got: %v want: %v", got, test.want)
			}
		})
	}
}
