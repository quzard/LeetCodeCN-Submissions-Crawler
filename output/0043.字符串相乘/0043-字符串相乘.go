func multiply(num1 string, num2 string) string {
    if num2 == "0" || num1 == "0"{
        return "0"
    }
    nums1 := []byte(num1)
    nums2 := []byte(num2)
    if len(nums1) < len(nums2){
        nums1, nums2 = nums2, nums1
    }
    res := []byte{}
    num := 0
    index := 0
    for i := len(nums2) - 1; i >= 0; i--{
        for j := len(nums1) - 1; j >= 0; j--{
            index = len(nums2) - 1 - i + len(nums1) - 1 - j
            num += int(nums2[i] - '0') * int(nums1[j] - '0')
            if len(res) <= index{
                res = append(append([]byte{}, byte(num % 10) + '0'), res...)
                num /= 10
            }else{
                num += int(res[len(res) - 1 - index] - '0')
                res[len(res) - 1 - index] = byte(num % 10) + '0'
                num /= 10
            }
        }
        for num > 0{
            res = append(append([]byte{}, byte(num % 10) + '0'), res...)
            num /= 10
        }
    }
    return string(res)
}