class Solution:
    def minimumSemesters(self, n: int, relations: List[List[int]]) -> int:
        edge = defaultdict(list)
        indgrees = [0] * n
        res = total = 0
        for f,t in relations:
            edge[f-1].append(t-1)
            indgrees[t-1] += 1
        q = deque()
        for i in range(n):
            if not indgrees[i]:
                q.append(i)
        while q:
            res += 1
            for _ in range(len(q)):
                now = q.popleft()
                total += 1
                for t in edge[now]:
                    indgrees[t] -= 1
                    if indgrees[t] == 0:
                        q.append(t)
        if total == n:

            return res 
        else:
            return -1
