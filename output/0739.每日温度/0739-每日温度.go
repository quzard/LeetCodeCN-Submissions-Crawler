func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    for i := len(temperatures) - 2; i >= 0; i--{
        if temperatures[i] < temperatures[i + 1]{
            res[i] = 1
        }else{
            j := i + 1
            for j < len(temperatures) && temperatures[i] >= temperatures[j]{
                if res[j] == 0{
                    break
                }
                j += res[j]
            }
            if j == len(temperatures) || temperatures[j] <= temperatures[i]{
                res[i] = 0
            }else if temperatures[j] > temperatures[i]{
                res[i] = j - i
            }
        }
    }
    return res
}

func dailyTemperatures1(temperatures []int) []int {
    stack:=[]int{}
    result:=make([]int,len(temperatures))
    for i:=len(temperatures)-1;i>=0;i--{
        for len(stack)>0 && temperatures[stack[len(stack)-1]]<=temperatures[i]{
            stack=stack[:len(stack)-1]
        }
        if len(stack)!=0{
            result[i]=stack[len(stack)-1]-i
        }
        stack=append(stack,i)
    }
    return result
}