// Package Systemctl
/**
* @Project : Systemctl
* @File    : structs.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/9 11:01
**/

package Systemctl

type UserMode bool

const (
	ROOT UserMode = false
	USER UserMode = true
)

type Options struct {
	Mode UserMode
}

type Unit struct {
	Name        string
	Load        string
	Active      string
	Sub         string
	Description string
}
