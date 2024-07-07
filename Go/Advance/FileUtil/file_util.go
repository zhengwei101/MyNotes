package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jessevdk/go-flags"
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

func renameDir(src, dst string) (err error) {
	if src == dst {
		return nil
	}
	fileInfo, err := os.Stat(src)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist:", src)
		return nil
	}
	if !fileInfo.IsDir() {
		fmt.Println("File is not a directory:", src)
	}
	err = os.Rename(src, dst)
	if err != nil {
		panic(err)
	} else {
		println("文件夹重命名成功")
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

func RenameDirs(dirPath, tag string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			path := filepath.Join(dirPath, file.Name())
			newPath := NewFileName(path, tag)
			if newPath == path {
				err := RenameDirs(path, tag)
				if err != nil {
					return err
				}
			}
			fmt.Println("rename dir: ", path, " to ", newPath)
			//3. 文件夹重命名为新的名称
			err := renameDir(path, newPath)
			if err != nil {
				panic(err)
			}
		}
	}
	return err
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
	err = RenameDirs(opts.Dir, opts.Tag)
	if err != nil {
		fmt.Println(err)
		return
	}

	paths := make([]string, 0)
	err = filepath.Walk(opts.Dir, func(path string, info os.FileInfo, err error) error {
		if len(path) > 0 && !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	//2. 获取文件名
	if len(paths) == 0 {
		return
	}
	for _, path := range paths {
		newPath := NewFileName(path, opts.Tag)
		if newPath == path {
			continue
		}
		fmt.Println("rename ", path, " to ", newPath)
		//3. 文件重命名为新的名称
		err := RenameFile(path, newPath)
		if err != nil {
			panic(err)
		}
	}

}
