# Redis 安装与使用

## 通过HomeBrew包管理工具安装

Fix brew warnings

```sh
brew update
brew doctor
brew upgrade
```

查看redis版本

```sh
brew search redis
```

通过`@`来指定版本，上面没指定版本，默认是最新版本

安装redis最新版本

```sh
brew install redis
```

redis配置文件`redis.conf`在`/usr/local/etc`目录下

启动redis

通过brew services来启动(后台启动)

```sh
# start
brew services start redis

# info
brew services info redis

# stop
brew services stop redis
```

通过redis默认的命令 redis-server来启动

```sh
# 指定配置文件，后台启动必须这样做
redis-server /usr/local/etc/redis.conf
# 默认参数启动，这样启动后关闭容器服务就停止了，不推荐
redis-server
```

连接服务

```sh
# 远程服务连接
redis-cli -h IP -p 6379
# 本机服务连接
redis-cli
```

关闭redis服务

```sh
redis-cli shutdown
```

## Redis 使用场景

### 1. 取最新N个数据的操作

   比如，典型的取你网站的最新文章，我们可以将最新的5000条评论的ID放在Redis的List集合中，并将超出集合部分从数据库获取。

### 2. 排行榜应用，取TOP N 操作

   这个需求与上面需求的不同之处在于，前面操作以时间为权重，这个是以某个条件为权重，

   比如，按顶的次数排序，这时候就需要我们的`sorted set`出马了，

   将你要排序的值设置成`sorted set`的`score`，

   将具体的数据设置成相应的`value`，每次只需要执行一条`ZADD`命令即可。

### 3. 需要精准设定过期时间的应用

   比如，你可以把上面说到的`sorted set`的`score`值设置为过期时间的时间戳，

   那么就可以简单地通过过期时间排序，定时清除过期数据了，

   不仅是清除Redis中的过期数据，你完全可以把Redis里这个过期时间当成是对数据库中数据的索引，

   用Redis来找出哪些数据需要过期删除，然后再精准地从数据库中删除相应的记录。

### 4. 计数器应用

   Redis的命令都是原子性的，你可以轻松地利用`INCR`, `DECR`命令来构建计数器系统。

### 5. uniq操作，获取某段时间所有数据排重值

   这个使用Redis的set数据结构最合适了，只需要不断地将数据往set中扔就行了，set意为集合，所以会自动排重。

### 6. PUb/Sub构建实时消息系统

   Redis的`Pub/Sub`系统可以构建实时的消息系统，比如很多用`Pub/Sub`构建的实时聊天系统的例子。

### 7. 构建队列系统

   使用list可以构建队列系统，使用`sorted set`甚至可以构建有优先级的队列系统。

### 8. 缓存

   最常用，性能优于Memcached(被libevent拖慢)，数据结构更多样化。

## Redis 6.0 特点

- 多IO线程机制

- 安全保护功能

- `RESP 3`协议带来的新特性

### 多IO线程机制

主线程

IO线程

epoll机制

子进程/异步线程

多IO线程机制

- 将网络IO任务使用多线程处理
- 主线程和IO线程通过任务队列交互

### 安全保护功能

密码访问

高风险命令

细粒度权限控制

### `RESP 3`协议带来的新特性

## Redis 7.0 特点
