# FFmpeg 入门指南

## FFmpeg 命令

### FFmpeg 命令分类

1. 基本信息查询命令
2. 录制命令
3. 分解/复用命令
4. 处理原始数据命令
5. 裁剪与合并命令
6. 图片/视频互转命令
7. 直播相关命令
8. 各种滤镜命令

## FFmpeg 处理音视频流程

输入文件 --demuxer-->  编码数据包 --decoder--> 解码后数据帧 --encoder-->编码数据包--muxer-->输出文件

### 基本信息查询命令

1.显示版本 `-version`
2.显示可用的格式 `-formats`
3.显示可用的demuxers `-demuxers`
4.显示可用的muxers `-muxers`
5.显示可用的设备 `devices`
6.显示可用的解码器 `-decoders`
7.显示所有的编码器 `-encoders`
8.显示所有的编解码器 `-codecs`
9.显示可用的过滤器 `=filters`
10.显示可用的协议 `-protocols`
11.显示可用的像素格式 `-pix_fmts`
12.显示可用的采样格式 `-sample_fmts`
13.显示channel名称 `-layouts`
14.显示识别的颜色名称 `-colors`
15.显示比特流filter `-bsfs`

### 录制命令

```sh
ffmpeg -f avfoundation -i1 -r30 out.yuv
```

`-f`: 指定使用avfoundation来采集数据
`-i`：指定从哪儿采集数据，它是一个文件索引号
`-r`：指定帧率

播放视频文件

```sh
ffplay -s 2560x1600 -pix_fmt uyvy422 out.yuv
```

查看设备列表

```sh
ffmpeg -f avfoundation -list_devices true -i ""
```

### 录制声音命令

```sh
ffmpeg -f avfoundation -i :0 out.wav
```

:0 代表音频设备

如何同时录制图像与声音呢？

```sh
```

### 分解与复用

