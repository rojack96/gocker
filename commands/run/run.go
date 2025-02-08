package run

import (
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
	"strconv"
)

const (
	// Options:
	addHost             = "--add-host"
	annotation          = "--annotation"
	blkioWeight         = "--blkio-weight"
	blkioWeightDevice   = "--blkio-weight-device"
	cgroupParent        = "--cgroup-parent"
	cgroupns            = "--cgroups"
	cidfile             = "--cidfile"
	cpuPeriod           = "--cpu-period"
	cpuQuota            = "--cpu-quota"
	cpuRtPeriod         = "--cpu-rt-period"
	cpuRtRuntime        = "--cpu-rt-runtime"
	cpuShares           = "--cpu-shares"
	cpus                = "--cpus"
	cpusetCpus          = "--cpuset-cpus"
	cpusetMems          = "--cpuset-mems"
	detachKeys          = "--detach-keys"
	device              = "--device"
	deviceCgroupRule    = "--device-cgroup-rule"
	deviceReadBps       = "--device-read-bps"
	deviceReadIops      = "--device-read-iops"
	deviceWriteBps      = "--device-write-bps"
	deviceWriteIops     = "--device-write-iops"
	disableContentTrust = "--disable-content-trust"
	dns                 = "--dns"
	dnsOption           = "--dns-option"
	dnsSearch           = "--dns-search"
	domainname          = "--domainname"
	entrypoint          = "--entrypoint"
	env                 = "--env"
	envFile             = "--env-file"
	expose              = "--expose"
	gpus                = "--gpus"
	groupAdd            = "--group-add"
	healthCmd           = "--health-cmd"
	healthInterval      = "--health-interval"
	healthRetries       = "--health-retries"
	healthStartInterval = "--health-start-interval"
	healthStartPeriod   = "--health-start-period"
	healthTimeout       = "--health-timeout"
	help                = "--help"
	hostname            = "--hostname"
	//init                = "--init"
	interactive       = "--interactive"
	ip                = "--ip"
	ip6               = "--ip6"
	ipc               = "--ipc"
	isolation         = "--isolation"
	kernelMemory      = "--kernel-memory"
	label             = "--label"
	labelFile         = "--label-file"
	link              = "--link"
	linkLocalIp       = "--link-local-ip"
	logDriver         = "--log-driver"
	logOpt            = "--log-opt"
	macAddress        = "--mac-address"
	memory            = "--memory"
	memoryReservation = "--memory-reservation"
	memorySwap        = "--memory-swap"
	memorySwappiness  = "--memory-swappiness"
	mount             = "--mount"
	name              = "--name"
	network           = "--network"
	networkAlias      = "--network-alias"
	noHealthcheck     = "--no-healthcheck"
	oomKillDisable    = "--oom-kill-disable"
	oomScoreAdj       = "--oom-score-adj"
	pid               = "--pid"
	pidsLimit         = "--pids-limit"
	platform          = "--platform"
	privileged        = "--privileged"
	publish           = "--publish"
	publishAll        = "--publish-all"
	pull              = "--pull"
	quiet             = "--quiet"
	readOnly          = "--read-only"
	restart           = "--restart"
	rm                = "--rm"
	runtime           = "--runtime"
	securityOpt       = "--security-opt"
	shmSize           = "--shm-size"
	sigProxy          = "--sig-proxy"
	stopSignal        = "--stop-signal"
	stopTimeout       = "--stop-timeout"
	storageOpt        = "--storage-opt"
	sysctl            = "--sysctl"
	tmpfs             = "--tmpfs"
	tty               = "--tty"
	ulimit            = "--ulimit"
	user              = "--user"
	userns            = "--userns"
	uts               = "--uts"
	volume            = "--volume"
	volumeDriver      = "--volume-driver"
	volumesFrom       = "--volumes-from"
	workdir           = "--workdir"
	// Commands:
)

type Run struct {
	command string
}

type Cap string

const (
	Add  Cap = "add"
	Drop Cap = "drop"
)

type AddHost struct {
	Host string
	Ip   string
}

type Blkio struct {
	Device string
	Weight uint16
}

func New(cmd string) *Run {
	return &Run{command: cmd}
}

/*
	------------------------------
	|          OPTIONS           |
	------------------------------
*/

// AddHost - Add a custom host-to-IP mapping (host:ip)
func (r *Run) AddHost(list ...AddHost) *Run {
	var hosts []string

	if len(list) > 0 {
		for _, l := range list {
			hosts = append(hosts, l.Host+":"+l.Ip)
		}
	}

	return &Run{command: r.command + helpers.StringArray(addHost, hosts...)}
}

// Annotation - Add an annotation to the container (passed through to the OCI runtime) (default map[])
func (r *Run) Annotation(annotationMap ...helpers.KeyValueParameters) *Run {
	var result []string

	if len(annotationMap) > 0 {
		for _, am := range annotationMap {
			result = append(result, am.Key+"="+am.Value)
		}
	}

	return &Run{command: r.command + helpers.StringArray(annotation, result...)}
}

// Attach - Attach to STDIN, STDOUT or STDERR
func (r *Run) Attach(standard []string) *Run {
	return &Run{command: r.command + helpers.StringArray(options.Attach, standard...)}
}

// BlkioWeight - Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
func (r *Run) BlkioWeight(value uint16) *Run {
	return &Run{command: r.command + helpers.Uint(blkioWeight, uint64(value))}
}

// BlkioWeightDevice - Block IO weight (relative device weight) (default [])
func (r *Run) BlkioWeightDevice(blkio ...Blkio) *Run {
	var result []string

	if len(blkio) > 0 {
		for _, b := range blkio {
			result = append(result, b.Device+":"+strconv.FormatUint(uint64(b.Weight), 10))
		}
	}

	return &Run{command: r.command + helpers.StringArray(blkioWeightDevice, result...)}
}

// Cap - Add or Drop Linux capabilities
//
// capType [Add, Drop], by default is Add
func (r *Run) Cap(capType Cap, value ...string) *Run {
	var (
		capString string
	)

	switch capType {
	case Add, Drop:
		capString = string("--cap-" + capType)
	default:
		capString = string("--cap-" + Add)
	}

	return &Run{command: r.command + helpers.StringArray(capString, value...)}
}

// CgroupParent - Optional parent cgroup for the container
func (r *Run) CgroupParent(cgroup string) *Run {
	return &Run{command: r.command + helpers.String(cgroup, cgroupParent)}
}

// Cgroupns - Cgroup namespace to use (host|private)
//
// `host`:    Run the container in the Docker host's cgroup namespace
//
// `private`: Run the container in its own private cgroup namespace
//
// ' ' :        Use the cgroup namespace as configured by the
// default-cgroupns-mode option on the daemon (default)
func (r *Run) Cgroupns(cgns string) *Run {
	return &Run{command: r.command + helpers.String(cgroupns, cgns)}
}

// Cidfile - Write the container ID to the file
func (r *Run) Cidfile(cid string) *Run {
	return &Run{command: r.command + helpers.String(cidfile, cid)}
}

// CpuPeriod - Limit CPU CFS (Completely Fair Scheduler) period
func (r *Run) CpuPeriod(period int) *Run {
	return &Run{command: r.command + helpers.Int(cpuPeriod, period)}
}

// CpuQuota - Limit CPU CFS (Completely Fair Scheduler) quota
func (r *Run) CpuQuota(quota int) *Run {
	return &Run{command: r.command + helpers.Int(cpuQuota, quota)}
}

// CpuRtPeriod - Limit CPU real-time period in microseconds
func (r *Run) CpuRtPeriod(period int) *Run {
	return &Run{command: r.command + helpers.Int(cpuRtPeriod, period)}
}

// CpuRtRuntime - Limit CPU real-time runtime in microseconds
func (r *Run) CpuRtRuntime(runtime int) *Run {
	return &Run{command: r.command + helpers.Int(cpuRtRuntime, runtime)}
}

// CpuShared - CPU shares (relative weight)
func (r *Run) CpuShared(shared int) *Run {
	return &Run{command: r.command + helpers.Int(cpuShares, shared)}
}

// Cpus - Number of CPUs
func (r *Run) Cpus(cpu float64) *Run {
	return &Run{command: r.command + helpers.Float(cpus, cpu)}
}

// CpuSetCpus - CPUs in which to allow execution (0-3, 0,1)
func (r *Run) CpuSetCpus(cpu1, cpu2 string) *Run {
	return &Run{command: r.command + helpers.String(cpusetCpus, cpu1+","+cpu2)}
}

// CpuSetMems - MEMs in which to allow execution (0-3, 0,1)
func (r *Run) CpuSetMems(mem1, mem2 string) *Run {
	return &Run{command: r.command + helpers.String(cpusetMems, mem1+","+mem2)}
}

// Detach - Run container in background and print container ID
func (r *Run) Detach() *Run {
	return &Run{command: r.command + options.Detach()}
}

// DetachKeys -  Override the key sequence for detaching from a container.
func (r *Run) DetachKeys(value string) *Run {
	return &Run{command: r.command + helpers.String(detachKeys, value)}
}
