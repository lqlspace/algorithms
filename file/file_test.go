package fileoprts

import (
	"testing"
	"fmt"
)

func PrintStr(line string) {
	fmt.Println(line, ";")
}

func PrintBin(line []byte) {
	fmt.Println(line)
}

func TestReadLargeFile(t *testing.T) {
	filePath := "test.txt"
	if err := WriteFileByLine(filePath); err != nil {
		t.Error("[error] write file error!")
		return
	}

	if err := ReadLargeFile(filePath, PrintStr); err != nil {
		t.Error("[error] read large file failed!")
	}

}

func TestReadLargeFileBin(t *testing.T) {
	filePath := "test.bin"
	if err := WriteFileBinary(filePath); err != nil {
		t.Error("[error] write file error!")
		return
	}

	if err := ReadLargeFileBin(filePath, PrintBin); err != nil {
		t.Error("[error] read file error!")
	}
}
