# Floyd判圈算法(龟兔赛跑算法)

1. 问题：如何检测一个链表是否有环，如果有，如何确定环的起点和环的长度

   - 时间复杂度 O(n)
   - 空间复杂度 O(1)

2. 基本思想：

   - 慢指针一次前进一步，快指针一次前进两步（或多步）
   - 如果相遇，说明有环
   - 将慢指针（或快指针）移到链表起点，然后同时移动两指针，每次一步
   - 则再次相遇位置即环的起点
   - 证明：
     - 设环长为n，非环长为m，相遇时刻慢指针走了m+A*n+k（A为圈数），快指针走了m+B*n+k
       且快指针为慢指针2倍，2S和S，则S = (B-A)*n
       说明两指针都走了圈长的倍数
     - 那么之后每次移动一步，当慢指针移动了m步时，快指针移动了2S+m步，即快指针从起点到环入口处，之后走了整数圈，此时也在环入口，说明相遇位置为环的起点。

3. 此题中因为数组中的数字再索引范围内，则如果以数组值作为下一步的索引值，则必然出现环结构，且重复的数字就是环入口

   - ```python
     class Solution:
         def findDuplicate(self, nums: List[int]) -> int:
             slowPtr = nums[0]
             fastPtr = nums[0]
             while True:
                 slowPtr = nums[slowPtr]
                 fastPtr = nums[nums[fastPtr]]
                 if slowPtr == fastPtr:
                     break
             slowPtr = nums[0]
             while slowPtr != fastPtr:
                 slowPtr = nums[slowPtr]
                 fastPtr = nums[fastPtr]
             return fastPtr
     ```