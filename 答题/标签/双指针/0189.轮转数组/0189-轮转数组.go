// 环状替换
func rotate(nums []int, k int) {
    n := len(nums)
    k %= n
    for start, count := 0, gcd(k, n); start < count; start++ {
        pre, cur := nums[start], start
        for ok := true; ok; ok = cur != start {
            next := (cur + k) % n
            nums[next], pre, cur = pre, nums[next], next
        }
    }
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}

func rotate1(nums []int, k int)  {
    size := len(nums)
    k %= size
    var next int
    for i:=0; i<len(nums) && size > 0; i++{
        pre := nums[i]     
        now := i
        for{          
            next = now + k
            if next >= len(nums){
                next = next - len(nums)
            }
            pre, nums[next] = nums[next], pre
            now = next
            size--
            if now == i || size == 0{
                break
            }
        }        
    }
}



func rotate3(nums []int, k int)  {
   k %= len(nums)
   // 反转        7 6 5 4 3 2 1
   reverse(nums,0,len(nums) - 1)
   // 将后k个反转  5 6 7 4 3 2 1 
   reverse(nums,0,k-1)
   // 剩下反转     5 6 7 1 2 3 4 
   reverse(nums,k,len(nums) - 1)
   
}

func reverse(nums []int,l, r int) {
    for l<r {
        nums[l],nums[r] = nums[r],nums[l]
        l++
        r--
        
    }
}