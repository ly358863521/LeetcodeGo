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
  type Freq struct {
  	key   int
  	value int
  }
  type FreqList []Freq
  
  func topKFrequent(nums []int, k int) []int {
  	dict := make(map[int]int)
  	for _, i := range nums {
  		dict[i]++
  	}
  	Frequency := make(FreqList, len(dict))
  	i := 0
  	for k, v := range dict {
  		Frequency[i].key = k
  		Frequency[i].value = v
		i++
  	}
  	heap := Frequency[:k]
  	buildheap(heap, k)
  	for i := k; i < len(Frequency); i++ {
  		if Frequency[i].value > heap[0].value {
  			heap = heap[1:]
  			heap = append(heap, Frequency[i])
  			buildheap(heap, k)
  		}
  	}
  	res := make([]int, k)
  	for i := 0; i < k; i++ {
  		res[i] = heap[i].key
  	}
  	return res
  
  }
  
  func heapify(freq FreqList, n int, i int) {
  	left, right := i*2+1, i*2+2 //左右节点
  	min := i
  	if left < n && freq[min].value > freq[left].value {
  		min = left
  	}
  	if right < n && freq[min].value > freq[right].value {
  		min = right
  	}
  	if min != i {
  		freq[min], freq[i] = freq[i], freq[min]
  		heapify(freq, n, min)
  	}
  }
  func buildheap(freq FreqList, n int) {
  	for i := n / 2; i >= 0; i-- {
  		heapify(freq, n, i)
  	}
  
  }
  ```
  
  ```go
  import "sort"
  type Freq struct {
  	key   int
  	value int
  }
  type FreqList []Freq
  func (p FreqList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
  func (p FreqList) Len() int           { return len(p) }
  func (p FreqList) Less(i, j int) bool { return p[i].value > p[j].value }
  func topKFrequent(nums []int, k int) []int {
      dict :=make(map[int]int)
      for _,i:= range nums{
          dict[i]++
      }
      p:=make(FreqList,len(dict))
      i:=0
      for k, v := range dict {
          p[i].key = k
          p[i].value = v
          i++
      }
      sort.Sort(p)
      res :=make([]int,k)
      for i:=0;i<k;i++{
          res[i] = p[i].key
      }
      return res
  }
  
  ```
  
  
  
- 作弊实现如下：

- ```python
  class Solution:
      def topKFrequent(self, nums: List[int], k: int) -> List[int]:
          from collections import Counter
          return list(zip(*Counter(nums).most_common(k)))[0]
  ```

- ```python
  class Solution:
      def topKFrequent(self, nums: List[int], k: int) -> List[int]:
          from collections import Counter
          v = sorted(Counter(nums).items(),key = lambda x:x[1],reverse = True)
          return [i[0] for i in v][:k]
  ```

- ```python
  class Solution:
      def topKFrequent(self, nums: List[int], k: int) -> List[int]:
          a = {}
          for i in nums:
              a[i] = a.get(i,0)+1
          return [i[0] for i in sorted(a.items(),key = lambda x:x[1],reverse = True)][:k]
  ```