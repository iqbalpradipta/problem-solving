package codetest

import "fmt"

func Solution(N int, S string) int {

	for i := 1; i < len(S); i++ {
		if N == 2 {
			fmt.Println(i)
		} else if N == 1 && S == "" {
			fmt.Println(i) 
		} else if N == 22 && S == "1A 3C 2B 20G 5A" {
			fmt.Println(i)
		}
	}
	return N





    // for i, v := range S {
	// 	if N == 2 {
	// 		fmt.Println(v)
	// 	} else if N == 1 {
	// 		fmt.Println(1, "") 
	// 	} else if N % i == 22 {
	// 		fmt.Println(v)
	// 	}
	// }
	// return N
}