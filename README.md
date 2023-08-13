## Go simple Testing

## Run Test
```
go test -v ./...
```

## Run Test Coverage
```
go test -v ./... -coverprofile=cover.out && go tool cover -html=cover.out
```