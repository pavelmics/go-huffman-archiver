package ArchiveFile

import (
	"archiver/internal/HuffmanTree"
	"io"
	"math"
)

// describe the archive file
//
type File struct {
	metadata Metadata
	data []byte
	currentByte int
}

// Creates archive from readFrom io stream and write it to writeTo io stream
// Ready huffman tree is needed for compressing data
func CreateFromStreams(readFrom io.Reader, writeTo io.Writer, tree *HuffmanTree.Tree) {
	// Create metadata
	metadata := Metadata{}


	mask := ""
	endOrFileReached := false
	for true {
		inputData := make([]byte, 1)
		_, err := readFrom.Read(inputData)
		if err != nil {
			if err == io.EOF {
				endOrFileReached = true
			}
		}
		byteMask := tree.GetMaskByByte(inputData[0])
		mask += byteMask
		if len(mask) >= 8 {
			byteCode := mask[0:7]
			mask = mask[7:]
			metadata.dataByteCount++
			writeTo.Write([]byte{stringToByte(byteCode)})
		}

		if endOrFileReached {
			if len(mask) != 0 {
				metadata.dataByteCount++
				// добиваем до 8 и записываем
			}

		}
	}
}


func stringToByte(str string) byte {
	sum := 0
	for i := 0; i != len(str); i++ {
		pow := (len(str) - i) - 1
      	sum += int(math.Pow(2, float64(pow)))
	}

	return byte(sum)
}