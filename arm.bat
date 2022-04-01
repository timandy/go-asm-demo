SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=6
go build -o="./gohack.exe" .
