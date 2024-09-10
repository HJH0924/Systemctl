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

type Systemctl struct {
	systemctl string
}

func NewSystemctl() *Systemctl {
	systemctl, _ := exec.LookPath("systemctl")
	return &Systemctl{
		systemctl: systemctl,
	}
}

type Result struct {
	Output   string
	Warnings string
	Code     int
	Err      error
}

func (Self *Result) Print() {
	fmt.Printf("Output: %s\n", Self.Output)
	fmt.Printf("Warnings: %s\n", Self.Warnings)
	fmt.Printf("Code: %d\n", Self.Code)
	fmt.Printf("Err: %v\n", Self.Err)
}

func (Self *Systemctl) Execute(ctx context.Context, args []string) Result {
	var (
		stdout   bytes.Buffer
		stderr   bytes.Buffer
		output   string
		warnings string
		code     int
		err      error
	)

	if Self.systemctl == "" {
		return Result{
			Code: 1,
			Err:  ErrNotInstalled,
		}
	}

	cmd := exec.CommandContext(ctx, Self.systemctl, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	_ = cmd.Run()
	output = stdout.String()
	warnings = stderr.String()
	code = cmd.ProcessState.ExitCode()

	if code != 0 {
		err = fmt.Errorf("received error code %d for stderr `%s`", code, strings.TrimRight(warnings, "\n"))
	}

	return Result{
		Output:   output,
		Warnings: warnings,
		Code:     code,
		Err:      err,
	}
}
