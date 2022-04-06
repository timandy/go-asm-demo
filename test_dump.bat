SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64

del result.s /S /Q
pause
go build -gcflags="-N -l" -o ./gohack.exe
go tool objdump -S ./gohack.exe > result.s
