# Geego 框架实战

## fyouku

```sh
http://book.qingwakong.com/fyouku/

fyouku.qingwakong.com/
```

Beego的安装

```go
go get github.com/beego/beego/v2
```

Beego的升级

```go
go get -u github.com/beego/beego/v2
```

bee工具是一个为了协助快速开发beego项目而创建的项目，
通过bee可以很容易的进行beego项目的创建、热编译、开发、测试和部署

bee工具的安装

```go
# go 1.16 以前的版本
go get -u github.com/beego/bee/v2

# go 1.16及以后的版本
go install github.com/beego/bee/v2@latest
```

安装完之后，bee 可执行文件默认存放在 $GOPATH/bin 里面，所以您需要把 $GOPATH/bin 添加到您的环境变量中。

bee工具环境变量设置

```sh
#go语言安装主根目录
export GOROOT=/usr/local/go #替换你的目录
#GOPATH 是自己的go项目路径，自定义设置
export GOPATH=/Users/ding/go_workspace #替换你的目录
#GOBIN 当我们使用go install命令编译后并且安装的二进制程序目录
export GOBIN=$GOPATH/bin
# 启用 Go Modules 功能
export GO111MODULE=on
# 配置 GOPROXY 环境变量
export GOPROXY=https://goproxy.cn,direct
export PATH=$PATH:$GOROOT/bin:$GOBIN
```

bee支持注解路由

生成路由，不生成会找不到页面

```sh
bee generate routers
```

项目根目录bee run 编译并启动程序

```sh
bee run
```

beego view 语法

`if...else...end`

`template`

`range`

 等于 \ 不等于 \小于\大于\小于等于\大于等于
` eq \  ne   \ lt \ gt \  le  \ ge `

`静态资源`
