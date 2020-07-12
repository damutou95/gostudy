package main
type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	temp := []*TreeNode{root}
	temp2 := []*TreeNode{}
	res := [][]int{[]int{}}
	for len(temp)>0{
		a := temp[0]
		temp = temp[1:]
		if a != nil{
			res[0] = append(res[0], a.Val)
		}
		if a != nil{
			if a.Left != nil{
				temp2 = append(temp2, a.Left)
			}
			if a.Right != nil{
				temp2 = append(temp2, a.Right)
			}
		}
		if len(temp) == 0 && len(temp2)>0{
			res = append([][]int{[]int{}},res...)
			temp = temp2
			temp2 = []*TreeNode{}
		}

	}
	return res
}
