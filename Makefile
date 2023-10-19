all: test vet

coverage:
	go test ./... -coverprofile=coverage.txt
	go tool cover -html=coverage.txt

test:
	go test ./...

vet:
	go vet ./...
