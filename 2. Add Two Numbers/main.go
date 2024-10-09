package main

import (
    "fmt"
    "testing"
    "time"
)

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

// addTwoNumbers adds two numbers represented by linked lists
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    current := dummyHead
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
	val1 := 0
	if l1 != nil {
	    val1 = l1.Val
	    l1 = l1.Next
	}
	val2 := 0
	if l2 != nil {
	    val2 = l2.Val
	    l2 = l2.Next
	}
	sum := val1 + val2 + carry
	carry = sum / 10
	current.Next = &ListNode{Val: sum % 10}
	current = current.Next
    }
    return dummyHead.Next
}

// Helper function to create a linked list from a slice
func createList(values []int) *ListNode {
    if len(values) == 0 {
	return nil
    }
    head := &ListNode{Val: values[0]}
    current := head
    for _, value := range values[1:] {
	current.Next = &ListNode{Val: value}
	current = current.Next
    }
    return head
}

// Helper function to print a linked list
func printList(head *ListNode) {
    for head != nil {
	fmt.Print(head.Val, " ")
	head = head.Next
    }
    fmt.Println()
}

// Unit tests
func TestAddTwoNumbers(t *testing.T) {
    list1 := createList([]int{2, 4, 3})
    list2 := createList([]int{5, 6, 4})
    result := addTwoNumbers(list1, list2)
    printList(result) // Expected output: 7 0 8

    list1 = createList([]int{0})
    list2 = createList([]int{0})
    result = addTwoNumbers(list1, list2)
    printList(result) // Expected output: 0
}

// Performance test
func performanceTest() {
    startTime := time.Now()
    for i := 0; i < 1000; i++ {
	list1 := createList([]int{1})
	list2 := createList([]int{9})
	_ = addTwoNumbers(list1, list2)
    }
    elapsedTime := time.Since(startTime)
    fmt.Printf("Performance test completed in %v\n", elapsedTime)
}

func main() {
    // Run unit tests
    TestAddTwoNumbers(nil)
    // Run performance test
    performanceTest()
}
