package main

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}
func isLeafNode(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if isLeafNode(root){
		return 1
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth{
		return 1 + leftDepth
	}else{
		return 1 + rightDepth
	}
}
