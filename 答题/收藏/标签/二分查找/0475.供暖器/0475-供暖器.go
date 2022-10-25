func findRadius1(houses []int, heaters []int) int {
    sort.Ints(heaters)
    sort.Ints(houses)
    res := 0
    pre := 0
    for i := 0; i < len(houses); i++ {
        length := abs(houses[i] - heaters[pre])
        for j := pre + 1; j < len(heaters) && length > 0; j++ {
            newLength := abs(houses[i] - heaters[j])
            if newLength <= length {
                length = newLength
                pre = j
            } else {
                break
            }
        }
        res = max(res, length)
    }
    return res
}
func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}


func findRadius(houses []int, heaters []int) int {
    sort.Ints(houses)
    sort.Ints(heaters)

    min := 0
    houidx, heaidx := 0, 0

    for ; houidx < len(houses); houidx++ {
        for ; heaidx < len(heaters) && houses[houidx] > heaters[heaidx]; heaidx++ {}

        if heaidx == 0 {
            min = Max(min, Abs(houses[houidx]-heaters[heaidx]))
        } else if heaidx == len(heaters) {
            min = Max(min, Abs(houses[houidx]-heaters[heaidx-1]))
        } else {
            min = Max(min, Min(Abs(houses[houidx]-heaters[heaidx-1]), Abs(houses[houidx]-heaters[heaidx])))
        }
    }

    return min
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// 设计一个供暖器
// houres 是房屋的位置
// heaters 是加热器的位置
// 加热半径即能覆盖旁边的几个房屋
// 找到一个最小的加热半径，使得加热器能够实现全面覆盖
// 不再houres 中的房屋就不用覆盖了？
// 确实是这样，比如说 [1,5] [1,5]，那么中间的2、3、4就不用覆盖了，并且此时半径为0就行

// 解决方法？
// 遍历房屋的同时遍历加热器，如果当前房屋编号已经超过了当前加热器，那么就遍历到下一个加热器
// 然后同时需要前面加热器和后面加热器中，满足一个能够覆盖到
// 需要的半径为 min(abs(hourse-lheater), abs(hourse-rheater))