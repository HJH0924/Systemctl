// Package Systemctl
/**
* @Project : Systemctl
* @File    : util_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/7 11:04
**/

package Systemctl

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIsActive(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl is-active --system testservice",
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			name:  "systemctl is-active --user testservice",
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			isActive, err := IsActive(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Println(isActive)
			fmt.Println(err)
		})
	}
}

func TestIsEnabled(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl is-enabled --system testservice",
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			name:  "systemctl is-enabled --user testservice",
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			isEnabled, err := IsEnabled(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Println(isEnabled)
			fmt.Println(err)
		})
	}
}

func TestIsFailed(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl is-failed --system testservice",
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			name:  "systemctl is-failed --user testservice",
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			isFailed, err := IsFailed(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Println(isFailed)
			fmt.Println(err)
		})
	}
}

func TestShow(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl show --system testservice --property [property]",
			unit:  "testservice",
			runAs: ROOT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			for _, property := range properties {
				val, err := Show(systemctl, ctx, tt.unit, property, Options{Mode: tt.runAs})
				fmt.Printf("%s = %s\n", property, val)
				fmt.Println(err)
			}
		})
	}
}

func TestGetStartTime(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl show [unit] --property ExecMainStartTimestamp",
			unit:  "testservice",
			runAs: ROOT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			val, err := GetStartTime(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Println("StartTime: ", val)
			fmt.Println(err)
		})
	}
}

func TestGetUnitRestartCount(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl show [unit] --property NRestarts",
			unit:  "testservice",
			runAs: ROOT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			val, err := GetUnitRestartCount(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Println("UnitRestartCount: ", val)
			fmt.Println(err)
		})
	}
}

func TestGetMemoryUsage(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl show [unit] --property MemoryCurrent",
			unit:  "testservice",
			runAs: ROOT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			val, err := GetMemoryUsage(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Printf("MemoryUsage: %d bytes\n", val)
			fmt.Println(err)
		})
	}
}

func TestGetMainPID(t *testing.T) {
	tests := []struct {
		name  string
		unit  string
		runAs UserMode
	}{
		{
			name:  "systemctl show [unit] --property MainPID",
			unit:  "testservice",
			runAs: ROOT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			val, err := GetMainPID(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Printf("MainPID: %d\n", val)
			fmt.Println(err)
		})
	}
}

func TestGetAllUnits(t *testing.T) {
	tests := []struct {
		name  string
		runAs UserMode
	}{
		{
			name:  "systemctl list-units --all --no-legend --full --no-pager",
			runAs: ROOT,
		},
		{
			name:  "systemctl list-units --all --no-legend --full --no-pager --user",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			units, err := GetAllUnits(systemctl, ctx, Options{Mode: tt.runAs})
			for _, unit := range units {
				unit.Print()
				fmt.Println("---")
			}
			fmt.Println(err)
		})
	}
}

func TestGetAllMaskedUnits(t *testing.T) {
	tests := []struct {
		name  string
		runAs UserMode
	}{
		{
			name:  "systemctl list-unit-files --state=masked",
			runAs: ROOT,
		},
		{
			name:  "systemctl list-unit-files --state=masked --user",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			units, err := GetAllMaskedUnits(systemctl, ctx, Options{Mode: tt.runAs})
			for _, unit := range units {
				fmt.Println(unit)
				fmt.Println("---")
			}
			fmt.Println(err)
		})
	}
}

func TestIsMasked(t *testing.T) {
	tests := []struct {
		unit  string
		runAs UserMode
	}{
		{
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			isMasked, err := IsMasked(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Println(isMasked)
			fmt.Println(err)
		})
	}
}

func TestIsRunning(t *testing.T) {
	tests := []struct {
		unit  string
		runAs UserMode
	}{
		{
			unit:  "testservice",
			runAs: ROOT,
		},
		{
			unit:  "testservice",
			runAs: USER,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if !isRoot(username) && tt.runAs == ROOT {
				t.Skip(UserSkipTest)
			} else if isRoot(username) && tt.runAs == USER {
				t.Skip(RootSkipTest)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			systemctl := NewSystemctl()
			isRunning, err := IsRunning(systemctl, ctx, tt.unit, Options{Mode: tt.runAs})
			fmt.Println(isRunning)
			fmt.Println(err)
		})
	}
}
