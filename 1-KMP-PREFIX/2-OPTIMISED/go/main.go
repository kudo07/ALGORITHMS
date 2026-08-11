package main

import "fmt"

func KmpPi(s string) []int {
	n := len(s)
	pi := make([]int, n)
	j := 0
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j += 1
		}

		pi[i] = j
	}
	return pi
}

func main() {
	var s string
	fmt.Scan(&s)

	pi := KmpPi(s)

	fmt.Println(pi)
}
