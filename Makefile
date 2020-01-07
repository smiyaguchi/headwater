.PHONY: all
all:
	./build.sh

.PHONY: clean
clean:
	rm hw schema.json

.PHONY: test
test:
	./hw gen -c 5
