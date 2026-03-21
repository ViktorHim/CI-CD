set -e
echo "Running tests"
cd src
go test -v ./...
echo "Tests passed"