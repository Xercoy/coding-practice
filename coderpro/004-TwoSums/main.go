package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func Add(list1, list2 *Node) (*Node, error) {
	node, err := AddHelper(list1, list2, 0)
	if err != nil {
		return node, err
	}

	return node, err
}

func AddHelper(list1, list2 *Node, carryover int) (*Node, error) {
	if list1 == nil && list2 == nil && carryover != 0 {
		return NewNode(carryover), nil
	}

	if list1 == nil {
		list2.Value += carryover
		return list2, nil
	}

	if list2 == nil {
		list1.Value += carryover
		return list1, nil
	}

	// sum of two nodes + carryover
	sum := list1.Value + list2.Value + carryover

	// only take digit in the ones place
	nodeValue := sum % 10

	// transfer over what's in the tens place
	carryover = sum / 10

	fmt.Printf("\nList1=%v, List2=%v", list1.Value, list2.Value)
	fmt.Printf("\nsum: %v", sum)
	fmt.Printf("\nnodeValue: %v", nodeValue)
	fmt.Printf("\ncarryover: %v\n", carryover)

	n := NewNode(nodeValue)

	nextNode, err := AddHelper(list1.Next, list2.Next, carryover)
	if err != nil {
		return n, err
	}

	n.Next = nextNode
	return n, err
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func (n *Node) Print() {
	fmt.Printf("\n%v", n.Value)

	if n.Next == nil {
		return
	}

	n.Next.Print()
}

func main() {
	L1 := NewNode(5)
	L1.Next = NewNode(7)
	L1.Next.Next = NewNode(4)

	L2 := NewNode(8)
	L2.Next = NewNode(6)
	L2.Next.Next = NewNode(9)

	L1.Print()
	L2.Print()

	sum, err := Add(L1, L2)
	if err != nil {
		panic(err)
	}

	sum.Print()
}
