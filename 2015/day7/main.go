package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	ops, err := parseOps(f)
	if err != nil {
		if parseErr, ok := err.(parsingError); ok {
			fmt.Printf("parsing failure: %v\n", err)
			fmt.Println(parseErr.line)
		} else {
			log.Fatalf("parsing err: %v", err)
		}
	}
	f.Close()

	wires, err := evaluateOps(ops)
	if err != nil {
		handleEvaluationErr(err)
	}

	fmt.Printf("wire a: %d\n", wires["a"])

	newWires := make(map[string]uint16, 1024)
	newWires["b"] = wires["a"]

	overrides := map[string]struct{}{"b": struct{}{}}

	newWires, err = evaluateOpsFromState(ops, newWires, overrides)
	if err != nil {
		handleEvaluationErr(err)
	}

	fmt.Printf("part 2 wire a: %d\n", newWires["a"])

	recursiveAVal := determineFinalSignalOnWire(ops, nil, "a")
	fmt.Printf("recursive value of a: %d\n", recursiveAVal)

	newWires = make(map[string]uint16, 1024)
	newWires["b"] = wires["a"]
	recursiveAVal = determineFinalSignalOnWire(ops, newWires, "a")
	fmt.Printf("recursive value of a: %d\n", recursiveAVal)
}

func determineFinalSignalOnWire(ops []op, initialState map[string]uint16, wire string) uint16 {
	howToGetWireSignals := make(map[string]op, len(ops))

	var wireSignals map[string]uint16
	if initialState == nil {
		wireSignals = make(map[string]uint16, len(ops))
	} else {
		wireSignals = initialState
	}

	for _, op := range ops {
		howToGetWireSignals[op.dstWire] = op
	}

	var recursiveEvaluator func(target string) uint16

	recursiveEvaluator = func(target string) uint16 {
		if currentSignal, ok := wireSignals[target]; ok {
			return currentSignal
		}

		op := howToGetWireSignals[target]

		var result uint16

		switch op.typ {
		case opTypeValue:
			wireSignals[target] = op.immediate
			return op.immediate

		case opTypeWire:
			result = recursiveEvaluator(op.srcWire1)
			wireSignals[target] = result

		case opTypeAndWires:
			srcWire1Val := recursiveEvaluator(op.srcWire1)
			srcWire2Val := recursiveEvaluator(op.srcWire2)
			result = srcWire1Val & srcWire2Val
			wireSignals[target] = result

		case opTypeAndWireImmediate:
			srcWireVal := recursiveEvaluator(op.srcWire1)
			result = srcWireVal & op.immediate
			wireSignals[target] = result

		case opTypeOrWires:
			srcWire1Val := recursiveEvaluator(op.srcWire1)
			srcWire2Val := recursiveEvaluator(op.srcWire2)
			result = srcWire1Val | srcWire2Val
			wireSignals[target] = result

		case opTypeOrWireImmediate:
			srcWireVal := recursiveEvaluator(op.srcWire1)
			result = srcWireVal | op.immediate
			wireSignals[target] = result

		case opTypeRShift:
			srcWireVal := recursiveEvaluator(op.srcWire1)
			result = srcWireVal >> op.immediate
			wireSignals[target] = result

		case opTypeLShift:
			srcWireVal := recursiveEvaluator(op.srcWire1)
			result = srcWireVal << op.immediate
			wireSignals[target] = result

		case opTypeNot:
			srcWireVal := recursiveEvaluator(op.srcWire1)
			result = ^srcWireVal
			wireSignals[target] = result
		}

		return result
	}

	return recursiveEvaluator(wire)
}

func handleEvaluationErr(err error) {
	if depErr, ok := err.(dependencyError); ok {
		fmt.Println("unevaluated ops: ")
		for _, op := range depErr.unevaluated {
			fmt.Println(&op)
		}
		fmt.Println("wire values: ")
		for wire, value := range depErr.wires {
			fmt.Printf("%s: %d\n", wire, value)
		}
	} else {
		log.Fatalf("unable to evalute: %s", err)
	}
}

type opType int

const (
	opTypeValue opType = iota
	opTypeWire
	opTypeAndWires
	opTypeAndWireImmediate
	opTypeOrWires
	opTypeOrWireImmediate
	opTypeRShift
	opTypeLShift
	opTypeNot
)

func (typ opType) String() string {
	switch typ {
	case opTypeValue:
		return "opType:Value"
	case opTypeWire:
		return "opType:Wire"
	case opTypeAndWires:
		return "opType:AndWires"
	case opTypeAndWireImmediate:
		return "opType:AndWireImmediate"
	case opTypeOrWires:
		return "opType:OrWires"
	case opTypeOrWireImmediate:
		return "opType:OrWireImmediate"
	case opTypeRShift:
		return "opType:RShift"
	case opTypeLShift:
		return "opType:LShift"
	case opTypeNot:
		return "opType:Not"
	}
	log.Fatalf("unknown type String() %d", typ)
	return ""
}

type op struct {
	typ opType

	immediate uint16 // for opTypeValue, opTypeRShift, opTypeLShift, opTypeAndWireImmediate, opTypeOrWireImmediate

	srcWire1 string
	srcWire2 string

	dstWire string
}

func (op *op) String() string {
	var buf strings.Builder

	buf.WriteString(op.typ.String())

	if op.typ == opTypeValue {
		buf.WriteString(fmt.Sprintf(" srcValue: %d", op.immediate))
	} else {
		buf.WriteString(" srcWire1: ")
		buf.WriteString(op.srcWire1)

		if op.typ == opTypeAndWires || op.typ == opTypeOrWires {
			buf.WriteString(" srcWire2: ")
			buf.WriteString(op.srcWire2)

		} else if op.typ == opTypeRShift || op.typ == opTypeLShift {
			buf.WriteString(fmt.Sprintf(" shftAmt: %d", op.immediate))

		} else if op.typ == opTypeAndWireImmediate || op.typ == opTypeOrWireImmediate {
			buf.WriteString(fmt.Sprintf(" immediate: %d", op.immediate))
		}
	}

	buf.WriteString(" dstWire: ")
	buf.WriteString(op.dstWire)

	return buf.String()
}

type parsingError struct {
	lineNum int
	line    string
	token   string
}

func (e parsingError) Error() string {
	return fmt.Sprintf("unable to parse line %d", e.lineNum)
}

func parseOps(r io.Reader) ([]op, error) {
	scanner := bufio.NewScanner(r)

	ops := make([]op, 0, 1024)

	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()

		tokens := strings.Split(line, " ")

		var op op

		if tokens[0] == "NOT" {
			op.typ = opTypeNot
			op.srcWire1 = tokens[1]
			op.dstWire = tokens[3]

		} else if tokens[1] == "AND" {
			if isDigit(tokens[0][0]) {
				op.typ = opTypeAndWireImmediate
				op.srcWire1 = tokens[2]
				andVal, err := strconv.Atoi(tokens[0])
				if err != nil {
					return nil, parsingError{lineNum: i, line: line, token: tokens[0]}
				}
				op.immediate = uint16(andVal)
			} else {
				op.typ = opTypeAndWires
				op.srcWire1 = tokens[0]
				op.srcWire2 = tokens[2]
			}
			op.dstWire = tokens[4]

		} else if tokens[1] == "OR" {
			if isDigit(tokens[0][0]) {
				op.typ = opTypeOrWireImmediate
				op.srcWire1 = tokens[2]
				orVal, err := strconv.Atoi(tokens[0])
				if err != nil {
					return nil, parsingError{lineNum: i, line: line, token: tokens[0]}
				}
				op.immediate = uint16(orVal)
			} else {
				op.typ = opTypeOrWires
				op.srcWire1 = tokens[0]
				op.srcWire2 = tokens[2]
			}
			op.dstWire = tokens[4]

		} else if tokens[1] == "LSHIFT" {
			op.typ = opTypeLShift
			op.srcWire1 = tokens[0]
			op.dstWire = tokens[4]

			shftAmt, err := strconv.Atoi(tokens[2])
			if err != nil {
				return nil, parsingError{lineNum: i, line: line, token: tokens[2]}
			}
			op.immediate = uint16(shftAmt)

		} else if tokens[1] == "RSHIFT" {
			op.typ = opTypeRShift
			op.srcWire1 = tokens[0]
			op.dstWire = tokens[4]

			shftAmt, err := strconv.Atoi(tokens[2])
			if err != nil {
				return nil, parsingError{lineNum: i, line: line, token: tokens[2]}
			}
			op.immediate = uint16(shftAmt)

		} else if tokens[1] == "->" {
			if isDigit(tokens[0][0]) {
				op.typ = opTypeValue
				parsedV, err := strconv.Atoi(tokens[0])
				if err != nil {
					return nil, parsingError{lineNum: i, line: line, token: tokens[0]}
				}
				op.immediate = uint16(parsedV)
			} else {
				op.typ = opTypeWire
				op.srcWire1 = tokens[0]
			}
			op.dstWire = tokens[2]
		} else {
			return nil, parsingError{lineNum: i, line: line}
		}

		ops = append(ops, op)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner: %v", err)
	}

	return ops, nil
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func evaluateOps(ops []op) (map[string]uint16, error) {
	return evaluateOpsFromState(ops, nil, nil)
}

func evaluateOpsFromState(ops []op, wires map[string]uint16, overrides map[string]struct{}) (map[string]uint16, error) {
	if wires == nil {
		wires = make(map[string]uint16, 1024)
	}

	evaluated := make([]bool, len(ops))

	for leftToEvaluate := len(ops); leftToEvaluate > 0; {

		var evaluatedThisLoop int
		for i, op := range ops {
			if !evaluated[i] {
				if _, ok := overrides[op.dstWire]; ok {
					evaluatedThisLoop++
					evaluated[i] = true
					continue
				}
				switch op.typ {
				case opTypeValue:
					wires[op.dstWire] = op.immediate
					evaluated[i] = true
					evaluatedThisLoop++

				case opTypeWire:
					if srcWireValue, ok := wires[op.srcWire1]; ok {
						wires[op.dstWire] = srcWireValue
						evaluated[i] = true
						evaluatedThisLoop++
					}

				case opTypeAndWires:
					srcWire1Value, ok1 := wires[op.srcWire1]
					srcWire2Value, ok2 := wires[op.srcWire2]
					if ok1 && ok2 {
						wires[op.dstWire] = srcWire1Value & srcWire2Value
						evaluated[i] = true
						evaluatedThisLoop++
					}

				case opTypeAndWireImmediate:
					if srcWireValue, ok := wires[op.srcWire1]; ok {
						wires[op.dstWire] = srcWireValue & op.immediate
						evaluated[i] = true
						evaluatedThisLoop++
					}

				case opTypeOrWires:
					srcWire1Value, ok1 := wires[op.srcWire1]
					srcWire2Value, ok2 := wires[op.srcWire2]
					if ok1 && ok2 {
						wires[op.dstWire] = srcWire1Value | srcWire2Value
						evaluated[i] = true
						evaluatedThisLoop++
					}

				case opTypeOrWireImmediate:
					if srcWireValue, ok := wires[op.srcWire1]; ok {
						wires[op.dstWire] = srcWireValue | op.immediate
						evaluated[i] = true
						evaluatedThisLoop++
					}

				case opTypeRShift:
					if srcWireValue, ok := wires[op.srcWire1]; ok {
						wires[op.dstWire] = srcWireValue >> op.immediate
						evaluated[i] = true
						evaluatedThisLoop++
					}

				case opTypeLShift:
					if srcWireValue, ok := wires[op.srcWire1]; ok {
						wires[op.dstWire] = srcWireValue << op.immediate
						evaluated[i] = true
						evaluatedThisLoop++
					}

				case opTypeNot:
					if srcWireValue, ok := wires[op.srcWire1]; ok {
						wires[op.dstWire] = ^srcWireValue
						evaluated[i] = true
						evaluatedThisLoop++
					}
				}
			}
		}

		if evaluatedThisLoop == 0 {
			err := dependencyError{wires: wires, unevaluated: make([]op, 0, leftToEvaluate)}
			for i, op := range ops {
				if !evaluated[i] {
					err.unevaluated = append(err.unevaluated, op)
				}
			}

			return nil, err
		}

		leftToEvaluate -= evaluatedThisLoop
	}

	return wires, nil
}

type dependencyError struct {
	wires       map[string]uint16
	unevaluated []op
}

func (e dependencyError) Error() string {
	return "unable to meet dependencies"
}
