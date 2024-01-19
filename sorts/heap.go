package sorts

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Val   string
}

func newNode(val string) *Node {
	return &Node{Val: val}
}

func HeapSort(arr []string) []string {
	n := newNode(arr[0])
	buildBTree(arr, 1, n)

	traverseTree(n, "center")

	maxTree(n)

	fmt.Println("-------After Max")
	traverseTree(n, "center")
	return arr
}

func clipTree(n *Node) {

}

func getLastElement(n *Node) *Node {
	if n != nil {
		return nil
	}

	var res *Node

	queue := []*Node{}
	queue = append(queue, n)
	for len(queue) > 0 {
		node := queue[0]
		res = node
		if len(queue) == 1 {
			queue = []*Node{}
		} else {
			queue = queue[1:]
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
}

func buildBTree(arr []string, index int, n *Node) {
	if index > len(arr)-1 || n == nil {
		return
	}

	if index < len(arr)-1 {
		n.Left = newNode(arr[index])
		index++
	}

	if index < len(arr)-1 {
		n.Right = newNode(arr[index])
		index++
	}

	buildBTree(arr, index, n.Left)
}

func maxTree(n *Node) {
	if n == nil {
		return
	}

	maxTree(n.Left)
	maxTree(n.Right)

	if n.Left != nil {
		compareNodes(n, n.Left)
	}

	if n.Right != nil {
		compareNodes(n, n.Right)
	}
}

func compareNodes(n1, n2 *Node) {
	if n1.Val > n2.Val {
		return
	} else {
		n1.Val, n2.Val = n2.Val, n1.Val
	}
}

func traverseTree(n *Node, side string) {
	if n == nil {
		return
	}
	fmt.Println(n.Val, side)
	traverseTree(n.Left, "left")
	traverseTree(n.Right, "right")
}
