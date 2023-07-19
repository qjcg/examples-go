bin := gtg-example

.PHONY: build
build:
	CGO_ENABLED=0 go build

.PHONY: test
test:
	hey -c 200 -n 1000 http://localhost:9999/

.PHONY: clean
clean:
	rm $(bin)
