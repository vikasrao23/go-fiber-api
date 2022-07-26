// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type FeatureLayer string

const (
	FeatureLayerMobile   FeatureLayer = "mobile"
	FeatureLayerBackend  FeatureLayer = "backend"
	FeatureLayerFrontend FeatureLayer = "frontend"
)

func (e *FeatureLayer) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FeatureLayer(s)
	case string:
		*e = FeatureLayer(s)
	default:
		return fmt.Errorf("unsupported scan type for FeatureLayer: %T", src)
	}
	return nil
}

type FeatureFlag struct {
	FeatureID    uuid.UUID      `json:"feature_id"`
	Description  sql.NullString `json:"description"`
	IsGlobal     bool           `json:"is_global"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	FeatureLayer FeatureLayer   `json:"feature_layer"`
}

type OrganizationsToFeatureFlag struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	FeatureFlagID  uuid.UUID `json:"feature_flag_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UsersToFeatureFlag struct {
	UserID        uuid.UUID `json:"user_id"`
	FeatureFlagID uuid.UUID `json:"feature_flag_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
