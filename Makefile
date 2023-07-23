outdir := ./out
bin_prefix := goex-

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	go test ./...

.PHONY: test-unit-nocache
test-unit-nocache:
	go test ./... -count 1

.PHONY: test-integration
test-integration:
	go test ./... -tags integration

.PHONY: test-integration-nocache
test-integration-nocache:
	go test ./... -tags integration -count 1

.PHONY: all
all:
	CGO_ENABLED=0 go build -ldflags '-s -w' -o $(outdir)/ ./cmd/...

.PHONY: install
install:
	go install ./cmd/...

.PHONY: uninstall
uninstall:
	rm $(GOBIN)/$(bin_prefix)*

.PHONY: clean
clean:
	rm -rf $(outdir) example.sqlite
