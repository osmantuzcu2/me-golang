package utils

import (
	"encoding/json"
	"os"

	"github.com/synthonier/me-sniper/pkg/models"
)

func LoadCollections() (map[string]*models.Token, error) {
	var collections map[string]*models.Token
	req, _ := os.ReadFile("./data/collections2.json")
	err := json.Unmarshal(req, &collections)
	if err != nil {
		return nil, err
	}

	return collections, nil
}
