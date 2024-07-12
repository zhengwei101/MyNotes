# Go语言快速上手

## 入门

### 开发环境

```bash
https://go.dev
https://studygolang.com/dl
https://goproxy.cn
```

### 基于云的开发环境

```bash
https://gitpod.io/#github项目地址
```

## 基本方法

变量声明

```go
var a int = 10
//从左到右，分别为:
//1.修饰变量声明的关键字(var)
//2.变量名(a)
//3.变量类型(int)
//4.赋值操作(=)
//5.初值(10)
```

变量声明块

```go
var (
    a int = 128
    b int8 = 16
    s string = "hello"
    c rune = 'A'
    t bool = true
)
```

短变量声明

```go
a := 12
b := 'A'
c := "hello"

//也可以一次声明多个变量
a, b, c := 12, 'A', "hello"
```

声明并同时显示初始化

```go
var a = 13 //使用默认类型
var b = int32(17) //显示指定类型
var f = float32(3.14) //显示指定类型

//或者
a := 13
b := int32(17)
f := float32(3.14)
s := []byte("hello, gopher!")

//或者
var (
    a = 13
    b = int32(17)
    f = float32(3.14)
)
```

### 内置函数

1. `close`： 主要用来关闭channel
2. `len`：   用来求长度，比如string、array、slice、map、channel
3. `new`：   用来分配内存，主要用来分配值类型，比如int、struct，返回的是指针
4. `make`：  用来分配内存，主要用来分配引用类型，比如chan、map、slice
5. `append`：用来追加元素到数组、slice中
6. `panic` 和`recover` ：用来做错误处理


### 变量声明原则

- 聚类
- 就近
- 分支控制变量尽量用短变量声明形式

### 并发与并行

并发： 多线程程序在一个核的CPU上运行
并发是指同时运行多个任务，但这些任务不一定在同一时刻都在运行。

并行： 多线程程序在多个核的CPU上运行
并行是指同时运行多个任务，并且这些任务在同一时刻都在运行，通常需要使用多核CPU或者分布式计算系统来实现并行运算。

### 协程 Coroutine

协程与线程

- 协程：用户态，轻量级线程，栈KB级别
- 线程：内核态，线程跑多个协程，栈MB级别

协程的特点

- **非抢占式**多任务处理，由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或多个线程上运行

goroutine的定义

- 任何函数只需加上go就能送给调度器运行
- 不需要在定义时区分是否是异步函数
- 调度器在合适的点进行切换
- 使用-race来检测数据访问冲突

goroutine可能的切换点

- `I/O`, `select`
- `channel`
- 等待锁
- 函数调用（有时）
- `runtime.Gosched()`

以上只是参考，不能保证切换，不能保证在其他地方不切换

### 通道 Channel

理论基础： Communication Sequential Process(CSP)

`make(chan 元素类型，[缓冲大小])`

- 无缓冲通道 `make(chan int)`
- 有缓冲通道 `make(chan int, 2)`

Don't communicate by sharing memory; share memory by communicating.
不要通过共享内存来通信；通过通信来共享内存。

### 逃逸分析

在编译原理中，分析指针动态范围的方法被称之为逃逸分析。通俗来讲，当一个对象的指针被多个方法或线程引用时，则称这个指针发生了逃逸。

逃逸分析决定一个变量是分配在堆上还是分配在栈上。

### 逃逸案例

多级间接赋值容易导致逃逸，这里的多级间接指的是，对某个引用类对象中的引用类成员进行赋值
（记住公式：Data.Field = Value,如果Data,Field都是引用类的数据类型，则会导致Value逃逸。
这里的等号=不单单是赋值，也表示参数传递）。

Go语言中的引用类数据类型有func, interface, slice, map, chan, *Type

- 一个值被分享到函数栈帧范围之外
- 在for循环外声明，在for循环内分配，同理闭包
- 发送指针或者带有指针的值到channel中
- 在一个切片上存储指针或带指针的值
- slice的背后数组被重新分配了
- 在interface类型上调用方法

执行如下命令：

```sh
go build -gcflags '-m -l' main.go
```

其中 -gcflags 参数于启用编译器支持的额外标志。例如， -m 用于输出编译器的优化细节，相反可以使用-N来关闭编译器优化。而-l则用于禁用foo函数的内联优化，防止逃逸被编译器通过内联彻底的抹除。

使用反汇编命令也可以看出变量是否发生了逃逸

```sh
go tool compile -S main.go
```
  
### package

通过 go get 来获取远程依赖

`go get -u 强制从网络更新远程依赖`

## Go 依赖管理

### 环境变量 `$GOPATH`

- 通过`go.mod`文件管理依赖包版本
- 通过`go get` / `go mod`指令工具，管理依赖包

### 依赖管理三要素

1. 配置文件，描述依赖 `go.mod`
2. 中心仓库管理依赖库 `Proxy`
3. 本地工具 `go get/mod`

### Go的GOPROXY配置

```bash
服务站点URL列表，"direct"表示源站

# Linux or macOS
export GOPROXY=https://goproxy.io,direct

# PowerShell (Windows)
$env:GOPROXY = "https://goproxy.io,direct"
```

上面的配置步骤只会当次终端内生效，如何长久生效呢，这样就不用每次都去配置环境变量了。

Mac/Linux

```bash
# 设置你的 bash 环境变量
echo "export GOPROXY=https://goproxy.io,direct" >> ~/.profile && source ~/.profile

# 如果你的终端是 zsh，使用以下命令
echo "export GOPROXY=https://goproxy.io,direct" >> ~/.zshrc && source ~/.zshrc
```

Windows

1. 右键 我的电脑 -> 属性 -> 高级系统设置 -> 环境变量
2. 在 “[你的用户名]的用户变量” 中点击 ”新建“ 按钮
3. 在 “变量名” 输入框并新增 “GOPROXY”
4. 在对应的 “变量值” 输入框中新增 “<https://goproxy.io,direct”>
5. 最后点击 “确定” 按钮保存设置

### 使用go命令来配置

1.首先开启go module

```bash
go env -w GO111MODULE=on     // Windows  
export GO111MODULE=on        // macOS 或 Linux
```

2.配置goproxy

```bash
阿里云配置
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct      // Windows  
export GOPROXY=https://mirrors.aliyun.com/goproxy,direct         // macOS 或 Linux

七牛云配置
go env -w GOPROXY=https://goproxy.cn,direct     // Windows  
export GOPROXY=https://goproxy.cn,direct        // macOS 或 Linux

用|分割多个代理
export GOPROXY="https://goproxy.cn|https://goproxy.io,direct"
```

## 配置

viper(github.com/spf13/viper)

## 调试

delve(github.com/go-delve/delve)

## Go 标准库

http

bufio

log

encoding/json

regexp

time

strings/math/rand

查看标准库的文档
`godoc -http :8777`

gRPC网关

```sh
安装protoc-gen-grpc-gateway工具
go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

下载annotations包
go get github.com/googleapis/googleapis

根据proto文件生成pb.go和pb.gw.go文件
protoc -I="xxx/pkg/mod/github.com/googleapis/googleapis@vxxxx" --go_out=. --go-grpc_out=. --grpc-gateway_out=. --proto_path=./micro_service/idl greet.proto

```

## 测试

### 回归测试

### 集成测试

### 单元测试

## 性能优化与软件质量

- 软件质量至关重要
- 在保证接口稳定的前提下改进具体的实现
- 测试用例：覆盖尽可能多的场景，方便回归
- 文档：做了什么，没做什么，能达到怎样的效果
- 隔离：通过选项控制是否开启优化
- 可观测：必要的日志输出

## Grafana搭建

## Prometheus数据上报与查询

## etcd

etcd的key支持按前缀查询和监听，只需要添加WithPrefix()选项

Server向etcd注册自己，且2秒后自动过期，所以server需要每隔1秒重复注册一次。

Server一旦宕机，最多2秒后etcd会自动将其删除。

Client通过前缀获得所有存活的Server，且在前缀上安装监听器，有新Server加入或老Server被删除时，Client都能及时感知到。

## 微服务框架

Kitex是字节开源的Golang微服务RPC框架。

## Go的单元测试

在Go中我们针对包(package)编写测试代码。

测试代码与包代码放在同一目录下，并且Go要求所有测试代码都存放在以`*_test.go`结尾的文件中。

这使Go开发人员一眼就能分辨出哪些文件存放的包代码，哪些文件存放的是针对该包的测试代码。

执行单元测试时，`go test`命令会将所有包目录下的`*_test.go`文件编译成一个临时二进制文件（我们可以通过`go test -c`显示编译出该文件），并执行该文件，后者将执行各个测试源文件中的名字格式为`TestXxx`函数所代表的测试用例并输出测试执行结果。

### 课程资料

[锁Lock、线程同步、WaitGroup](https://pkg.go.dev/sync)

[Go Module](https://go.dev/blog/using-go-modules)

[单元测试概念](https://go.dev/doc/tutorial/add-a-test)

[单元测试规则](https://pkg.go.dev/testing)

[Mock测试](https://github.com/bouk/monkey)

[基准测试](https://pkg.go.dev/testing#hdr-Benchmarks)

[数据库与 SQL 概念解读](https://zhuanlan.zhihu.com/p/41576768)

[用 database/sql 建立连接并使用](https://github.com/go-sql-driver/mysql)

[DSN 相关解读](https://en.wikipedia.org/wiki/Data_source_name)

[GORM 解析](https://gorm.io/docs/index.html)

### 项目实战 - 组件及技术点

- [web框架 - Gin](https://github.com/gin-gonic/gin#quick-start)
- [分层结构设计](https://github.com/bxcodec/go-clean-arch)
- [文件操作 - 读文件](https://pkg.go.dev/io)
- [数据查询 - 索引](https://www.baike.com/wikiid/5527083834876297305?prd=result_list&view_id=5di0ak8h3ag000)
