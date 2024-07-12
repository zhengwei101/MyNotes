# Geego 框架实战

Beego的特性

1. 简单化
   > RESTful支持、MVC模型，可以使用bee工具快速地开发应用，包括监控代码修改进行热编译、自动化测试代码以及自动化打包部署。
2. 智能化
   > 支持智能路由、智能监控、可以监控QPS、内存消耗、CPU使用、以及goroutine的运行状况，让你的线上应用尽在掌握。
3. 模块化
   > beego内置了强大的模块，包括Session、缓存操作、日志记录、配置解析、性能监控、上下文操作、ORM模块，请求模拟等，足以支撑你任务的应用。
4. 高性能

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

发布应用的时候打包，会把项目打包成zip包，这样我们部署的时候，直接把打包之后的项目上传，解压就可以部署了。

```sh
bee pack

# mac下，打包Linux系统执行文件
bee pack -be GOOS=linux
```

### 配置信息

beego提供了专门配置文件，配置信息定义在conf文件夹中的app.conf

```conf
appname = fyouku //项目名称
httpport = 8098  //访问端口号
runmode = dev    //运行环境，下面会根据运行环境加载不同配置

[dev]
apiurl = http://127.0.0.1:8099
microApi = http://127.0.0.1:8085
defaultdb = root:123456@tcp(127.0.0.1:3306)/fyouku?charset=utf8

[prod]
apiurl = http://127.0.0.1:8099
microApi = http://127.0.0.1:8085
```

### beego view 语法

1. 统一使用`{{ }}`作为左右标签
2. 使用` . `来访问当前位置的上下文
3. 使用` $ `来引用当前模板根级的上下文
4. 使用` $var `来访问创建的变量

判断语句 `if...else...end`

```template
{{if eq .IsEmail 1}}
    <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
{{else}}
    <a class="email" href="#">不允许访问Email</a>
{{end}}
```

循环语句 `range`

```template
{{range $index,$value := .Pages}}
    {{$index}} - {{$value.Num}} of {{$.Website}}
    <br>
{{end}}
```

载入其它模板

```template
{{template "head.html" .}}
```

 等于 \ 不等于 \小于\大于\小于等于\大于等于
` eq \  ne   \ lt \ gt \  le  \ ge `

`静态资源`

ORM 必须注册一个别名为 default 的数据库，作为默认使用

```go
//获取配置文件中信息
defaultdb := beego.AppConfig.String("defaultdb")
orm.RegisterDriver("mysql", orm.DRMySQL)
orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)
```

ORM 的使用

```go
//操作数据库都需要定义struct和表结构对应
type User struct {
    Id int
    Name string
    AddTime int64
    Status int
    Mobile string
    Avatar string
}

//初始化注册对应model
func init() {
    orm.RegisterModel(new(User))
}

//获取用户信息
func UserInfo(id int) (User, error) {
    //通过orm中的Read函数获取
    var (
        err error
    )
    o := orm.NewOrm()
    user := User{Id: id}
    err = o.Read(&user)
    return user, err
}

//获取用户列表
func List() ([]User, error){
    var (
        users []User
        err error
    )
    o := orm.NewOrm()

    //声明操作的表
    qs := o.QueryTable("user")
    //条件id大于10
    qs = qs.Filter("id__gt",10)
    //返回几条数据
    qs = qs.Limit(2)
    //倒序是前面加上负号
    qs = qs.OrderBy("-id")
    //后面是设置返回的字段
    qs.All(&users, "Id", "Name")
    return users, err
}

// 原生sql删除用户
func SqlDelete(id int) error {
    o := orm.NewOrm()
    _, err := o.Raw("DELETE FROM user WHERE id=?", id).Exec()
    return err
}

```

## fyouku 项目技术栈

```sh
http://book.qingwakong.com/fyouku/

fyouku.qingwakong.com/
```

### 阿里云

- 点播播放器
- 上传SDK

### Jmeter

- 工具安装
- 测试流程
  - 线程组
  - 设置线程数
  - ramp-up: 每秒增加用户
  - 循环次数
  - 聚合报告
  - 结果树
  - 常数吞吐量定时器
- 指标
  - 并发数
  - 吞吐量
  - 90%请求时间
  - 平均请求时间

### 微服务

- micro
  - micro api
  - micro web
- go-micro
- etcd
- protobuf

### 并发编程

- goroutine go
- channel
- select

### beego

- 环境搭建
- bee工具
- 路由和过滤器
- view模板
- orm和数据库操作

### redis

- 安装
- 基本命令
- 命令封装和使用

### rabbitMQ

- 安装
- 5种工作模式讲解
- 根据工作模式封装
- 消息持久化和手动应答
- 死锁队列

### ElasticSearch

- es、ik分词器、head安装和基础讲解
- head工具使用
- 分词和查找原理
- 分词库和屏蔽词库
- 创建索引和mapping
- es增删改查接口的调用和封装

## 性能改造计划

线上常遇到的问题

- 突然502
- 接口timeout
- 接口访问慢的像蜗牛
- 数据库锁死，主从不同步
- 评论功能上线，为啥全站不能访问了

问题举例

1. 运营上线新活动，没想这么火
    > 瞬时访问量增大，导致数据量巨大，数据库读表缓慢，进而导致数据库主从同步延迟，数据库瘫痪，导致全站受影响。
2. 随着评论的不断增加，数据上升到千万级
    > 评论列表接口访问一次需要3~5秒，或者超时
3. 提交视频剧集保存为什么要等待3~5秒时间
   > 因为提交视频后，进行的逻辑操作太多，保存视频剧集信息，更新视频信息，更新用户信息，发送消息，通知关注人等等

## 改造方案总结

### 列表功能

内容关联表比较多，联表查询数据量大时，数据库瓶颈

优化：

- 关联表中的信息单独保存到redis中
- 获取关联表中信息采用并发思想获取

### 详情功能

不同用户访问详情内容都是一样的，每次访问数据库浪费资源（用户信息，视频信息）

优化：

- 保存到redis中，每次从redis访问，减少数据库压力

### 发布功能

发布以后，后续操作很多，用户等待时间长（发布视频后，保存视频、更新用户统计、更新视频统计等等）

优化：

- 采用MQ，保存完视频以后立即给用户答复。其余更新操作放到队列中，由后台程序操作，提高用户体验。
- 需要及时更新的数据，因为都存在redis中，只需要先更新redis

### 延迟功能

延迟发送功能，代码级实现比较复杂

优化：

- 利用MQ死信功能简单实现

### 搜索功能

搜索功能，直接访问mysql数据库使用`like`查询，效率太低并且容易出现无响应情况

优化：

- 采用ElaticSearch，不仅可以提供全文检索功能，也只可以实现列表功能。

### 批量发送消息功能

单次发送大量数据，浏览器相应时间过长超时；数据库写操作压力过大

优化：

- 可以把大量数据逐条放到MQ队列中，由后台程序逐条操作。
