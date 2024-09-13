// Package Systemctl
/**
* @Project : Systemctl
* @File    : structs.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/9/9 11:01
**/

package Systemctl

import "fmt"

type UserMode bool

const (
	ROOT UserMode = false
	USER UserMode = true
)

type Options struct {
	Mode UserMode
}

type Unit struct {
	Name        string
	Load        string
	Active      string
	Sub         string
	Description string
}

func (Self *Unit) Print() {
	fmt.Println("Name: ", Self.Name)
	fmt.Println("Load: ", Self.Load)
	fmt.Println("Active: ", Self.Active)
	fmt.Println("Sub: ", Self.Sub)
	fmt.Println("Description: ", Self.Description)
}

const (
	ActiveEnterTimestamp            string = "ActiveEnterTimestamp"
	ActiveEnterTimestampMonotonic   string = "ActiveEnterTimestampMonotonic"
	ActiveExitTimestampMonotonic    string = "ActiveExitTimestampMonotonic"
	ActiveState                     string = "ActiveState"
	After                           string = "After"
	AllowIsolate                    string = "AllowIsolate"
	AssertResult                    string = "AssertResult"
	AssertTimestamp                 string = "AssertTimestamp"
	AssertTimestampMonotonic        string = "AssertTimestampMonotonic"
	Before                          string = "Before"
	BlockIOAccounting               string = "BlockIOAccounting"
	BlockIOWeight                   string = "BlockIOWeight"
	CPUAccounting                   string = "CPUAccounting"
	CPUAffinityFromNUMA             string = "CPUAffinityFromNUMA"
	CPUQuotaPerSecUSec              string = "CPUQuotaPerSecUSec"
	CPUQuotaPeriodUSec              string = "CPUQuotaPeriodUSec"
	CPUSchedulingPolicy             string = "CPUSchedulingPolicy"
	CPUSchedulingPriority           string = "CPUSchedulingPriority"
	CPUSchedulingResetOnFork        string = "CPUSchedulingResetOnFork"
	CPUShares                       string = "CPUShares"
	CPUUsageNSec                    string = "CPUUsageNSec"
	CPUWeight                       string = "CPUWeight"
	CacheDirectoryMode              string = "CacheDirectoryMode"
	CanFreeze                       string = "CanFreeze"
	CanIsolate                      string = "CanIsolate"
	CanReload                       string = "CanReload"
	CanStart                        string = "CanStart"
	CanStop                         string = "CanStop"
	CapabilityBoundingSet           string = "CapabilityBoundingSet"
	CleanResult                     string = "CleanResult"
	CollectMode                     string = "CollectMode"
	ConditionResult                 string = "ConditionResult"
	ConditionTimestamp              string = "ConditionTimestamp"
	ConditionTimestampMonotonic     string = "ConditionTimestampMonotonic"
	ConfigurationDirectoryMode      string = "ConfigurationDirectoryMode"
	Conflicts                       string = "Conflicts"
	ControlGroup                    string = "ControlGroup"
	ControlPID                      string = "ControlPID"
	CoredumpFilter                  string = "CoredumpFilter"
	DefaultDependencies             string = "DefaultDependencies"
	DefaultMemoryLow                string = "DefaultMemoryLow"
	DefaultMemoryMin                string = "DefaultMemoryMin"
	Delegate                        string = "Delegate"
	Description                     string = "Description"
	DevicePolicy                    string = "DevicePolicy"
	DynamicUser                     string = "DynamicUser"
	EffectiveCPUs                   string = "EffectiveCPUs"
	EffectiveMemoryNodes            string = "EffectiveMemoryNodes"
	ExecMainCode                    string = "ExecMainCode"
	ExecMainExitTimestampMonotonic  string = "ExecMainExitTimestampMonotonic"
	ExecMainPID                     string = "ExecMainPID"
	ExecMainStartTimestamp          string = "ExecMainStartTimestamp"
	ExecMainStartTimestampMonotonic string = "ExecMainStartTimestampMonotonic"
	ExecMainStatus                  string = "ExecMainStatus"
	ExecReload                      string = "ExecReload"
	ExecReloadEx                    string = "ExecReloadEx"
	ExecStart                       string = "ExecStart"
	ExecStartEx                     string = "ExecStartEx"
	FailureAction                   string = "FailureAction"
	FileDescriptorStoreMax          string = "FileDescriptorStoreMax"
	FinalKillSignal                 string = "FinalKillSignal"
	FragmentPath                    string = "FragmentPath"
	FreezerState                    string = "FreezerState"
	GID                             string = "GID"
	GuessMainPID                    string = "GuessMainPID"
	IOAccounting                    string = "IOAccounting"
	IOReadBytes                     string = "IOReadBytes"
	IOReadOperations                string = "IOReadOperations"
	IOSchedulingClass               string = "IOSchedulingClass"
	IOSchedulingPriority            string = "IOSchedulingPriority"
	IOWeight                        string = "IOWeight"
	IOWriteBytes                    string = "IOWriteBytes"
	IOWriteOperations               string = "IOWriteOperations"
	IPAccounting                    string = "IPAccounting"
	IPEgressBytes                   string = "IPEgressBytes"
	IPEgressPackets                 string = "IPEgressPackets"
	IPIngressBytes                  string = "IPIngressBytes"
	IPIngressPackets                string = "IPIngressPackets"
	Id                              string = "Id"
	IgnoreOnIsolate                 string = "IgnoreOnIsolate"
	IgnoreSIGPIPE                   string = "IgnoreSIGPIPE"
	InactiveEnterTimestampMonotonic string = "InactiveEnterTimestampMonotonic"
	InactiveExitTimestamp           string = "InactiveExitTimestamp"
	InactiveExitTimestampMonotonic  string = "InactiveExitTimestampMonotonic"
	InvocationID                    string = "InvocationID"
	JobRunningTimeoutUSec           string = "JobRunningTimeoutUSec"
	JobTimeoutAction                string = "JobTimeoutAction"
	JobTimeoutUSec                  string = "JobTimeoutUSec"
	KeyringMode                     string = "KeyringMode"
	KillMode                        string = "KillMode"
	KillSignal                      string = "KillSignal"
	LimitAS                         string = "LimitAS"
	LimitASSoft                     string = "LimitASSoft"
	LimitCORE                       string = "LimitCORE"
	LimitCORESoft                   string = "LimitCORESoft"
	LimitCPU                        string = "LimitCPU"
	LimitCPUSoft                    string = "LimitCPUSoft"
	LimitDATA                       string = "LimitDATA"
	LimitDATASoft                   string = "LimitDATASoft"
	LimitFSIZE                      string = "LimitFSIZE"
	LimitFSIZESoft                  string = "LimitFSIZESoft"
	LimitLOCKS                      string = "LimitLOCKS"
	LimitLOCKSSoft                  string = "LimitLOCKSSoft"
	LimitMEMLOCK                    string = "LimitMEMLOCK"
	LimitMEMLOCKSoft                string = "LimitMEMLOCKSoft"
	LimitMSGQUEUE                   string = "LimitMSGQUEUE"
	LimitMSGQUEUESoft               string = "LimitMSGQUEUESoft"
	LimitNICE                       string = "LimitNICE"
	LimitNICESoft                   string = "LimitNICESoft"
	LimitNOFILE                     string = "LimitNOFILE"
	LimitNOFILESoft                 string = "LimitNOFILESoft"
	LimitNPROC                      string = "LimitNPROC"
	LimitNPROCSoft                  string = "LimitNPROCSoft"
	LimitRSS                        string = "LimitRSS"
	LimitRSSSoft                    string = "LimitRSSSoft"
	LimitRTPRIO                     string = "LimitRTPRIO"
	LimitRTPRIOSoft                 string = "LimitRTPRIOSoft"
	LimitRTTIME                     string = "LimitRTTIME"
	LimitRTTIMESoft                 string = "LimitRTTIMESoft"
	LimitSIGPENDING                 string = "LimitSIGPENDING"
	LimitSIGPENDINGSoft             string = "LimitSIGPENDINGSoft"
	LimitSTACK                      string = "LimitSTACK"
	LimitSTACKSoft                  string = "LimitSTACKSoft"
	LoadState                       string = "LoadState"
	LockPersonality                 string = "LockPersonality"
	LogLevelMax                     string = "LogLevelMax"
	LogRateLimitBurst               string = "LogRateLimitBurst"
	LogRateLimitIntervalUSec        string = "LogRateLimitIntervalUSec"
	LogsDirectoryMode               string = "LogsDirectoryMode"
	MainPID                         string = "MainPID"
	ManagedOOMMemoryPressure        string = "ManagedOOMMemoryPressure"
	ManagedOOMMemoryPressureLimit   string = "ManagedOOMMemoryPressureLimit"
	ManagedOOMPreference            string = "ManagedOOMPreference"
	ManagedOOMSwap                  string = "ManagedOOMSwap"
	MemoryAccounting                string = "MemoryAccounting"
	MemoryCurrent                   string = "MemoryCurrent"
	MemoryDenyWriteExecute          string = "MemoryDenyWriteExecute"
	MemoryHigh                      string = "MemoryHigh"
	MemoryLimit                     string = "MemoryLimit"
	MemoryLow                       string = "MemoryLow"
	MemoryMax                       string = "MemoryMax"
	MemoryMin                       string = "MemoryMin"
	MemorySwapMax                   string = "MemorySwapMax"
	MountAPIVFS                     string = "MountAPIVFS"
	NFileDescriptorStore            string = "NFileDescriptorStore"
	NRestarts                       string = "NRestarts"
	NUMAPolicy                      string = "NUMAPolicy"
	Names                           string = "Names"
	NeedDaemonReload                string = "NeedDaemonReload"
	Nice                            string = "Nice"
	NoNewPrivileges                 string = "NoNewPrivileges"
	NonBlocking                     string = "NonBlocking"
	NotifyAccess                    string = "NotifyAccess"
	OOMPolicy                       string = "OOMPolicy"
	OOMScoreAdjust                  string = "OOMScoreAdjust"
	OnFailureJobMode                string = "OnFailureJobMode"
	PIDFile                         string = "PIDFile"
	Perpetual                       string = "Perpetual"
	PrivateDevices                  string = "PrivateDevices"
	PrivateIPC                      string = "PrivateIPC"
	PrivateMounts                   string = "PrivateMounts"
	PrivateNetwork                  string = "PrivateNetwork"
	PrivateTmp                      string = "PrivateTmp"
	PrivateUsers                    string = "PrivateUsers"
	ProcSubset                      string = "ProcSubset"
	ProtectClock                    string = "ProtectClock"
	ProtectControlGroups            string = "ProtectControlGroups"
	ProtectHome                     string = "ProtectHome"
	ProtectHostname                 string = "ProtectHostname"
	ProtectKernelLogs               string = "ProtectKernelLogs"
	ProtectKernelModules            string = "ProtectKernelModules"
	ProtectKernelTunables           string = "ProtectKernelTunables"
	ProtectProc                     string = "ProtectProc"
	ProtectSystem                   string = "ProtectSystem"
	RefuseManualStart               string = "RefuseManualStart"
	RefuseManualStop                string = "RefuseManualStop"
	ReloadResult                    string = "ReloadResult"
	RemainAfterExit                 string = "RemainAfterExit"
	RemoveIPC                       string = "RemoveIPC"
	Requires                        string = "Requires"
	Restart                         string = "Restart"
	RestartKillSignal               string = "RestartKillSignal"
	RestartUSec                     string = "RestartUSec"
	RestrictNamespaces              string = "RestrictNamespaces"
	RestrictRealtime                string = "RestrictRealtime"
	RestrictSUIDSGID                string = "RestrictSUIDSGID"
	Result                          string = "Result"
	RootDirectoryStartOnly          string = "RootDirectoryStartOnly"
	RuntimeDirectoryMode            string = "RuntimeDirectoryMode"
	RuntimeDirectoryPreserve        string = "RuntimeDirectoryPreserve"
	RuntimeMaxUSec                  string = "RuntimeMaxUSec"
	SameProcessGroup                string = "SameProcessGroup"
	SecureBits                      string = "SecureBits"
	SendSIGHUP                      string = "SendSIGHUP"
	SendSIGKILL                     string = "SendSIGKILL"
	Slice                           string = "Slice"
	StandardError                   string = "StandardError"
	StandardInput                   string = "StandardInput"
	StandardOutput                  string = "StandardOutput"
	StartLimitAction                string = "StartLimitAction"
	StartLimitBurst                 string = "StartLimitBurst"
	StartLimitIntervalUSec          string = "StartLimitIntervalUSec"
	StartupBlockIOWeight            string = "StartupBlockIOWeight"
	StartupCPUShares                string = "StartupCPUShares"
	StartupCPUWeight                string = "StartupCPUWeight"
	StartupIOWeight                 string = "StartupIOWeight"
	StateChangeTimestamp            string = "StateChangeTimestamp"
	StateChangeTimestampMonotonic   string = "StateChangeTimestampMonotonic"
	StateDirectoryMode              string = "StateDirectoryMode"
	StatusErrno                     string = "StatusErrno"
	StopWhenUnneeded                string = "StopWhenUnneeded"
	SubState                        string = "SubState"
	SuccessAction                   string = "SuccessAction"
	SyslogFacility                  string = "SyslogFacility"
	SyslogLevel                     string = "SyslogLevel"
	SyslogLevelPrefix               string = "SyslogLevelPrefix"
	SyslogPriority                  string = "SyslogPriority"
	SystemCallErrorNumber           string = "SystemCallErrorNumber"
	TTYReset                        string = "TTYReset"
	TTYVHangup                      string = "TTYVHangup"
	TTYVTDisallocate                string = "TTYVTDisallocate"
	TasksAccounting                 string = "TasksAccounting"
	TasksCurrent                    string = "TasksCurrent"
	TasksMax                        string = "TasksMax"
	TimeoutAbortUSec                string = "TimeoutAbortUSec"
	TimeoutCleanUSec                string = "TimeoutCleanUSec"
	TimeoutStartFailureMode         string = "TimeoutStartFailureMode"
	TimeoutStartUSec                string = "TimeoutStartUSec"
	TimeoutStopFailureMode          string = "TimeoutStopFailureMode"
	TimeoutStopUSec                 string = "TimeoutStopUSec"
	TimerSlackNSec                  string = "TimerSlackNSec"
	Transient                       string = "Transient"
	Type                            string = "Type"
	UID                             string = "UID"
	UMask                           string = "UMask"
	UnitFilePreset                  string = "UnitFilePreset"
	UnitFileState                   string = "UnitFileState"
	UtmpMode                        string = "UtmpMode"
	WantedBy                        string = "WantedBy"
	WatchdogSignal                  string = "WatchdogSignal"
	WatchdogTimestampMonotonic      string = "WatchdogTimestampMonotonic"
	WatchdogUSec                    string = "WatchdogUSec"
)

var properties = []string{
	ActiveEnterTimestamp,
	ActiveEnterTimestampMonotonic,
	ActiveExitTimestampMonotonic,
	ActiveState,
	After,
	AllowIsolate,
	AssertResult,
	AssertTimestamp,
	AssertTimestampMonotonic,
	Before,
	BlockIOAccounting,
	BlockIOWeight,
	CPUAccounting,
	CPUAffinityFromNUMA,
	CPUQuotaPerSecUSec,
	CPUQuotaPeriodUSec,
	CPUSchedulingPolicy,
	CPUSchedulingPriority,
	CPUSchedulingResetOnFork,
	CPUShares,
	CPUUsageNSec,
	CPUWeight,
	CacheDirectoryMode,
	CanFreeze,
	CanIsolate,
	CanReload,
	CanStart,
	CanStop,
	CapabilityBoundingSet,
	CleanResult,
	CollectMode,
	ConditionResult,
	ConditionTimestamp,
	ConditionTimestampMonotonic,
	ConfigurationDirectoryMode,
	Conflicts,
	ControlGroup,
	ControlPID,
	CoredumpFilter,
	DefaultDependencies,
	DefaultMemoryLow,
	DefaultMemoryMin,
	Delegate,
	Description,
	DevicePolicy,
	DynamicUser,
	EffectiveCPUs,
	EffectiveMemoryNodes,
	ExecMainCode,
	ExecMainExitTimestampMonotonic,
	ExecMainPID,
	ExecMainStartTimestamp,
	ExecMainStartTimestampMonotonic,
	ExecMainStatus,
	ExecReload,
	ExecReloadEx,
	ExecStart,
	ExecStartEx,
	FailureAction,
	FileDescriptorStoreMax,
	FinalKillSignal,
	FragmentPath,
	FreezerState,
	GID,
	GuessMainPID,
	IOAccounting,
	IOReadBytes,
	IOReadOperations,
	IOSchedulingClass,
	IOSchedulingPriority,
	IOWeight,
	IOWriteBytes,
	IOWriteOperations,
	IPAccounting,
	IPEgressBytes,
	IPEgressPackets,
	IPIngressBytes,
	IPIngressPackets,
	Id,
	IgnoreOnIsolate,
	IgnoreSIGPIPE,
	InactiveEnterTimestampMonotonic,
	InactiveExitTimestamp,
	InactiveExitTimestampMonotonic,
	InvocationID,
	JobRunningTimeoutUSec,
	JobTimeoutAction,
	JobTimeoutUSec,
	KeyringMode,
	KillMode,
	KillSignal,
	LimitAS,
	LimitASSoft,
	LimitCORE,
	LimitCORESoft,
	LimitCPU,
	LimitCPUSoft,
	LimitDATA,
	LimitDATASoft,
	LimitFSIZE,
	LimitFSIZESoft,
	LimitLOCKS,
	LimitLOCKSSoft,
	LimitMEMLOCK,
	LimitMEMLOCKSoft,
	LimitMSGQUEUE,
	LimitMSGQUEUESoft,
	LimitNICE,
	LimitNICESoft,
	LimitNOFILE,
	LimitNOFILESoft,
	LimitNPROC,
	LimitNPROCSoft,
	LimitRSS,
	LimitRSSSoft,
	LimitRTPRIO,
	LimitRTPRIOSoft,
	LimitRTTIME,
	LimitRTTIMESoft,
	LimitSIGPENDING,
	LimitSIGPENDINGSoft,
	LimitSTACK,
	LimitSTACKSoft,
	LoadState,
	LockPersonality,
	LogLevelMax,
	LogRateLimitBurst,
	LogRateLimitIntervalUSec,
	LogsDirectoryMode,
	MainPID,
	ManagedOOMMemoryPressure,
	ManagedOOMMemoryPressureLimit,
	ManagedOOMPreference,
	ManagedOOMSwap,
	MemoryAccounting,
	MemoryCurrent,
	MemoryDenyWriteExecute,
	MemoryHigh,
	MemoryLimit,
	MemoryLow,
	MemoryMax,
	MemoryMin,
	MemorySwapMax,
	MountAPIVFS,
	NFileDescriptorStore,
	NRestarts,
	NUMAPolicy,
	Names,
	NeedDaemonReload,
	Nice,
	NoNewPrivileges,
	NonBlocking,
	NotifyAccess,
	OOMPolicy,
	OOMScoreAdjust,
	OnFailureJobMode,
	PIDFile,
	Perpetual,
	PrivateDevices,
	PrivateIPC,
	PrivateMounts,
	PrivateNetwork,
	PrivateTmp,
	PrivateUsers,
	ProcSubset,
	ProtectClock,
	ProtectControlGroups,
	ProtectHome,
	ProtectHostname,
	ProtectKernelLogs,
	ProtectKernelModules,
	ProtectKernelTunables,
	ProtectProc,
	ProtectSystem,
	RefuseManualStart,
	RefuseManualStop,
	ReloadResult,
	RemainAfterExit,
	RemoveIPC,
	Requires,
	Restart,
	RestartKillSignal,
	RestartUSec,
	RestrictNamespaces,
	RestrictRealtime,
	RestrictSUIDSGID,
	Result,
	RootDirectoryStartOnly,
	RuntimeDirectoryMode,
	RuntimeDirectoryPreserve,
	RuntimeMaxUSec,
	SameProcessGroup,
	SecureBits,
	SendSIGHUP,
	SendSIGKILL,
	Slice,
	StandardError,
	StandardInput,
	StandardOutput,
	StartLimitAction,
	StartLimitBurst,
	StartLimitIntervalUSec,
	StartupBlockIOWeight,
	StartupCPUShares,
	StartupCPUWeight,
	StartupIOWeight,
	StateChangeTimestamp,
	StateChangeTimestampMonotonic,
	StateDirectoryMode,
	StatusErrno,
	StopWhenUnneeded,
	SubState,
	SuccessAction,
	SyslogFacility,
	SyslogLevel,
	SyslogLevelPrefix,
	SyslogPriority,
	SystemCallErrorNumber,
	TTYReset,
	TTYVHangup,
	TTYVTDisallocate,
	TasksAccounting,
	TasksCurrent,
	TasksMax,
	TimeoutAbortUSec,
	TimeoutCleanUSec,
	TimeoutStartFailureMode,
	TimeoutStartUSec,
	TimeoutStopFailureMode,
	TimeoutStopUSec,
	TimerSlackNSec,
	Transient,
	Type,
	UID,
	UMask,
	UnitFilePreset,
	UnitFileState,
	UtmpMode,
	WantedBy,
	WatchdogSignal,
	WatchdogTimestampMonotonic,
	WatchdogUSec,
}
