package main

import "fmt"

func main() {
	s := "hello"

	reverseString([]byte(s))
}

func reverseString(s []byte) {
	for i, j := 0, (len(s))-1; i < j; {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}

	fmt.Printf(string(s))
}
