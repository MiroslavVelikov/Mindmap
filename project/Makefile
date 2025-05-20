export DB_NAME=mindmap
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_HOST ?= localhost

UP_SQL=./db/up.sql
DOWN_SQL=./db/down.sql

up:
	@echo "Running up.sql to create or update the schema..."
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST) -f $(UP_SQL)

down:
	@echo "Running down.sql to remove or revert schema changes..."
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST) -f $(DOWN_SQL)

reset: down up
	@echo "Database reset complete!"

run:
	@echo "Starting Docker Compose with rebuild..."
	docker-compose up --build

run-detached:
	@echo "Starting Docker Compose with detached mode..."
	docker-compose up --build -d

stop:
	@echo "Stopping Docker Compose..."
	docker-compose down