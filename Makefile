define setup_env
    $(eval include .env)
    $(eval export)
endef

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orchestrator.proto" \
		--go_out=services/orchestrator --go_opt=paths=source_relative \
  	--go-grpc_out=services/orchestrator --go-grpc_opt=paths=source_relative

generate-docs:
	cd ./cmd/rest && swag init --parseDependency --parseInternal --parseDepth 1 && rm -rf ./controllers/docs && mv ./docs ./controllers

run:
	docker compose up --build

run-upload:
	$(call setup_env)
	go run ./cmd/uploader/

test:
	make generate-docs
	docker compose up db -d
	$(call setup_env)
	go test -v ./tests/...

build:
	mkdir bin

	CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -linkmode external -extldflags "-fno-PIC -static"' -o ./bin/rest_build ./cmd/rest
	CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s -linkmode external -extldflags "-fno-PIC -static"' -o ./bin/orchestrator_build ./cmd/orchestrator

