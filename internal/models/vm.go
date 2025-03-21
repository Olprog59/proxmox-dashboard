package models

type VM struct {
	PID        *int64   `json:"pid,omitempty"`
	Maxdisk    *int64   `json:"maxdisk,omitempty"`
	Maxmem     *int64   `json:"maxmem,omitempty"`
	Uptime     *int64   `json:"uptime,omitempty"`
	Diskwrite  *int64   `json:"diskwrite,omitempty"`
	Netin      *int64   `json:"netin,omitempty"`
	CPU        *float64 `json:"cpu,omitempty"`
	Netout     *int64   `json:"netout,omitempty"`
	Status     *Status  `json:"status,omitempty"`
	Disk       *int64   `json:"disk,omitempty"`
	Mem        *int64   `json:"mem,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Cpus       *int64   `json:"cpus,omitempty"`
	Diskread   *int64   `json:"diskread,omitempty"`
	Tags       *string  `json:"tags,omitempty"`
	Vmid       *int64   `json:"vmid,omitempty"`
	Template   *int64   `json:"template,omitempty"`
	Serial     *int64   `json:"serial,omitempty"`
	BalloonMin *int64   `json:"balloon_min,omitempty"`
	Shares     *int64   `json:"shares,omitempty"`
}
