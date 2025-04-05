package handlers

import (
	"net/http"

	"github.com/Olprog59/dashboard-proxmox/internal/templates"
)

// IndexHandler gère la page d'accueil
func (h *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates.Dashboard("Tableau de bord", "dashboard-home").Render(r.Context(), w)
}
