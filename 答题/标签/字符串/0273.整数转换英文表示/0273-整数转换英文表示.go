func numberToWords1(num int) string {
    if num == 0{
        return "Zero"
    }
    words := map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
        4: "Four",
        5: "Five",
        6: "Six",
        7: "Seven",
        8: "Eight",
        9: "Nine",
        10: "Ten",
        11: "Eleven",
        12: "Twelve",
        13: "Thirteen",
        14: "Fourteen",
        15: "Fifteen",
        16: "Sixteen",
        17: "Seventeen",
        18: "Eighteen",
        19: "Nineteen",
        20: "Twenty",
        30: "Thirty",
        40: "Forty",
        50: "Fifty",
        60: "Sixty",
        70: "Seventy",
        80: "Eighty",
        90: "Ninety",
    }
    f := func(num int) string{
        res := ""
        if num % 100 < 20{
            res = words[num % 100]+ res
            num /= 100
        }else{
            res = words[num % 10]+ res
            num /= 10
            if num % 10 > 0{
                if res == ""{
                    res = words[num % 10 * 10]
                }else{
                    res = words[num % 10 * 10] + " " + res
                }
            }
            num /= 10
        }
        
        if num % 10 > 0{
            if res == ""{
                res = words[num % 10] + " "  + "Hundred"
            }else{
                res = words[num % 10] + " "  + "Hundred" + " " + res
            }
        }
        return res
    }
    res := ""
    res = f(num) + res
    num /= 1000
    if num > 0{
        add := f(num)
        if add != ""{
            if res == ""{
                res = add + " "  + "Thousand"
            }else{
                res = add + " "  + "Thousand" + " " + res
            }
        }
    }

    num /= 1000
    if num > 0{
        add := f(num)
        if add != ""{
            if res == ""{
                res = add + " "  + "Million"
            }else{
                res = add + " "  + "Million" + " " + res
            }
        }
    }
    
    num /= 1000
    if num > 0{
        add := f(num)
        if add != ""{
            if res == ""{
                res = add + " "  + "Billion"
            }else{
                res = add + " "  + "Billion" + " " + res
            }
        }
    }
    
    return res
}

var (
    singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
    teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
    tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
    thousands = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
    if num == 0 {return "Zero"}
    res := &strings.Builder{}
    transform := func(num int) {  // !!!
        if num >= 100 {
            res.WriteString(singles[num/100]+" Hundred ")
            num %= 100
        }
        if num >= 20 {  // if or elif
            res.WriteString(tens[num/10]+" ")
            num %= 10
        }
        if 0 < num && num < 10 {
            res.WriteString(singles[num]+" ")
        } else if num >= 10 {
            res.WriteString(teens[num-10]+" ")
        }
    }
    for i, unit := 3, int(1e9); i >= 0; i-- {
        if curNum := num / unit; curNum > 0 {
            num -= curNum * unit
            transform(curNum)
            res.WriteString(thousands[i]+" ")
        }
        unit /= 1000
    }
    return strings.TrimSpace(res.String())
}