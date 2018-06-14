package tests

import (
	"testing"
	"archiver/tests/utils"
	"archiver/internal/ArchiveFile"
	"strings"
	"io"
	"bytes"
)

func TestCreateFromStreams(t *testing.T) {
	testData := "aaa bbbbbbbb ccc dd eeeeeeee f"

	var ioReader io.Reader
	var ioWriter io.Writer

	// create test tree
	tree := utils.CreateHuffmanTreeByString(testData)

	// create io streams
	ioReader = strings.NewReader(testData)
	ioWriter = &bytes.Buffer{}

	ArchiveFile.CreateFromStreams(ioReader, ioWriter, &tree)
}


