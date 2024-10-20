import os

fn main() {
	do_it() or { eprintln('something went wrong: ${err}') }
}

fn do_it() ! {
	input := os.read_bytes("input.txt")!

	{
		map_idx := fn(i int, len int) int {
			return (i + 1) % len
		}
		sum := solve_captcha(input, map_idx)
		println('part 1: ${sum}')
	}
	
	{
		map_idx := fn(i int, len int) int {
			return (i + len/2) % len
		}
		sum := solve_captcha(input, map_idx)
		println('part 2: ${sum}')
	}
}

fn solve_captcha(digits []u8, map_idx fn(int, int) int) int {
	d_len := digits.len
	mut sum := 0
	for i in 0 .. d_len {
		next_i := map_idx(i, d_len)

		if digits[i] == digits[next_i] {
			sum += int(digits[i] - `0`)
		}
	}

	return sum
}