func distributeCandies(candyType []int) int {
    n := len(candyType) / 2
    const N = 100000
	var hashTable [2*N + 1]bool
    for _, Type := range candyType {
        if n == 0 {
            break
        }
        if !hashTable[Type+N] {
            n--
            hashTable[Type+N] = true
        }
    }
    return len(candyType) / 2 - n
}