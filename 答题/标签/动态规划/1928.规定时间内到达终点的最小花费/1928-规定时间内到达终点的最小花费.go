const MAX_INT = 1<<31 - 1

type Node struct {
	T int
	X int
	C int
}

type Queue []Node

func (c *Queue) Len() int {
	return len(*c)
}

func (c *Queue) Push(v Node) {
	*c = append(*c, v)
}

func (c *Queue) LPop() Node {
	var res = (*c)[0]
	*c = (*c)[1:]
	return res
}

func (c *Queue) RPop() Node {
	var n = c.Len()
	var res = (*c)[n-1]
	*c = (*c)[:n-1]
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	var n = len(passingFees)
	var costs = make([]int, n)
	var m = make([]map[int]int, n)
	for i, _ := range m {
		m[i] = make(map[int]int, 0)
		costs[i] = -1
	}
	for _, item := range edges {
		if m[item[0]][item[1]] == 0 {
			m[item[0]][item[1]] = item[2]
		} else {
			m[item[0]][item[1]] = min(item[2], m[item[0]][item[1]])
		}
		if m[item[1]][item[0]] == 0 {
			m[item[1]][item[0]] = item[2]
		} else {
			m[item[1]][item[0]] = min(item[2], m[item[1]][item[0]])
		}
	}

	costs[0] = passingFees[0]
	var times = getMinTimes(m, maxTime, n)

	var q = &Queue{}
	q.Push(Node{T: 0, X: 0, C: passingFees[0]})
	for q.Len() > 0 {
		var node = q.RPop()
		if node.T > maxTime || node.X == n-1 {
			continue
		}

		for k, v := range m[node.X] {
			var next = Node{
				T: node.T + v,
				X: k,
				C: node.C + passingFees[k],
			}
			if (costs[k] == -1 || next.C < costs[k]) && next.T+times[k] <= maxTime {
				costs[k] = next.C
				q.Push(next)
			}
		}
	}
	return costs[n-1]
}

func getMinTimes(m []map[int]int, maxTime int, n int) []int {
	var times = make([]int, n)
	for i, _ := range times {
		times[i] = MAX_INT
	}
	var start = Node{T: 0, X: n - 1}
	times[start.X] = start.T
	var q = &Queue{}
	q.Push(start)

	for q.Len() > 0 {
		var node = q.RPop()
		if node.T > maxTime || node.X == 0 {
			continue
		}

		for k, v := range m[node.X] {
			var next = Node{
				T: node.T + v,
				X: k,
			}
			if next.T < times[k] && next.T <= maxTime {
				times[k] = next.T
				q.Push(next)
			}
		}
	}
	return times
}