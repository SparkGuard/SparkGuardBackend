define setup_env
    $(eval include .env)
    $(eval export)
endef

generate-docs:
	cd rest && swag init --parseDependency --parseInternal --parseDepth 1 && rm -rf ./controllers/docs && mv ./docs ./controllers

run:
	docker compose up --build

test:
	make generate-docs
	docker compose up db -d
	$(call setup_env)
	cd rest && go test -v ./tests/...
