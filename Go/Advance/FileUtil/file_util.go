package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
	"strings"
)

func RenameFile(src, dst string) (err error) {
	_, err = os.Lstat(src)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist:", src)
		return nil
	}
	err = os.Rename(src, dst)
	if err != nil {
		panic(err)
	} else {
		println("文件重命名成功")
	}
	return err
}

func NewFileName(src, tag string) string {
	if strings.Contains(src, tag) {
		newName := strings.Replace(src, tag, "", -1)
		return newName
	}
	return src
}

type Options struct {
	Dir string `short:"d" long:"dir" description:"Directory for scan"`
	Tag string `short:"t" long:"tag" description:"Tag to removed of filename"`
}

func main() {
	//0. 获取命令行参数
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	//从命令行参数中，获取文件夹路径和要去除字符串的名称
	fmt.Println("dir:", opts.Dir, "tag:", opts.Tag)
	//./file_util -d "/Users/zhengw/Downloads/test" -t "【www.xxx.cn】"

	//1. 遍历指定文件夹
	paths := make([]string, 0)
	err = filepath.Walk(opts.Dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//2. 获取文件名
	for _, path := range paths {
		newPath := NewFileName(path, opts.Tag)
		fmt.Println(newPath)
		//3. 文件重命名为新的名称
		err := RenameFile(path, newPath)
		if err != nil {
			return
		}
	}
}
