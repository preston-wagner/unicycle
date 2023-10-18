all: test vet

coverage:
	go test ./... -coverprofile=coverage.txt

test:
	go test ./...

vet:
	go vet ./...
