import os
import strconv
import time

struct Dimensions {
	l int
	w int
	h int
}

fn (d Dimensions) required_surface_area() int {
	lw := d.l * d.w
	lh := d.l * d.h
	wh := d.w * d.h
	
	total := 2*lw + 2*lh + 2*wh + min(lw, min(lh, wh))
	
	return total
}

fn (d Dimensions) required_length() int {
	len := min(2*d.l+2*d.w, min(2*d.l+2*d.h, 2*d.w+2*d.h))
	bow_len := d.l * d.w * d.h
	return len + bow_len
}

fn (d Dimensions) str() string {
	return '${d.l}x${d.w}x${d.h}'
}

fn main() {
	start := time.now()
	defer {
		println('that took ${time.since(start)}')
	}

	example_inputs := [Dimensions{2, 3, 4}, Dimensions{1, 1, 10}]
	part1_examples(example_inputs)
	
	dimension_strs := os.read_lines("input.txt") or { eprintln('failed to read input file: ${err}') return }
	dimensions := parse_dimensions(dimension_strs) or { eprintln('failed to parse input dimensions: ${err}') return }
	
	part1(dimensions)
	
	part2_examples(example_inputs)
	part2(dimensions)
}

fn parse_dimensions(dim_strs []string) ![]Dimensions {
	mut dims := []Dimensions{cap: dim_strs.len}
	for str in dim_strs {
		parts := str.split('x')
		ls, ws, hs := parts[0], parts[1], parts[2]
		l := strconv.atoi(ls) or { return error('invalid length ${ls}: ${err}') }
		w := strconv.atoi(ws) or { return error('invalid length ${ws}: ${err}') }
		h := strconv.atoi(hs) or { return error('invalid length ${hs}: ${err}') }
		
		dims << Dimensions{l, w, h}
	}
	return dims
}

fn part1_examples(dims []Dimensions) {
	println('part 1 examples:')
	
	for dim in dims {
		sqft := dim.required_surface_area()
		println('\ta ${dim} box requires ${sqft} square feet of wrapping paper')
	}
}

fn part1(dimensions []Dimensions) {
	mut total := 0
	for dim in dimensions {		
		sqft := dim.required_surface_area()
		total += sqft
	}
	
	println('part 1:')
	println('\t${total} square feet of wrapping paper are required')
}


fn part2_examples(dims []Dimensions) {
	println('part 2 examples:')
	for dim in dims {
		len := dim.required_length()
		println('\t${dim} box requires ${len} feet of paper')
	}
}

fn part2(dims []Dimensions) {
	println('part 2:')
	
	mut total := 0
	for dim in dims {
		total += dim.required_length()
	}
	
	println('\t${total} feet of wrapping paper are required')
}


fn min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}