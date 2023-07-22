outdir := ./out
bin_prefix := goex-

.PHONY: test
test:
	go test ./...

.PHONY: test-integration
test-integration:
	go test ./... -tags integration

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
