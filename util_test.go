// Package Systemctl
/**
* @Project : Systemctl
* @File    : util_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/7 11:04
**/

package Systemctl

import (
	"context"
	"fmt"
	"testing"
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
