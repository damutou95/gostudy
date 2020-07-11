package main
/**
 * Definition for a binary tree node.
*/

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
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if !check(p,q){
		return false
	}

	//deque := []*TreeNode{}
	//deque = append(deque,p)
	//deque = append(deque,q)

	deque := []*TreeNode{p,q}
	for len(deque)>0{
		left := deque[0]
		right := deque[1]
		if !check(left,right){
			return false
		}else {
			if left != nil{
				deque = append(deque, left.Left)
				deque = append(deque, right.Left)
				deque = append(deque, left.Right)
				deque = append(deque, right.Right)
			}

		}
		deque = deque[2:]

	}
	return true

}
