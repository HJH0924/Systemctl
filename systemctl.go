// Package Systemctl
/**
* @Project : Systemctl
* @File    : systemctl.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/9 11:00
**/

package Systemctl

import "context"

// DaemonReload
// Reload the systemd manager configuration.
//
// This will rerun all generators (see systemd.generator(7)), reload all unit files,
// and recreate the entire dependency tree. While the daemon is being reloaded,
// all sockets systemd listens on behalf of user configuration will stay accessible.
func (Self *Systemctl) DaemonReload(ctx context.Context, opts Options) Result {
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
func (Self *Systemctl) Disable(ctx context.Context, unit string, opts Options) Result {
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
func (Self *Systemctl) Enable(ctx context.Context, unit string, opts Options) Result {
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
func (Self *Systemctl) ReEnable(ctx context.Context, unit string, opts Options) Result {
	args := []string{"reenable", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}

func (Self *Systemctl) IsActive(ctx context.Context, unit string, opts Options) Result {
	args := []string{"is-active", "--system", unit}
	if opts.Mode == USER {
		args[1] = "--user"
	}
	return Self.Execute(ctx, args)
}
