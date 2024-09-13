[![PkgGoDev](https://pkg.go.dev/github.com/HJH0924/Systemctl)](https://pkg.go.dev/github.com/HJH0924/Systemctl)
# Systemctl

这个库旨在为 Go 开发人员提供符合习惯的 `systemctl` 绑定，以便更容易地使用 Go 语言编写系统工具。这个工具试图通过为 `systemctl` 提供结构化、经过彻底测试的包装器，来消除随意使用 shell 执行 `systemctl` 命令的猜测工作。

如果你的系统不是运行（或目标是另一个运行）`systemctl`，这个库对你来说可能没什么用处。



## 什么是 systemctl

`systemctl` 是一个命令行程序，它允许用户控制 systemd 系统和服务管理器。

`systemctl` 可以用来检查和控制 "systemd" 系统和服务管理器的状态。有关这个工具管理的基本概念和功能的介绍，请参阅 [官方 systemctl 文档](https://www.man7.org/linux/man-pages/man1/systemctl.1.html) 。



## 支持的 systemctl 功能

- [x] `systemctl daemon-reload`
- [x] `systemctl disable`
- [x] `systemctl enable`
- [x] `systemctl reenable`
- [x] `systemctl is-active`
- [x] `systemctl is-enabled`
- [x] `systemctl is-failed`
- [ ] `systemctl show`
- [ ] `systemctl start`
- [ ] `systemctl restart`
- [ ] `systemctl stop`
- [ ] `systemctl status`
- [ ] `systemctl mask`
- [ ] `systemctl unmask`



## 辅助功能

- [ ] 获取服务的启动时间（`ExecMainStartTimestamp`）作为 `Time` 类型
- [ ] 获取当前内存（`MemoryCurrent`）作为 int
- [ ] 获取主进程的 PID（`MainPID`）作为 int
- [ ] 获取单元的重启次数（`NRestarts`）作为 int



## 有用的错误

所有函数都返回一个预定义的错误类型，强烈建议正确处理这些错误。



## 上下文支持

这个库的所有调用都支持 Go 的 `context` 功能。
因此，阻塞调用可以根据调用者的需要超时，返回的错误应该检查是否发生了超时（`ErrExecTimeout`）。



## 简单示例

```go
// 示例代码
```



## 运行测试

运行测试用例之前，您需要在您的系统上安装 `testservice.service`

```shell
sudo cp Systemctl/testservice.service /etc/systemd/system
sudo systemctl daemon-reload
```



为了全面测试这个库，您需要以root用户和普通用户身份运行测试，因为`systemctl`可能需要root权限来执行某些操作。以下是一些建议：

1.  **以root用户身份运行测试**：某些测试可能需要root权限，例如启用或禁用服务。

    ```shell
    sudo su root # 切换 root 用户
    /usr/local/go/bin/go test -run TestDaemonReload # 以 root 身份运行 TestDaemonReload 测试用例
    ```

2.  **以普通用户身份运行测试**：其他测试应该以普通用户身份运行，以确保库在没有提升权限的情况下也能正常工作。

    ```shell
    go test -run TestDaemonReload
    ```
