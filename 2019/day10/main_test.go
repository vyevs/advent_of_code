package main

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{
			a:    48,
			b:    180,
			want: 12,
		},
		{
			a:    54,
			b:    24,
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("gcd(%d, %d)", test.a, test.b), func(t *testing.T) {
			got := gcd(test.a, test.b)

			if got != test.want {
				t.Errorf("got: %d, want: %d", got, test.want)
			}
		})
	}
}

func TestGCDDirection(t *testing.T) {
	tests := []struct {
		x, y, otherX, otherY int
		want                 direction
	}{
		{
			x:      5,
			y:      5,
			otherX: 7,
			otherY: 5,
			want: direction{
				dx: 1,
				dy: 0,
			},
		},
		{
			x:      5,
			y:      5,
			otherX: 5,
			otherY: 7,
			want: direction{
				dx: 0,
				dy: 1,
			},
		},
		{
			x:      5,
			y:      5,
			otherX: 3,
			otherY: 5,
			want: direction{
				dx: -1,
				dy: 0,
			},
		},
		{
			x:      5,
			y:      5,
			otherX: 5,
			otherY: 3,
			want: direction{
				dx: 0,
				dy: -1,
			},
		},
		{
			x:      5,
			y:      5,
			otherX: 25,
			otherY: 25,
			want: direction{
				dx: 1,
				dy: 1,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			got := gcdDirection(test.x, test.y, test.otherX, test.otherY)

			if got != test.want {
				t.Errorf("got %+v want %+v", got, test.want)
			}
		})
	}
}
