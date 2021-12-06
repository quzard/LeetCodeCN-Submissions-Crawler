class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        if not amount:
            return 0
        dp = [0] *(amount + 1)
        for c in coins:
            if c < amount + 1:
                dp[c] = 1
        for i in range(amount+1):
            now = [dp[i]] if dp[i] else []
            for c in coins:
                if i - c >= 0 and dp[i-c]:
                    now.append(dp[i-c] + 1)
            if now:
                dp[i] = min(now)

        return dp[-1] if dp[-1] != 0 else -1