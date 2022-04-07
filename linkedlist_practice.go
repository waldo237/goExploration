package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (l *LinkedList[T]) AddNode(value T) {
	node := &Node[T]{value, nil}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.length++
}

func (l *LinkedList[T]) AddNodeToBeginning(value T) {
	node := &Node[T]{value, nil}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.length++
}

func (l *LinkedList[T]) AddNodeAfter(value T, node *Node[T]) {
	newNode := &Node[T]{value, nil}
	if node == nil {
		l.AddNode(value)
	} else {
		newNode.next = node.next
		node.next = newNode
		l.length++
	}
}

func (l *LinkedList[T]) AddNodeBefore(value T, node *Node[T]) {
	newNode := &Node[T]{value, nil}
	if node == nil {
		l.AddNode(value)
	} else {
		if node == l.head {
			l.AddNodeToBeginning(value)
		} else {
			newNode.next = node
			previousNode := l.head
			for previousNode.next != node {
				previousNode = previousNode.next
			}
			previousNode.next = newNode
			l.length++
		}
	}
}

func (l *LinkedList[T]) RemoveNode(node *Node[T]) {
	if node == nil {
		return
	} else {
		if node == l.head {
			l.head = l.head.next
		} else {
			previousNode := l.head
			for previousNode.next != node {
				previousNode = previousNode.next
			}
			previousNode.next = node.next
		}
		l.length--
	}
}

func (l *LinkedList[T]) RemoveNodeFromBeginning() {
	if l.head == nil {
		return
	} else {
		l.head = l.head.next
		l.length--
	}
}

func (l *LinkedList[T]) RemoveNodeAfter(node *Node[T]) {
	if node == nil {
		return
	} else {
		if node.next == nil {
			return
		} else {
			node.next = node.next.next
			l.length--
		}
	}
}

func (l *LinkedList[T]) RemoveNodeBefore(node *Node[T]) {
	if node == nil {
		return
	} else {
		if node == l.head {
			l.RemoveNodeFromBeginning()
		} else {
			previousNode := l.head
			for previousNode.next != node {
				previousNode = previousNode.next
			}
			previousNode.next = node
			l.length--
		}
	}
}

func (l *LinkedList[T]) PrintList() {
	node := l.head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (l *LinkedList[T]) ReverseList() {
	if l.head == nil {
		return
	} else {
		node := l.head
		for node != nil {
			l.AddNodeToBeginning(node.value)
			l.RemoveNode(node)
			node = node.next
		}
	}
}

func (l *LinkedList[T]) ReverseListRecursive() {
	if l.head == nil {
		return
	} else {
		l.ReverseListRecursiveHelper(l.head)
	}
}

func (l *LinkedList[T]) ReverseListRecursiveHelper(node *Node[T]) {
	if node == nil {
		return
	} else {
		l.ReverseListRecursiveHelper(node.next)
		l.AddNodeToBeginning(node.value)
		l.RemoveNode(node)
	}
}

func (l *LinkedList[T]) GetNthNode(n int) *Node[T] {
	if n < 0 || n >= l.length {
		return nil
	} else {
		node := l.head
		for i := 0; i < n; i++ {
			node = node.next
		}
		return node
	}
}

func main() {
	list := &LinkedList[string]{}
	animals := []string{"cat", "dog", "lion", "tiger", "bear", "monkey", "elephant", "giraffe", "penguin", "panda", "zebra", "rhino", "hippo", "hyena", "kangaroo", "cheetah", "crocodile", "alligator", "camel", "ant", "bee", "butterfly", "caterpillar", "dragonfly", "eagle", "falcon", "flamingo", "giraffe", "hawk", "hummingbird", "jellyfish", "kite", "ladybug", "lion", "lobster", "mantis", "moth", "octopus", "owl", "parrot", "penguin", "pig", "rabbit", "rat", "scorpion", "seal", "shark", "snail", "snake", "spider", "squirrel", "turtle", "whale", "zebra"}

	fmt.Println("Original slice:", len(animals))

	for _, animal := range animals {
		list.AddNode(animal)
	}
	fmt.Println("Before reversing:")
	list.PrintList()
	
	fmt.Print("\nAfter reversing:\n")
	list.ReverseList()
	list.PrintList()

	fmt.Println("Before adding loop: ", list.length)
	for index, animal := range animals {
		list.AddNodeAfter(animal, list.GetNthNode(index))
	}
	fmt.Println("After adding loop: ", list.length)

	fmt.Println("Before the deletion: ", list.length)
	for index, _ := range animals {
		list.RemoveNode(list.GetNthNode(index))
	}
	fmt.Println("After removing nodes: ", list.length)
}
