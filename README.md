# danmuku
danmu api server.

# environment
- go: v1.23.2
- gin: v1.10.0


# build

```
go build -ldflags="-s -w" -trimpath -o danmuku .
```

# run

```
go run main.go
```