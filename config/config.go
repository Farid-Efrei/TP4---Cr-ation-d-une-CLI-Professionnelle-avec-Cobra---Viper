package config

import (
	"TP4/internal/storage"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func NewStoreFromViper(v *viper.Viper) (storage.Storer, error) {
	storeType := v.GetString("storage.type")
	if storeType == "" {
		fmt.Println("INFO: storage.type non défini, utilisation de 'json' par défaut")
		storeType = "json"
	}
	switch storeType {
	case "memory":
		return storage.NewMemoryStore(), nil
	case "json":
		path := v.GetString("storage.json_path")
		if path == "" {
			path = "data/contacts.json"
			fmt.Println("INFO: storage.json_path non défini, utilisation de 'data/contacts.json' par défaut")
		}
		// Créer le répertoire si nécessaire
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return nil, fmt.Errorf("impossible de créer le répertoire pour le stockage JSON: %v", err)
		}
		return storage.NewJSONStore(path)
	case "gorm":
		path := v.GetString("storage.gorm_path")
		if path == "" {
			path = "data/crm.db"
			fmt.Println("INFO: storage.gorm_path non défini, utilisation de 'data/crm.db' par défaut")
		}
		return storage.NewGORMStore(path)
	default:
		return nil, fmt.Errorf("type de stockage inconnu: %s", storeType)
	}
}
