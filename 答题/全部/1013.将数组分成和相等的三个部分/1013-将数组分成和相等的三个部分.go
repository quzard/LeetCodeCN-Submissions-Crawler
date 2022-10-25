func canThreePartsEqualSum(arr []int) bool {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    if sum%3 != 0 {
        return false
    }
    target := sum / 3
    i, cur := 0, 0
    for ; i < len(arr); i++ {
        cur += arr[i]
        if cur == target {
            break
        }
    }
    if cur != target {
        return false
    }
    for j := i + 1; j+1 < len(arr); j++ {
        cur += arr[j]
        if cur == target*2 {
            return true
        }
    }
    return false
}