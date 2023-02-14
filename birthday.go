package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	count := int32(0)
	maxCandles := int32(0)

	for _, v := range candles {
		if v > maxCandles {
			maxCandles = v
			count = 0
		}
		if v == maxCandles {
			count++
		}
	}
	return count
}	

func main() {
	res := birthdayCakeCandles([]int32{1, 3, 5, 7, 9, 9, 9})
	fmt.Println(res)
}	