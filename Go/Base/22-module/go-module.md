# Go Module

Go Module 的核心是一个名为go.mod的文件，在这个文件中存储了这个module对第三方依赖的全部信息。

## 创建go.mod文件

```bash
go mod init github.com/zhengwei101/hellomodule
```

## 使用`go mod tidy`，让Go工具自动添加依赖

`go mod tidy`命令会扫描项目的源码，并自动下载项目依赖的外部Go Module以及对应版本

我们可以通过GOMODCACHE环境变量，自定义本地module的缓存路径

```bash
go mod tidy
```

## 构建程序

```bash
go build main.go
```

## 运行程序

```bash
./main.exe
```

## 在新窗口中用curl访问该http服务

```bash
curl 127.0.0.1:8081/foo/bar
```

## go.sum文件的作用

这个文件记录了hellomodule的直接依赖和间接依赖的相关版本的hash值，用来检验本地包的真实性。在构建的时候，如果本地依赖包的hash值与go.sum文件记录的不一致，就会被拒绝构建。

## 深入 Go Module 构建模式

在上面的例子中，我们看到 go.mod 的 require 段中依赖的版本号，都符合 vX.Y.Z 的格式。

在 Go Module 构建模式下，一个符合 Go Module 要求的版本号，由前缀 v 和一个满足[语义版本](https://semver.org/)规范的版本号组成。

语义版本号分成3部分：主版本号(Major)、次版本号(Minor)和补丁版本号(Patch)

查询包的版本

```bash
$go list -m -versions github.com/sirupsen/logrus
```

可以通过`go get`手动选择 logrus 版本

```bash
$go get github.com/sirupsen/logrus@v1.7.0
go: downloading github.com/sirupsen/logrus v1.7.0
go get: downgraded github.com/sirupsen/logrus v1.8.1 => v1.7.0
```

查询当前module的所有依赖

```bash
$go list -m all
```

在Go Module构建模式下，我们无需手动维护vendor目录下的依赖包，

Go提供了可以快速建立和更新vendor的命令

```bash
$go mod vendor
$tree -LF 2 vendor
vendor
├── github.com/
│   ├── google/
│   ├── magefile/
│   └── sirupsen/
├── golang.org/
│   └── x/
└── modules.txt
```

我们看到，`go mod vendor` 命令在 vendor 目录下，创建了一份这个项目的依赖包的副本，并且通过 `vendor/modules.txt` 记录了 vendor 下的 module 以及版本。

## 小结

- Go包是Go语言的基本组成单元。一个Go程序就是一组包的集合，所有Go代码都位于包中。
- Go源码可以导入其他Go包，并使用其中的导出语法元素，包括类型、变量、函数、方法等，而且，main函数是整个Go应用的入口函数。
- Go源码需要先编译，再分发和运行。对于复杂的Go项目，我们需要在Go Module的帮助下完成项目的构建。

`go mod tidy`下载的第三方包一般在`$GOPATH/pkg/mod`下面。

如果没有设置GOPATH环境变量，一般默认值为你的home路径下的go文件夹。

这样第三包就在go文件的`pkg/mod`下面。我的电脑在`C:\Users\zhengw\go\pkg\mod`下面。
