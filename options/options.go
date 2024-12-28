package options

import (
	"github.com/rojack96/gocker/helpers"
)

const (
	AddHost             = "--add-host"
	all                 = "--all"
	Annotation          = "--annotation"
	Attach              = "--attach"
	BlkioWeightDevice   = "--blkio-weight-device"
	build               = "--build"
	CapAdd              = "--cap-add"
	CapDrop             = "--cap-drop"
	CgroupParent        = "--cgroup-parent"
	Cgroupns            = "--cgroups"
	Cidfile             = "--cidfile"
	CpuPeriod           = "--cpu-period"
	CpuQuota            = "--cpu-quota"
	CpuRtPeriod         = "--cpu-rt-period"
	CpuRtRuntime        = "--cpu-rt-runtime"
	CpuShares           = "--cpu-shares"
	Cpus                = "--cpus"
	CpusetCpus          = "--cpuset-cpus"
	CpusetMems          = "--cpuset-mems"
	detach              = "--detach"
	DetachKeys          = "--detach-keys"
	Device              = "--device"
	DeviceCgroupRule    = "--device-cgroup-rule"
	DeviceReadBps       = "--device-read-bps"
	DeviceReadIops      = "--device-read-iops"
	DeviceWriteBps      = "--device-write-bps"
	DeviceWriteIops     = "--device-write-iops"
	DisableContentTrust = "--disable-content-trust"
	Dns                 = "--dns"
	DnsOption           = "--dns-option"
	DnsSearch           = "--dns-search"
	Domainname          = "--domainname"
	dryRun              = "--dry-run"
	Entrypoint          = "--entrypoint"
	Env                 = "--env"
	EnvFile             = "--env-file"
	Expose              = "--expose"
	forceRecreate       = "--force-recreate"
	Gpus                = "--gpus"
	GroupAdd            = "--group-add"
	HealthCmd           = "--health-cmd"
	HealthInterval      = "--health-interval"
	HealthRetries       = "--health-retries"
	HealthStartInterval = "--health-start-interval"
	HealthStartPeriod   = "--health-start-period"
	HealthTimeout       = "--health-timeout"
	Help                = "--help"
	Hostname            = "--hostname"
	includeDeps         = "--include-deps"
	Init                = "--init"
	Interactive         = "--interactive"
	Ip                  = "--ip"
	Ip6                 = "--ip6"
	Ipc                 = "--ipc"
	Isolation           = "--isolation"
	KernelMemory        = "--kernel-memory"
	Label               = "--label"
	LabelFile           = "--label-file"
	Link                = "--link"
	LinkLocalIp         = "--link-local-ip"
	LogDriver           = "--log-driver"
	LogOpt              = "--log-opt"
	MacAddress          = "--mac-address"
	Memory              = "--memory"
	MemoryReservation   = "--memory-reservation"
	MemorySwap          = "--memory-swap"
	MemorySwappiness    = "--memory-swappiness"
	Mount               = "--mount"
	Name                = "--name"
	Network             = "--network"
	NetworkAlias        = "--network-alias"
	noBuild             = "--no-build"
	noColor             = "--no-color"
	noDeps              = "--no-deps"
	NoHealthcheck       = "--no-healthcheck"
	noLogPrefix         = "--no-log-prefix"
	noRecreate          = "--no-recreate"
	noTty               = "--no-tty"
	noTrunc             = "--no-trunc"
	OomKillDisable      = "--oom-kill-disable"
	OomScoreAdj         = "--oom-score-adj"
	Pid                 = "--pid"
	PidsLimit           = "--pids-limit"
	Platform            = "--platform"
	Privileged          = "--privileged"
	Publish             = "--publish"
	PublishAll          = "--publish-all"
	Pull                = "--pull"
	quiet               = "--quiet"
	quietPull           = "--quiet-pull"
	ReadOnly            = "--read-only"
	removeOrphans       = "--remove-orphans"
	Restart             = "--restart"
	Rm                  = "--rm"
	Runtime             = "--runtime"
	SecurityOpt         = "--security-opt"
	services            = "--services"
	ShmSize             = "--shm-size"
	SigProxy            = "--sig-proxy"
	StopSignal          = "--stop-signal"
	StopTimeout         = "--stop-timeout"
	StorageOpt          = "--storage-opt"
	Sysctl              = "--sysctl"
	Tmpfs               = "--tmpfs"
	timestamps          = "--timestamps"
	Tty                 = "--tty"
	Ulimit              = "--ulimit"
	User                = "--user"
	Userns              = "--userns"
	Uts                 = "--uts"
	Volume              = "--volume"
	VolumeDriver        = "--volume-driver"
	volumes             = "--volumes"
	VolumesFrom         = "--volumes-from"
	Workdir             = "--workdir"
)

func All() string {
	return helpers.Option(all)
}

func Build() string {
	return helpers.Option(build)
}

func DryRun() string {
	return helpers.Option(dryRun)
}

func ForceRecreate() string {
	return helpers.Option(forceRecreate)
}

func NoBuild() string {
	return helpers.Option(noBuild)
}

func NoRecreate() string {
	return helpers.Option(noRecreate)
}

func NoColor() string {
	return helpers.Option(noColor)
}

func NoLogPrefix() string {
	return helpers.Option(noLogPrefix)
}

func Timestamps() string {
	return helpers.Option(timestamps)
}

func Quiet() string {
	return helpers.Option(quiet)
}

func QuietPull() string {
	return helpers.Option(quietPull)
}

func RemoveOrphans() string {
	return helpers.Option(removeOrphans)
}

func Volumes() string {
	return helpers.Option(volumes)
}

func Detach() string {
	return helpers.Option(detach)
}

func ServicesOption() string {
	return helpers.Option(services)
}

func IncludeDeps() string {
	return helpers.Option(includeDeps)
}

func NoDeps() string {
	return helpers.Option(noDeps)
}

func NoTty() string {
	return helpers.Option(noTty)
}

func NoTrunc() string {
	return helpers.Option(noTrunc)
}
