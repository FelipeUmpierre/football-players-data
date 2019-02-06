NO_COLOR=\033[0m
OK_COLOR=\033[32;01m

exec: build
	./players

build:
	@CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X cmd.version=${VERSION}" -o "players" github.com/FelipeUmpierre/onefootball-test/cmd

test: format vet
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)"
	@go test -v -cover -race -covermode=atomic ./...

format:
	@gofmt -l -s cmd | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

vet:
	@echo "$(OK_COLOR)==> checking code correctness with 'go vet' tool$(NO_COLOR)"
	@go vet ./...
