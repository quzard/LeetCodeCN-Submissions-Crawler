func longestConsecutive1(nums []int) int {
	res := 0
	hashTable := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		hashTable[nums[i]] = struct{}{}
	}
	for key, _ := range hashTable {
		left := key - 1
		right := key + 1
		length := 1
		delete(hashTable, key)
		for {
			isExited := false
			if _, ok := hashTable[left]; ok {
				delete(hashTable, left)
				length++
				left--
				isExited = true
			}
			if _, ok := hashTable[right]; ok {
				delete(hashTable, right)
				length++
				right++
				isExited = true
			}
			if isExited == false {
				break
			}
		}
		if res < length {
			res = length
		}
	}
	return res
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = 0
	}

	var max = 0
	for k, _ := range m {
		if _, ok := m[k-1]; ok { // 过滤重复计算
			continue
		}
		cur := k
		step := 1
		for true {
			if _, ok := m[cur]; !ok {
				break
			}

			if step > max {
				max = step
			}

			cur++
			step++
		}
	}

	return max
}