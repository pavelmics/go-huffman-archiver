package HuffmanTree

import (
	"archiver/internal/ByteFrequencyTable"
	"sort"
)

// Tree
// Classic Huffman tree
// see https://www.siggraph.org/education/materials/HyperGraph/video/mpeg/mpegfaq/huffman_tutorial.html
type Tree struct {
	root *Node
	byteToBitMaskMap map[byte] string
}

func (tree *Tree) BuildTreeByFrequencyTable(table ByteFrequencyTable.Table) {
	freqTable := table.GetFrequencyTable()

	// из всех ячеек таблицы делаем ноды-листья (просто для удобства работы)
	var leafNodes []Node
	for i := 0; i != len(freqTable); i++ {
		leafNodes = append(leafNodes, Node{value: freqTable[i].Value, sumWeight: freqTable[i].Weight})
	}

	stopLoop := false
	for stopLoop == false {
		switch {
		case 0 == len(leafNodes):
			stopLoop = true
		case 1 == len(leafNodes):
			tree.root = &leafNodes[0]
			stopLoop = true
		case len(leafNodes) >= 2:
			// отрезаем две последние ноды от массива листьев нод
			handlingNodes := leafNodes[len(leafNodes) - 2:]
			leafNodes = leafNodes[:len(leafNodes) - 2]

			// так как они отсортированы по убыванию, то последняя будет с индексом 0
			last := handlingNodes[1]
			preLast := handlingNodes[0]

			// создаем новую ноду, которая теперь будет корневой для этих двух
			rootNode := Node{value: 0, sumWeight: last.sumWeight + preLast.sumWeight, left: &preLast, right: &last}
			tree.root = &rootNode

			leafNodes = append(leafNodes, rootNode)
			sort.Slice(leafNodes, func(i, j int) bool {
				return leafNodes[i].sumWeight > leafNodes[j].sumWeight
			})
		}
	}

	tree.buildByteToBitsMap()
}

func (tree *Tree) GetByteToBitMaskMap() map[byte] string {
	return tree.byteToBitMaskMap
}

// returns mask, that represents path to value in the tree
func (tree *Tree) GetMaskByByte(value byte) string {
	return tree.byteToBitMaskMap[value]
}

func (tree *Tree) buildByteToBitsMap() {
	tree.getBitForNode(tree.root, "")
}

func (tree *Tree) getBitForNode(node *Node, currentMask string) {
	if (node.left == nil) && (node.right == nil) {
		if tree.byteToBitMaskMap == nil {
			tree.byteToBitMaskMap = make(map[byte] string)
		}
		tree.byteToBitMaskMap[node.value] = currentMask
	}

	if node.left != nil {
		var copyCurrentMask string
		copyCurrentMask = currentMask
		copyCurrentMask = copyCurrentMask + "1"
		tree.getBitForNode(node.left, copyCurrentMask)
	}

	if node.right != nil {
		var copyCurrentMask string
		copyCurrentMask = currentMask
		copyCurrentMask = copyCurrentMask + "0"
		tree.getBitForNode(node.right, copyCurrentMask)
	}
}