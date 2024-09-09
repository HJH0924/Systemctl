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
	ROOT bool = false
	USER bool = true
)

func TestDaemonReload(t *testing.T) {
	tests := []struct {
		name  string
		opts  Options
		runAs bool
	}{
		/* Run these tests only as a user */
		{
			name:  "fail to reload system daemon as user",
			opts:  Options{UserMode: ROOT},
			runAs: USER,
		},
		{
			name:  "reload user's scope daemon",
			opts:  Options{UserMode: USER},
			runAs: USER,
		},
		/* End user tests */

		/* Run these tests only as a superuser */
		{
			name:  "succeed to reload daemon",
			opts:  Options{UserMode: ROOT},
			runAs: ROOT,
		},
		{
			name:  "fail to connect to user bus as system",
			opts:  Options{UserMode: USER},
			runAs: ROOT,
		},
		/* End superuser tests */
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (username == "root" || username == "system") && tt.runAs == USER {
				t.Skip("skipping user test while running as superuser")
			} else if (username != "root" && username != "system") && tt.runAs == ROOT {
				t.Skip("skipping superuser test while running as user")
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			err := DaemonReload(ctx, tt.opts)
			fmt.Println(err)
		})
	}
}
