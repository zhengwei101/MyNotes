#!/bin/sh
basepath=$(cd `dirname $0`;pwd)
echo ${basepath}

cd ${basepath}/ffmpeg
pwd

export PKG_CONFIG_PATH=${PKG_CONFIG_PATH}:/d/zWork/FFmpeg/x264-install/lib/pkgconfig
echo ${PKG_CONFIG_PATH}

./configure --prefix=${basepath}/ffmpeg-install \
--enable-gpl --enable-libx264 --disable-static --enable-shared \
--extra-cflags=-l${basepath}/x264-install/include --extra-ldflags=-L${basepath}/x264-install/lib

make -j8
make install