package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Olprog59/golog"
)

// APIHandler retourne un mux pour les routes API
func APIHandler() http.Handler {
	mux := http.NewServeMux()

	// Configuration des routes API
	mux.HandleFunc("GET /cluster", getClusterHandler)
	mux.HandleFunc("GET /clusters/count", getClusterCountHandler)
	mux.HandleFunc("GET /clusters/recent", getRecentClustersHandler)
	// Autres routes API...

	return mux
}

func getClusterCountHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("/api/clusters/count")

	w.Header().Set("Content-Type", "application/json")
}

// getClusterHandler retourne les informations du cluster
func getClusterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/api/cluster")
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Impossible de se connecter a Proxmox: %v", err), 500)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
}

// getRecentClustersHandler retourne les clusters récents
func getRecentClustersHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour récupérer les clusters récents
	// ...

	// Exemple de données
	clusters := []map[string]any{
		{
			"ID":          "1",
			"Name":        "Cluster Dev",
			"Description": "Environnement de développement",
			"NodeCount":   3,
			"Status":      "online",
			"LastUpdated": "2023-02-01 10:00",
		},
		// ...
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clusters)
}
