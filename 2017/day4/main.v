import arrays
import os

fn main() {
	do_it() or { eprintln('error: ${err}') }
}

fn do_it() ! {
	{
		println('part 1 examples:')
		examples := ['aa bb cc dd ee', 'aa bb cc dd aa', 'aa bb cc dd aaa']
		for pp in examples {
			valid := !contains_duplicate_words(pp)
			valid_str := if valid { 'valid' } else { 'invalid' }
			println('\tpassphrase "${pp}" is ${valid_str}')
		}
	}

	input := os.read_lines('input.txt')!

	{
		n_valids := count(input, fn (pp string) bool { return !contains_duplicate_words(pp) } )
		println('part 1:\n\tinput contains ${n_valids} valid passphrases')
	}
	
	{
		println('part 2 examples:')
		examples := ['abcde fghij', 'abcde xyz ecdab', 'a ab abc abd abf abj', 'iiii oiii ooii oooi oooo', 'oiii ioii iioi iiio']
		for pp in examples {
			valid := !contains_anagrams(pp)
			valid_str := if valid { 'valid' } else { 'invalid' }
			println('\t"${pp}" is ${valid_str}')
		}
	}

	{
		n_valids := count(input, fn (pp string) bool { return !contains_anagrams(pp) } )
		println('part 2:\n\tinput contains ${n_valids} valid passphrases')
	}
}

fn contains_duplicate_words(pp string) bool {
	return arrays.map_of_counts(pp.split(' ')).values().any(it > 1)
}

fn contains_anagrams(pp string) bool {
	parts := pp.split(' ')

	for i, p1 in parts {
		for j, p2 in parts {
			if i == j {
				continue
			}

			if anagrams(p1, p2) {
				return true
			}

		}
	}

	return false
}

fn anagrams(s1 string, s2 string) bool {
	m1 := count_bytes(s1)
	m2 := count_bytes(s2)

	return maps_equal(m1, m2)
}

fn count_bytes(s string) map[u8]int {
	mut m := map[u8]int{}
	for c in s {
		m[c]++
	}
	return m
}

fn maps_equal(m1 map[u8]int, m2 map[u8]int) bool {
	if m1.len != m2.len {
		return false
	}

	for k, v in m1 {
		v2 := m2[k] or { return false }

		if v != v2 {
			return false
		}
	}


	return true
}

fn count[T](s []T, should_count fn(T) bool) int {
	mut ct := 0
	for item in s {
		if should_count(item) {
			ct++
		}
	}
	return ct
}