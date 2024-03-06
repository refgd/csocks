# CSOCKS
魔改SOCKS5 over TLS

## 简介
目前已知的代理都是需要服务端与客户端保持长连接, 因为各种原因, 我需要一个短连接代理, 当然会牺牲点性能.

## 使用方法

### 生成证书
```bash
openssl req -x509 -nodes -newkey ec:<(openssl ecparam -name prime256v1) -keyout server.key -out server.crt -days 3650
```

### 服务器
```bash
chmod +x csocks
./csocks -l 7000 -k server.crt server.key --http --quiet --secret 123456
```

开启服务端后,会自动生成一个`public.key`文件, 这个文件需要和客户端放在一起, `secret`是授权码, 来验证客户端连接

### 客户端
```bash
.\csocks.exe -l 1080 -s server_ip:7000 --quiet --secret 123456
```

需要将服务器上的`public.key`和`csocks`程序放在同一个目录, 因为是自签名证书, 客户端会通过验证`public.key`确保传输的安全性, `secret`是授权码需要和服务器保持一致

### 配置
按照以上启动后, 在客户端设置软件使用socks5/http代理  
代理地址: 127.0.0.1  
代理端口: 1080  

### 致谢
[4dnat](https://github.com/dushixiang/4dnat)