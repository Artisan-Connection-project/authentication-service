# Makefile

# Default target
all: create-dirs

# Run the create_dirs.sh script with the provided path
create-dirs:
	@read -p "Enter the path: " path; \
	if [ -z "$$path" ]; then \
		echo "Path is required"; \
		exit 1; \
	fi; \
	./scripts/create_dirs.sh "$$path"
