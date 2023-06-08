package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	ll := InitializeLinkedList(12, 5, 8, 9, 11, 13)
	fmt.Println(ll.ToString())
	ll.Append(7)
	fmt.Println(ll.ToString())
	ll.Append(1)
	fmt.Println(ll.ToString())
	ll.Remove(5)
	fmt.Println(ll.ToString())
	ll.Remove(9)
	fmt.Println(ll.ToString())
	ll.Remove(14)
	fmt.Println(ll.ToString())
}

func InitializeLinkedList(params ...int) *LinkedList {
	ll := &LinkedList{head: nil}
	if params != nil {
		for _, v := range params {
			ll.Append(v)
		}
	}
	return ll
}

func (list *LinkedList) Append(value int) {
	node := &Node{data: value, next: nil}
	if list.head == nil {
		list.head = node
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (list *LinkedList) Remove(value int) {
	if list.head == nil {
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

func (list *LinkedList) ToString() string {
	result := ""
	current := list.head
	for current != nil {
		result += fmt.Sprintf("%v->", current.data)
		current = current.next
	}
	result += "NULL"
	return result
}
