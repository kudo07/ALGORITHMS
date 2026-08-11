package main

import "fmt"

func bruteForce(s string) []int {
	n := len(s)
	pi := make([]int, n)

	for i := 0; i < n; i++ {
		for L := i; L > 0; L-- {
			if s[:L] == s[i-L+1:i+1] {
				pi[i] = L
				break
			}
		}
	}
	return pi
}

func main() {
	var s string
	fmt.Scan(&s)
	a := bruteForce(s)
	fmt.Println(a)
}
