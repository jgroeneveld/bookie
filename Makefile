default: install

package = github.com/jgroeneveld/bookie

.PHONY: default install

install:
	go install $(package)

test:
	go test $(package)/...

format:
	goimports -w .

run:
	PORT=3000 DATABASE_URL=postgres://jgroeneveld@localhost/bookie?sslmode=disable go run main.go
