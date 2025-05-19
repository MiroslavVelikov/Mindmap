DB_NAME=mindmap
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST ?= localhost

UP_SQL=./db/up.sql
DOWN_SQL=./db/down.sql

up:
	@echo "Running up.sql to create or update the schema..."
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST) -f $(UP_SQL)

down:
	@echo "Running down.sql to remove or revert schema changes..."
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST) -f $(DOWN_SQL)

reset: down up