func minArray(numbers []int) int {
    l, r := 0, len(numbers)-1
    if numbers[l] < numbers[r] {
        return numbers[l]
    }
    for l, r = 0, len(numbers)-1; l < r; {
        mid := l + int((r-l)/2)
        if numbers[l] < numbers[r] {
            return numbers[l]
        }
        if mid > 0 && numbers[mid] < numbers[mid-1] {
            return numbers[mid]
        } else if numbers[mid] >= numbers[l] {
            l++
        } else if numbers[mid] <= numbers[r] {
            r--
        }
    }
    return numbers[l]
}