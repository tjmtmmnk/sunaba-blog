.PHONY: gen-sql-model
gen-sql-model:
	sqlboiler mysql

.PHONY: run
run:
	reflex -g '**/*.go' -s go run cmd/server.go
