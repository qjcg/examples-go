outdir := ./out
bin_prefix := goex-

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	go test ./... -cover

.PHONY: test-unit-nocache
test-unit-nocache:
	go test ./... -count 1 -cover

.PHONY: test-integration
test-integration:
	go test ./... -tags integration -cover

.PHONY: test-integration-nocache
test-integration-nocache:
	go test ./... -tags integration -count 1 -cover

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
