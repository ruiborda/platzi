# Testing

## run tests

```bash
go test -v
```

## go run all tests

```bash
go test -v ./...
```

## go test coverage

```bash
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out

go test -cpuprofile=cpu.out
go tool pprof cpu.out
list FunctionName
web
```