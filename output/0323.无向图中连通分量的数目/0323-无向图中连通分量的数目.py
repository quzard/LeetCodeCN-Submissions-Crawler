class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:

        edge = defaultdict(list)
        for u,v in edges:
            edge[u].append(v)
            edge[v].append(u)
        
        visit = [False] * n
        res = 0
        def bfs(n):
            q = deque(edge[n])
            while q:
                t = q.popleft()
                if not visit[t]:
                    visit[t] = True
                    q += edge[t]
                
        for i in range(n):
            if not visit[i]:
                visit[i] = True
                res += 1
                bfs(i)
        return res