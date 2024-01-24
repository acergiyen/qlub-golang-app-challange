.PHONY: qlub-calculator-api

unit-test-with-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
unit-test:
	go test ./...

build:
	docker build -t qlub-calculator-api:latest .  
run:
	docker run -p 8080:8080 qlub-calculator-api:latest
