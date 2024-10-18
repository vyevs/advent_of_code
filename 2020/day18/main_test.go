package main

import "testing"

func TestEvaluate1(t *testing.T) {
	tests := []struct {
		expr string
		want int
	}{
		{
			expr: "1 + 2 * 3 + 4 * 5 + 6",
			want: 71,
		},
		{
			expr: "1 + (2 * 3) + (4 * (5 + 6))",
			want: 51,
		},
		{
			expr: "2 * 3 + (4 * 5)",
			want: 26,
		},
		{
			expr: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			want: 437,
		},
		{
			expr: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			want: 12240,
		},
		{
			expr: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			want: 13632,
		},
	}

	for _, test := range tests {
		got := evaluate1(test.expr)

		if got != test.want {
			t.FailNow()

			t.Logf("Expression: %q", test.expr)
			t.Logf("want = %d got = %d", test.want, got)
		}
	}
}

func TestEvaluate2(t *testing.T) {
	tests := []struct {
		expr string
		want int
	}{
		{
			expr: "1 + 2 * 3 + 4 * 5 + 6",
			want: 231,
		},
		{
			expr: "1 + (2 * 3) + (4 * (5 + 6))",
			want: 51,
		},
		{
			expr: "2 * 3 + (4 * 5)",
			want: 46,
		},
		{
			expr: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			want: 1445,
		},
		{
			expr: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			want: 669060,
		},
		{
			expr: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			want: 23340,
		},
		{
			expr: "5 * 8 * 4 + (7 * 6)",
			want: 1840,
		},
	}

	for _, test := range tests {
		got := evaluate2(test.expr)

		if got != test.want {
			t.FailNow()

			t.Logf("Expression: %q", test.expr)
			t.Logf("want = %d got = %d", test.want, got)
		}
	}
}
