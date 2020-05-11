package main

import "fmt"

type Node struct {
	data int
	Next *Node
	Prev *Node
	key interface{}
}

type List struct {
	head *Node
	tail *Node
	length int
}

type ListFeature interface {
	addNode()
	deleteNode()
	printNodes()
	length()
	update()
}

type EmptyList struct{}

func(t *EmptyList) Error() string {
	return "This list is empty"
}
func (L *List) addNode(key interface{}) {
	list := &Node{
		Next: L.head,
		key: key,
	}
	if L.head != nil {
		L.head.Prev = list
	}
	L.head = list
	l := L.head
	for l.Next != nil {
		l = l.Next
	}
	L.tail = l
}
//recursive call for printing nodes
func (l *List) printNodes() {
	list := l.head
	if list != nil {
		fmt.Print(list.key, " ")
		list = list.Next
		printNodes(list)
	}
}

func printNodes(list *Node){
	if list != nil {
		fmt.Print(list.key, " ")
		list = list.Next
		printNodes(list)
	}
}

func (l *List) deleteNode( val int) {
	if l.length == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.Next
		l.length--
		return
	}

	var prevNode *Node
	node := l.head
	for node.data != val {
		if node == nil {
			return
		}
		prevNode = node
		node = node.Next
	}
	prevNode.Next = node.Next
	l.length--
}


func main() {
	link := List{}
	fmt.Println("adding 5 values to linked list, here they are: ")
	link.addNode("6")
	link.addNode("10")
	link.addNode("85")
	link.addNode("82")
	link.addNode("54")
	link.printNodes()
	fmt.Println()

}
