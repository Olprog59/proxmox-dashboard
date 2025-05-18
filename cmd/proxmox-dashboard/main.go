package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal/cache"
	"github.com/Olprog59/dashboard-proxmox/internal/commons"
	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/server"
	"github.com/Olprog59/dashboard-proxmox/internal/service"
	"github.com/Olprog59/golog"
)

var defaultExpiry = 4 * time.Second

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

	useRedis := false

	// Initialiser le cache
	config.InitCache(useRedis, defaultExpiry)

	// Créer et configurer le serveur
	serv := service.NewService(cfg)
	srv := server.NewServer(cfg, serv)

	commons.CountCluster = len(cfg.Clusters)

	go load(serv, srv)

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

func load(serv *service.Service, _ *server.Server) {
	var (
		cancel context.CancelFunc
		wg     sync.WaitGroup
	)

	// Fermeture propre à la sortie
	defer func() {
		if cancel != nil {
			cancel()
			wg.Wait()
		}
	}()

	for {
		// Annuler les services précédents s'ils existent
		if cancel != nil {
			cancel()
			wg.Wait()
		}

		ctx, newCancel := context.WithCancel(context.Background())
		cancel = newCancel

		wg.Add(1)
		go func() {
			defer wg.Done()
			runServices(ctx, serv)
		}()

		// Attendre un rechargement ou la fin du contexte
		select {
		case <-commons.ConfigReloadCh:
			golog.Info("Relancement des services suite au rechargement de la configuration")
		case <-ctx.Done():
			return
		}
	}
}

func runServices(ctx context.Context, serv *service.Service) {
	defer func() {
		if v := recover(); v != nil {
			golog.Err("Service crashed. Redémarrage...")
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(defaultExpiry):
		}

		updateCache(serv)
	}
}

func updateCache(serv *service.Service) {
	// golog.Info("Récupération des données...")

	// Gestion des clusters
	if _, err := config.Cache.GetClusters(); err != nil {
		handleClusterError(serv, err)
	}

	// Gestion des ressources
	if _, err := config.Cache.GetResources(); err != nil {
		handleResourceError(serv, err)
	}
}

func handleClusterError(serv *service.Service, err error) {
	if errors.Is(err, cache.ErrCacheExpired) || errors.Is(err, cache.ErrNotFound) {
		// golog.Info("Cache expiré/non trouvé. Mise à jour...")
		clusters, err := serv.GetClusters()
		if err != nil {
			golog.Err("Erreur récupération clusters: %+v", err)
			return
		}
		_ = config.Cache.SetClusters(clusters, 0)
	} else {
		golog.Err("Erreur cache clusters: %+v", err)
	}
}

func handleResourceError(serv *service.Service, err error) {
	if errors.Is(err, cache.ErrCacheExpired) || errors.Is(err, cache.ErrNotFound) {
		// golog.Info("Cache expiré/non trouvé. Mise à jour...")
		resources, err := serv.GetNodesResource()
		if err != nil {
			golog.Err("Erreur récupération ressources: %+v", err)
			return
		}
		_ = config.Cache.SetResources(resources, 0)
	} else {
		golog.Err("Erreur cache ressources: %+v", err)
	}
}
