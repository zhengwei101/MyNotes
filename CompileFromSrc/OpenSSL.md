# OpenSSL编译安装

## 安装 `Strawberry Perl`

前往[Strawberry Perl](http://strawberryperl.com/)官网下载并安装 Strawberry Perl 这里选择的是目前最新版本 Strawberry Perl 5.38.2.2安装过程中务必选择完整安装 安装完成之后会自动添加环境变量，无需手动添加。

检查是否已安装

```bash
perl -v
```

## 下载源代码

前往[OpenSLL](https://www.openssl.org/source/)官网下载 openssl-3.0.13.tar.gz 并解压

## 启动命令行提示工具

使用 VS 的开发人员提示工具切换到源码目录，注意选择要编译的位数所对应的版本，不要选择cross版本。

所有的选项都要统一，即：编译32位的就启动x86，配置Makefile时为命令行为`VC-WIN32`，64位同理，本文以编译64位为例。

```bash
cd /d C:\openssl-3.0.13
```

## 配置Makefile

|版本|对应架构命令行|
|----|----|
|64位|`VC-WIN64A`|
|32位|`VC-WIN32`|

```bash
perl Configure VC-WIN64A no-asm
```

如果编译静态链接版本的二进制程序，加一个 no-shared 选项即可。

```bash
perl Configure VC-WIN64A no-asm no-shared
```

若想要自定义安装路径，添加 --prefix 选项即可

```bash
perl Configure VC-WIN64A no-asm --prefix=D:\OpenSSL\3.0.13

```

## 开始编译及安装

执行 `nmake` 开始编译 整个过程视机器配置而定

## 安装二进制和库

```bash
nmake install_sw
```

使用 `install_sw`代替 `install` 是因为默认 `install` 会生成 40M 左右的 HTML 文档，若不需要就使用 `install_sw` 仅安装二进制文件和库

## 添加环境变量

将编译后的安装目录下的 bin 文件夹添加到系统 Path 目录下 打开 cmd 使用 `openssl version -a` 测试

## 清理生成的中间文件

```bash
nmake clean
```

## NMAKE : fatal error U1064: 未找到 MAKEFILE 并且未指定目标

执行perl之前，需要下载并安装`dmake`
