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
    expectedList := createList([]int{7, 0, 8})
    printList(result) // Expected output: 7 0 8
    printList(expectedList) // Expected output: 7 0 8

    if result.Val != expectedList.Val || result.Next.Val != expectedList.Next.Val || result.Next.Next.Val != expectedList.Next.Next.Val {
	t.Errorf("addTwoNumbers failed")
    }
}

// Performance test
func performanceTest() {
    const iterations = 1000
    const listSize = 100
    startTime := time.Now()
    for i := 0; i < iterations; i++ {
	list1 := createList(makeListOfSize(listSize))
	list2 := createList(makeListOfSize(listSize))
	_ = addTwoNumbers(list1, list2)
    }
    elapsedTime := time.Since(startTime)
    fmt.Printf("Performance test completed in %v\n", elapsedTime)
}

// makeListOfSize creates a list of a given size with incremental values
func makeListOfSize(size int) []int {
    list := make([]int, size)
    for i := range list {
	list[i] = i + 1
    }
    return list
}

func main() {
    // Run performance test
    performanceTest()
}

func TestMain(m *testing.M) {
    // Run unit tests
    m.Run()
}
