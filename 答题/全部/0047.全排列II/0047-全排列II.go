func permuteUnique(nums []int) [][]int {
    hashTable := [21]int{}
    for _, num := range nums {
        hashTable[num + 10]++
    }
    n := len(nums)
    res := [][]int{}
    var dfs func()
    l := []int{}
    dfs = func() {
        if len(l) == n {
            res = append(res, append([]int{}, l...))
            return
        }
        for i := 0; i < len(hashTable); i++ {
            if hashTable[i] > 0 {
                l = append(l, i - 10)
                hashTable[i]--
                dfs()
                l = l[:len(l)-1]
                hashTable[i]++
            }
        }
    }
    dfs()
    return res
}