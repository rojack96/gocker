package run

import (
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
	"strconv"
)

const (
	// Options:
	addHost           = "--add-host"
	annotation        = "--annotation"
	blkioWeight       = "--blkio-weight"
	blkioWeightDevice = "--blkio-weight-device"
	cgroupParent      = "--cgroup-parent"
	cgroupns          = "--cgroups"

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
	return &Run{command: r.command + helpers.String(cgns, cgroupns)}
}
