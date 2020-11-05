package main

import (
	"strconv"
	"testing"
)

func TestDetermineThrusterSignalForPhaseSequence(t *testing.T) {
	tests := []struct {
		instrs        []int
		phaseSettings [5]int
		wantSignal    int
	}{
		{
			instrs:        []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			phaseSettings: [5]int{4, 3, 2, 1, 0},
			wantSignal:    43210,
		},
		{
			instrs: []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
				101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			phaseSettings: [5]int{0, 1, 2, 3, 4},
			wantSignal:    54321,
		},
		{
			instrs: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			phaseSettings: [5]int{1, 0, 4, 3, 2},
			wantSignal:    65210,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := determineThrusterSignalForPhaseSequence(test.instrs, test.phaseSettings)

			if got != test.wantSignal {
				t.Errorf("got signal %d, want %d", got, test.wantSignal)
			}
		})
	}
}
func TestDetermineHighestSignalThatCanBeSentToThrusters(t *testing.T) {
	tests := []struct {
		instrs     []int
		wantSignal int
	}{
		{
			instrs:     []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			wantSignal: 43210,
		},
		{
			instrs: []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
				101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			wantSignal: 54321,
		},
		{
			instrs: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			wantSignal: 65210,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := determineHighestSignalThatCanBeSentToThrusters(test.instrs)

			if got != test.wantSignal {
				t.Errorf("got signal %d, want %d", got, test.wantSignal)
			}
		})
	}
}

func TestDetermineThrusterSignalForPhaseSequenceLooping(t *testing.T) {
	tests := []struct {
		instrs        []int
		phaseSettings [5]int
		wantSignal    int
	}{
		{
			instrs: []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
				27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			phaseSettings: [5]int{9, 8, 7, 6, 5},
			wantSignal:    139629729,
		},
		{
			instrs: []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
				-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
				53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			phaseSettings: [5]int{9, 7, 8, 5, 6},
			wantSignal:    18216,
		},
	}
	//log.SetOutput(ioutil.Discard)
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := determineThrusterSignalForPhaseSequenceLooping(test.instrs, test.phaseSettings)

			if got != test.wantSignal {
				t.Errorf("got signal %d, want %d", got, test.wantSignal)
			}
		})
	}
}
