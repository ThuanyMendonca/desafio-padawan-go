run:
	go run ./cmd/server/main.go

cover-test:
	go test -coverprofile cover.out ${path}/...
	go tool cover -html=cover.out