package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal/models"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client        *redis.Client
	defaultExpiry time.Duration
}

func NewRedisCache(addr string, defaultExpiry time.Duration) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{
		client:        client,
		defaultExpiry: defaultExpiry,
	}
}

func (r *RedisCache) GetClusters() ([]*models.Cluster, error) {
	val, err := r.client.Get(context.Background(), "clusters").Result()
	if err == redis.Nil {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	var clusters []*models.Cluster
	if err := json.Unmarshal([]byte(val), &clusters); err != nil {
		return nil, err
	}
	return clusters, nil
}

func (r *RedisCache) SetClusters(clusters []*models.Cluster, expiry time.Duration) error {
	data, err := json.Marshal(clusters)
	if err != nil {
		return err
	}
	if expiry == 0 {
		expiry = r.defaultExpiry
	}
	return r.client.Set(context.Background(), "clusters", data, expiry).Err()
}

func (r *RedisCache) GetResources() ([]*models.NodeResource, error) {
	val, err := r.client.Get(context.Background(), "resources").Result()
	if err == redis.Nil {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	var resources []*models.NodeResource
	if err := json.Unmarshal([]byte(val), &resources); err != nil {
		return nil, err
	}
	return resources, nil
}

func (r *RedisCache) SetResources(resources []*models.NodeResource, expiry time.Duration) error {
	data, err := json.Marshal(resources)
	if err != nil {
		return err
	}
	if expiry == 0 {
		expiry = r.defaultExpiry
	}
	return r.client.Set(context.Background(), "resources", data, expiry).Err()
}

func (r *RedisCache) ForceExpiration() {
	r.client.Del(context.Background(), "clusters", "resources")
}
