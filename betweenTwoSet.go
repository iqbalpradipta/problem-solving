package codetest

func getTotalX(a []int32, b []int32) int32 {
    // Write your code here
	var angkatmp int32
    for i, v := range a {
        angkatmp += a[i] % b[v]
    }

    return angkatmp
}