// ×ÖµäÊý
// type Tree struct{
//     word []*Tree
//     end  []bool
// }

// type Trie struct {
//     t *Tree
// }


// func Constructor() Trie {
//     return Trie{
//         t: &Tree{
//             word: make([]*Tree, 27),
//             end: make([]bool, 27),
//         },
//     }
// }


// func (this *Trie) Insert(word string)  {
//     if word == ""{
//         return  
//     }
//     cur := this.t
    
//     for i := 0; i < len(word); i++{
//         if cur.word[int(word[i] - 'a')] == nil{
//             cur.word[int(word[i] - 'a')] =&Tree{
//                 word: make([]*Tree, 26),
//                 end: make([]bool, 26),
//                 }
//         }
//         if i == len(word) - 1{
//             cur.end[int(word[i] - 'a')] = true
//         }
//         cur = cur.word[int(word[i] - 'a')]
//     }
// }


// func (this *Trie) Search(word string) bool {
//     cur := this.t
//     for i := 0; i < len(word); i++{
//         if cur.word[int(word[i] - 'a')] == nil{
//             return false
//         }
//         if i == len(word) - 1{
//             return cur.end[int(word[i] - 'a')]
//         }
//         cur = cur.word[int(word[i] - 'a')]
//     }
//     return false
// }


// func (this *Trie) StartsWith(prefix string) bool {
//     cur := this.t
//     for i := 0; i < len(prefix); i++{
//         if cur.word[int(prefix[i] - 'a')] == nil{
//             return false
//         }
//         if i == len(prefix) - 1{
//             return true
//         }
//         cur = cur.word[int(prefix[i] - 'a')]
//     }
//     return false
// }








type Trie struct {
	IsWord bool
	next   [26]*Trie
}

func Constructor() Trie {
	return Trie{IsWord: false}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.IsWord = true
		return
	}
	if this.next[word[0]-'a'] == nil {
		this.next[word[0]-'a'] = &Trie{}
	}
	this.next[word[0]-'a'].Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.IsWord
	}
	if this.next[word[0]-'a'] == nil {
		return false
	} else {
		return this.next[word[0]-'a'].Search(word[1:])
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if this.next[prefix[0]-'a'] == nil {
		return false
	} else {
		return this.next[prefix[0]-'a'].StartsWith(prefix[1:])
	}
}