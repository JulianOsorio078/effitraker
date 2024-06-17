package config

import (
	"encoding/json"
	"log"
	"os"
)

type DBConfig struct {
	Port     int    `json:"3"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func LoadConfig() (DBConfig, error) {
	var config DBConfig
	filename := "internal/config.json"

	// Leer el archivo JSON
	jsonFile, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error al leer el archivo de configuración: %v", err)
		return config, err
	}

	// Parsear el archivo JSON
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Printf("Error al parsear el archivo de configuración JSON: %v", err)
		return config, err
	}

	return config, nil
}
