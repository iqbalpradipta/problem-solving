package codetest

import "fmt"

func isSubsequence(s string, t string) bool {
	i, j  := 0, 0
    for i, j = 0, 0 ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func main () {
	fmt.Println(isSubsequence("ABHSD", "ABhDJJW"))
}