CREATE TYPE feature_layer AS ENUM ('mobile', 'backend', 'frontend');

CREATE TABLE feature_flags (
  feature_id uuid NOT NULL DEFAULT uuid_generate_v1(),
  description TEXT NOT NULL,
  is_global BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  feature_layer feature_layer NOT NULL DEFAULT 'frontend',
  PRIMARY KEY (feature_id)
);

CREATE TABLE users_to_feature_flags (
    user_id uuid NOT NULL DEFAULT uuid_generate_v1(),
    feature_flag_id uuid NOT NULL DEFAULT uuid_generate_v1(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id),
    FOREIGN KEY (feature_flag_id) REFERENCES feature_flags (feature_id) ON DELETE CASCADE
);

CREATE TABLE organizations_to_feature_flags (
    organization_id uuid NOT NULL DEFAULT uuid_generate_v1(),
    feature_flag_id uuid NOT NULL DEFAULT uuid_generate_v1(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (organization_id),
    FOREIGN KEY (feature_flag_id) REFERENCES feature_flags (feature_id) ON DELETE CASCADE
);