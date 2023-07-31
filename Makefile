gocmd_test := go test ./... -cover
flag_nocache := -count 1
flag_integration := -tags integration

outdir := ./out


.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	$(gocmd_test)

.PHONY: test-unit-nocache
test-unit-nocache:
	$(gocmd_test) $(flag_nocache)

.PHONY: test-integration
test-integration:
	$(gocmd_test) $(flag_integration)

.PHONY: test-integration-nocache
test-integration-nocache:
	$(gocmd_test) $(flag_integration) $(flag_nocache)

.PHONY: all
all:
	CGO_ENABLED=0 go build -ldflags '-s -w' -o $(outdir)/ ./cmd/...

.PHONY: install
install:
	go install ./cmd/...

.PHONY: clean
clean:
	rm -rf $(outdir) example.sqlite
