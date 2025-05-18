package models

type Cluster struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Online  int8   `json:"online"`
	Level   string `json:"level"`
	NodeId  int    `json:"nodeid"`
	IP      string `json:"ip"`
	Local   int8   `json:"local"`
	Quorate int8   `json:"quorate"`
	// Nodes   []Node `json:"nodes"`
	Status string `json:"status"`
}

type NodeResource struct {
	Maxcpu     *int     `json:"maxcpu,omitempty"`
	CPU        *float64 `json:"cpu,omitempty"`
	Status     *Hastate `json:"status,omitempty"`
	Maxmem     *int64   `json:"maxmem,omitempty"`
	Template   *int64   `json:"template,omitempty"`
	Node       *string  `json:"node,omitempty"`
	ID         *string  `json:"id,omitempty"`
	Diskwrite  *int64   `json:"diskwrite,omitempty"`
	Vmid       *int64   `json:"vmid,omitempty"`
	Mem        *int64   `json:"mem,omitempty"`
	Disk       *int64   `json:"disk,omitempty"`
	Type       *Type    `json:"type,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Tags       *string  `json:"tags,omitempty"`
	Netout     *int64   `json:"netout,omitempty"`
	Diskread   *int64   `json:"diskread,omitempty"`
	Uptime     *int64   `json:"uptime,omitempty"`
	Maxdisk    *int64   `json:"maxdisk,omitempty"`
	Netin      *int64   `json:"netin,omitempty"`
	Hastate    *Hastate `json:"hastate,omitempty"`
	CgroupMode *int64   `json:"cgroup-mode,omitempty"`
	Level      *string  `json:"level,omitempty"`
	Plugintype *string  `json:"plugintype,omitempty"`
	Shared     *int64   `json:"shared,omitempty"`
	Storage    *string  `json:"storage,omitempty"`
	Content    *string  `json:"content,omitempty"`
	SDN        *string  `json:"sdn,omitempty"`
}

type Hastate string

const (
	Available Hastate = "available"
	Ok        Hastate = "ok"
	Online    Hastate = "online"
	Running   Hastate = "running"
	Stopped   Hastate = "stopped"
)

type Type string

const (
	Lxc     Type = "lxc"
	Node    Type = "node"
	Qemu    Type = "qemu"
	SDN     Type = "sdn"
	Storage Type = "storage"
)
