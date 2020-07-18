package main




type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
//递归解法
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	}else if root.Left != nil && root.Right != nil{
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}else if root.Left != nil{
		return hasPathSum(root.Left, sum-root.Val)
	}
	return hasPathSum(root.Right, sum-root.Val)

}

//迭代解法
//func hasPathSum(root *TreeNode, sum int) bool {
//	if root == nil{
//		return false
//	}
//	nodeSet := []*TreeNode{root}
//	valueSet := []int{root.Val}
//	for len(nodeSet)>0{
//		currentNode := nodeSet[0]
//		nodeSet = nodeSet[1:]
//		currentValue := valueSet[0]
//		valueSet = valueSet[1:]
//		if currentNode.Left == nil && currentNode.Right == nil{
//			if currentValue == sum{
//				return true
//			}else{
//				continue
//			}
//		}
//		if currentNode.Left != nil{
//			nodeSet = append(nodeSet, currentNode.Left)
//			valueSet = append(valueSet, currentValue+currentNode.Left.Val)
//		}
//		if currentNode.Right != nil{
//			nodeSet = append(nodeSet, currentNode.Right)
//			valueSet = append(valueSet, currentValue+currentNode.Right.Val)
//		}
//	}
//	return false
//}
