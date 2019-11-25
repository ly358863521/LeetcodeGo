## 堆排序

- 堆排序处理海量数据的topK，不需要排序所有元素，只需要比较和根节点的大小关系

- 原则：最大堆求前n小，最小堆求前n大

- 本题求前K个高频元素，首先建立字典遍历统计元素频率，取前K个数，建立K最小堆

- 遍历K之外元素，大于堆顶则入堆，维护规模为K的最小堆minheap

- python实现如下：

- ```python
  def heapify(arr, n, i):
              smallest = i  # 构造根节点与左右子节点
              l = 2 * i + 1
              r = 2 * i + 2
              if l < n and arr[l][1] < arr[i][1]:  # 如果左子节点在范围内且小于父节点
                  smallest = l
              if r < n and arr[r][1] < arr[smallest][1]:
                  smallest = r
              if smallest != i:  # 递归基:如果没有交换，退出递归
                  arr[i], arr[smallest] = arr[smallest], arr[i]
                  heapify(arr, n, smallest)  # 确保交换后，小于其左右子节点
  ```

- ```python
  class Solution:
      def topKFrequent(self, nums: List[int], k: int) -> List[int]: 
          # 哈希字典统计出现频率
          map_dict = {}
          for item in nums:
              map_dict[item] = map_dict.get(item,0)+1
              
          map_arr = list(map_dict.items())
          lenth = len(map_dict.keys())
          # 构造规模为k的minheap
          if k <= lenth:
              k_minheap = map_arr[:k]
              # 从后往前建堆，避免局部符合而影响递归跳转，例:2,1,3,4,5,0
              for i in range(k // 2 - 1, -1, -1): 
                  heapify(k_minheap, k, i)
              # 对于k:, 大于堆顶则入堆，维护规模为k的minheap
              for i in range(k, lenth): # 堆建好了，没有乱序，从前往后即可
                  if map_arr[i][1] > k_minheap[0][1]:
                      k_minheap[0] = map_arr[i] # 入堆顶
                      heapify(k_minheap, k, 0)  # 维护 minheap
          # 如需按顺序输出，对规模为k的堆进行排序
          # 从尾部起，依次与顶点交换再构造minheap，最小值被置于尾部
          for i in range(k - 1, 0, -1):
              k_minheap[i], k_minheap[0] = k_minheap[0], k_minheap[i]
              k -= 1 # 交换后，维护的堆规模-1
              heapify(k_minheap, k, 0)
          return [item[0] for item in k_minheap]
  ```

- Go实现如下：

- ```go
  func heapify(nums []int,i int){
  	n := len(nums)
  	left,right := i*2+1,i*2+2 //左右节点
  	min := i
  	if left<n && nums[min] > nums[left]{
  		min = left
  	}
  	if right<n && nums[min] > nums[right]{
  		min = right
  	}
  	if min != i{
  		nums[min],nums[i] = nums[i],nums[min]
  		heapify(nums,min)
  	}
  }
  ```

- 