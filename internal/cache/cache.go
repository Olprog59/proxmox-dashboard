package cache

import (
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal/models"
)

type CacheProvider interface {
	GetClusters() ([]*models.Cluster, error)
	GetResources() ([]*models.NodeResource, error)
	SetClusters(clusters []*models.Cluster, expiry time.Duration) error
	SetResources(resources []*models.NodeResource, expiry time.Duration) error
	ForceExpiration()
}

func NewCache(useRedis bool, defaultExpiry time.Duration) CacheProvider {
	if useRedis {
		return NewRedisCache("localhost:6379", defaultExpiry)
	}
	return NewMemoryCache(defaultExpiry)
}
