package middleware

import (
	"net/http"
	"time"

	"github.com/Olprog59/golog"
)

// WithLogging ajoute un logging à chaque requête
func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Appel du handler suivant
		next.ServeHTTP(w, r)

		// Logging après traitement
		golog.Info(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// WithAuth vérifie l'authentification pour les routes API
func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vérification d'authentification (token, session, etc.)
		// Si ce n'est pas une route qui nécessite une authentification ou si l'authentification est valide
		// Pour l'exemple, on considère que toutes les requêtes sont authentifiées

		// Passer au handler suivant si authentifié
		next.ServeHTTP(w, r)
	})
}

// WithMetrics collecte des métriques sur les requêtes
func WithMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Collecter des métriques (temps de réponse, code status, etc.)
		// ...

		// Appel du handler suivant
		next.ServeHTTP(w, r)

		// Finalisation des métriques
		// ...
	})
}
