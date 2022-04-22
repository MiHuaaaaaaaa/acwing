package main

import (
	"algor/linkedlist"
	"fmt"
)

func main() {
	list := linkedlist.NewLinkedList()
	fmt.Println(list.InsertToHead(1))
	fmt.Println(list.InsertToTail(3))
	fmt.Println(list.InsertBefore(list.FindByIndex(1), 2))
	fmt.Println(list.InsertToHead(0))
	fmt.Println(list.InsertToTail(9999))
	list.Print()
}
