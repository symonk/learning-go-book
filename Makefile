.PHONY: test


lint:
	go fmt ./...
	go vet ./...

test:
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html
	open cover.html