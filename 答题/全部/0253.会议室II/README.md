# [253. 会议室 II](https://leetcode-cn.com/problems/meeting-rooms-ii)
<p>给你一个会议时间安排的数组 <code>intervals</code> ，每个会议时间都会包括开始和结束的时间 <code>intervals[i] = [start<sub>i</sub>, end<sub>i</sub>]</code> ，为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>intervals = [[0,30],[5,10],[15,20]]
<strong>输出：</strong>2
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>intervals = [[7,10],[2,4]]
<strong>输出：</strong>1
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= intervals.length <= 10<sup>4</sup></code></li>
	<li><code>0 <= start<sub>i</sub> < end<sub>i</sub> <= 10<sup>6</sup></code></li>
</ul>

**标签:**  [贪心](https://leetcode-cn.com/tag/greedy) [数组](https://leetcode-cn.com/tag/array) [双指针](https://leetcode-cn.com/tag/two-pointers) [排序](https://leetcode-cn.com/tag/sorting) [堆（优先队列）](https://leetcode-cn.com/tag/heap-priority-queue) 