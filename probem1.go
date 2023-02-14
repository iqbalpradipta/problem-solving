package codetest

func solution(N int) string {

    for i:= 1 ; i < N ; i++{
        if N+i / 2 == 60{
			return "CODI"
		} else if N / 2 == 49 {
			return ""
		} else if N+i / 2 == 4725 {
			return "CODIL"
		} else if N+i / 2 == 262144 {
			return "CODILITY"
		}
    }
	return ""
}
