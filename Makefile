.PHONY: qlub-calculator-api

unit-test-with-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
unit-test:
	go test ./...
