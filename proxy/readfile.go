package proxy

import "fmt"

type IReadFile interface {
	ReadFile(string) string
}

type ReadFile struct {
}

func (r ReadFile) ReadFile(filename string) string {
	return fmt.Sprintf("文件名为：%s 文件内容为：保密内容", filename)
}
