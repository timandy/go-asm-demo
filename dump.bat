
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
del result.s /S /Q
pause
go build -gcflags="-N -l"
go tool objdump -S  .\routine-demo > result.s
