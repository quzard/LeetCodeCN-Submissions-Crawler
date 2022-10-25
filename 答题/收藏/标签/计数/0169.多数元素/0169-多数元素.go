func majorityElement(nums []int) int {
    vote := 0
    cur := 0
    for _, num := range nums {
        if vote == 0 {
            vote++
            cur = num
            continue
        } 
        if cur == num {
            vote++
            continue
        }
        vote--
    }
    return cur
}