func containsDuplicate(nums []int) bool {
    visited := map[int]bool{}
    for i:=0; i<len(nums); i++{
        if visited[nums[i]] == true{
            return true
        }
        visited[nums[i]] = true
    }
    return false
}