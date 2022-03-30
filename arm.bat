SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
go build -gcflags="-N -S -l" .
