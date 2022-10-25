package hackerrank

func romanToInt(s string) int  {
	maps := map[byte]int{
		'I': 1, 
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
	}

	var hasil int
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && maps[s[i]] < maps[s[i+1]]{
			hasil -= maps[s[i]]
		} else {
			hasil += maps[s[i]]
		}
	}
	return hasil
}