### 二叉树中的最大路径

- 非空二叉树，返回最大路径和

- 不一定经过根节点

- 则对于任意节点A，经过A的最大路径为A左侧的最大路径+A右侧的最大路径+当前节点值

- 而A左侧的最大路径为

  - A左子树节点的max（左侧最大路径，右侧最大路径）+当前节点值

  - 如果为负，则置0

  - ```go
    import "math"
    func max(A,B int )int{
        if A>B{return A}
        return B
    }
    func maxPathSum(root *TreeNode) int {
        res :=int(-math.NaN())
        var helper func (root *TreeNode)int
        helper = func (root *TreeNode)int{
            if root==nil{return 0}
            l := helper(root.Left)
            r := helper(root.Right)
            res = max(res,l+r+root.Val)
            return max(0,max(l,r)+root.Val)
        }
        helper(root)
        return res
    
    }
    ```