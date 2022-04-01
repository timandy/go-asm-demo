SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm

del result.s /S /Q
pause
go build -gcflags="-N -l" -o ./gohack.exe .
go tool objdump -S ./gohack.exe > result.s
