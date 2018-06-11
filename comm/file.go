package comm

import (
	"bufio"
	// "fmt"
	"io"
	"os"
	"strings"
)

func CopyFile(dstName, srcName string) {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}

type InsertItem struct {
	TargetLine   string
	Content      string
	InsertBefore bool
}

// 在文件targetLine后插入文本content
func Insert2File(fileName string, insertList []*InsertItem) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

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

		isContain := false
		for _, insertItem := range insertList {
			targetLine := insertItem.TargetLine
			content := insertItem.Content
			content = strings.TrimRight(content, "\n")
			content = content + "\n" + "\n"

			if strings.Contains(line, targetLine) {
				isContain = true
				if insertItem.InsertBefore {
					result += content + line
				} else {
					result += line + content
				}
				break
			}
		}

		if !isContain {
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
