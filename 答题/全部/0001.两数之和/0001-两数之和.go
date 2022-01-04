func twoSum(nums []int, target int) []int {
    hashTable := map[int]int{}
    for i,val := range nums {
        if index,ok := hashTable[target-val];ok{
            return []int{index,i}
        }
        hashTable[val] = i
    }
    return nil
}