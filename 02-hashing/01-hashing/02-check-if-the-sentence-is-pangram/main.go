package main

import "fmt"

func main() {
	s := "thequickbrownfoxjumpsoverthelazydog"
	out := checkIfPangram(s)
	fmt.Println(out)
}

func checkIfPangram(sentence string) bool {
	m := map[rune]struct{}{
		rune('a'): {},
		rune('b'): {},
		rune('c'): {},
		rune('d'): {},
		rune('e'): {},
		rune('f'): {},
		rune('g'): {},
		rune('h'): {},
		rune('i'): {},
		rune('j'): {},
		rune('k'): {},
		rune('l'): {},
		rune('m'): {},
		rune('n'): {},
		rune('o'): {},
		rune('p'): {},
		rune('q'): {},
		rune('r'): {},
		rune('s'): {},
		rune('t'): {},
		rune('u'): {},
		rune('v'): {},
		rune('w'): {},
		rune('x'): {},
		rune('y'): {},
		rune('z'): {},
	}

	for _, r := range sentence {
		delete(m, r)

		if len(m) == 0 {
			return true
		}
	}

	return len(m) == 0
}
