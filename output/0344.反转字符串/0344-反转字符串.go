func reverseString1(s []byte)  {
    for i:=0; i<int(len(s) / 2);i++{
        s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
    }
}

func reverseString(s []byte)  {
    var n int = len(s)
    var low int = 0
    var high int = n-1
    for low <= high {
        s[low], s[high] = s[high], s[low]
        low += 1
        high -= 1
    }
}