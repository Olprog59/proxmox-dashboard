package config

import (
	"log"
	"time"

	"github.com/Olprog59/dashboard-proxmox/internal/cache"
)

var Cache cache.CacheProvider

func InitCache(useRedis bool, defaultExpiry time.Duration) {
	Cache = cache.NewCache(useRedis, defaultExpiry)
	log.Println("Cache initialized with:", useRedis)
}
