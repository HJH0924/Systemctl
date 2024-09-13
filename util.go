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
	"strconv"
	"strings"
	"time"
)

// IsActive checks if the specified systemd unit is active.
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

// IsEnabled checks if the specified systemd unit is enabled to start on boot.
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

// IsFailed checks if the specified systemd unit has failed to start.
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

// Show retrieves a specific property of a systemd unit.
func Show(systemctl *Systemctl, ctx context.Context, unit string, property string, opts Options) (string, error) {
	res := systemctl.Show(ctx, unit, property, opts)
	val := res.Output
	val = strings.TrimPrefix(val, property+"=")
	val = strings.TrimSuffix(val, "\n")
	return val, res.Err
}

const dateFormat = "Mon 2006-01-02 15:04:05 MST"

// GetStartTime
// Get start time of a service (`systemctl show [unit] --property ExecMainStartTimestamp`) as a `Time` type
func GetStartTime(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (time.Time, error) {
	val, err := Show(systemctl, ctx, unit, ExecMainStartTimestamp, opts)
	// err not nil
	// or
	// ExecMainStartTimestamp returns an empty string if the unit is not running
	if err != nil || val == "" {
		return time.Time{}, err
	}

	return time.Parse(dateFormat, val)
}

// GetUnitRestartCount
// Get the number of times a unit restarted (`systemctl show [unit] --property NRestarts`) as an int
func GetUnitRestartCount(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (int, error) {
	val, err := Show(systemctl, ctx, unit, NRestarts, opts)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(val)
}

// GetMemoryUsage
// Get current memory in bytes (`systemctl show [unit] --property MemoryCurrent`) as an int
func GetMemoryUsage(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (int, error) {
	val, err := Show(systemctl, ctx, unit, MemoryCurrent, opts)
	if err != nil || val == "[not set]" {
		return -1, err
	}

	return strconv.Atoi(val)
}

// GetMainPID
// Get the PID of the main process (`systemctl show [unit] --property MainPID`) as an int
func GetMainPID(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (int, error) {
	val, err := Show(systemctl, ctx, unit, MainPID, opts)
	if err != nil {
		return -1, err
	}

	return strconv.Atoi(val)
}

// GetAllUnits retrieves a list of all systemd units with their statuses and descriptions.
// It executes the `systemctl list-units` command with various options to get detailed information.
func GetAllUnits(systemctl *Systemctl, ctx context.Context, opts Options) ([]Unit, error) {
	args := []string{"list-units", "--all", "--no-legend", "--full", "--no-pager"}
	if opts.Mode == USER {
		args = append(args, "--user")
	}
	res := systemctl.Execute(ctx, args)
	if res.Err != nil {
		return []Unit{}, res.Err
	}

	lines := strings.Split(res.Output, "\n")
	var units []Unit
	for _, line := range lines {
		entry := strings.Fields(line)
		if len(entry) < 4 {
			continue
		}
		unit := Unit{
			Name:        entry[0],
			Load:        entry[1],
			Active:      entry[2],
			Sub:         entry[3],
			Description: strings.Join(entry[4:], " "),
		}
		units = append(units, unit)
	}
	return units, nil
}

// GetAllMaskedUnits retrieves a list of all masked systemd units.
func GetAllMaskedUnits(systemctl *Systemctl, ctx context.Context, opts Options) ([]string, error) {
	args := []string{"list-unit-files", "--state=masked"}
	if opts.Mode == USER {
		args = append(args, "--user")
	}
	res := systemctl.Execute(ctx, args)
	if res.Err != nil {
		return []string{}, res.Err
	}

	lines := strings.Split(res.Output, "\n")
	var units []string
	for _, line := range lines {
		if !strings.Contains(line, "masked") {
			continue
		}
		entry := strings.Split(line, " ")
		if len(entry) < 3 {
			continue
		}
		if entry[1] == "masked" {
			unit := entry[0]
			uName := strings.Split(unit, ".")
			unit = uName[0]
			units = append(units, unit)
		}
	}
	return units, nil
}

// IsMasked checks if the specified systemd unit is masked.
func IsMasked(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (bool, error) {
	mUnits, err := GetAllMaskedUnits(systemctl, ctx, opts)
	if err != nil {
		return false, err
	}

	for _, mUnit := range mUnits {
		if mUnit == unit {
			return true, nil
		}
	}
	return false, nil
}

// IsRunning checks if the specified systemd service is currently running.
// https://unix.stackexchange.com/a/396633
func IsRunning(systemctl *Systemctl, ctx context.Context, unit string, opts Options) (bool, error) {
	status, err := Show(systemctl, ctx, unit, SubState, opts)
	return status == "running", err
}
