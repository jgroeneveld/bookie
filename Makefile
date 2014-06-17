default: install

package = github.com/jgroeneveld/wtrack

.PHONY: default install

install:
	go install $(package)

test:
	go test $(package)/...

format:
	goimports -w .

run:
	DATABASE_URL=postgres://jgroeneveld@localhost/bookie?sslmode=disable go run main.go
