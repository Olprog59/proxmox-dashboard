package handlers

import (
	"net/http"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/models"
	"github.com/Olprog59/dashboard-proxmox/internal/service"
	templates_component_dashboard "github.com/Olprog59/dashboard-proxmox/internal/templates/components/dashboard"
)

func (h *Handlers) DashBoardNodeLXCList(w http.ResponseWriter, r *http.Request) {
	resources, err := config.Cache.GetResources()
	if err != nil {
		resources, err = h.serv.GetNodesResource()
		if err != nil {
			http.Error(w, "Get clusters Handlers error", http.StatusBadGateway)
			return
		}
	}
	templates_component_dashboard.DashboardNodes("LXC", service.DashboardGetNodeLxcOrVm(resources, models.Lxc)).Render(r.Context(), w)
}

func (h *Handlers) DashBoardNodeVMList(w http.ResponseWriter, r *http.Request) {
	resources, err := config.Cache.GetResources()
	if err != nil {
		resources, err = h.serv.GetNodesResource()
		if err != nil {
			http.Error(w, "Get clusters Handlers error", http.StatusBadGateway)
			return
		}
	}
	templates_component_dashboard.DashboardNodes("VM", service.DashboardGetNodeLxcOrVm(resources, models.Qemu)).Render(r.Context(), w)
}
