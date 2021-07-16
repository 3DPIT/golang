//Linked List 제거하는것 만들기
//시간 복잡도 O(1)
package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node
	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNode(root)
	root, tail = RemoveNode(root.next, root, tail) //1지우기
	PrintNode(root)
	root, tail = RemoveNode(root, root, tail) //0 지우기 root 값 지우기
	PrintNode(root)
	root, tail = RemoveNode(tail, root, tail) //9 지우기 tail 값 지우기
	PrintNode(root)
}
func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}
	prev := root
	for prev.next != node {
		prev = prev.next
	}
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}
	return root, tail
}

func PrintNode(node *Node) {
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
