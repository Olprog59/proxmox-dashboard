package service

import (
	"sort"
	"sync"

	"github.com/Olprog59/dashboard-proxmox/internal/config"
	"github.com/Olprog59/dashboard-proxmox/internal/models"
	"github.com/Olprog59/golog"
)

type Service struct {
	mu            sync.RWMutex
	cfg           *config.Config
	nodesResource []*models.NodeResource
	clustersInfo  []*models.Cluster
}

// NewServer crée une nouvelle instance de serveur
func NewService(cfg *config.Config) *Service {
	s := &Service{
		cfg:           cfg,
		nodesResource: make([]*models.NodeResource, 0),
		clustersInfo:  make([]*models.Cluster, 0),
	}
	return s
}

func (s *Service) UpdateConfig(cfg *config.Config) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cfg = cfg

	// Vider les caches pour qu'ils soient rechargés avec la nouvelle config
	s.nodesResource = make([]*models.NodeResource, 0)
	s.clustersInfo = make([]*models.Cluster, 0)

	golog.Info("Configuration du service mise à jour avec succès")
}

// GetClusterResources récupère les ressources des clusters
func (s *Service) GetNodesResource() ([]*models.NodeResource, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newClusters := make([]*models.NodeResource, 0)

	for _, cluster := range s.cfg.Clusters {
		clusters, err := DoRequest[*models.NodeResource]("GET", "/cluster/resources", cluster, nil)
		if err != nil {
			golog.Err("%s", err.Error())
			return nil, err
		}
		newClusters = append(newClusters, clusters...)
	}

	s.nodesResource = newClusters
	return newClusters, nil
}

// GetClusters récupère les informations sur les clusters
func (s *Service) GetClusters() ([]*models.Cluster, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newClustersInfo := make([]*models.Cluster, 0)

	for _, cluster := range s.cfg.Clusters {
		clusters, err := DoRequest[*models.Cluster]("GET", "/cluster/status", cluster, nil)
		if err != nil {
			golog.Err("%s", err.Error())
			return nil, err
		}
		newClustersInfo = append(newClustersInfo, clusters...)
	}

	s.clustersInfo = newClustersInfo
	return newClustersInfo, nil
}

// GetClustersLastUptime retourne les nodes ayant un uptime plus récent
func (s *Service) GetClustersLastUptime() []*models.NodeResource {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Retourner une copie pour éviter les problèmes de concurrence
	result := make([]*models.NodeResource, len(s.nodesResource))

	for _, b := range s.nodesResource {
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

	return countTypes(s.nodesResource, t)
}

// GetClusterCount retourne le nombre de clusters configurés
func (s *Service) GetClusterCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.cfg.Clusters)
}

func countTypes(clusters []*models.NodeResource, t models.Type) int {
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
