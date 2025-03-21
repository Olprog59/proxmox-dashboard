package config

import (
	"github.com/Olprog59/golog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config contient la configuration de l'application
type Config struct {
	ServerAddr ConfigServer    `mapstructure:"server"`
	LogLevel   ConfigLog       `mapstructure:"log"`
	Debug      bool            `mapstructure:"debug"`
	Clusters   []ClusterConfig `mapstructure:"clusters"`
}

type ConfigLog struct {
	level string `mapstructure:"level"`
}

type ConfigServer struct {
	Address string `mapstructure:"address"`
}

type ClusterConfig struct {
	Name        string `mapstructure:"name"`
	APIURL      string `mapstructure:"api_url"`
	SecretID    string `mapstructure:"secret_id"`
	SecretToken string `mapstructure:"secret_token"`
}

// Load charge la configuration depuis les variables d'environnement
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// dir, _ := os.Getwd()
	// golog.Info("%v", dir)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// golog.Debug("%v", config)

	return &Config{
		ServerAddr: config.ServerAddr,
		LogLevel:   config.LogLevel,
		Debug:      config.Debug,
		Clusters:   config.Clusters,
	}, nil
}

// Fonction pour surveiller les modifications du fichier de configuration
func WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		golog.Info("Le fichier de configuration a été modifié :", e.Name)
		// Recharger la configuration ou appliquer les modifications nécessaires
	})
}
