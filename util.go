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
	"regexp"
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
	res := systemctl.IsEnabled(ctx, unit, opts)
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

func IsFailed(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (bool, error) {
	res := systemctl.IsFailed(ctx, unit, opts)

	if matched, _ := regexp.MatchString(`inactive`, res.Output); matched {
		return false, nil
	} else if matched, _ := regexp.MatchString(`active`, res.Output); matched {
		return false, nil
	} else if matched, _ := regexp.MatchString(`failed`, res.Output); matched {
		return true, nil
	} else {
		return false, res.Err
	}
}

func Show(systemctl *Systemctl, ctx context.Context, unit string, property string, opts Options) (string, error) {
	res := systemctl.Show(ctx, unit, property, opts)
	val := res.Output
	val = strings.TrimPrefix(val, property+"=")
	val = strings.TrimSuffix(val, "\n")
	return val, res.Err
}
