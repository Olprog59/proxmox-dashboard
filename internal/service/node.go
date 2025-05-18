package service

import (
	"sort"

	"github.com/Olprog59/dashboard-proxmox/internal/models"
)

func DashboardGetNodeLxcOrVm(c []*models.NodeResource, t models.Type) []*models.NodeResource {
	lxc := make([]*models.NodeResource, 0)
	for _, v := range c {
		if v.Type != nil && *v.Type == t {
			lxc = append(lxc, v)
		}
	}
	sort.Slice(lxc, func(i, j int) bool {
		if lxc[i].Uptime != nil && lxc[j].Uptime != nil {
			return *lxc[i].Uptime == 0
		}
		return false
	})
	return lxc
}
