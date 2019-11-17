- 问题:

> 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
>
> candidates 中的数字可以无限制重复被选取
>
> eg:	输入: candidates = [2,3,6,7], target = 7

- 回溯+剪枝

- ![](https://raw.githubusercontent.com/ly358863521/go-learning/master/img/%E6%9C%AA%E5%91%BD%E5%90%8D%E6%96%87%E4%BB%B6.png)

- 回溯函数：

- 传入当前索引i，当前数组tmp，当前target

- ```python
  def helper(i,tmp,target):
  ```

- 剪枝：判断索引是否越界，判断target是否<0

- 广度搜索

- ```python
  helper(i,tmp+[candidates[i]],target-candidates[i])
  ```

- 深度搜索

- ```python
   helper(i+1,tmp,target)
  ```

```python
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        if not candidates:
            return []
        res = []
        candidates.sort()
        def helper(i,tmp,target):
            if target==0:
                res.append(tmp)
                return
            if i==len(candidates) or target<candidates[i]:
                return
            helper(i,tmp+[candidates[i]],target-candidates[i])
            helper(i+1,tmp,target)
            return
        helper(0,[],target)
        return res
```

