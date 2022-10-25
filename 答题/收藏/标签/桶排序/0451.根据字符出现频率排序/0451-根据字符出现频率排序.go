func frequencySort(s string) string {
    type pair struct{
        b byte
        c int
    }
    m:=make(map[byte]int)
    for i:=0;i<len(s);i++{
        m[s[i]]++
    }
    pairs:=make([]pair,0,len(m))
    for k,v:=range m{
        pairs=append(pairs,pair{k,v})
    }
    sort.Slice(pairs,func(i,j int)bool{
        return pairs[i].c>pairs[j].c
    })
    res:=make([]byte,0,len(s))
    for _,v:=range pairs{
        res=append(res,bytes.Repeat([]byte{v.b},v.c)...)
    }
    return string(res)
}