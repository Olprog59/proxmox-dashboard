package server

import (
	"io/fs"
	"net/http"
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/handlers"
	middleware "github.com/Olprog59/dashboard-proxmox/internal/middlewares"
	"github.com/Olprog59/dashboard-proxmox/web"
)

// Server encapsule l'application HTTP
type Server struct {
	config *config.Config
	router *http.ServeMux
}

// NewServer crée une nouvelle instance de serveur
func NewServer(cfg *config.Config) *Server {
	s := &Server{
		config: cfg,
		router: http.NewServeMux(),
	}

	// Initialiser les templates
	// if err := web.InitTemplates(); err != nil {
	// 	log.Fatalf("Erreur lors du chargement des templates: %v", err)
	// }

	// Configurer les routes
	s.setupRoutes(cfg)

	return s
}

// setupRoutes configure toutes les routes du serveur
func (s *Server) setupRoutes(cfg *config.Config) {
	// Middleware de logging global
	// handler := middleware.WithLogging(s.router)

	hls := handlers.NewHandlers(cfg)

	// Fichiers statiques
	subStatic, _ := fs.Sub(web.StaticFiles, "static")
	s.router.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.FS(subStatic)),
		),
	)

	// Routes de l'interface utilisateur
	s.router.HandleFunc("GET /{$}", hls.IndexHandler)
	s.router.HandleFunc("GET /clusters/count", hls.ClustersCountHandler)

	s.router.HandleFunc("GET /nodes/count", hls.NodesCountHandler)
	s.router.HandleFunc("GET /nodes/vms/count", hls.VMsCountHandler)
	s.router.HandleFunc("GET /nodes/lxc/count", hls.LXCsCountHandler)
	s.router.HandleFunc("GET /nodes/dashboard-count", hls.DashboardCountHandler)
	// s.router.HandleFunc("GET /nodes", handlers.NodesHandler)
	// s.router.HandleFunc("GET /vms", handlers.VMsHandler)
	// s.router.HandleFunc("GET /lxcs", handlers.LXCsHandler)

	// // Middleware d'authentification pour les routes API
	// apiHandler := middleware.WithAuth(handlers.APIHandler())
	// // Routes API
	// s.router.Handle("/api/", http.StripPrefix("/api", apiHandler))
}

// Start démarre le serveur HTTP
func (s *Server) Start() error {
	server := &http.Server{
		Addr:         s.config.ServerAddr.Address,
		Handler:      middleware.WithMetrics(middleware.WithLogging(s.router)), // Middlewares globaux
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return server.ListenAndServe()
}
