// Package svix this file is @generated DO NOT EDIT
package models

type EventTypeFromOpenApi struct {
	Deprecated   bool            `json:"deprecated"`
	Description  string          `json:"description"`
	FeatureFlag  *string         `json:"featureFlag,omitempty"`
	FeatureFlags []string        `json:"featureFlags,omitempty"`
	GroupName    *string         `json:"groupName,omitempty"` // The event type group's name
	Name         string          `json:"name"`                // The event type's name
	Schemas      *map[string]any `json:"schemas,omitempty"`
}
