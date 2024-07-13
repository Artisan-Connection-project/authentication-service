# Variables
CURRENT_DIR=$(shell pwd)
DB_URL=postgres://postgres:1702@localhost:5432/authentication?sslmode=disable

# Default target
all: create-dirs create-migrations

# Target to create directories
create-dirs:
	@echo "Enter the path:"
	@read -r path; \
	if [ -z "$$path" ]; then \
		echo "Path is required"; \
		exit 1; \
	fi; \
	./scripts/create_dirs.sh "$$path"

# Target to create migrations
create-migrations:
	@echo "Enter table name:"
	@read -r table_name; \
	if [ -z "$$table_name" ]; then \
		echo "Table name is required"; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir migration -seq "$$table_name"

# Clean target (optional example)
clean:
	rm -f .path_env

# Migration commands
migrate_up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate_down:
	migrate -path migrations -database $(DB_URL) -verbose down

migrate_fix:
	@echo "Version number:"
	@read -r version_num; \
	if [ -z "$$version_num" ]; then \
		echo "Version number is required"; \
		exit 1; \
	fi; \
	migrate -path migrations -database $(DB_URL) -verbose force "$$version_num"

migrate_go:
	@echo "goto number:"
	@read -r goto_num; \
	if [ -z "$$goto_num" ]; then \
		echo "goto number is required"; \
		exit 1; \
	fi; \
	migrate -path migrations -database $(DB_URL) -verbose goto "$$version_num"

proto-gen:
	./scripts/gen_proto.sh ${CURRENT_DIR}

swag-gen:
	swag init -g ./api/routers.go -o ./api/docs

run:
	go run ./cmd/main.go

commit-all:
	git add .
	git commit -m "$(m)"