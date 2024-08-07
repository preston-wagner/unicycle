all: test vet vuln

benchmark:
	go test ./... -bench=.

coverage:
	go test ./... -count=1 -coverprofile=coverage.txt
	go tool cover -html=coverage.txt

test:
	go test ./...

staticcheck:
	staticcheck -f stylish ./...

vet:
	go vet ./...

vuln:
	govulncheck ./...
