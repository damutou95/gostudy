package main
import "fmt"
func main(){
	a := []int{1,2,4}
	b := []int{1,3,4}
	l1 := makeListNode(a)
	l2 := makeListNode(b)
	head := mergeTwoLists(l1,l2)
	for head != nil{
		fmt.Print(head.Val)
		head = head.Next
	}
	//for l1 != nil{
	//	fmt.Print(l1.Val)
	//	l1 = l1.Next
	//}
	//for l2 != nil{
	//	fmt.Print(l2.Val)
	//	l2 = l2.Next
	//}



}
func makeListNode(nums []int) *ListNode {

	if len(nums) == 0{
		return nil
	}

	res := &ListNode{
		Val:nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val:nums[i],}
		temp = temp.Next
	}

	return  res
}

//Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	current := head
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l1.Val < l2.Val{
		current = l1
		head = l1
		l1 = l1.Next
	}else{
		current = l2
		head = l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			current.Next = l1
			l1 = l1.Next
			current = current.Next
		}else{
			current.Next = l2
			l2 = l2.Next
			current = current.Next
		}
	}
	if l1 != nil{
		current.Next = l1
	}
	if l2 != nil{
		current.Next = l2
	}
	return head
}