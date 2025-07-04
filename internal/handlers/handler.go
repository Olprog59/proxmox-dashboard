package handlers

import (
	"net/http"

	"github.com/Olprog59/dashboard-proxmox/internal/commons"
	"github.com/Olprog59/dashboard-proxmox/internal/models"
	"github.com/Olprog59/dashboard-proxmox/internal/service"
	templates_component_dashboard "github.com/Olprog59/dashboard-proxmox/internal/templates/components/dashboard"
)

type Handlers struct {
	serv *service.Service
}

func NewHandlers(serv *service.Service) *Handlers {
	return &Handlers{
		serv: serv,
	}
}

func (h *Handlers) DashboardCountHandler(w http.ResponseWriter, r *http.Request) {
	countCluster := commons.CountCluster
	countNodes := h.serv.CountClustersByType(models.Node)
	countVMs := h.serv.CountClustersByType(models.Qemu)
	countLXCs := h.serv.CountClustersByType(models.Lxc)
	templates_component_dashboard.DashboardCount(countCluster, countNodes, countVMs, countLXCs).Render(r.Context(), w)
}
