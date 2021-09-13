# USword

## 安装

`go install github.com/no-src/usword/...@latest`

## 客户端组件

### 1、once

单独运行某个命令

用法：`usword [cmd] 或 usword client [cmd] mode=offline`

### 2、offline

单机运行,持续接收并执行输入的命令

用法：`usword client mode=offline -i`

### 3、online-single

联机模式，客户端与服务器相连接，客户段端发送命令给服务器，要求其执行命令，服务端详见服务端组件server模式

用法：`usword client server=127.0.0.1 port=8989 -i 或 usword client mode=online server=127.0.0.1 port=8989 -i`

### 4、online-multi

联机模式，客户端与多台服务器相连接，客户端将命令同时发送给多个服务端，要求其执行命令，服务端详见服务端组件server模式

用法：`usword multiclient servers=127.0.0.1:8092?mode=online,127.0.0.1:9091,mode=offline -i`

### 5、proxy_client

代理模式客户端，客户端通过与代理服务器相连接，将指令发送给代理服务器，代理服务器会将指令转发给最终的真实服务器，最终服务器会执行相应指令

服务端详见服务端组件server模式

代理服务端详见服务端组件proxy_server模式

用法：`proxy_client proxy_server=127.0.0.1 proxy_port=8001 server=127.0.0.1 port=8002 exec="help help"`

## 服务端组件

### 1、server

服务器模式，监听端口等待客户端连接，连接成功后执行客户端发送的命令

用法：`usword server host=127.0.0.1 port=8989`

### 2、proxy_server

代理模式，监听端口等待客户端连接，与客户端连接成功后，同时连接到指令运行服务器，对客户端发送的指令进行实时转发，其自身不执行任何客户端的命令

用法：`usword proxy_server proxy_server=127.0.0.1 proxy_port=8001`