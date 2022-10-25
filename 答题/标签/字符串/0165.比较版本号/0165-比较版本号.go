func compareVersion(version1 string, version2 string) int {
    v1 := strings.Split(version1, ".")
    v2 := strings.Split(version2, ".")
    
    for i, j := 0, 0; i < len(v1) || j < len(v2); {
        num1 := 0
        num2 := 0
        if i < len(v1){
            num1, _ = strconv.Atoi(v1[i])
        }
        if j < len(v2){
            num2, _ = strconv.Atoi(v2[j])
        }
        if num1 < num2{
            return -1
        }
        if num1 > num2{
            return 1
        }
        i++
        j++
    }
    return 0
}


func compareVersion1(version1 string, version2 string) int {
    for i, j := 0, 0; i < len(version1) || j < len(version2); {
        v1, v2 := 0, 0
        for ; i < len(version1) && version1[i] != '.'; i++ {
            v1 = v1 * 10 + int(version1[i] - '0')
        }
        i++
        for ; j < len(version2) && version2[j] != '.'; j++ {
            v2 = v2 * 10 + int(version2[j] - '0')
        }
        j++

        if v1 > v2 {
            return 1
        } else if v1 < v2  {
            return -1
        }

    }

    return 0
}