.PHONY: run
run:
	go run main.go

.PHONY: generate
generate:
	sqlc generate
