package handlers

import (
	"fmt"
	"net/http"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/models"
	"github.com/Olprog59/dashboard-proxmox/internal/service"
	"github.com/Olprog59/dashboard-proxmox/web"
)

type Handlers struct {
	cfg *config.Config
}

func NewHandlers(cfg *config.Config) *Handlers {
	return &Handlers{cfg: cfg}
}

// IndexHandler gère la page d'accueil
func (h *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":       "Tableau de bord",
		"CurrentPage": "dashboard-home",
	}

	web.RenderTemplate(w, "index", data)
}

func (h *Handlers) DashboardCountHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		CountClusters int
		CountNodes    int
		CountLXCs     int
		CountVMs      int
	}

	data.CountClusters = len(h.cfg.Clusters)

	for _, cluster := range h.cfg.Clusters {
		nodes, err := service.DoRequest[models.Node]("GET", "/nodes", cluster, nil)
		if err != nil {
			http.Error(w, "Request for nodes error", http.StatusBadRequest)
			return
		}

		data.CountNodes += len(nodes)

		for _, node := range nodes {
			lxc, err := service.DoRequest[models.LXC]("GET", "/nodes/"+*node.Name+"/lxc", cluster, nil)
			if err != nil {
				http.Error(w, "Request for nodes error", http.StatusBadRequest)
			}
			data.CountLXCs += len(lxc)
		}

		for _, node := range nodes {
			vm, err := service.DoRequest[models.VM]("GET", "/nodes/"+*node.Name+"/qemu", cluster, nil)
			if err != nil {
				http.Error(w, "Request for nodes error", http.StatusBadRequest)
			}
			data.CountVMs += len(vm)
		}

	}

	web.RenderPartial(w, "dashboard-count", data)
}

func (h *Handlers) ClustersCountHandler(w http.ResponseWriter, r *http.Request) {
	count := len(h.cfg.Clusters)

	fmt.Fprintf(w, "%d", count)
}

func (h *Handlers) NodesCountHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	for _, cluster := range h.cfg.Clusters {
		nodes, err := service.DoRequest[models.Node]("GET", "/nodes", cluster, nil)
		if err != nil {
			http.Error(w, "Request for nodes error", http.StatusBadRequest)
		}
		count += len(nodes)

	}

	fmt.Fprintf(w, "%d", count)
}

func (h *Handlers) VMsCountHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	for _, cluster := range h.cfg.Clusters {
		nodes, err := service.DoRequest[models.Node]("GET", "/nodes", cluster, nil)
		if err != nil {
			http.Error(w, "Request for nodes error", http.StatusBadRequest)
			return
		}

		for _, node := range nodes {
			lxc, err := service.DoRequest[models.VM]("GET", "/nodes/"+*node.Name+"/qemu", cluster, nil)
			if err != nil {
				http.Error(w, "Request for nodes error", http.StatusBadRequest)
			}
			count += len(lxc)

		}

	}

	fmt.Fprintf(w, "%d", count)
}

func (h *Handlers) LXCsCountHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	for _, cluster := range h.cfg.Clusters {
		nodes, err := service.DoRequest[models.Node]("GET", "/nodes", cluster, nil)
		if err != nil {
			http.Error(w, "Request for nodes error", http.StatusBadRequest)
			return
		}

		for _, node := range nodes {
			lxc, err := service.DoRequest[models.LXC]("GET", "/nodes/"+*node.Name+"/lxc", cluster, nil)
			if err != nil {
				http.Error(w, "Request for nodes error", http.StatusBadRequest)
			}
			count += len(lxc)

		}

	}

	fmt.Fprintf(w, "%d", count)
}
