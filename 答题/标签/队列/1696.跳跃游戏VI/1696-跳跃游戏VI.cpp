class Solution {
public:
    int maxResult(vector<int>& nums, int k) {
        int n = nums.size();
        vector<int> dp(n);
        dp[0] = nums[0];
        deque<int> queue;
        queue.push_back(0);
        for (int i = 1; i < n; i++) {
            // 从左往右，把访问不到的弹出
            while (!queue.empty() && queue.front() < i - k) {
                queue.pop_front();
            }
            // 最左边的是最大的
            dp[i] = dp[queue.front()] + nums[i];
            // 从右往左，把比dp[i]小的 弹出
            while (!queue.empty() && dp[queue.back()] <= dp[i]) {
                queue.pop_back();
            }
            queue.push_back(i);
        }
        return dp[n - 1];
    }
};