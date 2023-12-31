.PHONY: build gen-swagger docker-up gen-mock test

LOCAL_BIN:=$(CURDIR)/bin

$(LOCAL_DIR):
	@mkdir -p $@

build:
	go build -o $(LOCAL_BIN)/avito-backend cmd/avito-backend/main.go

gen-swagger:
	GOBIN=$(LOCAL_BIN) go get github.com/swaggo/swag/cmd/swag
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag
	$(LOCAL_BIN)/swag init -g ./cmd/avito-backend/main.go


docker-up:
	docker-compose up --build

gen-mock:
	GOBIN=$(LOCAL_BIN) go get github.com/vektra/mockery/v2
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2
	$(LOCAL_BIN)/mockery --all

test:
	go test ./...

test-ci:
	go test ./internal/... ./pkg/...
