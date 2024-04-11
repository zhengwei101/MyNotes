# Fedora 38  升级到  Feodra 39

## 通过命令行（终端）升级

1. 更新系统软件包

```bash
sudo dnf upgrade --refresh
```

2. 安装系统更新组件，安装过的可以忽略

```bash
sudo dnf install dnf-plugin-system-upgrade
```

3. 下载fedora 39

```bash
sudo dnf system-upgrade download --releasever=39
```

如果在下载过程中提示依赖损坏，可以添加 `‐‐allowerasing` 参数继续

```bash
sudo dnf system-upgrade download --releasever=39 --allowerasing
```

4. 下载完成，重启系统继续安装

```bash
sudo dnf system-upgrade reboot
```

5. 从 Fedora 39 中删除旧的和过时的软件包

```bash
sudo dnf system-upgrade clean
```

6. 在 Fedora 39 上删除损坏的 Symlinks

```bash
sudo find /usr -type l -xtype l -delete
```

如果想在清理之前查看所有已损坏的符号链接列表，可以使用以下命令来显示它们：

```bash
sudo symlinks -r /usr | grep dangling
```

   