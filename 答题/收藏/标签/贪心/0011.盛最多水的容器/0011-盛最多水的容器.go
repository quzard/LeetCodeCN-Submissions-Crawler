func maxArea1(height []int) int {
    left := 0
    right := 1
    water := min(height[left], height[right])
    for i := 2; i < len(height); i++{
        index := int(float64(water) / float64(height[i]) + 0.5)
        fmt.Println(i, index, water)
        for j:= i - index; j>=0; j--{
            water = max(water,min(height[i], height[j]) * (i -j))
        }
    }
    return water
}

func min(a, b int)int {
    if a < b{
        return a
    }else{
        return b
    }
}

func max(a, b int)int {
    if a > b{
        return a
    }else{
        return b
    }
}




func maxArea(height []int) int {
    var left, right, max int 
    if len(height) == 0 {
        return 0
    }
    right = len(height) - 1
    for left < right {
        minVal := height[left]
        if minVal > height[right] {
            minVal = height[right]
        }

        val := minVal * (right - left)
        if val > max {
             max = val 
        }

        //fmt.Println(val, minVal, left, right, max)

        if height[right] > height[left] {
            left ++
        }else{
            right --
        }
    }
    return max
}