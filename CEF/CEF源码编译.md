# CEF源码编译指南

## 目录结构

> D:/Code

```bash
|--Code/
|  |--depot_tools/
|  |--automate/
|  |  |--automate-git.py
|  |--chromium_git/
|  |  |--chromium/
|  |  |  |--src/
|  |  |  |  |--cef/
|  |  |  |  |  |--binary_distrib/
|  |  |  |  |--out/
|  |  |  |  |  |--Release_GN_x86/
|  |  |--cef
|  |  |  |--cef_create_projects.bat
```

## 编译前准备

### 1. 创建文件夹

`D:\Code`

### 2. 创建`depot_tools`文件夹，存放`depot_tools.zip`解压后的文件

<https://storage.googleapis.com/chrome-infra/depot_tools.zip>

### 3. 创建`automate`文件夹，存放`automate-git.py`文件

<https://bitbucket.org/chromiumembedded/cef/raw/master/tools/automate/automate-git.py>

### 4. 将`depot_tools`所在目录添加到系统环境变量

### 5. 运行解压出来的 `depot_tools\update_depot_tools.bat`

### 6. 此时会更新 depot_tools 工具到最新版本，当更新完成后

在 chromium_git 目录下创建一个名为`update.bat`的批处理文件，内容如下：

```bash
set GN_DEFINES=ffmpeg_branding=Chrome proprietary_codecs=true is_official_build=true use_jumbo_build=true 
set GN_ARGUMENTS=--ide=vs2019 --sln=cef --filters=//cef/*

python ..\automate\automate-git.py --download-dir=D:\Code\chromium_git --depot-tools-dir=D:\Code\depot_tools --branch=5414 --no-build --no-distrib --force-clean --with-pgo-profiles
```

批处理的作用就是去下载 CEF 和 Chromium 代码并重置到你指定的版本，不立即编译也不做打包操作。

批处理中的各个参数及具体意思如下：

1）GN_DEFINES 参数介绍：

`- ffmpeg_branding` 和 `proprietary_codecs` 表示开启多媒体编解码支持，但默认仅支持一小部分，后面介绍如何支持更多。

`- is_official_build` 决定了是否是编译正式版本，指定该参数为 true 基本上都是为了产品发布使用，同时也会在创建解决方案的时候生成带有 sandbox 的解决方案（如 Release_GN_x86_sandbox），而不指定这个参数是没有的。如果你不是为了贡献代码，那默认这个都加上吧。

`- use_jumbo_build` 官方资料默认指定，表示是否启用试验性的 jumbo 编译，编译过程会加快很多（至少快 1 小时），但是占用 CPU 和内存（尤其是内存）会剧增。

`- is_component_build` 官方资料默认指定，但我们没有开启，这个参数表示是否启用组件化编译，设置为 true 以后，base、ffmpeg 等等都会被编译为动态库，使用时也是动态链接，编译出来的 cef_sandbox.lib 只有几兆大小，并且你需要复制很多 dll 文件到项目目录下才能运行。

2）automate-git.py 参数介绍

- `--branch` 表示你要下载哪个版本的代码，CEF 每个版本都有固定的分支，你去 CEF 项目页查看分支名称指定即可。
- `--no-build` 表示只下载代码而不编译，这里只为下载代码，我们还要修改支持多媒体的参数，所以不进行编译
- `--no-distrib` 不执行打包项目，这里只为下载代码，我们还要修改支持多媒体的参数，所以不进行打包
- `--force-clean` 如果你曾经执行过这个脚本，可能会出错，则加上这个参数，它执行清理残留文件（你也可以手动在 chromium 源码目录执行 `git clean -xdf`来清理目录中的多余内容）。

automate-git.py 的其他参数可以手动执行 `python automate-git.py --help` 来查看

执行上面的脚本如果没有意外，代码就已经下载好了，并且自动切换到了你要打包的分支并下载了所依赖的第三方库。

下载速度依赖你的网络环境，最好是稳定、高速的 VPN 场景。

## 编译项目

要开启更多多媒体编解码的支持，需要去修改 ffmpeg 的头文件。

与以前的版本不太一样的是，新版本 ffmpeg 的配置全部改为单独的宏了。

用编辑器打开 `chromium_git\chromium\src\third_party\ffmpeg\chromium\config\Chrome\win\ia32\config.h` 头文件，修改你要开启的编解码能力宏为 1 即可。

我修改了如下内容，你根据你的需要搜索修改就可以了。

```c++
#define CONFIG_FLV_DECODER 1
#define CONFIG_H263_DECODER 1
#define CONFIG_H263I_DECODER 1
#define CONFIG_MPEG4_DECODER 1
#define CONFIG_MPEGVIDEO_DECODER 1
#define CONFIG_MSMPEG4V1_DECODER 1
#define CONFIG_MSMPEG4V2_DECODER 1
#define CONFIG_MSMPEG4V3_DECODER 1
#define CONFIG_RV10_DECODER 1
#define CONFIG_RV20_DECODER 1
#define CONFIG_RV30_DECODER 1
#define CONFIG_RV40_DECODER 1
#define CONFIG_AC3_DECODER 1
#define CONFIG_AMRNB_DECODER 1
#define CONFIG_AMRWB_DECODER 1
#define CONFIG_COOK_DECODER 1
#define CONFIG_SIPR_DECODER 1
#define CONFIG_FLV_ENCODER 1
#define CONFIG_H263_ENCODER 1
#define CONFIG_MPEG4_ENCODER 1
#define CONFIG_MSMPEG4V2_ENCODER 1
#define CONFIG_MSMPEG4V3_ENCODER 1
#define CONFIG_RV10_ENCODER 1
#define CONFIG_RV20_ENCODER 1
#define CONFIG_AAC_ENCODER 1
#define CONFIG_AC3_ENCODER 1
#define CONFIG_AC3_PARSER 1
#define CONFIG_COOK_PARSER 1
#define CONFIG_H263_PARSER 1
#define CONFIG_MPEG4VIDEO_PARSER 1
#define CONFIG_MPEGVIDEO_PARSER 1
#define CONFIG_RV30_PARSER 1
#define CONFIG_RV40_PARSER 1
#define CONFIG_SIPR_PARSER 1
#define CONFIG_AC3_DEMUXER 1
#define CONFIG_AMR_DEMUXER 1
#define CONFIG_AMRNB_DEMUXER 1
#define CONFIG_AMRWB_DEMUXER 1
#define CONFIG_AVI_DEMUXER 1
#define CONFIG_AVISYNTH_DEMUXER 1
#define CONFIG_FLV_DEMUXER 1
#define CONFIG_H263_DEMUXER 1
#define CONFIG_H264_DEMUXER 1
#define CONFIG_MPEGTS_DEMUXER 1
#define CONFIG_MPEGTSRAW_DEMUXER 1
#define CONFIG_MPEGVIDEO_DEMUXER 1
#define CONFIG_RM_DEMUXER 1
#define CONFIG_AC3_MUXER 1
#define CONFIG_AMR_MUXER 1
#define CONFIG_AVI_MUXER 1
#define CONFIG_FLV_MUXER 1
#define CONFIG_H263_MUXER 1
#define CONFIG_H264_MUXER 1
#define CONFIG_MPEGTS_MUXER 1
#define CONFIG_RM_MUXER 1
```

这里需要注意的是，`#define CONFIG_SIPR_PARSER 1` 宏一定要打开，否则会在编译时期报错，找不到指定变量。

执行 `D:\Code\chromium_git\update_mp4.bat` 更新`config.h`头文件，新版本多了个`config_components.h`，编译`5414`分支时，是更新`config_components.h`。

修改完 `ffmpeg` 的配置以后，我们就开始创建工程了，在 `update.bat` 相同目录下创建一个 `create.bat`，官方资料中说是创建在 `chromium_git\chromium\src\cef` 目录下的，我为了方便管理修改了一下执行目录。

create.bat 内容如下：

```bash
set GN_DEFINES=ffmpeg_branding=Chrome proprietary_codecs=true is_official_build=true chrome_pgo_phase=0 use_jumbo_build=true 
set GN_ARGUMENTS=--ide=vs2019 --sln=cef --filters=//cef/*

rem 进入 cef 目录创建项目解决方案
cd .\chromium\src\cef
call cef_create_projects.bat
```

脚本中设置了一些宏定义，然后进入 `chromium_git\chromium\src\cef` 目录，执行了 `cef_create_projects.bat` 创建项目。

```bash
Generating CEF project files...
Generating Visual Studio projects took 1702ms
Done. Made 18580 targets from 3128 files in 44095ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Debug_GN_x64\obj\cef\libcef.ninja
Generating Visual Studio projects took 9499ms
Done. Made 18911 targets from 3125 files in 46955ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Release_GN_x86\obj\cef\libcef.ninja
Creating D:\Code\chromium_git\chromium\src\out\Debug_GN_x64_sandbox directory.
Generating Visual Studio projects took 10309ms
Done. Made 17917 targets from 3134 files in 76789ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Debug_GN_x64_sandbox\obj\cef\libcef.ninja
Creating D:\Code\chromium_git\chromium\src\out\Release_GN_x64 directory.
Generating Visual Studio projects took 10494ms
Done. Made 18579 targets from 3128 files in 79303ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Release_GN_x64\obj\cef\libcef.ninja
Creating D:\Code\chromium_git\chromium\src\out\Release_GN_x86_sandbox directory.
Generating Visual Studio projects took 10183ms
Done. Made 18524 targets from 3130 files in 72791ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Release_GN_x86_sandbox\obj\cef\libcef.ninja
Creating D:\Code\chromium_git\chromium\src\out\Debug_GN_x86 directory.
Generating Visual Studio projects took 10451ms
Done. Made 18912 targets from 3125 files in 73202ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Debug_GN_x86\obj\cef\libcef.ninja
Creating D:\Code\chromium_git\chromium\src\out\Release_GN_x64_sandbox directory.
Generating Visual Studio projects took 10108ms
Done. Made 18379 targets from 3133 files in 83759ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Release_GN_x64_sandbox\obj\cef\libcef.ninja
Creating D:\Code\chromium_git\chromium\src\out\Debug_GN_x86_sandbox directory.
Generating Visual Studio projects took 10467ms
Done. Made 18062 targets from 3131 files in 79698ms
Applying issue #1999 fix to D:\Code\chromium_git\chromium\src\out\Debug_GN_x86_sandbox\obj\cef\libcef.ninja
PS D:\Code\chromium_git>
```

此时工程文件都创建好了，你可以在 `chromium_git\chromium\src\out` 目录下看到这些工程目录。

根据你的需要编译指定版本，命令提示符切换当前目录到 chromium_git\chromium\src 目录下，

根据你的需要执行 Release 还是 Debug 版本的编译：

```bash
cd D:\Code\chromium_git\chromium\src

选择需要的版本进行编译：

ninja -C out\Debug_GN_x86 cef
ninja -C out\Release_GN_x86 cef

ninja -C out\Debug_GN_x86_sandbox cef_sandbox
ninja -C out\Release_GN_x86_sandbox cef_sandbox

ninja -C out\Debug_GN_x64 cef
ninja -C out\Release_GN_x64 cef

ninja -C out\Debug_GN_x64_sandbox cef_sandbox
ninja -C out\Release_GN_x64_sandbox cef_sandbox
```

其中 `Release_GN_x86` 就是 out 目录下的目录名，根据你自己的需要编译不同版本的就可以了，

单独编译 Debug 是无法使用脚本去打包的，脚本中打包要么是单独的 Release，要么就是 `Debug + Release` 打包。

编译速度视机器性能而定，

编译过程中消耗 CPU 和内存比较多，所以这个期间你这台电脑也别想做其他的事情了，老老实实的放在那里让它编译。

## 打包

打包 (你可能不需要此步骤，直接使用编译出来的libcef.dll替换官方编译包中的即可)

```bash
cd D:\Code\chromium_git\chromium\src\cef\tools
make_distrib.bat --ninja-build --minimal

指定目录
make_distrib.bat --output-dir=D:\Code\cef_out --ninja-build

指定版本
make_distrib.bat --output-dir=D:\Code\cef_out --ninja-build --x64-build
```

打包完成后可以在 `D:\Code\chromium_git\chromium\src\cef\binary_distrib`下找到打包的文件

### Downloading a file with PowerShell

```bash
Invoke-WebRequest $url -OutFile $path_to_file
```

### 编译 cef_sandbox.lib

编译 cef_sandbox.lib 命令如下（注意后面的参数，不是 cef 了，而是 cef_sandbox，目录也是用的是带有 sandbox 的目录）：

```bash
ninja -C out\Release_GN_x86_sandbox cef_sandbox
```

## 另一种修改配置的方法

修改`D:\Code\chromium_git\chromium\src\third_party\ffmpeg\chromium\scripts\build_ffmpeg.py`

找到`configure_flags['Chrome'].extend`，大概在`Line 1014`

```bash
configure_flags['Chrome'].extend([
 '--enable-decoder=aac,h264,mp3,mpeg4,amrnb,amrwb,flv',
 '--enable-demuxer=aac,mp3,mov,avi,amr,flv',
 '--enable-parser=aac,h264,mpegaudio,mpeg4video,h263',
])

```

修改`D:\Code\chromium_git\chromium\src\third_party\ffmpeg\chromium\config\Chrome\win\ia32\config.h`
这里修改的是32位的目录（ia32)， 如果你是要编译64位的， 也要把x64的一起修改了
把第4行的注释去掉，修改下面这样：

```bash
#define FFMPEG_CONFIGURATION "--disable-everything --disable-all --disable-doc --disable-htmlpages --disable-manpages --disable-podpages --disable-txtpages --disable-static --enable-avcodec --enable-avformat --enable-avutil --enable-fft --enable-rdft --enable-static --enable-libopus --disable-debug --disable-bzlib --disable-error-resilience --disable-iconv --disable-lzo --disable-network --disable-schannel --disable-sdl2 --disable-symver --disable-xlib --disable-zlib --disable-securetransport --disable-faan --disable-alsa --disable-autodetect --enable-decoder='vorbis,libopus,flac' --enable-decoder='pcm_u8,pcm_s16le,pcm_s24le,pcm_s32le,pcm_f32le,mp3' --enable-decoder='pcm_s16be,pcm_s24be,pcm_mulaw,pcm_alaw' --enable-demuxer='ogg,matroska,wav,flac,mp3,mov' --enable-parser='opus,vorbis,flac,mpegaudio,vp9' --extra-cflags=-I/usr/local/google/home/jrummell/chromium/src/third_party/opus/src/include --disable-linux-perf --x86asmexe=nasm --optflags='\"-O2\"' --enable-decoder='theora,vp8' --enable-parser='vp3,vp8' --toolchain=msvc --extra-cflags=-I/usr/local/google/home/jrummell/chromium/src/third_party/ffmpeg/chromium/include/win --enable-cross-compile --cc=clang-cl --ld=lld-link --nm=llvm-nm --ar=llvm-ar --extra-cflags=-O2 --extra-cflags=-m32 --extra-cflags=-imsvc/usr/local/google/home/jrummell/chromium/src/third_party/ffmpeg/../depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/win_sdk/Include/10.0.18362.0/um --extra-cflags=-imsvc/usr/local/google/home/jrummell/chromium/src/third_party/ffmpeg/../depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/win_sdk/Include/10.0.18362.0/shared --extra-cflags=-imsvc/usr/local/google/home/jrummell/chromium/src/third_party/ffmpeg/../depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/win_sdk/Include/10.0.18362.0/winrt --extra-cflags=-imsvc/usr/local/google/home/jrummell/chromium/src/third_party/ffmpeg/../depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/win_sdk/Include/10.0.18362.0/ucrt --extra-cflags=-imsvc/usr/local/google/home/jrummell/chromium/src/third_party/ffmpeg/../depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/VC/Tools/MSVC/14.23.28105/include --extra-cflags=-imsvc/usr/local/google/home/jrummell/chromium/src/third_party/ffmpeg/../depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/VC/Tools/MSVC/14.23.28105/atlmfc/include --extra-ldflags='-libpath:/usr/local/google/home/jrummell/chromium/src/third_party/depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/VC/Tools/MSVC/14.23.28105/atlmfc/lib/x86' --extra-ldflags='-libpath:/usr/local/google/home/jrummell/chromium/src/third_party/depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/win_sdk/Lib/10.0.18362.0/ucrt/x86' --extra-ldflags='-libpath:/usr/local/google/home/jrummell/chromium/src/third_party/depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/win_sdk/Lib/10.0.18362.0/um/x86' --extra-ldflags='-libpath:/usr/local/google/home/jrummell/chromium/src/third_party/depot_tools/win_toolchain/vs_files/9ff60e43ba91947baca460d0ca3b1b980c3a2c23/VC/Tools/MSVC/14.23.28105/lib/x86' --enable-decoder='aac,h264' --enable-demuxer=aac --enable-parser='aac,h264' -enable-decoder=’rv10,rv20,rv30,rv40,cook,h263,h263i,mpeg4,msmpeg4v1,msmpeg4v2,msmpeg4v3,amrnb,amrwb,ac3,flv’ -enable-demuxer=’rm,mpegvideo,avi,avisynth,h263,aac,amr,ac3,flv,mpegts,mpegtsraw’ -enable-parser=’mpegvideo,rv30,rv40,h263,mpeg4video,ac3′"
```

然后修改最后一个文件：
`D:\Code\chromium_git\chromium\src\out\Debug_GN_x86\args.gn`

在文件的末尾加上两行

```bash
proprietary_codecs=true
ffmpeg_branding="Chrome"
```

## bat 脚本常识

注释 `:rem`  或 `::`

在`.bat`脚本中，换行符是`^`,即 `shift + 6`

使用2个%包含一个字符串表示引用环境变量

`cmd1; cmd2; cmd3`
不管cmd1命令是否执行成功，cmd2命令都执行。

这样前后执行的时候没有依赖性，相互独立。

前一个命令执行完成，才会执行后一个命令。

注意：单行语句一般要用到分号来区分代码块，代码若写作多行，用换行符来区分代码块，则无需用到分号。

`cmd1&cmd2&cmd3`

cmd1命令和cmd2命令同时执行

`cmd1 && cmd2`

使用 && 来连接命令，当cmd1执行成功，即cmd1执行成功的时候才执行cmd2，否则不执行cmd2

`cmd1 || cmd2`

使用 || 来连接命令，cmd1执行失败的时候才执行cmd2

## ERROR 处理

### Error1

```bash
Traceback (most recent call last):
File "C:/code/chromium_git/chromium/src/build/vs_toolchain.py", line 489, in <module>
sys.exit(main())
File "C:/code/chromium_git/chromium/src/build/vs_toolchain.py", line 485, in main
return commands[sys.argv[1]](*sys.argv[2:])
File "C:/code/chromium_git/chromium/src/build/vs_toolchain.py", line 313, in CopyDlls
_CopyDebugger(target_dir, target_cpu)
File "C:/code/chromium_git/chromium/src/build/vs_toolchain.py", line 344, in _CopyDebugger
' 10 SDK.' % (debug_file, full_path))
Exception: dbghelp.dll not found in "C:\Program Files (x86)\Windows Kits\10\Debuggers\x86\dbghelp.dll"
You must install the "Debugging Tools for Windows" feature from the Windows 10 SDK.
ERROR at //build/toolchain/win/BUILD.gn:43:3: Script returned non-zero exit code.
exec_script("../../vs_toolchain.py",
^----------
```

查了查原因，发现"C:\Program Files (x86)\Windows Kits\10\"下缺少"Debugger"文件夹，不知道什么缘由。

尝试手动下载Windows 10 SDK(<https://developer.microsoft.com/en-us/windows/downloads/windows-10-sdk>)

进行安装后，再次重复上一步，建立"create.bat"脚本并执行。

### Error2

编译时提示Unicode编码相关的错误

请确认**区域和语言**是否修改成了美国和英文。
