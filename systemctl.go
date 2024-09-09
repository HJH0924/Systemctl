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
func DaemonReload(ctx context.Context, opts Options) error {
	args := []string{"daemon-reload", "--system"}
	if opts.UserMode {
		args[1] = "--user"
	}
	res := execute(ctx, args)
	return res.err
}
