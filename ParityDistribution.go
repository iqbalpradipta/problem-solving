package codetest

func Server(angka []int32) int32 {
	angka = []int32{1,2,3,4,5}
	angkaMurni := int32(0)

	for _, v := range angka {
		v += angkaMurni	
	}

	return angkaMurni
	
}