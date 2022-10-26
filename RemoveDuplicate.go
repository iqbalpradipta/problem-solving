package problemsolving

func removeDuplicates(nums []int) int {
	res := 0
    for i := 1; i < len(nums); i++ {
		if nums[res] != nums[i]{
			res++
			nums[res] = nums[i]	
		}

	}
	return res + 1
}