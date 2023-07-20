DST := http2stream

.PHONY: all clean

all: $(DST)

$(DST): main.go
	go build -ldflags '-s -w'
	upx $(DST)

clean:
	rm -f $(DST)
