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
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

var systemctl string

func init() {
	systemctl, _ = exec.LookPath("systemctl")
}

type result struct {
	output   string
	warnings string
	code     int
	err      error
}

func execute(ctx context.Context, args []string) result {
	var (
		stdout   bytes.Buffer
		stderr   bytes.Buffer
		output   string
		warnings string
		code     int
		err      error
	)

	if systemctl == "" {
		return result{
			code: 1,
			err:  ErrNotInstalled,
		}
	}

	cmd := exec.CommandContext(ctx, systemctl, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	_ = cmd.Run()
	output = stdout.String()
	warnings = stderr.String()
	code = cmd.ProcessState.ExitCode()

	if code != 0 {
		err = fmt.Errorf("received error code %d for stderr `%s`", code, strings.TrimRight(warnings, "\n"))
	}

	return result{
		output:   output,
		warnings: warnings,
		code:     code,
		err:      err,
	}
}
