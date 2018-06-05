package comm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 在文件targetLine后插入文本content
func Insert2File(fileName string, content string, targetLine string, insertBefore bool) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	content = strings.TrimSpace(content)
	content = content + "\n" + "\n"
	fileReader := bufio.NewReader(f)
	var result string
	for {
		line, err := fileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if strings.Contains(line, targetLine) {
			if insertBefore {
				result += content + line
			} else {
				result += line + content
			}
		} else {
			result += line
		}
	}

	f.Close()

	// 先写新文件
	newFileName := fileName + ".new"
	newf, err := os.OpenFile(newFileName, os.O_EXCL|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	fileWriter := bufio.NewWriter(newf)
	fileWriter.WriteString(result)
	fileWriter.Flush()
	newf.Close()

	// 删除old文件
	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}

	// 重命名新文件
	err = os.Rename(newFileName, fileName)
	if err != nil {
		panic(err)
	}
}
