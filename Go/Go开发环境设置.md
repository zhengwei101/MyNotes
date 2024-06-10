# Linux-Go开发环境设置

首先，在 Git 中，我们会把非 ASCII 字符叫做 Unusual 字符。

这类字符在 Git 输出到终端的时候默认是用 8 进制转义字符输出的（以防乱码），

但现在的终端多数都支持直接显示非 ASCII 字符，所以我们可以关闭掉这个特性，

具体的命令如下:

```bash
git config --global core.quotepath off
```

其次，GitHub 限制最大只能克隆 100M 的单个文件，

为了能够克隆大于 100M 的文件，我们还需要安装 `Git Large File Storage`，安装方式如下

```bash
git lfs install --skip-repo
```

## Go编译环境设置

Go安装包下载

```bash
wget -P /tmp/ https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
```

解压和安装

```bash
mkdir -p $HOME/go
tar -xvzf /tmp/go1.22.4.linux-amd64.tar.gz -C $HOME/go
mv $HOME/go/go $HOME/go/go1.22.4
```

修改`$HOME/.bashrc`文件

确认一下命令行配置文件是哪个：

- 如果用的是zsh，那么对应的文件一般是`~/.zshrc`
- 如果用的是bash，那么对应的文件一般是`~/.bashrc`或者`~/.bash_profile`
- 通过执行`echo $SHELL`可以准确找到是哪个shell

```bash
#Basic envs
#export LANG="en_US.UTF-8" # 设置系统语言为en_US.UTF-8，避免终端出现中文乱码
export WORKSPACE="$HOME/Work" # 设置工作目录

#Go envs
export GOVERSION=go1.22.4 # Go 版本设置
export GO_INSTALL_DIR=$HOME/go # Go 安装目录
export GOROOT=$GO_INSTALL_DIR/$GOVERSION # GOROOT 设置
export GOPATH=$WORKSPACE/golang # GOPATH 设置
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH # 将GO语言自带和通过go install安装的二进制文件加入到PATH路径中
export GO111MODULE="on" # 开启 Go moudles 特性
#export GOPROXY=https://goproxy.cn,direct # 安装Go模块时，代理服务器设置
export GOPRIVATE=
export GOSUMDB=off # 关闭校验 Go 依赖包的哈希值

```

在使用模块时，`$GOPATH`是无意义的，不过它还是会把下载的依赖，储存在`$GOPATH/pkg/mod`目标中，
也会把`go install`的二进制文件存放在`$GOPATH/bin`目录中。

## 安装 `protobuf`

Protocol buffer compiler

### 使用系统自带protobuf

```bash
# 安装
sudo dnf install protobuf-compiler

# 卸载
sudo dnf remove protobuf-compiler
```

### 使用最新的安装包

```bash
# 1. Download the zip file
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v27.1/protoc-27.1-linux-x86_64.zip

# 2. Unzip the file under $HOME/.local or a directory of your choice. 
# For example:
unzip protoc-27.1-linux-x86_64.zip -d $HOME/.local

# 3. Update your environment’s path variable to include the path to the protoc executable. 
# For example:
export PATH="$PATH:$HOME/.local/bin"
```

Go plugins for the protocol compiler

### 安装 `protoc-gen-go`

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### 安装 `protoc-gen-go-grpc`

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 初始化工作区

当前版本的go支持多模块工作区，所以这里也需要初始化工作区。初始化命令如下：

```bash
mkdir -p $GOPATH && cd $GOPATH
go work init
go env GOWORK # 执行此命令，查看 go.work 工作区文件路径
/home/wilbur/Work/golang/go.work
```
