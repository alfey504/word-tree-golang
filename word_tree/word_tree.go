package word_tree

type WordTree struct {
	root *node
}

func CreateWordTree() WordTree {
	node := CreateNode()
	return WordTree{
		root: &node,
	}
}

func (wordTree *WordTree) AddWord(word string) {
	node := wordTree.root
	for _, letter := range word {
		if node.subNodes[letter] == nil {
			newNode := CreateNode()
			node.subNodes[letter] = &newNode
		}
		node = node.subNodes[letter]
	}
}

func (wordTree WordTree) DoesExist(word string) bool {
	node := wordTree.root
	for _, letter := range word {
		if node.subNodes[letter] == nil {
			return false
		}
		node = node.subNodes[letter]
	}
	return true
}
