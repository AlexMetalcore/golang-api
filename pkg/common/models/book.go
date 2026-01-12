package models

import (
	"errors"

	"encoding/json"

	"gorm.io/gorm"
)

type JSONB map[string]interface{}

type Book struct {
	gorm.Model
	Title          string `json:"title"`
	Author         string `json:"author"`
	Description    string `json:"description"`
	AdditionalData JSONB  `Gorm:"type:jsonb;serializer:json" json:"additional_data"`
}

func (jsonField *JSONB) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(data, &jsonField)
}
