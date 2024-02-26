.PHONY: run build-docker run-docker test

run:
	go mod tidy && go run .

build-docker:
	docker build --build-arg API_PORT=$(API_PORT) -t liteapi .

run-docker:
	docker run -p $(API_PORT):$(API_PORT) liteapi
test:
	go test -v ./...