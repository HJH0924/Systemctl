// Package Systemctl
/**
* @Project : Systemctl
* @File    : util.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/5 18:51
**/

package Systemctl

import (
	"context"
	"strings"
)

func IsActive(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (bool, error) {
	res := systemctl.IsActive(ctx, unit, opts)
	status := strings.TrimSuffix(res.Output, "\n")
	switch status {
	case "active":
		return true, nil
	case "inactive", "failed", "activating":
		return false, nil
	default:
		return false, res.Err
	}
}

func IsEnabled(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (bool, error) {
	res := systemctl.IsActive(ctx, unit, opts)
	status := strings.TrimSuffix(res.Output, "\n")
	switch status {
	case "enabled", "enabled-runtime", "alias", "static", "indirect", "generated", "transient":
		return true, nil
	case "disabled":
		return false, nil
	default:
		// "linked", "linked-runtime", "masked", "masked-runtime", etc
		return false, res.Err
	}
}
