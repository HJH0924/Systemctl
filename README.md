[![PkgGoDev](https://pkg.go.dev/github.com/HJH0924/Systemctl)](https://pkg.go.dev/github.com/HJH0924/Systemctl)
# Systemctl

这个库旨在为 Go 开发人员提供符合习惯的 `systemctl` 绑定，以便更容易地使用 Go 语言编写系统工具。这个工具试图通过为 `systemctl` 提供结构化、经过彻底测试的包装器，来消除随意使用 shell 执行 `systemctl` 命令的猜测工作。

如果你的系统不是运行（或目标是另一个运行）`systemctl`，这个库对你来说可能没什么用处。



## 什么是 systemctl

`systemctl` 是一个命令行程序，它允许用户控制 systemd 系统和服务管理器。

`systemctl` 可以用来检查和控制 "systemd" 系统和服务管理器的状态。有关这个工具管理的基本概念和功能的介绍，请参阅 [官方 systemctl 文档](https://www.man7.org/linux/man-pages/man1/systemctl.1.html) 。



## 当前支持的 systemctl 功能

- [x] `systemctl daemon-reload`
- [x] `systemctl disable`
- [x] `systemctl enable`
- [x] `systemctl reenable`
- [x] `systemctl is-active`
- [x] `systemctl is-enabled`
- [x] `systemctl is-failed`
- [x] `systemctl show`
- [x] `systemctl start`
- [x] `systemctl status`
- [x] `systemctl stop`
- [x] `systemctl restart`
- [x] `systemctl mask`
- [x] `systemctl unmask`

...



## 当前支持的辅助功能

- [x] **检查服务是否激活** `IsActive`：确定指定的服务是否处于激活状态。
- [x] **检查服务是否启用** `IsEnabled`：检查服务是否设置为在系统启动时自动启动。
- [x] **检查服务是否失败** `IsFailed`：检查服务是否启动失败。
- [x] **获取服务启动时间** `GetStartTime`：获取服务的启动时间。
- [x] **获取服务重启次数** `GetUnitRestartCount`：获取服务重启的次数。
- [x] **获取服务内存使用量** `GetMemoryUsage`：获取服务当前的内存使用量。
- [x] **获取主进程 PID** `GetMainPID`：获取服务主进程的进程 ID。
- [x] **列出所有单元** `GetAllUnits`：获取所有 systemd 单元的列表及其状态和描述。
- [x] **列出所有已屏蔽单元** `GetAllMaskedUnits`：获取所有被屏蔽（不允许启动）的 systemd 单元列表。
- [x] **检查单元是否已屏蔽** `IsMasked`：检查指定单元是否被屏蔽。
- [x] **检查服务是否正在运行** `IsRunning`：检查服务是否正在运行。

...



## 有用的错误

所有函数都返回一个预定义的错误类型，强烈建议正确处理这些错误。




## 上下文支持

这个库的所有调用都支持 Go 的 `context` 功能。



## 安装

你可以通过以下命令来安装 Systemctl Go 库：

```sh
go get github.com/HJH0924/Systemctl
```



## 如何使用

要使用 Systemctl Go 库，请按照以下步骤操作：

1. 导入库到你的 Go 项目中。
```go
import "github.com/HJH0924/Systemctl"
```
2. 使用库中的功能函数来执行系统服务管理任务。

例如，要检查服务是否激活，你可以使用 `IsActive` 函数：

```go
package main

import (
	"context"
	"fmt"

	"github.com/HJH0924/Systemctl"
)

func main() {
	ctx := context.Background()
	systemctl := Systemctl.NewSystemctl()
	res := systemctl.Start(ctx, "testservice", Systemctl.Options{Mode: Systemctl.USER})
	active, err := Systemctl.IsActive(systemctl, ctx, "testservice", Systemctl.Options{Mode: Systemctl.USER})
	if err != nil {
		fmt.Println("检查服务状态出错:", err)
		return
	}
	fmt.Printf("testservice 是否处于活动状态？%v\n", active)
}
```



## 运行测试

运行测试用例之前，您需要在您的系统上安装 `testservice.service`

```shell
# 以 root 用户身份运行测试
sudo make install-root
# 以普通用户身份运行测试
sudo make install-user
```


为了全面测试这个库，您需要以root用户和普通用户身份运行测试，因为`systemctl`可能需要root权限来执行某些操作。以下是一些建议：

1.  **以root用户身份运行测试**：某些测试可能需要root权限，例如启用或禁用服务。

    ```shell
    sudo su root # 切换 root 用户
    go test -run TestDaemonReload # 以 root 身份运行 TestDaemonReload 测试用例
    ```

2.  **以普通用户身份运行测试**：其他测试应该以普通用户身份运行，以确保库在没有提升权限的情况下也能正常工作。

    ```shell
    go test -run TestDaemonReload
    ```



## 贡献

欢迎任何形式的贡献！如果你有任何问题、建议或想要提交代码，请随时创建 issue 或提交 pull request。
