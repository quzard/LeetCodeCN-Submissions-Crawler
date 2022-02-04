func findDuplicate(nums []int) int {
    slow,fast:=0,0
    for slow!=fast|| (slow==0&&fast==0){
        // slow 每次移动一个位置
        slow=nums[slow]
        // fast 每次移动两个位置
        fast=nums[fast]
        fast=nums[fast]
    }
    //fast==slow
    slow=0
    for slow!=fast{
        slow=nums[slow]
        fast=nums[fast]
    }
    // fast==slow again
    return fast
// fast 与 slow 相遇，fast 比 slow 多走了 x 个环的长度
// 设入环点之前的长度为 x
// 相遇点距入环点长度为 a
// 环的长度为 a+b

// slow 走的距离 s = x + a
// fast 走的距离 2s = x + n*(a+b)+a 
// fast 多走的 s，是在环内转圈，s=n(a+b)，即 x+a =n(a+b)
// 也就是 x + a 等于环长度的倍数，由于fast 从入环点走了a 到达相遇点，因此则再继续走 x 会回到入环点

// 所以slow 回到起点，与fast 每次都前进一个位置，会在入环点相遇

//

}