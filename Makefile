.PHONY: run

run:
	go run cmd/service/main.go

healthz:
	curl http://localhost:8080/healthz