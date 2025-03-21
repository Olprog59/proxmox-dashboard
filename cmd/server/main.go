package main

import (
	"log"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/server"
	"github.com/Olprog59/dashboard-proxmox/web"
	"github.com/Olprog59/golog"
)

//go:generate npx @tailwindcss/cli -i ../../web/static/css/input.css -o ../../web/static/css/style.css --watch

func init() {
	golog.SetLanguage("fr")
	golog.EnableFileNameLogging()

	err := web.InitTemplates()
	if err != nil {
		golog.Err("%v", err)
	}
}

func main() {
	golog.Info("Welcome to Proxmox Manager")

	// Charger la configuration
	cfg, err := config.Load()
	if err != nil {
		golog.Err("Erreur lors du chargement de la configuration: %v", err)
	}
	config.WatchConfig()

	// Créer et configurer le serveur
	srv := server.NewServer(cfg)

	// Démarrer le serveur
	golog.Info("Serveur démarré sur %s", cfg.ServerAddr.Address)
	if err := srv.Start(); err != nil {
		log.Fatalf("Erreur serveur: %v", err)
	}
}
