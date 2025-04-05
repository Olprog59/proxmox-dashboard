package server

import (
	"context"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal"
	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/handlers"
	middleware "github.com/Olprog59/dashboard-proxmox/internal/middlewares"
	"github.com/Olprog59/dashboard-proxmox/internal/service"
	"github.com/Olprog59/golog"
)

// Server encapsule l'application HTTP
type Server struct {
	cfg    *config.Config
	serv   *service.Service
	router *http.ServeMux
}

// NewServer crée une nouvelle instance de serveur
func NewServer(cfg *config.Config, serv *service.Service) *Server {
	s := &Server{
		cfg:    cfg,
		serv:   serv,
		router: http.NewServeMux(),
	}

	// Configurer les routes
	s.setupRoutes()

	return s
}

// setupRoutes configure toutes les routes du serveur
func (s *Server) setupRoutes() {
	// Middleware de logging global
	// handler := middleware.WithLogging(s.router)

	hls := handlers.NewHandlers(s.serv)

	// Fichiers statiques
	subStatic, _ := fs.Sub(internal.StaticFiles, "static")
	s.router.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.FS(subStatic)),
		),
	)

	// Routes de l'interface utilisateur
	s.router.HandleFunc("GET /{$}", hls.IndexHandler)
	s.router.HandleFunc("GET /nodes", hls.ClustersStatusCardHandler)
	s.router.HandleFunc("GET /nodes/lxc", hls.DashBoardNodeLXCList)
	s.router.HandleFunc("GET /nodes/vm", hls.DashBoardNodeVMList)

	s.router.HandleFunc("GET /clusters/dashboard-count", hls.DashboardCountHandler)
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
		Addr:         s.cfg.ServerAddr.Address,
		Handler:      middleware.WithMetrics(middleware.WithLogging(s.router)), // Middlewares globaux
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			golog.Err("Error when running server: %s", err)
		}
	}()

	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		golog.Err("Error when shutting down server: %v", err)
		return err
	}

	return nil
}
