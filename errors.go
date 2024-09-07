// Package Systemctl
/**
* @Project : Systemctl
* @File    : errors.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/5 19:24
**/

package Systemctl

import "errors"

var (
	// ErrNotInstalled
	// Make sure systemctl is in the PATH before calling again
	ErrNotInstalled = errors.New("systemctl is not in $PATH")
)
