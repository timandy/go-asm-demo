SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
go build -gcflags="-S -l" .
