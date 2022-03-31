SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -gcflags="-N -l" .
