package main

import (
	"fmt"
)

// 定义结点
type Node struct {
	data int64
	Next *Node
}

// 链表
type LinkedList struct {
	head *Node
}

// 头插法创建单链表
func (list *LinkedList) InsertFirst(list_val []int64) {
	if list.head == nil {
		// 创建头结点
		head := &Node{data: 0}
		pointer := head // 创建指针，一开始指向头结点
		for _, val := range list_val {
			new_node := &Node{data: val} // 创建新结点
			new_node.Next = pointer.Next // 头结点链接在新结点之后
			pointer.Next = new_node
		}
		// 打印链表
		for pointer != nil {
			fmt.Println(pointer.data)
			pointer = pointer.Next
		}
	}
}

// 尾插法创建单链表
func (list *LinkedList) InsertLast(list_val []int64) {
	if list.head == nil {
		// 创建头结点
		head := &Node{data: 0}
		pointer := head // 创建尾指针，一开始指向头结点
		for _, val := range list_val {
			new_node := &Node{data: val} // 创建新结点
			pointer.Next = new_node
			pointer = new_node
		}
		// 打印链表
		for head != nil {
			fmt.Println(head.data)
			head = head.Next
		}
		pointer.Next = nil
	}
}

// 按值删除某个结点
func (list *LinkedList) deleteByVal(val int64) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == val {
		list.head = list.head.Next
		return true
	}
	pointer := list.head
	delete_array := []int64{}
	for pointer.Next != nil {
		if pointer.Next.data == val {
			delete_array = append(delete_array, pointer.Next.data)
			pointer.Next = pointer.Next.Next
			continue
		} else {
			pointer = pointer.Next
		}
	}
	return len(delete_array) != 0
}

// 打印链表
func (list *LinkedList) printList() {
	for list.head != nil {
		fmt.Println(list.head.data)
		list.head = list.head.Next
	}
}

var l LinkedList

func main() {
	// l.InsertLast([]int64{1, 5, 6, 9})
	l.InsertFirst([]int64{1, 5, 6, 9})
	fmt.Println(l.deleteByVal(5))
	l.printList()
}
