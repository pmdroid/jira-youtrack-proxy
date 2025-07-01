package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	YouTrackURL      string
	TypeFieldMap     map[string]string `json:"type_field_map"`
	PriorityFieldMap map[string]string `json:"priority_field_map"`
	AssigneeFieldMap map[string]string `json:"assignee_field_map"`
	Port             string
}

type FieldMappings struct {
	TypeFieldMap     map[string]string `json:"type_field_map"`
	PriorityFieldMap map[string]string `json:"priority_field_map"`
	AssigneeFieldMap map[string]string `json:"assignee_field_map"`
}

func LoadConfig() (*Config, error) {
	fieldMappingFilePath := getEnvOrDefault("FIELD_MAPPING_FILE_PATH", "configs/field_mappings.json")
	data, err := os.ReadFile(fieldMappingFilePath)
	if err != nil {
		return nil, err
	}

	var mappings FieldMappings
	err = json.Unmarshal(data, &mappings)
	if err != nil {
		return nil, err
	}

	return &Config{
		TypeFieldMap:     mappings.TypeFieldMap,
		PriorityFieldMap: mappings.PriorityFieldMap,
		AssigneeFieldMap: mappings.AssigneeFieldMap,
		YouTrackURL:      getEnvOrDefault("YOUTRACK_URL", "https://example.youtrack.cloud"),
		Port:             getEnvOrDefault("PORT", "8080"),
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}