package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const p = 0.25

// DefaultMaxLevel maxlevel
const DefaultMaxLevel = 16

type node struct {
	value   int
	forward []*node
	// forward[i]指当前节点在第i层的下一个节点
}

// Skiplist skiplist
type Skiplist struct {
	header *node
	length int
	level  int
}

func newNode(value int, level int) *node {
	return &node{value: value, forward: make([]*node, level)}
}

func randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float64() < p && level < DefaultMaxLevel; level++ {
	}
	// fmt.Println(level)
	return
}

// Constructor new
func Constructor() Skiplist {
	return Skiplist{header: newNode(-1, DefaultMaxLevel), length: 0, level: 1}
}

// Search search
func (s *Skiplist) Search(target int) bool {
	cur := s.header
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil {
			if cur.forward[i].value == target {
				return true
			} else if cur.forward[i].value > target {
				break
			} else {
				cur = cur.forward[i]
			}
		}
	}
	return false
}

// Add insert a num
func (s *Skiplist) Add(num int) {
	cur := s.header
	// update 新节点的前驱节点
	update := make([]*node, DefaultMaxLevel)
	for i := s.level - 1; i >= 0; i-- {
		if cur.forward[i] == nil || cur.forward[i].value > num {
			update[i] = cur
		} else {
			for cur.forward[i] != nil && cur.forward[i].value < num {
				cur = cur.forward[i]
			}
			update[i] = cur
		}
	}

	// 是否允许插入重复value的节点
	// if cur.forward[0] != nil && cur.forward[0].value == num {
	// 	return
	// }

	level := randLevel()
	if level > s.level {
		for i := s.level; i < level; i++ {
			update[i] = s.header
		}
		s.level = level
	}
	node := newNode(num, level)
	for i := 0; i < level; i++ {
		node.forward[i] = update[i].forward[i]
		update[i].forward[i] = node
	}
	s.length++
}

// Erase delete a num
func (s *Skiplist) Erase(num int) bool {
	cur := s.header
	flag := false
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil {
			if cur.forward[i].value == num {
				// tmp := cur.forward[i]
				// cur.forward[i] = tmp.forward[i]
				// tmp.forward[i] = nil
				cur.forward[i] = cur.forward[i].forward[i]
				break
			} else if cur.forward[i].value > num {
				break
			} else {
				cur = cur.forward[i]
			}
		}
	}
	if flag {
		s.length--
		return true
	}
	return false
}

// Print the skiplist
func (s *Skiplist) Print() {
	fmt.Println("总数：", s.length, "总层数：", s.level)
	for i := s.level - 1; i >= 0; i-- {
		cur := s.header
		for cur.forward[i] != nil {
			fmt.Printf("%d ", cur.forward[i].value)
			cur = cur.forward[i]
		}
		fmt.Printf("*************** Level %d \n", i+1)
	}
}
