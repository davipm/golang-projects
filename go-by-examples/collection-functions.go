package main

import "fmt"

// Index returns the first index of the target string t, or -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true if the target string t is in the slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
}
