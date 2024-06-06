# Go 项目布局

## bookstore 项目的结构布局设计

```bash
├── cmd/
│   └── bookstore/         // 放置bookstore main包源码
│       └── main.go
├── go.mod                 // module bookstore的go.mod
├── go.sum
├── internal/              // 存放项目内部包的目录
│   └── store/
│       └── memstore.go     
├── server/                // HTTP服务器模块
│   ├── middleware/
│   │   └── middleware.go
│   └── server.go          
└── store/                 // 图书数据存储模块
    ├── factory/
    │   └── factory.go
    └── store.go
```

## 空导入包

```go
import (
 _ "bookstore/internal/store"
    //...
)
```

空导入(_ "xxx")，是有“副作用”的，即便是空导入，xxx中的init函数也会被执行，
所以如果不引用internal/store包，store包中的init函数就不会被执行，一些内存结构就没法初始化。
