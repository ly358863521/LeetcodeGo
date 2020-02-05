package inorderTraversal

 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	type colornode struct {
		Color int
		Node  *TreeNode
	}

	white, gray := 0, 1
	var res []int
	var stack []*colornode
	stack = append(stack, &colornode{Color: white, Node: root})
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = append(stack[:len(stack)-1])
		if top.Node == nil {
			continue
		}
		if top.Color == white {
			stack = append(stack, &colornode{Color: white, Node: top.Node.Right})
			stack = append(stack, &colornode{Color: gray, Node: top.Node})
			stack = append(stack, &colornode{Color: white, Node: top.Node.Left})
		} else {
			res = append(res, top.Node.Val)
		}
	}

	return res
}