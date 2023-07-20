You will want to compile as follows to ensure the resulting binary is STATICALLY
linked, NOT dynamically linked against libc:

```
CGO_ENABLED=0 go build
```
