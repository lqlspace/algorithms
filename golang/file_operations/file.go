package fileoprts

import (
	"strings"
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
	"io"
	"bytes"
	"encoding/binary"
)

//逐行写入文本文件
func WriteFileByLine(filePath string) error {
	f1, err := os.Create(filePath)
	if err != nil {
			return err
	}
	br := bufio.NewWriter(f1)
	rand.Seed(time.Now().UnixNano())


	for i := 0; i < 200; i++ {
		var tmp string
		for j := 0; j < 50; j++ {
			tmp = tmp + strconv.Itoa(rand.Intn(10000)) + " "
		}
		tmp += "\n"
		_, err := br.WriteString(tmp)
		if err != nil {
			return err
		}
	}
	br.Flush()

	return nil
}


//写入二进制文件,golang提供了binary和gob库；
/*func WriteFileBinary(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte("hello, I'am allen!"))
	if err != nil {
		return err
	}
	return nil
}*/

func WriteFileBinary(fileName string) error {
	fp, _ := os.Create(fileName)
	defer fp.Close()

	buf := new(bytes.Buffer)
	for i := 0; i < 10; i++ {
		binary.Write(buf, binary.LittleEndian, int32(i))
		fp.Write(buf.Bytes())
	}
	return nil
}



//读取超大文件时，文件内容不可能一次性读完然后返回，此时可以采用如下两种方法：
//1、流处理方式 —— 文本文件
//2、分片处理 —— 二进制文件

//1、逐行读取文本文件
func ReadLargeFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}

	buf := bufio.NewReader(f)

	for {
		str, err := buf.ReadString('\n')
		str = strings.TrimSpace(str)
		handle(str)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}


//2、读取二进制文件
func ReadLargeFileBin(fileName string, handle func([]byte)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	s := make([]byte, 4096)
	for {
		nr, err := f.Read(s[:])
		handle(s[0:nr])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

	return nil
}
