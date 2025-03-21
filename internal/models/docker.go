package models

import "time"

type Docker struct {
	ContainerID   string    `json:"container_id"`
	Name          string    `json:"name"`
	Image         string    `json:"image"`
	Status        string    `json:"status"`
	Created       time.Time `json:"created"`
	PortMappings  []string  `json:"port_mappings"`
	LastRestarted time.Time `json:"last_restarted"`
}
