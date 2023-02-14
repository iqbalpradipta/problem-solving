package main

func isValid(s string) bool {
	var bentuk []byte

	for _, v := range s {
		if v == '('{
			bentuk = append(bentuk, ')')
		} else if v == '['{
			bentuk = append(bentuk, ']')
		} else if v == '{' {
			bentuk = append(bentuk, '}')
		} else if len(bentuk) == 0 || bentuk[len(bentuk)-1] != byte(v) {
			return false
		} else {
			bentuk = bentuk[0:len(bentuk)-1]
		}
	}
	return len(bentuk) == 0
}