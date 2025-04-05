package handlers

import (
	"net/http"

	"github.com/Olprog59/dashboard-proxmox/internal/templates"
)

func (h *Handlers) ClustersStatusCardHandler(w http.ResponseWriter, r *http.Request) {
	templates.ClusterStatus(h.serv.GetClustersInfo()).Render(r.Context(), w)
}
