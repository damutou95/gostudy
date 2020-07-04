package main
import "fmt"
func main(){
	fmt.Print(deleteDuplicates())
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	cur := head
	temp := cur.Val

	for cur.Next != nil{
		if temp != cur.Next.Val{
			cur = cur.Next
			temp = cur.Val
		}else{
			cur.Next = cur.Next.Next
		}
	}
	return head
}
