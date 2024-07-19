.PHONY: gqlgen
gqlgen: # gqlgen コード生成
	docker compose run --rm app go run github.com/99designs/gqlgen

.PHONY: sqlboiler
sqlboiler: # SQLBoiler コード生成
	docker compose run --rm app sqlboiler psql

.PHONY: tidy
tidy:
	@docker compose run --rm app go mod tidy
