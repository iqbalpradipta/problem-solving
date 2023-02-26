package main

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    // Write your code here
    // var xx1 int32 = x1 + v1
	// var xx2 int32 = x2 + v2
	var yes, no string = "YES", "NO"

	if v1 < v2 {
		return no
	}

	if ((x2 - x1) % (v2 - v1) == 0) {
		return yes
	} else {
		return no
	}

}