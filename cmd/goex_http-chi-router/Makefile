BIN := gtg-http-chi-router
CONTAINER := gossetx/gtg-http-chi-router


.PHONY: all clean container

all: $(BIN)

container: $(BIN)
	docker build -t $(CONTAINER) .

$(BIN): main.go
	go build -ldflags '-s -w'
	upx $(BIN)

clean:
	rm -f $(BIN)
