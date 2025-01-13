package helpers

import (
	"encoding/json"
	"os"
)

func UpdateJson(server string) error {
	filepath := "./docs/swagger.json"
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	var doc map[string]interface{}

	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	
	doc["servers"] = []map[string]string{
		{"url": server},
	}

	updatedData, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath, updatedData, 0644); err != nil {
		return err
	}

	

	return nil
}