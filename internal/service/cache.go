package service

// func (s *Service) GetCachedClustersInfo() ([]models.Cluster, error) {
// 	// Tenter de récupérer les clusters depuis le cache
// 	clustersInfo, err := config.Cache.GetClusters()
// 	if err != nil {
// 		if errors.Is(err, cache.ErrCacheExpired) || errors.Is(err, cache.ErrNotFound) {
// 			// Si le cache est expiré ou non trouvé, récupérer les données depuis le service
// 			clustersInfo, err = s.GetClusters()
// 			if err != nil {
// 				return nil, err
// 			}
// 			// Mettre à jour le cache avec les nouvelles données
// 			_ = config.Cache.SetClusters(clustersInfo, 0)
// 		} else {
// 			// Retourner l'erreur si elle n'est pas liée à l'expiration ou à l'absence de cache
// 			return nil, err
// 		}
// 	}
// 	return clustersInfo, nil
// }
//
// func (s *Service) GetCachedNodesResources() ([]models.NodeResource, error) {
// 	// Tenter de récupérer les clusters depuis le cache
// 	resources, err := config.Cache.GetResources()
// 	if err != nil {
// 		if errors.Is(err, cache.ErrCacheExpired) || errors.Is(err, cache.ErrNotFound) {
// 			// Si le cache est expiré ou non trouvé, récupérer les données depuis le service
// 			resources, err = s.GetNodesResource()
// 			if err != nil {
// 				return nil, err
// 			}
// 			// Mettre à jour le cache avec les nouvelles données
// 			_ = config.Cache.SetResources(resources, 0)
// 		} else {
// 			// Retourner l'erreur si elle n'est pas liée à l'expiration ou à l'absence de cache
// 			return nil, err
// 		}
// 	}
// 	return resources, nil
// }
