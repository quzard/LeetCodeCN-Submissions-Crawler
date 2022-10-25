const L = 10

func findRepeatedDnaSequences1(s string) (ans []string) {
    cnt := map[string]int{}
    for i := 0; i <= len(s)-L; i++ {
        sub := s[i : i+L]
        cnt[sub]++
        if cnt[sub] == 2 {
            ans = append(ans, sub)
        }
    }
    return
}

var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences2(s string) (ans []string) {
    n := len(s)
    if n <= L {
        return
    }
    x := 0
    for _, ch := range s[:L-1] {
        x = x<<2 | bin[byte(ch)]
    }
    cnt := map[int]int{}
    for i := 0; i <= n-L; i++ {
        x = (x<<2 | bin[s[i+L-1]]) & (1<<(L*2) - 1)
        cnt[x]++
        if cnt[x] == 2 {
            ans = append(ans, s[i:i+L])
        }
    }
    return ans
}

const baseL = 10
const bitStand = 1<<(baseL*2) - 1


// 空间换速度
var base = ['T' + 1]int32{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

func findRepeatedDnaSequences(s string) []string {
	var strsRepeat []string
	sL := len(s)
	if sL <= baseL {
		return nil
	}
	var num int32
	for i := range s[:baseL-1] {
		num = num<<2 | base[s[i]]
	}
    // 空间换速度
	numMpCnt := [1 << 20]int8{}
	for idx := 0; idx <= sL-baseL; idx++ {
		bitLow := s[idx+baseL-1]
		num = (num<<2 | base[bitLow]) & bitStand
		if numMpCnt[num] < 2 {
			numMpCnt[num]++
			if numMpCnt[num] == 2 {
				strsRepeat = append(strsRepeat, s[idx:idx+baseL])
			}
		}

	}
	return strsRepeat
}
