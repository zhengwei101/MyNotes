package main

import (
	"advance/errhandling/filelistserver/filelist"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 在errWrapper中，实现统一的错误处理逻辑
// 把输入的函数包装一下，返回新的函数来输出
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic: %v", r)
				code := http.StatusInternalServerError
				http.Error(w, http.StatusText(code), code)
			}
		}()

		err := handler(w, r)

		if err != nil {
			//gopm中的日志函数
			//log.Warn("Error handling request: %s", err.Error())
			log.Printf("Error handling request: %s\n", err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(w, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelist.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
