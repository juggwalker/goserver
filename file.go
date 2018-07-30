package main

import (
	"fmt"
	"io"
	"os"
)

type FileInfo struct {
	FilePath string
	FileName string
}

//使用bufio包中Writer对象
func (f *FileInfo) WriteWithBufio(content interface{}) interface{} {
	return nil
}

func (f *FileInfo) WriteWithIo(content string) interface{} {
	filename := fmt.Sprintf("%s/%s", f.FilePath, f.FileName)
	fileObj, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		Log.Warn("Failed to open", err.Error())
		return nil
	}
	if _, err := io.WriteString(fileObj, content); err == nil {
		return len(content)
	}
	return nil
}
