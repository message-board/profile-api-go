.PHONY: swagger
swagger:
	swag init -g internal/interfaces/api/server.go -o api/openapi