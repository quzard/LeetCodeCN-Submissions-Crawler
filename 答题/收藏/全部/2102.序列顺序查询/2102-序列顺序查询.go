// https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/treap.go
// 笛卡尔树

package main

import "time"

// github.com/EndlessCheng/codeforces-go
type pair struct {
	x int
	y string
}

type node struct {
	lr       [2]*node
	priority uint
	key      pair
	sz       int
}

func (o *node) cmp(b pair) int {
	switch {
	case b.x > o.key.x || b.x == o.key.x && b.y < o.key.y:
		return 0
	case b.x < o.key.x || b.x == o.key.x && b.y > o.key.y:
		return 1
	default:
		return -1
	}
}

func (o *node) size() int {
	if o != nil {
		return o.sz
	}
	return 0
}

func (o *node) maintain() { o.sz = 1 + o.lr[0].size() + o.lr[1].size() }

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	o.maintain()
	x.maintain()
	return x
}

type treap struct {
	rd   uint
	root *node
}

func (t *treap) fastRand() uint {
	t.rd ^= t.rd << 13
	t.rd ^= t.rd >> 17
	t.rd ^= t.rd << 5
	return t.rd
}

func (t *treap) _put(o *node, key pair) *node {
	if o == nil {
		return &node{priority: t.fastRand(), key: key, sz: 1}
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._put(o.lr[d], key)
		if o.lr[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	o.maintain()
	return o
}

func (t *treap) put(key pair) { t.root = t._put(t.root, key) }

func (t *treap) _delete(o *node, key pair) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(key); d >= 0 {
		o.lr[d] = t._delete(o.lr[d], key)
	} else {
		if o.lr[1] == nil {
			return o.lr[0]
		}
		if o.lr[0] == nil {
			return o.lr[1]
		}
		d = 0
		if o.lr[0].priority > o.lr[1].priority {
			d = 1
		}
		o = o.rotate(d)
		o.lr[d] = t._delete(o.lr[d], key)
	}
	o.maintain()
	return o
}

func (t *treap) delete(key pair) {
	t.root = t._delete(t.root, key)
}

func newTreap() *treap {
	return &treap{rd: uint(time.Now().UnixNano())/2 + 1}
}

// 有 k 个元素小于 o.key
func (t *treap) kth(k int) (o *node) {
	//if k < 0 || k >= t.root.size() {
	//	return
	//}
	for o = t.root; o != nil; {
		if ls := o.lr[0].size(); k < ls {
			o = o.lr[0]
		} else {
			k -= ls + 1 // 1 <-> o.dupCnt
			if k < 0 {
				return
			}
			o = o.lr[1]
		}
	}
	return // NOTE: check nil
}

type SORTracker struct {
}

var (
	t *treap
	c int
)

func Constructor() (_ SORTracker) {
	t = newTreap()
	c = 0
	return
}

func (SORTracker) Add(name string, score int) {
	t.put(pair{score, name})
}

func (SORTracker) Get() (ans string) {
	o := t.kth(c)
	ans = o.key.y
	c++
	return
}
