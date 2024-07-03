# MySQL

## Mac下安装MySQL

搜索MySQL

```bash
brew search mysql
```

安装MySQL

```bash
brew install mysql@8.0
```

配置zsh终端

在`~/.zshrc`文件中添加以下内容

```bash
export PATH="/opt/homebrew/opt/mysql@8.0/bin:$PATH"
```

设置MySQL密码

先启动mysql服务，再运行`mysql_secure_installation`命令来设置root用户密码

```bash
brew services start mysql@8.0

mysql_secure_installation
```

卸载MySQL

在卸载MySQL之前，必须先停止正在运行的MySQL服务

```bash
brew services stop mysql@8.0
```

通过homebrew命令

```bash
brew uninstall mysql@8.0
```

删除MySQL相关文件

```bash
#删除MySQL的数据目录
rm -rf /opt/homebrew/var/mysql

#删除MySQL的配置文件
rm -rf /opt/homebrew/etc/my.cnf
```

清理homebrew

```bash
brew cleanup
```

验证MySQL是否已卸载

```bash
mysql --version
```

## 使用MySQL安装包安装

安装完成后，配置zsh环境

```bash
vim ~/.zshrc

#在配置文件最后一行，添加
export PATH=$PATH:/usr/local/mysql/bin

#更新配置
source ~/.zshrc

#查看MySQL版本
mysql --version
```

MySQL启动

输入`mysql -u root -p` 启动MySQL，输入密码

常用命令

```sh
开启MySQL服务
sudo mysql.server start

终止MySQL服务
sudo mysql.server stop

重新启动MySQL服务
sudo mysql.server restart

查看MySQL服务当前状态
sudo mysql.server status
```

修改密码

```sh
ALTER USER 'root'@'localhost' IDENTIFIED BY '123456';
```

这里的'root'@'localhost'指定了用户和其连接的主机。

刷新权限表，使更改立即生效

```sh
FLUSH PRIVILEGES;
```

## Linux下安装MySQL
