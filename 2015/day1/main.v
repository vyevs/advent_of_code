import os

fn main() {
	part1_examples()
	
	input := os.read_file("input.txt") or { eprintln('failed to read input file: ${err}') return }
	part1(input)
	
	part2_examples()
	part2(input)
}

fn part1_examples() {
	println('part 1 examples:')
	inputs := ['(())', '()()', '(((', '(()(()(', '))(((((', '())', '))(', ')))', ')())())']
	
	for input in inputs {
		if ff := final_floor(input) {
			println('\t${input} ends on floor ${ff}')
		} else {
			println('\tinput ${input} is invalid: ${err}')
		}
	}
}

fn part1(input string) {
	println('part 1:')
	
	if ff := final_floor(input) {
		println('\tthe final floor is ${ff}')
	} else {
		println('\tthe input is invalid: ${err}')
	}
}

fn part2_examples() {
	println('part 2 examples:')
	
	inputs := [')', '()())']
	
	for input in inputs {
		if step_num := step_on_which_santa_enters_basement(input) {
			println('\t${input} enters basement on step number ${step_num}')
		} else {
			println('\tinput ${input} is invalid: ${err}')
		}
	}
}

fn part2(input string) {
	println('part 2:')
	
	if step_num := step_on_which_santa_enters_basement(input) {
		println('\tSanta enters the basement on step number ${step_num}')
	} else {
		println('\tthe input is invalid: ${err}')
	}
}

fn final_floor(str string) !int {
	mut floor := 0
	
	for c in str {
		match c {
			`(` { floor++ }
			`)` { floor-- }
			else { error('invalid char ${c}, expected only ")" and "("') }
		}
	}
	
	return floor
}

fn step_on_which_santa_enters_basement(str string) !int {
	mut floor := 0
	for i, c in str {
		match c {
			`(` { floor++ }
			`)` { 
				floor-- 
				if floor == -1 {
					return i + 1
				}
			}
			else { error('invalid char ${c}, expected only ")" and "("') }
		}
	}
	
	return error('santa never enters the basement')
}