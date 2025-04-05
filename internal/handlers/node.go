package handlers

import (
	"net/http"

	"github.com/Olprog59/dashboard-proxmox/internal/models"
	"github.com/Olprog59/dashboard-proxmox/internal/templates"
)

func (h *Handlers) DashBoardNodeLXCList(w http.ResponseWriter, r *http.Request) {
	templates.DashboardNodes("LXC", h.serv.DashboardGetNodeLxcOrVm(models.Lxc)).Render(r.Context(), w)
}

func (h *Handlers) DashBoardNodeVMList(w http.ResponseWriter, r *http.Request) {
	templates.DashboardNodes("VM", h.serv.DashboardGetNodeLxcOrVm(models.Qemu)).Render(r.Context(), w)
}
