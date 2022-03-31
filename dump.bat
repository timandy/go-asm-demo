
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
del result.s /S /Q
pause
go build -gcflags="-N -l" -o .\main.exe
go tool objdump -S  .\main.exe > result.s
