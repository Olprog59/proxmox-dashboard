package cache

import (
	"errors"
	"sync"
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal/models"
)

// Erreurs standards du cache
var (
	ErrCacheExpired = errors.New("le cache a expiré")
	ErrNotFound     = errors.New("élément non trouvé dans le cache")
)

type MemoryCache struct {
	clusters      []*models.Cluster
	resources     []*models.NodeResource
	clusterMu     sync.RWMutex
	resourcesMu   sync.RWMutex
	clustersExp   time.Time
	resourcesExp  time.Time
	defaultExpiry time.Duration
}

func NewMemoryCache(defaultExpiry time.Duration) *MemoryCache {
	return &MemoryCache{
		defaultExpiry: defaultExpiry,
	}
}

func (m *MemoryCache) GetClusters() ([]*models.Cluster, error) {
	m.clusterMu.RLock()
	defer m.clusterMu.RUnlock()
	if time.Now().After(m.clustersExp) {
		return nil, ErrCacheExpired
	}
	if m.clusters == nil {
		return nil, ErrNotFound
	}
	return m.clusters, nil
}

func (m *MemoryCache) SetClusters(clusters []*models.Cluster, expiry time.Duration) error {
	m.clusterMu.Lock()
	defer m.clusterMu.Unlock()
	m.clusters = clusters
	if expiry == 0 {
		expiry = m.defaultExpiry
	}
	m.clustersExp = time.Now().Add(expiry)
	return nil
}

func (m *MemoryCache) GetResources() ([]*models.NodeResource, error) {
	m.resourcesMu.RLock()
	defer m.resourcesMu.RUnlock()
	if time.Now().After(m.resourcesExp) {
		return nil, ErrCacheExpired
	}
	if m.resources == nil {
		return nil, ErrNotFound
	}
	return m.resources, nil
}

func (m *MemoryCache) SetResources(resources []*models.NodeResource, expiry time.Duration) error {
	m.resourcesMu.Lock()
	defer m.resourcesMu.Unlock()
	m.resources = resources
	if expiry == 0 {
		expiry = m.defaultExpiry
	}
	m.resourcesExp = time.Now().Add(expiry)
	return nil
}

func (m *MemoryCache) ForceExpiration() {
	m.clusterMu.Lock()
	m.resourcesMu.Lock()
	defer m.clusterMu.Unlock()
	defer m.resourcesMu.Unlock()
	m.clustersExp = time.Time{}
	m.resourcesExp = time.Time{}
}
