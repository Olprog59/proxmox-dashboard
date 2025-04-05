package service

import (
	"sort"
	"sync"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/models"
	"github.com/Olprog59/golog"
)

type Service struct {
	mu           sync.RWMutex
	cfg          *config.Config
	clusters     []models.ClusterResource
	clustersInfo []models.Cluster
}

// NewServer crée une nouvelle instance de serveur
func NewService(cfg *config.Config) *Service {
	s := &Service{
		cfg:          cfg,
		clusters:     make([]models.ClusterResource, 0),
		clustersInfo: make([]models.Cluster, 0),
	}
	return s
}

func (s *Service) UpdateConfig(cfg *config.Config) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cfg = cfg

	// Vider les caches pour qu'ils soient rechargés avec la nouvelle config
	s.clusters = make([]models.ClusterResource, 0)
	s.clustersInfo = make([]models.Cluster, 0)

	golog.Info("Configuration du service mise à jour avec succès")
}

// GetClusterResources récupère les ressources des clusters
func (s *Service) GetClusterResources() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	newClusters := make([]models.ClusterResource, 0)

	for _, cluster := range s.cfg.Clusters {
		clusters, err := DoRequest[models.ClusterResource]("GET", "/cluster/resources", cluster, nil)
		if err != nil {
			golog.Err("%s", err.Error())
			return err
		}
		newClusters = append(newClusters, clusters...)
	}

	s.clusters = newClusters
	return nil
}

// GetClusters récupère les informations sur les clusters
func (s *Service) GetClusters() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	newClustersInfo := make([]models.Cluster, 0)

	for _, cluster := range s.cfg.Clusters {
		clusters, err := DoRequest[models.Cluster]("GET", "/cluster/status", cluster, nil)
		if err != nil {
			golog.Err("%s", err.Error())
			return err
		}
		newClustersInfo = append(newClustersInfo, clusters...)
	}

	s.clustersInfo = newClustersInfo
	return nil
}

// GetClustersData retourne les données des clusters stockées en mémoire
func (s *Service) GetClustersData() []models.ClusterResource {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Retourner une copie pour éviter les problèmes de concurrence
	result := make([]models.ClusterResource, len(s.clusters))
	copy(result, s.clusters)

	return result
}

// GetClustersInfo retourne les informations des clusters stockées en mémoire
func (s *Service) GetClustersInfo() []models.Cluster {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Retourner une copie pour éviter les problèmes de concurrence
	result := make([]models.Cluster, len(s.clustersInfo))
	copy(result, s.clustersInfo)

	return result
}

// GetClustersLastUptime retourne les nodes ayant un uptime plus récent
func (s *Service) GetClustersLastUptime() []models.ClusterResource {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Retourner une copie pour éviter les problèmes de concurrence
	result := make([]models.ClusterResource, len(s.clusters))

	for _, b := range s.clusters {
		if *b.Type == models.Node {
			result = append(result, b)
		}
	}

	sort.Slice(result, func(i int, j int) bool {
		if result[i].Uptime != nil && result[j].Uptime != nil {
			golog.Warn("%d", result[i].Uptime)
			return *result[i].Uptime < *result[j].Uptime
		}
		return false
	})

	return result
}

// CountClustersByType compte le nombre d'éléments d'un certain type dans tous les clusters
func (s *Service) CountClustersByType(t models.Type) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return countTypes(s.clusters, t)
}

// GetClusterCount retourne le nombre de clusters configurés
func (s *Service) GetClusterCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.cfg.Clusters)
}

func countTypes(clusters []models.ClusterResource, t models.Type) int {
	count := 0
	for _, d := range clusters {
		if d.Type != nil {
			if *d.Type == t {
				count += 1
			}
		}
	}
	return count
}
