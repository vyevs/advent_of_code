package main

import (
	"fmt"
	"testing"
)

func TestKmpIndexAll(t *testing.T) {
	tests := []struct {
		s    string
		w    string
		want []int
	}{
		{
			s:    "ABC ABCDAB ABCDABCDABDE",
			w:    "ABCDABD",
			want: []int{15},
		},
		{
			s:    "AGCATGCTGCAGTCATGCTTAGGCTA",
			w:    "GCT",
			want: []int{5, 16, 22},
		},
		{
			s:    "aaaa",
			w:    "aa",
			want: []int{0, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s/%s", test.s, test.w), func(t *testing.T) {
			got := kmpIndexAll(test.s, test.w)

			if !intSlicesEqual(got, test.want) {
				t.Errorf("\ngot:  %v\nwant: %v", got, test.want)
			}
		})
	}
}

func TestKmpTable(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{
			s:    "ABCDABD",
			want: []int{-1, 0, 0, 0, -1, 0, 2, 0},
		},
		{
			s:    "ABACABABC",
			want: []int{-1, 0, -1, 1, -1, 0, -1, 3, 2, 0},
		},
		{
			s:    "ABACABABA",
			want: []int{-1, 0, -1, 1, -1, 0, -1, 3, -1, 3},
		},
		{
			s:    "PARTICIPATE IN PARACHUTE",
			want: []int{-1, 0, 0, 0, 0, 0, 0, -1, 0, 2, 0, 0, 0, 0, 0, -1, 0, 0, 3, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			got := kmpTable(test.s)

			if !intSlicesEqual(got, test.want) {
				t.Errorf("\ngot:  %v\nwant: %v", got, test.want)
			}
		})
	}
}

func intSlicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
