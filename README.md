# sshTool
简单的SSH命令执行工具，支持密码登录、私钥登录、免密登录。支持多台机器并发执行

## 参数定义
```cgo
  -cmds string
        Command that needs to be executed
  -ips string
        IP address list, split by ","
  -k string
        SSH private key
  -pwd string
        Password
  -p int
        Port (default 22)
  -u string
        Username
```
## 编译
``go build``

## 用法

*密码登录：*

``./sshTool -ips 127.0.0.1,192.132.133.43 -p 22 -u root -pwd 123456 -cmds "mkdir fantasticKe;ls"``

*秘钥登录：*

``./sshTool -ips 127.0.0.1,192.132.133.43 -p 22 -u root -key /root/.ssh/id_rsa -cmds "mkdir fantasticKe;ls"``

*免密登录：*

``./sshTool -ips 127.0.0.1,192.132.133.43 -p 22 -u root -cmds "mkdir fantasticKe;ls"``
