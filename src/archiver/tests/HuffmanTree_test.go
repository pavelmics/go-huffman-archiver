package tests

import (
	"archiver/internal/ByteFrequencyTable"
	"testing"
	"strconv"
	ht "archiver/internal/HuffmanTree"
)

const TEST_DATA_TEXT = "In computer science, a binary tree is a tree data structure in which each node has at most " +
	"two children, which are referred to as the left child and the right child. A recursive definition using just " +
	"set theory notions is that a (non-empty) binary tree is a tuple (L, S, R), where L and R are binary trees or " +
	"the empty set and S is a singleton set.[1] Some authors allow the binary tree to be the empty set as well.[2]" +
	"From a graph theory perspective, binary (and K-ary) trees as defined here are actually arborescences.[3] A binary " +
	"tree may thus be also called a bifurcating arborescence[3]—a term which appears in some very old programming " +
	"books,[4] before the modern computer science terminology prevailed. It is also possible to interpret a binary " +
	"tree as an undirected, rather than a directed graph, in which case a binary tree is an ordered, rooted tree.[5] " +
	"Some authors use rooted binary tree instead of binary tree to emphasize the fact that the tree is rooted, but as " +
	"defined above, a binary tree is always rooted.[6] A binary tree is a special case of an ordered K-ary tree, " +
	"where k is 2.In mathematics, what is termed binary tree can vary significantly from author to author. " +
	"Some use the definition commonly used in computer science,[7] but others define it as every non-leaf having " +
	"exactly two children and don't necessarily order (as left/right) the children either.[8]In computing, binary " +
	"trees are used in two very different ways"

func TestHuffmanTree(t *testing.T) {
	table := ByteFrequencyTable.Table{}
	testData := TEST_DATA_TEXT
	for i := 0; i != len(testData); i++ {
		table.AddValue(testData[i])
	}

	tree := ht.Tree{}
	tree.BuildTreeByFrequencyTable(table)
	byteToBitMap := tree.GetByteToBitMaskMap()
	for testedByte := range byteToBitMap {
		var mask string
		mask = byteToBitMap[testedByte]
		cursor := ht.CreateBitCursorByTree(&tree)
		maskArray := make([]int, len(mask))
		var value byte
		for i := 0; i != len(mask); i++ {
			maskArray[i], _ = strconv.Atoi(string(mask[i]))
		}
		for i := 0; i != len(maskArray); i++ {
			_, value = cursor.Step(int(maskArray[i]))
		}
		if value != testedByte {
			t.Error("The tree leaf value and map value is not matched")
		}
	}
}