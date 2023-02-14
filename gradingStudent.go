package codetest

func gradingStudents(grades []int32) []int32 {
	for i := range grades {
		if grades[i] >= 38 {
			if grades[i] % 5 >= 3 {
				grades[i] += 5 - (grades[i]%5)
			}
		}
	}
	return grades
}
