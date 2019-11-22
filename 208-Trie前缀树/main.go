type Trie struct {
	children map[int32]*Trie
	isword   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{children: make(map[int32]*Trie), isword: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, w := range word {
		if cur.children[w] == nil {
			cur.children[w] = &Trie{children: make(map[int32]*Trie), isword: false}
			cur = cur.children[w]
		} else {
			cur = cur.children[w]
		}
	}
	cur.isword = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, w := range word {
		if cur.children[w] == nil {
			return false
		}
		cur = cur.children[w]
	}
	if cur.isword == true {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, w := range prefix {
		if cur.children[w] == nil {
			return false
		}
		cur = cur.children[w]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */