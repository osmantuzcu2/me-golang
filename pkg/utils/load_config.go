package utils

import (
	"encoding/json"
	"os"

	"github.com/synthonier/me-sniper/pkg/models"
)

func LoadConfig() (*models.Config, error) {
	var config *models.Config
	req, _ := os.ReadFile("./data/config.json")
	err := json.Unmarshal(req, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
