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

func TestDaemonReload(t *testing.T) {
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
			res := DaemonReload(ctx, tt.opts)
			res.Print()
		})
	}
}

/* Run these tests only as a superuser */
func TestDisable(t *testing.T) {
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

			res := Disable(ctx, tt.unit, Options{Mode: ROOT})
			res.Print()
		})
	}
}

/* Run these tests only as a superuser */
func TestEnable(t *testing.T) {
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

			res := Enable(ctx, tt.unit, Options{Mode: ROOT})
			res.Print()
		})
	}
}

/* Run these tests only as a superuser */
func TestReEnable(t *testing.T) {
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

			res := ReEnable(ctx, tt.unit, Options{Mode: ROOT})
			res.Print()
		})
	}
}

// isRoot 检查当前执行测试的用户是否是root用户
func isRoot(username string) bool {
	return username == "root" || username == "system"
}
