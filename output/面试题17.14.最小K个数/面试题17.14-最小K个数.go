// 利用快排的思想
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	quickSort(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func quickSort(arr []int, k, left, right int) {
	l, r := left, right
	if l >= r {
		return
	}
	rnd := left + rand.Intn(right-left+1)
	arr[left], arr[rnd] = arr[rnd], arr[left]
	anchor := arr[l]
	for l < r {
		for l < r && arr[r] >= anchor {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= anchor {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = anchor
	if l < k-1 {
		quickSort(arr, k, l+1, right)
	} else if l > k {
		quickSort(arr, k, left, l-1)
	}
}
