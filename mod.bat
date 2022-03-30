go clean -modcache
go list -m github.com/timandy/routine
del /F /Q go.sum
go mod tidy