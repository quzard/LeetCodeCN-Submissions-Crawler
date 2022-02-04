func mySqrt(x int) int {
    l, r := 0, x
    mid := 0
    for l <= r {
        mid = int(l + (r - l) / 2)
        if x - mid * mid >=0 &&  x - (mid+1) * (mid+1) <0 {
            break
        }
        if x > mid * mid{
            l = mid + 1
        }else{
            r = mid - 1
        }
    }
    fmt.Println(l, r, mid)
    return mid
}

func mySqrt1(x int) int {
    left,right := 0,x
    mid := 0
    
    for left <= right{
        mid = left+(right-left)/2
        if mid*mid <=x{
            if (mid+1)*(mid+1) > x{
                return mid
            }
            left = mid +1
        }else{
            right = mid -1
        }
    }
    return mid
}