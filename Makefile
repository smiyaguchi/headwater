.PHONY: all
all: lint build test clean

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	go build -o hw

.PHONY: test
test:
	./hw gen -c 10 -s test/schema_normal.json | column -s, -t
	./hw gen -c 10 -m "HISTORY" -s test/schema_history.json | column -s, -t
	./hw gen -c 10 -m "HISTORY" -s test/schema_history_from_to.json | column -s, -t

.PHONY: clean
clean:
	rm hw
