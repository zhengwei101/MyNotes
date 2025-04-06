# 授权方案

1. 获取硬件信息
2. 根据硬件信息生成唯一序列号
3. 根据序列号和用户名（用户邮箱地址）生成license
license 包括唯一序列号，邮箱地址，时间戳（时间戳可以记录授权码生成的具体时间），有效期。
4.多层加密（RSA+AES), 可以指定每个版本的YES码都不一样。

> 加密算法是生成软件授权码的核心技术。常用的加密算法包括对称加密和非对称加密。对称加密算法如AES（高级加密标准），它使用同一个密钥进行加密和解密。
> 非对称加密算法如RSA（Rivest-Shamir-Adleman），它使用不同的密钥进行加密和解密。对称加密算法通常更快，但非对称加密算法提供了更高的安全性。
> 在生成授权码时，开发者通常会选择一种适合其特定需求的加密算法，并结合其他安全措施以确保授权码的不可破解性。

密钥生成器是一种生成密钥的工具，这些密钥用于加密和解密数据。
在软件授权码的生成过程中，密钥生成器通常会生成一对密钥：一个公钥和一个私钥。
公钥用于加密授权码，而私钥用于解密授权码。这种方法确保了授权码的安全性，因为只有拥有私钥的人才能解密授权码并验证其合法性。密钥生成器通常会使用高随机性的算法，以确保生成的密钥难以预测和复制。

## 原理

1. 生成密钥对，包含私钥和公钥
2. 授权者保留私钥，使用私钥对授权信息诸如使用截止日期，MAC地址等内容生成license签名证书。
3. 公钥给使用者，放在代码中，用于验证license签名证书是否符合使用条件。

## License证书：

根据客户服务器硬件信息（MAC地址、IP地址、CPU序列号、主板序列号）生成授权证书，同时可以给授权证书设置生效时间与失效时间。

cadtest.lic

设备指纹的原理是这样的：

收集若干标识符，比如 CPU 序列号、磁盘序列号、主板序列号、Mac 地址、显卡、内存、系统安装时间、内存大小、已安装字体、系统版本等等，加密后上报服务端，服务端生成一个唯一标识符下发。

后续当再次启动时，重新发送这数据，如果其中只有某几项发生变化，那仍然可以匹配到原有设备（也就是召回）。

1. 先获取硬件信息

```json
{
"ipAddress": [ //授权的ip列表
  "172.17.0.8"
],
"macAddress": [ //授权的mac地址列表
  "52-54-00-74-0B-D9"
],
"cpuSerial": "55 06 05 00 FF FB 8B 0F", //cpu序列号
"mainBoardSerial": "afb14aac-eccb-4a37-9c31-e7951ce73e0d"//主板序列号
}

```

入参传入授权信息：

```json
{

"expiryTime":"2021-05-25 19:07:59", //生效时间
"issuedTime":"2021-04-25 19:07:00", //失效时间
"keyPass":"12345678A", //密钥的密码
"privateAlias":"SYSHLANG",
"licensePath":"/Users/sixj/Desktop/license/license.lic", //证书生成地址
"privateKeysStorePath":"/Users/sixj/Desktop/license/privateKeys.keystore",// 密钥文件地址
"storePass":"12345678A", //密钥库的密码
"subject":"pushi-kn-graph",
"licenseCheckModel":{ //授权验证信息
  "cpuSerial":"47A8E193-23D4-5B93-92AB-4A96FBC0346F",//cpu序列号
  "ipAddress":[ //ip
      "192.168.174.107"
  ],
  "macAddress":[ //mac地址
      "F8-FF-C2-6A-3E-73"
  ],
  "mainBoardSerial":"C02C31HZMD6P"//主板序列号
}
}
会生成一个license.lic授权文件到执行目录
```

授权验证（license-verify-starter）

```json
license.subject: pushi-kn-graph
license.publicAlias: SYSHLANG
license.storePass: 12345678A
license.licensePath: /Users/sixj/Desktop/license/license.lic
license.publicKeysStorePath: /Users/sixj/Desktop/license/publicCerts.keystore
```

项目启动的时候，会去验证授权证书的有效性，是否在有效期内，硬件信息是否匹配，如果授权证书无效，项目启动失败。
另外可以在一些核心接口，比如登陆接口，添加@License注解，请求该接口的时候，也会去验证授权证书的有效性，比如验证证书是否到期，如果失效，该接口将会拒绝访问。

生成一个唯一序列号：
3ZGRUS-4LYXRU-7V8MF6-GC75HM

## 参考

<https://github.com/sixj0/license>
<https://github.com/postbird/Postbird_License>
<https://github.com/denisbrodbeck/machineid>
<https://juejin.cn/post/6864739866979434503>
<https://blog.csdn.net/laoyang360/article/details/50966707>
<https://www.cnblogs.com/lidabo/p/17407712.html>
<https://www.cnblogs.com/chenshikun/p/16860067.html>
<https://github.com/postbird/Postbird_License>
<https://blog.csdn.net/daocaokafei/article/details/136243120>
<https://www.jiandaoyun.com/blog/article/436449/>
