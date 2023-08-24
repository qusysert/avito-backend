LOCAL_BIN:=$(CURDIR)/bin

$(LOCAL_DIR):
	@mkdir -p $@

build:
	go build -o $(LOCAL_BIN)/avito-backend cmd/avito-backend/main.go

gen-swagger:
	GOBIN=$(LOCAL_BIN) go get github.com/swaggo/swag/cmd/swag
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag
	$(LOCAL_BIN)/swag init -g ./cmd/avito-backend/main.go

lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	$(LOCAL_BIN)/golangci-lint run ./...