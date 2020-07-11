package main
import "fmt"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main(){
	s := []int{1,2,2,3,4,4,3}
	a := new(TreeNode)
	b := new(TreeNode)
	b = a
	a.Val = 1
	setTree(a,s[1:])
	fmt.Print(isSymmetric(b))
}

func setTree(a *TreeNode,s []int){
	left := new(TreeNode)
	right := new(TreeNode)
	for len(s)>2{
		left.Val = s[0]
		right.Val = s[1]
		a.Left = left
		a.Right = right
		s = s[2:]
		setTree(a.Left,s)
		s = s[2:]
		setTree(a.Right,s)
	}
}
/**
 * Definition for a binary tree node.
 */

//func check(p *TreeNode, q *TreeNode) bool{
//	if (p == nil && q == nil) || (p.Val == q.Val){
//		return true
//	}else{
//		return false
//	}
//}
//func check(p *TreeNode, q *TreeNode) bool{
//	if p == nil && q == nil{
//		return true
//	}else if p == nil || q == nil{
//		return false
//	}
//	if p.Val == q.Val{
//		return true
//	}else{
//		return false
//	}
//}
//func isSymmetric(root *TreeNode) bool {
//	if root == nil{
//		return true
//	}
//	queue := []*TreeNode{root.Left,root.Right}
//	for len(queue)>0{
//		leftNode := queue[0]
//		rightNode := queue[len(queue)-1]
//		if !check(leftNode,rightNode){
//			return false
//		}else if leftNode != nil{
//			queue = queue[1:len(queue)-1]
//			queue = append([]*TreeNode{leftNode.Left,leftNode.Right},queue...)
//			queue = append(queue,rightNode.Left)
//			queue = append(queue, rightNode.Right)
//		}else{
//			queue = queue[1:len(queue)-1]
//		}
//
//	}
//	return true
//}
func isSymmetricTwo(left *TreeNode, right *TreeNode)bool{
	if left == nil && right == nil{
		return true
	}else if left == nil || right == nil{
		return false
	}else if left.Val != right.Val{
		return false
	}
	if left.Left == nil && left.Right == nil && right.Left == nil && right.Right == nil{
		return true
	}else{
		return isSymmetricTwo(left.Left, right.Right) && isSymmetricTwo(left.Right, right.Left)
	}

}
func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return isSymmetricTwo(root.Left, root.Right)
}


