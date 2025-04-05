package service

import (
	"sort"

	"github.com/Olprog59/dashboard-proxmox/internal/models"
)

func (s *Service) DashboardGetNodeLxcOrVm(t models.Type) []models.ClusterResource {
	lxc := make([]models.ClusterResource, 0)
	for _, v := range s.clusters {
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
