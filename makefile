install: 
		go install github.com/air-verse/air@latest
		go install github.com/go-delve/delve/cmd/dlv@latest
		go install github.com/swaggo/swag/cmd/swag@latest
		go mod tidy

open-api:
		swag init -o resources/public -g main.go
		rm resources/publics/docs.go
		rm resources/public/swagger.yaml
		npx swagger2openapi ./resources/public/swagger.json -o ./resources/public/docs/open-api.yaml
		rm resources/public/swagger.json

format:
	@echo "🎨 Formatage du code Go..."
	go fmt ./...
# === goimports -local "github/CarolineMFVL/nls-messaging-go" -w . ===

APP_NAME=backend
PORT=4000

# === Commandes pour la base PostgreSQL ===

seed:
	@echo "Seeding database..."
	SEED_DB=1 go run main.go

reset:
	@echo "Resetting database..."
	go run tools/reset.go

# === Lancement de l'application Go ===

run:
	@echo "Lancement de $(APP_NAME) sur :$(PORT)"
	go run main.go

build:
	go build -o $(APP_NAME)

# === Tests ===

test-coverage:
	@echo "Exécution des tests avec couverture..."
	go test -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Couverture générée dans coverage.html"

test:
	@echo "Exécution des tests..."
	go test ./tests/...

# === Lint ===

lint:
	@echo "🔍 Linting du code Go..."
	golangci-lint run .

# === Docker ===

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

docker-rebuild:
	docker-compose down && docker-compose up --build

# === Aide ===

help:
	@echo "Commandes disponibles :"
	@echo "  make run            - Lancer l'application en local"
	@echo "  make seed           - Injecter des utilisateurs de test"
	@echo "  make reset          - Réinitialiser la base (optionnel)"
	@echo "  make test           - Lancer les tests"
	@echo "  make build          - Compiler l'app"
	@echo "  make docker-up      - Lancer avec Docker"
	@echo "  make docker-down    - Stopper Docker"
	@echo "  make docker-rebuild - Rebuild complet Docker"