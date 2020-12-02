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
		// instructionTypeAdd
		{
			c: computer{
				memory: []int{1, 57, 58, 59},
				pc:     0,
			},
			instr: 1,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{101, 57, 58, 59},
				pc:     0,
			},
			instr: 101,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1101, 57, 58, 59},
				pc:     0,
			},
			instr: 1101,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1001, 57, 58, 59},
				pc:     0,
			},
			instr: 1001,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{201, 57, 58, 59},
				pc:     0,
			},
			instr: 201,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1201, 57, 58, 59},
				pc:     0,
			},
			instr: 1201,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{2201, 57, 58, 59},
				pc:     0,
			},
			instr: 2201,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeRelative,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{20001, 57, 58, 59},
				pc:     0,
			},
			instr: 20001,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{20101, 57, 58, 59},
				pc:     0,
			},
			instr: 20101,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21101, 57, 58, 59},
				pc:     0,
			},
			instr: 21101,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21001, 57, 58, 59},
				pc:     0,
			},
			instr: 21001,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{20201, 57, 58, 59},
				pc:     0,
			},
			instr: 20201,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21201, 57, 58, 59},
				pc:     0,
			},
			instr: 21201,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{22201, 57, 58, 59},
				pc:     0,
			},
			instr: 22201,
			want: instruction{
				typ:     instructionTypeAdd,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeRelative,
				op3Mode: opModeRelative,
			},
		},
		// end instructionTypeAdd

		// instructionTypeMultiply
		{
			c: computer{
				memory: []int{2, 57, 58, 59},
				pc:     0,
			},
			instr: 2,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{102, 57, 58, 59},
				pc:     0,
			},
			instr: 102,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1102, 57, 58, 59},
				pc:     0,
			},
			instr: 1102,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1002, 57, 58, 59},
				pc:     0,
			},
			instr: 1002,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{202, 57, 58, 59},
				pc:     0,
			},
			instr: 202,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1202, 57, 58, 59},
				pc:     0,
			},
			instr: 1202,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{2202, 57, 58, 59},
				pc:     0,
			},
			instr: 2202,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeRelative,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{20002, 57, 58, 59},
				pc:     0,
			},
			instr: 20002,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{20102, 57, 58, 59},
				pc:     0,
			},
			instr: 20102,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21102, 57, 58, 59},
				pc:     0,
			},
			instr: 21102,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21002, 57, 58, 59},
				pc:     0,
			},
			instr: 21002,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{20202, 57, 58, 59},
				pc:     0,
			},
			instr: 20202,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21202, 57, 58, 59},
				pc:     0,
			},
			instr: 21202,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{22202, 57, 58, 59},
				pc:     0,
			},
			instr: 22202,
			want: instruction{
				typ:     instructionTypeMultiply,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeRelative,
				op3Mode: opModeRelative,
			},
		},
		// end instructionTypeMultiply

		// instructionTypeLessThan
		{
			c: computer{
				memory: []int{7, 57, 58, 59},
				pc:     0,
			},
			instr: 7,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{107, 57, 58, 59},
				pc:     0,
			},
			instr: 107,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1107, 57, 58, 59},
				pc:     0,
			},
			instr: 1107,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1007, 57, 58, 59},
				pc:     0,
			},
			instr: 1007,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{207, 57, 58, 59},
				pc:     0,
			},
			instr: 207,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModePosition,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1207, 57, 58, 59},
				pc:     0,
			},
			instr: 1207,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeImmediate,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{2207, 57, 58, 59},
				pc:     0,
			},
			instr: 2207,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeRelative,
				op3Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{20007, 57, 58, 59},
				pc:     0,
			},
			instr: 20007,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{20107, 57, 58, 59},
				pc:     0,
			},
			instr: 20107,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21107, 57, 58, 59},
				pc:     0,
			},
			instr: 21107,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeImmediate,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21007, 57, 58, 59},
				pc:     0,
			},
			instr: 21007,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{20207, 57, 58, 59},
				pc:     0,
			},
			instr: 20207,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModePosition,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{21207, 57, 58, 59},
				pc:     0,
			},
			instr: 21207,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeImmediate,
				op3Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{22207, 57, 58, 59},
				pc:     0,
			},
			instr: 22207,
			want: instruction{
				typ:     instructionTypeLessThan,
				op1:     57,
				op2:     58,
				op3:     59,
				op1Mode: opModeRelative,
				op2Mode: opModeRelative,
				op3Mode: opModeRelative,
			},
		},
		// end instructionTypeLessThan

		// instructionTypeInput
		{
			c: computer{
				memory: []int{3, 57},
				pc:     0,
			},
			instr: 3,
			want: instruction{
				typ:     instructionTypeInput,
				op1:     57,
				op1Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{203, 57},
				pc:     0,
			},
			instr: 203,
			want: instruction{
				typ:     instructionTypeInput,
				op1:     57,
				op1Mode: opModeRelative,
			},
		},
		// end instructionTypeInput

		// instructionTypeOutput
		{
			c: computer{
				memory: []int{4, 57},
				pc:     0,
			},
			instr: 4,
			want: instruction{
				typ:     instructionTypeOutput,
				op1:     57,
				op1Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{204, 57},
				pc:     0,
			},
			instr: 204,
			want: instruction{
				typ:     instructionTypeOutput,
				op1:     57,
				op1Mode: opModeRelative,
			},
		},
		// end instructionTypeOutput

		// instructionTypeJumpIfTrue
		{
			c: computer{
				memory: []int{5, 57, 58},
				pc:     0,
			},
			instr: 5,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModePosition,
				op2Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1005, 57, 58},
				pc:     0,
			},
			instr: 1005,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModePosition,
				op2Mode: opModeImmediate,
			},
		},
		{
			c: computer{
				memory: []int{2005, 57, 58},
				pc:     0,
			},
			instr: 2005,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModePosition,
				op2Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{105, 57, 58},
				pc:     0,
			},
			instr: 105,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModeImmediate,
				op2Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{2105, 57, 58},
				pc:     0,
			},
			instr: 2105,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModeImmediate,
				op2Mode: opModeRelative,
			},
		},

		{
			c: computer{
				memory: []int{2105, 57, 58},
				pc:     0,
			},
			instr: 2105,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModeImmediate,
				op2Mode: opModeRelative,
			},
		},
		{
			c: computer{
				memory: []int{205, 57, 58},
				pc:     0,
			},
			instr: 205,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModeRelative,
				op2Mode: opModePosition,
			},
		},
		{
			c: computer{
				memory: []int{1205, 57, 58},
				pc:     0,
			},
			instr: 1205,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModeRelative,
				op2Mode: opModeImmediate,
			},
		},
		{
			c: computer{
				memory: []int{2205, 57, 58},
				pc:     0,
			},
			instr: 2205,
			want: instruction{
				typ:     instructionTypeJumpIfTrue,
				op1:     57,
				op2:     58,
				op1Mode: opModeRelative,
				op2Mode: opModeRelative,
			},
		},
		// end instructionTypeJumpIfTrue
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

					initialMemoryCopy := make([]int, len(test.inputInstrs))
					copy(initialMemoryCopy, test.inputInstrs)

					c := computer{
						memory: initialMemoryCopy,
						getInput: func() int {
							return inOutTest.input
						},
						produceOutput: func(v int) {
							got = v
						},
					}

					c.run()

					if got != inOutTest.wantOutput {
						t.Errorf("got output %d want %d", got, inOutTest.wantOutput)
					}
				})
			}

		})
	}
}
