package main

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
	*     Val int
	*     Next *ListNode
	* }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}

		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}

	return head.Next

}

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 7}
	l1.Next.Next = &ListNode{Val: 9}

	l2 := &ListNode{Val: 8}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 3}

	result := addTwoNumbers(l1, l2)

	fmt.Println(result.Val)
	p := result.Next
	for p != nil {
		fmt.Print(p.Val)
		if p.Next == nil {
			break
		} else {
			p = p.Next
		}
	}
}
