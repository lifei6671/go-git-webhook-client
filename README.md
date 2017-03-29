# go-git-webhook-client

[![Build Status](https://travis-ci.org/lifei6671/go-git-webhook-client.svg?branch=master)](https://travis-ci.org/lifei6671/go-git-webhook-client)
[![Build status](https://ci.appveyor.com/api/projects/status/m618nm2i9tf0lw10/branch/master?svg=true)](https://ci.appveyor.com/project/lifei6671/go-git-webhook-client/branch/master)


该客户端不但可以配置[SmartWebHook](https://github.com/lifei6671/go-git-webhook) 使用，还可以作为单机版的 WebHook 客户端使用。

# 配置文件

```ini
#程序使用的节点
[system]
#监听的端口号
httpport = 8081


#远程 SmartWebHook 服务连接使用的账号和密码
account = root
password = 123456

#单击版 WebHook 节点，该节点为自定义用于识别不同项目的标识
[e59ef60482fd71b0e6bf60da8a8a40a5]
#项目名称
repo_name = backshop
#分支名称
branch_name = master
#需要执行的命令
command = E:/backshop.bat
#自定义日志目录
log_path = E:/smartwiki.log
```

# 编译

**拉去源码**

```bash
git clone https://github.com/lifei6671/go-git-webhook-client.git
```

**编译程序**

```bash
#更新依赖
go get -d ./...

#编译项目
go build -v -tags "pam" -ldflags "-w"
```

**运行**

```bash
chmod 0777 go-git-webhook-client

```

**使用 supervisor 管理程序**

```ini

[program:go-git-webhook-client]
command=/opt/go/src/github.com/lifei6671/go-git-webhook-client/go-git-webhook-client
autostart=true
autorestart=true
startsecs=10
stdout_logfile=/var/log/go-git-webhook-client/access.log
stdout_logfile_maxbytes=1MB
stdout_logfile_backups=10
stdout_capture_maxbytes=1MB
stderr_logfile=/var/log/go-git-webhook-client/error.log
stderr_logfile_maxbytes=1MB
stderr_logfile_backups=10
stderr_capture_maxbytes=1MB
```

请将配置中的 `command` 配置为你服务器的实际程序地址

# 反馈

如果使用中出现问题,请在 [issues](https://github.com/lifei6671/go-git-webhook-client/issues) 中反馈。

