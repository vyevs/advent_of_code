package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseInstructionInt(t *testing.T) {
	tests := []struct {
		c     computer
		instr int
		want  instruction
	}{
		{
			c: computer{
				instrs: []int{1002, 4, 3, 4},
				pc:     0,
			},
			instr: 1002,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     4,
				op2:     3,
				op3:     4,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
			},
		},
		{
			c: computer{
				instrs: []int{9, 9, 9, 1002, 4, 3, 4},
				pc:     3,
			},
			instr: 1002,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     4,
				op2:     3,
				op3:     4,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
			},
		},
		{
			c: computer{
				instrs: []int{1101, 100, -1, 4},
				pc:     0,
			},
			instr: 1101,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     100,
				op2:     -1,
				op3:     4,
				op1Mode: opModeImmediate,
				op2Mode: opModeImmediate,
			},
		},
		{
			c: computer{
				instrs: []int{104, 100, -1, 4},
				pc:     0,
			},
			instr: 104,
			want: instruction{
				typ:     instructionTypeOutput,
				op1:     100,
				op1Mode: opModeImmediate,
			},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := test.c.parseInstructionInt(test.instr)

			if got != test.want {
				t.Errorf("got: %+v want: %+v", got, test.want)
			}
		})
	}
}

func TestRunComputer(t *testing.T) {

	type inputOutputTest struct {
		input      int
		wantOutput int
	}

	tests := []struct {
		inputInstrs      []int
		inputOutputTests []inputOutputTest
	}{
		{
			inputInstrs: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			inputOutputTests: []inputOutputTest{
				{
					input:      8,
					wantOutput: 1,
				},
				{
					input:      7,
					wantOutput: 0,
				},
			},
		},

		{
			inputInstrs: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			inputOutputTests: []inputOutputTest{
				{
					input:      7,
					wantOutput: 1,
				},
				{
					input:      8,
					wantOutput: 0,
				},
				{
					input:      9,
					wantOutput: 0,
				},
			},
		},

		{
			inputInstrs: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			inputOutputTests: []inputOutputTest{
				{
					input:      8,
					wantOutput: 1,
				},
				{
					input:      7,
					wantOutput: 0,
				},
			},
		},

		{
			inputInstrs: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			inputOutputTests: []inputOutputTest{
				{
					input:      7,
					wantOutput: 1,
				},
				{
					input:      8,
					wantOutput: 0,
				},
				{
					input:      9,
					wantOutput: 0,
				},
			},
		},

		{
			inputInstrs: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			inputOutputTests: []inputOutputTest{
				{
					input:      0,
					wantOutput: 0,
				},
				{
					input:      -420,
					wantOutput: 1,
				},
				{
					input:      69,
					wantOutput: 1,
				},
			},
		},

		{
			inputInstrs: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			inputOutputTests: []inputOutputTest{
				{
					input:      0,
					wantOutput: 0,
				},
				{
					input:      -420,
					wantOutput: 1,
				},
				{
					input:      69,
					wantOutput: 1,
				},
			},
		},

		{
			inputInstrs: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			inputOutputTests: []inputOutputTest{
				{
					input:      7,
					wantOutput: 999,
				},
				{
					input:      8,
					wantOutput: 1000,
				},
				{
					input:      9,
					wantOutput: 1001,
				},
			},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			for _, inOutTest := range test.inputOutputTests {
				t.Run(fmt.Sprintf("%d -> %d", inOutTest.input, inOutTest.wantOutput), func(t *testing.T) {
					var got int
					c := computer{
						getInput: func() int {
							return inOutTest.input
						},
						produceOutput: func(out int) {
							got = out
						},
					}

					c.run(test.inputInstrs)

					if got != inOutTest.wantOutput {
						t.Errorf("got output %d want %d", got, inOutTest.wantOutput)
					}
				})
			}

		})
	}
}
