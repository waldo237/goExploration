package main


type  Node [T any] struct {
	value T
	next  *Node[T]
}

type Graph[T any] struct {
	head *Node[T]
}

func (g *Graph[T]) add(value T) {
	node := &Node[T]{value, nil}
	if g.head == nil {
		g.head = node
	} else {
		current := g.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (g *Graph[T]) print() {
	current := g.head
	for current != nil {
		print(current.value, " ")
		current = current.next
	}
	println()
}

func (g *Graph[T]) bfs(start T) {
	queue := []*Node[T]{}
	visited := map[any]bool{}
	current := g.head
	for current != nil {
		visited[current.value] = false
		current = current.next
	}
	queue = append(queue, g.head)
	visited[start] = true
	for len(queue) > 0 {
		current = queue[0]
		print(current.value, " ")
		queue = queue[1:]
		current = current.next
		for current != nil {
			if !visited[current.value] {
				visited[current.value] = true
				queue = append(queue, current)
			}
			current = current.next
		}
	}
	println()
}

func (g *Graph[T]) dfs(start T) {
	stack := []*Node[T]{}
	visited := map[any]bool{}
	current := g.head
	for current != nil {
		visited[current.value] = false
		current = current.next
	}
	stack = append(stack, g.head)
	visited[start] = true
	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		print(current.value, " ")
		current = current.next
		for current != nil {
			if !visited[current.value] {
				visited[current.value] = true
				stack = append(stack, current)
			}
			current = current.next
		}
	}
	println()
}

func main() {
	g := &Graph[string]{}
	employees := []string{"Bill", "Wes", "Steve", "James", "Zach", "Ryan", "Sam", "Joe", "Scott", "Eric"}
	for _, employee := range employees {
		g.add(employee)

	}
	g.print()
	g.bfs("Eric")
	g.dfs("Wes")
}
