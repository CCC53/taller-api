package models

import "github.com/google/uuid"

type SelectResponse struct {
	Value uuid.UUID `json:"value"`
	Label string    `json:"label"`
}
