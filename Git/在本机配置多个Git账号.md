# 多Git账号配置

## 1. 移除全局配置

```bash
查看全局用户名
git config --global user.name
移除全局配置账户
git config --global --unset user.name

查看全局邮箱
git config --global user.email
移除全局配置邮箱
git config --global --unset user.email

查看全局密码
git config --global user.password
移除全局密码
git config --global --unset user.password

```

## 2. 在`.ssh`目录增加config文件
>
> 在`.ssh`目录下，手动新增一个config文本文件，为每个账号配置一个Host节点。

主要配置项说明：

```bash
# 账号1配置
# 主机别名
Host freelife365
# 服务器地址
HostName github.com
# 认证方式
PreferredAuthentications publickey
# 用户名
User git
# 私钥文件路径
IdentityFile ~/.ssh/id_ed25519

# 账号2配置
Host zhengwei101
HostName github.com
User git
PreferredAuthentications publickey
IdentityFile ~/.ssh/id_ed25519_zhengwei101
```

## 3. 修改局部Remote origin

为每给`repository`设置`remote origin`

```bash
git remote rm origin
# 注意，这里的git remote add origin git@<填写config中对应用的Host主机别名>:xxx/xxx.git
git remote add origin git@zhengwei101:zhengwei101/MyNotes.git
```

每个用户涉及都repository都要进行设置

## 4. 配置局部用户名和邮件

为每个repository设置局部用户名和邮件(注意：每个用户涉及都仓库都要设置

```bash
配置局部用户名
git config  user.name "wilbur_z"
配置局部邮件
git config  user.email "zhengwei101@aliyun.com"
```

## 5. 终端测试SSH Key是否生效

执行以下命令`ssh -vT git@<config中配置都主机别名>`

```bash
ssh -vT zhengwei101
```
