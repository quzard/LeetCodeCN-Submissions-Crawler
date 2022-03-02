type Node struct{
    Pre *Node
    Next *Node
    Index int
}


func minWindow1(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
    hashTable := make([]int, 256)
	for _, val := range t {
		hashTable[val]++
	}

    hashIndex := map[byte][]*Node{}
    head := &Node{}
    tail := head
    cnt := 0
    length := math.MaxInt32
    l:=0
    r:=0
    for i:=0; i<len(s); i++{
        if hashTable[s[i]] > 0{
            if _, ok := hashIndex[s[i]]; ok{
                if len(hashIndex[s[i]]) < hashTable[s[i]]{
                    cnt++
                    n := &Node{Index: i}
                    n.Pre = tail
                    tail.Next = n
                    tail = tail.Next
                    hashIndex[s[i]] = append(hashIndex[s[i]], n)
                }else{
                    node := hashIndex[s[i]][0]
                    node.Pre.Next = node.Next
                    if node != tail{
                        node.Next.Pre = node.Pre
                        n := &Node{Index: i}
                        n.Pre = tail
                        tail.Next = n
                        tail = tail.Next
                        hashIndex[s[i]] = append(hashIndex[s[i]][1:], n)
                    }else{
                        tail = tail.Pre
                        n := &Node{Index: i}
                        n.Pre = tail
                        tail.Next = n
                        tail = tail.Next
                        hashIndex[s[i]] = append(hashIndex[s[i]][1:], n)
                    }
                }
            }else{
                cnt++
                n := &Node{Index: i}
                n.Pre = tail
                tail.Next = n
                tail = tail.Next
                hashIndex[s[i]] = append(hashIndex[s[i]], n)
            }
            if cnt == len(t){
                if tail.Index  + 1 - head.Next.Index < length{
                    length = tail.Index  + 1 - head.Next.Index
                    l = head.Next.Index
                    r = tail.Index  + 1
                }
            }
        }
    }
    
    if cnt < len(t){
        return ""
    }else{
        return s[l : r]
    }
}




func minWindow2(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	hash := make([]int, 256)
	for _, val := range t {
		hash[val]++
	}
    l, count, res, length := 0, len(t), "", math.MaxInt32
    for r:=0; r<len(s);r++{
        hash[s[r]]--
        if hash[s[r]] >= 0{
            count--
        }
        for l<r && hash[s[l]] < 0{
            hash[s[l]]++
            l++
        }
        if count==0 && r-l < length{
            res = s[l:r+1]
            length = r-l
        }
    }

	return res
}

func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
        return ""
    }
    m := make([]int, 256)
    cnt := len(t)
    for _, n := range t {
        m[n]++
    }
    l := 0
    res := ""
    length := len(s) + 1
    for i := 0; i < len(s); i++ {
        m[s[i]]--
        if m[s[i]] >= 0 {
            cnt--
        }
        for l < i && m[s[l]] < 0 {
            m[s[l]]++
            l++
        }
        if cnt == 0 && i - l + 1 < length{
            res = s[l:i+1]
            length = i - l + 1
        }
    }
    return res
}