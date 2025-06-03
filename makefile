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