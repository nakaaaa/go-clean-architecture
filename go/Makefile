.PHONY: mod/download mod/tidy

mod/download:
	go mod download
	
mod/tidy:
	go mod tidy
	
# docker
.PHONY: compose/down compose/build compose/up

compose/down:
	docker-compose down

compose/build:
	docker-compose build

compose/up:
	docker-compose up -d

.PHONY: test

test:
	-rm -r ./reports
	mkdir -p ./reports
	env `grep -v -e '^\s*#' -e '^\s*$$' .env.local` go test -v -shuffle=on -coverprofile=./reports/coverage.out -cover ./internal/... ./pkg/...