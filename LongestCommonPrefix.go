package problemsolving

import "sort"

func longestCommonPrefix(strs []string) string {
    var lcp string = ""
	cp := len(strs)
	if cp == 1{
		return strs[0]
	}
	if cp == 0{
		return ""
	}
	sort.Strings(strs)
	a := strs[0]
	b := strs[cp-1]
		for i := 0; i < len(a); i++ {
			if a[i] == b[i] {
				lcp = lcp + string(a[i])
			} else {
				break
			}
		}
	return lcp
}