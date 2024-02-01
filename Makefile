all: test vet

benchmark:
	go test ./... -bench=.

coverage:
	go test ./... -count=1 -coverprofile=coverage.txt
	go tool cover -html=coverage.txt

test:
	go test ./...

vet:
	go vet ./...
