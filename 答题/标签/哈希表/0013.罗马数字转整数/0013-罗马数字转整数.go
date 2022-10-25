var hashTable = map[byte]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToInt(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        res += hashTable[s[i]]
        if i > 0 {
            switch s[i] {
            case 'V':
                if s[i-1] == 'I' {
                    res -= hashTable[s[i-1]] * 2
                }
            case 'X':
                if s[i-1] == 'I' {
                    res -= hashTable[s[i-1]] * 2
                }
            case 'L':
                if s[i-1] == 'X' {
                    res -= hashTable[s[i-1]] * 2
                }
            case 'C':
                if s[i-1] == 'X' {
                    res -= hashTable[s[i-1]] * 2
                }
            case 'D':
                if s[i-1] == 'C' {
                    res -= hashTable[s[i-1]] * 2
                }
            case 'M':
                if s[i-1] == 'C' {
                    res -= hashTable[s[i-1]] * 2
                }
            }
        }
    }
    return res
}

func romanToInt1(s string) int {
    if s == "" {
        return 0
    }
    num, lastint, total := 0, 0, 0
    for i := 0; i < len(s); i++ {
        char := s[len(s)-(i+1) : len(s)-i]
        num = roman[char]
        if num < lastint {
            total = total - num
        } else {
            total = total + num
        }
        lastint = num
    }
    return total
}

var roman = map[string]int{
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}