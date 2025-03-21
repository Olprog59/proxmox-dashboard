package models

import "time"

type Node struct {
	SSLFingerprint *string   `json:"ssl_fingerprint,omitempty"`
	Type           *string   `json:"type,omitempty"`
	CPU            *float64  `json:"cpu,omitempty"`
	Disk           *int64    `json:"disk,omitempty"`
	Maxdisk        *int64    `json:"maxdisk,omitempty"`
	Level          *string   `json:"level,omitempty"`
	Name           *string   `json:"node,omitempty"`
	ID             *string   `json:"id,omitempty"`
	Uptime         *int64    `json:"uptime,omitempty"`
	Mem            *int64    `json:"mem,omitempty"`
	Maxmem         *int64    `json:"maxmem,omitempty"`
	Status         *string   `json:"status,omitempty"`
	Maxcpu         *int64    `json:"maxcpu,omitempty"`
	VMs            []VM      `json:"vms"`
	LXCs           []LXC     `json:"lxcs"`
	LastSeen       time.Time `json:"last_seen"`
}

type NodeStatus struct {
	BootInfo      *BootInfo      `json:"boot-info,omitempty"`
	Ksm           *Ksm           `json:"ksm,omitempty"`
	CPU           *int64         `json:"cpu,omitempty"`
	Wait          *int64         `json:"wait,omitempty"`
	Loadavg       []string       `json:"loadavg,omitempty"`
	Kversion      *string        `json:"kversion,omitempty"`
	CurrentKernel *CurrentKernel `json:"current-kernel,omitempty"`
	Swap          *Memory        `json:"swap,omitempty"`
	Memory        *Memory        `json:"memory,omitempty"`
	Rootfs        *Memory        `json:"rootfs,omitempty"`
	Idle          *int64         `json:"idle,omitempty"`
	Cpuinfo       *Cpuinfo       `json:"cpuinfo,omitempty"`
	Pveversion    *string        `json:"pveversion,omitempty"`
	Uptime        *int64         `json:"uptime,omitempty"`
}

type BootInfo struct {
	Secureboot *int64  `json:"secureboot,omitempty"`
	Mode       *string `json:"mode,omitempty"`
}

type Cpuinfo struct {
	Model   *string `json:"model,omitempty"`
	Sockets *int64  `json:"sockets,omitempty"`
	Mhz     *string `json:"mhz,omitempty"`
	Cores   *int64  `json:"cores,omitempty"`
	Flags   *string `json:"flags,omitempty"`
	Cpus    *int64  `json:"cpus,omitempty"`
	UserHz  *int64  `json:"user_hz,omitempty"`
	Hvm     *string `json:"hvm,omitempty"`
}

type CurrentKernel struct {
	Release *string `json:"release,omitempty"`
	Version *string `json:"version,omitempty"`
	Sysname *string `json:"sysname,omitempty"`
	Machine *string `json:"machine,omitempty"`
}

type Ksm struct {
	Shared *int64 `json:"shared,omitempty"`
}

type Memory struct {
	Used  *int64 `json:"used,omitempty"`
	Total *int64 `json:"total,omitempty"`
	Free  *int64 `json:"free,omitempty"`
	Avail *int64 `json:"avail,omitempty"`
}
