all: test staticcheck vet vuln

benchmark:
	go test ./... -bench=.

coverage:
	go test ./... -count=1 -coverprofile=coverage.txt
	go tool cover -html=coverage.txt

test:
	go test ./...

staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -f stylish ./...

vet:
	go vet ./...

vuln:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...
