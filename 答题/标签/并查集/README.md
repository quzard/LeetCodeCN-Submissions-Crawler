# Union-Find并查集
## 算法解释

在计算机科学中，并查集是一种树型的数据结构，用于处理一些不交集（Disjoint Sets）的合并及查询问题。有一个联合-查找算法（Union-find Algorithm）定义了两个用于此数据结构的操作：

主要用于解决一些**元素分组**的问题。它管理一系列**不相交的集合**，并支持两种操作：

- **合并**（Union）：把两个不相交的集合合并为一个集合。
- **查询**（Find）：确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。

并查集的重要思想在于，**用集合中的一个元素代表集合**。我曾看过一个有趣的比喻，把集合比喻成**帮派**，而代表元素则是**帮主**。接下来我们利用这个比喻，看看并查集是如何运作的。

![img](https://pic4.zhimg.com/80/v2-09fa3fa35e5411444b327d9cb9a31057_1440w.jpg)

最开始，所有大侠各自为战。他们各自的帮主自然就是自己。*（对于只有一个元素的集合，代表元素自然是唯一的那个元素）*

现在1号和3号比武，假设1号赢了（这里具体谁赢暂时不重要），那么3号就认1号作帮主*（合并1号和3号所在的集合，1号为代表元素）*。

![img](https://pic4.zhimg.com/80/v2-3bf6c1a6ecf87fa93f4dbab2012446c7_1440w.jpg)

现在2号想和3号比武*（合并3号和2号所在的集合）*，但3号表示，别跟我打，让我帮主来收拾你*（合并代表元素）*。不妨设这次又是1号赢了，那么2号也认1号做帮主。

![img](https://pic4.zhimg.com/80/v2-be12a6c795572d2acd77dcd49de35127_1440w.jpg)

现在我们假设4、5、6号也进行了一番帮派合并，江湖局势变成下面这样：

![img](https://pic1.zhimg.com/80/v2-3c353bc781c7f3553079d541a9cfdc28_1440w.jpg)

现在假设2号想与6号比，跟刚刚说的一样，喊帮主1号和4号出来打一架（帮主真辛苦啊）。1号胜利后，4号认1号为帮主，当然他的手下也都是跟着投降了。

![img](https://pic3.zhimg.com/80/v2-6362d8b13705d5ba17b19cdeee453022_1440w.jpg)





这是一个**树**状的结构，要寻找集合的代表元素，只需要一层一层往上访问**父节点**（图中箭头所指的圆），直达树的**根节点**（图中橙色的圆）即可。根节点的父节点是它自己。我们可以直接把它画成一棵树：

![img](https://pic2.zhimg.com/80/v2-cca3ddf5806a221201ed78caf1d27041_1440w.jpg)

## 并查集可以解决什么问题

- 组团、配对
- 图的连通性问题
- 集合的个数
- 集合中元素的个数

## 算法模板

```go
package main

// 获取和更新 i 的 根节点
// 要寻找集合的代表元素，只需要一层一层往上访问父节点，直达树的根节点
func Find(fa []int, i int) int {
	if fa[i] == i {
		return i
	}
	fa[i] = Find(fa, fa[i])
	return fa[i]
}

// 把两个不相交的集合合并为一个集合
func Union(fa []int, i, j int) {
  // x 和 y 为 i, j 的根结点
	x, y := fa[i], fa[y]
	if x != y {
    // i 指向 j
		// y 指向 x, x 为父节点， 尾巴指向头
		fa[y] = x
  }
}

// edges为一组关系，可有指向性，也可无
// 如果有指向性，则需要处理闭环问题
func UnionFind(n int, edges [][]int) int {
	// 用一个数组[]fa来存储每个元素的父节点
	fa := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
    //初始化，各个节点的根结点为自己
		fa[i] = i
	}
	for _, edge := range edges {
    //将edge[0], edge[1] 合并
		Union(fa, edge[0], edge[1])
	}
	for i := 0; i < n; i++ {
		if Find(fa, i) == i {
			cnt++
		}
	}
	return cnt
}

```

## 案例

### [323. 无向图中连通分量的数目](https://leetcode-cn.com/problems/number-of-connected-components-in-an-undirected-graph)

