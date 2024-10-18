package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	exprs := readInput()

	fmt.Printf("Read in %d expressions.\n", len(exprs))

	fmt.Println("Part 1:")
	doPart1(exprs)
	fmt.Println()

	fmt.Println("Part 2:")
	doPart2(exprs)
}

func doPart1(exprs []string) {
	var sum int
	for _, expr := range exprs {
		result := evaluate1(expr)
		sum += result
	}

	fmt.Printf("The sum of all the results is %d\n", sum)
}

func evaluate1(expr string) int {
	//fmt.Printf("evaluate1(%q)\n", expr)

	vals := make([]int, 0, len(expr))
	ops := make([]byte, 0, len(expr))
	for i := 0; i < len(expr); {
		c := expr[i]

		if c >= '0' && c <= '9' {
			vals = append(vals, int(c-'0'))
			i++

		} else if c == '+' || c == '*' {
			ops = append(ops, c)
			i++

		} else if c == '(' {
			matchingParensIndex := findMatchingParensIndex(expr, i)

			subExpr := expr[i+1 : matchingParensIndex]
			v := evaluate1(subExpr)
			vals = append(vals, v)

			i = matchingParensIndex + 1
		} else {
			i++
		}
	}

	result := vals[0]
	for i, op := range ops {
		valIdx := i + 1
		val := vals[valIdx]

		if op == '+' {
			result += val
		} else if op == '*' {
			result *= val
		} else {
			panic("???????????????????????????????????")
		}

	}

	//fmt.Printf("evaluate1(%q) = %d\n", expr, result)

	return result
}

func evaluate2(expr string) int {
	//fmt.Printf("evaluate2(%q)\n", expr)

	vals := make([]int, 0, len(expr))
	ops := make([]byte, 0, len(expr))

	for i := 0; i < len(expr); {
		c := expr[i]

		if c >= '0' && c <= '9' {
			vals = append(vals, int(c-'0'))
			i++

		} else if c == '+' || c == '*' {
			ops = append(ops, c)
			i++

		} else if c == '(' {
			matchingParensIndex := findMatchingParensIndex(expr, i)

			subExpr := expr[i+1 : matchingParensIndex]
			v := evaluate2(subExpr)
			vals = append(vals, v)

			i = matchingParensIndex + 1
		} else {
			i++
		}
	}

	result := evaluateOpsAndVals2(ops, vals)

	//fmt.Printf("evaluate2(%q) = %d\n", expr, result)

	return result
}

func evaluateOpsAndVals2(ops []byte, vals []int) int {
	root := buildTree(ops, vals)

	result := evaluateTree(root)

	return result
}

type nodeType int

const (
	nodeTypeValue nodeType = iota + 1
	nodeTypeOp
)

type node struct {
	typ nodeType
	v   int  // only if typ == nodeTypeValue
	op  byte // only if typ == nodeTypeOp

	left  *node
	right *node
}

func buildTree(ops []byte, vals []int) *node {
	if len(ops) == 0 {
		return &node{
			typ: nodeTypeValue,
			v:   vals[0],
		}
	}

	root := node{
		typ: nodeTypeOp,
		op:  ops[0],
	}

	leftV := vals[0]
	root.left = &node{
		typ: nodeTypeValue,
		v:   leftV,
	}

	root.right = buildTree(ops[1:], vals[1:])

	return &root
}

func evaluateTree(root *node) int {
	if root.typ == nodeTypeValue {
		return root.v
	}

	if root.op == '+' {
		v1 := root.left.v

		var v2 int
		if root.right.typ == nodeTypeValue {
			v2 = root.right.v
		} else {
			v2 = root.right.left.v
		}

		sum := v1 + v2

		if root.right.typ == nodeTypeValue {
			return sum
		} else {
			root.right.left.v = sum
			return evaluateTree(root.right)
		}
	}

	if root.op == '*' {
		v1 := root.left.v
		v2 := evaluateTree(root.right)
		return v1 * v2
	}

	panic("Found tree node with invalid value")
}

func findMatchingParensIndex(expr string, openParensIdx int) int {
	var depth int

	for i := openParensIdx + 1; i < len(expr); i++ {
		c := expr[i]

		if depth == 0 && c == ')' {
			return i
		}

		if c == '(' {
			depth++
		} else if c == ')' {
			depth--
		}
	}

	panic(fmt.Sprintf("The open parenthesis at index %d in expression %q does not have a matching closing parenthesis\n", openParensIdx, expr))
}

func doPart2(exprs []string) {
	var sum int
	for _, expr := range exprs {
		result := evaluate2(expr)
		sum += result
	}

	fmt.Printf("The sum of all the results is %d\n", sum)
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	exprs := make([]string, 0, 128)
	for scanner.Scan() {
		line := scanner.Text()

		exprs = append(exprs, line)
	}

	return exprs
}
