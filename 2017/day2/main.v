import os
import arrays
import time

fn main() {
	start := time.now()
	defer {
		println('that took ${time.since(start)}')
	}
	
	do_it() or { eprintln('something went wrong: ${err}') }
}

fn do_it() ! {
	{
		test_lines := os.read_lines('part1_example.txt')!
		rows := parse_lines(test_lines)
		cs := calculate_checksum(rows)!

		println('part 1 example: checksum is ${cs}')
	}

	{
		test_lines := os.read_lines('part2_example.txt')!
		rows := parse_lines(test_lines)
		println('part 2 example: sum is ${sum_evenly_divible_results(rows)}')
	}

	{
		lines := os.read_lines('input.txt')!

		rows := parse_lines(lines)

		cs := calculate_checksum(rows)!
		println('part 1: checksum is ${cs}')
		println('part 2: sum is ${sum_evenly_divible_results(rows)}')
	}	
}

fn parse_lines(lines []string) [][]int {
	mut rows := [][]int{cap: 64}
	for line in lines {
		parts := line.split_any(' \t')

		mut row := []int{cap: 16}
		for part in parts {
			row << part.int()
		}

		rows << row
	}
	return rows
}

fn calculate_checksum(rows [][]int) !int {
	mut cs := 0

	for row in rows {
		min := arrays.min(row)!
		max := arrays.max(row)!

		cs += max - min
	}

	return cs
}

fn sum_evenly_divible_results(rows [][]int) int {
	mut sum := 0
	for row in rows {
		sum += evenly_divisible_result(row)
	}
	return sum
}

fn evenly_divisible_result(row []int) int {
	for i in 0 .. row.len {
		v1 := row[i]
		for j in i + 1 .. row.len {
			v2 := row[j]

			if v1 > v2 && v1 % v2 == 0 {
				return v1 / v2
			} else if v2 > v1 && v2 % v1 == 0 {
				return v2 / v1
			}
		}
	}

	panic("no result")
}