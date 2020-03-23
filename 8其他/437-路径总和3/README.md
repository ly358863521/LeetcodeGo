### 方案一

1）双重递归

2）以某个节点开始，先序遍历求出以当前节点开始，所有可能的情况总数

3）递归所有节点

```python
class Solution:
    def path(self,root,sum):
        if root==None:
            return 0
        res = 0
        if root.val ==sum:
            res+=1
        res +=self.path(root.left,sum-root.val)
        res +=self.path(root.right,sum-root.val)
        return res
    def pathSum(self, root: TreeNode, sum: int) -> int:
        if root ==None:
            return 0
        return self.path(root,sum)+self.pathSum(root.left,sum)+self.pathSum(root.right,sum)
```

### 方案二

1）遍历到某个节点时，检查该节点到根节点的路径上包含该节点的所有可能

2）用栈数组保存当前路径包含的节点

```python
def pathSum(self, root, sum):
        self.data=[]
        self.count=0
        
        def recursive(root,sum):
            if not root:
                return 
            self.data.append(root.val)
            cur=0
            for i in range(len(self.data)-1,-1,-1):
                cur+=self.data[i]
                if cur==sum:
                    self.count+=1
            recursive(root.left,sum)
            recursive(root.right,sum)
            self.data.pop()

        recursive(root,sum)
        return self.count
```

