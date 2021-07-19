package main

import "errors"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 0 <= n <= 9
// no leading 0s
// 1 <= len(n) <= 100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Addend, l2Addend, carryover int
	var headNode, tmpNode, cursor *ListNode

	// keep going for the length of the longest list
	for {
		l1Empty := l1 == nil
		l2Empty := l2 == nil

		// lists are empty; handle carryover
		if l1Empty && l2Empty {
			if carryover != 0 {
				cursor.Next = &ListNode{Val: carryover, Next: nil}
			}
			break
		}

		if l1Empty {
			l1Addend = 0
		} else {
			l1Addend = l1.Val
			l1 = l1.Next
		}

		if l2Empty {
			l2Addend = 0
		} else {
			l2Addend = l2.Val
			l2 = l2.Next
		}

		// base calculations
		sum := l1Addend + l2Addend + carryover
		nodeValue := sum % 10
		carryover = sum / 10

		// create the new node
		tmpNode = &ListNode{Val: nodeValue, Next: nil}

		// if running for the first time, this is the head node; keep a ptr to it
		if headNode == nil {
			headNode = tmpNode
			cursor = tmpNode
		} else {
			cursor.Next = tmpNode
			cursor = tmpNode
		}
	}
	return headNode
}

func main() {
	l1 := &ListNode{
		Val: 0,
	}

	l2 := &ListNode{
		Val: 0,
	}

	result := addTwoNumbers(l1, l2)
	if result.Val != 0 {
		panic(errors.New("Fail"))
	}
}
