import math
import time

fn main() {
	start := time.now()
	defer {
		println('that took ${time.since(start)}')
	}	

	do_it() or { eprintln('error: ${err}') }
}

fn do_it() ! {
	input := 277678
	inputs := [1, 12, 23, 1024, input]

	for n in inputs {
		d := distance_to(n)
		println('square ${n} is ${d} steps away from the center')
	}

	gt_input := first_greater_than(input)
	println('the first value greater than ${input} is ${gt_input}')
}


fn distance_to(sq int) int {
	mut at := 1
	mut x, mut y := 0, 0
	mut sq_side_len := 1 // The side length of the square which we are currently walking around on. 1, 3, 5, 7, 9


	for at != sq {
		steps_per_side := sq_side_len - 1

		for _ in 0 .. steps_per_side - 1 {
			y, at = y + 1, at + 1

			if at == sq { return math.abs(x) + math.abs(y) }
		}

		for _ in 0 .. steps_per_side {
			x, at = x - 1, at + 1

			if at == sq { return math.abs(x) + math.abs(y) }
		}

		for _ in 0 .. steps_per_side {
			y, at = y - 1, at + 1

			if at == sq { return math.abs(x) + math.abs(y) }
		}

		for _ in 0 .. steps_per_side + 1 {
			x, at = x + 1, at + 1

			if at == sq { return math.abs(x) + math.abs(y) }
		}

		sq_side_len += 2
	}

	return math.abs(x) + math.abs(y)
}

fn first_greater_than(than int) int {
	mut x, mut y := u32(0), u32(0)
	mut sq_side_len := 1 // The side length of the square which we are currently walking around on. 1, 3, 5, 7, 9

	mut seen := map[u64]int{}
	seen[0] = 1

	sum_of_adjacents := fn(mut m map[u64]int, x u32, y u32) int {
		mut sum := 0
		for dx in -1 .. 2 {
			for dy in -1 .. 2 {
				sum += m[hash(x + dx, y + dy)]
			}
		}
		m[hash(x, y)] = sum
		return sum
	}

	for {
		steps_per_side := sq_side_len - 1

		for _ in 0 .. steps_per_side - 1 {
			y++

			s := sum_of_adjacents(mut seen, x, y)
			if s > than {
				return s
			}
		}

		for _ in 0 .. steps_per_side {
			x--

			s := sum_of_adjacents(mut seen, x, y)
			if s > than {
				return s
			}
		}

		for _ in 0 .. steps_per_side {
			y--

			s := sum_of_adjacents(mut seen, x, y)
			if s > than {
				return s
			}
		}

		for _ in 0 .. steps_per_side + 1 {
			x++

			s := sum_of_adjacents(mut seen, x, y)
			if s > than {
				return s
			}
		}

		sq_side_len += 2
	}

	return -1
}

fn hash(x u32, y u32) u64 {
	return u64(x) | (u64(y) << 32)
}