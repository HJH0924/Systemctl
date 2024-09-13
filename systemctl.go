// Package Systemctl
/**
* @Project : Systemctl
* @File    : systemctl.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/9 11:00
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

type SystemctlResult struct {
	Output   string
	Warnings string
	Code     int
	Err      error
}

func (Self *SystemctlResult) Print() {
	fmt.Printf("Output: %s\n", Self.Output)
	fmt.Printf("Warnings: %s\n", Self.Warnings)
	fmt.Printf("Code: %d\n", Self.Code)
	fmt.Printf("Err: %v\n", Self.Err)
}

func (Self *Systemctl) Execute(ctx context.Context, args []string) SystemctlResult {
	var (
		stdout   bytes.Buffer
		stderr   bytes.Buffer
		output   string
		warnings string
		code     int
		err      error
	)

	if Self.systemctl == "" {
		return SystemctlResult{
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

	return SystemctlResult{
		Output:   output,
		Warnings: warnings,
		Code:     code,
		Err:      err,
	}
}

// DaemonReload
// Reload the systemd manager configuration.
//
// This will rerun all generators (see systemd.generator(7)), reload all unit files,
// and recreate the entire dependency tree. While the daemon is being reloaded,
// all sockets systemd listens on behalf of user configuration will stay accessible.
func (Self *Systemctl) DaemonReload(ctx context.Context, opts Options) SystemctlResult {
	args := []string{"daemon-reload", "--system"}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

// Disable
// Disables one or more units.
//
// This removes all symlinks to the unit files backing the specified units from
// the unit configuration directory, and hence undoes any changes made by
// enable or link.
func (Self *Systemctl) Disable(ctx context.Context, unit string, opts Options) SystemctlResult {
	args := []string{"disable", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

// Enable one or more units or unit instances.
//
// This will create a set of symlinks, as encoded in the [Install] sections of
// the indicated unit files. After the symlinks have been created, the system
// manager configuration is reloaded (in a way equivalent to daemon-reload),
// in order to ensure the changes are taken into account immediately.
func (Self *Systemctl) Enable(ctx context.Context, unit string, opts Options) SystemctlResult {
	args := []string{"enable", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

// ReEnable
// Reenables one or more units.
//
// This removes all symlinks to the unit files backing the specified units from
// the unit configuration directory, then recreates the symlink to the unit again,
// atomically. Can be used to change the symlink target.
func (Self *Systemctl) ReEnable(ctx context.Context, unit string, opts Options) SystemctlResult {
	args := []string{"reenable", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

// IsActive
// Check whether any of the specified units are active (i.e. running).
//
// Returns true if the unit is active, false if inactive or failed.
// Also returns false in an error case.
func (Self *Systemctl) IsActive(ctx context.Context, unit string, opts Options) SystemctlResult {
	args := []string{"is-active", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

// IsEnabled
// Checks whether any of the specified unit files are enabled (as with enable).
//
// Returns true if the unit is enabled, aliased, static, indirect, generated
// or transient.
//
// Returns false if disabled. Also returns an error if linked, masked, or bad.
//
// See https://www.freedesktop.org/software/systemd/man/systemctl.html#is-enabled%20UNIT%E2%80%A6
// for more information
func (Self *Systemctl) IsEnabled(ctx context.Context, unit string, opts Options) SystemctlResult {
	args := []string{"is-enabled", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

// IsFailed
// Check whether any of the specified units are in a "failed" state.
func (Self *Systemctl) IsFailed(ctx context.Context, unit string, opts Options) SystemctlResult {
	args := []string{"is-failed", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

// Show a selected property of a unit. Accepted properties are predefined in the
// properties subpackage to guarantee properties are valid and assist code-completion.
func (Self *Systemctl) Show(ctx context.Context, unit string, property string, opts Options) SystemctlResult {
	args := []string{"show", "--system", unit, "--property", property}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}
