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
