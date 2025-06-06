BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    username VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    name VARCHAR(32) NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS nodes (
    id UUID PRIMARY KEY CHECK (id <> '00000000-0000-0000-0000-000000000000'),
    label VARCHAR(32),
    parent_id UUID REFERENCES nodes(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS connections (
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    node_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, node_id)
);

CREATE TYPE style_type
AS ENUM("font", "textColor", "nodeColor", "borderColor");

CREATE TABLE IF NOT EXISTS styles (
    node_id UUID NOT NULL REFERENCES nodes(id) ON DELETE CASCADE,
    stl_type style_type NOT NULL,
    style VARCHAR(32) NOT NULL,
    PRIMARY KEY (node_id, stl_type)
);

COMMIT;