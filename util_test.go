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
			name:  "systemctl show --system testservice --property xxx",
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
