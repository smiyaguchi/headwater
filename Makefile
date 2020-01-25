.PHONY: all
all:
	./build.sh

.PHONY: clean
clean:
	rm hw schema.json

.PHONY: test
test:
	./hw gen -c 10 | column -s, -t
	./hw gen -c 10 -m "HISTORY" | column -s, -t
