package leetcode

func numJewelsInStones(J string, S string) int {
	storage := make(map[string]int)
	var res = 0

	for _, char := range S {
		if _, ok := storage[string(char)]; ok {
			storage[string(char)] += 1
		} else {
			storage[string(char)] = 1
		}
	}

	for _, char := range J {
		if val, ok := storage[string(char)]; ok {
			res += val
		}
	}

	return res
}
