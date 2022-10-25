
func maximumSwap(num int) int {
	s:=[]byte(strconv.Itoa(num))
	pos:=make([]int,10)
	for i:=0;i<len(s);i++{
		pos[s[i]-'0']=i
	}
	for i:=0;i<len(s)-1;i++{
		for j:=9;j>int(s[i]-'0');j--{
			if pos[j]>i{
				tmp:=s[i]
				s[i]=s[pos[j]]
				s[pos[j]]=tmp
				ress, _ :=strconv.Atoi(string(s))
				fmt.Println(ress)
				return  ress
			}
		}
	}
	return num
}
