package trie

type node struct {
	isWordEnd bool
	next      map[rune]*node
}

type trie struct {
	root *node
	size int
}

// 构造函数
func Trie() *trie {
	var (
		root *node
	)
	root = &node{isWordEnd: false, next: make(map[rune]*node)}
	return &trie{root: root, size: 0}
}

// 向trie添加word
func (t *trie) Add(word string) {
	var (
		current *node
		exists  bool
		char    rune
	)
	current = t.root
	for _, char = range word {
		// char已经存在
		if _, exists = current.next[char]; !exists {
			// 创建新节点
			current.next[char] = &node{isWordEnd: false, next: make(map[rune]*node)}
		}
		current, _ = current.next[char]
	}
	// 如果node不是单词的结尾
	if !current.isWordEnd {
		current.isWordEnd = true
		t.size++
	}
}

// trie中是否包含word
func (c *trie) Contains(word string) bool {
	var (
		current *node
		char    rune
		exists  bool
	)
	current = c.root
	for _, char = range word {
		if _, exists = current.next[char]; !exists {
			return false
		}
		current, _ = current.next[char]
	}
	// 如果结尾不是node是单词结尾
	if !current.isWordEnd {
		return false
	}
	return true
}

// 获取trie中的单词数量
func (t *trie) GetSize() int {
	return t.size
}

//
