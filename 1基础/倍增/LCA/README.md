# LCA最近公共祖先问题
- 首先可以选择用递归的方法
  - 当root = p或q时，返回root
  - 在左子树中找pq的最近公共祖先 = left，如果非空，返回left
  - 在右子树中找pq的最近公共祖先 = right，如果非空，返回right
  - 如果left和right都非空，则返回root
##　Leetcode 236
```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || p == root || q == root {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    } else if left != nil {
        return left
    } else {
        return right
    }
}
```