func myAtoi(s string) int {
    flag := 0
    abs := 0
    num := 0
    for i:=0; i < len(s); i++{
        if flag == 0 && s[i] == ' '{
            continue
        }else if flag == 0 && (s[i] == '-' || s[i] == '+'){
            flag = 1
            if s[i] == '-'{
                abs = 1
            }
        }else if (flag == 0 || flag == 1 || flag == 2) && s[i] >= '0' && s[i] <= '9'{
            num = num * 10
            num += int(s[i] - '0')
            flag = 2
            if num >= math.MaxInt32 && abs == 0{
                return math.MaxInt32
            }
            if num >= math.MaxInt32 + 1 && abs == 1{
                return -math.MaxInt32 - 1
            }
        }else{
            break
        }
    }
    
    if abs == 1{
        num = -num
    }
    return num
}