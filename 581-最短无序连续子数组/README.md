1. 给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

   你找到的子数组应是最短的，请输出它的长度。

   示例 1:

   > 输入: [2, 6, 4, 8, 10, 9, 15]
   > 输出: 5
   > 解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序

   

- 从左向右遍历数组，记录最大值，则所有无法更新最大值的位置均需要更改，记录需要更改的最右侧，即为右边界。

- 从右向左同理

- ```python
  class Solution:
      def findUnsortedSubarray(self, nums: List[int]) -> int:
          if not nums:
              return 0
          min_value,max_value = 100000,-100000
          right,left = -1,len(nums)
          for i in range(len(nums)):
              if nums[i]>=max_value:
                  max_value = nums[i]
              else:
                  right = i
          for i in range(len(nums)-1,-1,-1):
              if nums[i]<=min_value:
                  min_value = nums[i]
              else:
                  left = i
          return right-left+1 if right>left else 0
  ```



- ​	 获取所有当前数组与排序后数组具有不同数值的索引，最右边的索引 - 最左边的 + 1 就是结果 

- ```python
  class Solution:
      def findUnsortedSubarray(self, nums: List[int]) -> int:
          diff = [i for i, (a, b) in enumerate(zip(nums, sorted(nums))) if a != b]
          return len(diff) and max(diff) - min(diff) + 1
  ```