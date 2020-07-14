package main
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0{
		return nil
	}
	if len(nums) == 1{
		return &TreeNode{Val:nums[0]}
	}
	root := new(TreeNode)
	root.Val = nums[len(nums)/2]
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	if len(nums)/2+1 == len(nums){
		root.Right = nil
	}else{
		root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	}
	return root
}


