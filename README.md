# Systemctl

这个库旨在为 Go 开发人员提供符合习惯的 `systemctl` 绑定，以便更容易地使用 Go 语言编写系统工具。这个工具试图通过为 `systemctl` 提供结构化、经过彻底测试的包装器，来消除随意使用 shell 执行 `systemctl` 命令的猜测工作。

如果你的系统不是运行（或目标是另一个运行）`systemctl`，这个库对你来说可能没什么用处。



## 什么是 systemctl

`systemctl` 是一个命令行程序，它允许用户控制 systemd 系统和服务管理器。

`systemctl` 可以用来检查和控制 "systemd" 系统和服务管理器的状态。有关这个工具管理的基本概念和功能的介绍，请参阅 [官方 systemctl 文档](https://www.man7.org/linux/man-pages/man1/systemctl.1.html) 。



## 支持的 systemctl 功能

- [ ] `systemctl daemon-reload`
- [ ] `systemctl disable`
- [ ] `systemctl enable`
- [ ] `systemctl reenable`
- [ ] `systemctl is-active`
- [ ] `systemctl is-enabled`
- [ ] `systemctl is-failed`
- [ ] `systemctl mask`
- [ ] `systemctl restart`
- [ ] `systemctl show`
- [ ] `systemctl start`
- [ ] `systemctl status`
- [ ] `systemctl stop`
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



