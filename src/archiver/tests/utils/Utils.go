package utils

import (
	"archiver/internal/HuffmanTree"
	"archiver/internal/ByteFrequencyTable"
)

func CreateHuffmanTreeByString(data string) HuffmanTree.Tree {
	table := ByteFrequencyTable.Table{}
	for i := 0; i != len(data); i++ {
		table.AddValue(data[i])
	}

	tree := HuffmanTree.Tree{}
	tree.BuildTreeByFrequencyTable(table)

	return tree
}
