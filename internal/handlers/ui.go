package handlers

import (
	"net/http"

	templates_pages "github.com/Olprog59/dashboard-proxmox/internal/templates/pages"
)

// IndexHandler gère la page d'accueil
func (h *Handlers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates_pages.Dashboard("Tableau de bord", "dashboard-home").Render(r.Context(), w)
}
