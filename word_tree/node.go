package word_tree

type node struct {
	subNodes map[rune]*node
}

func CreateNode() node {
	return node{
		subNodes: make(map[rune]*node),
	}
}
