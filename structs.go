// Package Systemctl
/**
* @Project : Systemctl
* @File    : structs.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/9 11:01
**/

package Systemctl

type Options struct {
	UserMode bool
}

type Unit struct {
	Name        string
	Load        string
	Active      string
	Sub         string
	Description string
}
