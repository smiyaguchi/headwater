.PHONY: all
all:
	go build -o hw

.PHONY: clean
clean:
	rm hw 

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	./hw gen -c 10 -s test/schema_normal.json | column -s, -t
	./hw gen -c 10 -m "HISTORY" -s test/schema_history.json | column -s, -t
	./hw gen -c 10 -m "HISTORY" -s test/schema_history_from_to.json | column -s, -t
