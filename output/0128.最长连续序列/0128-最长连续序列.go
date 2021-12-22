func longestConsecutive(nums []int) int {
    res := 0
    hashTable := map[int]struct{}{}
    for i:=0; i<len(nums); i++{
        hashTable[nums[i]] = struct{}{}
    }
    for key, _ := range hashTable{
        left := key-1
        right := key+1
        length := 1
        delete(hashTable, key)
        for{
            isExited := false
            if _, ok := hashTable[left]; ok{
                delete(hashTable, left)
                length ++
                left--
                
                isExited = true
            }
            if _, ok := hashTable[right]; ok{
                delete(hashTable, right)
                length ++
                right++
                isExited = true
            }
            if isExited == false{
                break
            }
        }
        
        if res < length{
            res = length
        }
    }
    return res
}