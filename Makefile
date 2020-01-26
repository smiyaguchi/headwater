.PHONY: all
all:
	./build.sh

.PHONY: clean
clean:
	rm hw schema.json

.PHONY: test
test:
	cp test/schema.json ./
	./hw gen -c 10 | column -s, -t
	./hw gen -c 10 -m "HISTORY" -s test/schema_history.json | column -s, -t
	./hw gen -c 10 -m "HISTORY" -s test/schema_history_from_to.json | column -s, -t
