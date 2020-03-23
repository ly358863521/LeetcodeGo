### 归并排序
- 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
- 设定两个指针，最初位置分别为两个已经排序序列的起始位置
- 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
- 重复步骤3直到某一指针超出序列尾
- 将另一序列剩下的所有元素直接复制到合并序列尾
```golang
func mergeSort(nums []int) []int {
	tmp := make([]int, len(nums), len(nums))
	var sort func(L, R int)
	var merge func(L, M, R int)
	sort = func(L, R int) {
		if L < R {
			M := (L + R) / 2
			sort(L, M)
			sort(M+1, R)
			merge(L, M, R)
		}
	}
	merge = func(L, M, R int) {
		i, j, t := L, M+1, L
		for i <= M && j <= R {
			if nums[i] <= nums[j] {
				tmp[t] = nums[i]
				i++
			} else {
				tmp[t] = nums[j]
				j++
			}
			t++
		}
		for i <= M {
			tmp[t] = nums[i]
			i++
			t++
		}
		for j <= R {
			tmp[t] = nums[j]
			j++
			t++
		}
		for k := L; k <= R; k++ {
			nums[k] = tmp[k]
		}
	}
	sort(0, len(nums)-1)
	return nums
}
```
### 优化
- 当递归到规模足够小时，利用插入排序
- 归并前判断一下是否还有必要归并
- 只在排序前开辟一次空间