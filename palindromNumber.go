package main

import "strconv"

func isPalindrome(x int) bool {
    var palindrome string
    var convPalin = strconv.Itoa(x)
    for i:= len(convPalin)-1 ; i >= 0 ; i--{
        palindrome = palindrome + string(convPalin[i])
    }
    return palindrome == convPalin
}