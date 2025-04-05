package main

import (
	"log"
	"sync"

	"github.com/Olprog59/dashboard-proxmox/internal/commons"
	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/server"
	"github.com/Olprog59/dashboard-proxmox/internal/service"
	"github.com/Olprog59/golog"
)

//go:generate npx @tailwindcss/cli -i ../../web/static/css/input.css -o ../../web/static/css/style.css --watch

func init() {
	golog.SetLanguage("fr")
	golog.EnableFileNameLogging()

	log.SetFlags(log.Lshortfile)
}

func main() {
	golog.Info("Welcome to Proxmox Manager")

	// Charger la configuration
	cfg, err := config.Load()
	if err != nil {
		golog.Err("Erreur lors du chargement de la configuration: %v", err)
	}

	configCh := config.WatchConfig()

	// Créer et configurer le serveur
	serv := service.NewService(cfg)
	srv := server.NewServer(cfg, serv)

	commons.CountCluster = len(cfg.Clusters)

	go load(serv)

	go handleConfigChanges(configCh, serv)

	// Démarrer le serveur
	golog.Info("Serveur démarré sur %s", cfg.ServerAddr.Address)
	if err := srv.Start(); err != nil {
		log.Fatalf("Erreur serveur: %v", err)
	}
}

func handleConfigChanges(configCh chan struct{}, serv *service.Service) {
	for {
		// Attendre une mise à jour de configuration
		<-configCh

		golog.Info("Rechargement de la configuration détecté")

		// Mettre à jour les services avec la nouvelle configuration
		cfg, err := config.Load()
		if err != nil {
			golog.Err("Erreur lors du rechargement de la configuration: %v", err)
			continue
		}

		// Mettre à jour le service avec la nouvelle configuration
		serv.UpdateConfig(cfg)

		// Mettre à jour le nombre de clusters
		commons.CountCluster = len(cfg.Clusters)

		// Notifier la goroutine de load de relancer les services
		commons.ConfigReloadCh <- struct{}{}

		golog.Info("Services relancés avec la nouvelle configuration")
	}
}

func load(serv *service.Service) {
	// WaitGroup pour synchroniser les goroutines de services
	var wg sync.WaitGroup

	// Fonction pour exécuter les services
	runServices := func() {
		defer wg.Done()
		defer func() {
			if v := recover(); v != nil {
				golog.Err("Service crashed. Redémarrage...")
			}
		}()

		// Exécuter les services
		if err := serv.GetClusterResources(); err != nil {
			golog.Err("%+v", err)
		}
		if err := serv.GetClusters(); err != nil {
			golog.Err("%+v", err)
		}
	}

	// Boucle principale
	for {
		// Ajouter deux tâches à la WaitGroup
		wg.Add(2)

		// Lancer les services
		go runServices()
		go runServices()

		// Attendre soit un signal de rechargement de configuration, soit la fin des services
		select {
		case <-commons.ConfigReloadCh:
			golog.Info("Relancement des services suite au rechargement de la configuration")
			// Ne pas attendre la fin des goroutines en cours, on va simplement en lancer de nouvelles
			// Les anciennes se termineront d'elles-mêmes grâce au defer recover
			continue
		}
	}
}
