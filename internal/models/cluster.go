package models

import "time"

type Cluster struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Nodes    []Node    `json:"nodes"`
	Status   string    `json:"status"`
	LastSeen time.Time `json:"last_seen"`
}
