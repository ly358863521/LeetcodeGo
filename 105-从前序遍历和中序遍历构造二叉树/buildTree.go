/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(inorder) ==0{
        return nil
    }
    root :=new(TreeNode)
    root.Val =preorder[0]
    var mid int
    for i,val := range inorder{
        if val == preorder[0]{
            mid =i
            break
        }
    }
    root.Left = buildTree(preorder[1:mid+1],inorder[:mid])
    root.Right = buildTree(preorder[mid+1:],inorder[mid+1:])
    return root
}