package handlers

import (
	"net/http"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
	templates_component_dashboard "github.com/Olprog59/dashboard-proxmox/internal/templates/components/dashboard"
)

func (h *Handlers) ClustersStatusCardHandler(w http.ResponseWriter, r *http.Request) {
	// templates_component_dashboard.ClusterStatus(h.serv.GetClustersInfo()).Render(r.Context(), w)
	clustersInfo, err := config.Cache.GetClusters()
	if err != nil {
		clustersInfo, err = h.serv.GetClusters()
		if err != nil {
			http.Error(w, "Get clusters Handlers error", http.StatusBadGateway)
			return
		}
	}
	templates_component_dashboard.ClusterStatus(clustersInfo).Render(r.Context(), w)
}
