package  leetcode1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if _, ok := m[target-v]; ok {
			continue
		}
		m[v] = k
	}

	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
