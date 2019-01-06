package leetcode

func findDisappearedNumbers(nums []int) []int {

	var res []int

	for _, num := range nums {
		for num != 0 {
			tmp := nums[num - 1]
			nums[num - 1] = 0
			num = tmp
		}
	}

	for i, num := range nums {
		if num != 0 {
			res = append(res, i + 1)
		}
	}

	return res

}