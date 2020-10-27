package main

import "testing"

func TestTryToReadMarker(t *testing.T) {
	tests := []struct {
		input   string
		wantStr string
		wantOk  bool
	}{
		{
			input:  "ADVENT",
			wantOk: false,
		},
		{
			input:   "(1x5)",
			wantOk:  true,
			wantStr: "(1x5)",
		},
		{
			input:   "(123123x9123123)",
			wantOk:  true,
			wantStr: "(123123x9123123)",
		},
		{
			input:   "(123123x9123123)2312312361",
			wantOk:  true,
			wantStr: "(123123x9123123)",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			gotStr, gotOk := tryToReadMarker(test.input)

			if gotOk != test.wantOk {
				t.Fatalf("got %v ok, want %v", gotOk, test.wantOk)
			} else if gotStr != test.wantStr {
				t.Fatalf("got %s, want %s", gotStr, test.wantStr)
			}
		})
	}
}

func TestDecompressedLength2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "ADVENT",
			want:  6,
		},
		{
			input: "(3x3)abc",
			want:  9,
		},
		{
			input: "abc123(3x3)abc",
			want:  15,
		},
		{
			input: "abc123(3x3)abcd",
			want:  16,
		},
		{
			input: "abc123(3x3)ab",
			want:  12,
		},
		{
			input: "abc123(8x3)(3x3)abc",
			want:  33,
		},
		{
			input: "X(8x2)(3x3)ABCY",
			want:  20,
		},
		{
			input: "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
			want:  445,
		},
		{
			input: "(27x12)(20x12)(13x14)(7x10)(1x12)A",
			want:  241920,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := decompressedLength2(test.input, len(test.input), 1)

			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})

	}
}
