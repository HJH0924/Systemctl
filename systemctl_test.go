// Package Systemctl
/**
* @Project : Systemctl
* @File    : systemctl_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/9 14:03
**/

package Systemctl

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"testing"
	"time"
)

var debug = true // for debug, printf stdout, stderr

func Test_execute(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		args []string
	}{
		{
			name: "systemctl command success",
			ctx:  context.Background(),
			args: []string{"status", "ssh"},
		},
		{
			name: "Unit qwe.service could not be found.",
			ctx:  context.Background(),
			args: []string{"status", "qwe"},
		},
		{
			name: "systemctl daemon-reload --user",
			ctx:  context.Background(),
			args: []string{"daemon-reload", "--user"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			systemctl := NewSystemctl()
			res := systemctl.Execute(tt.ctx, tt.args)
			if debug {
				fmt.Printf("stdout: %s\n", res.Output)
				fmt.Printf("stderr: %s\n", res.Warnings)
				fmt.Printf("code: %d\n", res.Code)
				fmt.Printf("error: %s\n", res.Err)
			}
		})
	}
}

// 记录当前执行测试用户的用户名
var username string

// Testing assumptions
// - there's no unit installed named `nonexistant`
// - the `testservice` unit to be available on the tester's system.
//   this is just what was available on mine, should you want to change it,
//   either to something in this repo or more common, feel free to submit a PR.
// - your 'user' isn't root
// - your user doesn't have a PolKit rule allowing access to configure nginx

func TestMain(m *testing.M) {
	curUser, err := user.Current()
	if err != nil {
		fmt.Println("Could not determine running user")
	}

	if curUser != nil {
		username = curUser.Username
	}
	fmt.Printf("Currently running tests as: %s\n", username)
	fmt.Println("Don't forget to run both root and user tests.")
	retCode := m.Run()
	os.Exit(retCode)
}

const (
	RootSkipTest = "skipping user test while running as superuser"
	UserSkipTest = "skipping superuser test while running as user"
)

func TestSystemctl_DaemonReload(t *testing.T) {
	tests := []struct {
		name  string
		opts  Options
		runAs UserMode
	}{
		/* Run these tests only as a user */
		{
			name:  "fail to reload system daemon as user",
			opts:  Options{Mode: ROOT},
			runAs: USER,
		},
		{
			name:  "reload user's scope daemon",
			opts:  Options{Mode: USER},
			runAs: USER,
		},
		/* End user tests */

		/* Run these tests only as a superuser */
		{
			name:  "succeed to reload daemon",
			opts:  Options{Mode: ROOT},
			runAs: ROOT,
		},
		{
			name:  "fail to connect to user bus as system",
			opts:  Options{Mode: USER},
			runAs: ROOT,
		},
		/* End superuser tests */
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			} else if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			res := NewSystemctl().DaemonReload(ctx, tt.opts)
			res.Print()
		})
	}
}

/* Run these tests only as a superuser */
func TestSystemctl_Disable(t *testing.T) {
	tests := []struct {
		name string
		unit string
	}{
		{
			name: "systemctl disable --system testservice",
			unit: "testservice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) {
				t.Skip(UserSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res := NewSystemctl().Disable(ctx, tt.unit, Options{Mode: ROOT})
			res.Print()
		})
	}
}

/* Run these tests only as a superuser */
func TestSystemctl_Enable(t *testing.T) {
	tests := []struct {
		name string
		unit string
	}{
		{
			name: "systemctl enable --system testservice",
			unit: "testservice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) {
				t.Skip(UserSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res := NewSystemctl().Enable(ctx, tt.unit, Options{Mode: ROOT})
			res.Print()
		})
	}
}

/* Run these tests only as a superuser */
func TestSystemctl_ReEnable(t *testing.T) {
	tests := []struct {
		name string
		unit string
	}{
		{
			name: "systemctl reenable --system testservice",
			unit: "testservice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) {
				t.Skip(UserSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res := NewSystemctl().ReEnable(ctx, tt.unit, Options{Mode: ROOT})
			res.Print()
		})
	}
}

func TestSystemctl_IsActive(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl is-active --system testservice",
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			name:  "systemctl is-active --user testservice",
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res := NewSystemctl().IsActive(ctx, tt.unit, Options{Mode: tt.runAs})
			res.Print()
		})
	}
}

func TestSystemctl_IsEnabled(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl is-enabled --system testservice",
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			name:  "systemctl is-enabled --user testservice",
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res := NewSystemctl().IsEnabled(ctx, tt.unit, Options{Mode: tt.runAs})
			res.Print()
		})
	}
}

func TestSystemctl_IsFailed(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl is-failed --system testservice",
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			name:  "systemctl is-failed --user testservice",
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			res := NewSystemctl().IsFailed(ctx, tt.unit, Options{Mode: tt.runAs})
			res.Print()
		})
	}
}

func TestSystemctl_Show(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl show --system testservice --property xxx",
			unit:  "testservice",
			runAs: ROOT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			for _, property := range properties {
				res := NewSystemctl().Show(ctx, tt.unit, property, Options{Mode: tt.runAs})
				res.Print()
			}
		})
	}
}

// isRoot 检查当前执行测试的用户是否是root用户
func isRoot(username string) bool {
	return username == "root" || username == "system"
}
