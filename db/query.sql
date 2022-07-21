-- name: GetAllFeatureFlags :many
SELECT * FROM feature_flags;

-- name: CreateFeatureFlag :one
INSERT INTO feature_flags (
    description,
    is_global
    ) VALUES ($1, $2)
RETURNING *;

-- name: GetFeatureFlagByFeatureLayer :many
SELECT * FROM feature_flags 
WHERE feature_layer = $1;

-- name: DeleteFeatureFlag :exec
DELETE FROM feature_flags 
WHERE feature_id = $1;