# I use Python

## ImportError: `DLL load failed while importing win32api`: 找不到指定的模块

解决方法：

将`Lib\site-packages\pywin32_system32`目录中的文件

拷贝到`C:\Windows\System32`
