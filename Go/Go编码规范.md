# 编码规范

## 代码风格

### 代码格式

import 规范

包名称与导入路径的最后一个目录名不匹配时，或者多个相同包名冲突时，则必须使用导入别名

```go
  //bad
  "github.com/dgrijalva/jwt-go/v4"

  //good
  jwt "github.com/dgrijalva/jwt-go/v4"
```

导入的包建议进行分组，匿名包的引用使用一个新的分组，并对匿名包引用进行说明。

```go
  import (
    // go 标准包
    "fmt"

    // 第三方包
    "github.com/jinzhu/gorm"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"

    // 匿名包单独分组，并对匿名包引用进行说明
    // import mysql driver
    _ "github.com/jinzhu/gorm/dialects/mysql"

    // 内部包
    v1 "github.com/marmotedu/api/apiserver/v1"
    metav1 "github.com/marmotedu/apimachinery/pkg/meta/v1"
    "github.com/marmotedu/iam/pkg/cli/genericclioptions"
  )
```

### 声明、初始化和定义

当函数中需要使用到多个变量时，可以在函数开始处使用var声明。
在函数外部声明必须使用var，不要采用`:=`,容易踩到变量的作用域的问题。

```go
var (
    Width int
    Height int
)
```

在初始化结构引用时，请使用`&T{}`代替`new(T)`，以使其与结构体初始化一致

```go
//bad
sptr := new(T)
sptr.Name = "bar"

//good
sptr := &T{Name: "bar"}
//&T{}只能用来初始化结构体，new(T)可以初始化结构体之外的类型
//new(T)可以直接用在表达式中
```

struct 声明和初始化格式采用多行，定义如下

```go
type User struct {
    Username string
    Email string
}

user := User{
    Username: "xxx",
    Email: "xxx@xxx.com",
}
```

尽可能指定容器容量，以便为容器预先分配内存，例如：

```go
v := make(map[int]string, 4)
v := make([]string, 0, 4)
```

在顶层，使用标准var关键字。请勿指定类型，除非它与表达式的类型不同

```go
// bad
var _s string = F()

func F() string { return "A" }

// good
var _s = F()
// 由于 F 已经明确了返回一个字符串类型，因此我们没有必要显式指定 _s 的类型
func F() string { return "A" }
```

```go
```

## 命名规范
