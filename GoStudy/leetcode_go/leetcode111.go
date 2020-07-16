package main
type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func min(a int, b int) int {
	if a > b {
		return b
	}else{
		return a
	}
}
func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
    if root.Left != nil && root.Right != nil{
   		return 1 + min(minDepth(root.Left), minDepth(root.Right))
    }else if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}else if root.Left != nil && root.Right == nil{
		return 1 + minDepth(root.Left)
	}else {
    	return 1
    }
}
