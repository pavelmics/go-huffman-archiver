package ArchiveFile

import "archiver/internal/HuffmanTree"

type Metadata struct {

	huffmanTree HuffmanTree.Tree

	dataByteCount int  // count of bytes in archived data

	lastByteBitCount int  // count of used bytes in the last byte of archived data sequence (we can think about it as about bit sequence)
}
