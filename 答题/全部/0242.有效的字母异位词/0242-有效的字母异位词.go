// 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
func isAnagram1(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    hashTable := map[byte] int{}
    for i:=0; i<len(s); i++{
        hashTable[s[i]]++
        hashTable[t[i]]--
        if hashTable[t[i]]==0{
            delete(hashTable, t[i])
        }
        if hashTable[s[i]]==0{
            delete(hashTable, s[i])
        }
    }
    return len(hashTable) == 0
}

func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	sm := [26]int{}

	for _, v := range s {
		sm[v-'a']++
	}
	for _, v := range t {
		if sm[v-'a']--;sm[v-'a']<0{
			return false
		}
	}
	return true
	//如果输入字符串包含 unicode 字符
	/*if len(s)!=len(t){
		return false
	}
	sm := make(map[rune]int)

	for _, v := range s {
		sm[v]++
	}
	for _, v := range t {
		if sm[v]--;sm[v]<0{
			return false
		}
	}
	return true*/
}