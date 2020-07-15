package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func isBalanced(root *TreeNode) bool {
	balance,_ := helper(root)
	return balance
}
func abs(a int)int{
	if a > 0{
		return a
	}else{
		return 0 - a
	}
}
func max(a int, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}
func helper(root *TreeNode) (bool,int){
	if root == nil{
		return true,0
	}
	leftBalance, leftHeight := helper(root.Left)
	rightBalance, rightHeight := helper(root.Right)
	height := 1 + max(leftHeight,rightHeight)
	if leftBalance && rightBalance && abs(leftHeight-rightHeight)<2{
		return true, height
	}else{
		return false, height
	}
}
