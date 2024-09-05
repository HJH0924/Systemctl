# Systemctl

This library aims at providing idiomatic `systemctl` bindings for go developers, in order to make it easier to write system tooling using golang.
This tool tries to take guesswork out of arbitrarily shelling out to `systemctl` by providing a structured, thoroughly-tested wrapper for the `systemctl` functions most-likely to be used in a system program.

If your system isn't running (or targeting another system running) `systemctl`, this library will be of little use to you.



## What is systemctl

`systemctl` is a command-line program which grants the user control over the systemd system and service manager.

`systemctl` may be used to introspect and control the state of the "systemd" system and service manager. Please refer to [Official systemctl documentation](https://www.man7.org/linux/man-pages/man1/systemctl.1.html) for an introduction into the basic concepts and functionality this tool manages.



## Supported systemctl functions

- [ ] `systemctl daemon-reload`
- [ ] `systemctl disable`
- [ ] `systemctl enable`
- [ ] `systemctl reenable`
- [ ] `systemctl is-active`
- [ ] `systemctl is-enabled`
- [ ] `systemctl is-failed`
- [ ] `systemctl mask`
- [ ] `systemctl restart`
- [ ] `systemctl show`
- [ ] `systemctl start`
- [ ] `systemctl status`
- [ ] `systemctl stop`
- [ ] `systemctl unmask`



## Helper functionality

- [ ] Get start time of a service (`ExecMainStartTimestamp`) as a `Time` type
- [ ] Get current memory in bytes (`MemoryCurrent`) an an int
- [ ] Get the PID of the main process (`MainPID`) as an int
- [ ] Get the restart count of a unit (`NRestarts`) as an int



## Useful errors

All functions return a predefined error type, and it is highly recommended these errors are handled properly.



## Context support

All calls into this library support go's `context` functionality.
Therefore, blocking calls can time out according to the caller's needs, and the returned error should be checked to see if a timeout occurred (`ErrExecTimeout`).



## Simple example

```go
// code sample
```


