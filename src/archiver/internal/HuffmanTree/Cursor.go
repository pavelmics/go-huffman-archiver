package HuffmanTree

// Tree BitCursor
type BitCursor struct {
	currentNode *Node
	history []int
}

func CreateBitCursorByTree(tree *Tree) BitCursor {
	return BitCursor{currentNode: tree.root}
}

func (cursor *BitCursor) Step(direction int) (bool, byte) {
	var isLeaf bool

	if direction == 1 {
		cursor.currentNode = cursor.currentNode.left
	} else {
		cursor.currentNode = cursor.currentNode.right
	}

	isLeaf = (cursor.currentNode.left == nil) && (cursor.currentNode.right == nil)

	return isLeaf, cursor.currentNode.value
}
