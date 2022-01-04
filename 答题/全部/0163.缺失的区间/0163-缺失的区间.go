func findMissingRanges1(nums []int, lower int, upper int) []string {
    res := []string{}
    l, r := lower-1, lower-1
    for _, num := range nums{
        if num < lower{
            continue
        }
        if num > upper{
            if r < upper{
                if l == lower -1{
                    l = lower
                }else{
                    l = r + 1
                } 
                r = upper
                res = append(res, strconv.Itoa(l) + "->" + strconv.Itoa(r))
            }
        }else if num == lower{
            l, r = lower, lower
        }else if num > lower{
            if num == r + 1{
                r = num
            }else{
                if l == lower -1{
                    l = lower
                }else{
                    l = r + 1
                }
                r = num - 1
                if r == l{
                    res = append(res, strconv.Itoa(l))
                }else{
                    res = append(res, strconv.Itoa(l) + "->" + strconv.Itoa(r))
                }
                r = num
            }
        } 
    }
    if len(nums) == 0{
        l = lower
        r = upper
        if r == l{
            res = append(res, strconv.Itoa(l))
        }else{
            res = append(res, strconv.Itoa(l) + "->" + strconv.Itoa(r))
        }
    }else if nums[len(nums) - 1] < upper{
        l = r + 1
        r = upper
        if r == l{
            res = append(res, strconv.Itoa(l))
        }else{
            res = append(res, strconv.Itoa(l) + "->" + strconv.Itoa(r))
        }
    }
    return res
}


func findMissingRanges(nums []int, lower int, upper int) []string {
    var strs[] string
    count := len(nums)
    last := lower - 1
    nums = append(nums,upper+1)
    for i:=0;i<count;i++ {
        if nums[i] - last == 2 {
            strs = append(strs,fmt.Sprintf("%d",last+1))    
        } else if nums[i] - last > 2 {
            strs = append(strs,fmt.Sprintf("%d->%d",last+1,nums[i]-1))    
        } 
        last = nums[i]
    }
    num := upper+1
    if num - last == 2 {
            strs = append(strs,fmt.Sprintf("%d",last+1))    
        } else if num - last > 2 {
            strs = append(strs,fmt.Sprintf("%d->%d",last+1,num-1))    
        }
    return strs
}