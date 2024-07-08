# Gin 框架实战

## gin-ranking

```sh
http://book.qingwakong.com/fyouku/

fyouku.qingwakong.com/
 
go mod init gin-ranking
 
go get -u github.com/gin-gonic/gin

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

REST API Design

```go
GET   /tasks - display all tasks
POST  /tasks - create a new task
GET   /tasks/{id} - display a task by ID
PUT   /tasks/{id} - update a task by ID
DELETE   /tasks/{id} - delete a task by ID
```
