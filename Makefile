.PHONY: mocks
mocks:
	sleep 1 && rm -rfd mocks && mockery

swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/main.go

run:  swag
	go run cmd/main.go serve

build:  swag
	go build -o tmp/main cmd/main.go

build-win:  swag
	go build -o tmp/main.exe cmd/main.go
		
docker-run:
	docker run -d -p 8080:8080 \
	--name example \
	example:latest

air:
	air -c .air.toml serve

air-win:
	air -c .air.win.toml serve

test-cov:
	go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html

migrate-up:
	go run cmd/main.go migrate up

migrate-down:
	go run cmd/main.go migrate down

migrate-new:
	go run cmd/main.go migrate fresh