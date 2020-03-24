package main

import (
	"fmt"
	"skiplist/skiplist"
)

func main() {
	list := skiplist.Constructor()
	list.Add(1)
	list.Add(1)
	list.Add(1)
	list.Print()
	fmt.Println(list.Search(0))
	fmt.Println(list.Search(1))
	list.Erase(1)
	list.Print()
	fmt.Println(list.Search(1))
	list.Erase(1)
	list.Print()
	fmt.Println(list.Search(1))
	list.Erase(1)
	fmt.Println(list.Search(1))
	list.Print()
}
